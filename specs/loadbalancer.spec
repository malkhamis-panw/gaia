# Model
model:
  rest_name: loadbalancer
  resource_name: loadbalancers
  entity_name: LoadBalancer
  package: squall
  group: policy/services
  description: Defines a generic external LoadBalancer that sits between 2 enforcers.
  aliases:
  - lb
  - lbs
  get:
    description: Retrieves the Load Balancer with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the service with the given ID.
  delete:
    description: Deletes the service with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@archivable'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@disabled'
  - '@timeable'
  - '@propagated'
  - '@servicebase'

# Indexes
indexes:
- - allTLSCertificateTags
- - namespace
  - allTLSCertificateTags

# Attributes
attributes:
  v1:
  - name: TLSCertificatesSelector
    description: |-
      A tag or tag expression that identifies the TLS certificates to be used by
      enforcers when exposing the service ran by the processing units.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - certificate=service-a
    validations:
    - $tagsExpression

  - name: allTLSCertificateTags
    description: This is a set of all selector tags for matching in the database.
    type: list
    subtype: string
    stored: true

  - name: proxyProtocolEnabled
    description: Enable trust in proxy protocols header.
    type: boolean
    exposed: true
    stored: true

  - name: proxyProtocolSubnets
    description: Only allow proxy protocols from the given subnets .
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidrs

  - name: trustedCertificateAuthorities
    description: |-
      PEM-encoded certificate authorities to trust when additional hops are needed. It
      must be set if the service must reach a service marked as `external` or must go
      through an additional TLS termination point like a layer 7 load balancer.
    type: string
    exposed: true
    stored: true
    validations:
    - $pem

  - name: type
    description: Type of the load balancer.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - HTTP
    - TCP

# Relations
relations:
- rest_name: processingunit
  get:
    description: Retrieves the processing units that implement this service.
