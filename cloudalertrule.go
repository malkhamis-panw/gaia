package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAlertRuleIdentity represents the Identity of the object.
var CloudAlertRuleIdentity = elemental.Identity{
	Name:     "cloudalertrule",
	Category: "cloudalertsrule",
	Package:  "vargid",
	Private:  false,
}

// CloudAlertRulesList represents a list of CloudAlertRules
type CloudAlertRulesList []*CloudAlertRule

// Identity returns the identity of the objects in the list.
func (o CloudAlertRulesList) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Copy returns a pointer to a copy the CloudAlertRulesList.
func (o CloudAlertRulesList) Copy() elemental.Identifiables {

	copy := append(CloudAlertRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudAlertRulesList.
func (o CloudAlertRulesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAlertRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAlertRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAlertRulesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAlertRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudAlertRulesList converted to SparseCloudAlertRulesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAlertRulesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAlertRulesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAlertRule)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAlertRulesList) Version() int {

	return 1
}

// CloudAlertRule represents the model of a cloudalertrule
type CloudAlertRule struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid,omitempty" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// List of accounts associated to alert rule.
	Accountids []string `json:"accountids" msgpack:"accountids" bson:"accountids" mapstructure:"accountids,omitempty"`

	// Prisma Cloud Alert Rule id.
	Alertruleid string `json:"alertruleid" msgpack:"alertruleid" bson:"alertruleid" mapstructure:"alertruleid,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags []string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags,omitempty" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// Alert rule description.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Is Alert Rule enabled.
	Enabled bool `json:"enabled" msgpack:"enabled" bson:"enabled" mapstructure:"enabled,omitempty"`

	// The time that the object was first ingested.
	IngestionTime time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime,omitempty" mapstructure:"ingestionTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the object (optional).
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID string `json:"nativeID" msgpack:"nativeID" bson:"nativeid" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences []string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// List of policy IDs associated to alert rule.
	Policyids []string `json:"policyids" msgpack:"policyids" bson:"policyids" mapstructure:"policyids,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// List of regions where the alert rule is enforced.
	Regions []string `json:"regions" msgpack:"regions" bson:"regions" mapstructure:"regions,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAlertRule returns a new *CloudAlertRule
func NewCloudAlertRule() *CloudAlertRule {

	return &CloudAlertRule{
		ModelVersion:     1,
		AssociatedTags:   []string{},
		CloudTags:        []string{},
		Accountids:       []string{},
		Annotations:      map[string][]string{},
		MigrationsLog:    map[string]string{},
		Policyids:        []string{},
		NormalizedTags:   []string{},
		PolicyReferences: []string{},
		Regions:          []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudAlertRule) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAlertRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAlertRule) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAlertRule{}

	s.APIID = o.APIID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.VPCID = o.VPCID
	s.AccountID = o.AccountID
	s.Accountids = o.Accountids
	s.Alertruleid = o.Alertruleid
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CloudTags = o.CloudTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CustomerID = o.CustomerID
	s.Description = o.Description
	s.Enabled = o.Enabled
	s.IngestionTime = o.IngestionTime
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.PolicyReferences = o.PolicyReferences
	s.Policyids = o.Policyids
	s.Protected = o.Protected
	s.RegionName = o.RegionName
	s.Regions = o.Regions
	s.ResourceID = o.ResourceID
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAlertRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.APIID = s.APIID
	o.ID = s.ID.Hex()
	o.VPCID = s.VPCID
	o.AccountID = s.AccountID
	o.Accountids = s.Accountids
	o.Alertruleid = s.Alertruleid
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CloudTags = s.CloudTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CustomerID = s.CustomerID
	o.Description = s.Description
	o.Enabled = s.Enabled
	o.IngestionTime = s.IngestionTime
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.PolicyReferences = s.PolicyReferences
	o.Policyids = s.Policyids
	o.Protected = s.Protected
	o.RegionName = s.RegionName
	o.Regions = s.Regions
	o.ResourceID = s.ResourceID
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudAlertRule) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAlertRule) BleveType() string {

	return "cloudalertrule"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAlertRule) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudAlertRule) Doc() string {

	return `Creates a Prisma Cloud alert rule along with policies and accounts associated
with the alert rule.`
}

