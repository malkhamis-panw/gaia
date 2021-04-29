package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LoadBalancerPublicationIdentity represents the Identity of the object.
var LoadBalancerPublicationIdentity = elemental.Identity{
	Name:     "loadbalancerpublication",
	Category: "loadbalancerpublications",
	Package:  "squall",
	Private:  false,
}

// LoadBalancerPublicationsList represents a list of LoadBalancerPublications
type LoadBalancerPublicationsList []*LoadBalancerPublication

// Identity returns the identity of the objects in the list.
func (o LoadBalancerPublicationsList) Identity() elemental.Identity {

	return LoadBalancerPublicationIdentity
}

// Copy returns a pointer to a copy the LoadBalancerPublicationsList.
func (o LoadBalancerPublicationsList) Copy() elemental.Identifiables {

	copy := append(LoadBalancerPublicationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LoadBalancerPublicationsList.
func (o LoadBalancerPublicationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LoadBalancerPublicationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LoadBalancerPublication))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LoadBalancerPublicationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LoadBalancerPublicationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the LoadBalancerPublicationsList converted to SparseLoadBalancerPublicationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LoadBalancerPublicationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLoadBalancerPublicationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLoadBalancerPublication)
	}

	return out
}

// Version returns the version of the content.
func (o LoadBalancerPublicationsList) Version() int {

	return 1
}

