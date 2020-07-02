# Model
model:
  rest_name: pingprobe
  resource_name: pingprobes
  entity_name: PingProbe
  package: guy
  group: core/enforcer
  description: |-
    Represents the result of a unique ping probe. They are aggregated into a
    PingResult.
  get:
    description: Retrieves a ping result.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@timeable'
  - '@identifiable-stored'

# Indexes
indexes:
- - pingID
- - namespace
  - pingID
- - namespace
  - pingID
  - iterationIndex

# Attributes
attributes:
  v1:
  - name: ACLPolicyAction
    description: Action of the ACL policy.
    type: string
    exposed: true
    stored: true

  - name: ACLPolicyID
    description: ID of the ACL policy.
    type: string
    exposed: true
    stored: true

  - name: RTT
    description: Time taken for a single request-response to complete.
    type: string
    exposed: true
    stored: true

  - name: applicationListening
    description: If true, application responded to the request.
    type: boolean
    exposed: true
    stored: true

  - name: claims
    description: Claims of the processing unit.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: claimsType
    description: Type of claims reported.
    type: enum
    exposed: true
    subtype: string
    stored: true
    required: true
    allowed_choices:
    - Transmitted
    - Received
    example_value:
    - Transmitted

  - name: enforcerID
    description: ID of the defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns

  - name: enforcerVersion
    description: Semantic version of the defender.
    type: string
    exposed: true
    stored: true

  - name: error
    description: A non-empty error indicates a failure.
    type: string
    exposed: true
    stored: true

  - name: excludedNetworks
    description: If true, destination IP is in `excludedNetworks`.
    type: boolean
    exposed: true
    stored: true

  - name: fourTuple
    description: Four tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true
    stored: true

  - name: isServer
    description: If true, the report was generated by the server.
    type: boolean
    exposed: true
    stored: true

  - name: iterationIndex
    description: Holds the iteration number this probe is attached to.
    type: integer
    exposed: true
    stored: true

  - name: payloadSize
    description: Size of the payload attached to the packet.
    type: integer
    exposed: true
    stored: true

  - name: payloadSizeType
    description: Type of the payload size.
    type: enum
    exposed: true
    subtype: string
    stored: true
    required: true
    allowed_choices:
    - Transmitted
    - Received
    example_value:
    - Transmitted

  - name: peerCertExpiry
    description: Represents the expiry of the peer certificate.
    type: string
    exposed: true
    stored: true

  - name: peerCertIssuer
    description: Represents the issuer of the peer certificate.
    type: string
    exposed: true
    stored: true

  - name: peerCertSubject
    description: Represents the subject of the peer certificate.
    type: string
    exposed: true
    stored: true

  - name: pingID
    description: PingID unique to a single ping control.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: policyAction
    description: Action of the policy.
    type: string
    exposed: true
    stored: true

  - name: policyID
    description: ID of the policy.
    type: string
    exposed: true
    stored: true

  - name: policyNamespace
    description: ID of the policy.
    type: string
    exposed: true
    stored: true

  - name: processingUnitID
    description: ID of the reporting processing unit.
    type: string
    exposed: true
    stored: true

  - name: protocol
    description: Protocol used for the communication.
    type: integer
    exposed: true
    stored: true

  - name: remoteController
    description: Controller of the remote endpoint.
    type: string
    exposed: true
    stored: true

  - name: remoteEndpointType
    description: Represents the remote endpoint type.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - External
    example_value:
    - External

  - name: remoteNamespace
    description: Namespace of the remote processing unit.
    type: string
    exposed: true
    stored: true

  - name: remoteNamespaceType
    description: Type of the namespace reported.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Plain
    - Hash
    example_value:
    - Plain

  - name: remoteProcessingUnitID
    description: ID of the remote processing unit.
    type: string
    exposed: true
    stored: true

  - name: seqNum
    description: Sequence number of the TCP packet. number.
    type: integer
    exposed: true
    stored: true

  - name: serviceID
    description: ID of the service If the service type is a proxy.
    type: string
    exposed: true
    stored: true

  - name: serviceType
    description: Type of the service.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: targetTCPNetworks
    description: If true, destination IP is in `targetTCPNetworks`.
    type: boolean
    exposed: true
    stored: true

  - name: type
    description: Type of the report.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Request
    - Response
    example_value:
    - Request
