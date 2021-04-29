package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TLSCertificateIdentity represents the Identity of the object.
var TLSCertificateIdentity = elemental.Identity{
	Name:     "tlscertificate",
	Category: "tlscertificates",
	Package:  "service",
	Private:  false,
}

// TLSCertificatesList represents a list of TLSCertificates
type TLSCertificatesList []*TLSCertificate

// Identity returns the identity of the objects in the list.
func (o TLSCertificatesList) Identity() elemental.Identity {

	return TLSCertificateIdentity
}

// Copy returns a pointer to a copy the TLSCertificatesList.
func (o TLSCertificatesList) Copy() elemental.Identifiables {

	copy := append(TLSCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TLSCertificatesList.
func (o TLSCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TLSCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TLSCertificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TLSCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TLSCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TLSCertificatesList converted to SparseTLSCertificatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TLSCertificatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTLSCertificatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTLSCertificate)
	}

	return out
}

// Version returns the version of the content.
func (o TLSCertificatesList) Version() int {

	return 1
}

// TLSCertificate represents the model of a tlscertificate
type TLSCertificate struct {
	// PEM-encoded certificate to expose to the clients for TLS. Only has effect and
	// required if `TLSType` is set to `External`.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// PEM-encoded certificate key associated with `TLSCertificate`. Only has effect
	// and required if `TLSType` is set to `External`.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTLSCertificate returns a new *TLSCertificate
func NewTLSCertificate() *TLSCertificate {

	return &TLSCertificate{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *TLSCertificate) Identity() elemental.Identity {

	return TLSCertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TLSCertificate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TLSCertificate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TLSCertificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTLSCertificate{}

	s.Certificate = o.Certificate
	s.Key = o.Key

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TLSCertificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTLSCertificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Certificate = s.Certificate
	o.Key = s.Key

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TLSCertificate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TLSCertificate) BleveType() string {

	return "tlscertificate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TLSCertificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TLSCertificate) Doc() string {

	return `Represents a certificate public and private key.`
}

func (o *TLSCertificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TLSCertificate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTLSCertificate{
			Certificate: &o.Certificate,
			Key:         &o.Key,
		}
	}

	sp := &SparseTLSCertificate{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "key":
			sp.Key = &(o.Key)
		}
	}

	return sp
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *TLSCertificate) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.Key, err = encrypter.EncryptString(o.Key); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'Key' for 'TLSCertificate' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *TLSCertificate) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.Key, err = encrypter.DecryptString(o.Key); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'Key' for 'TLSCertificate' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// Patch apply the non nil value of a *SparseTLSCertificate to the object.
func (o *TLSCertificate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTLSCertificate)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
}

// DeepCopy returns a deep copy if the TLSCertificate.
func (o *TLSCertificate) DeepCopy() *TLSCertificate {

	if o == nil {
		return nil
	}

	out := &TLSCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TLSCertificate.
func (o *TLSCertificate) DeepCopyInto(out *TLSCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TLSCertificate: %s", err))
	}

	*out = *target.(*TLSCertificate)
}

// Validate valides the current information stored into the structure.
func (o *TLSCertificate) Validate() error {

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
func (*TLSCertificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TLSCertificateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TLSCertificateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TLSCertificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TLSCertificateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TLSCertificate) ValueForAttribute(name string) interface{} {

	switch name {
	case "certificate":
		return o.Certificate
	case "key":
		return o.Key
	}

	return nil
}

// TLSCertificateAttributesMap represents the map of attribute for TLSCertificate.
var TLSCertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": {
		AllowedChoices: []string{},
		BSONFieldName:  "certificate",
		ConvertedName:  "Certificate",
		Description: `PEM-encoded certificate to expose to the clients for TLS. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "certificate",
		Stored:  true,
		Type:    "string",
	},
	"Key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description: `PEM-encoded certificate key associated with ` + "`" + `TLSCertificate` + "`" + `. Only has effect
and required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Encrypted: true,
		Exposed:   true,
		Name:      "key",
		Stored:    true,
		Type:      "string",
	},
}

// TLSCertificateLowerCaseAttributesMap represents the map of attribute for TLSCertificate.
var TLSCertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": {
		AllowedChoices: []string{},
		BSONFieldName:  "certificate",
		ConvertedName:  "Certificate",
		Description: `PEM-encoded certificate to expose to the clients for TLS. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "certificate",
		Stored:  true,
		Type:    "string",
	},
	"key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description: `PEM-encoded certificate key associated with ` + "`" + `TLSCertificate` + "`" + `. Only has effect
and required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Encrypted: true,
		Exposed:   true,
		Name:      "key",
		Stored:    true,
		Type:      "string",
	},
}

// SparseTLSCertificatesList represents a list of SparseTLSCertificates
type SparseTLSCertificatesList []*SparseTLSCertificate

// Identity returns the identity of the objects in the list.
func (o SparseTLSCertificatesList) Identity() elemental.Identity {

	return TLSCertificateIdentity
}

// Copy returns a pointer to a copy the SparseTLSCertificatesList.
func (o SparseTLSCertificatesList) Copy() elemental.Identifiables {

	copy := append(SparseTLSCertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTLSCertificatesList.
func (o SparseTLSCertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTLSCertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTLSCertificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTLSCertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTLSCertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTLSCertificatesList converted to TLSCertificatesList.
func (o SparseTLSCertificatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTLSCertificatesList) Version() int {

	return 1
}

// SparseTLSCertificate represents the sparse version of a tlscertificate.
type SparseTLSCertificate struct {
	// PEM-encoded certificate to expose to the clients for TLS. Only has effect and
	// required if `TLSType` is set to `External`.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"certificate,omitempty" mapstructure:"certificate,omitempty"`

	// PEM-encoded certificate key associated with `TLSCertificate`. Only has effect
	// and required if `TLSType` is set to `External`.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTLSCertificate returns a new  SparseTLSCertificate.
func NewSparseTLSCertificate() *SparseTLSCertificate {
	return &SparseTLSCertificate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTLSCertificate) Identity() elemental.Identity {

	return TLSCertificateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTLSCertificate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTLSCertificate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTLSCertificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTLSCertificate{}

	if o.Certificate != nil {
		s.Certificate = o.Certificate
	}
	if o.Key != nil {
		s.Key = o.Key
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTLSCertificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTLSCertificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Certificate != nil {
		o.Certificate = s.Certificate
	}
	if s.Key != nil {
		o.Key = s.Key
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTLSCertificate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTLSCertificate) ToPlain() elemental.PlainIdentifiable {

	out := NewTLSCertificate()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.Key != nil {
		out.Key = *o.Key
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseTLSCertificate) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.Key, err = encrypter.EncryptString(*o.Key); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'Key' for 'SparseTLSCertificate' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseTLSCertificate) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.Key, err = encrypter.DecryptString(*o.Key); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'Key' for 'SparseTLSCertificate' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DeepCopy returns a deep copy if the SparseTLSCertificate.
func (o *SparseTLSCertificate) DeepCopy() *SparseTLSCertificate {

	if o == nil {
		return nil
	}

	out := &SparseTLSCertificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTLSCertificate.
func (o *SparseTLSCertificate) DeepCopyInto(out *SparseTLSCertificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTLSCertificate: %s", err))
	}

	*out = *target.(*SparseTLSCertificate)
}

type mongoAttributesTLSCertificate struct {
	Certificate string `bson:"certificate"`
	Key         string `bson:"key"`
}
type mongoAttributesSparseTLSCertificate struct {
	Certificate *string `bson:"certificate,omitempty"`
	Key         *string `bson:"key,omitempty"`
}
