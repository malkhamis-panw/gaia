package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// RoleIdentity represents the Identity of the object
var RoleIdentity = elemental.Identity{
	Name:     "role",
	Category: "roles",
}

// RolesList represents a list of Roles
type RolesList []*Role

// ContentIdentity returns the identity of the objects in the list.
func (o RolesList) ContentIdentity() elemental.Identity {

	return RoleIdentity
}

// Copy returns a pointer to a copy the RolesList.
func (o RolesList) Copy() elemental.ContentIdentifiable {

	copy := append(RolesList{}, o...)
	return &copy
}

// List converts the object to an elemental.IdentifiablesList.
func (o RolesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RolesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o RolesList) Version() int {

	return 1
}

// Role represents the model of a role
type Role struct {
	// Authorizations of the role.
	Authorizations map[string][]string `json:"authorizations" bson:"-"`

	// Description is the description of the role.
	Description string `json:"description" bson:"-"`

	// Key is the of the role.
	Key string `json:"key" bson:"-"`

	// Name of the role.
	Name string `json:"name" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRole returns a new *Role
func NewRole() *Role {

	return &Role{
		ModelVersion:   1,
		Authorizations: map[string][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Role) Identity() elemental.Identity {

	return RoleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Role) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Role) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Role) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Role) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Role) Doc() string {
	return `Roles returns the available roles that can be used with API Authorization Policies.`
}

func (o *Role) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Role) Validate() error {

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
func (*Role) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RoleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RoleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Role) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RoleAttributesMap
}

// RoleAttributesMap represents the map of attribute for Role.
var RoleAttributesMap = map[string]elemental.AttributeSpecification{
	"Authorizations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Authorizations of the role.`,
		Exposed:        true,
		Name:           "authorizations",
		ReadOnly:       true,
		SubType:        "authorization_map",
		Type:           "external",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Description is the description of the role.`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		ReadOnly:       true,
		Type:           "string",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key is the of the role.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Name of the role.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
}

// RoleLowerCaseAttributesMap represents the map of attribute for Role.
var RoleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"authorizations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Authorizations of the role.`,
		Exposed:        true,
		Name:           "authorizations",
		ReadOnly:       true,
		SubType:        "authorization_map",
		Type:           "external",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Description is the description of the role.`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		ReadOnly:       true,
		Type:           "string",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key is the of the role.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Name of the role.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
}
