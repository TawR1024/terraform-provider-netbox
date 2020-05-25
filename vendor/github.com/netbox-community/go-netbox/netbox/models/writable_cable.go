// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCable writable cable
// swagger:model WritableCable
type WritableCable struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	// Max Length: 100
	Label string `json:"label,omitempty"`

	// Length
	// Maximum: 32767
	// Minimum: 0
	Length *int64 `json:"length,omitempty"`

	// Length unit
	// Enum: [1200 1100 2100 2000]
	LengthUnit *int64 `json:"length_unit,omitempty"`

	// Status
	// Enum: [false true]
	Status bool `json:"status,omitempty"`

	// Termination a
	// Read Only: true
	TerminationA map[string]string `json:"termination_a,omitempty"`

	// Termination a id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationAID *int64 `json:"termination_a_id"`

	// Termination a type
	// Required: true
	TerminationAType *string `json:"termination_a_type"`

	// Termination b
	// Read Only: true
	TerminationB map[string]string `json:"termination_b,omitempty"`

	// Termination b id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationBID *int64 `json:"termination_b_id"`

	// Termination b type
	// Required: true
	TerminationBType *string `json:"termination_b_type"`

	// Type
	// Enum: [1300 1500 1510 1600 1610 1700 1800 1810 3000 3010 3020 3030 3040 3500 3510 3520 3800 5000]
	Type *int64 `json:"type,omitempty"`
}

// Validate validates this writable cable
func (m *WritableCable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationAID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationAType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationBID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationBType(formats); err != nil {
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

func (m *WritableCable) validateColor(formats strfmt.Registry) error {

	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", string(m.Color), 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", string(m.Color), `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateLength(formats strfmt.Registry) error {

	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if err := validate.MinimumInt("length", "body", int64(*m.Length), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "body", int64(*m.Length), 32767, false); err != nil {
		return err
	}

	return nil
}

var writableCableTypeLengthUnitPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1200,1100,2100,2000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeLengthUnitPropEnum = append(writableCableTypeLengthUnitPropEnum, v)
	}
}

// prop value enum
func (m *WritableCable) validateLengthUnitEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableCableTypeLengthUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateLengthUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.LengthUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateLengthUnitEnum("length_unit", "body", *m.LengthUnit); err != nil {
		return err
	}

	return nil
}

var writableCableTypeStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeStatusPropEnum = append(writableCableTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableCable) validateStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writableCableTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationAID(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_id", "body", m.TerminationAID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_a_id", "body", int64(*m.TerminationAID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_a_id", "body", int64(*m.TerminationAID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationAType(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_type", "body", m.TerminationAType); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationBID(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_id", "body", m.TerminationBID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_b_id", "body", int64(*m.TerminationBID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_b_id", "body", int64(*m.TerminationBID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationBType(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_type", "body", m.TerminationBType); err != nil {
		return err
	}

	return nil
}

var writableCableTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1300,1500,1510,1600,1610,1700,1800,1810,3000,3010,3020,3030,3040,3500,3510,3520,3800,5000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeTypePropEnum = append(writableCableTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *WritableCable) validateTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableCableTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCable) UnmarshalBinary(b []byte) error {
	var res WritableCable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}