// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProxyStatus Status of proxy
// swagger:model ProxyStatus
type ProxyStatus struct {

	// IP address that the proxy listens on
	IP string `json:"ip,omitempty"`

	// Port range used for proxying
	PortRange string `json:"port-range,omitempty"`
}

// Validate validates this proxy status
func (m *ProxyStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProxyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyStatus) UnmarshalBinary(b []byte) error {
	var res ProxyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
