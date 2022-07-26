// Code generated by ent, DO NOT EDIT.

package ent

import (
	"stilyng94/fiber-crm/ent/lead"
	"stilyng94/fiber-crm/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	leadMixin := schema.Lead{}.Mixin()
	leadMixinFields0 := leadMixin[0].Fields()
	_ = leadMixinFields0
	leadFields := schema.Lead{}.Fields()
	_ = leadFields
	// leadDescCreateTime is the schema descriptor for create_time field.
	leadDescCreateTime := leadMixinFields0[0].Descriptor()
	// lead.DefaultCreateTime holds the default value on creation for the create_time field.
	lead.DefaultCreateTime = leadDescCreateTime.Default.(func() time.Time)
	// leadDescUpdateTime is the schema descriptor for update_time field.
	leadDescUpdateTime := leadMixinFields0[1].Descriptor()
	// lead.DefaultUpdateTime holds the default value on creation for the update_time field.
	lead.DefaultUpdateTime = leadDescUpdateTime.Default.(func() time.Time)
	// lead.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	lead.UpdateDefaultUpdateTime = leadDescUpdateTime.UpdateDefault.(func() time.Time)
	// leadDescName is the schema descriptor for name field.
	leadDescName := leadFields[0].Descriptor()
	// lead.NameValidator is a validator for the "name" field. It is called by the builders before save.
	lead.NameValidator = leadDescName.Validators[0].(func(string) error)
	// leadDescCompany is the schema descriptor for company field.
	leadDescCompany := leadFields[1].Descriptor()
	// lead.CompanyValidator is a validator for the "company" field. It is called by the builders before save.
	lead.CompanyValidator = leadDescCompany.Validators[0].(func(string) error)
	// leadDescEmail is the schema descriptor for email field.
	leadDescEmail := leadFields[2].Descriptor()
	// lead.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	lead.EmailValidator = leadDescEmail.Validators[0].(func(string) error)
	// leadDescPhone is the schema descriptor for phone field.
	leadDescPhone := leadFields[3].Descriptor()
	// lead.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	lead.PhoneValidator = leadDescPhone.Validators[0].(func(string) error)
}
