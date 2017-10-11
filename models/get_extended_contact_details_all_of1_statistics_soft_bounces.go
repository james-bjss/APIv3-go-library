// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetExtendedContactDetailsAllOf1StatisticsSoftBounces Listing of the softbounes generated by the contact
// swagger:model getExtendedContactDetailsAllOf1StatisticsSoftBounces

type GetExtendedContactDetailsAllOf1StatisticsSoftBounces []*GetExtendedContactDetailsAllOf1StatisticsSoftBouncesItems

// Validate validates this get extended contact details all of1 statistics soft bounces
func (m GetExtendedContactDetailsAllOf1StatisticsSoftBounces) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
