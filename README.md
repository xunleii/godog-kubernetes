[![Go Reference](https://pkg.go.dev/badge/github.com/xunleii/godog-kubernetes.svg)](https://pkg.go.dev/github.com/xunleii/godog-kubernetes)
![](https://github.com/xunleii/godog-kubernetes/actions/workflows/golang.yml/badge.svg)
![](https://github.com/xunleii/godog-kubernetes/actions/workflows/code-quality.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/xunleii/godog-kubernetes)](https://goreportcard.com/report/github.com/xunleii/godog-kubernetes)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=xunleii_godog-kubernetes&metric=alert_status)](https://sonarcloud.io/dashboard?id=xunleii_godog-kubernetes)

# Godog - Kubernetes feature context

Package **godog-kubernetes** provides a *feature context* to test code that require a Kubernetes cluster, 
like controllers or operators. It provides some **Gherkin** rules to easily create, get, list, update
and delete resources on Kubernetes.

## How to use it

To inject the Kubernetes feature context, you just need to call `NewFeatureContext` on the `InitializeScenario`
function.  

For example :
```go
package main

import (
	"github.com/cucumber/godog"
	kubernetes_ctx "github.com/xunleii/godog-kubernetes"
	kubernetes_ctx_helpers "github.com/xunleii/godog-kubernetes/helpers"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	// This line injects the kubernetes context feature to the current *godog.ScenarioContext
	kubectx, _ := kubernetes_ctx.NewFeatureContext(ctx, kubernetes_ctx.WithFakeClient(runtime.NewScheme()))

	// You can use kubectx to make want you want on Kubernetes ...
	ns, _ := kubectx.Get(schema.FromAPIVersionAndKind("v1", "Namespace"), types.NamespacedName{Name: "default"})                                                   // get the default namespace object
	_ = kubectx.Patch(ns.GroupVersionKind(), types.NamespacedName{Name: ns.GetName()}, types.MergePatchType, []byte(`{"metadata": {"labels": {"key": "value"}}}`)) // add label `key=value` to namespace default
	_, _ = kubectx.Delete(ns.GroupVersionKind(), types.NamespacedName{Name: ns.GetName()})                                                                         // remove the default namespace

	// .. or use it with your controller
	client = kubectx.Client() // get the Kubernetes client if you want to use it in your controller
	ctrl := NewMyController(client)
	ctx.Step(
		`^my controller reconciles '(`+kubernetes_ctx.RxNamespacedName+`)'$`,
		func(target string) error {
			nn, err := kubernetes_ctx_helpers.NamespacedNameFrom(target)
			if err != nil {
				return err
			}

			_, err := ctrl.Request{NamespacedName: nn}
			return err
		},
	)
}
```

## List of available rules

| Type | Rule | Description |
|:----:|:----:|:------------:|
| **CREATE** | `Kubernetes must have <ApiGroupVersionKind> '<NamespacedName>'` | Creates a new resource, without any specific fields. |
| **CREATE** | `Kubernetes creates a new <ApiGroupVersionKind> '<NamespacedName>'` | Creates a new resource, without any specific fields. |
| **CREATE** | `Kubernetes must have <ApiGroupVersionKind> '<NamespacedName>' with <YAML>` | Creates a new resource, with the given definition. |
| **CREATE** | `Kubernetes creates a new <ApiGroupVersionKind> '<NamespacedName>' with <YAML>` | Creates a new resource, with the given definition. |
| **CREATE** | `Kubernetes must have <ApiGroupVersionKind> '<NamespacedName>' from <filename>` | Creates a new resource, with then definition available in the given filename. |
| **CREATE** | `Kubernetes creates a new <ApiGroupVersionKind> '<NamespacedName>' from <filename>` | Creates a new resource, with then definition available in the given filename. |
| **CREATE** | `Kubernetes must have the following resources <RESOURCES_TABLE>` | Creates several resources in a row, without any specific fields (useful for Namespaces). |
| **CREATE** | `Kubernetes creates the following resources <RESOURCES_TABLE>` | Creates several resources in a row, without any specific fields (useful for Namespaces). |
| **LIST** | `Kubernetes has <NumberResources> <ApiGroupVersionKind>` | Compare the current number of a specific resource with the given number. |
| **LIST** | `Kubernetes has <NumberResources> <ApiGroupVersionKind> in namespace '<Namespace>'` | Compare the current number of a specific resource with the given number. |
| **GET** | `Kubernetes has <ApiGroupVersionKind> '<NamespacedName>'` | Validates the fact that Kubernetes has the specified resource. |
| **GET** | `Kubernetes doesn't have <ApiGroupVersionKind> '<NamespacedName>'` | Validates the fact that Kubernetes doesn't have the specified resource. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is similar to '<NamespacedName>'` | Compares two resources in order to determine if they are similar. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is not similar to '<NamespacedName>'` | Compares two resources in order to determine if they are not similar. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is equal to '<NamespacedName>'` | Compares two resources in order to determine if they are equal. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is not equal to '<NamespacedName>'` | Compares two resources in order to determine if they are not equal. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has '<FieldPath>'` | Validates the fact that the specific resource has the field. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' doesn't have '<FieldPath>'` | Validates the fact that the specific resource doesn't have the field. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has '<FieldPath>=<FieldValue>'` | Validates the fact that the specific resource field has the given value. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has '<FieldPath>!=<FieldValue>'` | Validates the fact that the specific resource field is different than the given value. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has label '<LabelName>'` | Validates the fact that the specific resource has the given label. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' doesn't have label '<LabelName>'` | Validates the fact that the specific resource doesn't have the given label. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has label '<LabelName>=<LabelValue>'` | Validates the fact that the specific resource label has the given value. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has label '<LabelName>!=<LabelValue>'` | Validates the fact that the specific resource label doesn't have the given value. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has annotation '<AnnotationName>'` | Validates the fact that the specific resource has the given annotation. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' doesn't have annotation '<AnnotationName>'` | Validates the fact that the specific resource doesn't have the given annotation. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has annotation '<AnnotationName>=<AnnotationValue>'` | Validates the fact that the specific resource annotation has the given value. |
| **GET** | `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has annotation '<AnnotationName>!=<AnnotationValue>'` | Validates the fact that the specific resource annotation doesn't have the given value. |
| **PATCH** | `Kubernetes labelizes <ApiGroupVersionKind> '<NamespacedName>' with '<LabelName>=<LabelValue>'` | Adds or modifies a specific resource label with the given value. |
| **PATCH** | `Kubernetes removes label <LabelName> on <ApiGroupVersionKind> '<NamespacedName>'` | Removes the given label on the specified resource. |
| **PATCH** | `Kubernetes updates label <LabelName> on <ApiGroupVersionKind> '<NamespacedName>' with '<LabelValue>'` | Updates the given label on the specified resource with the given value. |
| **PATCH** | `Kubernetes annotates <ApiGroupVersionKind> '<NamespacedName>' with '<AnnotationName>=<AnnotationValue>'` | Adds or modifies a specific resource annotation with the given value. |
| **PATCH** | `Kubernetes removes annotation <AnnotationName> on <ApiGroupVersionKind> '<NamespacedName>'` | Removes the given annotation on the specified resource. |
| **PATCH** | `Kubernetes updates annotation <AnnotationName> on <ApiGroupVersionKind> '<NamespacedName>' with '<AnnotationValue>'` | Updates the given annotation on the specified resource with the given value. |
| **PATCH** | `Kubernetes patches <ApiGroupVersionKind> '<NamespacedName>' with <YAML>` | Patches a specific resource with the given patch (it use StrategicMergePatchType... |
| **DELETE** | `Kubernetes removes <ApiGroupVersionKind> '<NamespacedName>'` | Removes the specified resource. |
| **DELETE** | `Kubernetes removes the following resources <RESOURCES_TABLE>` | Creates several resources in a row. |

## LICENSE
**godog-kubernetes** is licensed under the [Apache v2](https://www.apache.org/licenses/LICENSE-2.0).