# Model
model:
  rest_name: cloudnetworkqueryreply
  resource_name: cloudnetworkqueryreplies
  entity_name: CloudNetworkQueryReply
  package: vargid
  group: pcn/infrastructure
  description: The reply object for a Cloud Network Query from the Graph Calculator.
  extends:
  - '@base'
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: alertRuleID
    description: Prisma Cloud Alert Rule id.
    type: string
    exposed: true
    subtype: string

  - name: cloudGraphResult
    description: The result of the cloud network query.
    type: ref
    exposed: true
    subtype: cloudgraph
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
