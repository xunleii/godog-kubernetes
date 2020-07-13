package kubernetes_ctx

import (
	"fmt"

	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/xunleii/godog-kubernetes/helpers"
)

// ResourceExists implements the GoDoc step
// - `Kubernetes has <ApiGroupVersionKind> '<NamespacedName>'`
// It validates the fact that Kubernetes has the specified resource.
func ResourceExists(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes has (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, name string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, _ := helpers.NamespacedNameFrom(name)

			_, err = ctx.Get(groupVersionKind, namespacedName)
			return err
		},
	)
}

// ResourceNotExists implements the GoDoc step
// - `Kubernetes doesn't have <ApiGroupVersionKind> '<NamespacedName>'`
// It validates the fact that Kubernetes doesn't have the specified resource.
func ResourceNotExists(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes doesn't have (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, name string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, _ := helpers.NamespacedNameFrom(name)

			_, err = ctx.Get(groupVersionKind, namespacedName)
			switch {
			case errors.IsNotFound(err):
				return nil
			case err != nil:
				return err
			default:
				return fmt.Errorf(`%s "%s" found`, groupVersionKindStr, name)
			}
		},
	)
}

// ResourceIsSimilarTo implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is similar to '<NamespacedName>'`
// It compares two resources in order to determine if they are similar.
//
// NOTE: Two resources are similar if all fields except 'medatata' are the same.
func ResourceIsSimilarTo(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' is similar to '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, lname, rname string) (err error) {
			lobj, err := getWithoutMetadata(ctx, groupVersionKindStr, lname)
			if err != nil {
				return err
			}

			robj, err := getWithoutMetadata(ctx, groupVersionKindStr, rname)
			if err != nil {
				return err
			}

			diff := diffObjects(lobj, robj)
			switch {
			case diff != "":
				return fmt.Errorf("resources %s '%s' and '%s' are not similar: %s", groupVersionKindStr, lname, rname, diff)
			}
			return nil
		},
	)
}

// ResourceIsNotSimilarTo implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is not similar to '<NamespacedName>'`
// It compares two resources in order to determine if they are not similar.
//
// NOTE: Two resources are similar if all fields except 'medatata' are the same.
func ResourceIsNotSimilarTo(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' is not similar to '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, lname, rname string) (err error) {
			lobj, err := getWithoutMetadata(ctx, groupVersionKindStr, lname)
			if err != nil {
				return err
			}

			robj, err := getWithoutMetadata(ctx, groupVersionKindStr, rname)
			if err != nil {
				return err
			}

			diff := gojsondiff.New().CompareObjects(lobj.Object, robj.Object)
			if !diff.Modified() {
				return fmt.Errorf("resources %s '%s' and '%s' are similar", groupVersionKindStr, lname, rname)
			}
			return nil
		},
	)
}

// getWithoutMetadata returns resource without metadata field.
func getWithoutMetadata(ctx *FeatureContext, groupVersionKindStr, name string) (*unstructured.Unstructured, error) {
	groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
	if err != nil {
		return nil, err
	}

	namespacedName, _ := helpers.NamespacedNameFrom(name)

	obj, err := ctx.Get(groupVersionKind, namespacedName)
	if err != nil {
		return nil, err
	}
	delete(obj.Object, "metadata")

	return obj, nil
}

// ResourceIsEqualTo implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is equal to '<NamespacedName>'`
// It compares two resources in order to determine if they are equal.
//
// NOTE: Two resources are equal if all fields except unique fields ('metadata.name',
//       'metadata.namespace', 'metadata.uid' and 'metadata.resourceVersion') are the same.
func ResourceIsEqualTo(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' is equal to '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, lname, rname string) (err error) {
			lobj, err := getWithoutUniqueFields(ctx, groupVersionKindStr, lname)
			if err != nil {
				return err
			}

			robj, err := getWithoutUniqueFields(ctx, groupVersionKindStr, rname)
			if err != nil {
				return err
			}

			diff := diffObjects(lobj, robj)
			switch {
			case diff != "":
				return fmt.Errorf("resources %s '%s' and '%s' are not equal: %s", groupVersionKindStr, lname, rname, diff)
			}
			return nil
		},
	)
}

// ResourceIsNotEqualTo implements the GoDoc step
// - `Kubernetes resource <ApiGroupVersionKind> '<NamespacedName>' is not equal to '<NamespacedName>'`
// It compares two resources in order to determine if they are not equal.
//
// NOTE: Two resources are equal if all fields except unique fields ('metadata.name',
//       'metadata.namespace', 'metadata.uid' and 'metadata.resourceVersion') are the same.
func ResourceIsNotEqualTo(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes resource (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)' is not equal to '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, lname, rname string) (err error) {
			lobj, err := getWithoutUniqueFields(ctx, groupVersionKindStr, lname)
			if err != nil {
				return err
			}

			robj, err := getWithoutUniqueFields(ctx, groupVersionKindStr, rname)
			if err != nil {
				return err
			}

			diff := gojsondiff.New().CompareObjects(lobj.Object, robj.Object)
			if !diff.Modified() {
				return fmt.Errorf("resources %s '%s' and '%s' are equal", groupVersionKindStr, lname, rname)
			}
			return nil
		},
	)
}

// getWithoutUniqFields returns resources without unique fields ('metadata.name',
// 'metadata.namespace', 'metadata.uid' and 'metadata.resourceVersion').
func getWithoutUniqueFields(ctx *FeatureContext, groupVersionKindStr, name string) (*unstructured.Unstructured, error) {
	groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
	if err != nil {
		return nil, err
	}

	namespacedName, _ := helpers.NamespacedNameFrom(name)

	obj, err := ctx.Get(groupVersionKind, namespacedName)
	if err != nil {
		return nil, err
	}
	obj.SetName("")
	obj.SetNamespace("")
	obj.SetUID("")
	obj.SetResourceVersion("")

	return obj, nil
}

// diffObjects return a readable diff if the given objects are different.
func diffObjects(lobj, robj *unstructured.Unstructured) string {
	diff := gojsondiff.New().CompareObjects(lobj.Object, robj.Object)
	if diff.Modified() {
		outDiff, _ := formatter.
			NewAsciiFormatter(lobj.Object, formatter.AsciiFormatterConfig{Coloring: false, ShowArrayIndex: true}).
			Format(diff)

		return fmt.Sprintf("\n%s", outDiff)
	}
	return ""
}
