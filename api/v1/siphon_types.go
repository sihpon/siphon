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

// SiphonSpec defines the desired state of Siphon.
type SiphonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Namespace   string           `json:"namespace"`
	Deployments []DeploymentInfo `json:"deployments,omitempty"`
}

func (s *SiphonSpec) UpdateDeploymentInfo(info DeploymentInfo) bool {
	for _, deployment := range s.Deployments {
		if deployment.Name == info.Name {
			deployment = info
			return true
		}
	}

	s.Deployments = append(s.Deployments, info)

	return true
}

type DeploymentInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
	Status    string `json:"status"`
}

// SiphonStatus defines the observed state of Siphon.
type SiphonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Siphon is the Schema for the siphons API.
type Siphon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiphonSpec   `json:"spec,omitempty"`
	Status SiphonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiphonList contains a list of Siphon.
type SiphonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Siphon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Siphon{}, &SiphonList{})
}
