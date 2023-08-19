package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "mygroup.example.com"

var SchemeGroupVersion = schema.GroupVersion{
	Group:   GroupName,
	Version: "v1alpha1",
}

var (
	SchemaBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemebuilder = &SchemaBuilder
	AddToScheme        = localSchemebuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &MyResource{}, &MyResourceList{})
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
