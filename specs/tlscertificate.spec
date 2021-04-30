# Model
model:
  rest_name: tlscertificate
  resource_name: tlscertificates
  entity_name: TLSCertificate
  package: service
  group: policy/services
  description: Represents a certificate public and private key.
  aliases: []
  get:
    description: Retrieves the certificate with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the service with the given ID.
  delete:
    description: Deletes the service with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@timeable'
  - '@propagated'

# Attributes
attributes:
  v1:
  - name: certificate
    description: PEM-encoded TLS certificate.
    type: string
    exposed: true
    stored: true
    validations:
    - $pem

  - name: key
    description: PEM-encoded TLS certificate key associated with `certificate`.
    type: string
    exposed: true
    stored: true
    encrypted: true
    validations:
    - $pem
