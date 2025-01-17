# Model
model:
  rest_name: authz
  resource_name: authz
  entity_name: Authz
  package: cid
  group: core/authorization
  description: This is an internal API that is used to resolve to API authorization.
  private: true

# Attributes
attributes:
  v1:
  - name: claims
    description: The list of verified claims.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value:
    - '@auth:account=acme'
    - '@auth:realm=vince'

  - name: clientIP
    description: IP of the client.
    type: string
    exposed: true

  - name: error
    description: Return an eventual error.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: permissions
    description: |-
      If the parameter permissions=1 is set, targetIdentity and targetOperation are
      ignored and this attribute will contain all the permission for the given claims.
    type: external
    exposed: true
    subtype: map[string]map[string]bool
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: policyRulesList
    description: |-
      Contains the raw resolved and unfiltered PolicyRuleList. It will be empty unless
      the query parameter `forwardpolicyrules=true` is set.
      Note that this contains the straight rules that matched the user claims in the
      target namespace. It can contain rules that are filtered out in later stage
      because it applies to children namespace or because the rule only applies to a
      certain source network, or because the token contains restrictions. Unless you
      really know why you need it, you mostly don't.
    type: refList
    exposed: true
    subtype: policyrule
    omit_empty: true

  - name: restrictedNamespace
    description: Sets the namespace restrictions that should apply.
    type: string
    exposed: true
    example_value: /namespace

  - name: restrictedNetworks
    description: Sets the networks restrictions that should apply.
    type: list
    exposed: true
    subtype: string
    example_value:
    - 10.0.0.0/8

  - name: restrictedPermissions
    description: Sets the permissions restrictions that should apply.
    type: list
    exposed: true
    subtype: string
    example_value:
    - '@auth:role=enforcer'

  - name: targetID
    description: The ID of the object to check permission for.
    type: string
    exposed: true

  - name: targetNamespace
    description: The namespace where to check permission from.
    type: string
    exposed: true
    required: true
    example_value: /acme
