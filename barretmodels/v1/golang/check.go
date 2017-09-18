package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// CheckIdentity represents the Identity of the object
var CheckIdentity = elemental.Identity{
	Name:     "check",
	Category: "checks",
}

// ChecksList represents a list of Checks
type ChecksList []*Check

// ContentIdentity returns the identity of the objects in the list.
func (o ChecksList) ContentIdentity() elemental.Identity {

	return CheckIdentity
}

// Copy returns a pointer to a copy the ChecksList.
func (o ChecksList) Copy() elemental.ContentIdentifiable {

	copy := append(ChecksList{}, o...)
	return &copy
}

// List converts the object to an elemental.IdentifiablesList.
func (o ChecksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ChecksList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ChecksList) Version() int {

	return 1
}

// Check represents the model of a check
type Check struct {
	// ID contains the certificate serialNumber
	ID string `json:"ID" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCheck returns a new *Check
func NewCheck() *Check {

	return &Check{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Check) Identity() elemental.Identity {

	return CheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Check) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Check) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Check) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Check) DefaultOrder() []string {

	return []string{}
}

func (o *Check) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Check) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Check) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Check) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CheckAttributesMap
}

// CheckAttributesMap represents the map of attribute for Check.
var CheckAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ID contains the certificate serialNumber`,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
}

// CheckLowerCaseAttributesMap represents the map of attribute for Check.
var CheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ID contains the certificate serialNumber`,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
}
