# Model
model:
  rest_name: cloudnetworkquery
  resource_name: cloudnetworkqueries
  entity_name: CloudNetworkQuery
  package: yeul
  group: pcn/infrastructure
  description: Provides the parameters for an effective network permissions query.
  get:
    description: Retrieves the cloud query with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the cloud query with the given ID.
  delete:
    description: Deletes the the cloud query with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  validations:
  - $cloudnetworkquery

# Attributes
attributes:
  v1:
  - name: destinationIP
    description: The destination IP of a trace route request. Might not always be
      an endpoint.
    type: string
    exposed: true
    stored: true
    validations:
    - $optionalcidrorip

  - name: destinationSelector
    description: A filter for selecting destinations for the query.
    type: ref
    exposed: true
    subtype: cloudnetworkqueryfilter
    stored: true
    extensions:
      refMode: pointer

  - name: effectiveAction
    description: |-
      Filters the results based on the effective action. 'ReachableAndAllowed' means
      that a destination is both reachable and allowed by security rules.
      'UnreachableOrRejected' means that the destination is either not reachable or
      rejected by security rules. 'ReachableOnly' means that all destinations that are
      reachable will be returned irrespective of their security verdict.
      'UnreachableOnly' means that only unreachable destinations will be returned.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ReachableAndAllowed
    - ReachableOnly
    - All
    default_value: ReachableOnly

  - name: excludeEnterpriseIPs
    description: |-
      If set, the evaluation will exclude enterprise IPs from the effective
      permissions.
    type: boolean
    exposed: true
    stored: true

  - name: excludedNetworks
    description: |-
      List of networks that should be a excluded from the calculation if the source or
      destination is a network.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidroriplist

  - name: protocolPorts
    description: |-
      Represents the ports and protocols this policy applies to. Protocol/ports are
      defined as tcp/80, udp/22. For protocols that do not have ports, the port
      designation
      is not allowed.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $serviceports

  - name: rawRQL
    description: The RQL string for this query as a reference.
    type: string
    stored: true

  - name: sourceIP
    description: The source IP of a trace route request. Might not be always and endpoint.
    type: string
    exposed: true
    stored: true
    validations:
    - $optionalcidrorip

  - name: sourceSelector
    description: A filter for selecting the sources of the request.
    type: ref
    exposed: true
    subtype: cloudnetworkqueryfilter
    stored: true
    extensions:
      refMode: pointer

  - name: type
    description: Indicates the type of results that should be provided by the query.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Summary
    - CompressedGraph
    - FullGraph
    default_value: Summary

# Relations
relations:
- rest_name: cloudgraph
  get:
    description: Initiates a calculation of the query and retrieves the results in
      CloudGraph.

- rest_name: cloudpolicy
  get:
    description: Retrieves the policies associated with this query.
  create:
    description: Creates a policy associated with this query.
