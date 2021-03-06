// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyAppVersionRequest openpitrix modify app version request
// swagger:model openpitrixModifyAppVersionRequest
type OpenpitrixModifyAppVersionRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// home
	Home string `json:"home,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// keywords
	Keywords string `json:"keywords,omitempty"`

	// maintainers
	Maintainers string `json:"maintainers,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// package
	Package strfmt.Base64 `json:"package,omitempty"`

	// package name
	PackageName string `json:"package_name,omitempty"`

	// readme
	Readme string `json:"readme,omitempty"`

	// screenshots
	Screenshots string `json:"screenshots,omitempty"`

	// sequence
	Sequence int64 `json:"sequence,omitempty"`

	// sources
	Sources string `json:"sources,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix modify app version request
func (m *OpenpitrixModifyAppVersionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyAppVersionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyAppVersionRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyAppVersionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
