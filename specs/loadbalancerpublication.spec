# Model
model:
  rest_name: loadbalancerpublication
  resource_name: loadbalancerpublications
  entity_name: LoadBalancerPublication
  package: squall
  group: policy/services
  description: |-
    Allows api users to publish a LoadBalancer in the namespace. It will create the
    Load Balancer as well as a API Authorization policy giving edit permission on
    the created object.
  aliases: []

# Attributes
attributes:
  v1:
  - name: loadBalancer
    description: LoadBalancer definition.
    type: ref
    exposed: true
    subtype: loadbalancer
    extensions:
      refMode: pointer

  - name: loadBalancerID
    description: Populated in response to give the newly created object ID.
    type: string
    exposed: true
