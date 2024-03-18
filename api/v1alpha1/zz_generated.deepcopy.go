//go:build !ignore_autogenerated

/*
Copyright 2024 zncdata-labs.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthTlsSpec) DeepCopyInto(out *AuthTlsSpec) {
	*out = *in
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(CASpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthTlsSpec.
func (in *AuthTlsSpec) DeepCopy() *AuthTlsSpec {
	if in == nil {
		return nil
	}
	out := new(AuthTlsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendSpec) DeepCopyInto(out *BackendSpec) {
	*out = *in
	if in.AuthTls != nil {
		in, out := &in.AuthTls, &out.AuthTls
		*out = new(AuthTlsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.K8sSearch != nil {
		in, out := &in.K8sSearch, &out.K8sSearch
		*out = new(K8sSearchSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Kerberos != nil {
		in, out := &in.Kerberos, &out.Kerberos
		*out = new(KerberosSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendSpec.
func (in *BackendSpec) DeepCopy() *BackendSpec {
	if in == nil {
		return nil
	}
	out := new(BackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASpec) DeepCopyInto(out *CASpec) {
	*out = *in
	if in.ExistingSecret != nil {
		in, out := &in.ExistingSecret, &out.ExistingSecret
		*out = new(ExistingSecretSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASpec.
func (in *CASpec) DeepCopy() *CASpec {
	if in == nil {
		return nil
	}
	out := new(CASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIPluginSpec) DeepCopyInto(out *CSIPluginSpec) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIPluginSpec.
func (in *CSIPluginSpec) DeepCopy() *CSIPluginSpec {
	if in == nil {
		return nil
	}
	out := new(CSIPluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIProvisionerSpec) DeepCopyInto(out *CSIProvisionerSpec) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIProvisionerSpec.
func (in *CSIProvisionerSpec) DeepCopy() *CSIProvisionerSpec {
	if in == nil {
		return nil
	}
	out := new(CSIProvisionerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExistingSecretSpec) DeepCopyInto(out *ExistingSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExistingSecretSpec.
func (in *ExistingSecretSpec) DeepCopy() *ExistingSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ExistingSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sSearchSpec) DeepCopyInto(out *K8sSearchSpec) {
	*out = *in
	if in.SearchNamespace != nil {
		in, out := &in.SearchNamespace, &out.SearchNamespace
		*out = new(SearchNamespaceSpec)
		**out = **in
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sSearchSpec.
func (in *K8sSearchSpec) DeepCopy() *K8sSearchSpec {
	if in == nil {
		return nil
	}
	out := new(K8sSearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KerberosSpec) DeepCopyInto(out *KerberosSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KerberosSpec.
func (in *KerberosSpec) DeepCopy() *KerberosSpec {
	if in == nil {
		return nil
	}
	out := new(KerberosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LivenessProbeSpec) DeepCopyInto(out *LivenessProbeSpec) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LivenessProbeSpec.
func (in *LivenessProbeSpec) DeepCopy() *LivenessProbeSpec {
	if in == nil {
		return nil
	}
	out := new(LivenessProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSpec) DeepCopyInto(out *LoggingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSpec.
func (in *LoggingSpec) DeepCopy() *LoggingSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDriverRegisterSpec) DeepCopyInto(out *NodeDriverRegisterSpec) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDriverRegisterSpec.
func (in *NodeDriverRegisterSpec) DeepCopy() *NodeDriverRegisterSpec {
	if in == nil {
		return nil
	}
	out := new(NodeDriverRegisterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchNamespaceSpec) DeepCopyInto(out *SearchNamespaceSpec) {
	*out = *in
	out.Pod = in.Pod
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchNamespaceSpec.
func (in *SearchNamespaceSpec) DeepCopy() *SearchNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(SearchNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCSI) DeepCopyInto(out *SecretCSI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCSI.
func (in *SecretCSI) DeepCopy() *SecretCSI {
	if in == nil {
		return nil
	}
	out := new(SecretCSI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretCSI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCSIList) DeepCopyInto(out *SecretCSIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretCSI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCSIList.
func (in *SecretCSIList) DeepCopy() *SecretCSIList {
	if in == nil {
		return nil
	}
	out := new(SecretCSIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretCSIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCSISpec) DeepCopyInto(out *SecretCSISpec) {
	*out = *in
	if in.CSIPlugin != nil {
		in, out := &in.CSIPlugin, &out.CSIPlugin
		*out = new(CSIPluginSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeDriverRegister != nil {
		in, out := &in.NodeDriverRegister, &out.NodeDriverRegister
		*out = new(NodeDriverRegisterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CSIProvisioner != nil {
		in, out := &in.CSIProvisioner, &out.CSIProvisioner
		*out = new(CSIProvisionerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(LivenessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCSISpec.
func (in *SecretCSISpec) DeepCopy() *SecretCSISpec {
	if in == nil {
		return nil
	}
	out := new(SecretCSISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCSIStatus) DeepCopyInto(out *SecretCSIStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCSIStatus.
func (in *SecretCSIStatus) DeepCopy() *SecretCSIStatus {
	if in == nil {
		return nil
	}
	out := new(SecretCSIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretClass) DeepCopyInto(out *SecretClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretClass.
func (in *SecretClass) DeepCopy() *SecretClass {
	if in == nil {
		return nil
	}
	out := new(SecretClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretClassList) DeepCopyInto(out *SecretClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretClassList.
func (in *SecretClassList) DeepCopy() *SecretClassList {
	if in == nil {
		return nil
	}
	out := new(SecretClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretClassSpec) DeepCopyInto(out *SecretClassSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretClassSpec.
func (in *SecretClassSpec) DeepCopy() *SecretClassSpec {
	if in == nil {
		return nil
	}
	out := new(SecretClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretClassStatus) DeepCopyInto(out *SecretClassStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretClassStatus.
func (in *SecretClassStatus) DeepCopy() *SecretClassStatus {
	if in == nil {
		return nil
	}
	out := new(SecretClassStatus)
	in.DeepCopyInto(out)
	return out
}
