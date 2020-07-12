Feature: Create resources
  In order to test creation features
  As feature context
  I need to be able to create new resources

  Scenario: should create resource
    Given Kubernetes must have v1/Namespace 'kube-lease'
    When Kubernetes creates a new v1/Service 'kube-lease/svc'
    Then Kubernetes has v1/Namespace 'kube-lease'
    And Kubernetes has v1/Service 'kube-lease/svc'

  Scenario: should create resource from YAML
    Given Kubernetes must have v1/Namespace 'kube-lease'
    And Kubernetes must have v1/Service 'kube-lease/svc' with
      """
      metadata:
        annotations:
          key: value
      """
    When Kubernetes creates a new v1/Service 'kube-lease/svc-lb' with
      """
      spec:
        type: LoadBalancer
      """
    Then Kubernetes has v1/Namespace 'kube-lease'
    And Kubernetes has v1/Service 'kube-lease/svc'
    And Kubernetes resource v1/Service 'kube-lease/svc' has annotation 'key=value'
    And Kubernetes has v1/Service 'kube-lease/svc-lb'
    And Kubernetes resource v1/Service 'kube-lease/svc-lb' has 'spec.type=LoadBalancer'

  Scenario: should create resource from file
    Given Kubernetes must have v1/Namespace 'kube-lease'
    And Kubernetes must have v1/Service 'kube-lease/svc' from features/resources/kube-lease/svc.yaml
    When Kubernetes creates a new v1/Service 'kube-lease/svc-lb' from features/resources/kube-lease/svc-lb.yaml
    Then Kubernetes has v1/Namespace 'kube-lease'
    And Kubernetes has v1/Service 'kube-lease/svc'
    And Kubernetes resource v1/Service 'kube-lease/svc' has annotation 'key=value'
    And Kubernetes has v1/Service 'kube-lease/svc-lb'
    And Kubernetes resource v1/Service 'kube-lease/svc-lb' has 'spec.type=LoadBalancer'

  Scenario: should create several resources at a time
    Given Kubernetes must have the following resources
      | ApiGroupVersion | Kind      | Namespace | Name        |
      | v1              | Namespace |           | kube-lease  |
      | v1              | Namespace |           | not-default |
      | v1              | Namespace |           | no-idea     |
    When Kubernetes creates the following resources
      | ApiGroupVersion | Kind      | Namespace  | Name       |
      | v1              | Service   | kube-lease | svc        |
      | v1              | Service   | kube-lease | svc-lb     |
    Then Kubernetes has v1/Namespace 'kube-lease'
    And Kubernetes has v1/Namespace 'not-default'
    And Kubernetes has v1/Namespace 'no-idea'
    And Kubernetes has v1/Service 'kube-lease/svc'
    And Kubernetes has v1/Service 'kube-lease/svc-lb'
