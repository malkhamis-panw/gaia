package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// AccountStatusValue represents the possible values for attribute "status".
type AccountStatusValue string

const (
	// AccountStatusActive represents the value Active.
	AccountStatusActive AccountStatusValue = "Active"

	// AccountStatusDisabled represents the value Disabled.
	AccountStatusDisabled AccountStatusValue = "Disabled"

	// AccountStatusInvited represents the value Invited.
	AccountStatusInvited AccountStatusValue = "Invited"

	// AccountStatusPending represents the value Pending.
	AccountStatusPending AccountStatusValue = "Pending"
)

// AccountIdentity represents the Identity of the object
var AccountIdentity = elemental.Identity{
	Name:     "account",
	Category: "accounts",
}

// AccountsList represents a list of Accounts
type AccountsList []*Account

// ContentIdentity returns the identity of the objects in the list.
func (o AccountsList) ContentIdentity() elemental.Identity {

	return AccountIdentity
}

// Copy returns a pointer to a copy the AccountsList.
func (o AccountsList) Copy() elemental.ContentIdentifiable {

	copy := append(AccountsList{}, o...)
	return &copy
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccountsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccountsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AccountsList) Version() int {

	return 1
}

// Account represents the model of a account
type Account struct {
	// ID of the object.
	ID string `json:"ID" bson:"_id"`

	// LDAPAddress holds the account authentication account's private ldap server.
	LDAPAddress string `json:"LDAPAddress" bson:"ldapaddress"`

	// LDAPBaseDN holds the base DN to use to ldap queries.
	LDAPBaseDN string `json:"LDAPBaseDN" bson:"ldapbasedn"`

	// LDAPBindDN holds the account's internal LDAP bind dn for querying auth.
	LDAPBindDN string `json:"LDAPBindDN" bson:"ldapbinddn"`

	// LDAPBindPassword holds the password to the LDAPBindDN.
	LDAPBindPassword string `json:"LDAPBindPassword" bson:"ldapbindpassword"`

	// LDAPCertificateAuthority contains the optional certificate authority that will be used to connect to the LDAP server. It is not needed if the TLS certificate of the LDAP is issued from a public truster CA.
	LDAPCertificateAuthority string `json:"LDAPCertificateAuthority" bson:"ldapcertificateauthority"`

	// LDAPEnabled triggers if the account uses it's own LDAP for authentication.
	LDAPEnabled bool `json:"LDAPEnabled" bson:"ldapenabled"`

	// AccessEnabled defines if the account holder should have access to the systems.
	AccessEnabled bool `json:"accessEnabled" bson:"accessenabled"`

	// ActivationExpiration contains the expiration date of the activation token.
	ActivationExpiration time.Time `json:"-" bson:"activationexpiration"`

	// ActivationToken contains the activation token.
	ActivationToken string `json:"-" bson:"activationtoken"`

	// AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.
	AssociatedAPIAuthPolicyID string `json:"associatedAPIAuthPolicyID" bson:"associatedapiauthpolicyid"`

	// AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.
	AssociatedAWSPolicies map[string]string `json:"-" bson:"associatedawspolicies"`

	// AssociatedNSMappingPolicyID contains the ID of the associated Namespace Mapping Policy.
	AssociatedNSMappingPolicyID string `json:"associatedNSMappingPolicyID" bson:"associatednsmappingpolicyid"`

	// AssociatedNamespaceID contains the ID of the associated namespace.
	AssociatedNamespaceID string `json:"associatedNamespaceID" bson:"associatednamespaceid"`

	// AssociatedPlanKey contains the plan key that is associated to this account.
	AssociatedPlanKey string `json:"associatedPlanKey" bson:"associatedplankey"`

	// AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.
	AssociatedQuotaPolicies map[string]string `json:"-" bson:"associatedquotapolicies"`

	// BumpToken contains the tag processing unit must use to be placed in the account's namespace.
	BumpToken string `json:"bumpToken" bson:"bumptoken"`

	// Company of the account user.
	Company string `json:"company" bson:"company"`

	// createdAt represents the creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Email of the account holder.
	Email string `json:"email" bson:"email"`

	// First Name of the account user.
	FirstName string `json:"firstName" bson:"firstname"`

	// Last Name of the account user.
	LastName string `json:"lastName" bson:"lastname"`

	// Name of the account.
	Name string `json:"name" bson:"name"`

	// Password for the account.
	Password string `json:"password" bson:"password"`

	// ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.
	ReCAPTCHAKey string `json:"reCAPTCHAKey" bson:"-"`

	// ResetPasswordExpiration contains the expiration time for reseting the password.
	ResetPasswordExpiration time.Time `json:"-" bson:"resetpasswordexpiration"`

	// ResetPasswordToken contains the token to use for resetting password.
	ResetPasswordToken string `json:"-" bson:"resetpasswordtoken"`

	// Status of the account.
	Status AccountStatusValue `json:"status" bson:"status"`

	// UpdateTime represents the last update date of the objct.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAccount returns a new *Account
func NewAccount() *Account {

	return &Account{
		ModelVersion:            1,
		AssociatedAWSPolicies:   map[string]string{},
		AssociatedPlanKey:       "aporeto.plan.free",
		AssociatedQuotaPolicies: map[string]string{},
		Status:                  "Pending",
	}
}

// Identity returns the Identity of the object.
func (o *Account) Identity() elemental.Identity {

	return AccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Account) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Account) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Account) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Account) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Account) Doc() string {
	return `Manage your Account.`
}

