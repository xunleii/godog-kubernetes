package kubernetes_ctx

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// FeatureContextOptionFnc wraps a function to implement
// the FeatureContextOption interface.
type FeatureContextOptionFnc func(*FeatureContext)

func (fnc FeatureContextOptionFnc) ApplyToFeatureContext(ctx *FeatureContext) {
	if fnc != nil {
		fnc(ctx)
	}
}

// WithClient inject the given client inside the Kubernetes
// feature context.
func WithClient(scheme Scheme, client client.Client) FeatureContextOptionFnc {
	return func(ctx *FeatureContext) {
		ctx.scheme = scheme
		ctx.client = client
	}
}

// WithFakeClient instantiate a new Kubernetes client
// with the given scheme. It automatically inject the
// NaiveGC as garbage collector if any is provided.
func WithFakeClient(scheme *runtime.Scheme) FeatureContextOptionFnc {
	return func(ctx *FeatureContext) {
		ctx.scheme = scheme
		ctx.client = fake.NewFakeClientWithScheme(scheme)
		if ctx.gc == nil {
			ctx.gc = NaiveGC
		}
	}
}

// WithFakeRuntimeClient instantiate a new Kubernetes client
// with the default scheme. It automatically inject the
// NaiveGC as garbage collector if any is provided.
func WithFakeRuntimeClient() FeatureContextOptionFnc {
	return WithFakeClient(scheme.Scheme)
}

// WithContext inject a context to the Kubernetes feature context, used
// during the client's call.
func WithContext(goctx context.Context) FeatureContextOptionFnc {
	return func(ctx *FeatureContext) { ctx.ctx = goctx }
}

// WithCustomGarbageCollector inject the given GarbageCollector
// to the feature context. This garbage collector is used to
// Delete children objects when a parent is removed. It is
// required because the fake client doesn't implement it.
func WithCustomGarbageCollector(gc func(*FeatureContext, *unstructured.Unstructured) error) FeatureContextOptionFnc {
	return func(ctx *FeatureContext) { ctx.gc = gc }
}
