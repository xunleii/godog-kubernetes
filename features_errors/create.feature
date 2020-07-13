Feature: Create resources with errors
  In order to test creation features
  As feature context
  I need to be able to manage creation errors

  Scenario: should failed due to invalid GroupVersionKind on resource creation
    When Kubernetes creates a new InvalidGVK 'kube-lease/svc'

  Scenario: should failed due to unknown GroupVersionKind on resource creation
    When Kubernetes creates a new v1/Unknown 'kube-lease/svc'

  Scenario: should failed due to invalid GroupVersionKind on resource creation from YAML
    When Kubernetes creates a new InvalidGVK 'kube-lease/svc-lb' with
      """
      spec:
        type: LoadBalancer
      """

  Scenario: should failed due to unknown GroupVersionKind on resource creation from YAML
    When Kubernetes creates a new v1/Unknown 'kube-lease/svc-lb' with
      """
      spec:
        type: LoadBalancer
      """

  Scenario: should failed due to invalid YAML on resource creation from YAML
    When Kubernetes creates a new v1/Service 'kube-lease/svc-lb' with
      """
      invalid yaml
      """

  Scenario: should failed due to invalid GroupVersionKind on resource creation from file
    When Kubernetes creates a new InvalidGVK 'kube-lease/svc-lb' from features/resources/kube-lease/svc-lb.yaml

  Scenario: should failed due to unknown GroupVersionKind on resource creation from file
    When Kubernetes creates a new v1/Unknown 'kube-lease/svc-lb' from features/resources/kube-lease/svc-lb.yaml

  Scenario: should failed due to nonexistent file on resource creation from file
    When Kubernetes creates a new v1/Service 'kube-lease/svc-lb' from features_errors/resources/kube-lease/svc-lb.yaml

  Scenario: should failed due to invalid file on resource creation from file
    When Kubernetes creates a new v1/Service 'kube-lease/svc-lb' from features_errors/resources/kube-lease/svc.yaml

  Scenario: should failed due to invalid table on multi resource creation
    When Kubernetes creates the following resources
      | InvalidTable |
      | LineA        |
      | LineB        |

  Scenario: should failed due to unknown GroupVersionKind on multi resource creation
    When Kubernetes creates the following resources
      | ApiGroupVersion | Kind      | Namespace  | Name       |
      | v1              | Unknown   | kube-lease | svc        |
