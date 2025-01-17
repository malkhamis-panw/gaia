package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AuthzIdentity represents the Identity of the object.
var AuthzIdentity = elemental.Identity{
	Name:     "authz",
	Category: "authz",
	Package:  "cid",
	Private:  true,
}

// AuthzsList represents a list of Authzs
type AuthzsList []*Authz

// Identity returns the identity of the objects in the list.
func (o AuthzsList) Identity() elemental.Identity {

	return AuthzIdentity
}

// Copy returns a pointer to a copy the AuthzsList.
func (o AuthzsList) Copy() elemental.Identifiables {

	copy := append(AuthzsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthzsList.
func (o AuthzsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuthzsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Authz))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthzsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthzsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuthzsList converted to SparseAuthzsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuthzsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAuthzsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAuthz)
	}

	return out
}

// Version returns the version of the content.
func (o AuthzsList) Version() int {

	return 1
}

// Authz represents the model of a authz
type Authz struct {
	// The list of verified claims.
	Claims []string `json:"claims" msgpack:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// IP of the client.
	ClientIP string `json:"clientIP" msgpack:"clientIP" bson:"-" mapstructure:"clientIP,omitempty"`

	// Return an eventual error.
	Error string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// If the parameter permissions=1 is set, targetIdentity and targetOperation are
	// ignored and this attribute will contain all the permission for the given claims.
	Permissions map[string]map[string]bool `json:"permissions,omitempty" msgpack:"permissions,omitempty" bson:"-" mapstructure:"permissions,omitempty"`

	// Contains the raw resolved and unfiltered PolicyRuleList. It will be empty unless
	// the query parameter `forwardpolicyrules=true` is set.
	// Note that this contains the straight rules that matched the user claims in the
	// target namespace. It can contain rules that are filtered out in later stage
	// because it applies to children namespace or because the rule only applies to a
	// certain source network, or because the token contains restrictions. Unless you
	// really know why you need it, you mostly don't.
	PolicyRulesList PolicyRulesList `json:"policyRulesList,omitempty" msgpack:"policyRulesList,omitempty" bson:"-" mapstructure:"policyRulesList,omitempty"`

	// Sets the namespace restrictions that should apply.
	RestrictedNamespace string `json:"restrictedNamespace" msgpack:"restrictedNamespace" bson:"-" mapstructure:"restrictedNamespace,omitempty"`

	// Sets the networks restrictions that should apply.
	RestrictedNetworks []string `json:"restrictedNetworks" msgpack:"restrictedNetworks" bson:"-" mapstructure:"restrictedNetworks,omitempty"`

	// Sets the permissions restrictions that should apply.
	RestrictedPermissions []string `json:"restrictedPermissions" msgpack:"restrictedPermissions" bson:"-" mapstructure:"restrictedPermissions,omitempty"`

	// The ID of the object to check permission for.
	TargetID string `json:"targetID" msgpack:"targetID" bson:"-" mapstructure:"targetID,omitempty"`

	// The namespace where to check permission from.
	TargetNamespace string `json:"targetNamespace" msgpack:"targetNamespace" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAuthz returns a new *Authz
func NewAuthz() *Authz {

	return &Authz{
		ModelVersion:          1,
		Claims:                []string{},
		Permissions:           map[string]map[string]bool{},
		PolicyRulesList:       PolicyRulesList{},
		RestrictedNetworks:    []string{},
		RestrictedPermissions: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Authz) Identity() elemental.Identity {

	return AuthzIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Authz) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Authz) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Authz) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAuthz{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Authz) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAuthz{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Authz) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Authz) BleveType() string {

	return "authz"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Authz) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Authz) Doc() string {

	return `This is an internal API that is used to resolve to API authorization.`
}

