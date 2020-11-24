// +build !ignore_autogenerated

/*
Copyright 2019 The etcd-operator Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ABSBackupSource) DeepCopyInto(out *ABSBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ABSBackupSource.
func (in *ABSBackupSource) DeepCopy() *ABSBackupSource {
	if in == nil {
		return nil
	}
	out := new(ABSBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ABSRestoreSource) DeepCopyInto(out *ABSRestoreSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ABSRestoreSource.
func (in *ABSRestoreSource) DeepCopy() *ABSRestoreSource {
	if in == nil {
		return nil
	}
	out := new(ABSRestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicy) DeepCopyInto(out *BackupPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicy.
func (in *BackupPolicy) DeepCopy() *BackupPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSource) DeepCopyInto(out *BackupSource) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3BackupSource)
		**out = **in
	}
	if in.ABS != nil {
		in, out := &in.ABS, &out.ABS
		*out = new(ABSBackupSource)
		**out = **in
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(GCSBackupSource)
		**out = **in
	}
	if in.OSS != nil {
		in, out := &in.OSS, &out.OSS
		*out = new(OSSBackupSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSource.
func (in *BackupSource) DeepCopy() *BackupSource {
	if in == nil {
		return nil
	}
	out := new(BackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.EtcdEndpoints != nil {
		in, out := &in.EtcdEndpoints, &out.EtcdEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BackupPolicy != nil {
		in, out := &in.BackupPolicy, &out.BackupPolicy
		*out = new(BackupPolicy)
		**out = **in
	}
	in.BackupSource.DeepCopyInto(&out.BackupSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	in.LastSuccessDate.DeepCopyInto(&out.LastSuccessDate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(PodPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		copy(*out, *in)
	}
	in.Members.DeepCopyInto(&out.Members)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdBackup) DeepCopyInto(out *EtcdBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdBackup.
func (in *EtcdBackup) DeepCopy() *EtcdBackup {
	if in == nil {
		return nil
	}
	out := new(EtcdBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdBackupList) DeepCopyInto(out *EtcdBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EtcdBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdBackupList.
func (in *EtcdBackupList) DeepCopy() *EtcdBackupList {
	if in == nil {
		return nil
	}
	out := new(EtcdBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdCluster) DeepCopyInto(out *EtcdCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdCluster.
func (in *EtcdCluster) DeepCopy() *EtcdCluster {
	if in == nil {
		return nil
	}
	out := new(EtcdCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdClusterList) DeepCopyInto(out *EtcdClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EtcdCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdClusterList.
func (in *EtcdClusterList) DeepCopy() *EtcdClusterList {
	if in == nil {
		return nil
	}
	out := new(EtcdClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdClusterRef) DeepCopyInto(out *EtcdClusterRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdClusterRef.
func (in *EtcdClusterRef) DeepCopy() *EtcdClusterRef {
	if in == nil {
		return nil
	}
	out := new(EtcdClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdRestore) DeepCopyInto(out *EtcdRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdRestore.
func (in *EtcdRestore) DeepCopy() *EtcdRestore {
	if in == nil {
		return nil
	}
	out := new(EtcdRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdRestoreList) DeepCopyInto(out *EtcdRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EtcdRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdRestoreList.
func (in *EtcdRestoreList) DeepCopy() *EtcdRestoreList {
	if in == nil {
		return nil
	}
	out := new(EtcdRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSBackupSource) DeepCopyInto(out *GCSBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSBackupSource.
func (in *GCSBackupSource) DeepCopy() *GCSBackupSource {
	if in == nil {
		return nil
	}
	out := new(GCSBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSRestoreSource) DeepCopyInto(out *GCSRestoreSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSRestoreSource.
func (in *GCSRestoreSource) DeepCopy() *GCSRestoreSource {
	if in == nil {
		return nil
	}
	out := new(GCSRestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSecret) DeepCopyInto(out *MemberSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSecret.
func (in *MemberSecret) DeepCopy() *MemberSecret {
	if in == nil {
		return nil
	}
	out := new(MemberSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersStatus) DeepCopyInto(out *MembersStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Unready != nil {
		in, out := &in.Unready, &out.Unready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersStatus.
func (in *MembersStatus) DeepCopy() *MembersStatus {
	if in == nil {
		return nil
	}
	out := new(MembersStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSSBackupSource) DeepCopyInto(out *OSSBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSSBackupSource.
func (in *OSSBackupSource) DeepCopy() *OSSBackupSource {
	if in == nil {
		return nil
	}
	out := new(OSSBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSSRestoreSource) DeepCopyInto(out *OSSRestoreSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSSRestoreSource.
func (in *OSSRestoreSource) DeepCopy() *OSSRestoreSource {
	if in == nil {
		return nil
	}
	out := new(OSSRestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodPolicy) DeepCopyInto(out *PodPolicy) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EtcdEnv != nil {
		in, out := &in.EtcdEnv, &out.EtcdEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodPolicy.
func (in *PodPolicy) DeepCopy() *PodPolicy {
	if in == nil {
		return nil
	}
	out := new(PodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSource) DeepCopyInto(out *RestoreSource) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3RestoreSource)
		**out = **in
	}
	if in.ABS != nil {
		in, out := &in.ABS, &out.ABS
		*out = new(ABSRestoreSource)
		**out = **in
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(GCSRestoreSource)
		**out = **in
	}
	if in.OSS != nil {
		in, out := &in.OSS, &out.OSS
		*out = new(OSSRestoreSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSource.
func (in *RestoreSource) DeepCopy() *RestoreSource {
	if in == nil {
		return nil
	}
	out := new(RestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSpec) DeepCopyInto(out *RestoreSpec) {
	*out = *in
	in.RestoreSource.DeepCopyInto(&out.RestoreSource)
	out.EtcdCluster = in.EtcdCluster
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSpec.
func (in *RestoreSpec) DeepCopy() *RestoreSpec {
	if in == nil {
		return nil
	}
	out := new(RestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreStatus) DeepCopyInto(out *RestoreStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreStatus.
func (in *RestoreStatus) DeepCopy() *RestoreStatus {
	if in == nil {
		return nil
	}
	out := new(RestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BackupSource) DeepCopyInto(out *S3BackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BackupSource.
func (in *S3BackupSource) DeepCopy() *S3BackupSource {
	if in == nil {
		return nil
	}
	out := new(S3BackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3RestoreSource) DeepCopyInto(out *S3RestoreSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3RestoreSource.
func (in *S3RestoreSource) DeepCopy() *S3RestoreSource {
	if in == nil {
		return nil
	}
	out := new(S3RestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticTLS) DeepCopyInto(out *StaticTLS) {
	*out = *in
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(MemberSecret)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticTLS.
func (in *StaticTLS) DeepCopy() *StaticTLS {
	if in == nil {
		return nil
	}
	out := new(StaticTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPolicy) DeepCopyInto(out *TLSPolicy) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(StaticTLS)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPolicy.
func (in *TLSPolicy) DeepCopy() *TLSPolicy {
	if in == nil {
		return nil
	}
	out := new(TLSPolicy)
	in.DeepCopyInto(out)
	return out
}
