package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudNetworkQueryEffectiveActionValue represents the possible values for attribute "effectiveAction".
type CloudNetworkQueryEffectiveActionValue string

const (
	// CloudNetworkQueryEffectiveActionAll represents the value All.
	CloudNetworkQueryEffectiveActionAll CloudNetworkQueryEffectiveActionValue = "All"

	// CloudNetworkQueryEffectiveActionReachableAndAllowed represents the value ReachableAndAllowed.
	CloudNetworkQueryEffectiveActionReachableAndAllowed CloudNetworkQueryEffectiveActionValue = "ReachableAndAllowed"

	// CloudNetworkQueryEffectiveActionReachableOnly represents the value ReachableOnly.
	CloudNetworkQueryEffectiveActionReachableOnly CloudNetworkQueryEffectiveActionValue = "ReachableOnly"
)

// CloudNetworkQueryTypeValue represents the possible values for attribute "type".
type CloudNetworkQueryTypeValue string

const (
	// CloudNetworkQueryTypeCompressedGraph represents the value CompressedGraph.
	CloudNetworkQueryTypeCompressedGraph CloudNetworkQueryTypeValue = "CompressedGraph"

	// CloudNetworkQueryTypeFullGraph represents the value FullGraph.
	CloudNetworkQueryTypeFullGraph CloudNetworkQueryTypeValue = "FullGraph"

	// CloudNetworkQueryTypeSummary represents the value Summary.
	CloudNetworkQueryTypeSummary CloudNetworkQueryTypeValue = "Summary"
)

// CloudNetworkQueryIdentity represents the Identity of the object.
var CloudNetworkQueryIdentity = elemental.Identity{
	Name:     "cloudnetworkquery",
	Category: "cloudnetworkqueries",
	Package:  "yeul",
	Private:  false,
}

// CloudNetworkQueriesList represents a list of CloudNetworkQueries
type CloudNetworkQueriesList []*CloudNetworkQuery

// Identity returns the identity of the objects in the list.
func (o CloudNetworkQueriesList) Identity() elemental.Identity {

	return CloudNetworkQueryIdentity
}

// Copy returns a pointer to a copy the CloudNetworkQueriesList.
func (o CloudNetworkQueriesList) Copy() elemental.Identifiables {

	copy := append(CloudNetworkQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudNetworkQueriesList.
func (o CloudNetworkQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudNetworkQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudNetworkQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudNetworkQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudNetworkQueriesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudNetworkQueriesList converted to SparseCloudNetworkQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudNetworkQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudNetworkQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudNetworkQuery)
	}

	return out
}

// Version returns the version of the content.
func (o CloudNetworkQueriesList) Version() int {

	return 1
}

// CloudNetworkQuery represents the model of a cloudnetworkquery
type CloudNetworkQuery struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The destination IP of a trace route request. Might not always be an endpoint.
	DestinationIP string `json:"destinationIP" msgpack:"destinationIP" bson:"destinationip" mapstructure:"destinationIP,omitempty"`

	// A filter for selecting destinations for the query.
	DestinationSelector *CloudNetworkQueryFilter `json:"destinationSelector" msgpack:"destinationSelector" bson:"destinationselector" mapstructure:"destinationSelector,omitempty"`

	// Filters the results based on the effective action. 'ReachableAndAllowed' means
	// that a destination is both reachable and allowed by security rules.
	// 'UnreachableOrRejected' means that the destination is either not reachable or
	// rejected by security rules. 'ReachableOnly' means that all destinations that are
	// reachable will be returned irrespective of their security verdict.
	// 'UnreachableOnly' means that only unreachable destinations will be returned.
	EffectiveAction CloudNetworkQueryEffectiveActionValue `json:"effectiveAction" msgpack:"effectiveAction" bson:"effectiveaction" mapstructure:"effectiveAction,omitempty"`

	// If set, the evaluation will exclude enterprise IPs from the effective
	// permissions.
	ExcludeEnterpriseIPs bool `json:"excludeEnterpriseIPs" msgpack:"excludeEnterpriseIPs" bson:"excludeenterpriseips" mapstructure:"excludeEnterpriseIPs,omitempty"`

	// List of networks that should be a excluded from the calculation if the source or
	// destination is a network.
	ExcludedNetworks []string `json:"excludedNetworks" msgpack:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts []string `json:"protocolPorts" msgpack:"protocolPorts" bson:"protocolports" mapstructure:"protocolPorts,omitempty"`

	// The RQL string for this query as a reference.
	RawRQL string `json:"-" msgpack:"-" bson:"rawrql" mapstructure:"-,omitempty"`

	// The source IP of a trace route request. Might not be always and endpoint.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"sourceip" mapstructure:"sourceIP,omitempty"`

	// A filter for selecting the sources of the request.
	SourceSelector *CloudNetworkQueryFilter `json:"sourceSelector" msgpack:"sourceSelector" bson:"sourceselector" mapstructure:"sourceSelector,omitempty"`

	// Indicates the type of results that should be provided by the query.
	Type CloudNetworkQueryTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudNetworkQuery returns a new *CloudNetworkQuery
