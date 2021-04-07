package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAlertRuleCNQIdentity represents the Identity of the object.
var CloudAlertRuleCNQIdentity = elemental.Identity{
	Name:     "cloudalertrulecnq",
	Category: "cloudalertrulecnqs",
	Package:  "vargid",
	Private:  false,
}

// CloudAlertRuleCNQsList represents a list of CloudAlertRuleCNQs
type CloudAlertRuleCNQsList []*CloudAlertRuleCNQ

// Identity returns the identity of the objects in the list.
func (o CloudAlertRuleCNQsList) Identity() elemental.Identity {

	return CloudAlertRuleCNQIdentity
}

// Copy returns a pointer to a copy the CloudAlertRuleCNQsList.
func (o CloudAlertRuleCNQsList) Copy() elemental.Identifiables {

	copy := append(CloudAlertRuleCNQsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudAlertRuleCNQsList.
func (o CloudAlertRuleCNQsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAlertRuleCNQsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAlertRuleCNQ))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAlertRuleCNQsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAlertRuleCNQsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudAlertRuleCNQsList converted to SparseCloudAlertRuleCNQsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAlertRuleCNQsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAlertRuleCNQsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAlertRuleCNQ)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAlertRuleCNQsList) Version() int {

	return 1
}

// CloudAlertRuleCNQ represents the model of a cloudalertrulecnq
type CloudAlertRuleCNQ struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid,omitempty" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Prisma Cloud Alert Rule id.
	AlertRuleID string `json:"alertRuleID" msgpack:"alertRuleID" bson:"alertruleid" mapstructure:"alertRuleID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The cloud network query that should be used.
	CloudNetworkQuery *CloudNetworkQuery `json:"cloudNetworkQuery" msgpack:"cloudNetworkQuery" bson:"cloudnetworkquery" mapstructure:"cloudNetworkQuery,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags []string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags,omitempty" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

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

	// Prisma Cloud Policy id.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences []string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// Result of the last successfully run query.
	ResultTimestamp time.Time `json:"resultTimestamp" msgpack:"resultTimestamp" bson:"ag" mapstructure:"resultTimestamp,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAlertRuleCNQ returns a new *CloudAlertRuleCNQ
func NewCloudAlertRuleCNQ() *CloudAlertRuleCNQ {

	return &CloudAlertRuleCNQ{
		ModelVersion:      1,
		MigrationsLog:     map[string]string{},
		CloudNetworkQuery: NewCloudNetworkQuery(),
		CloudTags:         []string{},
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		PolicyReferences:  []string{},
		NormalizedTags:    []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudAlertRuleCNQ) Identity() elemental.Identity {

	return CloudAlertRuleCNQIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAlertRuleCNQ) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAlertRuleCNQ) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRuleCNQ) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAlertRuleCNQ{}

	s.APIID = o.APIID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.VPCID = o.VPCID
	s.AccountID = o.AccountID
	s.AlertRuleID = o.AlertRuleID
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CloudNetworkQuery = o.CloudNetworkQuery
	s.CloudTags = o.CloudTags
	s.CloudType = o.CloudType
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CustomerID = o.CustomerID
	s.Description = o.Description
	s.IngestionTime = o.IngestionTime
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NativeID = o.NativeID
	s.NormalizedTags = o.NormalizedTags
	s.PolicyID = o.PolicyID
	s.PolicyReferences = o.PolicyReferences
	s.Protected = o.Protected
	s.RegionName = o.RegionName
	s.ResourceID = o.ResourceID
	s.ResultTimestamp = o.ResultTimestamp
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRuleCNQ) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAlertRuleCNQ{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.APIID = s.APIID
	o.ID = s.ID.Hex()
	o.VPCID = s.VPCID
	o.AccountID = s.AccountID
	o.AlertRuleID = s.AlertRuleID
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CloudNetworkQuery = s.CloudNetworkQuery
	o.CloudTags = s.CloudTags
	o.CloudType = s.CloudType
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CustomerID = s.CustomerID
	o.Description = s.Description
	o.IngestionTime = s.IngestionTime
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NativeID = s.NativeID
	o.NormalizedTags = s.NormalizedTags
	o.PolicyID = s.PolicyID
	o.PolicyReferences = s.PolicyReferences
	o.Protected = s.Protected
	o.RegionName = s.RegionName
	o.ResourceID = s.ResourceID
	o.ResultTimestamp = s.ResultTimestamp
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudAlertRuleCNQ) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAlertRuleCNQ) BleveType() string {

	return "cloudalertrulecnq"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAlertRuleCNQ) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudAlertRuleCNQ) Doc() string {

	return `The result object of executing policies in an alert rule.`
}

func (o *CloudAlertRuleCNQ) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAPIID returns the APIID of the receiver.
func (o *CloudAlertRuleCNQ) GetAPIID() int {

	return o.APIID
}

// SetAPIID sets the property APIID of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetAPIID(APIID int) {

	o.APIID = APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *CloudAlertRuleCNQ) GetVPCID() string {

	return o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetVPCID(VPCID string) {

	o.VPCID = VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *CloudAlertRuleCNQ) GetAccountID() string {

	return o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetAccountID(accountID string) {

	o.AccountID = accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudAlertRuleCNQ) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudAlertRuleCNQ) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *CloudAlertRuleCNQ) GetCloudTags() []string {

	return o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetCloudTags(cloudTags []string) {

	o.CloudTags = cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *CloudAlertRuleCNQ) GetCloudType() string {

	return o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetCloudType(cloudType string) {

	o.CloudType = cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudAlertRuleCNQ) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *CloudAlertRuleCNQ) GetCustomerID() int {

	return o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetCustomerID(customerID int) {

	o.CustomerID = customerID
}

// GetDescription returns the Description of the receiver.
func (o *CloudAlertRuleCNQ) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetDescription(description string) {

	o.Description = description
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *CloudAlertRuleCNQ) GetIngestionTime() time.Time {

	return o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudAlertRuleCNQ) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudAlertRuleCNQ) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudAlertRuleCNQ) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *CloudAlertRuleCNQ) GetNativeID() string {

	return o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetNativeID(nativeID string) {

	o.NativeID = nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudAlertRuleCNQ) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *CloudAlertRuleCNQ) GetPolicyReferences() []string {

	return o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *CloudAlertRuleCNQ) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetProtected(protected bool) {

	o.Protected = protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *CloudAlertRuleCNQ) GetRegionName() string {

	return o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetRegionName(regionName string) {

	o.RegionName = regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *CloudAlertRuleCNQ) GetResourceID() int {

	return o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetResourceID(resourceID int) {

	o.ResourceID = resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudAlertRuleCNQ) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudAlertRuleCNQ) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudAlertRuleCNQ) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudAlertRuleCNQ) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAlertRuleCNQ) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAlertRuleCNQ{
			APIID:                &o.APIID,
			ID:                   &o.ID,
			VPCID:                &o.VPCID,
			AccountID:            &o.AccountID,
			AlertRuleID:          &o.AlertRuleID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CloudNetworkQuery:    o.CloudNetworkQuery,
			CloudTags:            &o.CloudTags,
			CloudType:            &o.CloudType,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CustomerID:           &o.CustomerID,
			Description:          &o.Description,
			IngestionTime:        &o.IngestionTime,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NativeID:             &o.NativeID,
			NormalizedTags:       &o.NormalizedTags,
			PolicyID:             &o.PolicyID,
			PolicyReferences:     &o.PolicyReferences,
			Protected:            &o.Protected,
			RegionName:           &o.RegionName,
			ResourceID:           &o.ResourceID,
			ResultTimestamp:      &o.ResultTimestamp,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudAlertRuleCNQ{}
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
		case "alertRuleID":
			sp.AlertRuleID = &(o.AlertRuleID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "cloudNetworkQuery":
			sp.CloudNetworkQuery = o.CloudNetworkQuery
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
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "policyReferences":
			sp.PolicyReferences = &(o.PolicyReferences)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regionName":
			sp.RegionName = &(o.RegionName)
		case "resourceID":
			sp.ResourceID = &(o.ResourceID)
		case "resultTimestamp":
			sp.ResultTimestamp = &(o.ResultTimestamp)
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

// Patch apply the non nil value of a *SparseCloudAlertRuleCNQ to the object.
func (o *CloudAlertRuleCNQ) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAlertRuleCNQ)
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
	if so.AlertRuleID != nil {
		o.AlertRuleID = *so.AlertRuleID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CloudNetworkQuery != nil {
		o.CloudNetworkQuery = so.CloudNetworkQuery
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
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.PolicyReferences != nil {
		o.PolicyReferences = *so.PolicyReferences
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.RegionName != nil {
		o.RegionName = *so.RegionName
	}
	if so.ResourceID != nil {
		o.ResourceID = *so.ResourceID
	}
	if so.ResultTimestamp != nil {
		o.ResultTimestamp = *so.ResultTimestamp
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

// DeepCopy returns a deep copy if the CloudAlertRuleCNQ.
func (o *CloudAlertRuleCNQ) DeepCopy() *CloudAlertRuleCNQ {

	if o == nil {
		return nil
	}

	out := &CloudAlertRuleCNQ{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAlertRuleCNQ.
func (o *CloudAlertRuleCNQ) DeepCopyInto(out *CloudAlertRuleCNQ) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAlertRuleCNQ: %s", err))
	}

	*out = *target.(*CloudAlertRuleCNQ)
}

// Validate valides the current information stored into the structure.
func (o *CloudAlertRuleCNQ) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if o.CloudNetworkQuery != nil {
		elemental.ResetDefaultForZeroValues(o.CloudNetworkQuery)
		if err := o.CloudNetworkQuery.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
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
func (*CloudAlertRuleCNQ) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAlertRuleCNQAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAlertRuleCNQLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAlertRuleCNQ) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAlertRuleCNQAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAlertRuleCNQ) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIID":
		return o.APIID
	case "ID":
		return o.ID
	case "VPCID":
		return o.VPCID
	case "accountID":
		return o.AccountID
	case "alertRuleID":
		return o.AlertRuleID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "cloudNetworkQuery":
		return o.CloudNetworkQuery
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
	case "policyID":
		return o.PolicyID
	case "policyReferences":
		return o.PolicyReferences
	case "protected":
		return o.Protected
	case "regionName":
		return o.RegionName
	case "resourceID":
		return o.ResourceID
	case "resultTimestamp":
		return o.ResultTimestamp
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudAlertRuleCNQAttributesMap represents the map of attribute for CloudAlertRuleCNQ.
var CloudAlertRuleCNQAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AlertRuleID": {
		AllowedChoices: []string{},
		BSONFieldName:  "alertruleid",
		ConvertedName:  "AlertRuleID",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "alertRuleID",
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
	"CloudNetworkQuery": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudnetworkquery",
		ConvertedName:  "CloudNetworkQuery",
		Description:    `The cloud network query that should be used.`,
		Exposed:        true,
		Name:           "cloudNetworkQuery",
		Stored:         true,
		SubType:        "cloudnetworkquery",
		Type:           "ref",
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
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// CloudAlertRuleCNQLowerCaseAttributesMap represents the map of attribute for CloudAlertRuleCNQ.
var CloudAlertRuleCNQLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"alertruleid": {
		AllowedChoices: []string{},
		BSONFieldName:  "alertruleid",
		ConvertedName:  "AlertRuleID",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "alertRuleID",
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
	"cloudnetworkquery": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudnetworkquery",
		ConvertedName:  "CloudNetworkQuery",
		Description:    `The cloud network query that should be used.`,
		Exposed:        true,
		Name:           "cloudNetworkQuery",
		Stored:         true,
		SubType:        "cloudnetworkquery",
		Type:           "ref",
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
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SparseCloudAlertRuleCNQsList represents a list of SparseCloudAlertRuleCNQs
type SparseCloudAlertRuleCNQsList []*SparseCloudAlertRuleCNQ

// Identity returns the identity of the objects in the list.
func (o SparseCloudAlertRuleCNQsList) Identity() elemental.Identity {

	return CloudAlertRuleCNQIdentity
}

// Copy returns a pointer to a copy the SparseCloudAlertRuleCNQsList.
func (o SparseCloudAlertRuleCNQsList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAlertRuleCNQsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAlertRuleCNQsList.
func (o SparseCloudAlertRuleCNQsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAlertRuleCNQsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAlertRuleCNQ))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAlertRuleCNQsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAlertRuleCNQsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudAlertRuleCNQsList converted to CloudAlertRuleCNQsList.
func (o SparseCloudAlertRuleCNQsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAlertRuleCNQsList) Version() int {

	return 1
}

// SparseCloudAlertRuleCNQ represents the sparse version of a cloudalertrulecnq.
type SparseCloudAlertRuleCNQ struct {
	// Prisma Cloud API ID (matches Prisma Cloud API ID).
	APIID *int `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"apiid,omitempty" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the host VPC.
	VPCID *string `json:"VPCID,omitempty" msgpack:"VPCID,omitempty" bson:"vpcid,omitempty" mapstructure:"VPCID,omitempty"`

	// Cloud account ID associated with the entity (matches Prisma Cloud accountID).
	AccountID *string `json:"accountId,omitempty" msgpack:"accountId,omitempty" bson:"accountid,omitempty" mapstructure:"accountId,omitempty"`

	// Prisma Cloud Alert Rule id.
	AlertRuleID *string `json:"alertRuleID,omitempty" msgpack:"alertRuleID,omitempty" bson:"alertruleid,omitempty" mapstructure:"alertRuleID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The cloud network query that should be used.
	CloudNetworkQuery *CloudNetworkQuery `json:"cloudNetworkQuery,omitempty" msgpack:"cloudNetworkQuery,omitempty" bson:"cloudnetworkquery,omitempty" mapstructure:"cloudNetworkQuery,omitempty"`

	// Internal representation of object tags retrieved from the cloud provider.
	CloudTags *[]string `json:"cloudTags,omitempty" msgpack:"cloudTags,omitempty" bson:"cloudtags,omitempty" mapstructure:"cloudTags,omitempty"`

	// Cloud type of the entity.
	CloudType *string `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"cloudtype,omitempty" mapstructure:"cloudType,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Customer ID as identified by Prisma Cloud.
	CustomerID *int `json:"customerID,omitempty" msgpack:"customerID,omitempty" bson:"customerid,omitempty" mapstructure:"customerID,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

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

	// Prisma Cloud Policy id.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// A list of policy references associated with this cloud node.
	PolicyReferences *[]string `json:"policyReferences,omitempty" msgpack:"policyReferences,omitempty" bson:"policyreferences,omitempty" mapstructure:"policyReferences,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Region name associated with the entity.
	RegionName *string `json:"regionName,omitempty" msgpack:"regionName,omitempty" bson:"regionname,omitempty" mapstructure:"regionName,omitempty"`

	// Prisma Cloud Resource ID.
	ResourceID *int `json:"resourceID,omitempty" msgpack:"resourceID,omitempty" bson:"resourceid,omitempty" mapstructure:"resourceID,omitempty"`

	// Result of the last successfully run query.
	ResultTimestamp *time.Time `json:"resultTimestamp,omitempty" msgpack:"resultTimestamp,omitempty" bson:"ag,omitempty" mapstructure:"resultTimestamp,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudAlertRuleCNQ returns a new  SparseCloudAlertRuleCNQ.
func NewSparseCloudAlertRuleCNQ() *SparseCloudAlertRuleCNQ {
	return &SparseCloudAlertRuleCNQ{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAlertRuleCNQ) Identity() elemental.Identity {

	return CloudAlertRuleCNQIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRuleCNQ) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRuleCNQ) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertRuleCNQ) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAlertRuleCNQ{}

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
	if o.AlertRuleID != nil {
		s.AlertRuleID = o.AlertRuleID
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CloudNetworkQuery != nil {
		s.CloudNetworkQuery = o.CloudNetworkQuery
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
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
	}
	if o.PolicyReferences != nil {
		s.PolicyReferences = o.PolicyReferences
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.RegionName != nil {
		s.RegionName = o.RegionName
	}
	if o.ResourceID != nil {
		s.ResourceID = o.ResourceID
	}
	if o.ResultTimestamp != nil {
		s.ResultTimestamp = o.ResultTimestamp
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
func (o *SparseCloudAlertRuleCNQ) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAlertRuleCNQ{}
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
	if s.AlertRuleID != nil {
		o.AlertRuleID = s.AlertRuleID
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CloudNetworkQuery != nil {
		o.CloudNetworkQuery = s.CloudNetworkQuery
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
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
	}
	if s.PolicyReferences != nil {
		o.PolicyReferences = s.PolicyReferences
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.RegionName != nil {
		o.RegionName = s.RegionName
	}
	if s.ResourceID != nil {
		o.ResourceID = s.ResourceID
	}
	if s.ResultTimestamp != nil {
		o.ResultTimestamp = s.ResultTimestamp
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
func (o *SparseCloudAlertRuleCNQ) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAlertRuleCNQ) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAlertRuleCNQ()
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
	if o.AlertRuleID != nil {
		out.AlertRuleID = *o.AlertRuleID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CloudNetworkQuery != nil {
		out.CloudNetworkQuery = o.CloudNetworkQuery
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
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.PolicyReferences != nil {
		out.PolicyReferences = *o.PolicyReferences
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.RegionName != nil {
		out.RegionName = *o.RegionName
	}
	if o.ResourceID != nil {
		out.ResourceID = *o.ResourceID
	}
	if o.ResultTimestamp != nil {
		out.ResultTimestamp = *o.ResultTimestamp
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
func (o *SparseCloudAlertRuleCNQ) GetAPIID() (out int) {

	if o.APIID == nil {
		return
	}

	return *o.APIID
}

// SetAPIID sets the property APIID of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetAPIID(APIID int) {

	o.APIID = &APIID
}

// GetVPCID returns the VPCID of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetVPCID() (out string) {

	if o.VPCID == nil {
		return
	}

	return *o.VPCID
}

// SetVPCID sets the property VPCID of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetVPCID(VPCID string) {

	o.VPCID = &VPCID
}

// GetAccountID returns the AccountID of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetAccountID() (out string) {

	if o.AccountID == nil {
		return
	}

	return *o.AccountID
}

// SetAccountID sets the property AccountID of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetAccountID(accountID string) {

	o.AccountID = &accountID
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCloudTags returns the CloudTags of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetCloudTags() (out []string) {

	if o.CloudTags == nil {
		return
	}

	return *o.CloudTags
}

// SetCloudTags sets the property CloudTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetCloudTags(cloudTags []string) {

	o.CloudTags = &cloudTags
}

// GetCloudType returns the CloudType of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetCloudType() (out string) {

	if o.CloudType == nil {
		return
	}

	return *o.CloudType
}

// SetCloudType sets the property CloudType of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetCloudType(cloudType string) {

	o.CloudType = &cloudType
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCustomerID returns the CustomerID of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetCustomerID() (out int) {

	if o.CustomerID == nil {
		return
	}

	return *o.CustomerID
}

// SetCustomerID sets the property CustomerID of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetCustomerID(customerID int) {

	o.CustomerID = &customerID
}

// GetDescription returns the Description of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetDescription(description string) {

	o.Description = &description
}

// GetIngestionTime returns the IngestionTime of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetIngestionTime() (out time.Time) {

	if o.IngestionTime == nil {
		return
	}

	return *o.IngestionTime
}

// SetIngestionTime sets the property IngestionTime of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetIngestionTime(ingestionTime time.Time) {

	o.IngestionTime = &ingestionTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNativeID returns the NativeID of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetNativeID() (out string) {

	if o.NativeID == nil {
		return
	}

	return *o.NativeID
}

// SetNativeID sets the property NativeID of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetNativeID(nativeID string) {

	o.NativeID = &nativeID
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPolicyReferences returns the PolicyReferences of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetPolicyReferences() (out []string) {

	if o.PolicyReferences == nil {
		return
	}

	return *o.PolicyReferences
}

// SetPolicyReferences sets the property PolicyReferences of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetPolicyReferences(policyReferences []string) {

	o.PolicyReferences = &policyReferences
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetRegionName returns the RegionName of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetRegionName() (out string) {

	if o.RegionName == nil {
		return
	}

	return *o.RegionName
}

// SetRegionName sets the property RegionName of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetRegionName(regionName string) {

	o.RegionName = &regionName
}

// GetResourceID returns the ResourceID of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetResourceID() (out int) {

	if o.ResourceID == nil {
		return
	}

	return *o.ResourceID
}

// SetResourceID sets the property ResourceID of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetResourceID(resourceID int) {

	o.ResourceID = &resourceID
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudAlertRuleCNQ) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudAlertRuleCNQ) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudAlertRuleCNQ.
func (o *SparseCloudAlertRuleCNQ) DeepCopy() *SparseCloudAlertRuleCNQ {

	if o == nil {
		return nil
	}

	out := &SparseCloudAlertRuleCNQ{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAlertRuleCNQ.
func (o *SparseCloudAlertRuleCNQ) DeepCopyInto(out *SparseCloudAlertRuleCNQ) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAlertRuleCNQ: %s", err))
	}

	*out = *target.(*SparseCloudAlertRuleCNQ)
}

type mongoAttributesCloudAlertRuleCNQ struct {
	APIID                int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	VPCID                string              `bson:"vpcid,omitempty"`
	AccountID            string              `bson:"accountid,omitempty"`
	AlertRuleID          string              `bson:"alertruleid"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CloudNetworkQuery    *CloudNetworkQuery  `bson:"cloudnetworkquery"`
	CloudTags            []string            `bson:"cloudtags,omitempty"`
	CloudType            string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CustomerID           int                 `bson:"customerid,omitempty"`
	Description          string              `bson:"description"`
	IngestionTime        time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        map[string]string   `bson:"migrationslog,omitempty"`
	Name                 string              `bson:"name,omitempty"`
	Namespace            string              `bson:"namespace"`
	NativeID             string              `bson:"nativeid"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	PolicyID             string              `bson:"policyid"`
	PolicyReferences     []string            `bson:"policyreferences,omitempty"`
	Protected            bool                `bson:"protected"`
	RegionName           string              `bson:"regionname,omitempty"`
	ResourceID           int                 `bson:"resourceid,omitempty"`
	ResultTimestamp      time.Time           `bson:"ag"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseCloudAlertRuleCNQ struct {
	APIID                *int                 `bson:"apiid,omitempty"`
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	VPCID                *string              `bson:"vpcid,omitempty"`
	AccountID            *string              `bson:"accountid,omitempty"`
	AlertRuleID          *string              `bson:"alertruleid,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CloudNetworkQuery    *CloudNetworkQuery   `bson:"cloudnetworkquery,omitempty"`
	CloudTags            *[]string            `bson:"cloudtags,omitempty"`
	CloudType            *string              `bson:"cloudtype,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CustomerID           *int                 `bson:"customerid,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	IngestionTime        *time.Time           `bson:"ingestiontime,omitempty"`
	MigrationsLog        *map[string]string   `bson:"migrationslog,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NativeID             *string              `bson:"nativeid,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	PolicyID             *string              `bson:"policyid,omitempty"`
	PolicyReferences     *[]string            `bson:"policyreferences,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	RegionName           *string              `bson:"regionname,omitempty"`
	ResourceID           *int                 `bson:"resourceid,omitempty"`
	ResultTimestamp      *time.Time           `bson:"ag,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}
