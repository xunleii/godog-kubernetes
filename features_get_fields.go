package kubernetes_ctx

import (
	"fmt"

	"github.com/stretchr/objx"

	"github.com/xunleii/godog-kubernetes/helpers"
)

// ResourceHasField implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has '<FieldPath>'`
// It validates the fact that the specific resource has the field.
func ResourceHasField(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' has '(`+RxFieldPath+`)'$`,
		func(groupVersionKindStr, name, field string) (err error) {
			_, exists, err := getResourceField(ctx, groupVersionKindStr, name, field)
			switch {
			case err != nil:
				return err
			case !exists:
				return fmt.Errorf("field '%s' not found", field)
			}
			return nil
		},
	)
}

// ResourceDoesntHaveField implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' doesn't have '<FieldPath>'`
// It validates the fact that the specific resource doesn't have the field.
func ResourceDoesntHaveField(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' doesn't have '(`+RxFieldPath+`)'$`,
		func(groupVersionKindStr, name, field string) (err error) {
			_, exists, err := getResourceField(ctx, groupVersionKindStr, name, field)
			switch {
			case err != nil:
				return err
			case exists:
				return fmt.Errorf("field '%s' found", field)
			}
			return nil
		},
	)
}

// ResourceHasFieldEqual implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has '<FieldPath>=<FieldValue>'`
// It validates the fact that the specific resource field has the given value.
func ResourceHasFieldEqual(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' has '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, name, field, value string) (err error) {
			rval, exists, err := getResourceField(ctx, groupVersionKindStr, name, field)
			switch {
			case err != nil:
				return err
			case !exists:
				return fmt.Errorf("field '%s' not found", field)
			case rval != value:
				return fmt.Errorf("field '%s' not equal to %s (current: %s)", field, value, rval)
			}
			return nil
		},
	)
}

// ResourceHasFieldNotEqual implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has '<FieldPath>!=<FieldValue>'`
// It validates the fact that the specific resource field is different than the given value.
func ResourceHasFieldNotEqual(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' doesn't have '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, name, field, value string) (err error) {
			rval, exists, err := getResourceField(ctx, groupVersionKindStr, name, field)
			switch {
			case err != nil:
				return err
			case !exists:
				return nil
			case rval == value:
				return fmt.Errorf("field '%s' equal to %s", field, value)
			}
			return nil
		},
	)
}

// getResourceField returns the resource field value and if it exists.
func getResourceField(ctx *FeatureContext, groupVersionKindStr, name, field string) (string, bool, error) {
	groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
	if err != nil {
		return "", false, err
	}
	namespacedName, _ := helpers.NamespacedNameFrom(name)

	obj, err := ctx.Get(groupVersionKind, namespacedName)
	if err != nil {
		return "", false, err
	}

	xmap := objx.Map(obj.Object)
	if xmap.Has(field) {
		return xmap.Get(field).String(), true, nil
	}
	return "", false, nil
}

// ResourceHasLabel implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has label '<LabelName>'`
// It validates the fact that the specific resource has the given label.
func ResourceHasLabel(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' has label '(`+RxFieldPath+`)'$`,
		func(groupVersionKindStr, name, label string) (err error) {
			_, exists, err := getResourceLabel(ctx, groupVersionKindStr, name, label)
			switch {
			case err != nil:
				return err
			case !exists:
				return fmt.Errorf("label '%s' not found", label)
			}
			return nil
		},
	)
}

// ResourceDoesntHaveLabel implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' doesn't have label '<LabelName>'`
// It validates the fact that the specific resource doesn't have the given label.
func ResourceDoesntHaveLabel(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' doesn't have label '(`+RxFieldPath+`)'$`,
		func(groupVersionKindStr, name, label string) (err error) {
			_, exists, err := getResourceLabel(ctx, groupVersionKindStr, name, label)
			switch {
			case err != nil && err.Error() != "no label found":
				return err
			case err == nil && exists:
				return fmt.Errorf("label '%s' found", label)
			}
			return nil
		},
	)
}

// ResourceHasLabelEqual implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has label '<LabelName>=<LabelValue>'`
// It validates the fact that the specific resource label has the given value.
func ResourceHasLabelEqual(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' has label '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, name, label, value string) (err error) {
			rval, exists, err := getResourceLabel(ctx, groupVersionKindStr, name, label)
			switch {
			case err != nil:
				return err
			case !exists:
				return fmt.Errorf("label '%s' not found", label)
			case rval != value:
				return fmt.Errorf("label '%s' not equal to %s (current: %s)", label, value, rval)
			}
			return nil
		},
	)
}

