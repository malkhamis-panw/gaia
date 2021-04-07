# Model
model:
  rest_name: cloudaddress
  resource_name: cloudaddresses
  entity_name: CloudAddress
  package: yeul
  group: pcn/infrastructure
  description: Managed the list of IP addresses associated with an interface.
  detached: true

# Attributes
attributes:
  v1:
  - name: IPVersion
    description: Designates IPv4 or IPv6.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - IPv4
    - IPv6
    example_value: IPv4

  - name: primary
    description: Designates the IP address as the primary IP address.
    type: boolean
    exposed: true
    stored: true
    example_value: true

  - name: privateDNSName
    description: The private DNS name associated with the address.
    type: string
    exposed: true
    stored: true
    example_value: ip-172-20-53-29.us-west-2.compute.internal

  - name: privateIP
    description: The private IP address value.
    type: string
    exposed: true
    stored: true
    example_value: 10.1.1.2
    validations:
    - $optionalcidrorip

  - name: publicDNSName
    description: The private DNS name associated with the address.
    type: string
    exposed: true
    stored: true
    example_value: ip-172-20-53-29.us-west-2.compute.internal

  - name: publicIP
    description: The private IP address value.
    type: string
    exposed: true
    stored: true
    example_value: 10.1.1.2
    validations:
    - $optionalcidrorip