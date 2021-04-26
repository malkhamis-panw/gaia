# Model
model:
  rest_name: cloudalertrulecnq
  resource_name: cloudalertrulecnqs
  entity_name: CloudAlertRuleCNQ
  package: vargid
  group: pcn/infrastructure
  description: The result object of executing policies in an alert rule.
  get:
    description: Retrieves the last run result of a given policy in an alert rule.
    global_parameters:
    - $filtering
  update:
    description: Updates the CNQ Result of a given policy in an alert rule.
  delete:
    description: Deletes the CNQ Result of a given policy in an alert rule.
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
- - namespace
  - alertruleid
  - policyid
- - namespace
  - resulttimestamp

# Attributes
attributes:
  v1:
  - name: alertRuleID
    description: Prisma Cloud Alert Rule id.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: cloudNetworkQuery
    description: The cloud network query that should be used.
    type: ref
    exposed: true
    subtype: cloudnetworkquery
    stored: true
    extensions:
      refMode: pointer

  - name: policyID
    description: Prisma Cloud Policy id.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: resultTimestamp
    description: Result of the last successfully run query.
    type: time
    exposed: true
    stored: true
    orderable: true
    extensions:
      bson_name: ag
