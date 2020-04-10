// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	status "github.com/operator-framework/operator-sdk/pkg/status"
	v1 "k8s.io/api/batch/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarketplaceConfig) DeepCopyInto(out *MarketplaceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarketplaceConfig.
func (in *MarketplaceConfig) DeepCopy() *MarketplaceConfig {
	if in == nil {
		return nil
	}
	out := new(MarketplaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MarketplaceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarketplaceConfigList) DeepCopyInto(out *MarketplaceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MarketplaceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarketplaceConfigList.
func (in *MarketplaceConfigList) DeepCopy() *MarketplaceConfigList {
	if in == nil {
		return nil
	}
	out := new(MarketplaceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MarketplaceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarketplaceConfigSpec) DeepCopyInto(out *MarketplaceConfigSpec) {
	*out = *in
	if in.DeploySecretName != nil {
		in, out := &in.DeploySecretName, &out.DeploySecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarketplaceConfigSpec.
func (in *MarketplaceConfigSpec) DeepCopy() *MarketplaceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MarketplaceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MarketplaceConfigStatus) DeepCopyInto(out *MarketplaceConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MarketplaceConfigStatus.
func (in *MarketplaceConfigStatus) DeepCopy() *MarketplaceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(MarketplaceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterBase) DeepCopyInto(out *MeterBase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterBase.
func (in *MeterBase) DeepCopy() *MeterBase {
	if in == nil {
		return nil
	}
	out := new(MeterBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeterBase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterBaseList) DeepCopyInto(out *MeterBaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MeterBase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterBaseList.
func (in *MeterBaseList) DeepCopy() *MeterBaseList {
	if in == nil {
		return nil
	}
	out := new(MeterBaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeterBaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterBaseSpec) DeepCopyInto(out *MeterBaseSpec) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterBaseSpec.
func (in *MeterBaseSpec) DeepCopy() *MeterBaseSpec {
	if in == nil {
		return nil
	}
	out := new(MeterBaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterBaseStatus) DeepCopyInto(out *MeterBaseStatus) {
	*out = *in
	if in.PrometheusNodes != nil {
		in, out := &in.PrometheusNodes, &out.PrometheusNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterBaseStatus.
func (in *MeterBaseStatus) DeepCopy() *MeterBaseStatus {
	if in == nil {
		return nil
	}
	out := new(MeterBaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metering) DeepCopyInto(out *Metering) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metering.
func (in *Metering) DeepCopy() *Metering {
	if in == nil {
		return nil
	}
	out := new(Metering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Metering) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringList) DeepCopyInto(out *MeteringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Metering, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringList.
func (in *MeteringList) DeepCopy() *MeteringList {
	if in == nil {
		return nil
	}
	out := new(MeteringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeteringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringSpec) DeepCopyInto(out *MeteringSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringSpec.
func (in *MeteringSpec) DeepCopy() *MeteringSpec {
	if in == nil {
		return nil
	}
	out := new(MeteringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeteringStatus) DeepCopyInto(out *MeteringStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeteringStatus.
func (in *MeteringStatus) DeepCopy() *MeteringStatus {
	if in == nil {
		return nil
	}
	out := new(MeteringStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSpec) DeepCopyInto(out *PrometheusSpec) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Storage.DeepCopyInto(&out.Storage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSpec.
func (in *PrometheusSpec) DeepCopy() *PrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RazeeDeployment) DeepCopyInto(out *RazeeDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RazeeDeployment.
func (in *RazeeDeployment) DeepCopy() *RazeeDeployment {
	if in == nil {
		return nil
	}
	out := new(RazeeDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RazeeDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RazeeDeploymentList) DeepCopyInto(out *RazeeDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RazeeDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RazeeDeploymentList.
func (in *RazeeDeploymentList) DeepCopy() *RazeeDeploymentList {
	if in == nil {
		return nil
	}
	out := new(RazeeDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RazeeDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RazeeDeploymentSpec) DeepCopyInto(out *RazeeDeploymentSpec) {
	*out = *in
	if in.DeploySecretName != nil {
		in, out := &in.DeploySecretName, &out.DeploySecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RazeeDeploymentSpec.
func (in *RazeeDeploymentSpec) DeepCopy() *RazeeDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(RazeeDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RazeeDeploymentStatus) DeepCopyInto(out *RazeeDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new(v1.JobCondition)
		(*in).DeepCopyInto(*out)
	}
	in.JobState.DeepCopyInto(&out.JobState)
	if in.MissingValuesFromSecret != nil {
		in, out := &in.MissingValuesFromSecret, &out.MissingValuesFromSecret
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.LocalSecretVarsPopulated != nil {
		in, out := &in.LocalSecretVarsPopulated, &out.LocalSecretVarsPopulated
		*out = new(bool)
		**out = **in
	}
	if in.RazeePrerequisitesCreated != nil {
		in, out := &in.RazeePrerequisitesCreated, &out.RazeePrerequisitesCreated
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.RedHatMarketplaceSecretFound != nil {
		in, out := &in.RedHatMarketplaceSecretFound, &out.RedHatMarketplaceSecretFound
		*out = new(bool)
		**out = **in
	}
	if in.RazeeJobInstall != nil {
		in, out := &in.RazeeJobInstall, &out.RazeeJobInstall
		*out = new(RazeeJobInstallStruct)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RazeeDeploymentStatus.
func (in *RazeeDeploymentStatus) DeepCopy() *RazeeDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(RazeeDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RazeeJobInstallStruct) DeepCopyInto(out *RazeeJobInstallStruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RazeeJobInstallStruct.
func (in *RazeeJobInstallStruct) DeepCopy() *RazeeJobInstallStruct {
	if in == nil {
		return nil
	}
	out := new(RazeeJobInstallStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(string)
		**out = **in
	}
	out.Size = in.Size.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
