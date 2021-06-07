package kubernetes_ctx_helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

func TestGroupVersionKindFrom(t *testing.T) {
	tests := []struct {
		gvk    string
		expect schema.GroupVersionKind
		err    error
	}{
		{gvk: "v1/Namespace", expect: schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}},
		{gvk: "apps/v1/Deployment", expect: schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}},
		{gvk: "", err: fmt.Errorf("invalid GroupVersionKind ''")},
		{gvk: "a/b/c/d", err: fmt.Errorf("invalid GroupVersionKind 'a/b/c/d'")},
	}

	for _, tt := range tests {
		t.Run(tt.gvk, func(t *testing.T) {
			gvk, err := GroupVersionKindFrom(tt.gvk)

			switch {
			case tt.err != nil:
				assert.EqualError(t, err, tt.err.Error())
			default:
				assert.Equal(t, tt.expect, gvk)
			}
		})
	}
}

func TestNamespacedNameFrom(t *testing.T) {
	tests := []struct {
		name   string
		expect types.NamespacedName
		err    error
	}{
		{name: "default", expect: types.NamespacedName{Name: "default"}},
		{name: "default/app", expect: types.NamespacedName{Namespace: "default", Name: "app"}},
		{name: "", err: fmt.Errorf("invalid NamespacedName ''")},
		{name: "a/b/c", err: fmt.Errorf("invalid NamespacedName 'a/b/c'")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, err := NamespacedNameFrom(tt.name)

			switch {
			case tt.err != nil:
				assert.EqualError(t, err, tt.err.Error())
			default:
				assert.Equal(t, tt.expect, name)
			}
		})
	}
}
