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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackup) DeepCopyInto(out *CassandraBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackup.
func (in *CassandraBackup) DeepCopy() *CassandraBackup {
	if in == nil {
		return nil
	}
	out := new(CassandraBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupList) DeepCopyInto(out *CassandraBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupList.
func (in *CassandraBackupList) DeepCopy() *CassandraBackupList {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupSpec) DeepCopyInto(out *CassandraBackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupSpec.
func (in *CassandraBackupSpec) DeepCopy() *CassandraBackupSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupStatus) DeepCopyInto(out *CassandraBackupStatus) {
	*out = *in
	if in.CassdcTemplateSpec != nil {
		in, out := &in.CassdcTemplateSpec, &out.CassdcTemplateSpec
		*out = new(CassandraDatacenterTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.FinishTime.DeepCopyInto(&out.FinishTime)
	if in.InProgress != nil {
		in, out := &in.InProgress, &out.InProgress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Finished != nil {
		in, out := &in.Finished, &out.Finished
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupStatus.
func (in *CassandraBackupStatus) DeepCopy() *CassandraBackupStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterConfig) DeepCopyInto(out *CassandraDatacenterConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterConfig.
func (in *CassandraDatacenterConfig) DeepCopy() *CassandraDatacenterConfig {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterTemplateSpec) DeepCopyInto(out *CassandraDatacenterTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterTemplateSpec.
func (in *CassandraDatacenterTemplateSpec) DeepCopy() *CassandraDatacenterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestore) DeepCopyInto(out *CassandraRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestore.
func (in *CassandraRestore) DeepCopy() *CassandraRestore {
	if in == nil {
		return nil
	}
	out := new(CassandraRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestoreList) DeepCopyInto(out *CassandraRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestoreList.
func (in *CassandraRestoreList) DeepCopy() *CassandraRestoreList {
	if in == nil {
		return nil
	}
	out := new(CassandraRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestoreSpec) DeepCopyInto(out *CassandraRestoreSpec) {
	*out = *in
	out.CassandraDatacenter = in.CassandraDatacenter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestoreSpec.
func (in *CassandraRestoreSpec) DeepCopy() *CassandraRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestoreStatus) DeepCopyInto(out *CassandraRestoreStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.FinishTime.DeepCopyInto(&out.FinishTime)
	if in.InProgress != nil {
		in, out := &in.InProgress, &out.InProgress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Finished != nil {
		in, out := &in.Finished, &out.Finished
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestoreStatus.
func (in *CassandraRestoreStatus) DeepCopy() *CassandraRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraRestoreStatus)
	in.DeepCopyInto(out)
	return out
}
