Feature: Get resource fields
  In order to test gathering features
  As feature context
  I need to be able to gather existing resource fields

  Scenario: should find existing or valid resource field
    Given Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-system' has 'metadata.name'
    And Kubernetes resource v1/Namespace 'kube-system' has 'metadata.annotations.key=value'

  Scenario: should not find non-existing resource field or detect invalid resource field
    Given Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-system' doesn't have 'metadata.annotations.oops'
    And Kubernetes resource v1/Namespace 'kube-system' doesn't have 'metadata.annotations.key=error'

  Scenario: should find existing or valid label
    Given Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-system' has label 'key'
    And Kubernetes resource v1/Namespace 'kube-system' has label 'key=value'

  Scenario: should not find non-existing label or detect invalid label
    Given Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-system' doesn't have label 'oops'
    And Kubernetes resource v1/Namespace 'kube-system' doesn't have label 'key=error'

  Scenario: should find existing or valid annotation
    Given Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-system' has annotation 'key'
    And Kubernetes resource v1/Namespace 'kube-system' has annotation 'key=value'

  Scenario: should not find non-existing label or detect invalid label
    Given Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes resource v1/Namespace 'kube-system' doesn't have annotation 'oops'
    And Kubernetes resource v1/Namespace 'kube-system' doesn't have annotation 'key=error'