func (o *CloudAlertRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAPIID returns the APIID of the receiver.
func (o *CloudAlertRule) GetAPIID() int {

	return o.APIID
}

// SetAPIID sets the property APIID of the receiver using the given value.
func (o *CloudAlertRule) SetAPIID(APIID int) {

	o.APIID = APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *CloudAlertRule) GetVPCID() string {

	return o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the given value.
func (o *CloudAlertRule) SetVPCID(VPCID string) {

	o.VPCID = VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudAlertRule) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudAlertRule) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudAlertRule) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudAlertRule) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudAlertRule) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudAlertRule) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *CloudAlertRule) GetCloudTags() []string {

	return o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the given value.
func (o *CloudAlertRule) SetCloudTags(cloudTags []string) {

	o.CloudTags = cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudAlertRule) GetCloudType() string {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudAlertRule) SetCloudType(cloudType string) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudAlertRule) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRule) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudAlertRule) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudAlertRule) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *CloudAlertRule) GetIngestionTime() time.Time {

	return o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the given value.
func (o *CloudAlertRule) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudAlertRule) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudAlertRule) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudAlertRule) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudAlertRule) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudAlertRule) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudAlertRule) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudAlertRule) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudAlertRule) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudAlertRule) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudAlertRule) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *CloudAlertRule) GetPolicyReferences() []string {

	return o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the given value.
func (o *CloudAlertRule) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *CloudAlertRule) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudAlertRule) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudAlertRule) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudAlertRule) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudAlertRule) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudAlertRule) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudAlertRule) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRule) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudAlertRule) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudAlertRule) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudAlertRule) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudAlertRule) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAlertRule) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAlertRule{
			APIID:                &o.APIID,
			ID:                   &o.ID,
			VPCID:                &o.VPCID,
			AccountID:            &o.AccountID,
			Accountids:           &o.Accountids,
			Alertruleid:          &o.Alertruleid,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CloudTags:            &o.CloudTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CustomerID:           &o.CustomerID,
			Description:          &o.Description,
			Enabled:              &o.Enabled,
			IngestionTime:        &o.IngestionTime,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			PolicyReferences:     &o.PolicyReferences,
			Policyids:            &o.Policyids,
			Protected:            &o.Protected,
			RegionName:           &o.RegionName,
			Regions:              &o.Regions,
			ResourceID:           &o.ResourceID,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudAlertRule{}
	for _, f := range fields {
		switch f {
		case "APIID":
			sp.APIID = &(o.APIID)
		case "ID":
			sp.ID = &(o.ID)
		case "VPCID":
			sp.VPCID = &(o.VPCID)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "accountids":
			sp.Accountids = &(o.Accountids)
		case "alertruleid":
			sp.Alertruleid = &(o.Alertruleid)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "cloudTags":
			sp.CloudTags = &(o.CloudTags)
		case "cloudType":
			sp.CloudType = &(o.CloudType)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "customerID":
			sp.CustomerID = &(o.CustomerID)
		case "description":
			sp.Description = &(o.Description)
		case "enabled":
			sp.Enabled = &(o.Enabled)
		case "ingestionTime":
			sp.IngestionTime = &(o.IngestionTime)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "nativeID":
			sp.NativeID = &(o.NativeID)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "policyReferences":
			sp.PolicyReferences = &(o.PolicyReferences)
		case "policyids":
			sp.Policyids = &(o.Policyids)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regionName":
			sp.RegionName = &(o.RegionName)
		case "regions":
			sp.Regions = &(o.Regions)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudAlertRule to the object.
func (o *CloudAlertRule) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAlertRule)
	if so.APIID != nil {
		o.APIID = *so.APIID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.VPCID != nil {
		o.VPCID = *so.VPCID
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.Accountids != nil {
		o.Accountids = *so.Accountids
	}
	if so.Alertruleid != nil {
		o.Alertruleid = *so.Alertruleid
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CloudTags != nil {
		o.CloudTags = *so.CloudTags
	}
	if so.CloudType != nil {
		o.CloudType = *so.CloudType
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CustomerID != nil {
		o.CustomerID = *so.CustomerID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Enabled != nil {
		o.Enabled = *so.Enabled
	}
	if so.IngestionTime != nil {
		o.IngestionTime = *so.IngestionTime
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NativeID != nil {
		o.NativeID = *so.NativeID
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.PolicyReferences != nil {
		o.PolicyReferences = *so.PolicyReferences
	}
	if so.Policyids != nil {
		o.Policyids = *so.Policyids
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.RegionName != nil {
		o.RegionName = *so.RegionName
	}
	if so.Regions != nil {
		o.Regions = *so.Regions
	}
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the CloudAlertRule.
func (o *CloudAlertRule) DeepCopy() *CloudAlertRule {

	if o == nil {
		return nil
	}

	out := &CloudAlertRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAlertRule.
func (o *CloudAlertRule) DeepCopyInto(out *CloudAlertRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAlertRule: %s", err))
	}

	*out = *target.(*CloudAlertRule)
}

// Validate valides the current information stored into the structure.
func (o *CloudAlertRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("nativeID", o.NativeID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("nativeID", o.NativeID, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateNativeID("nativeID", o.NativeID); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("regionName", o.RegionName, 256, false); err != nil {
		errors = errors.Append(err)
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
func (*CloudAlertRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAlertRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAlertRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAlertRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAlertRuleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAlertRule) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIID":
		return o.APIID
	case "ID":
		return o.ID
	case "VPCID":
		return o.VPCID
	case "accountID":
		return o.AccountID
	case "accountids":
		return o.Accountids
	case "alertruleid":
		return o.Alertruleid
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "cloudTags":
		return o.CloudTags
	case "cloudType":
		return o.CloudType
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "customerID":
		return o.CustomerID
	case "description":
		return o.Description
	case "enabled":
		return o.Enabled
	case "ingestionTime":
		return o.IngestionTime
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "nativeID":
		return o.NativeID
	case "normalizedTags":
		return o.NormalizedTags
	case "policyReferences":
		return o.PolicyReferences
	case "policyids":
		return o.Policyids
	case "protected":
		return o.Protected
	case "regionName":
		return o.RegionName
	case "regions":
		return o.Regions
	case "resourceID":
		return o.ResourceID
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudAlertRuleAttributesMap represents the map of attribute for CloudAlertRule.
var CloudAlertRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"APIID": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiid",
		ConvertedName:  "APIID",
		Description:    `Prisma Cloud API ID (matches Prisma Cloud API ID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "APIID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"VPCID": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "VPCID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"AccountID": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "accountID",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Accountids": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountids",
		ConvertedName:  "Accountids",
		Description:    `List of accounts associated to alert rule.`,
		Exposed:        true,
		Name:           "accountids",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Alertruleid": {
		AllowedChoices: []string{},
		BSONFieldName:  "alertruleid",
		ConvertedName:  "Alertruleid",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "alertruleid",
		Stored:         true,
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
	"CloudTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtags",
		ConvertedName:  "CloudTags",
		Description:    `Internal representation of object tags retrieved from the cloud provider.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CloudType": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtype",
		ConvertedName:  "CloudType",
		Description:    `Cloud type of the entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudType",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"CustomerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Getter:         true,
		Name:           "customerID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Alert rule description.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"Enabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "enabled",
		ConvertedName:  "Enabled",
		Description:    `Is Alert Rule enabled.`,
		Exposed:        true,
		Name:           "enabled",
		Stored:         true,
		SubType:        "boolean",
		Type:           "boolean",
	},
	"IngestionTime": {
		AllowedChoices: []string{},
		BSONFieldName:  "ingestiontime",
		ConvertedName:  "IngestionTime",
		Description:    `The time that the object was first ingested.`,
		Exposed:        true,
		Getter:         true,
		Name:           "ingestionTime",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the object (optional).`,
		Exposed:        true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
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
	"NativeID": {
		AllowedChoices: []string{},
		BSONFieldName:  "nativeid",
		ConvertedName:  "NativeID",
		Description:    `ID of the cloud provider object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "nativeID",
		Orderable:      true,
		Required:       true,
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
	"PolicyReferences": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyreferences",
		ConvertedName:  "PolicyReferences",
		Description:    `A list of policy references associated with this cloud node.`,
		Exposed:        true,
		Getter:         true,
		Name:           "policyReferences",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Policyids": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyids",
		ConvertedName:  "Policyids",
		Description:    `List of policy IDs associated to alert rule.`,
		Exposed:        true,
		Name:           "policyids",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"RegionName": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionname",
		ConvertedName:  "RegionName",
		Description:    `Region name associated with the entity.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Regions": {
		AllowedChoices: []string{},
		BSONFieldName:  "regions",
		ConvertedName:  "Regions",
		Description:    `List of regions where the alert rule is enforced.`,
		Exposed:        true,
		Name:           "regions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ResourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// CloudAlertRuleLowerCaseAttributesMap represents the map of attribute for CloudAlertRule.
var CloudAlertRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiid": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiid",
		ConvertedName:  "APIID",
		Description:    `Prisma Cloud API ID (matches Prisma Cloud API ID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "APIID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"vpcid": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `ID of the host VPC.`,
		Exposed:        true,
		Getter:         true,
		Name:           "VPCID",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"accountid": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountid",
		ConvertedName:  "AccountID",
		Description:    `Cloud account ID associated with the entity (matches Prisma Cloud accountID).`,
		Exposed:        true,
		Getter:         true,
		Name:           "accountID",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"accountids": {
		AllowedChoices: []string{},
		BSONFieldName:  "accountids",
		ConvertedName:  "Accountids",
		Description:    `List of accounts associated to alert rule.`,
		Exposed:        true,
		Name:           "accountids",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"alertruleid": {
		AllowedChoices: []string{},
		BSONFieldName:  "alertruleid",
		ConvertedName:  "Alertruleid",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "alertruleid",
		Stored:         true,
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
	"cloudtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtags",
		ConvertedName:  "CloudTags",
		Description:    `Internal representation of object tags retrieved from the cloud provider.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"cloudtype": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudtype",
		ConvertedName:  "CloudType",
		Description:    `Cloud type of the entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "cloudType",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"customerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "customerid",
		ConvertedName:  "CustomerID",
		Description:    `Customer ID as identified by Prisma Cloud.`,
		Exposed:        true,
		Getter:         true,
		Name:           "customerID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Alert rule description.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"enabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "enabled",
		ConvertedName:  "Enabled",
		Description:    `Is Alert Rule enabled.`,
		Exposed:        true,
		Name:           "enabled",
		Stored:         true,
		SubType:        "boolean",
		Type:           "boolean",
	},
	"ingestiontime": {
		AllowedChoices: []string{},
		BSONFieldName:  "ingestiontime",
		ConvertedName:  "IngestionTime",
		Description:    `The time that the object was first ingested.`,
		Exposed:        true,
		Getter:         true,
		Name:           "ingestionTime",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the object (optional).`,
		Exposed:        true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
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
	"nativeid": {
		AllowedChoices: []string{},
		BSONFieldName:  "nativeid",
		ConvertedName:  "NativeID",
		Description:    `ID of the cloud provider object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "nativeID",
		Orderable:      true,
		Required:       true,
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
	"policyreferences": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyreferences",
		ConvertedName:  "PolicyReferences",
		Description:    `A list of policy references associated with this cloud node.`,
		Exposed:        true,
		Getter:         true,
		Name:           "policyReferences",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"policyids": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyids",
		ConvertedName:  "Policyids",
		Description:    `List of policy IDs associated to alert rule.`,
		Exposed:        true,
		Name:           "policyids",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"regionname": {
		AllowedChoices: []string{},
		BSONFieldName:  "regionname",
		ConvertedName:  "RegionName",
		Description:    `Region name associated with the entity.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      256,
		Name:           "regionName",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"regions": {
		AllowedChoices: []string{},
		BSONFieldName:  "regions",
		ConvertedName:  "Regions",
		Description:    `List of regions where the alert rule is enforced.`,
		Exposed:        true,
		Name:           "regions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"resourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "resourceid",
		ConvertedName:  "ResourceID",
		Description:    `Prisma Cloud Resource ID.`,
		Exposed:        true,
		Getter:         true,
		Name:           "resourceID",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
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
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseCloudAlertRulesList represents a list of SparseCloudAlertRules
type SparseCloudAlertRulesList []*SparseCloudAlertRule

// Identity returns the identity of the objects in the list.
func (o SparseCloudAlertRulesList) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Copy returns a pointer to a copy the SparseCloudAlertRulesList.
func (o SparseCloudAlertRulesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAlertRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAlertRulesList.
func (o SparseCloudAlertRulesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAlertRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAlertRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAlertRulesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAlertRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudAlertRulesList converted to CloudAlertRulesList.
func (o SparseCloudAlertRulesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAlertRulesList) Version() int {

	return 1
}

// SparseCloudAlertRule represents the sparse version of a cloudalertrule.
type SparseCloudAlertRule struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID *int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID *string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid,omitempty" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID *string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// List of accounts associated to alert rule.
	Accountids *[]string `json:"accountids,omitempty" msgpack:"accountids,omitempty" bson:"accountids,omitempty" mapstructure:"accountids,omitempty"`

	// Prisma Cloud Alert Rule id.
	Alertruleid *string `json:"alertruleid,omitempty" msgpack:"alertruleid,omitempty" bson:"alertruleid,omitempty" mapstructure:"alertruleid,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags *[]string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags,omitempty" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType *string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID *int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// Alert rule description.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Is Alert Rule enabled.
	Enabled *bool `json:"enabled,omitempty" msgpack:"enabled,omitempty" bson:"enabled,omitempty" mapstructure:"enabled,omitempty"`

	// The time that the object was first ingested.
	IngestionTime *time.Time `json:"ingestionTime,omitempty" msgpack:"ingestionTime,omitempty" bson:"ingestiontime,omitempty" mapstructure:"ingestionTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the object (optional).
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// ID of the cloud provider object.
	NativeID *string `json:"nativeID,omitempty" msgpack:"nativeID,omitempty" bson:"nativeid,omitempty" mapstructure:"nativeID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences *[]string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// List of policy IDs associated to alert rule.
	Policyids *[]string `json:"policyids,omitempty" msgpack:"policyids,omitempty" bson:"policyids,omitempty" mapstructure:"policyids,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName *string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// List of regions where the alert rule is enforced.
	Regions *[]string `json:"regions,omitempty" msgpack:"regions,omitempty" bson:"regions,omitempty" mapstructure:"regions,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID *int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudAlertRule returns a new  SparseCloudAlertRule.
func NewSparseCloudAlertRule() *SparseCloudAlertRule {
	return &SparseCloudAlertRule{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAlertRule) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRule) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRule) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAlertRule{}

	if o.APIID != nil {
		s.APIID = o.APIID
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.VPCID != nil {
		s.VPCID = o.VPCID
	}
	if o.AccountID != nil {
		s.AccountID = o.AccountID
	}
	if o.Accountids != nil {
		s.Accountids = o.Accountids
	}
	if o.Alertruleid != nil {
		s.Alertruleid = o.Alertruleid
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CloudTags != nil {
		s.CloudTags = o.CloudTags
	}
	if o.CloudType != nil {
		s.CloudType = o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CustomerID != nil {
		s.CustomerID = o.CustomerID
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Enabled != nil {
		s.Enabled = o.Enabled
	}
	if o.IngestionTime != nil {
		s.IngestionTime = o.IngestionTime
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NativeID != nil {
		s.NativeID = o.NativeID
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.PolicyReferences != nil {
		s.PolicyReferences = o.PolicyReferences
	}
	if o.Policyids != nil {
		s.Policyids = o.Policyids
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.RegionName != nil {
		s.RegionName = o.RegionName
	}
	if o.Regions != nil {
		s.Regions = o.Regions
	}
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAlertRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.APIID != nil {
		o.APIID = s.APIID
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.VPCID != nil {
		o.VPCID = s.VPCID
	}
	if s.AccountID != nil {
		o.AccountID = s.AccountID
	}
	if s.Accountids != nil {
		o.Accountids = s.Accountids
	}
	if s.Alertruleid != nil {
		o.Alertruleid = s.Alertruleid
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CloudTags != nil {
		o.CloudTags = s.CloudTags
	}
	if s.CloudType != nil {
		o.CloudType = s.CloudType
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CustomerID != nil {
		o.CustomerID = s.CustomerID
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Enabled != nil {
		o.Enabled = s.Enabled
	}
	if s.IngestionTime != nil {
		o.IngestionTime = s.IngestionTime
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NativeID != nil {
		o.NativeID = s.NativeID
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.PolicyReferences != nil {
		o.PolicyReferences = s.PolicyReferences
	}
	if s.Policyids != nil {
		o.Policyids = s.Policyids
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.RegionName != nil {
		o.RegionName = s.RegionName
	}
	if s.Regions != nil {
		o.Regions = s.Regions
	}
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudAlertRule) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAlertRule) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAlertRule()
	if o.APIID != nil {
		out.APIID = *o.APIID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.VPCID != nil {
		out.VPCID = *o.VPCID
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.Accountids != nil {
		out.Accountids = *o.Accountids
	}
	if o.Alertruleid != nil {
		out.Alertruleid = *o.Alertruleid
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CloudTags != nil {
		out.CloudTags = *o.CloudTags
	}
	if o.CloudType != nil {
		out.CloudType = *o.CloudType
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CustomerID != nil {
		out.CustomerID = *o.CustomerID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Enabled != nil {
		out.Enabled = *o.Enabled
	}
	if o.IngestionTime != nil {
		out.IngestionTime = *o.IngestionTime
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NativeID != nil {
		out.NativeID = *o.NativeID
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.PolicyReferences != nil {
		out.PolicyReferences = *o.PolicyReferences
	}
	if o.Policyids != nil {
		out.Policyids = *o.Policyids
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.RegionName != nil {
		out.RegionName = *o.RegionName
	}
	if o.Regions != nil {
		out.Regions = *o.Regions
	}
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAPIID returns the APIID of the receiver.
func (o *SparseCloudAlertRule) GetAPIID() (out int) {

	if o.APIID == nil {
		return
	}

	return *o.APIID
}

// SetAPIID sets the property APIID of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetAPIID(APIID int) {

	o.APIID = &APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *SparseCloudAlertRule) GetVPCID() (out string) {

	if o.VPCID == nil {
		return
	}

	return *o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetVPCID(VPCID string) {

	o.VPCID = &VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudAlertRule) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudAlertRule) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudAlertRule) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *SparseCloudAlertRule) GetCloudTags() (out []string) {

	if o.CloudTags == nil {
		return
	}

	return *o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetCloudTags(cloudTags []string) {

	o.CloudTags = &cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudAlertRule) GetCloudType() (out string) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetCloudType(cloudType string) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRule) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudAlertRule) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *SparseCloudAlertRule) GetIngestionTime() (out time.Time) {

	if o.IngestionTime == nil {
		return
	}

	return *o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = &ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudAlertRule) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudAlertRule) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudAlertRule) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudAlertRule) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudAlertRule) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *SparseCloudAlertRule) GetPolicyReferences() (out []string) {

	if o.PolicyReferences == nil {
		return
	}

	return *o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = &policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudAlertRule) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudAlertRule) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudAlertRule) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRule) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudAlertRule) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudAlertRule) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudAlertRule.
func (o *SparseCloudAlertRule) DeepCopy() *SparseCloudAlertRule {

	if o == nil {
		return nil
	}

	out := &SparseCloudAlertRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAlertRule.
func (o *SparseCloudAlertRule) DeepCopyInto(out *SparseCloudAlertRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAlertRule: %s", err))
	}

	*out = *target.(*SparseCloudAlertRule)
}

type mongoAttributesCloudAlertRule struct {
	APIID                int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	VPCID                string              `bson:"vpcid,omitempty"`
	AccountID            string              `bson:"accountid,omitempty"`
	Accountids           []string            `bson:"accountids"`
	Alertruleid          string              `bson:"alertruleid"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CloudTags            []string            `bson:"cloudtags,omitempty"`
	CloudType            string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CustomerID           int                 `bson:"customerid,omitempty"`
	Description          string              `bson:"description"`
	Enabled              bool                `bson:"enabled"`
	IngestionTime        time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        map[string]string   `bson:"migrationslog,omitempty"`
	Name                 string              `bson:"name,omitempty"`
	Namespace            string              `bson:"namespace"`
	NativeID             string              `bson:"nativeid"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	PolicyReferences     []string            `bson:"policyreferences,omitempty"`
	Policyids            []string            `bson:"policyids"`
	Protected            bool                `bson:"protected"`
	RegionName           string              `bson:"regionname,omitempty"`
	Regions              []string            `bson:"regions"`
	ResourceID           int                 `bson:"resourceid,omitempty"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseCloudAlertRule struct {
	APIID                *int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	VPCID                *string              `bson:"vpcid,omitempty"`
	AccountID            *string              `bson:"accountid,omitempty"`
	Accountids           *[]string            `bson:"accountids,omitempty"`
	Alertruleid          *string              `bson:"alertruleid,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CloudTags            *[]string            `bson:"cloudtags,omitempty"`
	CloudType            *string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CustomerID           *int                 `bson:"customerid,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	Enabled              *bool                `bson:"enabled,omitempty"`
	IngestionTime        *time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        *map[string]string   `bson:"migrationslog,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NativeID             *string              `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	PolicyReferences     *[]string            `bson:"policyreferences,omitempty"`
	Policyids            *[]string            `bson:"policyids,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	RegionName           *string              `bson:"regionname,omitempty"`
	Regions              *[]string            `bson:"regions,omitempty"`
	ResourceID           *int                 `bson:"resourceid,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}
