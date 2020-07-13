Feature: Patch resources with errors
  In order to test patching features
  As feature context
  I need to be able to manage patching errors

  Scenario: should failed due to invalid GroupVersionKind on resource patching
    When Kubernetes patches InvalidGVK 'default/kubernetes' with
      """
      spec:
        type: LoadBalancer
      """

  Scenario: should failed due to unknown GroupVersionKind on resource patching
    When Kubernetes patches v1/Unknown 'default/kubernetes' with
      """
      spec:
        type: LoadBalancer
      """

  Scenario: should failed due to non-existent resource on resource patching
    When Kubernetes patches v1/Service 'default/unknown' with
      """
      spec:
        type: LoadBalancer
      """

  Scenario: should failed due to invalid YAML on resource patching
    When Kubernetes patches v1/Service 'default/unknown' with
      """
      invalidYAML
      """
