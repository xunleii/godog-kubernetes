Feature: Patch resource fields with errors
  In order to test patching features
  As feature context
  I need to be able to manage patching errors

  # Label patching
  Scenario: should failed due to invalid GroupVersionKind on label creation
    When Kubernetes labelizes InvalidGVK 'default/svc' with 'key=value'

  Scenario: should failed due to unknown GroupVersionKind on label creation
    When Kubernetes labelizes v1/Unknown 'default/svc' with 'key=value'

  Scenario: should failed due to non-existent resource on label creation
    When Kubernetes labelizes v1/Service 'default/svc' with 'key=value'

  Scenario: should failed due to invalid GroupVersionKind on label update
    When Kubernetes updates label 'key' on InvalidGVK 'default/svc' with 'value'

  Scenario: should failed due to unknown GroupVersionKind on label update
    When Kubernetes updates label 'key' on v1/Unknown 'default/svc' with 'value'

  Scenario: should failed due to non-existent resource on label update
    When Kubernetes updates label 'key' on v1/Service 'default/svc' with 'value'

  Scenario: should failed due to non-existent labels on label update
    When Kubernetes updates label 'key' on v1/Service 'default/default' with 'value'

  Scenario: should failed due to non-existent label on label update
    When Kubernetes updates label 'unknown' on v1/Service 'default/kubernetes' with 'value'

  Scenario: should failed due to invalid GroupVersionKind on label deletion
    When Kubernetes removes label 'key' on InvalidGVK 'default/svc'

  Scenario: should failed due to unknown GroupVersionKind on label deletion
    When Kubernetes removes label 'key' on v1/Unknown 'default/svc'

  Scenario: should failed due to non-existent resource on label deletion
    When Kubernetes removes label 'key' on v1/Service 'default/svc'

  Scenario: should failed due to non-existent labels on label deletion
    When Kubernetes removes label 'key' on v1/Service 'default/default'

  Scenario: should failed due to non-existent label on label deletion
    When Kubernetes removes label 'unknown' on v1/Service 'default/kubernetes'

  # Annotation patching
  Scenario: should failed due to invalid GroupVersionKind on annotation creation
    When Kubernetes annotates InvalidGVK 'default/svc' with 'key=value'

  Scenario: should failed due to unknown GroupVersionKind on annotation creation
    When Kubernetes annotates v1/Unknown 'default/svc' with 'key=value'

  Scenario: should failed due to non-existent resource on annotation creation
    When Kubernetes annotates v1/Service 'default/svc' with 'key=value'

  Scenario: should failed due to invalid GroupVersionKind on annotation update
    When Kubernetes updates annotation 'key' on InvalidGVK 'default/svc' with 'value'

  Scenario: should failed due to unknown GroupVersionKind on annotation update
    When Kubernetes updates annotation 'key' on v1/Unknown 'default/svc' with 'value'

  Scenario: should failed due to non-existent resource on annotation update
    When Kubernetes updates annotation 'key' on v1/Service 'default/svc' with 'value'

  Scenario: should failed due to non-existent annotations on annotation update
    When Kubernetes updates annotation 'key' on v1/Service 'default/default' with 'value'

  Scenario: should failed due to non-existent annotation on annotation update
    When Kubernetes updates annotation 'unknown' on v1/Service 'default/kubernetes' with 'value'

  Scenario: should failed due to invalid GroupVersionKind on annotation deletion
    When Kubernetes removes annotation 'key' on InvalidGVK 'default/svc'

  Scenario: should failed due to unknown GroupVersionKind on annotation deletion
    When Kubernetes removes annotation 'key' on v1/Unknown 'default/svc'

  Scenario: should failed due to non-existent resource on annotation deletion
    When Kubernetes removes annotation 'key' on v1/Service 'default/svc'

  Scenario: should failed due to non-existent annotations on annotation deletion
    When Kubernetes removes annotation 'key' on v1/Service 'default/default'

  Scenario: should failed due to non-existent annotation on annotation deletion
    When Kubernetes removes annotation 'unknown' on v1/Service 'default/kubernetes'
