// +build !ignore_autogenerated

/*

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
	"k8s.io/apimachinery/pkg/labels"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkFailureInjection) DeepCopyInto(out *NetworkFailureInjection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
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
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkFailureInjection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
	if in.NumPodsToTarget != nil {
		in, out := &in.NumPodsToTarget, &out.NumPodsToTarget
		*out = new(int)
		**out = **in
	}
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
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLatencyInjection) DeepCopyInto(out *NetworkLatencyInjection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLatencyInjection.
func (in *NetworkLatencyInjection) DeepCopy() *NetworkLatencyInjection {
	if in == nil {
		return nil
	}
	out := new(NetworkLatencyInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkLatencyInjection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLatencyInjectionList) DeepCopyInto(out *NetworkLatencyInjectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkLatencyInjection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLatencyInjectionList.
func (in *NetworkLatencyInjectionList) DeepCopy() *NetworkLatencyInjectionList {
	if in == nil {
		return nil
	}
	out := new(NetworkLatencyInjectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkLatencyInjectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLatencyInjectionSpec) DeepCopyInto(out *NetworkLatencyInjectionSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(labels.Set, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLatencyInjectionSpec.
func (in *NetworkLatencyInjectionSpec) DeepCopy() *NetworkLatencyInjectionSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkLatencyInjectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLatencyInjectionStatus) DeepCopyInto(out *NetworkLatencyInjectionStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLatencyInjectionStatus.
func (in *NetworkLatencyInjectionStatus) DeepCopy() *NetworkLatencyInjectionStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkLatencyInjectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFailureInjection) DeepCopyInto(out *NodeFailureInjection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFailureInjection.
func (in *NodeFailureInjection) DeepCopy() *NodeFailureInjection {
	if in == nil {
		return nil
	}
	out := new(NodeFailureInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFailureInjection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFailureInjectionList) DeepCopyInto(out *NodeFailureInjectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFailureInjection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFailureInjectionList.
func (in *NodeFailureInjectionList) DeepCopy() *NodeFailureInjectionList {
	if in == nil {
		return nil
	}
	out := new(NodeFailureInjectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFailureInjectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFailureInjectionSpec) DeepCopyInto(out *NodeFailureInjectionSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(labels.Set, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Quantity != nil {
		in, out := &in.Quantity, &out.Quantity
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFailureInjectionSpec.
func (in *NodeFailureInjectionSpec) DeepCopy() *NodeFailureInjectionSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFailureInjectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFailureInjectionStatus) DeepCopyInto(out *NodeFailureInjectionStatus) {
	*out = *in
	if in.Injected != nil {
		in, out := &in.Injected, &out.Injected
		*out = make([]NodeFailureInjectionStatusInjectedEntry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFailureInjectionStatus.
func (in *NodeFailureInjectionStatus) DeepCopy() *NodeFailureInjectionStatus {
	if in == nil {
		return nil
	}
	out := new(NodeFailureInjectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFailureInjectionStatusInjectedEntry) DeepCopyInto(out *NodeFailureInjectionStatusInjectedEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFailureInjectionStatusInjectedEntry.
func (in *NodeFailureInjectionStatusInjectedEntry) DeepCopy() *NodeFailureInjectionStatusInjectedEntry {
	if in == nil {
		return nil
	}
	out := new(NodeFailureInjectionStatusInjectedEntry)
	in.DeepCopyInto(out)
	return out
}
