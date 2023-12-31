// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusNode status node
//
// swagger:model statusNode
type StatusNode struct {

	// done count
	// Required: true
	DoneCount *int64 `json:"DoneCount"`

	// error
	// Required: true
	Error *string `json:"Error"`

	// finished at
	// Required: true
	FinishedAt *string `json:"FinishedAt"`

	// log
	// Required: true
	Log *string `json:"Log"`

	// retry count
	// Required: true
	RetryCount *int64 `json:"RetryCount"`

	// started at
	// Required: true
	StartedAt *string `json:"StartedAt"`

	// status
	// Required: true
	Status *int64 `json:"Status"`

	// status text
	// Required: true
	StatusText *string `json:"StatusText"`

	// step
	// Required: true
	Step *StepObject `json:"Step"`
}

// Validate validates this status node
func (m *StatusNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDoneCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetryCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStep(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusNode) validateDoneCount(formats strfmt.Registry) error {

	if err := validate.Required("DoneCount", "body", m.DoneCount); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateError(formats strfmt.Registry) error {

	if err := validate.Required("Error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateFinishedAt(formats strfmt.Registry) error {

	if err := validate.Required("FinishedAt", "body", m.FinishedAt); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("Log", "body", m.Log); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateRetryCount(formats strfmt.Registry) error {

	if err := validate.Required("RetryCount", "body", m.RetryCount); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("StartedAt", "body", m.StartedAt); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateStatusText(formats strfmt.Registry) error {

	if err := validate.Required("StatusText", "body", m.StatusText); err != nil {
		return err
	}

	return nil
}

func (m *StatusNode) validateStep(formats strfmt.Registry) error {

	if err := validate.Required("Step", "body", m.Step); err != nil {
		return err
	}

	if m.Step != nil {
		if err := m.Step.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Step")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Step")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status node based on the context it is used
func (m *StatusNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusNode) contextValidateStep(ctx context.Context, formats strfmt.Registry) error {

	if m.Step != nil {

		if err := m.Step.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Step")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Step")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusNode) UnmarshalBinary(b []byte) error {
	var res StatusNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
