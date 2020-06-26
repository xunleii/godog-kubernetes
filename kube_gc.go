package kubernetes_ctx

import (
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// NaiveGC performs a manual and naive garbage collector using the given object as
// owner.
func NaiveGC(ctx *FeatureContext, owner *unstructured.Unstructured) error {
	for kind, ktype := range scheme.Scheme.AllKnownTypes() {
		if !strings.HasSuffix(kind.Kind, "List") {
			// ignore non List
			continue
		}
		if kind.Group == "" && strings.HasPrefix(kind.Kind, "API") {
			// ignore API...List
			continue
		}

		kobj := reflect.New(ktype)
		if _, isRuntimeObject := kobj.Interface().(runtime.Object); !isRuntimeObject {
			continue
		}

		err := ctx.client.List(ctx.ctx, kobj.Interface().(runtime.Object))
		if err != nil {
			return err
		}

		{
			// Pre-check if items are available (this check cost less than unmarshalling the entire object)
			items := kobj.Elem().FieldByName("Items")
			if !items.IsValid() || items.IsZero() {
				continue
			}
		}

		obj := unstructured.Unstructured{}
		obj.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(kobj.Interface().(runtime.Object))
		if err != nil {
			return err
		}

		items, _, err := unstructured.NestedSlice(obj.Object, "items")
		if err != nil {
			continue
		}

		for _, item := range items {
			if _, isObj := item.(map[string]interface{}); !isObj {
				continue
			}

			obj := unstructured.Unstructured{}
			obj.Object = item.(map[string]interface{})
			if len(obj.GetOwnerReferences()) == 0 {
				continue
			}

			// NOTE: how it works on when several owner exists ?
			for _, ownerReference := range obj.GetOwnerReferences() {
				if ownerReference.UID == owner.GetUID() {
					err := ctx.Client().Delete(ctx.ctx, &obj)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
