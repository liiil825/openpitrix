// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateAppVersionResponse openpitrix create app version response
// swagger:model openpitrixCreateAppVersionResponse
type OpenpitrixCreateAppVersionResponse struct {

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix create app version response
func (m *OpenpitrixCreateAppVersionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateAppVersionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateAppVersionResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateAppVersionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
