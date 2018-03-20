// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRepoTask openpitrix repo task
// swagger:model openpitrixRepoTask
type OpenpitrixRepoTask struct {

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// repo id
	RepoID string `json:"repo_id,omitempty"`

	// repo task id
	RepoTaskID string `json:"repo_task_id,omitempty"`

	// result
	Result string `json:"result,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`
}

// Validate validates this openpitrix repo task
func (m *OpenpitrixRepoTask) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRepoTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRepoTask) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRepoTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}