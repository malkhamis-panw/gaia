# Model
model:
  rest_name: networkrulesetpolicy
  resource_name: networkrulesetpolicies
  entity_name: NetworkRuleSetPolicy
  package: squall
  group: policy/networking
  description: |-
    Allows you to define network rule sets to allow or prevent processing units
    identified by their tags to talk to other processing units or external networks
    (also identified by their tags).
  aliases:
  - netruleset
  - netrulesets
  - netset
  - netsets
  - networkruleset
  - networkrulesets
  get:
    description: Retrieves the policy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the policy with the given ID.
  delete:
    description: Deletes the policy with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: incomingRules
    description: |-
      The set of rules to apply to incoming traffic (traffic coming to the Processing
      Unit matching the subject).
    type: refList
    exposed: true
    subtype: networkrule
    extensions:
      refMode: pointer

  - name: outgoingRules
    description: |-
      The set of rules to apply to outgoing traffic (traffic coming from the
      Processing Unit matching the subject).
    type: refList
    exposed: true
    subtype: networkrule
    extensions:
      refMode: pointer

  - name: subject
    description: |-
      A tag expression identifying used to match processing units to which this policy
      applies to.
    type: external
    exposed: true
    subtype: '[][]string'
    validations:
    - $atLeastOneSubExpression
    - $subExpressionsNotEmpty
    - $noDuplicateTagsInEachSubExpression
    - $tagsExpression

# Relations
relations:
- rest_name: externalnetwork
  get:
    description: Returns the list of external networks affected by a network rule
      set policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object

- rest_name: processingunit
  get:
    description: Returns the list of processing units affected by a network rule set
      policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object

- rest_name: service
  get:
    description: Returns the list of services affected by a network rule set policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object
