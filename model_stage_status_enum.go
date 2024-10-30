/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"fmt"
)

// StageStatusEnum the model 'StageStatusEnum'
type StageStatusEnum string

// List of StageStatusEnum
const (
	STAGESTATUSENUM_QUEUED   StageStatusEnum = "QUEUED"
	STAGESTATUSENUM_ONGOING  StageStatusEnum = "ONGOING"
	STAGESTATUSENUM_DONE     StageStatusEnum = "DONE"
	STAGESTATUSENUM_ERROR    StageStatusEnum = "ERROR"
	STAGESTATUSENUM_SKIPPED  StageStatusEnum = "SKIPPED"
	STAGESTATUSENUM_CANCELED StageStatusEnum = "CANCELED"
)

// All allowed values of StageStatusEnum enum
var AllowedStageStatusEnumEnumValues = []StageStatusEnum{
	"QUEUED",
	"ONGOING",
	"DONE",
	"ERROR",
	"SKIPPED",
	"CANCELED",
}

func (v *StageStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StageStatusEnum(value)
	for _, existing := range AllowedStageStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StageStatusEnum", value)
}

// NewStageStatusEnumFromValue returns a pointer to a valid StageStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStageStatusEnumFromValue(v string) (*StageStatusEnum, error) {
	ev := StageStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StageStatusEnum: valid values are %v", v, AllowedStageStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StageStatusEnum) IsValid() bool {
	for _, existing := range AllowedStageStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StageStatusEnum value
func (v StageStatusEnum) Ptr() *StageStatusEnum {
	return &v
}

type NullableStageStatusEnum struct {
	value *StageStatusEnum
	isSet bool
}

func (v NullableStageStatusEnum) Get() *StageStatusEnum {
	return v.value
}

func (v *NullableStageStatusEnum) Set(val *StageStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStageStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStageStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageStatusEnum(val *StageStatusEnum) *NullableStageStatusEnum {
	return &NullableStageStatusEnum{value: val, isSet: true}
}

func (v NullableStageStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}