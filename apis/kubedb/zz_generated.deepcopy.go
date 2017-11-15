// +build !ignore_autogenerated

/*
Copyright 2017 The KubeDB Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package kubedb

import (
	reflect "reflect"

	api "github.com/appscode/kutil/tools/monitoring/api"
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureSpec).DeepCopyInto(out.(*AzureSpec))
			return nil
		}, InType: reflect.TypeOf(&AzureSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupScheduleSpec).DeepCopyInto(out.(*BackupScheduleSpec))
			return nil
		}, InType: reflect.TypeOf(&BackupScheduleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DormantDatabase).DeepCopyInto(out.(*DormantDatabase))
			return nil
		}, InType: reflect.TypeOf(&DormantDatabase{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DormantDatabaseList).DeepCopyInto(out.(*DormantDatabaseList))
			return nil
		}, InType: reflect.TypeOf(&DormantDatabaseList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DormantDatabaseSpec).DeepCopyInto(out.(*DormantDatabaseSpec))
			return nil
		}, InType: reflect.TypeOf(&DormantDatabaseSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DormantDatabaseStatus).DeepCopyInto(out.(*DormantDatabaseStatus))
			return nil
		}, InType: reflect.TypeOf(&DormantDatabaseStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Elasticsearch).DeepCopyInto(out.(*Elasticsearch))
			return nil
		}, InType: reflect.TypeOf(&Elasticsearch{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchList).DeepCopyInto(out.(*ElasticsearchList))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchSpec).DeepCopyInto(out.(*ElasticsearchSpec))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ElasticsearchStatus).DeepCopyInto(out.(*ElasticsearchStatus))
			return nil
		}, InType: reflect.TypeOf(&ElasticsearchStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GCSSpec).DeepCopyInto(out.(*GCSSpec))
			return nil
		}, InType: reflect.TypeOf(&GCSSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InitSpec).DeepCopyInto(out.(*InitSpec))
			return nil
		}, InType: reflect.TypeOf(&InitSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LocalSpec).DeepCopyInto(out.(*LocalSpec))
			return nil
		}, InType: reflect.TypeOf(&LocalSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MySQL).DeepCopyInto(out.(*MySQL))
			return nil
		}, InType: reflect.TypeOf(&MySQL{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MySQLList).DeepCopyInto(out.(*MySQLList))
			return nil
		}, InType: reflect.TypeOf(&MySQLList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MySQLSpec).DeepCopyInto(out.(*MySQLSpec))
			return nil
		}, InType: reflect.TypeOf(&MySQLSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MySQLStatus).DeepCopyInto(out.(*MySQLStatus))
			return nil
		}, InType: reflect.TypeOf(&MySQLStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Origin).DeepCopyInto(out.(*Origin))
			return nil
		}, InType: reflect.TypeOf(&Origin{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*OriginSpec).DeepCopyInto(out.(*OriginSpec))
			return nil
		}, InType: reflect.TypeOf(&OriginSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Postgres).DeepCopyInto(out.(*Postgres))
			return nil
		}, InType: reflect.TypeOf(&Postgres{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PostgresList).DeepCopyInto(out.(*PostgresList))
			return nil
		}, InType: reflect.TypeOf(&PostgresList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PostgresSchemaInfo).DeepCopyInto(out.(*PostgresSchemaInfo))
			return nil
		}, InType: reflect.TypeOf(&PostgresSchemaInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PostgresSpec).DeepCopyInto(out.(*PostgresSpec))
			return nil
		}, InType: reflect.TypeOf(&PostgresSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PostgresStatus).DeepCopyInto(out.(*PostgresStatus))
			return nil
		}, InType: reflect.TypeOf(&PostgresStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PostgresSummary).DeepCopyInto(out.(*PostgresSummary))
			return nil
		}, InType: reflect.TypeOf(&PostgresSummary{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PostgresTableInfo).DeepCopyInto(out.(*PostgresTableInfo))
			return nil
		}, InType: reflect.TypeOf(&PostgresTableInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*S3Spec).DeepCopyInto(out.(*S3Spec))
			return nil
		}, InType: reflect.TypeOf(&S3Spec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ScriptSourceSpec).DeepCopyInto(out.(*ScriptSourceSpec))
			return nil
		}, InType: reflect.TypeOf(&ScriptSourceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Snapshot).DeepCopyInto(out.(*Snapshot))
			return nil
		}, InType: reflect.TypeOf(&Snapshot{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SnapshotList).DeepCopyInto(out.(*SnapshotList))
			return nil
		}, InType: reflect.TypeOf(&SnapshotList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SnapshotSourceSpec).DeepCopyInto(out.(*SnapshotSourceSpec))
			return nil
		}, InType: reflect.TypeOf(&SnapshotSourceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SnapshotSpec).DeepCopyInto(out.(*SnapshotSpec))
			return nil
		}, InType: reflect.TypeOf(&SnapshotSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SnapshotStatus).DeepCopyInto(out.(*SnapshotStatus))
			return nil
		}, InType: reflect.TypeOf(&SnapshotStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SnapshotStorageSpec).DeepCopyInto(out.(*SnapshotStorageSpec))
			return nil
		}, InType: reflect.TypeOf(&SnapshotStorageSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SwiftSpec).DeepCopyInto(out.(*SwiftSpec))
			return nil
		}, InType: reflect.TypeOf(&SwiftSpec{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSpec) DeepCopyInto(out *AzureSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSpec.
func (in *AzureSpec) DeepCopy() *AzureSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleSpec) DeepCopyInto(out *BackupScheduleSpec) {
	*out = *in
	in.SnapshotStorageSpec.DeepCopyInto(&out.SnapshotStorageSpec)
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleSpec.
func (in *BackupScheduleSpec) DeepCopy() *BackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DormantDatabase) DeepCopyInto(out *DormantDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DormantDatabase.
func (in *DormantDatabase) DeepCopy() *DormantDatabase {
	if in == nil {
		return nil
	}
	out := new(DormantDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DormantDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DormantDatabaseList) DeepCopyInto(out *DormantDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DormantDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DormantDatabaseList.
func (in *DormantDatabaseList) DeepCopy() *DormantDatabaseList {
	if in == nil {
		return nil
	}
	out := new(DormantDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DormantDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DormantDatabaseSpec) DeepCopyInto(out *DormantDatabaseSpec) {
	*out = *in
	in.Origin.DeepCopyInto(&out.Origin)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DormantDatabaseSpec.
func (in *DormantDatabaseSpec) DeepCopy() *DormantDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DormantDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DormantDatabaseStatus) DeepCopyInto(out *DormantDatabaseStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.PausingTime != nil {
		in, out := &in.PausingTime, &out.PausingTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.WipeOutTime != nil {
		in, out := &in.WipeOutTime, &out.WipeOutTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DormantDatabaseStatus.
func (in *DormantDatabaseStatus) DeepCopy() *DormantDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DormantDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elasticsearch) DeepCopyInto(out *Elasticsearch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elasticsearch.
func (in *Elasticsearch) DeepCopy() *Elasticsearch {
	if in == nil {
		return nil
	}
	out := new(Elasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Elasticsearch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchList) DeepCopyInto(out *ElasticsearchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Elasticsearch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchList.
func (in *ElasticsearchList) DeepCopy() *ElasticsearchList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PersistentVolumeClaimSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		if *in == nil {
			*out = nil
		} else {
			*out = new(InitSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.BackupSchedule != nil {
		in, out := &in.BackupSchedule, &out.BackupSchedule
		if *in == nil {
			*out = nil
		} else {
			*out = new(BackupScheduleSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		if *in == nil {
			*out = nil
		} else {
			*out = new(api.AgentSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchStatus) DeepCopyInto(out *ElasticsearchStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchStatus.
func (in *ElasticsearchStatus) DeepCopy() *ElasticsearchStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSSpec) DeepCopyInto(out *GCSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSSpec.
func (in *GCSSpec) DeepCopy() *GCSSpec {
	if in == nil {
		return nil
	}
	out := new(GCSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitSpec) DeepCopyInto(out *InitSpec) {
	*out = *in
	if in.ScriptSource != nil {
		in, out := &in.ScriptSource, &out.ScriptSource
		if *in == nil {
			*out = nil
		} else {
			*out = new(ScriptSourceSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SnapshotSource != nil {
		in, out := &in.SnapshotSource, &out.SnapshotSource
		if *in == nil {
			*out = nil
		} else {
			*out = new(SnapshotSourceSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitSpec.
func (in *InitSpec) DeepCopy() *InitSpec {
	if in == nil {
		return nil
	}
	out := new(InitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpec) DeepCopyInto(out *LocalSpec) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpec.
func (in *LocalSpec) DeepCopy() *LocalSpec {
	if in == nil {
		return nil
	}
	out := new(LocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQL) DeepCopyInto(out *MySQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQL.
func (in *MySQL) DeepCopy() *MySQL {
	if in == nil {
		return nil
	}
	out := new(MySQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLList) DeepCopyInto(out *MySQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLList.
func (in *MySQLList) DeepCopy() *MySQLList {
	if in == nil {
		return nil
	}
	out := new(MySQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLSpec) DeepCopyInto(out *MySQLSpec) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PersistentVolumeClaimSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DatabaseSecret != nil {
		in, out := &in.DatabaseSecret, &out.DatabaseSecret
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretVolumeSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		if *in == nil {
			*out = nil
		} else {
			*out = new(InitSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.BackupSchedule != nil {
		in, out := &in.BackupSchedule, &out.BackupSchedule
		if *in == nil {
			*out = nil
		} else {
			*out = new(BackupScheduleSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		if *in == nil {
			*out = nil
		} else {
			*out = new(api.AgentSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLSpec.
func (in *MySQLSpec) DeepCopy() *MySQLSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLStatus) DeepCopyInto(out *MySQLStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLStatus.
func (in *MySQLStatus) DeepCopy() *MySQLStatus {
	if in == nil {
		return nil
	}
	out := new(MySQLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Origin) DeepCopyInto(out *Origin) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Origin.
func (in *Origin) DeepCopy() *Origin {
	if in == nil {
		return nil
	}
	out := new(Origin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginSpec) DeepCopyInto(out *OriginSpec) {
	*out = *in
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		if *in == nil {
			*out = nil
		} else {
			*out = new(ElasticsearchSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		if *in == nil {
			*out = nil
		} else {
			*out = new(PostgresSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.MySQL != nil {
		in, out := &in.MySQL, &out.MySQL
		if *in == nil {
			*out = nil
		} else {
			*out = new(MySQLSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginSpec.
func (in *OriginSpec) DeepCopy() *OriginSpec {
	if in == nil {
		return nil
	}
	out := new(OriginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Postgres) DeepCopyInto(out *Postgres) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Postgres.
func (in *Postgres) DeepCopy() *Postgres {
	if in == nil {
		return nil
	}
	out := new(Postgres)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Postgres) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresList) DeepCopyInto(out *PostgresList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Postgres, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresList.
func (in *PostgresList) DeepCopy() *PostgresList {
	if in == nil {
		return nil
	}
	out := new(PostgresList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSchemaInfo) DeepCopyInto(out *PostgresSchemaInfo) {
	*out = *in
	if in.Table != nil {
		in, out := &in.Table, &out.Table
		*out = make(map[string]*PostgresTableInfo, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(PostgresTableInfo)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSchemaInfo.
func (in *PostgresSchemaInfo) DeepCopy() *PostgresSchemaInfo {
	if in == nil {
		return nil
	}
	out := new(PostgresSchemaInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSpec) DeepCopyInto(out *PostgresSpec) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PersistentVolumeClaimSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DatabaseSecret != nil {
		in, out := &in.DatabaseSecret, &out.DatabaseSecret
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretVolumeSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		if *in == nil {
			*out = nil
		} else {
			*out = new(InitSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.BackupSchedule != nil {
		in, out := &in.BackupSchedule, &out.BackupSchedule
		if *in == nil {
			*out = nil
		} else {
			*out = new(BackupScheduleSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		if *in == nil {
			*out = nil
		} else {
			*out = new(api.AgentSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSpec.
func (in *PostgresSpec) DeepCopy() *PostgresSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresStatus) DeepCopyInto(out *PostgresStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresStatus.
func (in *PostgresStatus) DeepCopy() *PostgresStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSummary) DeepCopyInto(out *PostgresSummary) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = make(map[string]*PostgresSchemaInfo, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(PostgresSchemaInfo)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSummary.
func (in *PostgresSummary) DeepCopy() *PostgresSummary {
	if in == nil {
		return nil
	}
	out := new(PostgresSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresTableInfo) DeepCopyInto(out *PostgresTableInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresTableInfo.
func (in *PostgresTableInfo) DeepCopy() *PostgresTableInfo {
	if in == nil {
		return nil
	}
	out := new(PostgresTableInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Spec) DeepCopyInto(out *S3Spec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Spec.
func (in *S3Spec) DeepCopy() *S3Spec {
	if in == nil {
		return nil
	}
	out := new(S3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptSourceSpec) DeepCopyInto(out *ScriptSourceSpec) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptSourceSpec.
func (in *ScriptSourceSpec) DeepCopy() *ScriptSourceSpec {
	if in == nil {
		return nil
	}
	out := new(ScriptSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Snapshot) DeepCopyInto(out *Snapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Snapshot.
func (in *Snapshot) DeepCopy() *Snapshot {
	if in == nil {
		return nil
	}
	out := new(Snapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Snapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotList) DeepCopyInto(out *SnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Snapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotList.
func (in *SnapshotList) DeepCopy() *SnapshotList {
	if in == nil {
		return nil
	}
	out := new(SnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSourceSpec) DeepCopyInto(out *SnapshotSourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSourceSpec.
func (in *SnapshotSourceSpec) DeepCopy() *SnapshotSourceSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpec) DeepCopyInto(out *SnapshotSpec) {
	*out = *in
	in.SnapshotStorageSpec.DeepCopyInto(&out.SnapshotStorageSpec)
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpec.
func (in *SnapshotSpec) DeepCopy() *SnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotStatus) DeepCopyInto(out *SnapshotStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotStatus.
func (in *SnapshotStatus) DeepCopy() *SnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotStorageSpec) DeepCopyInto(out *SnapshotStorageSpec) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		if *in == nil {
			*out = nil
		} else {
			*out = new(LocalSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		if *in == nil {
			*out = nil
		} else {
			*out = new(S3Spec)
			**out = **in
		}
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		if *in == nil {
			*out = nil
		} else {
			*out = new(GCSSpec)
			**out = **in
		}
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		if *in == nil {
			*out = nil
		} else {
			*out = new(AzureSpec)
			**out = **in
		}
	}
	if in.Swift != nil {
		in, out := &in.Swift, &out.Swift
		if *in == nil {
			*out = nil
		} else {
			*out = new(SwiftSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotStorageSpec.
func (in *SnapshotStorageSpec) DeepCopy() *SnapshotStorageSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftSpec) DeepCopyInto(out *SwiftSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftSpec.
func (in *SwiftSpec) DeepCopy() *SwiftSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftSpec)
	in.DeepCopyInto(out)
	return out
}
