package v1beta1

import (
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	builders.Scheme.AddKnownTypes(SchemeGroupVersion, &FooBarOptions{})
	builders.ParameterScheme.AddKnownTypes(SchemeGroupVersion, &FooBarOptions{})
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooBarOptions FooBarOptions
type FooBarOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Arg1              string `json:"arg1"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooBarOptionsList FooBarOptionsList
type FooBarOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FooBarOptions `json:"items"`
}
