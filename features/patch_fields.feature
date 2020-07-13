Feature: Patch resource fields
  In order to test patching features
  As feature context
  I need to be able to list patching resource fields

  Scenario: should add label
    Given Kubernetes resource v1/Service 'default/default' doesn't have label 'key'
    When Kubernetes labelizes v1/Service 'default/default' with 'key=value'
    Then Kubernetes resource v1/Service 'default/default' has label 'key=value'

  Scenario: should update label
    Given Kubernetes resource v1/Service 'default/kubernetes' has label 'key=value'
    When Kubernetes updates label 'key' on v1/Service 'default/kubernetes' with 'error'
    Then Kubernetes resource v1/Service 'default/kubernetes' has label 'key=error'

  Scenario: should remove label
    Given Kubernetes resource v1/Service 'default/kubernetes' has label 'key=value'
    When Kubernetes removes label 'key' on v1/Service 'default/kubernetes'
    Then Kubernetes resource v1/Service 'default/kubernetes' doesn't have label 'key'

  Scenario: should add annotation
    Given Kubernetes resource v1/Service 'default/default' doesn't have annotation 'key'
    When Kubernetes annotates v1/Service 'default/default' with 'key=value'
    Then Kubernetes resource v1/Service 'default/default' has annotation 'key=value'

  Scenario: should update annotation
    Given Kubernetes resource v1/Service 'default/kubernetes' has annotation 'key=value'
    When Kubernetes updates annotation 'key' on v1/Service 'default/kubernetes' with 'error'
    Then Kubernetes resource v1/Service 'default/kubernetes' has annotation 'key=error'

  Scenario: should remove annotation
    Given Kubernetes resource v1/Service 'default/kubernetes' has annotation 'key=value'
    When Kubernetes removes annotation 'key' on v1/Service 'default/kubernetes'
    Then Kubernetes resource v1/Service 'default/kubernetes' doesn't have annotation 'key'
