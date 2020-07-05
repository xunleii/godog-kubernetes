package kubernetes_ctx

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"kubernetes_ctx/helpers"
)

// CreateSingleResource implements the GoDoc step
// - `Kubernetes must have <ApiGroupVersionKind> '<NamespacedName>'`
// - `Kubernetes creates a new <ApiGroupVersionKind> '<NamespacedName>'`
// It creates a new resource, without any specific fields.
func CreateSingleResource(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes (?:must have|creates a new) (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, resourceName string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := helpers.NamespacedNameFrom(resourceName)
			if err != nil {
				return err
			}

			return ctx.Create(groupVersionKind, namespacedName, &unstructured.Unstructured{})
		},
	)
}

// CreateSingleResourceWith implements the GoDoc step
// - `Kubernetes must have <ApiGroupVersionKind> '<NamespacedName>' with <YAML>`
// - `Kubernetes creates a new <ApiGroupVersionKind> '<NamespacedName>' with <YAML>`
// It creates a new resource, with the given definition.
func CreateSingleResourceWith(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes (?:must have|creates a new) (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' with$`,
		func(groupVersionKindStr, resourceName string, yamlObj helpers.YamlDocString) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := helpers.NamespacedNameFrom(resourceName)
			if err != nil {
				return err
			}

			obj, err := helpers.UnmarshalYamlDocString(yamlObj)
			if err != nil {
				return err
			}

			return ctx.Create(groupVersionKind, namespacedName, &unstructured.Unstructured{Object: obj})
		},
	)
}

// CreateSingleResourceFrom implements the GoDoc step
// - `Kubernetes must have <ApiGroupVersionKind> '<NamespacedName>' from <filename>`
// - `Kubernetes creates a new <ApiGroupVersionKind> '<NamespacedName>' from <filename>`
// It creates a new resource, with then definition available in the given filename.
func CreateSingleResourceFrom(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes (?:must have|creates a new) (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' from (.+)$`,
		func(groupVersionKindStr, resourceName, fileName string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := helpers.NamespacedNameFrom(resourceName)
			if err != nil {
				return err
			}

			data, err := ioutil.ReadFile(fileName)
			if err != nil {
				return err
			}

			var obj unstructured.Unstructured
			err = yaml.Unmarshal(data, &obj.Object)
			if err != nil {
				return err
			}

			return ctx.Create(groupVersionKind, namespacedName, &obj)
		},
	)
}

// CreateMultiResources implements the GoDoc step
// - `Kubernetes must have the following resources <RESOURCES_TABLE>`
// - `Kubernetes creates the following resources <RESOURCES_TABLE>`
// It creates several resources in a row, without any specific fields (useful for Namespaces).
func CreateMultiResources(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes (?:must have|creates) the following resources$`,
		func(table helpers.ResourceTable) error {
			resources, err := helpers.UnmarshalResourceTable(table)
			if err != nil {
				return err
			}

			for _, resource := range resources {
				err := ctx.Create(resource.GroupVersionKind(), resource.NamespacedName(), &unstructured.Unstructured{})
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}
