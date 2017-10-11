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

// GetExtendedListAllOf1CampaignStatsItems get extended list all of1 campaign stats items
// swagger:model getExtendedListAllOf1CampaignStatsItems

type GetExtendedListAllOf1CampaignStatsItems struct {

	// ID of the campaign
	// Required: true
	CampaignID *int64 `json:"campaignId"`

	// stats
	// Required: true
	Stats *GetCampaignStats `json:"stats"`
}

/* polymorph getExtendedListAllOf1CampaignStatsItems campaignId false */

/* polymorph getExtendedListAllOf1CampaignStatsItems stats false */

// Validate validates this get extended list all of1 campaign stats items
func (m *GetExtendedListAllOf1CampaignStatsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetExtendedListAllOf1CampaignStatsItems) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaignId", "body", m.CampaignID); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedListAllOf1CampaignStatsItems) validateStats(formats strfmt.Registry) error {

	if err := validate.Required("stats", "body", m.Stats); err != nil {
		return err
	}

	if m.Stats != nil {

		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetExtendedListAllOf1CampaignStatsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetExtendedListAllOf1CampaignStatsItems) UnmarshalBinary(b []byte) error {
	var res GetExtendedListAllOf1CampaignStatsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
