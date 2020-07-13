Feature: Delete resources with errors
  In order to test deletion features
  As feature context
  I need to be able to manage deletion errors

  Scenario: should failed due to invalid GroupVersionKind on resource deletion
    When Kubernetes removes InvalidGVK 'default/default'

  Scenario: should failed due to unknown GroupVersionKind on resource deletion
    When Kubernetes removes v1/Unknown 'default/default'

  Scenario: should failed due to unknown GroupVersionKind on resource deletion
    Given Kubernetes doesn't have v1/Service 'default/svc'
    When Kubernetes removes v1/Service 'default/svc'

  Scenario: should failed due to invalid table on multi resource deletion
    When Kubernetes removes the following resources
      | InvalidTable |
      | LineA        |
      | LineB        |

  Scenario: should failed due to unknown GroupVersionKind on multi resource deletion
    When Kubernetes removes the following resources
      | ApiGroupVersion | Kind      | Namespace  | Name       |
      | v1              | Unknown   | kube-lease | svc        |
