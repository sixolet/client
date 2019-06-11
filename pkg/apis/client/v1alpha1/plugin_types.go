/*
Copyright 2018 The Knative Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Plugin struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec PluginSpec `json:"spec,omitempty"`
	// +optional
	Status PluginStatus `json:"status,omitempty"`
}

type Implementation struct {
	// +optional
	Url string `json:"url,omitempty"`
	// +optional
	Digest string `json:"digest,omitempty"`
	// +optional
	Platform string `json:"platform,omitempty"`
}

type PluginSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Implementations []Implementation `json:"implementations,omitempty"`
	// +optional
	Version string `json:"version,omitempty"`
	// +optional
	Command []string `json:"command,omitempty"`
}

type PluginStatus struct {
}

// PluginList is a list of Plugin resources

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Plugin `json:"items"`
}
