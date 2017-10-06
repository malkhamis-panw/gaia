package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "github.com/aporeto-inc/gaia/squallmodels/v1/golang/constants"

// RenderedPolicyIdentity represents the Identity of the object
var RenderedPolicyIdentity = elemental.Identity{
	Name:     "renderedpolicy",
	Category: "renderedpolicies",
}

// RenderedPoliciesList represents a list of RenderedPolicies
type RenderedPoliciesList []*RenderedPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o RenderedPoliciesList) ContentIdentity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Copy returns a pointer to a copy the RenderedPoliciesList.
func (o RenderedPoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(RenderedPoliciesList{}, o...)
	return &copy
}

// List converts the object to an elemental.IdentifiablesList.
func (o RenderedPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RenderedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o RenderedPoliciesList) Version() int {

	return 1
}

// RenderedPolicy represents the model of a renderedpolicy
type RenderedPolicy struct {
	// EgressPolicies lists all the egress policies attached to ProcessingUnit
	EgressPolicies map[string]PolicyRulesList `json:"egressPolicies" bson:"-"`

	// IngressPolicies lists all the ingress policies attached to ProcessingUnit
	IngressPolicies map[string]PolicyRulesList `json:"ingressPolicies" bson:"-"`

	// MatchingTags contains the list of tags that matched the policies.
	MatchingTags []string `json:"matchingTags" bson:"-"`

	// Can be set during a POST operation to render a policy on a Processing Unit that has not been created yet.
	ProcessingUnit *ProcessingUnit `json:"processingUnit" bson:"-"`

	// Identifier of the ProcessingUnit
	ProcessingUnitID string `json:"processingUnitID" bson:"-"`

	// Profile is the trust profile of the processing unit that should be used during all communications.
	Profile map[string]string `json:"profile" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRenderedPolicy returns a new *RenderedPolicy
func NewRenderedPolicy() *RenderedPolicy {

	return &RenderedPolicy{
		ModelVersion:    1,
		EgressPolicies:  map[string]PolicyRulesList{string(constants.RenderedPolicyTypeNetwork): PolicyRulesList{}, string(constants.RenderedPolicyTypeFile): PolicyRulesList{}, string(constants.RenderedPolicyTypeSystemCall): PolicyRulesList{}},
		IngressPolicies: map[string]PolicyRulesList{string(constants.RenderedPolicyTypeNetwork): PolicyRulesList{}, string(constants.RenderedPolicyTypeFile): PolicyRulesList{}, string(constants.RenderedPolicyTypeSystemCall): PolicyRulesList{}},
	}
}

// Identity returns the Identity of the object.
func (o *RenderedPolicy) Identity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RenderedPolicy) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RenderedPolicy) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *RenderedPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *RenderedPolicy) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RenderedPolicy) Doc() string {
	return `Retrieve the aggregated policies applied to a particular processing unit.`
}

func (o *RenderedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *RenderedPolicy) Validate() error {

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
func (*RenderedPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RenderedPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RenderedPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RenderedPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RenderedPolicyAttributesMap
}

// RenderedPolicyAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"EgressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `EgressPolicies lists all the egress policies attached to ProcessingUnit`,
		Exposed:        true,
		Name:           "egressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"IngressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `IngressPolicies lists all the ingress policies attached to ProcessingUnit`,
		Exposed:        true,
		Name:           "ingressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"MatchingTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `MatchingTags contains the list of tags that matched the policies.`,
		Exposed:        true,
		Name:           "matchingTags",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"ProcessingUnit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Can be set during a POST operation to render a policy on a Processing Unit that has not been created yet.`,
		Exposed:        true,
		Name:           "processingUnit",
		SubType:        "processingunit",
		Type:           "external",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Identifier of the ProcessingUnit`,
		Exposed:        true,
		Format:         "free",
		Name:           "processingUnitID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Profile": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Profile is the trust profile of the processing unit that should be used during all communications.`,
		Exposed:        true,
		Name:           "profile",
		ReadOnly:       true,
		SubType:        "trust_profile",
		Type:           "external",
	},
}

// RenderedPolicyLowerCaseAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"egresspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `EgressPolicies lists all the egress policies attached to ProcessingUnit`,
		Exposed:        true,
		Name:           "egressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"ingresspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `IngressPolicies lists all the ingress policies attached to ProcessingUnit`,
		Exposed:        true,
		Name:           "ingressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"matchingtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `MatchingTags contains the list of tags that matched the policies.`,
		Exposed:        true,
		Name:           "matchingTags",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"processingunit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Can be set during a POST operation to render a policy on a Processing Unit that has not been created yet.`,
		Exposed:        true,
		Name:           "processingUnit",
		SubType:        "processingunit",
		Type:           "external",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Identifier of the ProcessingUnit`,
		Exposed:        true,
		Format:         "free",
		Name:           "processingUnitID",
		ReadOnly:       true,
		Type:           "string",
	},
	"profile": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Profile is the trust profile of the processing unit that should be used during all communications.`,
		Exposed:        true,
		Name:           "profile",
		ReadOnly:       true,
		SubType:        "trust_profile",
		Type:           "external",
	},
}
