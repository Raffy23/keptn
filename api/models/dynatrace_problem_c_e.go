// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DynatraceProblemCE dynatrace problem c e
// swagger:model DynatraceProblemCE
type DynatraceProblemCE struct {
	CEWithoutDataWithKeptncontext

	// data
	Data *DynatraceProblemCEAO1Data `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DynatraceProblemCE) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CEWithoutDataWithKeptncontext
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CEWithoutDataWithKeptncontext = aO0

	// AO1
	var dataAO1 struct {
		Data *DynatraceProblemCEAO1Data `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Data = dataAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DynatraceProblemCE) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CEWithoutDataWithKeptncontext)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Data *DynatraceProblemCEAO1Data `json:"data,omitempty"`
	}

	dataAO1.Data = m.Data

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this dynatrace problem c e
func (m *DynatraceProblemCE) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CEWithoutDataWithKeptncontext
	if err := m.CEWithoutDataWithKeptncontext.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynatraceProblemCE) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DynatraceProblemCE) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynatraceProblemCE) UnmarshalBinary(b []byte) error {
	var res DynatraceProblemCE
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DynatraceProblemCEAO1Data dynatrace problem c e a o1 data
// swagger:model DynatraceProblemCEAO1Data
type DynatraceProblemCEAO1Data struct {

	// impactedentities
	// Required: true
	Impactedentities []*DynatraceProblemCEAO1DataImpactedentitiesItems0 `json:"impactedentities"`

	// impactedentity
	// Required: true
	Impactedentity *string `json:"impactedentity"`

	// pid
	// Required: true
	Pid *string `json:"pid"`

	// problemdetails
	// Required: true
	Problemdetails *DynatraceProblemCEAO1DataProblemdetails `json:"problemdetails"`

	// problemid
	// Required: true
	Problemid *string `json:"problemid"`

	// problemtitle
	// Required: true
	Problemtitle *string `json:"problemtitle"`

	// state
	// Required: true
	State *string `json:"state"`
}

// Validate validates this dynatrace problem c e a o1 data
func (m *DynatraceProblemCEAO1Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImpactedentities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpactedentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProblemdetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProblemid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProblemtitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynatraceProblemCEAO1Data) validateImpactedentities(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"impactedentities", "body", m.Impactedentities); err != nil {
		return err
	}

	for i := 0; i < len(m.Impactedentities); i++ {
		if swag.IsZero(m.Impactedentities[i]) { // not required
			continue
		}

		if m.Impactedentities[i] != nil {
			if err := m.Impactedentities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "impactedentities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DynatraceProblemCEAO1Data) validateImpactedentity(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"impactedentity", "body", m.Impactedentity); err != nil {
		return err
	}

	return nil
}

func (m *DynatraceProblemCEAO1Data) validatePid(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"pid", "body", m.Pid); err != nil {
		return err
	}

	return nil
}

func (m *DynatraceProblemCEAO1Data) validateProblemdetails(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"problemdetails", "body", m.Problemdetails); err != nil {
		return err
	}

	if m.Problemdetails != nil {
		if err := m.Problemdetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "problemdetails")
			}
			return err
		}
	}

	return nil
}

func (m *DynatraceProblemCEAO1Data) validateProblemid(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"problemid", "body", m.Problemid); err != nil {
		return err
	}

	return nil
}

func (m *DynatraceProblemCEAO1Data) validateProblemtitle(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"problemtitle", "body", m.Problemtitle); err != nil {
		return err
	}

	return nil
}

func (m *DynatraceProblemCEAO1Data) validateState(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DynatraceProblemCEAO1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynatraceProblemCEAO1Data) UnmarshalBinary(b []byte) error {
	var res DynatraceProblemCEAO1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DynatraceProblemCEAO1DataImpactedentitiesItems0 dynatrace problem c e a o1 data impactedentities items0
// swagger:model DynatraceProblemCEAO1DataImpactedentitiesItems0
type DynatraceProblemCEAO1DataImpactedentitiesItems0 struct {

	// entity
	// Required: true
	Entity *string `json:"entity"`

	// name
	// Required: true
	Name *string `json:"name"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this dynatrace problem c e a o1 data impactedentities items0
func (m *DynatraceProblemCEAO1DataImpactedentitiesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynatraceProblemCEAO1DataImpactedentitiesItems0) validateEntity(formats strfmt.Registry) error {

	if err := validate.Required("entity", "body", m.Entity); err != nil {
		return err
	}

	return nil
}

func (m *DynatraceProblemCEAO1DataImpactedentitiesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DynatraceProblemCEAO1DataImpactedentitiesItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DynatraceProblemCEAO1DataImpactedentitiesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynatraceProblemCEAO1DataImpactedentitiesItems0) UnmarshalBinary(b []byte) error {
	var res DynatraceProblemCEAO1DataImpactedentitiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DynatraceProblemCEAO1DataProblemdetails dynatrace problem c e a o1 data problemdetails
// swagger:model DynatraceProblemCEAO1DataProblemdetails
type DynatraceProblemCEAO1DataProblemdetails struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this dynatrace problem c e a o1 data problemdetails
func (m *DynatraceProblemCEAO1DataProblemdetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynatraceProblemCEAO1DataProblemdetails) validateID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"problemdetails"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DynatraceProblemCEAO1DataProblemdetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynatraceProblemCEAO1DataProblemdetails) UnmarshalBinary(b []byte) error {
	var res DynatraceProblemCEAO1DataProblemdetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
