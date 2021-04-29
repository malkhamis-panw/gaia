package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LoadBalancerTypeValue represents the possible values for attribute "type".
type LoadBalancerTypeValue string

const (
	// LoadBalancerTypeHTTP represents the value HTTP.
	LoadBalancerTypeHTTP LoadBalancerTypeValue = "HTTP"

	// LoadBalancerTypeTCP represents the value TCP.
	LoadBalancerTypeTCP LoadBalancerTypeValue = "TCP"
)

// LoadBalancerIdentity represents the Identity of the object.
var LoadBalancerIdentity = elemental.Identity{
	Name:     "loadbalancer",
	Category: "loadbalancers",
	Package:  "squall",
	Private:  false,
}

// LoadBalancersList represents a list of LoadBalancers
type LoadBalancersList []*LoadBalancer

// Identity returns the identity of the objects in the list.
func (o LoadBalancersList) Identity() elemental.Identity {

	return LoadBalancerIdentity
}

// Copy returns a pointer to a copy the LoadBalancersList.
func (o LoadBalancersList) Copy() elemental.Identifiables {

	copy := append(LoadBalancersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LoadBalancersList.
func (o LoadBalancersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LoadBalancersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LoadBalancer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LoadBalancersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LoadBalancersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the LoadBalancersList converted to SparseLoadBalancersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LoadBalancersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLoadBalancersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLoadBalancer)
	}

	return out
}

// Version returns the version of the content.
func (o LoadBalancersList) Version() int {

	return 1
}

