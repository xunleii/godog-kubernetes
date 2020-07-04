package kubernetes_ctx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

var (
	namespaceGVK     = schema.GroupVersionKind{Version: "v1", Kind: "Namespace"}
	namespaceDefault = types.NamespacedName{Name: "default"}
	notFoundGVK      = schema.GroupVersionKind{Version: "v1", Kind: "NotFound"}
)

func TestFeatureContext_Create(t *testing.T) {
	ctx := initFakeScenario(t)

	err := ctx.Create(namespaceGVK, namespaceDefault, &unstructured.Unstructured{})
	require.NoError(t, err)

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{"apiVersion": "v1", "kind": "Namespace"},
	}

	err = ctx.Client().Get(ctx.GoContext(), namespaceDefault, obj)
	require.NoError(t, err)
}

func TestFeatureContext_Create_KindNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	err := ctx.Create(notFoundGVK, namespaceDefault, &unstructured.Unstructured{})
	assert.EqualError(t, err, "no kind \"NotFound\" is registered for version \"v1\" in scheme \"pkg/runtime/scheme.go:101\"")
}

func TestFeatureContext_Create_ResourceAlreadyExists(t *testing.T) {
	ctx := initFakeScenarioWithNamespaces(t)
	err := ctx.Create(namespaceGVK, namespaceDefault, &unstructured.Unstructured{})
	assert.EqualError(t, err, "namespaces \"default\" already exists")
}

func TestFeatureContext_Get(t *testing.T) {
	ctx := initFakeScenarioWithNamespaces(t)

	_, err := ctx.Get(namespaceGVK, namespaceDefault)
	require.NoError(t, err)
}

func TestFeatureContext_Get_KindNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	_, err := ctx.Get(notFoundGVK, namespaceDefault)
	assert.EqualError(t, err, "no kind \"NotFound\" is registered for version \"v1\" in scheme \"pkg/runtime/scheme.go:101\"")
}

func TestFeatureContext_Get_ResourceNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	_, err := ctx.Get(namespaceGVK, namespaceDefault)
	assert.EqualError(t, err, "namespaces \"default\" not found")
}

func TestFeatureContext_List(t *testing.T) {
	ctx := initFakeScenarioWithNamespaces(t)

	objs, err := ctx.List(namespaceGVK)
	require.NoError(t, err)
	assert.Len(t, objs, 3)
}

func TestFeatureContext_List_KindNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	_, err := ctx.List(notFoundGVK)
	assert.EqualError(t, err, "no kind \"NotFoundList\" is registered for version \"v1\" in scheme \"pkg/runtime/scheme.go:101\"")
}

func TestFeatureContext_Update(t *testing.T) {
	ctx := initFakeScenarioWithNamespaces(t)

	obj, _ := ctx.Get(namespaceGVK, namespaceDefault)
	obj.SetLabels(map[string]string{"new-label": ""})

	err := ctx.Update(namespaceGVK, namespaceDefault, obj)
	require.NoError(t, err)

	obj = &unstructured.Unstructured{
		Object: map[string]interface{}{"apiVersion": "v1", "kind": "Namespace"},
	}

	err = ctx.Client().Get(ctx.GoContext(), namespaceDefault, obj)
	require.NoError(t, err)
	assert.Contains(t, obj.GetLabels(), "new-label")
}

func TestFeatureContext_Update_KindNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	err := ctx.Update(notFoundGVK, namespaceDefault, &unstructured.Unstructured{})
	assert.EqualError(t, err, "no kind \"NotFound\" is registered for version \"v1\" in scheme \"pkg/runtime/scheme.go:101\"")
}

func TestFeatureContext_Update_ResourceNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	err := ctx.Update(namespaceGVK, namespaceDefault, &unstructured.Unstructured{})
	assert.EqualError(t, err, "namespaces \"default\" not found")
}

func TestFeatureContext_Patch(t *testing.T) {
	ctx := initFakeScenarioWithNamespaces(t)

	patch := []byte(`{"metadata":{"labels":{"new-label": ""}}}`)
	err := ctx.Patch(namespaceGVK, namespaceDefault, types.StrategicMergePatchType, patch)
	require.NoError(t, err)

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{"apiVersion": "v1", "kind": "Namespace"},
	}

	err = ctx.Client().Get(ctx.GoContext(), namespaceDefault, obj)
	require.NoError(t, err)
	assert.Contains(t, obj.GetLabels(), "new-label")
}

func TestFeatureContext_Patch_ResourceNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	err := ctx.Patch(namespaceGVK, namespaceDefault, types.JSONPatchType, nil)
	assert.EqualError(t, err, "namespaces \"default\" not found")
}

func TestFeatureContext_Delete(t *testing.T) {
	ctx := initFakeScenarioWithNamespaces(t)

	_, err := ctx.Delete(namespaceGVK, namespaceDefault)
	require.NoError(t, err)

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{"apiVersion": "v1", "kind": "Namespace"},
	}

	err = ctx.Client().Get(ctx.GoContext(), types.NamespacedName{Name: "default"}, obj)
	assert.True(t, errors.IsNotFound(err))
}

func TestFeatureContext_Delete_KindNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	_, err := ctx.Delete(notFoundGVK, namespaceDefault)
	assert.EqualError(t, err, "no kind \"NotFound\" is registered for version \"v1\" in scheme \"pkg/runtime/scheme.go:101\"")
}

func TestFeatureContext_Delete_ResourceNotFound(t *testing.T) {
	ctx := initFakeScenario(t)
	_, err := ctx.Delete(namespaceGVK, namespaceDefault)
	assert.EqualError(t, err, "namespaces \"default\" not found")
}
