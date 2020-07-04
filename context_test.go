package kubernetes_ctx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"kubernetes_ctx"
)

var (
	scheme   = runtime.NewScheme()
	client   = fake.NewFakeClientWithScheme(scheme)
	goctx, _ = context.WithCancel(context.TODO())
	gc       = func(*kubernetes_ctx.FeatureContext, *unstructured.Unstructured) error { return nil }
)

func TestNewFeatureContext(t *testing.T) {
	scenarioCtx := MockScenarioContext()

	ctx, err := kubernetes_ctx.NewFeatureContext(scenarioCtx, kubernetes_ctx.WithFakeRuntimeClient())
	require.NoError(t, err)
	assert.NotNil(t, ctx)

	assert.Len(t, scenarioCtx.beforeScenarioList, 1)
	assert.Len(t, scenarioCtx.stepList, 0) // NOTE: Do not forget to update this value
}

func TestNewEmptyFeatureContext(t *testing.T) {
	scenarioCtx := MockScenarioContext()

	ctx, err := kubernetes_ctx.NewEmptyFeatureContext(scenarioCtx, kubernetes_ctx.WithFakeRuntimeClient())
	require.NoError(t, err)
	assert.NotNil(t, ctx)

	assert.Len(t, scenarioCtx.beforeScenarioList, 1)
	assert.Len(t, scenarioCtx.stepList, 0)
}

func TestFeatureContextOption_Builders(t *testing.T) {
	tests := map[string]struct {
		opts   []kubernetes_ctx.FeatureContextOption
		assert func(t *testing.T, ctx *kubernetes_ctx.FeatureContext)
	}{
		"WithClient": {
			opts: []kubernetes_ctx.FeatureContextOption{kubernetes_ctx.WithClient(scheme, client)},
			assert: func(t *testing.T, ctx *kubernetes_ctx.FeatureContext) {
				// value only initialized during BeforeScenario
				assert.Equal(t, nil, ctx.Scheme())
				assert.Equal(t, nil, ctx.Client())
			},
		},
		"WithFakeClient": {
			opts: []kubernetes_ctx.FeatureContextOption{kubernetes_ctx.WithFakeClient(scheme)},
			assert: func(t *testing.T, ctx *kubernetes_ctx.FeatureContext) {
				// value only initialized during BeforeScenario
				assert.Equal(t, nil, ctx.Scheme())
				assert.Equal(t, nil, ctx.Client())
			},
		},
		"WithFakeRuntimeClient": {
			opts: []kubernetes_ctx.FeatureContextOption{kubernetes_ctx.WithFakeRuntimeClient()},
			assert: func(t *testing.T, ctx *kubernetes_ctx.FeatureContext) {
				// value only initialized during BeforeScenario
				assert.Equal(t, nil, ctx.Scheme())
				assert.Equal(t, nil, ctx.Client())
			},
		},
		"WithContext": {
			opts: []kubernetes_ctx.FeatureContextOption{
				kubernetes_ctx.WithFakeRuntimeClient(), // client required
				kubernetes_ctx.WithContext(goctx),
			},
			assert: func(t *testing.T, ctx *kubernetes_ctx.FeatureContext) {
				// value only initialized during BeforeScenario
				assert.Equal(t, context.TODO(), ctx.GoContext())
			},
		},
		"WithCustomGarbageCollector": {
			opts: []kubernetes_ctx.FeatureContextOption{
				kubernetes_ctx.WithFakeRuntimeClient(), // client required
				kubernetes_ctx.WithCustomGarbageCollector(gc),
			},
			assert: func(t *testing.T, ctx *kubernetes_ctx.FeatureContext) {
				// value only initialized during BeforeScenario
				assert.Nil(t, ctx.GarbageCollector())
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			scenarioCtx := MockScenarioContext()

			ctx, err := kubernetes_ctx.NewEmptyFeatureContext(scenarioCtx, tt.opts...)
			require.NoError(t, err)
			tt.assert(t, ctx)
		})
	}
}

func TestFeatureContext_Validation(t *testing.T) {
	scenarioCtx := MockScenarioContext()

	_, err := kubernetes_ctx.NewFeatureContext(scenarioCtx)
	assert.Error(t, err)
}

func TestFeatureContext_OnBeforeScenarioInit(t *testing.T) {
	scenarioCtx := MockScenarioContext()

	ctx, err := kubernetes_ctx.NewEmptyFeatureContext(
		scenarioCtx,
		kubernetes_ctx.WithClient(scheme, client),
		kubernetes_ctx.WithContext(goctx),
		kubernetes_ctx.WithCustomGarbageCollector(gc),
	)
	require.NoError(t, err)
	assert.Equal(t, nil, ctx.Client())
	assert.Equal(t, nil, ctx.Scheme())
	assert.Equal(t, context.TODO(), ctx.GoContext())
	assert.Nil(t, ctx.GarbageCollector())

	// call the before scenario initialisation
	scenarioCtx.RunScenario()
	assert.Equal(t, client, ctx.Client())
	assert.Equal(t, scheme, ctx.Scheme())
	assert.Equal(t, goctx, ctx.GoContext())
	assert.NotNil(t, ctx.GarbageCollector())
}
