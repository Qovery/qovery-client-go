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

// DatabaseModeEnum the model 'DatabaseModeEnum'
type DatabaseModeEnum string

// List of DatabaseModeEnum
const (
	MANAGED   DatabaseModeEnum = "MANAGED"
	CONTAINER DatabaseModeEnum = "CONTAINER"
)

// All allowed values of DatabaseModeEnum enum
var AllowedDatabaseModeEnumEnumValues = []DatabaseModeEnum{
	"MANAGED",
	"CONTAINER",
}

func (v *DatabaseModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatabaseModeEnum(value)
	for _, existing := range AllowedDatabaseModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatabaseModeEnum", value)
}

// NewDatabaseModeEnumFromValue returns a pointer to a valid DatabaseModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatabaseModeEnumFromValue(v string) (*DatabaseModeEnum, error) {
	ev := DatabaseModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatabaseModeEnum: valid values are %v", v, AllowedDatabaseModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatabaseModeEnum) IsValid() bool {
	for _, existing := range AllowedDatabaseModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseModeEnum value
func (v DatabaseModeEnum) Ptr() *DatabaseModeEnum {
	return &v
}

type NullableDatabaseModeEnum struct {
	value *DatabaseModeEnum
	isSet bool
}

func (v NullableDatabaseModeEnum) Get() *DatabaseModeEnum {
	return v.value
}

func (v *NullableDatabaseModeEnum) Set(val *DatabaseModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseModeEnum(val *DatabaseModeEnum) *NullableDatabaseModeEnum {
	return &NullableDatabaseModeEnum{value: val, isSet: true}
}

func (v NullableDatabaseModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}