func (o *Authz) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Authz) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuthz{
			Claims:                &o.Claims,
			ClientIP:              &o.ClientIP,
			Error:                 &o.Error,
			Permissions:           &o.Permissions,
			PolicyRulesList:       &o.PolicyRulesList,
			RestrictedNamespace:   &o.RestrictedNamespace,
			RestrictedNetworks:    &o.RestrictedNetworks,
			RestrictedPermissions: &o.RestrictedPermissions,
			TargetID:              &o.TargetID,
			TargetNamespace:       &o.TargetNamespace,
		}
	}

	sp := &SparseAuthz{}
	for _, f := range fields {
		switch f {
		case "claims":
			sp.Claims = &(o.Claims)
		case "clientIP":
			sp.ClientIP = &(o.ClientIP)
		case "error":
			sp.Error = &(o.Error)
		case "permissions":
			sp.Permissions = &(o.Permissions)
		case "policyRulesList":
			sp.PolicyRulesList = &(o.PolicyRulesList)
		case "restrictedNamespace":
			sp.RestrictedNamespace = &(o.RestrictedNamespace)
		case "restrictedNetworks":
			sp.RestrictedNetworks = &(o.RestrictedNetworks)
		case "restrictedPermissions":
			sp.RestrictedPermissions = &(o.RestrictedPermissions)
		case "targetID":
			sp.TargetID = &(o.TargetID)
		case "targetNamespace":
			sp.TargetNamespace = &(o.TargetNamespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAuthz to the object.
func (o *Authz) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuthz)
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.ClientIP != nil {
		o.ClientIP = *so.ClientIP
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Permissions != nil {
		o.Permissions = *so.Permissions
	}
	if so.PolicyRulesList != nil {
		o.PolicyRulesList = *so.PolicyRulesList
	}
	if so.RestrictedNamespace != nil {
		o.RestrictedNamespace = *so.RestrictedNamespace
	}
	if so.RestrictedNetworks != nil {
		o.RestrictedNetworks = *so.RestrictedNetworks
	}
	if so.RestrictedPermissions != nil {
		o.RestrictedPermissions = *so.RestrictedPermissions
	}
	if so.TargetID != nil {
		o.TargetID = *so.TargetID
	}
	if so.TargetNamespace != nil {
		o.TargetNamespace = *so.TargetNamespace
	}
}

// DeepCopy returns a deep copy if the Authz.
func (o *Authz) DeepCopy() *Authz {

	if o == nil {
		return nil
	}

	out := &Authz{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Authz.
func (o *Authz) DeepCopyInto(out *Authz) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Authz: %s", err))
	}

	*out = *target.(*Authz)
}

// Validate valides the current information stored into the structure.
func (o *Authz) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("claims", o.Claims); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	for _, sub := range o.PolicyRulesList {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredString("targetNamespace", o.TargetNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*Authz) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthzAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthzLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Authz) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthzAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Authz) ValueForAttribute(name string) interface{} {

	switch name {
	case "claims":
		return o.Claims
	case "clientIP":
		return o.ClientIP
	case "error":
		return o.Error
	case "permissions":
		return o.Permissions
	case "policyRulesList":
		return o.PolicyRulesList
	case "restrictedNamespace":
		return o.RestrictedNamespace
	case "restrictedNetworks":
		return o.RestrictedNetworks
	case "restrictedPermissions":
		return o.RestrictedPermissions
	case "targetID":
		return o.TargetID
	case "targetNamespace":
		return o.TargetNamespace
	}

	return nil
}

// AuthzAttributesMap represents the map of attribute for Authz.
var AuthzAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `The list of verified claims.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"ClientIP": {
		AllowedChoices: []string{},
		ConvertedName:  "ClientIP",
		Description:    `IP of the client.`,
		Exposed:        true,
		Name:           "clientIP",
		Type:           "string",
	},
	"Error": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Return an eventual error.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"Permissions": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Permissions",
		Description: `If the parameter permissions=1 is set, targetIdentity and targetOperation are
ignored and this attribute will contain all the permission for the given claims.`,
		Exposed:  true,
		Name:     "permissions",
		ReadOnly: true,
		SubType:  "map[string]map[string]bool",
		Type:     "external",
	},
	"PolicyRulesList": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyRulesList",
		Description: `Contains the raw resolved and unfiltered PolicyRuleList. It will be empty unless
the query parameter ` + "`" + `forwardpolicyrules=true` + "`" + ` is set.
Note that this contains the straight rules that matched the user claims in the
target namespace. It can contain rules that are filtered out in later stage
because it applies to children namespace or because the rule only applies to a
certain source network, or because the token contains restrictions. Unless you
really know why you need it, you mostly don't.`,
		Exposed: true,
		Name:    "policyRulesList",
		SubType: "policyrule",
		Type:    "refList",
	},
	"RestrictedNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNamespace",
		Description:    `Sets the namespace restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNamespace",
		Type:           "string",
	},
	"RestrictedNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNetworks",
		Description:    `Sets the networks restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNetworks",
		SubType:        "string",
		Type:           "list",
	},
	"RestrictedPermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedPermissions",
		Description:    `Sets the permissions restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedPermissions",
		SubType:        "string",
		Type:           "list",
	},
	"TargetID": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		Description:    `The ID of the object to check permission for.`,
		Exposed:        true,
		Name:           "targetID",
		Type:           "string",
	},
	"TargetNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `The namespace where to check permission from.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}

