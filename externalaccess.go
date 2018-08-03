package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// ExternalAccessIndexes lists the attribute compound indexes.
var ExternalAccessIndexes = [][]string{}

// ExternalAccessIdentity represents the Identity of the object.
var ExternalAccessIdentity = elemental.Identity{
	Name:     "externalaccess",
	Category: "externalaccesses",
	Private:  false,
}

// ExternalAccessList represents a list of ExternalAccess
type ExternalAccessList []*ExternalAccess

// Identity returns the identity of the objects in the list.
func (o ExternalAccessList) Identity() elemental.Identity {

	return ExternalAccessIdentity
}

// Copy returns a pointer to a copy the ExternalAccessList.
func (o ExternalAccessList) Copy() elemental.Identifiables {

	copy := append(ExternalAccessList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ExternalAccessList.
func (o ExternalAccessList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ExternalAccessList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ExternalAccess))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ExternalAccessList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ExternalAccessList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ExternalAccessList) Version() int {

	return 1
}

// ExternalAccess represents the model of a externalaccess
type ExternalAccess struct {
	// IPRecords refers to a list of IPRecord that contains the IP information.
	IPRecords []*types.IPRecord `json:"IPRecords" bson:"-" mapstructure:"IPRecords,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewExternalAccess returns a new *ExternalAccess
func NewExternalAccess() *ExternalAccess {

	return &ExternalAccess{
		ModelVersion: 1,
		IPRecords:    []*types.IPRecord{},
	}
}

// Identity returns the Identity of the object.
func (o *ExternalAccess) Identity() elemental.Identity {

	return ExternalAccessIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ExternalAccess) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ExternalAccess) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *ExternalAccess) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ExternalAccess) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ExternalAccess) Doc() string {
	return `ExternalAccess allows to retrieve connection from or to an external service.`
}

func (o *ExternalAccess) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *ExternalAccess) Validate() error {

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
func (*ExternalAccess) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ExternalAccessAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ExternalAccessLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ExternalAccess) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ExternalAccessAttributesMap
}

// ExternalAccessAttributesMap represents the map of attribute for ExternalAccess.
var ExternalAccessAttributesMap = map[string]elemental.AttributeSpecification{
	"IPRecords": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IPRecords",
		Description:    `IPRecords refers to a list of IPRecord that contains the IP information.`,
		Exposed:        true,
		Name:           "IPRecords",
		ReadOnly:       true,
		SubType:        "ip_records",
		Type:           "external",
	},
}

// ExternalAccessLowerCaseAttributesMap represents the map of attribute for ExternalAccess.
var ExternalAccessLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"iprecords": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IPRecords",
		Description:    `IPRecords refers to a list of IPRecord that contains the IP information.`,
		Exposed:        true,
		Name:           "IPRecords",
		ReadOnly:       true,
		SubType:        "ip_records",
		Type:           "external",
	},
}
