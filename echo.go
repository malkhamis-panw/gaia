package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EchoIdentity represents the Identity of the object.
var EchoIdentity = elemental.Identity{
	Name:     "echo",
	Category: "echo",
	Package:  "core",
	Private:  false,
}

// EchosList represents a list of Echos
type EchosList []*Echo

// Identity returns the identity of the objects in the list.
func (o EchosList) Identity() elemental.Identity {

	return EchoIdentity
}

// Copy returns a pointer to a copy the EchosList.
func (o EchosList) Copy() elemental.Identifiables {

	copy := append(EchosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EchosList.
func (o EchosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EchosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Echo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EchosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EchosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EchosList converted to SparseEchosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EchosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEchosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEcho)
	}

	return out
}

// Version returns the version of the content.
func (o EchosList) Version() int {

	return 1
}

// Echo represents the model of a echo
type Echo struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The description of the pizza.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEcho returns a new *Echo
func NewEcho() *Echo {

	return &Echo{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Echo) Identity() elemental.Identity {

	return EchoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Echo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Echo) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Echo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEcho{}

	s.Description = o.Description

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Echo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEcho{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Description = s.Description

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Echo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Echo) BleveType() string {

	return "echo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Echo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Echo) Doc() string {

	return `This is an echo.`
}

func (o *Echo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Echo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEcho{
			ID:          &o.ID,
			Description: &o.Description,
		}
	}

	sp := &SparseEcho{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "description":
			sp.Description = &(o.Description)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEcho to the object.
func (o *Echo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEcho)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
}

// DeepCopy returns a deep copy if the Echo.
func (o *Echo) DeepCopy() *Echo {

	if o == nil {
		return nil
	}

	out := &Echo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Echo.
func (o *Echo) DeepCopyInto(out *Echo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Echo: %s", err))
	}

	*out = *target.(*Echo)
}

// Validate valides the current information stored into the structure.
func (o *Echo) Validate() error {

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
func (*Echo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EchoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EchoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Echo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EchoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Echo) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "description":
		return o.Description
	}

	return nil
}

// EchoAttributesMap represents the map of attribute for Echo.
var EchoAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `The description of the pizza.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
}

// EchoLowerCaseAttributesMap represents the map of attribute for Echo.
var EchoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `The description of the pizza.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
}

// SparseEchosList represents a list of SparseEchos
type SparseEchosList []*SparseEcho

// Identity returns the identity of the objects in the list.
func (o SparseEchosList) Identity() elemental.Identity {

	return EchoIdentity
}

// Copy returns a pointer to a copy the SparseEchosList.
func (o SparseEchosList) Copy() elemental.Identifiables {

	copy := append(SparseEchosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEchosList.
func (o SparseEchosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEchosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEcho))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEchosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEchosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEchosList converted to EchosList.
func (o SparseEchosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEchosList) Version() int {

	return 1
}

// SparseEcho represents the sparse version of a echo.
type SparseEcho struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The description of the pizza.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEcho returns a new  SparseEcho.
func NewSparseEcho() *SparseEcho {
	return &SparseEcho{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEcho) Identity() elemental.Identity {

	return EchoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEcho) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEcho) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEcho) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEcho{}

	if o.Description != nil {
		s.Description = o.Description
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEcho) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEcho{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Description != nil {
		o.Description = s.Description
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseEcho) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEcho) ToPlain() elemental.PlainIdentifiable {

	out := NewEcho()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}

	return out
}

// DeepCopy returns a deep copy if the SparseEcho.
func (o *SparseEcho) DeepCopy() *SparseEcho {

	if o == nil {
		return nil
	}

	out := &SparseEcho{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEcho.
func (o *SparseEcho) DeepCopyInto(out *SparseEcho) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEcho: %s", err))
	}

	*out = *target.(*SparseEcho)
}

type mongoAttributesEcho struct {
	Description string `bson:"description"`
}
type mongoAttributesSparseEcho struct {
	Description *string `bson:"description,omitempty"`
}
