package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// InvoiceRecordIndexes lists the attribute compound indexes.
var InvoiceRecordIndexes = [][]string{}

// InvoiceRecordIdentity represents the Identity of the object.
var InvoiceRecordIdentity = elemental.Identity{
	Name:     "invoicerecord",
	Category: "invoicerecords",
	Private:  false,
}

// InvoiceRecordsList represents a list of InvoiceRecords
type InvoiceRecordsList []*InvoiceRecord

// Identity returns the identity of the objects in the list.
func (o InvoiceRecordsList) Identity() elemental.Identity {

	return InvoiceRecordIdentity
}

// Copy returns a pointer to a copy the InvoiceRecordsList.
func (o InvoiceRecordsList) Copy() elemental.Identifiables {

	copy := append(InvoiceRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InvoiceRecordsList.
func (o InvoiceRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InvoiceRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*InvoiceRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InvoiceRecordsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InvoiceRecordsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o InvoiceRecordsList) Version() int {

	return 1
}

// InvoiceRecord represents the model of a invoicerecord
type InvoiceRecord struct {
	// ID is the id of this invoice record.
	ID string `json:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// InvoiceID references the id of the invoice that this invoice record provides
	// details for.
	InvoiceID string `json:"invoiceID" bson:"invoiceid" mapstructure:"invoiceID,omitempty"`

	// InvoiceRecords provides details about billing units.
	InvoiceRecords []string `json:"invoiceRecords" bson:"invoicerecords" mapstructure:"invoiceRecords,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewInvoiceRecord returns a new *InvoiceRecord
func NewInvoiceRecord() *InvoiceRecord {

	return &InvoiceRecord{
		ModelVersion:   1,
		InvoiceRecords: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *InvoiceRecord) Identity() elemental.Identity {

	return InvoiceRecordIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InvoiceRecord) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InvoiceRecord) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *InvoiceRecord) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *InvoiceRecord) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *InvoiceRecord) Doc() string {
	return `This api allows to view detailed records of invoices for Aporeto customers.`
}

func (o *InvoiceRecord) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *InvoiceRecord) Validate() error {

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
func (*InvoiceRecord) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InvoiceRecordAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InvoiceRecordLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*InvoiceRecord) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InvoiceRecordAttributesMap
}

// InvoiceRecordAttributesMap represents the map of attribute for InvoiceRecord.
var InvoiceRecordAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the id of this invoice record.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"InvoiceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceID",
		Description: `InvoiceID references the id of the invoice that this invoice record provides
details for.`,
		Exposed:    true,
		Filterable: true,
		Name:       "invoiceID",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"InvoiceRecords": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceRecords",
		Description:    `InvoiceRecords provides details about billing units.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "invoiceRecords",
		Orderable:      true,
		Stored:         true,
		SubType:        "invoicerecord_list",
		Type:           "external",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// InvoiceRecordLowerCaseAttributesMap represents the map of attribute for InvoiceRecord.
var InvoiceRecordLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the id of this invoice record.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"invoiceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceID",
		Description: `InvoiceID references the id of the invoice that this invoice record provides
details for.`,
		Exposed:    true,
		Filterable: true,
		Name:       "invoiceID",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"invoicerecords": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceRecords",
		Description:    `InvoiceRecords provides details about billing units.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "invoiceRecords",
		Orderable:      true,
		Stored:         true,
		SubType:        "invoicerecord_list",
		Type:           "external",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