// AuthzLowerCaseAttributesMap represents the map of attribute for Authz.
var AuthzLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `The list of verified claims.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"clientip": {
		AllowedChoices: []string{},
		ConvertedName:  "ClientIP",
		Description:    `IP of the client.`,
		Exposed:        true,
		Name:           "clientIP",
		Type:           "string",
	},
	"error": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Return an eventual error.`,
		Exposed:        true,
		Name:           "error",
		ReadOnly:       true,
		Type:           "string",
	},
	"permissions": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Permissions",
		Description: `If the parameter permissions=1 is set, targetIdentity and targetOperation are
ignored and this attribute will contain all the permission for the given claims.`,
		Exposed:  true,
		Name:     "permissions",
		ReadOnly: true,
		SubType:  "map[string]map[string]bool",
		Type:     "external",
	},
	"policyruleslist": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyRulesList",
		Description: `Contains the raw resolved and unfiltered PolicyRuleList. It will be empty unless
the query parameter ` + "`" + `forwardpolicyrules=true` + "`" + ` is set.
Note that this contains the straight rules that matched the user claims in the
target namespace. It can contain rules that are filtered out in later stage
because it applies to children namespace or because the rule only applies to a
certain source network, or because the token contains restrictions. Unless you
really know why you need it, you mostly don't.`,
		Exposed: true,
		Name:    "policyRulesList",
		SubType: "policyrule",
		Type:    "refList",
	},
	"restrictednamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNamespace",
		Description:    `Sets the namespace restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNamespace",
		Type:           "string",
	},
	"restrictednetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedNetworks",
		Description:    `Sets the networks restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedNetworks",
		SubType:        "string",
		Type:           "list",
	},
	"restrictedpermissions": {
		AllowedChoices: []string{},
		ConvertedName:  "RestrictedPermissions",
		Description:    `Sets the permissions restrictions that should apply.`,
		Exposed:        true,
		Name:           "restrictedPermissions",
		SubType:        "string",
		Type:           "list",
	},
	"targetid": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		Description:    `The ID of the object to check permission for.`,
		Exposed:        true,
		Name:           "targetID",
		Type:           "string",
	},
	"targetnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description:    `The namespace where to check permission from.`,
		Exposed:        true,
		Name:           "targetNamespace",
		Required:       true,
		Type:           "string",
	},
}

// SparseAuthzsList represents a list of SparseAuthzs
type SparseAuthzsList []*SparseAuthz

// Identity returns the identity of the objects in the list.
func (o SparseAuthzsList) Identity() elemental.Identity {

	return AuthzIdentity
}

