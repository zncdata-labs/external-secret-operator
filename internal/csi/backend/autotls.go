package backend

import (
	"context"
	"crypto/x509"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/zncdatadev/secret-operator/pkg/pod_info"
	"github.com/zncdatadev/secret-operator/pkg/util"
	"github.com/zncdatadev/secret-operator/pkg/volume"
	"sigs.k8s.io/controller-runtime/pkg/client"

	secretsv1alpha1 "github.com/zncdatadev/secret-operator/api/v1alpha1"
	"github.com/zncdatadev/secret-operator/internal/csi/backend/ca"
)

const (
	KeystoreP12FileName   = "keystore.p12"
	TruststoreP12FileName = "truststore.p12"
	PEMTlsCertFileName    = "tls.crt"
	PEMTlsKeyFileName     = "tls.key"
	PEMCaCertFileName     = "ca.crt"
)

const (
	DefaultCertLifeTime time.Duration = 24 * 7 * time.Hour
	DefaultCertJitter   float64       = 0.2
	DefaultCertBuffer   time.Duration = 8 * time.Hour
)

var _ IBackend = &AutoTlsBackend{}

type AutoTlsBackend struct {
	client                 client.Client
	podInfo                *pod_info.PodInfo
	volumeContext          *volume.SecretVolumeContext
	maxCertificateLifeTime time.Duration

	ca          *secretsv1alpha1.CASpec
	certManager *ca.CertificateManager
}

func NewAutoTlsBackend(config *BackendConfig) (IBackend, error) {
	autotls := config.SecretClass.Spec.Backend.AutoTls
	maxCertificateLifeTime, err := time.ParseDuration(autotls.MaxCertificateLifeTime)
	if err != nil {
		return nil, err
	}

	caCertificateLifeTime, err := time.ParseDuration(autotls.CA.CACertificateLifeTime)
	if err != nil {
		return nil, err
	}

	return &AutoTlsBackend{
		client:                 config.Client,
		podInfo:                config.PodInfo,
		volumeContext:          config.VolumeContext,
		maxCertificateLifeTime: maxCertificateLifeTime,
		ca:                     autotls.CA,

		certManager: ca.NewCertificateManager(
			config.Client,
			caCertificateLifeTime,
			autotls.CA.AutoGenerated,
			autotls.CA.Secret.Name,
			autotls.CA.Secret.Namespace,
		),
	}, nil
}

// use AutoTlsCertLifetime and AutoTlsCertJitterFactor to calculate the certificate lifetime
func (a *AutoTlsBackend) getCertLife() (time.Duration, error) {
	now := time.Now()

	certLife := a.volumeContext.AutoTlsCertLifetime
	if certLife == 0 {
		logger.Info("Certificate lifetime is not set, using default certificate lifetime", "defaultCertLifeTime", DefaultCertLifeTime)
		certLife = DefaultCertLifeTime
	}
	restarterBuffer := a.volumeContext.AutoTlsCertRestartBuffer
	if restarterBuffer == 0 {
		logger.Info("Certificate restart buffer is not set, using default certificate restart buffer", "defaultCertBuffer", DefaultCertBuffer)
		restarterBuffer = DefaultCertBuffer
	}

	if certLife > a.maxCertificateLifeTime {
		logger.Info("Certificate lifetime is greater than the maximum certificate lifetime, using the maximum certificate lifetime",
			"certLife", certLife,
			"maxCertificateLifeTime", a.maxCertificateLifeTime,
		)
		certLife = a.maxCertificateLifeTime
	}

	jitterFactor := a.volumeContext.AutoTlsCertJitterFactor

	jitterFactorAllowedRange := 0.0 < jitterFactor && jitterFactor < 1.0
	if !jitterFactorAllowedRange {
		logger.Info("Invalid jitter factor, using default value", "jitterFactor", jitterFactor)
		jitterFactor = DefaultCertJitter
	}

	randomJitterFactor := rand.Float64() * jitterFactor
	jitterLife := time.Duration(float64(certLife) * jitterFactor)
	jitteredCertLife := certLife - jitterLife

	logger.Info("Jittered certificate lifetime",
		"certLife", certLife,
		"jitteredCertLife", jitteredCertLife,
		"jitterLife", jitterLife,
		"jitterFactor", jitterFactor,
		"randomJitterFactor", randomJitterFactor,
	)

	notAfter := now.Add(jitteredCertLife)
	podExpires := notAfter.Add(-restarterBuffer)
	if podExpires.Before(now) {
		return 0, fmt.Errorf("certificate lifetime is too short, pod will restart before certificate expiration. "+
			"'Now': %v, 'Expires': %v, 'Restart': %v", now, notAfter, podExpires,
		)
	}

	return certLife, nil
}