// ResourceHasLabelNotEqual implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has label '<LabelName>!=<LabelValue>'`
// It validates the fact that the specific resource label doesn't have the given value.
func ResourceHasLabelNotEqual(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' doesn't have label '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, name, label, value string) (err error) {
			rval, exists, err := getResourceLabel(ctx, groupVersionKindStr, name, label)
			switch {
			case err != nil:
				return err
			case !exists:
				return nil
			case rval == value:
				return fmt.Errorf("label '%s' equal to %s", label, value)
			}
			return nil
		},
	)
}

// getResourceLabel returns the resource label value and if it exists.
func getResourceLabel(ctx *FeatureContext, groupVersionKindStr, name, label string) (string, bool, error) {
	groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
	if err != nil {
		return "", false, err
	}
	namespacedName, _ := helpers.NamespacedNameFrom(name)

	obj, err := ctx.Get(groupVersionKind, namespacedName)
	if err != nil {
		return "", false, err
	}

	labels := obj.GetLabels()
	if labels == nil {
		return "", false, fmt.Errorf("no label found")
	}

	value, exists := labels[label]
	return value, exists, nil
}

// ResourceHasAnnotation implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has annotation '<AnnotationName>'`
// It validates the fact that the specific resource has the given annotation.
func ResourceHasAnnotation(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' has annotation '(`+RxFieldPath+`)'$`,
		func(groupVersionKindStr, name, annotation string) (err error) {
			_, exists, err := getResourceAnnotation(ctx, groupVersionKindStr, name, annotation)
			switch {
			case err != nil:
				return err
			case !exists:
				return fmt.Errorf("annotation '%s' not found", annotation)
			}
			return nil
		},
	)
}

// ResourceDoesntHaveAnnotation implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' doesn't have annotation '<AnnotationName>'`
// It validates the fact that the specific resource doesn't have the given annotation.
func ResourceDoesntHaveAnnotation(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' doesn't have annotation '(`+RxFieldPath+`)'$`,
		func(groupVersionKindStr, name, annotation string) (err error) {
			_, exists, err := getResourceAnnotation(ctx, groupVersionKindStr, name, annotation)
			switch {
			case err != nil && err.Error() != "no annotation found":
				return err
			case err == nil && exists:
				return fmt.Errorf("annotation '%s' found", annotation)
			}
			return nil
		},
	)
}

// ResourceHasAnnotationEqual implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has annotation '<AnnotationName>=<AnnotationValue>'`
// It validates the fact that the specific resource annotation has the given value.
func ResourceHasAnnotationEqual(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' has annotation '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, name, annotation, value string) (err error) {
			rval, exists, err := getResourceAnnotation(ctx, groupVersionKindStr, name, annotation)
			switch {
			case err != nil:
				return err
			case !exists:
				return fmt.Errorf("annotation '%s' not found", annotation)
			case rval != value:
				return fmt.Errorf("annotation '%s' not equal to %s (current: %s)", annotation, value, rval)
			}
			return nil
		},
	)
}

// ResourceHasAnnotationNotEqual implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' has annotation '<AnnotationName>!=<AnnotationValue>'`
// It validates the fact that the specific resource annotation doesn't have the given value.
func ResourceHasAnnotationNotEqual(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' doesn't have annotation '(`+RxFieldPath+`)=(.*)'$`,
		func(groupVersionKindStr, name, annotation, value string) (err error) {
			rval, exists, err := getResourceAnnotation(ctx, groupVersionKindStr, name, annotation)
			switch {
			case err != nil:
				return err
			case !exists:
				return nil
			case rval == value:
				return fmt.Errorf("annotation '%s' equal to %s", annotation, value)
			}
			return nil
		},
	)
}

// getResourceAnnotation returns the resource annotation value and if it exists.
func getResourceAnnotation(ctx *FeatureContext, groupVersionKindStr, name, annotation string) (string, bool, error) {
	groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
	if err != nil {
		return "", false, err
	}
	namespacedName, _ := helpers.NamespacedNameFrom(name)

	obj, err := ctx.Get(groupVersionKind, namespacedName)
	if err != nil {
		return "", false, err
	}

	annotations := obj.GetAnnotations()
	if annotations == nil {
		return "", false, fmt.Errorf("no annotation found")
	}

	value, exists := annotations[annotation]
	return value, exists, nil
}
