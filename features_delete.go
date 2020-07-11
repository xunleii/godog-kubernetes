package kubernetes_ctx

import "kubernetes_ctx/helpers"

// RemoveResource implements the GoDoc step
// - `Kubernetes removes <ApiGroupVersionKind> '<NamespacedName>'`
// It removes the specified resource.
func RemoveResource(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes removes (`+RxGroupVersionKind+`) '(`+RxNamespacedName+`)'$`,
		func(groupVersionKindStr, name string) error {
			groupVersionKind, err := helpers.GroupVersionKindFrom(groupVersionKindStr)
			if err != nil {
				return err
			}
			namespacedName, err := helpers.NamespacedNameFrom(name)
			if err != nil {
				return err
			}

			_, err = ctx.Delete(groupVersionKind, namespacedName)
			return err
		},
	)
}

// RemoveMultiResource implements the GoDoc step
// - `Kubernetes removes the following resources <RESOURCES_TABLE>`
// It creates several resources in a row.
func RemoveMultiResource(ctx *FeatureContext, s ScenarioContext) {
	s.Step(
		`^Kubernetes removes the following resources$`,
		func(table helpers.ResourceTable) error {
			resources, err := helpers.UnmarshalResourceTable(table)
			if err != nil {
				return err
			}

			for _, resource := range resources {
				_, err := ctx.Delete(resource.GroupVersionKind(), resource.NamespacedName())
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}
