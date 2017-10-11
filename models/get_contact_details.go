// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetContactDetails get contact details
// swagger:model getContactDetails

type GetContactDetails struct {

	// attributes
	// Required: true
	Attributes map[string]string `json:"attributes"`

	// Email address of the contact for which you requested the details
	// Required: true
	Email *strfmt.Email `json:"email"`

	// Blacklist status for email campaigns (true=blacklisted, false=not blacklisted)
	// Required: true
	EmailBlacklisted *bool `json:"emailBlacklisted"`

	// ID of the contact for which you requested the details
	// Required: true
	ID *int64 `json:"id"`

	// list ids
	// Required: true
	ListIds []int64 `json:"listIds"`

	// list unsubscribed
	ListUnsubscribed []int64 `json:"listUnsubscribed"`

	// Last modification date of the contact (YYYY-MM-DDTHH:mm:ss.SSSZ)
	// Required: true
	ModifiedAt *strfmt.DateTime `json:"modifiedAt"`

	// Blacklist status for SMS campaigns (true=blacklisted, false=not blacklisted)
	// Required: true
	SmsBlacklisted *bool `json:"smsBlacklisted"`
}

/* polymorph getContactDetails attributes false */

/* polymorph getContactDetails email false */

/* polymorph getContactDetails emailBlacklisted false */

/* polymorph getContactDetails id false */

/* polymorph getContactDetails listIds false */

/* polymorph getContactDetails listUnsubscribed false */

/* polymorph getContactDetails modifiedAt false */

/* polymorph getContactDetails smsBlacklisted false */

// Validate validates this get contact details
func (m *GetContactDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmailBlacklisted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateListUnsubscribed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModifiedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSmsBlacklisted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetContactDetails) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	return nil
}

func (m *GetContactDetails) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetContactDetails) validateEmailBlacklisted(formats strfmt.Registry) error {

	if err := validate.Required("emailBlacklisted", "body", m.EmailBlacklisted); err != nil {
		return err
	}

	return nil
}

func (m *GetContactDetails) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetContactDetails) validateListIds(formats strfmt.Registry) error {

	if err := validate.Required("listIds", "body", m.ListIds); err != nil {
		return err
	}

	return nil
}

func (m *GetContactDetails) validateListUnsubscribed(formats strfmt.Registry) error {

	if swag.IsZero(m.ListUnsubscribed) { // not required
		return nil
	}

	return nil
}

func (m *GetContactDetails) validateModifiedAt(formats strfmt.Registry) error {

	if err := validate.Required("modifiedAt", "body", m.ModifiedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("modifiedAt", "body", "date-time", m.ModifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetContactDetails) validateSmsBlacklisted(formats strfmt.Registry) error {

	if err := validate.Required("smsBlacklisted", "body", m.SmsBlacklisted); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetContactDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetContactDetails) UnmarshalBinary(b []byte) error {
	var res GetContactDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
