package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApiStatusSpec defines the desired state of ApiStatus
// +k8s:openapi-gen=true
type ApiStatusSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	//Image ...
	Image string `json:"image"`
	//ImageName ...
	ImageName string `json:"imageName"`
	//Size ...
	Size int32 `json:"size"`
}

// ApiStatusStatus defines the observed state of ApiStatus
// +k8s:openapi-gen=true
type ApiStatusStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApiStatus is the Schema for the apistatuses API
// +k8s:openapi-gen=true
type ApiStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiStatusSpec   `json:"spec,omitempty"`
	Status ApiStatusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApiStatusList contains a list of ApiStatus
type ApiStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiStatus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApiStatus{}, &ApiStatusList{})
}
