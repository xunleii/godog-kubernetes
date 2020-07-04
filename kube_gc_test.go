package kubernetes_ctx_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"kubernetes_ctx"
)

func TestNaiveGC(t *testing.T) {
	const (
		rawService = `apiVersion: v1
kind: Service
metadata:
  name: ownerService
  namespace: default
spec:
  clusterIP: None`
		rawEndpoints = `apiVersion: v1
kind: Endpoints
metadata:
  name: ownedEndpoints
  namespace: default
  ownerReferences:
    - apiVersion: v1
      kind: Service
      name: ownerService
      uid: %s`
	)

	ctx := initFakeScenarioWithNamespaces(t)

	svc := yamlToUnstructured(t, rawService)
	err := ctx.Create(svc.GroupVersionKind(), types.NamespacedName{Namespace: svc.GetNamespace(), Name: svc.GetName()}, svc)
	require.NoError(t, err)
	svc, err = ctx.Get(svc.GroupVersionKind(), types.NamespacedName{Namespace: svc.GetNamespace(), Name: svc.GetName()})
	require.NoError(t, err)

	endpoints := yamlToUnstructured(t, fmt.Sprintf(rawEndpoints, svc.GetUID()))
	err = ctx.Create(endpoints.GroupVersionKind(), types.NamespacedName{Namespace: endpoints.GetNamespace(), Name: endpoints.GetName()}, endpoints)
	require.NoError(t, err)
	endpoints, err = ctx.Get(endpoints.GroupVersionKind(), types.NamespacedName{Namespace: endpoints.GetNamespace(), Name: endpoints.GetName()})
	assert.False(t, errors.IsNotFound(err))

	err = kubernetes_ctx.NaiveGC(ctx, svc)
	require.NoError(t, err)

	_, err = ctx.Get(endpoints.GroupVersionKind(), types.NamespacedName{Namespace: endpoints.GetNamespace(), Name: endpoints.GetName()})
	assert.True(t, errors.IsNotFound(err))
}