// LoadBalancer represents the model of a loadbalancer
type LoadBalancer struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The list of IP addresses where the service can be accessed. This is an optional
	// attribute and is only required if no host names are provided. The system will
	// automatically resolve IP addresses from host names otherwise.
	IPs []string `json:"IPs" msgpack:"IPs" bson:"ips" mapstructure:"IPs,omitempty"`

	// A tag or tag expression that identifies the TLS certificates to be used by
	// enforcers when exposing the service ran by the processing units.
	TLSCertificatesSelector [][]string `json:"TLSCertificatesSelector" msgpack:"TLSCertificatesSelector" bson:"tlscertificatesselector" mapstructure:"TLSCertificatesSelector,omitempty"`

	// This is a set of all selector tags for matching in the database.
	AllProcessingUnitTags []string `json:"-" msgpack:"-" bson:"allprocessingunittags" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the database.
	AllTLSCertificateTags []string `json:"-" msgpack:"-" bson:"alltlscertificatetags" mapstructure:"-,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Defines if the object is archived.
	Archived bool `json:"-" msgpack:"-" bson:"archived" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// The port that the service can be accessed on. Note that this is different from
	// the `port` attribute that describes the port that the service is actually
	// listening on. For example if a load balancer is used, the `exposedPort` is
	// the port that the load balancer is listening for the service, whereas the
	// port that the implementation is listening on can be different.
	ExposedPort int `json:"exposedPort" msgpack:"exposedPort" bson:"exposedport" mapstructure:"exposedPort,omitempty"`

	// Indicates that the exposed service is TLS. This means that the enforcer has to
	// initiate a TLS session in order to forward traffic to the service.
	ExposedServiceIsTLS bool `json:"exposedServiceIsTLS" msgpack:"exposedServiceIsTLS" bson:"exposedserviceistls" mapstructure:"exposedServiceIsTLS,omitempty"`

	// The host names that the service can be accessed on.
	Hosts []string `json:"hosts" msgpack:"hosts" bson:"hosts" mapstructure:"hosts,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// The port that the implementation of the service is listening to. It can be
	// different than `exposedPort`. This is needed for port mapping use cases
	// where there are private and public ports.
	Port int `json:"port" msgpack:"port" bson:"port" mapstructure:"port,omitempty"`

	// Tag expression that identifies the processing unit that implements this
	// particular service.
	ProcessingUnitSelector [][]string `json:"processingUnitSelector" msgpack:"processingUnitSelector" bson:"processingunitselector" mapstructure:"processingUnitSelector,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Enable trust in proxy protocols header.
	ProxyProtocolEnabled bool `json:"proxyProtocolEnabled" msgpack:"proxyProtocolEnabled" bson:"proxyprotocolenabled" mapstructure:"proxyProtocolEnabled,omitempty"`

	// Only allow proxy protocols from the given subnets .
	ProxyProtocolSubnets []string `json:"proxyProtocolSubnets" msgpack:"proxyProtocolSubnets" bson:"proxyprotocolsubnets" mapstructure:"proxyProtocolSubnets,omitempty"`

	// A new virtual port that the service can be accessed on using HTTPS. Since the
	// enforcer transparently inserts TLS in the application path, you might want
	// to declare a new port where the enforcer listens for TLS. However, the
	// application does not need to be modified and the enforcer will map the
	// traffic to the correct application port. This is useful when
	// an application is being accessed from a public network.
	PublicApplicationPort int `json:"publicApplicationPort" msgpack:"publicApplicationPort" bson:"publicapplicationport" mapstructure:"publicApplicationPort,omitempty"`

	// PEM-encoded certificate authorities to trust when additional hops are needed. It
	// must be set if the service must reach a service marked as `external` or must go
	// through an additional TLS termination point like a layer 7 load balancer.
	TrustedCertificateAuthorities string `json:"trustedCertificateAuthorities" msgpack:"trustedCertificateAuthorities" bson:"trustedcertificateauthorities" mapstructure:"trustedCertificateAuthorities,omitempty"`

	// Type of the load balancer.
	Type LoadBalancerTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLoadBalancer returns a new *LoadBalancer
func NewLoadBalancer() *LoadBalancer {

	return &LoadBalancer{
		ModelVersion:            1,
		AssociatedTags:          []string{},
		AllProcessingUnitTags:   []string{},
		AllTLSCertificateTags:   []string{},
		Annotations:             map[string][]string{},
		Hosts:                   []string{},
		ExposedServiceIsTLS:     false,
		MigrationsLog:           map[string]string{},
		TLSCertificatesSelector: [][]string{},
		Metadata:                []string{},
		ProxyProtocolSubnets:    []string{},
		NormalizedTags:          []string{},
		ProcessingUnitSelector:  [][]string{},
		IPs:                     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *LoadBalancer) Identity() elemental.Identity {

	return LoadBalancerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LoadBalancer) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LoadBalancer) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LoadBalancer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLoadBalancer{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.IPs = o.IPs
	s.TLSCertificatesSelector = o.TLSCertificatesSelector
	s.AllProcessingUnitTags = o.AllProcessingUnitTags
	s.AllTLSCertificateTags = o.AllTLSCertificateTags
	s.Annotations = o.Annotations
	s.Archived = o.Archived
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.ExposedPort = o.ExposedPort
	s.ExposedServiceIsTLS = o.ExposedServiceIsTLS
	s.Hosts = o.Hosts
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Port = o.Port
	s.ProcessingUnitSelector = o.ProcessingUnitSelector
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.ProxyProtocolEnabled = o.ProxyProtocolEnabled
	s.ProxyProtocolSubnets = o.ProxyProtocolSubnets
	s.PublicApplicationPort = o.PublicApplicationPort
	s.TrustedCertificateAuthorities = o.TrustedCertificateAuthorities
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LoadBalancer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLoadBalancer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.IPs = s.IPs
	o.TLSCertificatesSelector = s.TLSCertificatesSelector
	o.AllProcessingUnitTags = s.AllProcessingUnitTags
	o.AllTLSCertificateTags = s.AllTLSCertificateTags
	o.Annotations = s.Annotations
	o.Archived = s.Archived
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.ExposedPort = s.ExposedPort
	o.ExposedServiceIsTLS = s.ExposedServiceIsTLS
	o.Hosts = s.Hosts
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Port = s.Port
	o.ProcessingUnitSelector = s.ProcessingUnitSelector
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.ProxyProtocolEnabled = s.ProxyProtocolEnabled
	o.ProxyProtocolSubnets = s.ProxyProtocolSubnets
	o.PublicApplicationPort = s.PublicApplicationPort
	o.TrustedCertificateAuthorities = s.TrustedCertificateAuthorities
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *LoadBalancer) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *LoadBalancer) BleveType() string {

	return "loadbalancer"
}

// DefaultOrder returns the list of default ordering fields.
func (o *LoadBalancer) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *LoadBalancer) Doc() string {

	return `Defines a generic external LoadBalancer that sits between 2 enforcers.`
}

func (o *LoadBalancer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *LoadBalancer) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *LoadBalancer) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *LoadBalancer) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the property Archived of the receiver using the given value.
func (o *LoadBalancer) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *LoadBalancer) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *LoadBalancer) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *LoadBalancer) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *LoadBalancer) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *LoadBalancer) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *LoadBalancer) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *LoadBalancer) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *LoadBalancer) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *LoadBalancer) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *LoadBalancer) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *LoadBalancer) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *LoadBalancer) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *LoadBalancer) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *LoadBalancer) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *LoadBalancer) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *LoadBalancer) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *LoadBalancer) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *LoadBalancer) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *LoadBalancer) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *LoadBalancer) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *LoadBalancer) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *LoadBalancer) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *LoadBalancer) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *LoadBalancer) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *LoadBalancer) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *LoadBalancer) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *LoadBalancer) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *LoadBalancer) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *LoadBalancer) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *LoadBalancer) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *LoadBalancer) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *LoadBalancer) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LoadBalancer) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLoadBalancer{
			ID:                            &o.ID,
			IPs:                           &o.IPs,
			TLSCertificatesSelector:       &o.TLSCertificatesSelector,
			AllProcessingUnitTags:         &o.AllProcessingUnitTags,
			AllTLSCertificateTags:         &o.AllTLSCertificateTags,
			Annotations:                   &o.Annotations,
			Archived:                      &o.Archived,
			AssociatedTags:                &o.AssociatedTags,
			CreateIdempotencyKey:          &o.CreateIdempotencyKey,
			CreateTime:                    &o.CreateTime,
			Description:                   &o.Description,
			Disabled:                      &o.Disabled,
			ExposedPort:                   &o.ExposedPort,
			ExposedServiceIsTLS:           &o.ExposedServiceIsTLS,
			Hosts:                         &o.Hosts,
			Metadata:                      &o.Metadata,
			MigrationsLog:                 &o.MigrationsLog,
			Name:                          &o.Name,
			Namespace:                     &o.Namespace,
			NormalizedTags:                &o.NormalizedTags,
			Port:                          &o.Port,
			ProcessingUnitSelector:        &o.ProcessingUnitSelector,
			Propagate:                     &o.Propagate,
			Protected:                     &o.Protected,
			ProxyProtocolEnabled:          &o.ProxyProtocolEnabled,
			ProxyProtocolSubnets:          &o.ProxyProtocolSubnets,
			PublicApplicationPort:         &o.PublicApplicationPort,
			TrustedCertificateAuthorities: &o.TrustedCertificateAuthorities,
			Type:                          &o.Type,
			UpdateIdempotencyKey:          &o.UpdateIdempotencyKey,
			UpdateTime:                    &o.UpdateTime,
			ZHash:                         &o.ZHash,
			Zone:                          &o.Zone,
		}
	}

	sp := &SparseLoadBalancer{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "IPs":
			sp.IPs = &(o.IPs)
		case "TLSCertificatesSelector":
			sp.TLSCertificatesSelector = &(o.TLSCertificatesSelector)
		case "allProcessingUnitTags":
			sp.AllProcessingUnitTags = &(o.AllProcessingUnitTags)
		case "allTLSCertificateTags":
			sp.AllTLSCertificateTags = &(o.AllTLSCertificateTags)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "archived":
			sp.Archived = &(o.Archived)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "exposedPort":
			sp.ExposedPort = &(o.ExposedPort)
		case "exposedServiceIsTLS":
			sp.ExposedServiceIsTLS = &(o.ExposedServiceIsTLS)
		case "hosts":
			sp.Hosts = &(o.Hosts)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "port":
			sp.Port = &(o.Port)
		case "processingUnitSelector":
			sp.ProcessingUnitSelector = &(o.ProcessingUnitSelector)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "proxyProtocolEnabled":
			sp.ProxyProtocolEnabled = &(o.ProxyProtocolEnabled)
		case "proxyProtocolSubnets":
			sp.ProxyProtocolSubnets = &(o.ProxyProtocolSubnets)
		case "publicApplicationPort":
			sp.PublicApplicationPort = &(o.PublicApplicationPort)
		case "trustedCertificateAuthorities":
			sp.TrustedCertificateAuthorities = &(o.TrustedCertificateAuthorities)
		case "type":
			sp.Type = &(o.Type)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLoadBalancer to the object.
func (o *LoadBalancer) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLoadBalancer)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IPs != nil {
		o.IPs = *so.IPs
	}
	if so.TLSCertificatesSelector != nil {
		o.TLSCertificatesSelector = *so.TLSCertificatesSelector
	}
	if so.AllProcessingUnitTags != nil {
		o.AllProcessingUnitTags = *so.AllProcessingUnitTags
	}
	if so.AllTLSCertificateTags != nil {
		o.AllTLSCertificateTags = *so.AllTLSCertificateTags
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.Archived != nil {
		o.Archived = *so.Archived
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.ExposedPort != nil {
		o.ExposedPort = *so.ExposedPort
	}
	if so.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = *so.ExposedServiceIsTLS
	}
	if so.Hosts != nil {
		o.Hosts = *so.Hosts
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
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
	if so.Port != nil {
		o.Port = *so.Port
	}
	if so.ProcessingUnitSelector != nil {
		o.ProcessingUnitSelector = *so.ProcessingUnitSelector
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ProxyProtocolEnabled != nil {
		o.ProxyProtocolEnabled = *so.ProxyProtocolEnabled
	}
	if so.ProxyProtocolSubnets != nil {
		o.ProxyProtocolSubnets = *so.ProxyProtocolSubnets
	}
	if so.PublicApplicationPort != nil {
		o.PublicApplicationPort = *so.PublicApplicationPort
	}
	if so.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = *so.TrustedCertificateAuthorities
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the LoadBalancer.
func (o *LoadBalancer) DeepCopy() *LoadBalancer {

	if o == nil {
		return nil
	}

	out := &LoadBalancer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LoadBalancer.
func (o *LoadBalancer) DeepCopyInto(out *LoadBalancer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LoadBalancer: %s", err))
	}

	*out = *target.(*LoadBalancer)
}

// Validate valides the current information stored into the structure.
func (o *LoadBalancer) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateOptionalIPAddressList("IPs", o.IPs); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("TLSCertificatesSelector", o.TLSCertificatesSelector); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("exposedPort", o.ExposedPort); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("exposedPort", o.ExposedPort, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("port", o.Port, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("processingUnitSelector", o.ProcessingUnitSelector); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOptionalCIDRList("proxyProtocolSubnets", o.ProxyProtocolSubnets); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("publicApplicationPort", o.PublicApplicationPort, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePEM("trustedCertificateAuthorities", o.TrustedCertificateAuthorities); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"HTTP", "TCP"}, false); err != nil {
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
func (*LoadBalancer) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LoadBalancerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LoadBalancerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LoadBalancer) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LoadBalancerAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LoadBalancer) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "IPs":
		return o.IPs
	case "TLSCertificatesSelector":
		return o.TLSCertificatesSelector
	case "allProcessingUnitTags":
		return o.AllProcessingUnitTags
	case "allTLSCertificateTags":
		return o.AllTLSCertificateTags
	case "annotations":
		return o.Annotations
	case "archived":
		return o.Archived
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "exposedPort":
		return o.ExposedPort
	case "exposedServiceIsTLS":
		return o.ExposedServiceIsTLS
	case "hosts":
		return o.Hosts
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "port":
		return o.Port
	case "processingUnitSelector":
		return o.ProcessingUnitSelector
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "proxyProtocolEnabled":
		return o.ProxyProtocolEnabled
	case "proxyProtocolSubnets":
		return o.ProxyProtocolSubnets
	case "publicApplicationPort":
		return o.PublicApplicationPort
	case "trustedCertificateAuthorities":
		return o.TrustedCertificateAuthorities
	case "type":
		return o.Type
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// LoadBalancerAttributesMap represents the map of attribute for LoadBalancer.
var LoadBalancerAttributesMap = map[string]elemental.AttributeSpecification{
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
	"IPs": {
		AllowedChoices: []string{},
		BSONFieldName:  "ips",
		ConvertedName:  "IPs",
		Description: `The list of IP addresses where the service can be accessed. This is an optional
