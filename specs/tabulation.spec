# Model
model:
  rest_name: tabulation
  resource_name: tabulations
  entity_name: Tabulation
  package: elena
  description: |-
    Tabulate API allows you to retrieve a custom table view for any identity using
    any tags you like as columns.
  aliases:
  - table
  - tables
  - tabs
  - tab

# Attributes
attributes:
  v1:
  - name: headers
    description: Headers contains the requests headers that matched.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true

  - name: rows
    description: Rows contains the tabulated data.
    type: external
    exposed: true
    subtype: tabulated_data
    read_only: true
    autogenerated: true

  - name: targetIdentity
    description: TargetIdentity contains the requested target identity.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    format: free