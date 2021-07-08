/*
Copyright 2021.

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

// SubsSpec defines the desired state of Subs
type SubsSpec struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

// SubsStatus defines the observed state of Subs
type SubsStatus struct {
	Output int `json:"result"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:JSONPath=".spec.num1",name=Number 1,type=integer
//+kubebuilder:printcolumn:JSONPath=".spec.num2",name=Number 2,type=integer
//+kubebuilder:printcolumn:JSONPath=".status.result",name=Result,type=integer
// Subs is the Schema for the subs API
type Subs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubsSpec   `json:"spec,omitempty"`
	Status SubsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SubsList contains a list of Subs
type SubsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subs `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Subs{}, &SubsList{})
}
