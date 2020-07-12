Feature: Get resources
  In order to test gathering features
  As feature context
  I need to be able to gather existing resources

  Scenario: should find existing resources
    Given Kubernetes has v1/Namespace 'default'
    And Kubernetes has v1/Namespace 'kube-public'
    And Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes has v1/Service 'default/default'
    And Kubernetes has v1/Service 'default/kubernetes'

  Scenario: should not find non-existing resources
    Given Kubernetes doesn't have v1/Namespace 'kube-lease'

  Scenario: should find similarity between two resources
    Given Kubernetes has v1/Namespace 'kube-public'
    And Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-public' is similar to 'kube-system'

  Scenario: should not find similarity between two resources
    Given Kubernetes has v1/Service 'default/default'
    And Kubernetes has v1/Service 'default/kubernetes'
    And Kubernetes resource v1/Service 'default/default' is not similar to 'default/kubernetes'

  Scenario: should find equality between two resources
    Given Kubernetes has v1/Namespace 'default'
    And Kubernetes has v1/Namespace 'kube-public'
    And Kubernetes resource v1/Namespace 'default' is equal to 'kube-public'

  Scenario: should not find equality between two resources
    Given Kubernetes has v1/Namespace 'kube-public'
    And Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-public' is not equal to 'kube-system'
