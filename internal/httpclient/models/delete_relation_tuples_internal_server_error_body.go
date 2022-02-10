// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteRelationTuplesInternalServerErrorBody DeleteRelationTuplesInternalServerErrorBody DeleteRelationTuplesInternalServerErrorBody DeleteRelationTuplesInternalServerErrorBody DeleteRelationTuplesInternalServerErrorBody delete relation tuples internal server error body
//
// swagger:model DeleteRelationTuplesInternalServerErrorBody
type DeleteRelationTuplesInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this delete relation tuples internal server error body
func (m *DeleteRelationTuplesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete relation tuples internal server error body based on context it is used
func (m *DeleteRelationTuplesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteRelationTuplesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteRelationTuplesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteRelationTuplesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
