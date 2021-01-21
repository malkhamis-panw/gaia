# Model
model:
  rest_name: echo
  resource_name: echo
  entity_name: Echo
  package: core
  group: core
  description: This is an echo.
  extends:
  - '@identifiable-not-stored'

# Attributes
attributes:
  v1:
  - name: description
    description: The description of the pizza.
    type: string
    exposed: true
    stored: true
