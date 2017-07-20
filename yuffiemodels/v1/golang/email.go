package yuffiemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// EmailIdentity represents the Identity of the object
var EmailIdentity = elemental.Identity{
	Name:     "email",
	Category: "emails",
}

// EmailsList represents a list of Emails
type EmailsList []*Email

// ContentIdentity returns the identity of the objects in the list.
func (o EmailsList) ContentIdentity() elemental.Identity {

	return EmailIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o EmailsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EmailsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o EmailsList) Version() int {

	return 1.0
}

// Email represents the model of a email
type Email struct {
	// Attachments is a list of attachments to send
	Attachments map[string]string `json:"attachments" bson:"-"`

	// Bcc represents email that should be in copy but hidden
	Bcc []string `json:"bcc" bson:"-"`

	// Cc represents the addresses that should be in copy
	Cc []string `json:"cc" bson:"-"`

	// Content of the email to send
	Content string `json:"content" bson:"-"`

	// From represents the sender of the email
	From string `json:"from" bson:"-"`

	// Subject represents the subject of the email
	Subject string `json:"subject" bson:"-"`

	// To represents receivers of the email
	To []string `json:"to" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEmail returns a new *Email
func NewEmail() *Email {

	return &Email{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Email) Identity() elemental.Identity {

	return EmailIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Email) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Email) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Email) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Email) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Email) Doc() string {
	return `Email is a message that can be send via email`
}

func (o *Email) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Email) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("from", o.From); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("from", o.From); err != nil {
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
func (*Email) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EmailAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EmailLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Email) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EmailAttributesMap
}

// EmailAttributesMap represents the map of attribute for Email.
var EmailAttributesMap = map[string]elemental.AttributeSpecification{
	"Attachments": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Attachments is a list of attachments to send`,
		Exposed:        true,
		Name:           "attachments",
		SubType:        "list_attachments",
		Type:           "external",
	},
	"Bcc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Bcc represents email that should be in copy but hidden `,
		Exposed:        true,
		Name:           "bcc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Cc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Cc represents the addresses that should be in copy`,
		Exposed:        true,
		Name:           "cc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Content of the email to send`,
		Exposed:        true,
		Format:         "free",
		Name:           "content",
		Type:           "string",
	},
	"From": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `From represents the sender of the email`,
		Exposed:        true,
		Format:         "email",
		Name:           "from",
		Required:       true,
		Type:           "string",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject represents the subject of the email`,
		Exposed:        true,
		Format:         "free",
		Name:           "subject",
		Type:           "string",
	},
	"To": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `To represents receivers of the email `,
		Exposed:        true,
		Name:           "to",
		SubType:        "list_emails",
		Type:           "external",
	},
}

// EmailLowerCaseAttributesMap represents the map of attribute for Email.
var EmailLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"attachments": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Attachments is a list of attachments to send`,
		Exposed:        true,
		Name:           "attachments",
		SubType:        "list_attachments",
		Type:           "external",
	},
	"bcc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Bcc represents email that should be in copy but hidden `,
		Exposed:        true,
		Name:           "bcc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"cc": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Cc represents the addresses that should be in copy`,
		Exposed:        true,
		Name:           "cc",
		SubType:        "list_emails",
		Type:           "external",
	},
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Content of the email to send`,
		Exposed:        true,
		Format:         "free",
		Name:           "content",
		Type:           "string",
	},
	"from": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `From represents the sender of the email`,
		Exposed:        true,
		Format:         "email",
		Name:           "from",
		Required:       true,
		Type:           "string",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject represents the subject of the email`,
		Exposed:        true,
		Format:         "free",
		Name:           "subject",
		Type:           "string",
	},
	"to": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `To represents receivers of the email `,
		Exposed:        true,
		Name:           "to",
		SubType:        "list_emails",
		Type:           "external",
	},
}
