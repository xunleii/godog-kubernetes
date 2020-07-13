package kubernetes_ctx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	kubernetes_ctx "github.com/xunleii/godog-kubernetes"
)

// initFakeScenario generates a godoc ScenarioContext and a FeatureContext
// with a fake client.
func initFakeScenario(t *testing.T) *kubernetes_ctx.FeatureContext {
	scenarioContextMock := MockScenarioContext()
	ctx, err := kubernetes_ctx.NewEmptyFeatureContext(scenarioContextMock, kubernetes_ctx.WithFakeRuntimeClient())
	require.NoError(t, err)
	scenarioContextMock.RunScenario()
	return ctx
}

// initFakeScenarioWithNamespaces generates a godoc ScenarioContext,
// a FeatureContext with a fake client and with default Kubernetes namespaces.
func initFakeScenarioWithNamespaces(t *testing.T) *kubernetes_ctx.FeatureContext {
	ctx := initFakeScenario(t)

	for _, namespace := range []string{"kube-system", "kube-public", "default"} {
		err := ctx.Create(schema.GroupVersionKind{Version: "v1", Kind: "Namespace"}, types.NamespacedName{Name: namespace}, &unstructured.Unstructured{})
		require.NoError(t, err)
	}
	return ctx
}

// yamlToUnstructured returns an unmarshalled unstructured.Unstructured based
// on the given YAML.
func yamlToUnstructured(t *testing.T, rawYaml string) *unstructured.Unstructured {
	obj := &unstructured.Unstructured{}
	require.NoError(t, yaml.Unmarshal([]byte(rawYaml), &obj.Object))
	return obj
}
