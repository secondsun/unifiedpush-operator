// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidVariant) DeepCopyInto(out *AndroidVariant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidVariant.
func (in *AndroidVariant) DeepCopy() *AndroidVariant {
	if in == nil {
		return nil
	}
	out := new(AndroidVariant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AndroidVariant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidVariantList) DeepCopyInto(out *AndroidVariantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AndroidVariant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidVariantList.
func (in *AndroidVariantList) DeepCopy() *AndroidVariantList {
	if in == nil {
		return nil
	}
	out := new(AndroidVariantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AndroidVariantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidVariantSpec) DeepCopyInto(out *AndroidVariantSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidVariantSpec.
func (in *AndroidVariantSpec) DeepCopy() *AndroidVariantSpec {
	if in == nil {
		return nil
	}
	out := new(AndroidVariantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AndroidVariantStatus) DeepCopyInto(out *AndroidVariantStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AndroidVariantStatus.
func (in *AndroidVariantStatus) DeepCopy() *AndroidVariantStatus {
	if in == nil {
		return nil
	}
	out := new(AndroidVariantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOSVariant) DeepCopyInto(out *IOSVariant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOSVariant.
func (in *IOSVariant) DeepCopy() *IOSVariant {
	if in == nil {
		return nil
	}
	out := new(IOSVariant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOSVariant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOSVariantList) DeepCopyInto(out *IOSVariantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOSVariant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOSVariantList.
func (in *IOSVariantList) DeepCopy() *IOSVariantList {
	if in == nil {
		return nil
	}
	out := new(IOSVariantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOSVariantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOSVariantSpec) DeepCopyInto(out *IOSVariantSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOSVariantSpec.
func (in *IOSVariantSpec) DeepCopy() *IOSVariantSpec {
	if in == nil {
		return nil
	}
	out := new(IOSVariantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOSVariantStatus) DeepCopyInto(out *IOSVariantStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOSVariantStatus.
func (in *IOSVariantStatus) DeepCopy() *IOSVariantStatus {
	if in == nil {
		return nil
	}
	out := new(IOSVariantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushApplication) DeepCopyInto(out *PushApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushApplication.
func (in *PushApplication) DeepCopy() *PushApplication {
	if in == nil {
		return nil
	}
	out := new(PushApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PushApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushApplicationList) DeepCopyInto(out *PushApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PushApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushApplicationList.
func (in *PushApplicationList) DeepCopy() *PushApplicationList {
	if in == nil {
		return nil
	}
	out := new(PushApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PushApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushApplicationSpec) DeepCopyInto(out *PushApplicationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushApplicationSpec.
func (in *PushApplicationSpec) DeepCopy() *PushApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(PushApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushApplicationStatus) DeepCopyInto(out *PushApplicationStatus) {
	*out = *in
	if in.Variants != nil {
		in, out := &in.Variants, &out.Variants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushApplicationStatus.
func (in *PushApplicationStatus) DeepCopy() *PushApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(PushApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServer) DeepCopyInto(out *UnifiedPushServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServer.
func (in *UnifiedPushServer) DeepCopy() *UnifiedPushServer {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnifiedPushServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerBackup) DeepCopyInto(out *UnifiedPushServerBackup) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerBackup.
func (in *UnifiedPushServerBackup) DeepCopy() *UnifiedPushServerBackup {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerList) DeepCopyInto(out *UnifiedPushServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UnifiedPushServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerList.
func (in *UnifiedPushServerList) DeepCopy() *UnifiedPushServerList {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnifiedPushServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerSpec) DeepCopyInto(out *UnifiedPushServerSpec) {
	*out = *in
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]UnifiedPushServerBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerSpec.
func (in *UnifiedPushServerSpec) DeepCopy() *UnifiedPushServerSpec {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnifiedPushServerStatus) DeepCopyInto(out *UnifiedPushServerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnifiedPushServerStatus.
func (in *UnifiedPushServerStatus) DeepCopy() *UnifiedPushServerStatus {
	if in == nil {
		return nil
	}
	out := new(UnifiedPushServerStatus)
	in.DeepCopyInto(out)
	return out
}
