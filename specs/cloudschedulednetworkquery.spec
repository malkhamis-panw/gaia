# Model
model:
  rest_name: cloudschedulednetworkquery
  resource_name: cloudschedulednetworkqueries
  entity_name: CloudScheduledNetworkQuery
  package: vargid
  group: pcn/infrastructure
  description: |-
    CloudSchedulednNetworkQuery represents a CloudNetworkQuery that will be
    scheduled periodically.
  private: true
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'

# Indexes
indexes:
- - alertruleid
- - lastexecutiontimestamp
- - alertruleid
  - policyid

# Attributes
attributes:
  v1:
  - name: alertRuleID
    description: Prisma Cloud Alert Rule ID.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: cloudGraphResult
    description: The result of the cloud network query.
    type: ref
    exposed: true
    subtype: cloudgraph
    extensions:
      refMode: pointer

  - name: cloudNetworkQuery
    description: The cloud network query that should be used.
    type: ref
    exposed: true
    subtype: cloudnetworkquery
    stored: true
    extensions:
      refMode: pointer

  - name: disabled
    description: Represents whether the associated policy was disabled.
    type: boolean
    stored: true
    getter: true
    setter: true

  - name: lastExecutionTimestamp
    description: Result of the last successfully run query.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: policyID
    description: Prisma Cloud Policy ID.
    type: string
    exposed: true
    subtype: string
    stored: true
