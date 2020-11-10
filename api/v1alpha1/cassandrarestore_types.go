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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type CassandraDatacenterConfig struct {
	// The name to give the new, restored CassandraDatacenter
	Name string `json:"name"`

	// The name to give the C* cluster.
	ClusterName string `json:"clusterName"`
}

// CassandraRestoreSpec defines the desired state of CassandraRestore
type CassandraRestoreSpec struct {

	// The name of the CassandraBackup to restore
	Backup string `json:"backup"`

	CassandraDatacenter CassandraDatacenterConfig `json:"cassandraDatacenter"`
}

// CassandraRestoreStatus defines the observed state of CassandraRestore
type CassandraRestoreStatus struct {
	StartTime metav1.Time `json:"startTime,omitempty"`

	FinishTime metav1.Time `json:"finishTime,omitempty"`

	InProgress []string `json:"inProgress,omitempty"`

	Finished []string `json:"finished,omitempty"`

	Failed []string `json:"failed,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CassandraRestore is the Schema for the cassandrarestores API
type CassandraRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CassandraRestoreSpec   `json:"spec,omitempty"`
	Status CassandraRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraRestoreList contains a list of CassandraRestore
type CassandraRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CassandraRestore{}, &CassandraRestoreList{})
}