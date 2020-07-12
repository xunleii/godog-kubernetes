Feature: Delete resources
  In order to test deletion features
  As feature context
  I need to be able to delete new resources

  Scenario: should delete resource
    Given Kubernetes has v1/Service 'default/default'
    When Kubernetes removes v1/Service 'default/default'
    Then Kubernetes doesn't have v1/Service 'default/default'

  Scenario: should delete several resources
    Given Kubernetes has v1/Namespace 'default'
    And Kubernetes has v1/Namespace 'kube-public'
    And Kubernetes has v1/Namespace 'kube-system'
    And Kubernetes has v1/Service 'default/default'
    And Kubernetes has v1/Service 'default/kubernetes'
    When Kubernetes removes the following resources
      | ApiGroupVersion | Kind      | Namespace  | Name         |
      | v1              | Namespace |            | default      |
      | v1              | Namespace |            | kube-public  |
      | v1              | Namespace |            | kube-system  |
      | v1              | Service   | default    | default      |
      | v1              | Service   | default    | kubernetes   |
    Then Kubernetes doesn't have v1/Namespace 'default'
    And Kubernetes doesn't have v1/Namespace 'kube-public'
    And Kubernetes doesn't have v1/Namespace 'kube-system'
    And Kubernetes doesn't have v1/Service 'default/default'
    And Kubernetes doesn't have v1/Service 'default/kubernetes'
