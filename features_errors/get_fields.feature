Feature: Get resource fields with errors
  In order to test gathering features
  As feature context
  I need to be able to manage gathering errors

  # Resource fields gathering
  Scenario: should failed due to invalid GroupVersionKind on resource field gathering
    When Kubernetes resource InvalidGVK 'default/svc' has 'spec'

  Scenario: should failed due to unknown GroupVersionKind on resource field gathering
    When Kubernetes resource v1/Unknown 'default/svc' has 'spec'

  Scenario: should failed due to non-existent resource on resource field gathering
    When Kubernetes resource v1/Service 'default/svc' has 'spec'

  Scenario: should failed due to non-existent resource field on resource field gathering
    When Kubernetes resource v1/Service 'default/default' has 'unknown'

  Scenario: should failed due to invalid GroupVersionKind on non-existent resource field gathering
    When Kubernetes resource InvalidGVK 'default/svc' doesn't have 'spec'

  Scenario: should failed due to unknown GroupVersionKind on non-existent resource field gathering
    When Kubernetes resource v1/Unknown 'default/svc' doesn't have 'spec'

  Scenario: should failed due to non-existent resource on non-existent resource field gathering
    When Kubernetes resource v1/Service 'default/svc' doesn't have 'spec'

  Scenario: should failed due to existing resource field on non-existent resource field gathering
    When Kubernetes resource v1/Service 'default/default' doesn't have 'kind'

  Scenario: should failed due to invalid GroupVersionKind on resource field comparison
    When Kubernetes resource InvalidGVK 'default/svc' has 'metadata.name=svc'

  Scenario: should failed due to unknown GroupVersionKind on resource field comparison
    When Kubernetes resource v1/Unknown 'default/svc' has 'metadata.name=svc'

  Scenario: should failed due to non-existent resource on resource field comparison
    When Kubernetes resource v1/Service 'default/svc' has 'metadata.name=svc'

  Scenario: should failed due to non-existent resource field on resource field comparison
    When Kubernetes resource v1/Service 'default/default' has 'spec.unknown=svc'

  Scenario: should failed due to non-equal resource field value on resource field comparison
    When Kubernetes resource v1/Service 'default/default' has 'metadata.name=svc'

  Scenario: should failed due to invalid GroupVersionKind on resource field differentiation
    When Kubernetes resource InvalidGVK 'default/svc' doesn't have 'metadata.name=svc'

  Scenario: should failed due to unknown GroupVersionKind on resource field differentiation
    When Kubernetes resource v1/Unknown 'default/svc' doesn't have 'metadata.name=svc'

  Scenario: should failed due to non-existent resource on resource field differentiation
    When Kubernetes resource v1/Service 'default/svc' doesn't have 'metadata.name=svc'

  Scenario: should failed due to non-existent resource field on resource field differentiation
    When Kubernetes resource v1/Service 'default/default' doesn't have 'spec.unknown=svc'

  Scenario: should failed due to equal resource field value on resource field differentiation
    When Kubernetes resource v1/Service 'default/default' doesn't have 'metadata.name=default'

  # Resource labels gathering
  Scenario: should failed due to invalid GroupVersionKind on label gathering
    When Kubernetes resource InvalidGVK 'default/svc' has label 'key'

  Scenario: should failed due to unknown GroupVersionKind on label gathering
    When Kubernetes resource v1/Unknown 'default/svc' has label 'key'

  Scenario: should failed due to non-existent resource on label gathering
    When Kubernetes resource v1/Service 'default/svc' has label 'key'

  Scenario: should failed due to non-existent labels on label gathering
    When Kubernetes resource v1/Service 'default/default' has label 'key'

  Scenario: should failed due to non-existent label on label gathering
    When Kubernetes resource v1/Service 'default/kubernetes' has label 'unknown'

  Scenario: should failed due to invalid GroupVersionKind on non-existent label gathering
    When Kubernetes resource InvalidGVK 'default/svc' doesn't have label 'key'

  Scenario: should failed due to unknown GroupVersionKind on non-existent label gathering
    When Kubernetes resource v1/Unknown 'default/svc' doesn't have label 'key'

  Scenario: should failed due to non-existent resource on non-existent label gathering
    When Kubernetes resource v1/Service 'default/svc' doesn't have label 'key'

  Scenario: should failed due to existing resource label on non-existent label gathering
    When Kubernetes resource v1/Service 'default/kubernetes' doesn't have label 'key'

  Scenario: should failed due to invalid GroupVersionKind on label comparison
    When Kubernetes resource InvalidGVK 'default/svc' has label 'key=value'

  Scenario: should failed due to unknown GroupVersionKind on label comparison
    When Kubernetes resource v1/Unknown 'default/svc' has label 'key=value'

  Scenario: should failed due to non-existent resource on label comparison
    When Kubernetes resource v1/Service 'default/svc' has label 'key=value'

  Scenario: should failed due to non-existent labels on label comparison
    When Kubernetes resource v1/Service 'default/default' has label 'key=value'

  Scenario: should failed due to non-existent label on label comparison
    When Kubernetes resource v1/Service 'default/kubernetes' has label 'unknown=value'

  Scenario: should failed due to non-equal label value on label comparison
    When Kubernetes resource v1/Service 'default/kubernetes' has label 'key=notvalue'

  Scenario: should failed due to invalid GroupVersionKind on label differentiation
    When Kubernetes resource InvalidGVK 'default/svc' doesn't have label 'key=value'

  Scenario: should failed due to unknown GroupVersionKind on label differentiation
    When Kubernetes resource v1/Unknown 'default/svc' doesn't have label 'key=value'

  Scenario: should failed due to non-existent resource on label differentiation
    When Kubernetes resource v1/Service 'default/svc' doesn't have label 'key=value'

  Scenario: should failed due to equal label value on label differentiation
    When Kubernetes resource v1/Service 'default/kubernetes' doesn't have label 'key=value'

  # Resource annotations gathering
  Scenario: should failed due to invalid GroupVersionKind on annotation gathering
    When Kubernetes resource InvalidGVK 'default/svc' has annotation 'key'

  Scenario: should failed due to unknown GroupVersionKind on annotation gathering
    When Kubernetes resource v1/Unknown 'default/svc' has annotation 'key'

  Scenario: should failed due to non-existent resource on annotation gathering
    When Kubernetes resource v1/Service 'default/svc' has annotation 'key'

  Scenario: should failed due to non-existent annotations on annotation gathering
    When Kubernetes resource v1/Service 'default/default' has annotation 'key'

  Scenario: should failed due to non-existent annotation on annotation gathering
    When Kubernetes resource v1/Service 'default/kubernetes' has annotation 'unknown'

  Scenario: should failed due to invalid GroupVersionKind on non-existent annotation gathering
    When Kubernetes resource InvalidGVK 'default/svc' doesn't have annotation 'key'

  Scenario: should failed due to unknown GroupVersionKind on non-existent annotation gathering
    When Kubernetes resource v1/Unknown 'default/svc' doesn't have annotation 'key'

  Scenario: should failed due to non-existent resource on non-existent annotation gathering
    When Kubernetes resource v1/Service 'default/svc' doesn't have annotation 'key'

  Scenario: should failed due to existing resource annotation on non-existent annotation gathering
    When Kubernetes resource v1/Service 'default/kubernetes' doesn't have annotation 'key'

  Scenario: should failed due to invalid GroupVersionKind on annotation comparison
    When Kubernetes resource InvalidGVK 'default/svc' has annotation 'key=value'

  Scenario: should failed due to unknown GroupVersionKind on annotation comparison
    When Kubernetes resource v1/Unknown 'default/svc' has annotation 'key=value'

  Scenario: should failed due to non-existent resource on annotation comparison
    When Kubernetes resource v1/Service 'default/svc' has annotation 'key=value'

  Scenario: should failed due to non-existent annotations on annotation comparison
    When Kubernetes resource v1/Service 'default/default' has annotation 'key=value'

  Scenario: should failed due to non-existent annotation on annotation comparison
    When Kubernetes resource v1/Service 'default/kubernetes' has annotation 'unknown=value'

  Scenario: should failed due to non-equal annotation value on annotation comparison
    When Kubernetes resource v1/Service 'default/kubernetes' has annotation 'key=notvalue'

  Scenario: should failed due to invalid GroupVersionKind on annotation differentiation
    When Kubernetes resource InvalidGVK 'default/svc' doesn't have annotation 'key=value'

  Scenario: should failed due to unknown GroupVersionKind on annotation differentiation
    When Kubernetes resource v1/Unknown 'default/svc' doesn't have annotation 'key=value'

  Scenario: should failed due to non-existent resource on annotation differentiation
    When Kubernetes resource v1/Service 'default/svc' doesn't have annotation 'key=value'

  Scenario: should failed due to equal annotation value on annotation differentiation
    When Kubernetes resource v1/Service 'default/kubernetes' doesn't have annotation 'key=value'
