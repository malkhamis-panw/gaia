package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ServicePublicationIdentity represents the Identity of the object.
var ServicePublicationIdentity = elemental.Identity{
	Name:     "servicepublication",
	Category: "servicepublications",
	Package:  "squall",
	Private:  false,
}

// ServicePublicationsList represents a list of ServicePublications
type ServicePublicationsList []*ServicePublication

// Identity returns the identity of the objects in the list.
func (o ServicePublicationsList) Identity() elemental.Identity {

	return ServicePublicationIdentity
}

// Copy returns a pointer to a copy the ServicePublicationsList.
func (o ServicePublicationsList) Copy() elemental.Identifiables {

	copy := append(ServicePublicationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ServicePublicationsList.
func (o ServicePublicationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ServicePublicationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ServicePublication))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServicePublicationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServicePublicationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ServicePublicationsList converted to SparseServicePublicationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ServicePublicationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseServicePublicationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseServicePublication)
	}

	return out
}

// Version returns the version of the content.
func (o ServicePublicationsList) Version() int {

	return 1
}

// ServicePublication represents the model of a servicepublication
type ServicePublication struct {
	// The service object that will be published.
	Service *Service `json:"service" msgpack:"service" bson:"-" mapstructure:"service,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewServicePublication returns a new *ServicePublication
func NewServicePublication() *ServicePublication {

	return &ServicePublication{
		ModelVersion: 1,
		Service:      NewService(),
	}
}

// Identity returns the Identity of the object.
func (o *ServicePublication) Identity() elemental.Identity {

	return ServicePublicationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ServicePublication) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServicePublication) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ServicePublication) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesServicePublication{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ServicePublication) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesServicePublication{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ServicePublication) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ServicePublication) BleveType() string {

	return "servicepublication"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ServicePublication) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ServicePublication) Doc() string {

	return `Encapsulates a service object that is ought to be published so it can be used
in a sibling namespace.`
}

func (o *ServicePublication) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ServicePublication) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseServicePublication{
			Service: o.Service,
		}
	}

	sp := &SparseServicePublication{}
	for _, f := range fields {
		switch f {
		case "service":
			sp.Service = o.Service
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseServicePublication to the object.
func (o *ServicePublication) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseServicePublication)
	if so.Service != nil {
		o.Service = so.Service
	}
}

// DeepCopy returns a deep copy if the ServicePublication.
func (o *ServicePublication) DeepCopy() *ServicePublication {

	if o == nil {
		return nil
	}

	out := &ServicePublication{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ServicePublication.
func (o *ServicePublication) DeepCopyInto(out *ServicePublication) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ServicePublication: %s", err))
	}

	*out = *target.(*ServicePublication)
}

// Validate valides the current information stored into the structure.
func (o *ServicePublication) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Service != nil {
		elemental.ResetDefaultForZeroValues(o.Service)
		if err := o.Service.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*ServicePublication) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServicePublicationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServicePublicationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ServicePublication) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServicePublicationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ServicePublication) ValueForAttribute(name string) interface{} {

	switch name {
	case "service":
		return o.Service
	}

	return nil
}

// ServicePublicationAttributesMap represents the map of attribute for ServicePublication.
var ServicePublicationAttributesMap = map[string]elemental.AttributeSpecification{
	"Service": {
		AllowedChoices: []string{},
		ConvertedName:  "Service",
		Description:    `The service object that will be published.`,
		Exposed:        true,
		Name:           "service",
		Required:       true,
		SubType:        "service",
		Type:           "ref",
	},
}

// ServicePublicationLowerCaseAttributesMap represents the map of attribute for ServicePublication.
var ServicePublicationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"service": {
		AllowedChoices: []string{},
		ConvertedName:  "Service",
		Description:    `The service object that will be published.`,
		Exposed:        true,
		Name:           "service",
		Required:       true,
		SubType:        "service",
		Type:           "ref",
	},
}

// SparseServicePublicationsList represents a list of SparseServicePublications
type SparseServicePublicationsList []*SparseServicePublication

// Identity returns the identity of the objects in the list.
func (o SparseServicePublicationsList) Identity() elemental.Identity {

	return ServicePublicationIdentity
}

// Copy returns a pointer to a copy the SparseServicePublicationsList.
func (o SparseServicePublicationsList) Copy() elemental.Identifiables {

	copy := append(SparseServicePublicationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseServicePublicationsList.
func (o SparseServicePublicationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseServicePublicationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseServicePublication))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseServicePublicationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseServicePublicationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseServicePublicationsList converted to ServicePublicationsList.
func (o SparseServicePublicationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseServicePublicationsList) Version() int {

	return 1
}

// SparseServicePublication represents the sparse version of a servicepublication.
type SparseServicePublication struct {
	// The service object that will be published.
	Service *Service `json:"service,omitempty" msgpack:"service,omitempty" bson:"-" mapstructure:"service,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseServicePublication returns a new  SparseServicePublication.
func NewSparseServicePublication() *SparseServicePublication {
	return &SparseServicePublication{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseServicePublication) Identity() elemental.Identity {

	return ServicePublicationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseServicePublication) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseServicePublication) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseServicePublication) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseServicePublication{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseServicePublication) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseServicePublication{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseServicePublication) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseServicePublication) ToPlain() elemental.PlainIdentifiable {

	out := NewServicePublication()
	if o.Service != nil {
		out.Service = o.Service
	}

	return out
}

// DeepCopy returns a deep copy if the SparseServicePublication.
func (o *SparseServicePublication) DeepCopy() *SparseServicePublication {

	if o == nil {
		return nil
	}

	out := &SparseServicePublication{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseServicePublication.
func (o *SparseServicePublication) DeepCopyInto(out *SparseServicePublication) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseServicePublication: %s", err))
	}

	*out = *target.(*SparseServicePublication)
}

type mongoAttributesServicePublication struct {
}
type mongoAttributesSparseServicePublication struct {
}
