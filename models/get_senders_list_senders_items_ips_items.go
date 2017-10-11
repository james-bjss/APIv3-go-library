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

// GetSendersListSendersItemsIpsItems get senders list senders items ips items
// swagger:model getSendersListSendersItemsIpsItems

type GetSendersListSendersItemsIpsItems struct {

	// Domain of the IP
	// Required: true
	Domain *string `json:"domain"`

	// Dedicated IP available in your account
	// Required: true
	IP *string `json:"ip"`

	// Weight of the IP for this sender
	// Required: true
	Weight *int64 `json:"weight"`
}

/* polymorph getSendersListSendersItemsIpsItems domain false */

/* polymorph getSendersListSendersItemsIpsItems ip false */

/* polymorph getSendersListSendersItemsIpsItems weight false */

// Validate validates this get senders list senders items ips items
func (m *GetSendersListSendersItemsIpsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSendersListSendersItemsIpsItems) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *GetSendersListSendersItemsIpsItems) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *GetSendersListSendersItemsIpsItems) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSendersListSendersItemsIpsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSendersListSendersItemsIpsItems) UnmarshalBinary(b []byte) error {
	var res GetSendersListSendersItemsIpsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
