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

// GetReportsReportsItems get reports reports items
// swagger:model getReportsReportsItems

type GetReportsReportsItems struct {

	// Number of blocked emails for the date
	// Required: true
	Blocked *int64 `json:"blocked"`

	// Number of clicks for the date
	// Required: true
	Clicks *int64 `json:"clicks"`

	// Date of the statistics
	// Required: true
	Date *strfmt.Date `json:"date"`

	// Number of delivered emails for the date
	// Required: true
	Delivered *int64 `json:"delivered"`

	// Number of hardbounces for the date
	// Required: true
	HardBounces *int64 `json:"hardBounces"`

	// Number of invalid emails for the date
	// Required: true
	Invalid *int64 `json:"invalid"`

	// Number of openings for the date
	// Required: true
	Opens *int64 `json:"opens"`

	// Number of requests for the date
	// Required: true
	Requests *int64 `json:"requests"`

	// Number of softbounces for the date
	// Required: true
	SoftBounces *int64 `json:"softBounces"`

	// Number of complaints (spam reports) for the date
	// Required: true
	SpamReports *int64 `json:"spamReports"`

	// Reminder of the specified tag (only available if a specific tag has been specified in the request)
	// Required: true
	Tag *string `json:"tag"`

	// Number of unique clicks for the date
	// Required: true
	UniqueClicks *int64 `json:"uniqueClicks"`

	// Number of unique openings for the date
	// Required: true
	UniqueOpens *int64 `json:"uniqueOpens"`
}

/* polymorph getReportsReportsItems blocked false */

/* polymorph getReportsReportsItems clicks false */

/* polymorph getReportsReportsItems date false */

/* polymorph getReportsReportsItems delivered false */

/* polymorph getReportsReportsItems hardBounces false */

/* polymorph getReportsReportsItems invalid false */

/* polymorph getReportsReportsItems opens false */

/* polymorph getReportsReportsItems requests false */

/* polymorph getReportsReportsItems softBounces false */

/* polymorph getReportsReportsItems spamReports false */

/* polymorph getReportsReportsItems tag false */

/* polymorph getReportsReportsItems uniqueClicks false */

/* polymorph getReportsReportsItems uniqueOpens false */

// Validate validates this get reports reports items
func (m *GetReportsReportsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlocked(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClicks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDelivered(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHardBounces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvalid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOpens(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSoftBounces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpamReports(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUniqueClicks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUniqueOpens(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetReportsReportsItems) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateClicks(formats strfmt.Registry) error {

	if err := validate.Required("clicks", "body", m.Clicks); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateDelivered(formats strfmt.Registry) error {

	if err := validate.Required("delivered", "body", m.Delivered); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateHardBounces(formats strfmt.Registry) error {

	if err := validate.Required("hardBounces", "body", m.HardBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateInvalid(formats strfmt.Registry) error {

	if err := validate.Required("invalid", "body", m.Invalid); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateOpens(formats strfmt.Registry) error {

	if err := validate.Required("opens", "body", m.Opens); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateRequests(formats strfmt.Registry) error {

	if err := validate.Required("requests", "body", m.Requests); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateSoftBounces(formats strfmt.Registry) error {

	if err := validate.Required("softBounces", "body", m.SoftBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateSpamReports(formats strfmt.Registry) error {

	if err := validate.Required("spamReports", "body", m.SpamReports); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateUniqueClicks(formats strfmt.Registry) error {

	if err := validate.Required("uniqueClicks", "body", m.UniqueClicks); err != nil {
		return err
	}

	return nil
}

func (m *GetReportsReportsItems) validateUniqueOpens(formats strfmt.Registry) error {

	if err := validate.Required("uniqueOpens", "body", m.UniqueOpens); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetReportsReportsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetReportsReportsItems) UnmarshalBinary(b []byte) error {
	var res GetReportsReportsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
