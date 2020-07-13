Feature: List resources
  In order to test listing features
  As feature context
  I need to be able to list existing resources

  Scenario: should list the number of namespace
    Given Kubernetes has v1/Namespace 'default'
    And Kubernetes has v1/Namespace 'kube-public'
    And Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes has 3 v1/Namespace

  Scenario: should list the number of services
    Given Kubernetes has v1/Service 'default/default'
    And Kubernetes has v1/Service 'default/kubernetes'
    And Kubernetes has 2 v1/Service in namespace 'default'
