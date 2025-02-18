/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SiphonManagedWorkloadSpec defines the desired state of SiphonManagedWorkload.
type SiphonManagedWorkloadSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ID          string      `json:"id"`
	Namespace   string      `json:"namespace"`
	Name        string      `json:"name,omitempty"`
	Version     string      `json:"version,omitempty"`
	Description string      `json:"description,omitempty"`
	Labels      []string    `json:"labels,omitempty"`
	Artifacts   []Artifact  `json:"artifacts,omitempty"`
	CreatedAt   metav1.Time `json:"createdAt,omitempty"`
	UpdatedAt   metav1.Time `json:"updatedAt,omitempty"`
}

type Artifact struct {
	ArtifactType ArtifactType `json:"type,omitempty"`
	Name         string       `json:"name,omitempty"`
	Tag          string       `json:"tag,omitempty"`
}

type ArtifactType string

const (
	Container ArtifactType = "container"
	Asset     ArtifactType = "asset"
)

// SiphonManagedWorkloadStatus defines the observed state of SiphonManagedWorkload.
type SiphonManagedWorkloadStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SiphonManagedWorkload is the Schema for the siphonmanagedworkloads API.
type SiphonManagedWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiphonManagedWorkloadSpec   `json:"spec,omitempty"`
	Status SiphonManagedWorkloadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiphonManagedWorkloadList contains a list of SiphonManagedWorkload.
type SiphonManagedWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiphonManagedWorkload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiphonManagedWorkload{}, &SiphonManagedWorkloadList{})
}
