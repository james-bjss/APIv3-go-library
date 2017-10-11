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

// SendSms send sms
// swagger:model sendSms

type SendSms struct {

	// reference
	// Required: true
	Reference *string `json:"reference"`
}

/* polymorph sendSms reference false */

// Validate validates this send sms
func (m *SendSms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSms) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSms) UnmarshalBinary(b []byte) error {
	var res SendSms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
