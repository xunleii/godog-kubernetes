package kubernetes_ctx

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	RxDNSChar          = `[a-z0-9\-.]`
	RxGroupVersionKind = `[\w/]+`
	RxNamespacedName   = RxDNSChar + `+(?:/` + RxDNSChar + `+)?`
	RxFieldPath        = `[^=:]+?`
)

type (
	// FeatureContext extends GoDoc in order to implement step definitions
	// for Kubernetes.
	FeatureContext struct {
		ctx    context.Context
		scheme Scheme
		client client.Client

		gc func(*FeatureContext, *unstructured.Unstructured) error
	}

	// FeatureContextOption is some configuration that modifies options for
	// the Kubernetes feature context.
	FeatureContextOption interface {
		ApplyToFeatureContext(*FeatureContext)
	}

	// Scheme abstracts the implementation of common operations on objects.
	Scheme interface {
		runtime.ObjectCreater
		runtime.ObjectTyper
	}
)

// NewFeatureContext returns a new instance of the Kubernetes feature context,
// with all step injected. If you want to choose which steps must be enabled,
// use NewEmptyFeatureContext instead.
func NewFeatureContext(s ScenarioContext, opts ...FeatureContextOption) (*FeatureContext, error) {
	ctx, err := NewEmptyFeatureContext(s, opts...)
	if err != nil {
		return nil, err
	}

	CreateSingleResource(ctx, s)
	CreateSingleResourceWith(ctx, s)
	CreateSingleResourceFrom(ctx, s)
	CreateMultiResources(ctx, s)

	ResourceExists(ctx, s)
	ResourceNotExists(ctx, s)
	ResourceIsSimilarTo(ctx, s)
	ResourceIsNotSimilarTo(ctx, s)
	ResourceIsEqualTo(ctx, s)
	ResourceIsNotEqualTo(ctx, s)
	ResourceHasField(ctx, s)
	ResourceDoesntHaveField(ctx, s)
	ResourceHasFieldEqual(ctx, s)
	ResourceHasFieldNotEqual(ctx, s)
	ResourceHasLabel(ctx, s)
	ResourceDoesntHaveLabel(ctx, s)
	ResourceHasLabelEqual(ctx, s)
	ResourceHasLabelNotEqual(ctx, s)
	ResourceHasAnnotation(ctx, s)
	ResourceDoesntHaveAnnotation(ctx, s)
	ResourceHasAnnotationEqual(ctx, s)
	ResourceHasAnnotationNotEqual(ctx, s)

	CountResources(ctx, s)
	CountNamespacedResources(ctx, s)

	PatchResourceWith(ctx, s)
	LabelizeResource(ctx, s)
	RemoveResourceLabel(ctx, s)
	UpdateResourceLabel(ctx, s)
	AnnotateResource(ctx, s)
	RemoveResourceAnnotation(ctx, s)
	UpdateResourceAnnotation(ctx, s)

	RemoveResource(ctx, s)
	RemoveMultiResource(ctx, s)

	return ctx, nil
}

// NewEmptyFeatureContext returns a new instance of the Kubernetes feature context,
// without any step injected.
func NewEmptyFeatureContext(s ScenarioContext, opts ...FeatureContextOption) (*FeatureContext, error) {
	// preflight checks
	dummy := &FeatureContext{}
	for _, opt := range opts {
		opt.ApplyToFeatureContext(dummy)
	}
	switch {
	case dummy.client == nil:
		return nil, fmt.Errorf("kubernetes client must be instanciated")
	}

	ctx := &FeatureContext{ctx: context.TODO()}
	s.BeforeScenario(func(*godog.Scenario) {
		for _, opt := range opts {
			opt.ApplyToFeatureContext(ctx)
		}
	})

	return ctx, nil
}

// Client returns the Kubernetes client used by the FeatureContext.
func (ctx FeatureContext) Client() client.Client { return ctx.client }

// Scheme returns the Kubernetes scheme used by the FeatureContext.
func (ctx FeatureContext) Scheme() Scheme { return ctx.scheme }

// GoContext returns the golang context used by the FeatureContext.
func (ctx FeatureContext) GoContext() context.Context { return ctx.ctx }

// GarbageCollector returns the garbage collector implementation used
// by the FeatureContext.
func (ctx FeatureContext) GarbageCollector() func(*FeatureContext, *unstructured.Unstructured) error {
	return ctx.gc
}

func (ctx *FeatureContext) callGC(obj *unstructured.Unstructured) error {
	if ctx.gc == nil {
		return nil
	}
	return ctx.gc(ctx, obj)
}
