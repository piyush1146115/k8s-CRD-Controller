package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	///_ "k8s.io/code-generator"
)

// +genclient
// +groupName=statefulcontroller.k8s.io
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooDB is a specification for a Foo resource
type FooDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooDBSpec   `json:"spec"`
	Status FooDBStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource
type FooDBSpec struct {
	StatefulsetName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// FooStatus is the status for a Foo resource
type FooDBStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type FooDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []FooDB `json:"items"`
}