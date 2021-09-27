// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/goharbor/harbor/src/pkg/quota/types"
)

// ProjectSummaryQuota project summary quota
//
// swagger:model ProjectSummaryQuota
type ProjectSummaryQuota struct {

	// The hard limits of the quota
	Hard types.ResourceList `json:"hard,omitempty"`

	// The used status of the quota
	Used types.ResourceList `json:"used,omitempty"`
}

// Validate validates this project summary quota
func (m *ProjectSummaryQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectSummaryQuota) validateHard(formats strfmt.Registry) error {
	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	if err := m.Hard.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hard")
		}
		return err
	}

	return nil
}

func (m *ProjectSummaryQuota) validateUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if err := m.Used.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("used")
		}
		return err
	}

	return nil
}

// ContextValidate validate this project summary quota based on the context it is used
func (m *ProjectSummaryQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectSummaryQuota) contextValidateHard(ctx context.Context, formats strfmt.Registry) error {

	// if err := m.Hard.ContextValidate(ctx, formats); err != nil {
	// 	if ve, ok := err.(*errors.Validation); ok {
	// 		return ve.ValidateName("hard")
	// 	}
	// 	return err
	// }

	return nil
}

func (m *ProjectSummaryQuota) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	// if err := m.Used.ContextValidate(ctx, formats); err != nil {
	// 	if ve, ok := err.(*errors.Validation); ok {
	// 		return ve.ValidateName("used")
	// 	}
	// 	return err
	// }

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectSummaryQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectSummaryQuota) UnmarshalBinary(b []byte) error {
	var res ProjectSummaryQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
