package kubernetes_ctx

import (
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"kubernetes_ctx/helpers"
)

// CountResources implements the GoDoc step
// - `Kubernetes has <NumberResources> <ApiGroupVersionKind>`
// It compare the current number of a specific resource with the given number.
func CountResources(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes has (\d+) (`+RxGroupVersionKind+`)$`,
		func(n int, groupVersionKindStr string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}

			objs, err := ctx.List(groupVersionKind)
			if err != nil {
				return err
			}

			if len(objs) != n {
				if len(objs) == 0 {
					return fmt.Errorf("no %s found", groupVersionKindStr)
				}

				items := funk.Map(objs, func(obj *unstructured.Unstructured) string {
					if obj.GetNamespace() == "" {
						return obj.GetName()
					}
					return obj.GetNamespace() + "/" + obj.GetName()
				})
				return fmt.Errorf("%d %s found (%s)", len(objs), groupVersionKindStr, strings.Join(items.([]string), ","))
			}
			return nil
		},
	)
}

// CountNamespacedResources implements the GoDoc step
// - `Kubernetes has <NumberResources> <ApiGroupVersionKind> in namespace '<Namespace>'`
// It compare the current number of a specific resource with the given number.
func CountNamespacedResources(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes has (\d+) (`+RxGroupVersionKind+`) in namespace '(`+RxDNSChar+`+)'$`,
		func(n int, groupVersionKindStr, namespace string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}

			objs, err := ctx.List(groupVersionKind, &client.ListOptions{Namespace: namespace})
			if err != nil {
				return err
			}

			if len(objs) != n {
				if len(objs) == 0 {
					return fmt.Errorf("no %s found in namespace '%s'", groupVersionKindStr, namespace)
				}

				items := funk.Map(objs, func(obj *unstructured.Unstructured) string {
					return obj.GetNamespace() + "/" + obj.GetName()
				})
				return fmt.Errorf("%d %s found in namespace '%s' (%s)", len(objs), groupVersionKindStr, namespace, strings.Join(items.([]string), ","))
			}
			return nil
		},
	)
}