func (o *Account) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Account) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[^\*\=]*$`, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Disabled", "Invited", "Pending"}, true); err != nil {
		errors = append(errors, err)
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
func (*Account) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccountAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccountLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Account) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccountAttributesMap
}

// AccountAttributesMap represents the map of attribute for Account.
var AccountAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPAddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPAddress holds the account authentication account's private ldap server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBaseDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBaseDN holds the base DN to use to ldap queries.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBindDN holds the account's internal LDAP bind dn for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindPassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBindPassword holds the password to the LDAPBindDN. `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPCertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPCertificateAuthority contains the optional certificate authority that will be used to connect to the LDAP server. It is not needed if the TLS certificate of the LDAP is issued from a public truster CA.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPCertificateAuthority",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPEnabled triggers if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"AccessEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ActivationExpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ActivationExpiration contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ActivationToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ActivationToken contains the activation token.`,
		Format:         "free",
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedAPIAuthPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAWSPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.`,
		Name:           "associatedAWSPolicies",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"AssociatedNSMappingPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedNSMappingPolicyID contains the ID of the associated Namespace Mapping Policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedNSMappingPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedNamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedNamespaceID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedPlanKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		DefaultValue:   "aporeto.plan.free",
		Description:    `AssociatedPlanKey contains the plan key that is associated to this account.`,
		Exposed:        true,
		Format:         "free",
		Name:           "associatedPlanKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedQuotaPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.`,
		Name:           "associatedQuotaPolicies",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"BumpToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `BumpToken contains the tag processing unit must use to be placed in the account's namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "bumpToken",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Company": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "company",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `createdAt represents the creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"FirstName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `First Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "firstName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LastName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Last Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "lastName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Password for the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ReCAPTCHAKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.`,
		Exposed:        true,
		Format:         "free",
		Name:           "reCAPTCHAKey",
		Type:           "string",
	},
	"ResetPasswordExpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ResetPasswordExpiration contains the expiration time for reseting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ResetPasswordToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ResetPasswordToken contains the token to use for resetting password.`,
		Format:         "free",
		Name:           "resetPasswordToken",
		Stored:         true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		DefaultValue:   AccountStatusValue("Pending"),
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime represents the last update date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// AccountLowerCaseAttributesMap represents the map of attribute for Account.
var AccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ldapaddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPAddress holds the account authentication account's private ldap server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbasedn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBaseDN holds the base DN to use to ldap queries.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbinddn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBindDN holds the account's internal LDAP bind dn for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbindpassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBindPassword holds the password to the LDAPBindDN. `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapcertificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPCertificateAuthority contains the optional certificate authority that will be used to connect to the LDAP server. It is not needed if the TLS certificate of the LDAP is issued from a public truster CA.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPCertificateAuthority",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPEnabled triggers if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"accessenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"activationexpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ActivationExpiration contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"activationtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ActivationToken contains the activation token.`,
		Format:         "free",
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"associatedapiauthpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedAPIAuthPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedawspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.`,
		Name:           "associatedAWSPolicies",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"associatednsmappingpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedNSMappingPolicyID contains the ID of the associated Namespace Mapping Policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedNSMappingPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatednamespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedNamespaceID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedplankey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		DefaultValue:   "aporeto.plan.free",
		Description:    `AssociatedPlanKey contains the plan key that is associated to this account.`,
		Exposed:        true,
		Format:         "free",
		Name:           "associatedPlanKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedquotapolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.`,
		Name:           "associatedQuotaPolicies",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"bumptoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `BumpToken contains the tag processing unit must use to be placed in the account's namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "bumpToken",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"company": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "company",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `createdAt represents the creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"firstname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `First Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "firstName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"lastname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Last Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "lastName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Password for the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"recaptchakey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.`,
		Exposed:        true,
		Format:         "free",
		Name:           "reCAPTCHAKey",
		Type:           "string",
	},
	"resetpasswordexpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ResetPasswordExpiration contains the expiration time for reseting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"resetpasswordtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ResetPasswordToken contains the token to use for resetting password.`,
		Format:         "free",
		Name:           "resetPasswordToken",
		Stored:         true,
		Type:           "string",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		DefaultValue:   AccountStatusValue("Pending"),
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime represents the last update date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
