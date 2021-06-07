package kubernetes_ctx

import (
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// NaiveGC performs a manual and naive garbage collector using the given object as
// owner.
func NaiveGC(ctx *FeatureContext, owner *unstructured.Unstructured) error {
	for kind := range scheme.Scheme.AllKnownTypes() {
		if kind.Kind == "List" || !strings.HasSuffix(kind.Kind, "List") {
			// ignore non List
			continue
		}
		if kind.Group == "" && strings.HasPrefix(kind.Kind, "API") {
			// ignore API...List
			continue
		}

		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(kind)
		err := ctx.client.List(ctx.ctx, list)
		if err != nil {
			return err
		}

		err = list.EachListItem(func(object runtime.Object) error {
			obj := object.(*unstructured.Unstructured)
			if len(obj.GetOwnerReferences()) == 0 {
				return nil
			}

			// NOTE: how it works on when several owner exists ?
			for _, ownerReference := range obj.GetOwnerReferences() {
				if ownerReference.UID == owner.GetUID() {
					err := ctx.Client().Delete(ctx.ctx, obj)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}
	}
	return nil
}