func (a *AutoTlsBackend) certificateFormat() volume.SecretFormat {
	return a.volumeContext.Format
}

// Convert the certificate to the format required by the volume
// If the format is PKCS12, the certificate will be converted to PKCS12 format,
// otherwise it will be converted to PEM format.
func (a *AutoTlsBackend) certificateConvert(cert *ca.Certificate) (map[string]string, error) {
	format := a.certificateFormat()

	trustAnchors := a.certManager.GetTrustAnchors()

	if format == volume.SecretFormatTLSP12 {
		logger.Info("Converting certificate to PKCS12 format")
		password := a.volumeContext.TlsPKCS12Password

		caCerts := make([]*x509.Certificate, 0, len(trustAnchors))
		for _, caCert := range trustAnchors {
			caCerts = append(caCerts, caCert.Certificate)
		}

		truststore, err := cert.TrustStoreP12(password, caCerts)
		if err != nil {
			return nil, err
		}
		keyStore, err := cert.KeyStoreP12(password, caCerts)
		if err != nil {
			return nil, err
		}
		return map[string]string{
			KeystoreP12FileName:   string(keyStore),
			TruststoreP12FileName: string(truststore),
		}, nil
	}

	pemCACerts := make([]string, 0, len(trustAnchors))

	for _, caCert := range trustAnchors {
		pemCACerts = append(pemCACerts, string(caCert.CertificatePEM()))
	}

	logger.Info("Converting certificate to PEM format")
	return map[string]string{
		PEMTlsCertFileName: string(cert.CertificatePEM()),
		PEMTlsKeyFileName:  string(cert.PrivateKeyPEM()),
		PEMCaCertFileName:  strings.Join(pemCACerts, "\n"),
	}, nil
}

func (k *AutoTlsBackend) GetQualifiedNodeNames(ctx context.Context) ([]string, error) {
	// Default implementation, return nil
	return nil, nil
}

func (a *AutoTlsBackend) GetSecretData(ctx context.Context) (*util.SecretContent, error) {
	certificateAuthority, err := a.getCertificateAuthority(ctx)
	if err != nil {
		return nil, err
	}

	addresses, err := a.getAddresses(ctx)
	if err != nil {
		return nil, err
	}

	certLife, err := a.getCertLife()
	if err != nil {
		return nil, err
	}

	notAfter := time.Now().Add(certLife)

	cert, err := certificateAuthority.SignServerCertificate(addresses, notAfter)

	if err != nil {
		return nil, err
	}

	logger.Info("Signed certificate", "notAfter", notAfter, "addresses", addresses, "certLife", certLife, "certSerialNumber", cert.SerialNumber())

	data, err := a.certificateConvert(cert)
	if err != nil {
		return nil, err
	}

	expiresTime := notAfter

	return &util.SecretContent{
		Data:        data,
		ExpiresTime: &expiresTime,
	}, nil
}

func (a *AutoTlsBackend) getAddresses(ctx context.Context) ([]pod_info.Address, error) {
	return a.podInfo.GetScopedAddresses(ctx)
}

// Get the certificate authority before the expiration time, the expiration time from the secret class
func (a *AutoTlsBackend) getCertificateAuthority(ctx context.Context) (*ca.CertificateAuthority, error) {
	atAfter := time.Now().Add(a.maxCertificateLifeTime) // server cert lifetime in secret class configed
	certificateAuthority, err := a.certManager.GetCertificateAuthority(ctx, atAfter)
	if err != nil {
		return nil, err
	}

	return certificateAuthority, nil
}
