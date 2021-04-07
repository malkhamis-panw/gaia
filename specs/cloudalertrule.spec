# Model
model:
  rest_name: cloudalertrule
  resource_name: cloudalertsrule
  entity_name: CloudAlertRule
  package: vargid
  group: pcn/infrastructure
  description: |-
    Creates a Prisma Cloud alert rule along with policies and accounts associated
    with the alert rule.
  get:
    description: Retrieves the Prisma Cloud Alert Rule with a given id.
    global_parameters:
    - $filtering
  update:
    description: Updates the Prisma Cloud Alert Rule with a given id.
  delete:
    description: Deletes the Prisma Cloud Alert Rule with a given id.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@prismabase'
  - '@named'

# Indexes
indexes:
- - namespace
  - alertruleid
  - enabled

# Attributes
attributes:
  v1:
  - name: accountids
    description: List of accounts associated to alert rule.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: alertruleid
    description: Prisma Cloud Alert Rule id.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: description
    description: Alert rule description.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: enabled
    description: Is Alert Rule enabled.
    type: boolean
    exposed: true
    subtype: boolean
    stored: true

  - name: policyids
    description: List of policy IDs associated to alert rule.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: regions
    description: List of regions where the alert rule is enforced.
    type: list
    exposed: true
    subtype: string
    stored: true
