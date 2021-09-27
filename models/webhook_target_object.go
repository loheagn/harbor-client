// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhookTargetObject The webhook policy target object.
//
// swagger:model WebhookTargetObject
type WebhookTargetObject struct {

	// The webhook target address.
	Address string `json:"address,omitempty"`

	// The webhook auth header.
	AuthHeader string `json:"auth_header,omitempty"`

	// Whether or not to skip cert verify.
	SkipCertVerify bool `json:"skip_cert_verify,omitempty"`

	// The webhook target notify type.
	Type string `json:"type,omitempty"`
}

// Validate validates this webhook target object
func (m *WebhookTargetObject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webhook target object based on context it is used
func (m *WebhookTargetObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookTargetObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookTargetObject) UnmarshalBinary(b []byte) error {
	var res WebhookTargetObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
