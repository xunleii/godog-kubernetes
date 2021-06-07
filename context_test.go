package kubernetes_ctx_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	kubernetes_ctx "github.com/xunleii/godog-kubernetes"
)

var (
	scheme   = runtime.NewScheme()
	client   = fake.NewClientBuilder().WithScheme(scheme).Build()
	goctx, _ = context.WithCancel(context.TODO())
	gc       = func(*kubernetes_ctx.FeatureContext, *unstructured.Unstructured) error { return nil }
)

func TestNewFeatureContext(t *testing.T) {
	scenarioCtx := MockScenarioContext()

	ctx, err := kubernetes_ctx.NewFeatureContext(scenarioCtx, kubernetes_ctx.WithFakeRuntimeClient())
	require.NoError(t, err)
	assert.NotNil(t, ctx)

	assert.Len(t, scenarioCtx.beforeScenarioList, 1)
	assert.Len(t, scenarioCtx.stepList, 33) // NOTE: Do not forget to update this value
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

// TestMain allows us to use GoDog tests with the go test framework.
func TestMain(m *testing.M) {
	opts := godog.Options{Output: colors.Colored(os.Stdout)}
	godog.BindFlags("godog.", flag.CommandLine, &opts)
	flag.Parse()

	// ScenarioInitializer defines how all scenarios will be initialized.
	// To easily test this package with GoDog, this initializer will create
	// some Kubernetes resources before running the tests:
	// - 3 Namespaces (default, kube-public & kube-system)
	// - 2 Services (default/default & default/Kubernetes)
	scenarioInitializer := func(scenarioContext *godog.ScenarioContext) {
		ctx, _ := kubernetes_ctx.NewFeatureContext(scenarioContext, kubernetes_ctx.WithFakeRuntimeClient())
		scenarioContext.BeforeScenario(func(sc *godog.Scenario) {
			// create default namespace
			_ = ctx.Create(schema.GroupVersionKind{Version: "v1", Kind: "Namespace"}, types.NamespacedName{Name: "default"}, &unstructured.Unstructured{})
			// create kube-public namespace
			_ = ctx.Create(schema.GroupVersionKind{Version: "v1", Kind: "Namespace"}, types.NamespacedName{Name: "kube-public"}, &unstructured.Unstructured{})
			// create kube-system namespace
			_ = ctx.Create(
				schema.GroupVersionKind{Version: "v1", Kind: "Namespace"},
				types.NamespacedName{Name: "kube-system"},
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"annotations": map[string]interface{}{"key": "value"},
							"labels":      map[string]interface{}{"key": "value"},
						},
					},
				},
			)

			// create service default/default
			_ = ctx.Create(schema.GroupVersionKind{Version: "v1", Kind: "Service"}, types.NamespacedName{Namespace: "default", Name: "default"}, &unstructured.Unstructured{})
			// create service default/Kubernetes
			_ = ctx.Create(
				schema.GroupVersionKind{Version: "v1", Kind: "Service"},
				types.NamespacedName{Namespace: "default", Name: "kubernetes"},
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"annotations": map[string]interface{}{"key": "value"},
							"labels":      map[string]interface{}{"key": "value"},
						},
						"spec": map[string]interface{}{"type": "ClusterIP", "clusterIP": "None"},
					},
				},
			)
		})
	}

	// GoDog test suite for features context
	var status int
	{
		overridedOpts := opts
		overridedOpts.Paths = []string{"features"}
		status = godog.TestSuite{
			Name:                "kubernetes_ctx",
			ScenarioInitializer: scenarioInitializer,
			Options:             &overridedOpts,
		}.Run()
	}

	// GoDog test suite for error management of features context
	{
		overridedOpts := opts
		overridedOpts.Paths = []string{"features_errors"}
		overridedOpts.StopOnFailure = false
		_ = godog.TestSuite{
			Name:                "kubernetes_ctx::errors",
			ScenarioInitializer: scenarioInitializer,
			Options:             &overridedOpts,
		}.Run()
	}

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
