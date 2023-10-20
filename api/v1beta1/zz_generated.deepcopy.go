//go:build !ignore_autogenerated

/*
Copyright 2023.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushEncryptedSecret) DeepCopyInto(out *PushEncryptedSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushEncryptedSecret.
func (in *PushEncryptedSecret) DeepCopy() *PushEncryptedSecret {
	if in == nil {
		return nil
	}
	out := new(PushEncryptedSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PushEncryptedSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushEncryptedSecretData) DeepCopyInto(out *PushEncryptedSecretData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushEncryptedSecretData.
func (in *PushEncryptedSecretData) DeepCopy() *PushEncryptedSecretData {
	if in == nil {
		return nil
	}
	out := new(PushEncryptedSecretData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushEncryptedSecretList) DeepCopyInto(out *PushEncryptedSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PushEncryptedSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushEncryptedSecretList.
func (in *PushEncryptedSecretList) DeepCopy() *PushEncryptedSecretList {
	if in == nil {
		return nil
	}
	out := new(PushEncryptedSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PushEncryptedSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushEncryptedSecretSpec) DeepCopyInto(out *PushEncryptedSecretSpec) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]PushEncryptedSecretData, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushEncryptedSecretSpec.
func (in *PushEncryptedSecretSpec) DeepCopy() *PushEncryptedSecretSpec {
	if in == nil {
		return nil
	}
	out := new(PushEncryptedSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushEncryptedSecretStatus) DeepCopyInto(out *PushEncryptedSecretStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushEncryptedSecretStatus.
func (in *PushEncryptedSecretStatus) DeepCopy() *PushEncryptedSecretStatus {
	if in == nil {
		return nil
	}
	out := new(PushEncryptedSecretStatus)
	in.DeepCopyInto(out)
	return out
}