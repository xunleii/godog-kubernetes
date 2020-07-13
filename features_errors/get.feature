Feature: Get resource fields with errors
  In order to test gathering features
  As feature context
  I need to be able to manage gathering errors

  Scenario: should failed due to invalid GroupVersionKind on resource gathering
    When Kubernetes has InvalidGVK 'kube-lease/svc'

  Scenario: should failed due to unknown GroupVersionKind on resource gathering
    When Kubernetes has v1/Unknown 'kube-lease/svc'

  Scenario: should failed due to non-existent resource on resource gathering
    When Kubernetes has v1/Service 'default/svc'

  Scenario: should failed due to invalid GroupVersionKind on non-existent resource gathering
    When Kubernetes doesn't have InvalidGVK 'kube-lease/svc'

  Scenario: should failed due to unknown GroupVersionKind on non-existent resource gathering
    When Kubernetes doesn't have v1/Unknown 'kube-lease/svc'

  Scenario: should failed due to existing resource on non-existent resource gathering
    When Kubernetes doesn't have v1/Service 'default/kubernetes'

  Scenario: should failed due to invalid GroupVersionKind on resource comparision (similarity)
    When Kubernetes resource InvalidGVK 'default/default' is similar to 'default/kubernetes'

  Scenario: should failed due to unknown GroupVersionKind on resource comparision (similarity)
    When Kubernetes resource v1/Unknown 'default/default' is similar to 'default/kubernetes'

  Scenario: should failed due to non-existing resource on resource comparision (similarity)
    When Kubernetes resource v1/Service 'default/svc' is similar to 'default/kubernetes'

  Scenario: should failed due to non-existing resource on resource comparision (similarity)
    When Kubernetes resource v1/Service 'default/kubernetes' is similar to 'default/svc'

  Scenario: should failed due to non-existing resource on resource comparision (similarity)
    When Kubernetes resource v1/Service 'default/default' is similar to 'default/kubernetes'

  Scenario: should failed due to invalid GroupVersionKind on resource diff (similarity)
    When Kubernetes resource InvalidGVK 'default/default' is not similar to 'default/default'

  Scenario: should failed due to unknown GroupVersionKind on resource diff (similarity)
    When Kubernetes resource v1/Unknown 'default/default' is not similar to 'default/default'

  Scenario: should failed due to non-existing resource on resource diff (similarity)
    When Kubernetes resource v1/Service 'default/svc' is not similar to 'default/default'

  Scenario: should failed due to non-existing resource on resource diff (similarity)
    When Kubernetes resource v1/Service 'default/kubernetes' is not similar to 'default/svc'

  Scenario: should failed due to non-existing resource on resource diff (similarity)
    When Kubernetes resource v1/Service 'default/default' is not similar to 'default/default'

  Scenario: should failed due to invalid GroupVersionKind on resource comparision (equality)
    When Kubernetes resource InvalidGVK 'default/default' is equal to 'default/kubernetes'

  Scenario: should failed due to unknown GroupVersionKind on resource comparision (equality)
    When Kubernetes resource v1/Unknown 'default/default' is equal to 'default/kubernetes'

  Scenario: should failed due to non-existing resource on resource comparision (equality)
    When Kubernetes resource v1/Service 'default/svc' is equal to 'default/kubernetes'

  Scenario: should failed due to non-existing resource on resource comparision (equality)
    When Kubernetes resource v1/Service 'default/kubernetes' is equal to 'default/svc'

  Scenario: should failed due to non-existing resource on resource comparision (equality)
    When Kubernetes resource v1/Service 'default/default' is equal to 'default/kubernetes'

  Scenario: should failed due to invalid GroupVersionKind on resource diff (equality)
    When Kubernetes resource InvalidGVK 'default/default' is not equal to 'default/default'

  Scenario: should failed due to unknown GroupVersionKind on resource diff (equality)
    When Kubernetes resource v1/Unknown 'default/default' is not equal to 'default/default'

  Scenario: should failed due to non-existing resource on resource diff (equality)
    When Kubernetes resource v1/Service 'default/svc' is not equal to 'default/default'

  Scenario: should failed due to non-existing resource on resource diff (equality)
    When Kubernetes resource v1/Service 'default/kubernetes' is not equal to 'default/svc'

  Scenario: should failed due to non-existing resource on resource diff (equality)
    When Kubernetes resource v1/Service 'default/default' is not equal to 'default/default'
