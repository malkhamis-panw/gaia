package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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

	return []string{}
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
	// The list of IP addresses where the service can be accessed. This is an optional
	// attribute and is only required if no host names are provided. The system will
	// automatically resolve IP addresses from host names otherwise.
	IPs []string `json:"IPs" msgpack:"IPs" bson:"ips" mapstructure:"IPs,omitempty"`

	// A tag or tag expression that identifies the TLS certificates to be used by
	// enforcers when exposing the service ran by the processing units.
	TLSCertificateSelector [][]string `json:"TLSCertificateSelector" msgpack:"TLSCertificateSelector" bson:"tlscertificateselector" mapstructure:"TLSCertificateSelector,omitempty"`

	// This is a set of all selector tags for matching in the database.
	AllServiceTags []string `json:"-" msgpack:"-" bson:"allservicetags" mapstructure:"-,omitempty"`

	// Resolves the API endpoints that the service is exposing. Only valid during
	// policy rendering.
	Endpoints []*Endpoint `json:"endpoints" msgpack:"endpoints" bson:"-" mapstructure:"endpoints,omitempty"`

	// The port that the service can be accessed on. Note that this is different from
	// the `port` attribute that describes the port that the service is actually
	// listening on. For example if a load balancer is used, the `exposedPort` is
	// the port that the load balancer is listening for the service, whereas the
	// port that the implementation is listening on can be different.
	ExposedPort int `json:"exposedPort" msgpack:"exposedPort" bson:"exposedport" mapstructure:"exposedPort,omitempty"`

	// Indicates that the exposed service is TLS. This means that the enforcer has to
	// initiate a TLS session in order to forward traffic to the service.
	ExposedServiceIsTLS bool `json:"exposedServiceIsTLS" msgpack:"exposedServiceIsTLS" bson:"exposedserviceistls" mapstructure:"exposedServiceIsTLS,omitempty"`

	// Indicates if this is an external service.
	External bool `json:"external" msgpack:"external" bson:"external" mapstructure:"external,omitempty"`

	// The host names that the service can be accessed on.
	Hosts []string `json:"hosts" msgpack:"hosts" bson:"hosts" mapstructure:"hosts,omitempty"`

	// The port that the implementation of the service is listening to. It can be
	// different than `exposedPort`. This is needed for port mapping use cases
	// where there are private and public ports.
	Port int `json:"port" msgpack:"port" bson:"port" mapstructure:"port,omitempty"`

	// Tag expression that identifies the processing unit that implements this
	// particular service.
	ProcessingUnitSelector [][]string `json:"processingUnitSelector" msgpack:"processingUnitSelector" bson:"processingunitselector" mapstructure:"processingUnitSelector,omitempty"`

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

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLoadBalancer returns a new *LoadBalancer
func NewLoadBalancer() *LoadBalancer {

	return &LoadBalancer{
		ModelVersion:           1,
		External:               false,
		AllServiceTags:         []string{},
		Hosts:                  []string{},
		Endpoints:              []*Endpoint{},
		ExposedServiceIsTLS:    false,
		IPs:                    []string{},
		TLSCertificateSelector: [][]string{},
		ProcessingUnitSelector: [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *LoadBalancer) Identity() elemental.Identity {

	return LoadBalancerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LoadBalancer) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LoadBalancer) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LoadBalancer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLoadBalancer{}

	s.IPs = o.IPs
	s.TLSCertificateSelector = o.TLSCertificateSelector
	s.AllServiceTags = o.AllServiceTags
	s.ExposedPort = o.ExposedPort
	s.ExposedServiceIsTLS = o.ExposedServiceIsTLS
	s.External = o.External
	s.Hosts = o.Hosts
	s.Port = o.Port
	s.ProcessingUnitSelector = o.ProcessingUnitSelector
	s.PublicApplicationPort = o.PublicApplicationPort
	s.TrustedCertificateAuthorities = o.TrustedCertificateAuthorities

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

	o.IPs = s.IPs
	o.TLSCertificateSelector = s.TLSCertificateSelector
	o.AllServiceTags = s.AllServiceTags
	o.ExposedPort = s.ExposedPort
	o.ExposedServiceIsTLS = s.ExposedServiceIsTLS
	o.External = s.External
	o.Hosts = s.Hosts
	o.Port = s.Port
	o.ProcessingUnitSelector = s.ProcessingUnitSelector
	o.PublicApplicationPort = s.PublicApplicationPort
	o.TrustedCertificateAuthorities = s.TrustedCertificateAuthorities

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

	return []string{}
}

