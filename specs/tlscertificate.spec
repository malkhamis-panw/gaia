# Model
model:
  rest_name: tlscertificate
  resource_name: tlscertificates
  entity_name: TLSCertificate
  package: service
  group: policy/services
  description: Represents a certificate public and private key.
  aliases: []

# Attributes
attributes:
  v1:
  - name: certificate
    description: |-
      PEM-encoded certificate to expose to the clients for TLS. Only has effect and
      required if `TLSType` is set to `External`.
    type: string
    exposed: true
    stored: true

  - name: key
    description: |-
      PEM-encoded certificate key associated with `TLSCertificate`. Only has effect
      and required if `TLSType` is set to `External`.
    type: string
    exposed: true
    stored: true
    encrypted: true