// Copy returns a pointer to a copy the SparseAuthzsList.
func (o SparseAuthzsList) Copy() elemental.Identifiables {

	copy := append(SparseAuthzsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuthzsList.
func (o SparseAuthzsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuthzsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuthz))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuthzsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuthzsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuthzsList converted to AuthzsList.
func (o SparseAuthzsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuthzsList) Version() int {

	return 1
}

// SparseAuthz represents the sparse version of a authz.
type SparseAuthz struct {
	// The list of verified claims.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// IP of the client.
	ClientIP *string `json:"clientIP,omitempty" msgpack:"clientIP,omitempty" bson:"-" mapstructure:"clientIP,omitempty"`

	// Return an eventual error.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// If the parameter permissions=1 is set, targetIdentity and targetOperation are
	// ignored and this attribute will contain all the permission for the given claims.
	Permissions *map[string]map[string]bool `json:"permissions,omitempty" msgpack:"permissions,omitempty" bson:"-" mapstructure:"permissions,omitempty"`

	// Contains the raw resolved and unfiltered PolicyRuleList. It will be empty unless
	// the query parameter `forwardpolicyrules=true` is set.
	// Note that this contains the straight rules that matched the user claims in the
	// target namespace. It can contain rules that are filtered out in later stage
	// because it applies to children namespace or because the rule only applies to a
	// certain source network, or because the token contains restrictions. Unless you
	// really know why you need it, you mostly don't.
	PolicyRulesList *PolicyRulesList `json:"policyRulesList,omitempty" msgpack:"policyRulesList,omitempty" bson:"-" mapstructure:"policyRulesList,omitempty"`

	// Sets the namespace restrictions that should apply.
	RestrictedNamespace *string `json:"restrictedNamespace,omitempty" msgpack:"restrictedNamespace,omitempty" bson:"-" mapstructure:"restrictedNamespace,omitempty"`

	// Sets the networks restrictions that should apply.
	RestrictedNetworks *[]string `json:"restrictedNetworks,omitempty" msgpack:"restrictedNetworks,omitempty" bson:"-" mapstructure:"restrictedNetworks,omitempty"`

	// Sets the permissions restrictions that should apply.
	RestrictedPermissions *[]string `json:"restrictedPermissions,omitempty" msgpack:"restrictedPermissions,omitempty" bson:"-" mapstructure:"restrictedPermissions,omitempty"`

	// The ID of the object to check permission for.
	TargetID *string `json:"targetID,omitempty" msgpack:"targetID,omitempty" bson:"-" mapstructure:"targetID,omitempty"`

	// The namespace where to check permission from.
	TargetNamespace *string `json:"targetNamespace,omitempty" msgpack:"targetNamespace,omitempty" bson:"-" mapstructure:"targetNamespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAuthz returns a new  SparseAuthz.
func NewSparseAuthz() *SparseAuthz {
	return &SparseAuthz{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuthz) Identity() elemental.Identity {

	return AuthzIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuthz) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuthz) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuthz) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAuthz{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuthz) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAuthz{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAuthz) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuthz) ToPlain() elemental.PlainIdentifiable {

	out := NewAuthz()
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.ClientIP != nil {
		out.ClientIP = *o.ClientIP
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Permissions != nil {
		out.Permissions = *o.Permissions
	}
	if o.PolicyRulesList != nil {
		out.PolicyRulesList = *o.PolicyRulesList
	}
	if o.RestrictedNamespace != nil {
		out.RestrictedNamespace = *o.RestrictedNamespace
	}
	if o.RestrictedNetworks != nil {
		out.RestrictedNetworks = *o.RestrictedNetworks
	}
	if o.RestrictedPermissions != nil {
		out.RestrictedPermissions = *o.RestrictedPermissions
	}
	if o.TargetID != nil {
		out.TargetID = *o.TargetID
	}
	if o.TargetNamespace != nil {
		out.TargetNamespace = *o.TargetNamespace
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAuthz.
func (o *SparseAuthz) DeepCopy() *SparseAuthz {

	if o == nil {
		return nil
	}

	out := &SparseAuthz{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuthz.
func (o *SparseAuthz) DeepCopyInto(out *SparseAuthz) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuthz: %s", err))
	}

	*out = *target.(*SparseAuthz)
}

type mongoAttributesAuthz struct {
}
type mongoAttributesSparseAuthz struct {
}