attribute and is only required if no host names are provided. The system will
automatically resolve IP addresses from host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"TLSCertificatesSelector": {
		AllowedChoices: []string{},
		BSONFieldName:  "tlscertificatesselector",
		ConvertedName:  "TLSCertificatesSelector",
		Description: `A tag or tag expression that identifies the TLS certificates to be used by
enforcers when exposing the service ran by the processing units.`,
		Exposed: true,
		Name:    "TLSCertificatesSelector",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"AllProcessingUnitTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allprocessingunittags",
		ConvertedName:  "AllProcessingUnitTags",
		Description:    `This is a set of all selector tags for matching in the database.`,
		Name:           "allProcessingUnitTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AllTLSCertificateTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "alltlscertificatetags",
		ConvertedName:  "AllTLSCertificateTags",
		Description:    `This is a set of all selector tags for matching in the database.`,
		Name:           "allTLSCertificateTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Archived": {
		AllowedChoices: []string{},
		BSONFieldName:  "archived",
		ConvertedName:  "Archived",
		Description:    `Defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
	"Disabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "disabled",
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"ExposedPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "exposedport",
		ConvertedName:  "ExposedPort",
		Description: `The port that the service can be accessed on. Note that this is different from
the ` + "`" + `port` + "`" + ` attribute that describes the port that the service is actually
listening on. For example if a load balancer is used, the ` + "`" + `exposedPort` + "`" + ` is
the port that the load balancer is listening for the service, whereas the
port that the implementation is listening on can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"ExposedServiceIsTLS": {
		AllowedChoices: []string{},
		BSONFieldName:  "exposedserviceistls",
		ConvertedName:  "ExposedServiceIsTLS",
		Description: `Indicates that the exposed service is TLS. This means that the enforcer has to
initiate a TLS session in order to forward traffic to the service.`,
		Exposed:    true,
		Filterable: true,
		Name:       "exposedServiceIsTLS",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"Hosts": {
		AllowedChoices: []string{},
		BSONFieldName:  "hosts",
		ConvertedName:  "Hosts",
		Description:    `The host names that the service can be accessed on.`,
		Exposed:        true,
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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
	"Port": {
		AllowedChoices: []string{},
		BSONFieldName:  "port",
		ConvertedName:  "Port",
		Description: `The port that the implementation of the service is listening to. It can be
different than ` + "`" + `exposedPort` + "`" + `. This is needed for port mapping use cases
where there are private and public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "port",
		Stored:   true,
		Type:     "integer",
	},
	"ProcessingUnitSelector": {
		AllowedChoices: []string{},
		BSONFieldName:  "processingunitselector",
		ConvertedName:  "ProcessingUnitSelector",
		Description: `Tag expression that identifies the processing unit that implements this
particular service.`,
		Exposed: true,
		Name:    "processingUnitSelector",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"Propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"ProxyProtocolEnabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "proxyprotocolenabled",
		ConvertedName:  "ProxyProtocolEnabled",
		Description:    `Enable trust in proxy protocols header.`,
		Exposed:        true,
		Name:           "proxyProtocolEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"ProxyProtocolSubnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "proxyprotocolsubnets",
		ConvertedName:  "ProxyProtocolSubnets",
		Description:    `Only allow proxy protocols from the given subnets .`,
		Exposed:        true,
		Name:           "proxyProtocolSubnets",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"PublicApplicationPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicapplicationport",
		ConvertedName:  "PublicApplicationPort",
		Description: `A new virtual port that the service can be accessed on using HTTPS. Since the
enforcer transparently inserts TLS in the application path, you might want
to declare a new port where the enforcer listens for TLS. However, the
application does not need to be modified and the enforcer will map the
traffic to the correct application port. This is useful when
an application is being accessed from a public network.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "publicApplicationPort",
		Stored:   true,
		Type:     "integer",
	},
	"TrustedCertificateAuthorities": {
		AllowedChoices: []string{},
		BSONFieldName:  "trustedcertificateauthorities",
		ConvertedName:  "TrustedCertificateAuthorities",
		Description: `PEM-encoded certificate authorities to trust when additional hops are needed. It
must be set if the service must reach a service marked as ` + "`" + `external` + "`" + ` or must go
through an additional TLS termination point like a layer 7 load balancer.`,
		Exposed: true,
		Name:    "trustedCertificateAuthorities",
		Stored:  true,
		Type:    "string",
	},
	"Type": {
		AllowedChoices: []string{"HTTP", "TCP"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the load balancer.`,
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
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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

// LoadBalancerLowerCaseAttributesMap represents the map of attribute for LoadBalancer.
var LoadBalancerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ips": {
		AllowedChoices: []string{},
		BSONFieldName:  "ips",
		ConvertedName:  "IPs",
		Description: `The list of IP addresses where the service can be accessed. This is an optional
attribute and is only required if no host names are provided. The system will
automatically resolve IP addresses from host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"tlscertificatesselector": {
		AllowedChoices: []string{},
		BSONFieldName:  "tlscertificatesselector",
		ConvertedName:  "TLSCertificatesSelector",
		Description: `A tag or tag expression that identifies the TLS certificates to be used by
enforcers when exposing the service ran by the processing units.`,
		Exposed: true,
		Name:    "TLSCertificatesSelector",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"allprocessingunittags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allprocessingunittags",
		ConvertedName:  "AllProcessingUnitTags",
		Description:    `This is a set of all selector tags for matching in the database.`,
		Name:           "allProcessingUnitTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"alltlscertificatetags": {
		AllowedChoices: []string{},
		BSONFieldName:  "alltlscertificatetags",
		ConvertedName:  "AllTLSCertificateTags",
		Description:    `This is a set of all selector tags for matching in the database.`,
		Name:           "allTLSCertificateTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"archived": {
		AllowedChoices: []string{},
		BSONFieldName:  "archived",
		ConvertedName:  "Archived",
		Description:    `Defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
	"disabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "disabled",
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"exposedport": {
		AllowedChoices: []string{},
		BSONFieldName:  "exposedport",
		ConvertedName:  "ExposedPort",
		Description: `The port that the service can be accessed on. Note that this is different from
the ` + "`" + `port` + "`" + ` attribute that describes the port that the service is actually
listening on. For example if a load balancer is used, the ` + "`" + `exposedPort` + "`" + ` is
the port that the load balancer is listening for the service, whereas the
port that the implementation is listening on can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"exposedserviceistls": {
		AllowedChoices: []string{},
		BSONFieldName:  "exposedserviceistls",
		ConvertedName:  "ExposedServiceIsTLS",
		Description: `Indicates that the exposed service is TLS. This means that the enforcer has to
initiate a TLS session in order to forward traffic to the service.`,
		Exposed:    true,
		Filterable: true,
		Name:       "exposedServiceIsTLS",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"hosts": {
		AllowedChoices: []string{},
		BSONFieldName:  "hosts",
		ConvertedName:  "Hosts",
		Description:    `The host names that the service can be accessed on.`,
		Exposed:        true,
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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
	"port": {
		AllowedChoices: []string{},
		BSONFieldName:  "port",
		ConvertedName:  "Port",
		Description: `The port that the implementation of the service is listening to. It can be
different than ` + "`" + `exposedPort` + "`" + `. This is needed for port mapping use cases
where there are private and public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "port",
		Stored:   true,
		Type:     "integer",
	},
	"processingunitselector": {
		AllowedChoices: []string{},
		BSONFieldName:  "processingunitselector",
		ConvertedName:  "ProcessingUnitSelector",
		Description: `Tag expression that identifies the processing unit that implements this
particular service.`,
		Exposed: true,
		Name:    "processingUnitSelector",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"proxyprotocolenabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "proxyprotocolenabled",
		ConvertedName:  "ProxyProtocolEnabled",
		Description:    `Enable trust in proxy protocols header.`,
		Exposed:        true,
		Name:           "proxyProtocolEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"proxyprotocolsubnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "proxyprotocolsubnets",
		ConvertedName:  "ProxyProtocolSubnets",
		Description:    `Only allow proxy protocols from the given subnets .`,
		Exposed:        true,
		Name:           "proxyProtocolSubnets",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"publicapplicationport": {
		AllowedChoices: []string{},
		BSONFieldName:  "publicapplicationport",
		ConvertedName:  "PublicApplicationPort",
		Description: `A new virtual port that the service can be accessed on using HTTPS. Since the
enforcer transparently inserts TLS in the application path, you might want
to declare a new port where the enforcer listens for TLS. However, the
application does not need to be modified and the enforcer will map the
traffic to the correct application port. This is useful when
an application is being accessed from a public network.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "publicApplicationPort",
		Stored:   true,
		Type:     "integer",
	},
	"trustedcertificateauthorities": {
		AllowedChoices: []string{},
		BSONFieldName:  "trustedcertificateauthorities",
		ConvertedName:  "TrustedCertificateAuthorities",
		Description: `PEM-encoded certificate authorities to trust when additional hops are needed. It
must be set if the service must reach a service marked as ` + "`" + `external` + "`" + ` or must go
through an additional TLS termination point like a layer 7 load balancer.`,
		Exposed: true,
		Name:    "trustedCertificateAuthorities",
		Stored:  true,
		Type:    "string",
	},
	"type": {
		AllowedChoices: []string{"HTTP", "TCP"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the load balancer.`,
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
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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

// SparseLoadBalancersList represents a list of SparseLoadBalancers
type SparseLoadBalancersList []*SparseLoadBalancer

// Identity returns the identity of the objects in the list.
func (o SparseLoadBalancersList) Identity() elemental.Identity {

	return LoadBalancerIdentity
}

// Copy returns a pointer to a copy the SparseLoadBalancersList.
func (o SparseLoadBalancersList) Copy() elemental.Identifiables {

	copy := append(SparseLoadBalancersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLoadBalancersList.
func (o SparseLoadBalancersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLoadBalancersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLoadBalancer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLoadBalancersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLoadBalancersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseLoadBalancersList converted to LoadBalancersList.
func (o SparseLoadBalancersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLoadBalancersList) Version() int {

	return 1
}

// SparseLoadBalancer represents the sparse version of a loadbalancer.
type SparseLoadBalancer struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The list of IP addresses where the service can be accessed. This is an optional
	// attribute and is only required if no host names are provided. The system will
	// automatically resolve IP addresses from host names otherwise.
	IPs *[]string `json:"IPs,omitempty" msgpack:"IPs,omitempty" bson:"ips,omitempty" mapstructure:"IPs,omitempty"`

	// A tag or tag expression that identifies the TLS certificates to be used by
	// enforcers when exposing the service ran by the processing units.
	TLSCertificatesSelector *[][]string `json:"TLSCertificatesSelector,omitempty" msgpack:"TLSCertificatesSelector,omitempty" bson:"tlscertificatesselector,omitempty" mapstructure:"TLSCertificatesSelector,omitempty"`

	// This is a set of all selector tags for matching in the database.
	AllProcessingUnitTags *[]string `json:"-" msgpack:"-" bson:"allprocessingunittags,omitempty" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the database.
	AllTLSCertificateTags *[]string `json:"-" msgpack:"-" bson:"alltlscertificatetags,omitempty" mapstructure:"-,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Defines if the object is archived.
	Archived *bool `json:"-" msgpack:"-" bson:"archived,omitempty" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// The port that the service can be accessed on. Note that this is different from
	// the `port` attribute that describes the port that the service is actually
	// listening on. For example if a load balancer is used, the `exposedPort` is
	// the port that the load balancer is listening for the service, whereas the
	// port that the implementation is listening on can be different.
	ExposedPort *int `json:"exposedPort,omitempty" msgpack:"exposedPort,omitempty" bson:"exposedport,omitempty" mapstructure:"exposedPort,omitempty"`

	// Indicates that the exposed service is TLS. This means that the enforcer has to
	// initiate a TLS session in order to forward traffic to the service.
	ExposedServiceIsTLS *bool `json:"exposedServiceIsTLS,omitempty" msgpack:"exposedServiceIsTLS,omitempty" bson:"exposedserviceistls,omitempty" mapstructure:"exposedServiceIsTLS,omitempty"`

	// The host names that the service can be accessed on.
	Hosts *[]string `json:"hosts,omitempty" msgpack:"hosts,omitempty" bson:"hosts,omitempty" mapstructure:"hosts,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// The port that the implementation of the service is listening to. It can be
	// different than `exposedPort`. This is needed for port mapping use cases
	// where there are private and public ports.
	Port *int `json:"port,omitempty" msgpack:"port,omitempty" bson:"port,omitempty" mapstructure:"port,omitempty"`

	// Tag expression that identifies the processing unit that implements this
	// particular service.
	ProcessingUnitSelector *[][]string `json:"processingUnitSelector,omitempty" msgpack:"processingUnitSelector,omitempty" bson:"processingunitselector,omitempty" mapstructure:"processingUnitSelector,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Enable trust in proxy protocols header.
	ProxyProtocolEnabled *bool `json:"proxyProtocolEnabled,omitempty" msgpack:"proxyProtocolEnabled,omitempty" bson:"proxyprotocolenabled,omitempty" mapstructure:"proxyProtocolEnabled,omitempty"`

	// Only allow proxy protocols from the given subnets .
	ProxyProtocolSubnets *[]string `json:"proxyProtocolSubnets,omitempty" msgpack:"proxyProtocolSubnets,omitempty" bson:"proxyprotocolsubnets,omitempty" mapstructure:"proxyProtocolSubnets,omitempty"`

	// A new virtual port that the service can be accessed on using HTTPS. Since the
	// enforcer transparently inserts TLS in the application path, you might want
	// to declare a new port where the enforcer listens for TLS. However, the
	// application does not need to be modified and the enforcer will map the
	// traffic to the correct application port. This is useful when
	// an application is being accessed from a public network.
	PublicApplicationPort *int `json:"publicApplicationPort,omitempty" msgpack:"publicApplicationPort,omitempty" bson:"publicapplicationport,omitempty" mapstructure:"publicApplicationPort,omitempty"`

	// PEM-encoded certificate authorities to trust when additional hops are needed. It
	// must be set if the service must reach a service marked as `external` or must go
	// through an additional TLS termination point like a layer 7 load balancer.
	TrustedCertificateAuthorities *string `json:"trustedCertificateAuthorities,omitempty" msgpack:"trustedCertificateAuthorities,omitempty" bson:"trustedcertificateauthorities,omitempty" mapstructure:"trustedCertificateAuthorities,omitempty"`

	// Type of the load balancer.
	Type *LoadBalancerTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLoadBalancer returns a new  SparseLoadBalancer.
func NewSparseLoadBalancer() *SparseLoadBalancer {
	return &SparseLoadBalancer{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLoadBalancer) Identity() elemental.Identity {

	return LoadBalancerIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLoadBalancer) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLoadBalancer) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLoadBalancer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLoadBalancer{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.IPs != nil {
		s.IPs = o.IPs
	}
	if o.TLSCertificatesSelector != nil {
		s.TLSCertificatesSelector = o.TLSCertificatesSelector
	}
	if o.AllProcessingUnitTags != nil {
		s.AllProcessingUnitTags = o.AllProcessingUnitTags
	}
	if o.AllTLSCertificateTags != nil {
		s.AllTLSCertificateTags = o.AllTLSCertificateTags
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.Archived != nil {
		s.Archived = o.Archived
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.ExposedPort != nil {
		s.ExposedPort = o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		s.ExposedServiceIsTLS = o.ExposedServiceIsTLS
	}
	if o.Hosts != nil {
		s.Hosts = o.Hosts
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
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
	if o.Port != nil {
		s.Port = o.Port
	}
	if o.ProcessingUnitSelector != nil {
		s.ProcessingUnitSelector = o.ProcessingUnitSelector
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.ProxyProtocolEnabled != nil {
		s.ProxyProtocolEnabled = o.ProxyProtocolEnabled
	}
	if o.ProxyProtocolSubnets != nil {
		s.ProxyProtocolSubnets = o.ProxyProtocolSubnets
	}
	if o.PublicApplicationPort != nil {
		s.PublicApplicationPort = o.PublicApplicationPort
	}
	if o.TrustedCertificateAuthorities != nil {
		s.TrustedCertificateAuthorities = o.TrustedCertificateAuthorities
	}
	if o.Type != nil {
		s.Type = o.Type
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
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
func (o *SparseLoadBalancer) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseLoadBalancer{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.IPs != nil {
		o.IPs = s.IPs
	}
	if s.TLSCertificatesSelector != nil {
		o.TLSCertificatesSelector = s.TLSCertificatesSelector
	}
	if s.AllProcessingUnitTags != nil {
		o.AllProcessingUnitTags = s.AllProcessingUnitTags
	}
	if s.AllTLSCertificateTags != nil {
		o.AllTLSCertificateTags = s.AllTLSCertificateTags
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.Archived != nil {
		o.Archived = s.Archived
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.ExposedPort != nil {
		o.ExposedPort = s.ExposedPort
	}
	if s.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = s.ExposedServiceIsTLS
	}
	if s.Hosts != nil {
		o.Hosts = s.Hosts
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
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
	if s.Port != nil {
		o.Port = s.Port
	}
	if s.ProcessingUnitSelector != nil {
		o.ProcessingUnitSelector = s.ProcessingUnitSelector
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.ProxyProtocolEnabled != nil {
		o.ProxyProtocolEnabled = s.ProxyProtocolEnabled
	}
	if s.ProxyProtocolSubnets != nil {
		o.ProxyProtocolSubnets = s.ProxyProtocolSubnets
	}
	if s.PublicApplicationPort != nil {
		o.PublicApplicationPort = s.PublicApplicationPort
	}
	if s.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = s.TrustedCertificateAuthorities
	}
	if s.Type != nil {
		o.Type = s.Type
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
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
func (o *SparseLoadBalancer) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLoadBalancer) ToPlain() elemental.PlainIdentifiable {

	out := NewLoadBalancer()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IPs != nil {
		out.IPs = *o.IPs
	}
	if o.TLSCertificatesSelector != nil {
		out.TLSCertificatesSelector = *o.TLSCertificatesSelector
	}
	if o.AllProcessingUnitTags != nil {
		out.AllProcessingUnitTags = *o.AllProcessingUnitTags
	}
	if o.AllTLSCertificateTags != nil {
		out.AllTLSCertificateTags = *o.AllTLSCertificateTags
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.Archived != nil {
		out.Archived = *o.Archived
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.ExposedPort != nil {
		out.ExposedPort = *o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		out.ExposedServiceIsTLS = *o.ExposedServiceIsTLS
	}
	if o.Hosts != nil {
		out.Hosts = *o.Hosts
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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
	if o.Port != nil {
		out.Port = *o.Port
	}
	if o.ProcessingUnitSelector != nil {
		out.ProcessingUnitSelector = *o.ProcessingUnitSelector
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ProxyProtocolEnabled != nil {
		out.ProxyProtocolEnabled = *o.ProxyProtocolEnabled
	}
	if o.ProxyProtocolSubnets != nil {
		out.ProxyProtocolSubnets = *o.ProxyProtocolSubnets
	}
	if o.PublicApplicationPort != nil {
		out.PublicApplicationPort = *o.PublicApplicationPort
	}
	if o.TrustedCertificateAuthorities != nil {
		out.TrustedCertificateAuthorities = *o.TrustedCertificateAuthorities
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
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
func (o *SparseLoadBalancer) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetArchived returns the Archived of the receiver.
func (o *SparseLoadBalancer) GetArchived() (out bool) {

	if o.Archived == nil {
		return
	}

	return *o.Archived
}

// SetArchived sets the property Archived of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetArchived(archived bool) {

	o.Archived = &archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseLoadBalancer) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseLoadBalancer) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseLoadBalancer) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseLoadBalancer) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseLoadBalancer) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseLoadBalancer) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseLoadBalancer) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseLoadBalancer) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseLoadBalancer) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseLoadBalancer) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseLoadBalancer) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseLoadBalancer) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseLoadBalancer) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseLoadBalancer) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseLoadBalancer) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseLoadBalancer) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseLoadBalancer) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseLoadBalancer.
func (o *SparseLoadBalancer) DeepCopy() *SparseLoadBalancer {

	if o == nil {
		return nil
	}

	out := &SparseLoadBalancer{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLoadBalancer.
func (o *SparseLoadBalancer) DeepCopyInto(out *SparseLoadBalancer) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLoadBalancer: %s", err))
	}

	*out = *target.(*SparseLoadBalancer)
}

type mongoAttributesLoadBalancer struct {
	ID                            bson.ObjectId         `bson:"_id,omitempty"`
	IPs                           []string              `bson:"ips"`
	TLSCertificatesSelector       [][]string            `bson:"tlscertificatesselector"`
	AllProcessingUnitTags         []string              `bson:"allprocessingunittags"`
	AllTLSCertificateTags         []string              `bson:"alltlscertificatetags"`
	Annotations                   map[string][]string   `bson:"annotations"`
	Archived                      bool                  `bson:"archived"`
	AssociatedTags                []string              `bson:"associatedtags"`
	CreateIdempotencyKey          string                `bson:"createidempotencykey"`
	CreateTime                    time.Time             `bson:"createtime"`
	Description                   string                `bson:"description"`
	Disabled                      bool                  `bson:"disabled"`
	ExposedPort                   int                   `bson:"exposedport"`
	ExposedServiceIsTLS           bool                  `bson:"exposedserviceistls"`
	Hosts                         []string              `bson:"hosts"`
	Metadata                      []string              `bson:"metadata"`
	MigrationsLog                 map[string]string     `bson:"migrationslog,omitempty"`
	Name                          string                `bson:"name"`
	Namespace                     string                `bson:"namespace"`
	NormalizedTags                []string              `bson:"normalizedtags"`
	Port                          int                   `bson:"port"`
	ProcessingUnitSelector        [][]string            `bson:"processingunitselector"`
	Propagate                     bool                  `bson:"propagate"`
	Protected                     bool                  `bson:"protected"`
	ProxyProtocolEnabled          bool                  `bson:"proxyprotocolenabled"`
	ProxyProtocolSubnets          []string              `bson:"proxyprotocolsubnets"`
	PublicApplicationPort         int                   `bson:"publicapplicationport"`
	TrustedCertificateAuthorities string                `bson:"trustedcertificateauthorities"`
	Type                          LoadBalancerTypeValue `bson:"type"`
	UpdateIdempotencyKey          string                `bson:"updateidempotencykey"`
	UpdateTime                    time.Time             `bson:"updatetime"`
	ZHash                         int                   `bson:"zhash"`
	Zone                          int                   `bson:"zone"`
}
type mongoAttributesSparseLoadBalancer struct {
	ID                            bson.ObjectId          `bson:"_id,omitempty"`
	IPs                           *[]string              `bson:"ips,omitempty"`
	TLSCertificatesSelector       *[][]string            `bson:"tlscertificatesselector,omitempty"`
	AllProcessingUnitTags         *[]string              `bson:"allprocessingunittags,omitempty"`
	AllTLSCertificateTags         *[]string              `bson:"alltlscertificatetags,omitempty"`
	Annotations                   *map[string][]string   `bson:"annotations,omitempty"`
	Archived                      *bool                  `bson:"archived,omitempty"`
	AssociatedTags                *[]string              `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey          *string                `bson:"createidempotencykey,omitempty"`
	CreateTime                    *time.Time             `bson:"createtime,omitempty"`
	Description                   *string                `bson:"description,omitempty"`
	Disabled                      *bool                  `bson:"disabled,omitempty"`
	ExposedPort                   *int                   `bson:"exposedport,omitempty"`
	ExposedServiceIsTLS           *bool                  `bson:"exposedserviceistls,omitempty"`
	Hosts                         *[]string              `bson:"hosts,omitempty"`
	Metadata                      *[]string              `bson:"metadata,omitempty"`
	MigrationsLog                 *map[string]string     `bson:"migrationslog,omitempty"`
	Name                          *string                `bson:"name,omitempty"`
	Namespace                     *string                `bson:"namespace,omitempty"`
	NormalizedTags                *[]string              `bson:"normalizedtags,omitempty"`
	Port                          *int                   `bson:"port,omitempty"`
	ProcessingUnitSelector        *[][]string            `bson:"processingunitselector,omitempty"`
	Propagate                     *bool                  `bson:"propagate,omitempty"`
	Protected                     *bool                  `bson:"protected,omitempty"`
	ProxyProtocolEnabled          *bool                  `bson:"proxyprotocolenabled,omitempty"`
	ProxyProtocolSubnets          *[]string              `bson:"proxyprotocolsubnets,omitempty"`
	PublicApplicationPort         *int                   `bson:"publicapplicationport,omitempty"`
	TrustedCertificateAuthorities *string                `bson:"trustedcertificateauthorities,omitempty"`
	Type                          *LoadBalancerTypeValue `bson:"type,omitempty"`
	UpdateIdempotencyKey          *string                `bson:"updateidempotencykey,omitempty"`
	UpdateTime                    *time.Time             `bson:"updatetime,omitempty"`
	ZHash                         *int                   `bson:"zhash,omitempty"`
	Zone                          *int                   `bson:"zone,omitempty"`
}
