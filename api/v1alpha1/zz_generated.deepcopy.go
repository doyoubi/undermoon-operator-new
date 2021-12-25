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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Undermoon) DeepCopyInto(out *Undermoon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Undermoon.
func (in *Undermoon) DeepCopy() *Undermoon {
	if in == nil {
		return nil
	}
	out := new(Undermoon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Undermoon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UndermoonList) DeepCopyInto(out *UndermoonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Undermoon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UndermoonList.
func (in *UndermoonList) DeepCopy() *UndermoonList {
	if in == nil {
		return nil
	}
	out := new(UndermoonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UndermoonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UndermoonSpec) DeepCopyInto(out *UndermoonSpec) {
	*out = *in
	if in.BrokerEnvVar != nil {
		in, out := &in.BrokerEnvVar, &out.BrokerEnvVar
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CoordinatorEnvVar != nil {
		in, out := &in.CoordinatorEnvVar, &out.CoordinatorEnvVar
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProxyEnvVar != nil {
		in, out := &in.ProxyEnvVar, &out.ProxyEnvVar
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RedisEnvVar != nil {
		in, out := &in.RedisEnvVar, &out.RedisEnvVar
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.BrokerResources.DeepCopyInto(&out.BrokerResources)
	in.CoordinatorResources.DeepCopyInto(&out.CoordinatorResources)
	in.ProxyResources.DeepCopyInto(&out.ProxyResources)
	in.RedisResources.DeepCopyInto(&out.RedisResources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageVolumes != nil {
		in, out := &in.StorageVolumes, &out.StorageVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RedisVolumeMounts != nil {
		in, out := &in.RedisVolumeMounts, &out.RedisVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RedisVolumeDevices != nil {
		in, out := &in.RedisVolumeDevices, &out.RedisVolumeDevices
		*out = make([]v1.VolumeDevice, len(*in))
		copy(*out, *in)
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]v1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UndermoonSpec.
func (in *UndermoonSpec) DeepCopy() *UndermoonSpec {
	if in == nil {
		return nil
	}
	out := new(UndermoonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UndermoonStatus) DeepCopyInto(out *UndermoonStatus) {
	*out = *in
	in.ScaleDownWaitTimestamp.DeepCopyInto(&out.ScaleDownWaitTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UndermoonStatus.
func (in *UndermoonStatus) DeepCopy() *UndermoonStatus {
	if in == nil {
		return nil
	}
	out := new(UndermoonStatus)
	in.DeepCopyInto(out)
	return out
}
