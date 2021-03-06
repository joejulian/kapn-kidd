/*
Copyright 2019 Joe Julian <me@joejulian.name>.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReleaseSpec defines the desired state of Release
type ReleaseSpec struct {
	Chart                *string                     `json:"chart,omitempty"`
	Devel                *bool                       `json:"devel,omitempty"`
	GenerateNameTemplate *string                     `json:"name-template,omitempty"`
	NoHooks              *bool                       `json:"no-hooks,omitempty"`
	PasswordRef          corev1.SecretReference      `json:"password,omitempty"`
	Repo                 *string                     `json:"repo,omitempty"`
	Settings             []corev1.EnvVar             `json:"settings,omitempty"`
	SetFromObj           corev1.LocalObjectReference `json:"set-from-obj,omitempty"`
	Timeout              *string                     `json:"timeout,omitempty"`
	Username             *string                     `json:"username,omitempty"`
	Values               *string                     `json:"values,omitempty"`
	ValuesRef            corev1.LocalObjectReference `json:"values-ref,omitempty"`
	Version              *string                     `json:"version"`
}

// ReleaseStatus defines the observed state of Release
type ReleaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Release is the Schema for the releases API
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReleaseSpec   `json:"spec,omitempty"`
	Status ReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReleaseList contains a list of Release
type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Release `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Release{}, &ReleaseList{})
}
