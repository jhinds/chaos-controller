// +build !ignore_autogenerated

/*
Copyright 2019 Datadog.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureInjection) DeepCopyInto(out *NetworkFailureInjection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFailureInjection.
func (in *NetworkFailureInjection) DeepCopy() *NetworkFailureInjection {
	if in == nil {
		return nil
	}
	out := new(NetworkFailureInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkFailureInjection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureInjectionList) DeepCopyInto(out *NetworkFailureInjectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkFailureInjection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFailureInjectionList.
func (in *NetworkFailureInjectionList) DeepCopy() *NetworkFailureInjectionList {
	if in == nil {
		return nil
	}
	out := new(NetworkFailureInjectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkFailureInjectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureInjectionSpec) DeepCopyInto(out *NetworkFailureInjectionSpec) {
	*out = *in
	out.Failure = in.Failure
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(labels.Set, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFailureInjectionSpec.
func (in *NetworkFailureInjectionSpec) DeepCopy() *NetworkFailureInjectionSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkFailureInjectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureInjectionSpecFailure) DeepCopyInto(out *NetworkFailureInjectionSpecFailure) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFailureInjectionSpecFailure.
func (in *NetworkFailureInjectionSpecFailure) DeepCopy() *NetworkFailureInjectionSpecFailure {
	if in == nil {
		return nil
	}
	out := new(NetworkFailureInjectionSpecFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureInjectionStatus) DeepCopyInto(out *NetworkFailureInjectionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkFailureInjectionStatus.
func (in *NetworkFailureInjectionStatus) DeepCopy() *NetworkFailureInjectionStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkFailureInjectionStatus)
	in.DeepCopyInto(out)
	return out
}
