Feature: Patch resources
  In order to test patching features
  As feature context
  I need to be able to list patching resources

  Scenario: should patch resource
    Given Kubernetes resource v1/Service 'default/kubernetes' doesn't have 'spec.type=LoadBalancer'
    When Kubernetes patches v1/Service 'default/kubernetes' with
      """
      spec:
        type: LoadBalancer
      """
    Then Kubernetes resource v1/Service 'default/kubernetes' has 'spec.type=LoadBalancer'
