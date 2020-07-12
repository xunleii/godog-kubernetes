package kubernetes_ctx

import (
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Create creates a Kubernetes resource based on the given APIVersion/Kind,
// the name and the object definition itself. It allows us to easily manage
// all resources through Unstructured object with the "official" Kubernetes
// client.Client interface.
func (ctx *FeatureContext) Create(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
	obj *unstructured.Unstructured,
	opts ...client.CreateOption,
) error {
	obj.SetGroupVersionKind(groupVersionKind)
	obj.SetUID(types.UID(uuid.New().String()))
	obj.SetName(namespacedName.Name)
	obj.SetNamespace(namespacedName.Namespace)

	kobj, err := ctx.scheme.New(groupVersionKind)
	if err != nil {
		return err
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, kobj)
	if err != nil {
		return err
	}
	return ctx.client.Create(ctx.ctx, kobj, opts...)
}

// Get fetches the Kubernetes resource using the given APIVersion/Kind and
// the name. It wraps the Get method of the "official" Kubernetes client.Client
// interface, but returns an Unstructured object, more easier to use.
func (ctx *FeatureContext) Get(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
) (*unstructured.Unstructured, error) {
	kobj, err := ctx.get(groupVersionKind, namespacedName)
	if err != nil {
		return nil, err
	}

	obj := unstructured.Unstructured{}
	obj.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(kobj)
	return &obj, err
}

// Get fetches the Kubernetes resource using the given APIVersion/Kind and
// the name. It wraps the Get method of the "official" Kubernetes client.Client
// interface, and returns a runtime.Object.
func (ctx *FeatureContext) get(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
) (runtime.Object, error) {
	kobj, err := ctx.scheme.New(groupVersionKind)
	if err != nil {
		return nil, err
	}

	err = ctx.client.Get(ctx.ctx, namespacedName, kobj)
	if err != nil {
		return nil, err
	}

	return kobj, nil
}

// List returns all Kubernetes resources based on the given APIVersion/Kind. It
// returns a List of Unstructured object, more easier to use.
func (ctx *FeatureContext) List(
	groupVersionKind schema.GroupVersionKind,
	opts ...client.ListOption,
) ([]*unstructured.Unstructured, error) {
	// NOTE: can be dangerous but seems working...
	groupVersionKind.Kind += "List"

	kobj, err := ctx.scheme.New(groupVersionKind)
	if err != nil {
		return nil, err
	}

	err = ctx.client.List(ctx.ctx, kobj, opts...)
	if err != nil {
		return nil, err
	}

	list := unstructured.Unstructured{}
	list.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(kobj)
	if err != nil {
		return nil, err
	}

	var objs []*unstructured.Unstructured
	return objs, list.EachListItem(func(object runtime.Object) error {
		objs = append(objs, object.(*unstructured.Unstructured))
		return nil
	})
}

// Update updates a Kubernetes resource based on the given APIVersion/Kind
// and the name with the given Unstructured object.
func (ctx *FeatureContext) Update(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
	obj *unstructured.Unstructured,
	opts ...client.UpdateOption,
) error {
	obj.SetGroupVersionKind(groupVersionKind)
	obj.SetName(namespacedName.Name)
	obj.SetNamespace(namespacedName.Namespace)

	kobj, err := ctx.scheme.New(obj.GroupVersionKind())
	if err != nil {
		return err
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, kobj)
	if err != nil {
		return err
	}

	return ctx.client.Update(ctx.ctx, obj, opts...)
}

// Patch patches a Kubernetes resource based on the given APIVersion/Kind
// and the name with the given Patch value.
func (ctx *FeatureContext) Patch(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
	pt types.PatchType,
	data []byte,
) error {
	obj, err := ctx.get(groupVersionKind, namespacedName)
	if err != nil {
		return err
	}

	return ctx.client.Patch(ctx.ctx, obj, client.RawPatch(pt, data))
}

// Delete deletes a Kubernetes resource based on the given APIVersion/Kind
// and the name, and returns the removed object. If a garbage collector is
// set to the context, it will call it on the removed resource.
func (ctx *FeatureContext) Delete(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
) (*unstructured.Unstructured, error) {
	obj, err := ctx.DeleteWithoutGC(groupVersionKind, namespacedName)
	if err != nil {
		return nil, err
	}

	return obj, ctx.callGC(obj)
}

// DeleteWithoutGC deletes a Kubernetes resource based on the given
// APIVersion/Kind and the name, and returns the removed object.
// Therefore, it never calls the garbage collector.
func (ctx *FeatureContext) DeleteWithoutGC(
	groupVersionKind schema.GroupVersionKind,
	namespacedName types.NamespacedName,
) (*unstructured.Unstructured, error) {
	kobj, err := ctx.get(groupVersionKind, namespacedName)
	if err != nil {
		return nil, err
	}

	obj := unstructured.Unstructured{}
	obj.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(kobj)
	if err != nil {
		return nil, err
	}

	return &obj, ctx.client.Delete(ctx.ctx, kobj)
}
