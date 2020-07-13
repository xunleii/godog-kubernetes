Feature: List resources with errors
  In order to test listing features
  As feature context
  I need to be able to manage listing errors

  Scenario: should failed due to invalid GroupVersionKind on resource listing
    When Kubernetes has 2 InvalidGVK

  Scenario: should failed due to unknown GroupVersionKind on resource listing
    When Kubernetes has 2 v1/Unknown

  Scenario: should failed due to the wrong count on resource listing (namespaced)
    Given Kubernetes has 2 v1/Service
    When Kubernetes has 1 v1/Service

  Scenario: should failed due to the wrong count on resource listing
    Given Kubernetes has 3 v1/Namespace
    When Kubernetes has 1 v1/Namespace

  Scenario: should failed due to the non-existent on resource listing
    When Kubernetes has 2 v1/Pod

  Scenario: should failed due to invalid GroupVersionKind on namespaced resource listing
    When Kubernetes has 2 InvalidGVK in namespace 'default'

  Scenario: should failed due to unknown GroupVersionKind on namespaced resource listing
    When Kubernetes has 2 v1/Unknown in namespace 'default'

  Scenario: should failed due to the wrong count on namespaced resource listing
    Given Kubernetes has 2 v1/Service in namespace 'default'
    When Kubernetes has 1 v1/Service in namespace 'default'

  Scenario: should failed due to the non-existent on namespaced resource listing
    When Kubernetes has 2 v1/Pod in namespace 'default'