// Doc returns the documentation for the object
func (o *LoadBalancer) Doc() string {

	return `Defines a generic external LoadBalancer that sits between 2 enforcers.`
}

func (o *LoadBalancer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LoadBalancer) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLoadBalancer{
			IPs:                           &o.IPs,
			TLSCertificateSelector:        &o.TLSCertificateSelector,
			AllServiceTags:                &o.AllServiceTags,
			Endpoints:                     &o.Endpoints,
			ExposedPort:                   &o.ExposedPort,
			ExposedServiceIsTLS:           &o.ExposedServiceIsTLS,
			External:                      &o.External,
			Hosts:                         &o.Hosts,
			Port:                          &o.Port,
			ProcessingUnitSelector:        &o.ProcessingUnitSelector,
			PublicApplicationPort:         &o.PublicApplicationPort,
			TrustedCertificateAuthorities: &o.TrustedCertificateAuthorities,
		}
	}

	sp := &SparseLoadBalancer{}
	for _, f := range fields {
		switch f {
		case "IPs":
			sp.IPs = &(o.IPs)
		case "TLSCertificateSelector":
			sp.TLSCertificateSelector = &(o.TLSCertificateSelector)
		case "allServiceTags":
			sp.AllServiceTags = &(o.AllServiceTags)
		case "endpoints":
			sp.Endpoints = &(o.Endpoints)
		case "exposedPort":
			sp.ExposedPort = &(o.ExposedPort)
		case "exposedServiceIsTLS":
			sp.ExposedServiceIsTLS = &(o.ExposedServiceIsTLS)
		case "external":
			sp.External = &(o.External)
		case "hosts":
			sp.Hosts = &(o.Hosts)
		case "port":
			sp.Port = &(o.Port)
		case "processingUnitSelector":
			sp.ProcessingUnitSelector = &(o.ProcessingUnitSelector)
		case "publicApplicationPort":
			sp.PublicApplicationPort = &(o.PublicApplicationPort)
		case "trustedCertificateAuthorities":
			sp.TrustedCertificateAuthorities = &(o.TrustedCertificateAuthorities)
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
	if so.IPs != nil {
		o.IPs = *so.IPs
	}
	if so.TLSCertificateSelector != nil {
		o.TLSCertificateSelector = *so.TLSCertificateSelector
	}
	if so.AllServiceTags != nil {
		o.AllServiceTags = *so.AllServiceTags
	}
	if so.Endpoints != nil {
		o.Endpoints = *so.Endpoints
	}
	if so.ExposedPort != nil {
		o.ExposedPort = *so.ExposedPort
	}
	if so.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = *so.ExposedServiceIsTLS
	}
	if so.External != nil {
		o.External = *so.External
	}
	if so.Hosts != nil {
		o.Hosts = *so.Hosts
	}
	if so.Port != nil {
		o.Port = *so.Port
	}
	if so.ProcessingUnitSelector != nil {
		o.ProcessingUnitSelector = *so.ProcessingUnitSelector
	}
	if so.PublicApplicationPort != nil {
		o.PublicApplicationPort = *so.PublicApplicationPort
	}
	if so.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = *so.TrustedCertificateAuthorities
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

	if err := ValidateTagsExpression("TLSCertificateSelector", o.TLSCertificateSelector); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.Endpoints {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredInt("exposedPort", o.ExposedPort); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("exposedPort", o.ExposedPort, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("port", o.Port, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("processingUnitSelector", o.ProcessingUnitSelector); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("publicApplicationPort", o.PublicApplicationPort, int(65535), false); err != nil {
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
	case "IPs":
		return o.IPs
	case "TLSCertificateSelector":
		return o.TLSCertificateSelector
	case "allServiceTags":
		return o.AllServiceTags
	case "endpoints":
		return o.Endpoints
	case "exposedPort":
		return o.ExposedPort
	case "exposedServiceIsTLS":
		return o.ExposedServiceIsTLS
	case "external":
		return o.External
	case "hosts":
		return o.Hosts
	case "port":
		return o.Port
	case "processingUnitSelector":
		return o.ProcessingUnitSelector
	case "publicApplicationPort":
		return o.PublicApplicationPort
	case "trustedCertificateAuthorities":
		return o.TrustedCertificateAuthorities
	}

	return nil
}

// LoadBalancerAttributesMap represents the map of attribute for LoadBalancer.
var LoadBalancerAttributesMap = map[string]elemental.AttributeSpecification{
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
	"TLSCertificateSelector": {
		AllowedChoices: []string{},
		BSONFieldName:  "tlscertificateselector",
		ConvertedName:  "TLSCertificateSelector",
		Description: `A tag or tag expression that identifies the TLS certificates to be used by
enforcers when exposing the service ran by the processing units.`,
		Exposed: true,
		Name:    "TLSCertificateSelector",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"AllServiceTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allservicetags",
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the database.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Endpoints": {
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Resolves the API endpoints that the service is exposing. Only valid during
policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "endpoint",
		Type:     "refList",
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
	"External": {
		AllowedChoices: []string{},
		BSONFieldName:  "external",
		ConvertedName:  "External",
		Description:    `Indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
}

// LoadBalancerLowerCaseAttributesMap represents the map of attribute for LoadBalancer.
var LoadBalancerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"tlscertificateselector": {
		AllowedChoices: []string{},
		BSONFieldName:  "tlscertificateselector",
		ConvertedName:  "TLSCertificateSelector",
		Description: `A tag or tag expression that identifies the TLS certificates to be used by
enforcers when exposing the service ran by the processing units.`,
		Exposed: true,
		Name:    "TLSCertificateSelector",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"allservicetags": {
		AllowedChoices: []string{},
		BSONFieldName:  "allservicetags",
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the database.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"endpoints": {
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Resolves the API endpoints that the service is exposing. Only valid during
policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "endpoint",
		Type:     "refList",
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
	"external": {
		AllowedChoices: []string{},
		BSONFieldName:  "external",
		ConvertedName:  "External",
		Description:    `Indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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

	return []string{}
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
	// The list of IP addresses where the service can be accessed. This is an optional
	// attribute and is only required if no host names are provided. The system will
	// automatically resolve IP addresses from host names otherwise.
	IPs *[]string `json:"IPs,omitempty" msgpack:"IPs,omitempty" bson:"ips,omitempty" mapstructure:"IPs,omitempty"`

	// A tag or tag expression that identifies the TLS certificates to be used by
	// enforcers when exposing the service ran by the processing units.
	TLSCertificateSelector *[][]string `json:"TLSCertificateSelector,omitempty" msgpack:"TLSCertificateSelector,omitempty" bson:"tlscertificateselector,omitempty" mapstructure:"TLSCertificateSelector,omitempty"`

	// This is a set of all selector tags for matching in the database.
	AllServiceTags *[]string `json:"-" msgpack:"-" bson:"allservicetags,omitempty" mapstructure:"-,omitempty"`

	// Resolves the API endpoints that the service is exposing. Only valid during
	// policy rendering.
	Endpoints *[]*Endpoint `json:"endpoints,omitempty" msgpack:"endpoints,omitempty" bson:"-" mapstructure:"endpoints,omitempty"`

	// The port that the service can be accessed on. Note that this is different from
	// the `port` attribute that describes the port that the service is actually
	// listening on. For example if a load balancer is used, the `exposedPort` is
	// the port that the load balancer is listening for the service, whereas the
	// port that the implementation is listening on can be different.
	ExposedPort *int `json:"exposedPort,omitempty" msgpack:"exposedPort,omitempty" bson:"exposedport,omitempty" mapstructure:"exposedPort,omitempty"`

	// Indicates that the exposed service is TLS. This means that the enforcer has to
	// initiate a TLS session in order to forward traffic to the service.
	ExposedServiceIsTLS *bool `json:"exposedServiceIsTLS,omitempty" msgpack:"exposedServiceIsTLS,omitempty" bson:"exposedserviceistls,omitempty" mapstructure:"exposedServiceIsTLS,omitempty"`

	// Indicates if this is an external service.
	External *bool `json:"external,omitempty" msgpack:"external,omitempty" bson:"external,omitempty" mapstructure:"external,omitempty"`

	// The host names that the service can be accessed on.
	Hosts *[]string `json:"hosts,omitempty" msgpack:"hosts,omitempty" bson:"hosts,omitempty" mapstructure:"hosts,omitempty"`

	// The port that the implementation of the service is listening to. It can be
	// different than `exposedPort`. This is needed for port mapping use cases
	// where there are private and public ports.
	Port *int `json:"port,omitempty" msgpack:"port,omitempty" bson:"port,omitempty" mapstructure:"port,omitempty"`

	// Tag expression that identifies the processing unit that implements this
	// particular service.
	ProcessingUnitSelector *[][]string `json:"processingUnitSelector,omitempty" msgpack:"processingUnitSelector,omitempty" bson:"processingunitselector,omitempty" mapstructure:"processingUnitSelector,omitempty"`

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

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLoadBalancer) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseLoadBalancer) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseLoadBalancer{}

	if o.IPs != nil {
		s.IPs = o.IPs
	}
	if o.TLSCertificateSelector != nil {
		s.TLSCertificateSelector = o.TLSCertificateSelector
	}
	if o.AllServiceTags != nil {
		s.AllServiceTags = o.AllServiceTags
	}
	if o.ExposedPort != nil {
		s.ExposedPort = o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		s.ExposedServiceIsTLS = o.ExposedServiceIsTLS
	}
	if o.External != nil {
		s.External = o.External
	}
	if o.Hosts != nil {
		s.Hosts = o.Hosts
	}
	if o.Port != nil {
		s.Port = o.Port
	}
	if o.ProcessingUnitSelector != nil {
		s.ProcessingUnitSelector = o.ProcessingUnitSelector
	}
	if o.PublicApplicationPort != nil {
		s.PublicApplicationPort = o.PublicApplicationPort
	}
	if o.TrustedCertificateAuthorities != nil {
		s.TrustedCertificateAuthorities = o.TrustedCertificateAuthorities
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

	if s.IPs != nil {
		o.IPs = s.IPs
	}
	if s.TLSCertificateSelector != nil {
		o.TLSCertificateSelector = s.TLSCertificateSelector
	}
	if s.AllServiceTags != nil {
		o.AllServiceTags = s.AllServiceTags
	}
	if s.ExposedPort != nil {
		o.ExposedPort = s.ExposedPort
	}
	if s.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = s.ExposedServiceIsTLS
	}
	if s.External != nil {
		o.External = s.External
	}
	if s.Hosts != nil {
		o.Hosts = s.Hosts
	}
	if s.Port != nil {
		o.Port = s.Port
	}
	if s.ProcessingUnitSelector != nil {
		o.ProcessingUnitSelector = s.ProcessingUnitSelector
	}
	if s.PublicApplicationPort != nil {
		o.PublicApplicationPort = s.PublicApplicationPort
	}
	if s.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = s.TrustedCertificateAuthorities
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
	if o.IPs != nil {
		out.IPs = *o.IPs
	}
	if o.TLSCertificateSelector != nil {
		out.TLSCertificateSelector = *o.TLSCertificateSelector
	}
	if o.AllServiceTags != nil {
		out.AllServiceTags = *o.AllServiceTags
	}
	if o.Endpoints != nil {
		out.Endpoints = *o.Endpoints
	}
	if o.ExposedPort != nil {
		out.ExposedPort = *o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		out.ExposedServiceIsTLS = *o.ExposedServiceIsTLS
	}
	if o.External != nil {
		out.External = *o.External
	}
	if o.Hosts != nil {
		out.Hosts = *o.Hosts
	}
	if o.Port != nil {
		out.Port = *o.Port
	}
	if o.ProcessingUnitSelector != nil {
		out.ProcessingUnitSelector = *o.ProcessingUnitSelector
	}
	if o.PublicApplicationPort != nil {
		out.PublicApplicationPort = *o.PublicApplicationPort
	}
	if o.TrustedCertificateAuthorities != nil {
		out.TrustedCertificateAuthorities = *o.TrustedCertificateAuthorities
	}

	return out
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
	IPs                           []string   `bson:"ips"`
	TLSCertificateSelector        [][]string `bson:"tlscertificateselector"`
	AllServiceTags                []string   `bson:"allservicetags"`
	ExposedPort                   int        `bson:"exposedport"`
	ExposedServiceIsTLS           bool       `bson:"exposedserviceistls"`
	External                      bool       `bson:"external"`
	Hosts                         []string   `bson:"hosts"`
	Port                          int        `bson:"port"`
	ProcessingUnitSelector        [][]string `bson:"processingunitselector"`
	PublicApplicationPort         int        `bson:"publicapplicationport"`
	TrustedCertificateAuthorities string     `bson:"trustedcertificateauthorities"`
}
type mongoAttributesSparseLoadBalancer struct {
	IPs                           *[]string   `bson:"ips,omitempty"`
	TLSCertificateSelector        *[][]string `bson:"tlscertificateselector,omitempty"`
	AllServiceTags                *[]string   `bson:"allservicetags,omitempty"`
	ExposedPort                   *int        `bson:"exposedport,omitempty"`
	ExposedServiceIsTLS           *bool       `bson:"exposedserviceistls,omitempty"`
	External                      *bool       `bson:"external,omitempty"`
	Hosts                         *[]string   `bson:"hosts,omitempty"`
	Port                          *int        `bson:"port,omitempty"`
	ProcessingUnitSelector        *[][]string `bson:"processingunitselector,omitempty"`
	PublicApplicationPort         *int        `bson:"publicapplicationport,omitempty"`
	TrustedCertificateAuthorities *string     `bson:"trustedcertificateauthorities,omitempty"`
}