// LoadBalancerPublication represents the model of a loadbalancerpublication
type LoadBalancerPublication struct {
	// LoadBalancer definition.
	LoadBalancer *LoadBalancer `json:"loadBalancer" msgpack:"loadBalancer" bson:"-" mapstructure:"loadBalancer,omitempty"`

	// Populated in response to give the newly created object ID.
	LoadBalancerID string `json:"loadBalancerID" msgpack:"loadBalancerID" bson:"-" mapstructure:"loadBalancerID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLoadBalancerPublication returns a new *LoadBalancerPublication
func NewLoadBalancerPublication() *LoadBalancerPublication {

	return &LoadBalancerPublication{
		ModelVersion: 1,
		LoadBalancer: NewLoadBalancer(),
	}
}

// Identity returns the Identity of the object.
func (o *LoadBalancerPublication) Identity() elemental.Identity {

	return LoadBalancerPublicationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LoadBalancerPublication) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LoadBalancerPublication) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LoadBalancerPublication) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLoadBalancerPublication{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LoadBalancerPublication) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLoadBalancerPublication{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *LoadBalancerPublication) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *LoadBalancerPublication) BleveType() string {

	return "loadbalancerpublication"
}

// DefaultOrder returns the list of default ordering fields.
func (o *LoadBalancerPublication) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *LoadBalancerPublication) Doc() string {

	return `Allows api users to publish a LoadBalancer in the namespace. It will create the
Load Balancer as well as a API Authorization policy giving edit permission on
the created object.`
}

func (o *LoadBalancerPublication) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LoadBalancerPublication) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLoadBalancerPublication{
			LoadBalancer:   o.LoadBalancer,
			LoadBalancerID: &o.LoadBalancerID,
		}
	}

	sp := &SparseLoadBalancerPublication{}
	for _, f := range fields {
		switch f {
		case "loadBalancer":
			sp.LoadBalancer = o.LoadBalancer
		case "loadBalancerID":
			sp.LoadBalancerID = &(o.LoadBalancerID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLoadBalancerPublication to the object.
func (o *LoadBalancerPublication) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLoadBalancerPublication)
	if so.LoadBalancer != nil {
		o.LoadBalancer = so.LoadBalancer
	}
	if so.LoadBalancerID != nil {
		o.LoadBalancerID = *so.LoadBalancerID
	}
}

// DeepCopy returns a deep copy if the LoadBalancerPublication.
func (o *LoadBalancerPublication) DeepCopy() *LoadBalancerPublication {

	if o == nil {
		return nil
	}

	out := &LoadBalancerPublication{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LoadBalancerPublication.
func (o *LoadBalancerPublication) DeepCopyInto(out *LoadBalancerPublication) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LoadBalancerPublication: %s", err))
	}

	*out = *target.(*LoadBalancerPublication)
}

// Validate valides the current information stored into the structure.
func (o *LoadBalancerPublication) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.LoadBalancer != nil {
		elemental.ResetDefaultForZeroValues(o.LoadBalancer)
		if err := o.LoadBalancer.Validate(); err != nil {
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
func (*LoadBalancerPublication) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LoadBalancerPublicationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LoadBalancerPublicationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LoadBalancerPublication) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LoadBalancerPublicationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LoadBalancerPublication) ValueForAttribute(name string) interface{} {

	switch name {
	case "loadBalancer":
		return o.LoadBalancer
	case "loadBalancerID":
		return o.LoadBalancerID
	}

	return nil
}

// LoadBalancerPublicationAttributesMap represents the map of attribute for LoadBalancerPublication.
var LoadBalancerPublicationAttributesMap = map[string]elemental.AttributeSpecification{
	"LoadBalancer": {
		AllowedChoices: []string{},
		ConvertedName:  "LoadBalancer",
		Description:    `LoadBalancer definition.`,
		Exposed:        true,
		Name:           "loadBalancer",
		SubType:        "loadbalancer",
		Type:           "ref",
	},
	"LoadBalancerID": {
		AllowedChoices: []string{},
		ConvertedName:  "LoadBalancerID",
		Description:    `Populated in response to give the newly created object ID.`,
		Exposed:        true,
		Name:           "loadBalancerID",
		Type:           "string",
	},
}

// LoadBalancerPublicationLowerCaseAttributesMap represents the map of attribute for LoadBalancerPublication.
var LoadBalancerPublicationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"loadbalancer": {
		AllowedChoices: []string{},
		ConvertedName:  "LoadBalancer",
		Description:    `LoadBalancer definition.`,
		Exposed:        true,
		Name:           "loadBalancer",
		SubType:        "loadbalancer",
		Type:           "ref",
	},
	"loadbalancerid": {
		AllowedChoices: []string{},
		ConvertedName:  "LoadBalancerID",
		Description:    `Populated in response to give the newly created object ID.`,
		Exposed:        true,
		Name:           "loadBalancerID",
		Type:           "string",
	},
}

// SparseLoadBalancerPublicationsList represents a list of SparseLoadBalancerPublications
type SparseLoadBalancerPublicationsList []*SparseLoadBalancerPublication

// Identity returns the identity of the objects in the list.
func (o SparseLoadBalancerPublicationsList) Identity() elemental.Identity {

	return LoadBalancerPublicationIdentity
}

// Copy returns a pointer to a copy the SparseLoadBalancerPublicationsList.
func (o SparseLoadBalancerPublicationsList) Copy() elemental.Identifiables {

	copy := append(SparseLoadBalancerPublicationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLoadBalancerPublicationsList.
func (o SparseLoadBalancerPublicationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLoadBalancerPublicationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLoadBalancerPublication))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLoadBalancerPublicationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLoadBalancerPublicationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseLoadBalancerPublicationsList converted to LoadBalancerPublicationsList.
func (o SparseLoadBalancerPublicationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLoadBalancerPublicationsList) Version() int {

	return 1
}

// SparseLoadBalancerPublication represents the sparse version of a loadbalancerpublication.
type SparseLoadBalancerPublication struct {
	// LoadBalancer definition.
	LoadBalancer *LoadBalancer `json:"loadBalancer,omitempty" msgpack:"loadBalancer,omitempty" bson:"-" mapstructure:"loadBalancer,omitempty"`

	// Populated in response to give the newly created object ID.
	LoadBalancerID *string `json:"loadBalancerID,omitempty" msgpack:"loadBalancerID,omitempty" bson:"-" mapstructure:"loadBalancerID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLoadBalancerPublication returns a new  SparseLoadBalancerPublication.
func NewSparseLoadBalancerPublication() *SparseLoadBalancerPublication {
	return &SparseLoadBalancerPublication{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLoadBalancerPublication) Identity() elemental.Identity {

	return LoadBalancerPublicationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLoadBalancerPublication) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLoadBalancerPublication) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLoadBalancerPublication) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLoadBalancerPublication{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLoadBalancerPublication) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLoadBalancerPublication{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseLoadBalancerPublication) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLoadBalancerPublication) ToPlain() elemental.PlainIdentifiable {

	out := NewLoadBalancerPublication()
	if o.LoadBalancer != nil {
		out.LoadBalancer = o.LoadBalancer
	}
	if o.LoadBalancerID != nil {
		out.LoadBalancerID = *o.LoadBalancerID
	}

	return out
}

// DeepCopy returns a deep copy if the SparseLoadBalancerPublication.
func (o *SparseLoadBalancerPublication) DeepCopy() *SparseLoadBalancerPublication {

	if o == nil {
		return nil
	}

	out := &SparseLoadBalancerPublication{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLoadBalancerPublication.
func (o *SparseLoadBalancerPublication) DeepCopyInto(out *SparseLoadBalancerPublication) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLoadBalancerPublication: %s", err))
	}

	*out = *target.(*SparseLoadBalancerPublication)
}

type mongoAttributesLoadBalancerPublication struct {
}
type mongoAttributesSparseLoadBalancerPublication struct {
}
