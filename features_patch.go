package kubernetes_ctx

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/xunleii/godog-kubernetes/helpers"
)

// PatchResourceWith implements the GoDoc step
// - `Kubernetes patches <ApiGroupVersionKind> '<NamespacedName>' with <YAML>`
// It patches a specific resource with the given patch (it use StrategicMergePatchType...
// see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/strategic-merge-patch.md
// for more information).
func PatchResourceWith(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes patches (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' with$`,
		func(groupVersionKindStr, resourceName string, content kubernetes_ctx_helpers.YamlDocString) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, _ := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)

			patch, err := kubernetes_ctx_helpers.YamlToJson(content.Content)
			if err != nil {
				return err
			}

			return ctx.Patch(groupVersionKind, namespacedName, types.MergePatchType, patch)
		},
	)
}
