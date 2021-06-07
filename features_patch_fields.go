package kubernetes_ctx

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"

	"github.com/xunleii/godog-kubernetes/helpers"
)

// LabelizeResource implements the GoDoc step
// - `Kubernetes labelizes <ApiGroupVersionKind> '<NamespacedName>' with '<LabelName>=<LabelValue>'`
// It adds or modifies a specific resource label with the given value.
func LabelizeResource(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes labelizes (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' with '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, resourceName, labelName, labelValue string) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, _ := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)

			patch := fmt.Sprintf(`{"metadata":{"labels":{"%s":"%s"}}}`, labelName, labelValue)
			return ctx.Patch(groupVersionKind, namespacedName, types.MergePatchType, []byte(patch))
		},
	)
}

// RemoveResourceLabel implements the GoDoc step
// - `Kubernetes removes label <LabelName> on <ApiGroupVersionKind> '<NamespacedName>'`
// It removes the given label on the specified resource.
func RemoveResourceLabel(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes removes label '(`+RxFieldPath+`)' on (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)'$`,
		func(label, groupVersionKindStr, resourceName string) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, _ := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)

			patch := fmt.Sprintf(`[{"op":"remove","path":"/metadata/labels/%s"}]`, kubernetes_ctx_helpers.SanitizeJsonPatch(label))
			return ctx.Patch(groupVersionKind, namespacedName, types.JSONPatchType, []byte(patch))
		},
	)
}

// UpdateResourceLabel implements the GoDoc step
// - `Kubernetes updates label <LabelName> on <ApiGroupVersionKind> '<NamespacedName>' with '<LabelValue>'`
// It updates the given label on the specified resource with the given value.
func UpdateResourceLabel(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes updates label '(`+RxFieldPath+`)' on (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' with '(.*)'$`,
		func(label, groupVersionKindStr, resourceName, value string) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, _ := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)

			obj, err := ctx.Get(groupVersionKind, namespacedName)
			if err != nil {
				return err
			}
			if _, exists := obj.GetLabels()[kubernetes_ctx_helpers.SanitizeJsonPatch(label)]; !exists {
				return fmt.Errorf("label `%s` not found on %s %s", kubernetes_ctx_helpers.SanitizeJsonPatch(label), groupVersionKind, namespacedName)
			}

			patch := fmt.Sprintf(`[{"op":"replace","path":"/metadata/labels/%s","value":"%s"}]`, kubernetes_ctx_helpers.SanitizeJsonPatch(label), value)
			return ctx.Patch(groupVersionKind, namespacedName, types.JSONPatchType, []byte(patch))
		},
	)
}

// AnnotateResource implements the GoDoc step
// - `Kubernetes annotates <ApiGroupVersionKind> '<NamespacedName>' with '<AnnotationName>=<AnnotationValue>'`
// It adds or modifies a specific resource annotation with the given value.
func AnnotateResource(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes annotates (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' with '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, resourceName, annotationName, annotationValue string) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)
			if err != nil {
				return err
			}

			patch := fmt.Sprintf(`{"metadata":{"annotations":{"%s":"%s"}}}`, annotationName, annotationValue)
			return ctx.Patch(groupVersionKind, namespacedName, types.MergePatchType, []byte(patch))
		},
	)
}

// RemoveResourceAnnotation implements the GoDoc step
// - `Kubernetes removes annotation <AnnotationName> on <ApiGroupVersionKind> '<NamespacedName>'`
// It removes the given annotation on the specified resource.
func RemoveResourceAnnotation(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes removes annotation '(`+RxFieldPath+`)' on (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)'$`,
		func(annotation, groupVersionKindStr, resourceName string) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)
			if err != nil {
				return err
			}

			patch := fmt.Sprintf(`[{"op":"remove","path":"/metadata/annotations/%s"}]`, kubernetes_ctx_helpers.SanitizeJsonPatch(annotation))
			return ctx.Patch(groupVersionKind, namespacedName, types.JSONPatchType, []byte(patch))
		},
	)
}

// UpdateResourceAnnotation implements the GoDoc step
// - `Kubernetes updates annotation <AnnotationName> on <ApiGroupVersionKind> '<NamespacedName>' with '<AnnotationValue>'`
// It updates the given annotation on the specified resource with the given value.
func UpdateResourceAnnotation(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes updates annotation '(`+RxFieldPath+`)' on (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' with '(.*)'$`,
		func(annotation, groupVersionKindStr, resourceName, value string) error {
			groupVersionKind, err := kubernetes_ctx_helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := kubernetes_ctx_helpers.NamespacedNameFrom(resourceName)
			if err != nil {
				return err
			}

			obj, err := ctx.Get(groupVersionKind, namespacedName)
			if err != nil {
				return err
			}
			if _, exists := obj.GetAnnotations()[kubernetes_ctx_helpers.SanitizeJsonPatch(annotation)]; !exists {
				return fmt.Errorf("annotation `%s` not found on %s %s", kubernetes_ctx_helpers.SanitizeJsonPatch(annotation), groupVersionKind, namespacedName)
			}

			patch := fmt.Sprintf(`[{"op":"replace","path":"/metadata/annotations/%s","value":"%s"}]`, kubernetes_ctx_helpers.SanitizeJsonPatch(annotation), value)
			return ctx.Patch(groupVersionKind, namespacedName, types.JSONPatchType, []byte(patch))
		},
	)
}
