// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RequestContactImport request contact import
// swagger:model requestContactImport

type RequestContactImport struct {

	// Mandatory if fileUrl is not defined. CSV content to be imported. Use semicolon to separate multiple attributes
	FileBody string `json:"fileBody,omitempty"`

	// Mandatory if fileBody not defined. URL of the file to be imported (no local file). Possible file types: .txt, .csv
	FileURL string `json:"fileUrl,omitempty"`

	// Manadatory if newList is not defined. Ids of the lists in which to add the contacts
	ListIds []int64 `json:"listIds"`

	// new list
	NewList *RequestContactImportNewList `json:"newList,omitempty"`

	// URL that will be called once the export process is finished
	NotifyURL string `json:"notifyUrl,omitempty"`
}

/* polymorph requestContactImport fileBody false */

/* polymorph requestContactImport fileUrl false */

/* polymorph requestContactImport listIds false */

/* polymorph requestContactImport newList false */

/* polymorph requestContactImport notifyUrl false */

// Validate validates this request contact import
func (m *RequestContactImport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestContactImport) validateListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ListIds) { // not required
		return nil
	}

	return nil
}

func (m *RequestContactImport) validateNewList(formats strfmt.Registry) error {

	if swag.IsZero(m.NewList) { // not required
		return nil
	}

	if m.NewList != nil {

		if err := m.NewList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newList")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestContactImport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestContactImport) UnmarshalBinary(b []byte) error {
	var res RequestContactImport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
