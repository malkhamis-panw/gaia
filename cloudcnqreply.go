package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudCNQReplyIdentity represents the Identity of the object.
var CloudCNQReplyIdentity = elemental.Identity{
	Name:     "cloudcnqreply",
	Category: "cloudcnqreplyies",
	Package:  "vargid",
	Private:  false,
}

// CloudCNQRepliesList represents a list of CloudCNQReplies
type CloudCNQRepliesList []*CloudCNQReply

// Identity returns the identity of the objects in the list.
func (o CloudCNQRepliesList) Identity() elemental.Identity {

	return CloudCNQReplyIdentity
}

// Copy returns a pointer to a copy the CloudCNQRepliesList.
func (o CloudCNQRepliesList) Copy() elemental.Identifiables {

	copy := append(CloudCNQRepliesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudCNQRepliesList.
func (o CloudCNQRepliesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudCNQRepliesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudCNQReply))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudCNQRepliesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudCNQRepliesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudCNQRepliesList converted to SparseCloudCNQRepliesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudCNQRepliesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudCNQRepliesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudCNQReply)
	}

	return out
}

// Version returns the version of the content.
func (o CloudCNQRepliesList) Version() int {

	return 1
}

// CloudCNQReply represents the model of a cloudcnqreply
type CloudCNQReply struct {
	// Prisma Cloud Alert Rule id.
	AlertRuleID string `json:"alertRuleID" msgpack:"alertRuleID" bson:"-" mapstructure:"alertRuleID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The result of the cloud network query.
	CloudGraphResult *CloudGraph `json:"cloudGraphResult" msgpack:"cloudGraphResult" bson:"-" mapstructure:"cloudGraphResult,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Prisma Cloud Policy id.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Result of the last successfully run query.
	ResultTimestamp time.Time `json:"resultTimestamp" msgpack:"resultTimestamp" bson:"ag" mapstructure:"resultTimestamp,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudCNQReply returns a new *CloudCNQReply
func NewCloudCNQReply() *CloudCNQReply {

	return &CloudCNQReply{
		ModelVersion:     1,
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		CloudGraphResult: NewCloudGraph(),
		NormalizedTags:   []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudCNQReply) Identity() elemental.Identity {

	return CloudCNQReplyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudCNQReply) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudCNQReply) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudCNQReply) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudCNQReply{}

	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.PolicyID = o.PolicyID
	s.Protected = o.Protected
	s.ResultTimestamp = o.ResultTimestamp
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudCNQReply) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudCNQReply{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.PolicyID = s.PolicyID
	o.Protected = s.Protected
	o.ResultTimestamp = s.ResultTimestamp
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudCNQReply) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudCNQReply) BleveType() string {

	return "cloudcnqreply"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudCNQReply) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudCNQReply) Doc() string {

	return `The reply object for a Cloud Network Query from the Graph Calculator.`
}

func (o *CloudCNQReply) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudCNQReply) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudCNQReply) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudCNQReply) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudCNQReply) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudCNQReply) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudCNQReply) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudCNQReply) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudCNQReply) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudCNQReply) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudCNQReply) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudCNQReply) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudCNQReply) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudCNQReply) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudCNQReply) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudCNQReply) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudCNQReply{
			AlertRuleID:          &o.AlertRuleID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CloudGraphResult:     o.CloudGraphResult,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			PolicyID:             &o.PolicyID,
			Protected:            &o.Protected,
			ResultTimestamp:      &o.ResultTimestamp,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
		}
	}

	sp := &SparseCloudCNQReply{}
	for _, f := range fields {
		switch f {
		case "alertRuleID":
			sp.AlertRuleID = &(o.AlertRuleID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "cloudGraphResult":
			sp.CloudGraphResult = o.CloudGraphResult
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "protected":
			sp.Protected = &(o.Protected)
		case "resultTimestamp":
			sp.ResultTimestamp = &(o.ResultTimestamp)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudCNQReply to the object.
func (o *CloudCNQReply) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudCNQReply)
	if so.AlertRuleID != nil {
		o.AlertRuleID = *so.AlertRuleID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CloudGraphResult != nil {
		o.CloudGraphResult = so.CloudGraphResult
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ResultTimestamp != nil {
		o.ResultTimestamp = *so.ResultTimestamp
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
}

// DeepCopy returns a deep copy if the CloudCNQReply.
func (o *CloudCNQReply) DeepCopy() *CloudCNQReply {

	if o == nil {
		return nil
	}

	out := &CloudCNQReply{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudCNQReply.
func (o *CloudCNQReply) DeepCopyInto(out *CloudCNQReply) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudCNQReply: %s", err))
	}

	*out = *target.(*CloudCNQReply)
}

// Validate valides the current information stored into the structure.
func (o *CloudCNQReply) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if o.CloudGraphResult != nil {
		elemental.ResetDefaultForZeroValues(o.CloudGraphResult)
		if err := o.CloudGraphResult.Validate(); err != nil {
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
func (*CloudCNQReply) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudCNQReplyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudCNQReplyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudCNQReply) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudCNQReplyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudCNQReply) ValueForAttribute(name string) interface{} {

	switch name {
	case "alertRuleID":
		return o.AlertRuleID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "cloudGraphResult":
		return o.CloudGraphResult
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "policyID":
		return o.PolicyID
	case "protected":
		return o.Protected
	case "resultTimestamp":
		return o.ResultTimestamp
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	}

	return nil
}

// CloudCNQReplyAttributesMap represents the map of attribute for CloudCNQReply.
var CloudCNQReplyAttributesMap = map[string]elemental.AttributeSpecification{
	"AlertRuleID": {
		AllowedChoices: []string{},
		ConvertedName:  "AlertRuleID",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "alertRuleID",
		SubType:        "string",
		Type:           "string",
	},
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CloudGraphResult": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudGraphResult",
		Description:    `The result of the cloud network query.`,
		Exposed:        true,
		Name:           "cloudGraphResult",
		SubType:        "cloudgraph",
		Type:           "ref",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyid",
		ConvertedName:  "PolicyID",
		Description:    `Prisma Cloud Policy id.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"ResultTimestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "ag",
		ConvertedName:  "ResultTimestamp",
		Description:    `Result of the last successfully run query.`,
		Exposed:        true,
		Name:           "resultTimestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// CloudCNQReplyLowerCaseAttributesMap represents the map of attribute for CloudCNQReply.
var CloudCNQReplyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"alertruleid": {
		AllowedChoices: []string{},
		ConvertedName:  "AlertRuleID",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "alertRuleID",
		SubType:        "string",
		Type:           "string",
	},
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"cloudgraphresult": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudGraphResult",
		Description:    `The result of the cloud network query.`,
		Exposed:        true,
		Name:           "cloudGraphResult",
		SubType:        "cloudgraph",
		Type:           "ref",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"policyid": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyid",
		ConvertedName:  "PolicyID",
		Description:    `Prisma Cloud Policy id.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"resulttimestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "ag",
		ConvertedName:  "ResultTimestamp",
		Description:    `Result of the last successfully run query.`,
		Exposed:        true,
		Name:           "resultTimestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseCloudCNQRepliesList represents a list of SparseCloudCNQReplies
type SparseCloudCNQRepliesList []*SparseCloudCNQReply

// Identity returns the identity of the objects in the list.
func (o SparseCloudCNQRepliesList) Identity() elemental.Identity {

	return CloudCNQReplyIdentity
}

// Copy returns a pointer to a copy the SparseCloudCNQRepliesList.
func (o SparseCloudCNQRepliesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudCNQRepliesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudCNQRepliesList.
func (o SparseCloudCNQRepliesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudCNQRepliesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudCNQReply))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudCNQRepliesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudCNQRepliesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudCNQRepliesList converted to CloudCNQRepliesList.
func (o SparseCloudCNQRepliesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudCNQRepliesList) Version() int {

	return 1
}

// SparseCloudCNQReply represents the sparse version of a cloudcnqreply.
type SparseCloudCNQReply struct {
	// Prisma Cloud Alert Rule id.
	AlertRuleID *string `json:"alertRuleID,omitempty" msgpack:"alertRuleID,omitempty" bson:"-" mapstructure:"alertRuleID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The result of the cloud network query.
	CloudGraphResult *CloudGraph `json:"cloudGraphResult,omitempty" msgpack:"cloudGraphResult,omitempty" bson:"-" mapstructure:"cloudGraphResult,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Prisma Cloud Policy id.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Result of the last successfully run query.
	ResultTimestamp *time.Time `json:"resultTimestamp,omitempty" msgpack:"resultTimestamp,omitempty" bson:"ag,omitempty" mapstructure:"resultTimestamp,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudCNQReply returns a new  SparseCloudCNQReply.
func NewSparseCloudCNQReply() *SparseCloudCNQReply {
	return &SparseCloudCNQReply{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudCNQReply) Identity() elemental.Identity {

	return CloudCNQReplyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudCNQReply) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudCNQReply) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudCNQReply) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudCNQReply{}

	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.ResultTimestamp != nil {
		s.ResultTimestamp = o.ResultTimestamp
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudCNQReply) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudCNQReply{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.ResultTimestamp != nil {
		o.ResultTimestamp = s.ResultTimestamp
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudCNQReply) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudCNQReply) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudCNQReply()
	if o.AlertRuleID != nil {
		out.AlertRuleID = *o.AlertRuleID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CloudGraphResult != nil {
		out.CloudGraphResult = o.CloudGraphResult
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ResultTimestamp != nil {
		out.ResultTimestamp = *o.ResultTimestamp
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudCNQReply) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudCNQReply) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudCNQReply) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudCNQReply) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudCNQReply) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudCNQReply) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudCNQReply) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudCNQReply) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// DeepCopy returns a deep copy if the SparseCloudCNQReply.
func (o *SparseCloudCNQReply) DeepCopy() *SparseCloudCNQReply {

	if o == nil {
		return nil
	}

	out := &SparseCloudCNQReply{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudCNQReply.
func (o *SparseCloudCNQReply) DeepCopyInto(out *SparseCloudCNQReply) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudCNQReply: %s", err))
	}

	*out = *target.(*SparseCloudCNQReply)
}

type mongoAttributesCloudCNQReply struct {
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	Namespace            string              `bson:"namespace"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	PolicyID             string              `bson:"policyid"`
	Protected            bool                `bson:"protected"`
	ResultTimestamp      time.Time           `bson:"ag"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
}
type mongoAttributesSparseCloudCNQReply struct {
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	PolicyID             *string              `bson:"policyid,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	ResultTimestamp      *time.Time           `bson:"ag,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
}