func NewCloudNetworkQuery() *CloudNetworkQuery {

	return &CloudNetworkQuery{
		ModelVersion:        1,
		Annotations:         map[string][]string{},
		DestinationSelector: NewCloudNetworkQueryFilter(),
		AssociatedTags:      []string{},
		ExcludedNetworks:    []string{},
		EffectiveAction:     CloudNetworkQueryEffectiveActionReachableOnly,
		NormalizedTags:      []string{},
		ProtocolPorts:       []string{},
		MigrationsLog:       map[string]string{},
		SourceSelector:      NewCloudNetworkQueryFilter(),
		Type:                CloudNetworkQueryTypeSummary,
	}
}

// Identity returns the Identity of the object.
func (o *CloudNetworkQuery) Identity() elemental.Identity {

	return CloudNetworkQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudNetworkQuery) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudNetworkQuery) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudNetworkQuery{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.Description = o.Description
	s.DestinationIP = o.DestinationIP
	s.DestinationSelector = o.DestinationSelector
	s.EffectiveAction = o.EffectiveAction
	s.ExcludeEnterpriseIPs = o.ExcludeEnterpriseIPs
	s.ExcludedNetworks = o.ExcludedNetworks
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.ProtocolPorts = o.ProtocolPorts
	s.RawRQL = o.RawRQL
	s.SourceIP = o.SourceIP
	s.SourceSelector = o.SourceSelector
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudNetworkQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudNetworkQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.Description = s.Description
	o.DestinationIP = s.DestinationIP
	o.DestinationSelector = s.DestinationSelector
	o.EffectiveAction = s.EffectiveAction
	o.ExcludeEnterpriseIPs = s.ExcludeEnterpriseIPs
	o.ExcludedNetworks = s.ExcludedNetworks
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.ProtocolPorts = s.ProtocolPorts
	o.RawRQL = s.RawRQL
	o.SourceIP = s.SourceIP
	o.SourceSelector = s.SourceSelector
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudNetworkQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudNetworkQuery) BleveType() string {

	return "cloudnetworkquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudNetworkQuery) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudNetworkQuery) Doc() string {

	return `Provides the parameters for an effective network permissions query.`
}

