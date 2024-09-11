package volume

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zncdatadev/operator-go/pkg/constants"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	logger = ctrl.Log.WithName("volume")
)

const (
	KerberosServiceNamesSplitter string = ","
)

type SecretFormat string

const (
	SecretFormatTLSPEM   SecretFormat = "tls-pem"
	SecretFormatTLSP12   SecretFormat = "tls-p12"
	SecretFormatKerberos SecretFormat = "kerberos"
)

const (
	// kubernetes and sig defained annotations for PVC
	CSIStoragePodName                       string = "csi.storage.k8s.io/pod.name"
	CSIStoragePodNamespace                  string = "csi.storage.k8s.io/pod.namespace"
	CSIStoragePodUid                        string = "csi.storage.k8s.io/pod.uid"
	CSIStorageServiceAccountName            string = "csi.storage.k8s.io/serviceAccount.name"
	CSIStorageEphemeral                     string = "csi.storage.k8s.io/ephemeral"
	StorageKubernetesCSIProvisionerIdentity string = "storage.kubernetes.io/csiProvisionerIdentity"
	VolumeKubernetesStorageProvisioner      string = "volume.kubernetes.io/storage-provisioner"
	// https://kubernetes.io/docs/reference/labels-annotations-taints/#volume-beta-kubernetes-io-storage-provisioner-deprecated
	DeprecatedVolumeKubernetesStorageProvisioner string = "volume.beta.kubernetes.io/storage-provisioner"
)

type SecretVolumeSelector struct {
	// Default values for volume context
	Pod                      string `json:"csi.storage.k8s.io/pod.name"`
	PodNamespace             string `json:"csi.storage.k8s.io/pod.namespace"`
	PodUID                   string `json:"csi.storage.k8s.io/pod.uid"`
	ServiceAccountName       string `json:"csi.storage.k8s.io/serviceAccount.name"`
	Ephemeral                string `json:"csi.storage.k8s.io/ephemeral"`
	CSIProvisionerIdIdentiry string `json:"storage.kubernetes.io/csiProvisionerIdentity"`
	Provisioner              string `json:"volume.kubernetes.io/storage-provisioner"`

	Class  string       `json:"secrets.zncdata.dev/class"`
	Scope  SecretScope  `json:"secrets.zncdata.dev/scope"`
	Format SecretFormat `json:"secrets.zncdata.dev/format"`

	TlsPKCS12Password       string        `json:"secrets.zncdata.dev/tlsPKCS12Password"`
	AutoTlsCertLifetime     time.Duration `json:"secrets.zncdata.dev/autoTlsCertLifetime"`
	AutoTlsCertJitterFactor float64       `json:"secrets.zncdata.dev/autoTlsCertJitterFactor"`

	KerberosServiceNames []string `json:"secrets.zncdata.dev/kerberosServiceNames"`
}

type ListScope string

const (
	ScopePod            ListScope = "pod"
	ScopeNode           ListScope = "node"
	ScopeService        string    = "service"
	ScopeListenerVolume string    = "listener-volume"
)

type SecretScope struct {
	Pod  ListScope `json:"pod"`
	Node ListScope `json:"node"`
	// this field is k-v pair, key is service name, value is service type
	Services []string `json:"service"`
	// this field is k-v pair, key is listener volume name, value is listener volume type
	ListenerVolumes []string `json:"listener-volume"`
}

func (v SecretVolumeSelector) ToMap() map[string]string {
	out := make(map[string]string)
	if v.Pod != "" {
		out[CSIStoragePodName] = v.Pod
	}
	if v.PodNamespace != "" {
		out[CSIStoragePodNamespace] = v.PodNamespace
	}
	if v.PodUID != "" {
		out[CSIStoragePodUid] = v.PodUID
	}
	if v.ServiceAccountName != "" {
		out[CSIStorageServiceAccountName] = v.ServiceAccountName
	}
	if v.Ephemeral != "" {
		out[CSIStorageEphemeral] = v.Ephemeral
	}
	if v.Provisioner != "" {
		out[StorageKubernetesCSIProvisionerIdentity] = v.Provisioner
	}
	if v.Class != "" {
		out[constants.AnnotationSecretsClass] = v.Class
	}
	if v.encodeScope() != "" {
		out[constants.AnnotationSecretsScope] = v.encodeScope()
	}
	if v.Format != "" {
		out[constants.AnnotationSecretsFormat] = string(v.Format)
	}
	if len(v.KerberosServiceNames) > 0 {
		out[constants.AnnotationSecretsKerberosServiceNames] = strings.Join(v.KerberosServiceNames, KerberosServiceNamesSplitter)
	}
	if v.TlsPKCS12Password != "" {
		out[constants.AnnotationSecretsPKCS12Password] = v.TlsPKCS12Password
	}
	if v.AutoTlsCertLifetime != 0 {
		out[constants.AnnonationSecretExpirationTimeName] = v.AutoTlsCertLifetime.String()
	}
	if v.AutoTlsCertJitterFactor != 0 {
		out[constants.AnnonationSecretExpirationTimeName] = fmt.Sprintf("%f", v.AutoTlsCertJitterFactor)
	}
	return out
}

