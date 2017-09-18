package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// RevocationIdentity represents the Identity of the object
var RevocationIdentity = elemental.Identity{
	Name:     "revocation",
	Category: "revocations",
}

// RevocationsList represents a list of Revocations
type RevocationsList []*Revocation

// ContentIdentity returns the identity of the objects in the list.
func (o RevocationsList) ContentIdentity() elemental.Identity {

	return RevocationIdentity
}

// Copy returns a pointer to a copy the RevocationsList.
func (o RevocationsList) Copy() elemental.ContentIdentifiable {

	copy := append(RevocationsList{}, o...)
	return &copy
}

// List converts the object to an elemental.IdentifiablesList.
func (o RevocationsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RevocationsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o RevocationsList) Version() int {

	return 1
}

// Revocation represents the model of a revocation
type Revocation struct {
	// ID contains the ID of the revocation.
	ID string `json:"-" bson:"_id"`

	// Status of the revocation.
	Revoked bool `json:"revoked" bson:"revoked"`

	// SerialNumber of the revoked certificate.
	SerialNumber string `json:"serialNumber" bson:"serialnumber"`

	// Subject of the certificate related to the revocation.
	Subject string `json:"subject" bson:"subject"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRevocation returns a new *Revocation
func NewRevocation() *Revocation {

	return &Revocation{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Revocation) Identity() elemental.Identity {

	return RevocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Revocation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Revocation) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Revocation) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Revocation) DefaultOrder() []string {

	return []string{}
}

func (o *Revocation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Revocation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Revocation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RevocationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RevocationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Revocation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RevocationAttributesMap
}

// RevocationAttributesMap represents the map of attribute for Revocation.
var RevocationAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID contains the ID of the revocation.`,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Revoked": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Status of the revocation.`,
		Exposed:        true,
		Name:           "revoked",
		Stored:         true,
		Type:           "boolean",
	},
	"SerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `SerialNumber of the revoked certificate.`,
		Exposed:        true,
		Format:         "free",
		Name:           "serialNumber",
		Stored:         true,
		Type:           "string",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Subject of the certificate related to the revocation.`,
		Exposed:        true,
		Format:         "free",
		Name:           "subject",
		Stored:         true,
		Type:           "string",
	},
}

// RevocationLowerCaseAttributesMap represents the map of attribute for Revocation.
var RevocationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID contains the ID of the revocation.`,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"revoked": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Status of the revocation.`,
		Exposed:        true,
		Name:           "revoked",
		Stored:         true,
		Type:           "boolean",
	},
	"serialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `SerialNumber of the revoked certificate.`,
		Exposed:        true,
		Format:         "free",
		Name:           "serialNumber",
		Stored:         true,
		Type:           "string",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Subject of the certificate related to the revocation.`,
		Exposed:        true,
		Format:         "free",
		Name:           "subject",
		Stored:         true,
		Type:           "string",
	},
}
