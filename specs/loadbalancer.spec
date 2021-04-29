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

# Indexes
indexes:
- - allProcessingUnitTags
- - namespace
  - allProcessingUnitTags

# Attributes
attributes:
  v1:
  - name: IPs
    description: |-
      The list of IP addresses where the service can be accessed. This is an optional
      attribute and is only required if no host names are provided. The system will
      automatically resolve IP addresses from host names otherwise.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalipaddresslist

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

  - name: allProcessingUnitTags
    description: This is a set of all selector tags for matching in the database.
    type: list
    subtype: string
    stored: true
    read_only: true

  - name: allTLSCertificateTags
    description: This is a set of all selector tags for matching in the database.
    type: list
    subtype: string
    stored: true
    read_only: true

  - name: exposedPort
    description: |-
      The port that the service can be accessed on. Note that this is different from
      the `port` attribute that describes the port that the service is actually
      listening on. For example if a load balancer is used, the `exposedPort` is
      the port that the load balancer is listening for the service, whereas the
      port that the implementation is listening on can be different.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 443
    max_value: 65535

  - name: exposedServiceIsTLS
    description: |-
      Indicates that the exposed service is TLS. This means that the enforcer has to
      initiate a TLS session in order to forward traffic to the service.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    filterable: true
    orderable: true

  - name: hosts
    description: The host names that the service can be accessed on.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: port
    description: |-
      The port that the implementation of the service is listening to. It can be
      different than `exposedPort`. This is needed for port mapping use cases
      where there are private and public ports.
    type: integer
    exposed: true
    stored: true
    example_value: 443
    max_value: 65535

  - name: processingUnitSelector
    description: |-
      Tag expression that identifies the processing unit that implements this
      particular service.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - $identity=processingunit
    validations:
    - $tagsExpression

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

  - name: publicApplicationPort
    description: |-
      A new virtual port that the service can be accessed on using HTTPS. Since the
      enforcer transparently inserts TLS in the application path, you might want
      to declare a new port where the enforcer listens for TLS. However, the
      application does not need to be modified and the enforcer will map the
      traffic to the correct application port. This is useful when
      an application is being accessed from a public network.
    type: integer
    exposed: true
    stored: true
    example_value: 443
    max_value: 65535

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