func (o *CloudNetworkQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudNetworkQuery) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudNetworkQuery) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudNetworkQuery) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudNetworkQuery) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudNetworkQuery) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudNetworkQuery) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetDescription returns the Description of the receiver.
func (o *CloudNetworkQuery) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *CloudNetworkQuery) SetDescription(description string) {

	o.Description = description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudNetworkQuery) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudNetworkQuery) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudNetworkQuery) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudNetworkQuery) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudNetworkQuery) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudNetworkQuery) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudNetworkQuery) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudNetworkQuery) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudNetworkQuery) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudNetworkQuery) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudNetworkQuery) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudNetworkQuery) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudNetworkQuery) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudNetworkQuery) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudNetworkQuery) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudNetworkQuery) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudNetworkQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudNetworkQuery{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			Description:          &o.Description,
			DestinationIP:        &o.DestinationIP,
			DestinationSelector:  o.DestinationSelector,
			EffectiveAction:      &o.EffectiveAction,
			ExcludeEnterpriseIPs: &o.ExcludeEnterpriseIPs,
			ExcludedNetworks:     &o.ExcludedNetworks,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			ProtocolPorts:        &o.ProtocolPorts,
			RawRQL:               &o.RawRQL,
			SourceIP:             &o.SourceIP,
			SourceSelector:       o.SourceSelector,
			Type:                 &o.Type,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseCloudNetworkQuery{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "description":
			sp.Description = &(o.Description)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "destinationSelector":
			sp.DestinationSelector = o.DestinationSelector
		case "effectiveAction":
			sp.EffectiveAction = &(o.EffectiveAction)
		case "excludeEnterpriseIPs":
			sp.ExcludeEnterpriseIPs = &(o.ExcludeEnterpriseIPs)
		case "excludedNetworks":
			sp.ExcludedNetworks = &(o.ExcludedNetworks)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "protocolPorts":
			sp.ProtocolPorts = &(o.ProtocolPorts)
		case "rawRQL":
			sp.RawRQL = &(o.RawRQL)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "sourceSelector":
			sp.SourceSelector = o.SourceSelector
		case "type":
			sp.Type = &(o.Type)
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

// Patch apply the non nil value of a *SparseCloudNetworkQuery to the object.
func (o *CloudNetworkQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudNetworkQuery)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
	}
	if so.DestinationSelector != nil {
		o.DestinationSelector = so.DestinationSelector
	}
	if so.EffectiveAction != nil {
		o.EffectiveAction = *so.EffectiveAction
	}
	if so.ExcludeEnterpriseIPs != nil {
		o.ExcludeEnterpriseIPs = *so.ExcludeEnterpriseIPs
	}
	if so.ExcludedNetworks != nil {
		o.ExcludedNetworks = *so.ExcludedNetworks
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ProtocolPorts != nil {
		o.ProtocolPorts = *so.ProtocolPorts
	}
	if so.RawRQL != nil {
		o.RawRQL = *so.RawRQL
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.SourceSelector != nil {
		o.SourceSelector = so.SourceSelector
	}
	if so.Type != nil {
		o.Type = *so.Type
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

// DeepCopy returns a deep copy if the CloudNetworkQuery.
func (o *CloudNetworkQuery) DeepCopy() *CloudNetworkQuery {

	if o == nil {
		return nil
	}

	out := &CloudNetworkQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudNetworkQuery.
func (o *CloudNetworkQuery) DeepCopyInto(out *CloudNetworkQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudNetworkQuery: %s", err))
	}

	*out = *target.(*CloudNetworkQuery)
}

// Validate valides the current information stored into the structure.
func (o *CloudNetworkQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("destinationIP", o.DestinationIP); err != nil {
		errors = errors.Append(err)
	}

	if o.DestinationSelector != nil {
		elemental.ResetDefaultForZeroValues(o.DestinationSelector)
		if err := o.DestinationSelector.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("effectiveAction", string(o.EffectiveAction), []string{"ReachableAndAllowed", "ReachableOnly", "All"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIPList("excludedNetworks", o.ExcludedNetworks); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateServicePorts("protocolPorts", o.ProtocolPorts); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRorIP("sourceIP", o.SourceIP); err != nil {
		errors = errors.Append(err)
	}

	if o.SourceSelector != nil {
		elemental.ResetDefaultForZeroValues(o.SourceSelector)
		if err := o.SourceSelector.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Summary", "CompressedGraph", "FullGraph"}, false); err != nil {
		errors = errors.Append(err)
	}

	// Custom object validation.
	if err := ValidateCloudNetworkQueryEntity(o); err != nil {
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
func (*CloudNetworkQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudNetworkQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudNetworkQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudNetworkQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudNetworkQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudNetworkQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "description":
		return o.Description
	case "destinationIP":
		return o.DestinationIP
	case "destinationSelector":
		return o.DestinationSelector
	case "effectiveAction":
		return o.EffectiveAction
	case "excludeEnterpriseIPs":
		return o.ExcludeEnterpriseIPs
	case "excludedNetworks":
		return o.ExcludedNetworks
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "protocolPorts":
		return o.ProtocolPorts
	case "rawRQL":
		return o.RawRQL
	case "sourceIP":
		return o.SourceIP
	case "sourceSelector":
		return o.SourceSelector
	case "type":
		return o.Type
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudNetworkQueryAttributesMap represents the map of attribute for CloudNetworkQuery.
var CloudNetworkQueryAttributesMap = map[string]elemental.AttributeSpecification{
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
	"DestinationIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationip",
		ConvertedName:  "DestinationIP",
		Description:    `The destination IP of a trace route request. Might not always be an endpoint.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"DestinationSelector": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationselector",
		ConvertedName:  "DestinationSelector",
		Description:    `A filter for selecting destinations for the query.`,
		Exposed:        true,
		Name:           "destinationSelector",
		Stored:         true,
		SubType:        "cloudnetworkqueryfilter",
		Type:           "ref",
	},
	"EffectiveAction": {
		AllowedChoices: []string{"ReachableAndAllowed", "ReachableOnly", "All"},
		BSONFieldName:  "effectiveaction",
		ConvertedName:  "EffectiveAction",
		DefaultValue:   CloudNetworkQueryEffectiveActionReachableOnly,
		Description: `Filters the results based on the effective action. 'ReachableAndAllowed' means
that a destination is both reachable and allowed by security rules.
'UnreachableOrRejected' means that the destination is either not reachable or
rejected by security rules. 'ReachableOnly' means that all destinations that are
reachable will be returned irrespective of their security verdict.
'UnreachableOnly' means that only unreachable destinations will be returned.`,
		Exposed: true,
		Name:    "effectiveAction",
		Stored:  true,
		Type:    "enum",
	},
	"ExcludeEnterpriseIPs": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludeenterpriseips",
		ConvertedName:  "ExcludeEnterpriseIPs",
		Description: `If set, the evaluation will exclude enterprise IPs from the effective
permissions.`,
		Exposed: true,
		Name:    "excludeEnterpriseIPs",
		Stored:  true,
		Type:    "boolean",
	},
	"ExcludedNetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludednetworks",
		ConvertedName:  "ExcludedNetworks",
		Description: `List of networks that should be a excluded from the calculation if the source or
destination is a network.`,
		Exposed: true,
		Name:    "excludedNetworks",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
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
	"ProtocolPorts": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocolports",
		ConvertedName:  "ProtocolPorts",
		Description: `Represents the ports and protocols this policy applies to. Protocol/ports are
defined as tcp/80, udp/22. For protocols that do not have ports, the port
designation
is not allowed.`,
		Exposed: true,
		Name:    "protocolPorts",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"RawRQL": {
		AllowedChoices: []string{},
		BSONFieldName:  "rawrql",
		ConvertedName:  "RawRQL",
		Description:    `The RQL string for this query as a reference.`,
		Name:           "rawRQL",
		Stored:         true,
		Type:           "string",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceip",
		ConvertedName:  "SourceIP",
		Description:    `The source IP of a trace route request. Might not be always and endpoint.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"SourceSelector": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceselector",
		ConvertedName:  "SourceSelector",
		Description:    `A filter for selecting the sources of the request.`,
		Exposed:        true,
		Name:           "sourceSelector",
		Stored:         true,
		SubType:        "cloudnetworkqueryfilter",
		Type:           "ref",
	},
	"Type": {
		AllowedChoices: []string{"Summary", "CompressedGraph", "FullGraph"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		DefaultValue:   CloudNetworkQueryTypeSummary,
		Description:    `Indicates the type of results that should be provided by the query.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
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

// CloudNetworkQueryLowerCaseAttributesMap represents the map of attribute for CloudNetworkQuery.
var CloudNetworkQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"destinationip": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationip",
		ConvertedName:  "DestinationIP",
		Description:    `The destination IP of a trace route request. Might not always be an endpoint.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"destinationselector": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationselector",
		ConvertedName:  "DestinationSelector",
		Description:    `A filter for selecting destinations for the query.`,
		Exposed:        true,
		Name:           "destinationSelector",
		Stored:         true,
		SubType:        "cloudnetworkqueryfilter",
		Type:           "ref",
	},
	"effectiveaction": {
		AllowedChoices: []string{"ReachableAndAllowed", "ReachableOnly", "All"},
		BSONFieldName:  "effectiveaction",
		ConvertedName:  "EffectiveAction",
		DefaultValue:   CloudNetworkQueryEffectiveActionReachableOnly,
		Description: `Filters the results based on the effective action. 'ReachableAndAllowed' means
that a destination is both reachable and allowed by security rules.
'UnreachableOrRejected' means that the destination is either not reachable or
rejected by security rules. 'ReachableOnly' means that all destinations that are
reachable will be returned irrespective of their security verdict.
'UnreachableOnly' means that only unreachable destinations will be returned.`,
		Exposed: true,
		Name:    "effectiveAction",
		Stored:  true,
		Type:    "enum",
	},
	"excludeenterpriseips": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludeenterpriseips",
		ConvertedName:  "ExcludeEnterpriseIPs",
		Description: `If set, the evaluation will exclude enterprise IPs from the effective
permissions.`,
		Exposed: true,
		Name:    "excludeEnterpriseIPs",
		Stored:  true,
		Type:    "boolean",
	},
	"excludednetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludednetworks",
		ConvertedName:  "ExcludedNetworks",
		Description: `List of networks that should be a excluded from the calculation if the source or
destination is a network.`,
		Exposed: true,
		Name:    "excludedNetworks",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
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
	"protocolports": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocolports",
		ConvertedName:  "ProtocolPorts",
		Description: `Represents the ports and protocols this policy applies to. Protocol/ports are
defined as tcp/80, udp/22. For protocols that do not have ports, the port
designation
is not allowed.`,
		Exposed: true,
		Name:    "protocolPorts",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"rawrql": {
		AllowedChoices: []string{},
		BSONFieldName:  "rawrql",
		ConvertedName:  "RawRQL",
		Description:    `The RQL string for this query as a reference.`,
		Name:           "rawRQL",
		Stored:         true,
		Type:           "string",
	},
	"sourceip": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceip",
		ConvertedName:  "SourceIP",
		Description:    `The source IP of a trace route request. Might not be always and endpoint.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"sourceselector": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceselector",
		ConvertedName:  "SourceSelector",
		Description:    `A filter for selecting the sources of the request.`,
		Exposed:        true,
		Name:           "sourceSelector",
		Stored:         true,
		SubType:        "cloudnetworkqueryfilter",
		Type:           "ref",
	},
	"type": {
		AllowedChoices: []string{"Summary", "CompressedGraph", "FullGraph"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		DefaultValue:   CloudNetworkQueryTypeSummary,
		Description:    `Indicates the type of results that should be provided by the query.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
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

// SparseCloudNetworkQueriesList represents a list of SparseCloudNetworkQueries
type SparseCloudNetworkQueriesList []*SparseCloudNetworkQuery

// Identity returns the identity of the objects in the list.
func (o SparseCloudNetworkQueriesList) Identity() elemental.Identity {

	return CloudNetworkQueryIdentity
}

// Copy returns a pointer to a copy the SparseCloudNetworkQueriesList.
func (o SparseCloudNetworkQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudNetworkQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudNetworkQueriesList.
func (o SparseCloudNetworkQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudNetworkQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudNetworkQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudNetworkQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudNetworkQueriesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudNetworkQueriesList converted to CloudNetworkQueriesList.
func (o SparseCloudNetworkQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudNetworkQueriesList) Version() int {

	return 1
}

// SparseCloudNetworkQuery represents the sparse version of a cloudnetworkquery.
type SparseCloudNetworkQuery struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The destination IP of a trace route request. Might not always be an endpoint.
	DestinationIP *string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"destinationip,omitempty" mapstructure:"destinationIP,omitempty"`

	// A filter for selecting destinations for the query.
	DestinationSelector *CloudNetworkQueryFilter `json:"destinationSelector,omitempty" msgpack:"destinationSelector,omitempty" bson:"destinationselector,omitempty" mapstructure:"destinationSelector,omitempty"`

	// Filters the results based on the effective action. 'ReachableAndAllowed' means
	// that a destination is both reachable and allowed by security rules.
	// 'UnreachableOrRejected' means that the destination is either not reachable or
	// rejected by security rules. 'ReachableOnly' means that all destinations that are
	// reachable will be returned irrespective of their security verdict.
	// 'UnreachableOnly' means that only unreachable destinations will be returned.
	EffectiveAction *CloudNetworkQueryEffectiveActionValue `json:"effectiveAction,omitempty" msgpack:"effectiveAction,omitempty" bson:"effectiveaction,omitempty" mapstructure:"effectiveAction,omitempty"`

	// If set, the evaluation will exclude enterprise IPs from the effective
	// permissions.
	ExcludeEnterpriseIPs *bool `json:"excludeEnterpriseIPs,omitempty" msgpack:"excludeEnterpriseIPs,omitempty" bson:"excludeenterpriseips,omitempty" mapstructure:"excludeEnterpriseIPs,omitempty"`

	// List of networks that should be a excluded from the calculation if the source or
	// destination is a network.
	ExcludedNetworks *[]string `json:"excludedNetworks,omitempty" msgpack:"excludedNetworks,omitempty" bson:"excludednetworks,omitempty" mapstructure:"excludedNetworks,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Represents the ports and protocols this policy applies to. Protocol/ports are
	// defined as tcp/80, udp/22. For protocols that do not have ports, the port
	// designation
	// is not allowed.
	ProtocolPorts *[]string `json:"protocolPorts,omitempty" msgpack:"protocolPorts,omitempty" bson:"protocolports,omitempty" mapstructure:"protocolPorts,omitempty"`

	// The RQL string for this query as a reference.
	RawRQL *string `json:"-" msgpack:"-" bson:"rawrql,omitempty" mapstructure:"-,omitempty"`

	// The source IP of a trace route request. Might not be always and endpoint.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"sourceip,omitempty" mapstructure:"sourceIP,omitempty"`

	// A filter for selecting the sources of the request.
	SourceSelector *CloudNetworkQueryFilter `json:"sourceSelector,omitempty" msgpack:"sourceSelector,omitempty" bson:"sourceselector,omitempty" mapstructure:"sourceSelector,omitempty"`

	// Indicates the type of results that should be provided by the query.
	Type *CloudNetworkQueryTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudNetworkQuery returns a new  SparseCloudNetworkQuery.
func NewSparseCloudNetworkQuery() *SparseCloudNetworkQuery {
	return &SparseCloudNetworkQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudNetworkQuery) Identity() elemental.Identity {

	return CloudNetworkQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudNetworkQuery) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudNetworkQuery) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudNetworkQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudNetworkQuery{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.DestinationIP != nil {
		s.DestinationIP = o.DestinationIP
	}
	if o.DestinationSelector != nil {
		s.DestinationSelector = o.DestinationSelector
	}
	if o.EffectiveAction != nil {
		s.EffectiveAction = o.EffectiveAction
	}
	if o.ExcludeEnterpriseIPs != nil {
		s.ExcludeEnterpriseIPs = o.ExcludeEnterpriseIPs
	}
	if o.ExcludedNetworks != nil {
		s.ExcludedNetworks = o.ExcludedNetworks
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
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.ProtocolPorts != nil {
		s.ProtocolPorts = o.ProtocolPorts
	}
	if o.RawRQL != nil {
		s.RawRQL = o.RawRQL
	}
	if o.SourceIP != nil {
		s.SourceIP = o.SourceIP
	}
	if o.SourceSelector != nil {
		s.SourceSelector = o.SourceSelector
	}
	if o.Type != nil {
		s.Type = o.Type
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
func (o *SparseCloudNetworkQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudNetworkQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.DestinationIP != nil {
		o.DestinationIP = s.DestinationIP
	}
	if s.DestinationSelector != nil {
		o.DestinationSelector = s.DestinationSelector
	}
	if s.EffectiveAction != nil {
		o.EffectiveAction = s.EffectiveAction
	}
	if s.ExcludeEnterpriseIPs != nil {
		o.ExcludeEnterpriseIPs = s.ExcludeEnterpriseIPs
	}
	if s.ExcludedNetworks != nil {
		o.ExcludedNetworks = s.ExcludedNetworks
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
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.ProtocolPorts != nil {
		o.ProtocolPorts = s.ProtocolPorts
	}
	if s.RawRQL != nil {
		o.RawRQL = s.RawRQL
	}
	if s.SourceIP != nil {
		o.SourceIP = s.SourceIP
	}
	if s.SourceSelector != nil {
		o.SourceSelector = s.SourceSelector
	}
	if s.Type != nil {
		o.Type = s.Type
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
func (o *SparseCloudNetworkQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudNetworkQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudNetworkQuery()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
	}
	if o.DestinationSelector != nil {
		out.DestinationSelector = o.DestinationSelector
	}
	if o.EffectiveAction != nil {
		out.EffectiveAction = *o.EffectiveAction
	}
	if o.ExcludeEnterpriseIPs != nil {
		out.ExcludeEnterpriseIPs = *o.ExcludeEnterpriseIPs
	}
	if o.ExcludedNetworks != nil {
		out.ExcludedNetworks = *o.ExcludedNetworks
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ProtocolPorts != nil {
		out.ProtocolPorts = *o.ProtocolPorts
	}
	if o.RawRQL != nil {
		out.RawRQL = *o.RawRQL
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.SourceSelector != nil {
		out.SourceSelector = o.SourceSelector
	}
	if o.Type != nil {
		out.Type = *o.Type
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

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudNetworkQuery) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudNetworkQuery) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudNetworkQuery) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetDescription returns the Description of the receiver.
func (o *SparseCloudNetworkQuery) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetDescription(description string) {

	o.Description = &description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudNetworkQuery) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudNetworkQuery) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudNetworkQuery) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudNetworkQuery) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudNetworkQuery) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudNetworkQuery) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudNetworkQuery) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudNetworkQuery) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudNetworkQuery) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudNetworkQuery.
func (o *SparseCloudNetworkQuery) DeepCopy() *SparseCloudNetworkQuery {

	if o == nil {
		return nil
	}

	out := &SparseCloudNetworkQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudNetworkQuery.
func (o *SparseCloudNetworkQuery) DeepCopyInto(out *SparseCloudNetworkQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudNetworkQuery: %s", err))
	}

	*out = *target.(*SparseCloudNetworkQuery)
}

type mongoAttributesCloudNetworkQuery struct {
	ID                   bson.ObjectId                         `bson:"_id,omitempty"`
	Annotations          map[string][]string                   `bson:"annotations"`
	AssociatedTags       []string                              `bson:"associatedtags"`
	CreateIdempotencyKey string                                `bson:"createidempotencykey"`
	Description          string                                `bson:"description"`
	DestinationIP        string                                `bson:"destinationip"`
	DestinationSelector  *CloudNetworkQueryFilter              `bson:"destinationselector"`
	EffectiveAction      CloudNetworkQueryEffectiveActionValue `bson:"effectiveaction"`
	ExcludeEnterpriseIPs bool                                  `bson:"excludeenterpriseips"`
	ExcludedNetworks     []string                              `bson:"excludednetworks"`
	MigrationsLog        map[string]string                     `bson:"migrationslog,omitempty"`
	Name                 string                                `bson:"name"`
	Namespace            string                                `bson:"namespace"`
	NormalizedTags       []string                              `bson:"normalizedtags"`
	Protected            bool                                  `bson:"protected"`
	ProtocolPorts        []string                              `bson:"protocolports"`
	RawRQL               string                                `bson:"rawrql"`
	SourceIP             string                                `bson:"sourceip"`
	SourceSelector       *CloudNetworkQueryFilter              `bson:"sourceselector"`
	Type                 CloudNetworkQueryTypeValue            `bson:"type"`
	UpdateIdempotencyKey string                                `bson:"updateidempotencykey"`
	ZHash                int                                   `bson:"zhash"`
	Zone                 int                                   `bson:"zone"`
}
type mongoAttributesSparseCloudNetworkQuery struct {
	ID                   bson.ObjectId                          `bson:"_id,omitempty"`
	Annotations          *map[string][]string                   `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                              `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string                                `bson:"createidempotencykey,omitempty"`
	Description          *string                                `bson:"description,omitempty"`
	DestinationIP        *string                                `bson:"destinationip,omitempty"`
	DestinationSelector  *CloudNetworkQueryFilter               `bson:"destinationselector,omitempty"`
	EffectiveAction      *CloudNetworkQueryEffectiveActionValue `bson:"effectiveaction,omitempty"`
	ExcludeEnterpriseIPs *bool                                  `bson:"excludeenterpriseips,omitempty"`
	ExcludedNetworks     *[]string                              `bson:"excludednetworks,omitempty"`
	MigrationsLog        *map[string]string                     `bson:"migrationslog,omitempty"`
	Name                 *string                                `bson:"name,omitempty"`
	Namespace            *string                                `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                              `bson:"normalizedtags,omitempty"`
	Protected            *bool                                  `bson:"protected,omitempty"`
	ProtocolPorts        *[]string                              `bson:"protocolports,omitempty"`
	RawRQL               *string                                `bson:"rawrql,omitempty"`
	SourceIP             *string                                `bson:"sourceip,omitempty"`
	SourceSelector       *CloudNetworkQueryFilter               `bson:"sourceselector,omitempty"`
	Type                 *CloudNetworkQueryTypeValue            `bson:"type,omitempty"`
	UpdateIdempotencyKey *string                                `bson:"updateidempotencykey,omitempty"`
	ZHash                *int                                   `bson:"zhash,omitempty"`
	Zone                 *int                                   `bson:"zone,omitempty"`
}