func (v SecretVolumeSelector) encodeScope() string {
	var scopes []string
	if v.Scope.Pod != "" && v.Scope.Pod == ScopePod {
		scopes = append(scopes, string(v.Scope.Pod))
	}
	if v.Scope.Node != "" {
		scopes = append(scopes, string(v.Scope.Node))
	}
	if v.Scope.Services != nil {
		for _, services := range v.Scope.Services {
			scopes = append(scopes, fmt.Sprintf("%s=%s", ScopeService, services))
		}
	}
	if v.Scope.ListenerVolumes != nil {
		for _, listenerVolume := range v.Scope.ListenerVolumes {
			scopes = append(scopes, fmt.Sprintf("%s=%s", ScopeListenerVolume, listenerVolume))
		}
	}
	return strings.Join(scopes, ",")
}

func (v SecretVolumeSelector) decodeScope(scope string) SecretScope {
	secretScope := SecretScope{}

	for _, scopes := range strings.Split(scope, ",") {
		kv := strings.Split(scopes, "=")
		switch kv[0] {
		case string(ScopePod):
			secretScope.Pod = ScopePod
		case string(ScopeNode):
			secretScope.Node = ScopeNode
		case ScopeService:
			secretScope.Services = append(secretScope.Services, kv[1])
		case ScopeListenerVolume:
			secretScope.ListenerVolumes = append(secretScope.ListenerVolumes, kv[1])
		default:
			logger.V(0).Info("Unknown scope, skip it", "scope name", kv[0], "scope value", kv[1])
		}
	}
	return secretScope
}

func NewVolumeSelectorFromMap(parameters map[string]string) (*SecretVolumeSelector, error) {
	v := &SecretVolumeSelector{}
	for key, value := range parameters {
		switch key {
		case CSIStoragePodName:
			v.Pod = value
		case CSIStoragePodNamespace:
			v.PodNamespace = value
		case CSIStoragePodUid:
			v.PodUID = value
		case CSIStorageServiceAccountName:
			v.ServiceAccountName = value
		case CSIStorageEphemeral:
			v.Ephemeral = value
		case StorageKubernetesCSIProvisionerIdentity:
			v.CSIProvisionerIdIdentiry = value
		case VolumeKubernetesStorageProvisioner:
			v.Provisioner = value
		case DeprecatedVolumeKubernetesStorageProvisioner:
			logger.V(0).Info("Deprecated key since v1.23, please use new key",
				"key", key,
				"value", value,
				"new key", VolumeKubernetesStorageProvisioner,
				"reference", "https://kubernetes.io/docs/reference/labels-annotations-taints/#volume-beta-kubernetes-io-storage-provisioner-deprecated",
			)
		case constants.AnnotationSecretsClass:
			v.Class = value
		case constants.AnnotationSecretsScope:
			v.Scope = v.decodeScope(value)
		case constants.AnnotationSecretsFormat:
			v.Format = SecretFormat(value)
		case constants.AnnotationSecretsKerberosServiceNames:
			v.KerberosServiceNames = strings.Split(value, KerberosServiceNamesSplitter)
		case constants.AnnotationSecretsPKCS12Password:
			v.TlsPKCS12Password = value
		case constants.AnnotationSecretCertLifeTime:
			d, err := time.ParseDuration(value)
			if err != nil {
				return nil, err
			}
			v.AutoTlsCertLifetime = d
		case constants.AnnotationSecretsCertJitterFactor:
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			v.AutoTlsCertJitterFactor = float64(i)
		default:
			logger.V(0).Info("Unknown key, skip it", "key", key, "value", value)
		}
	}
	return v, nil
}
