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

// WeekdayEnum the model 'WeekdayEnum'
type WeekdayEnum string

// List of WeekdayEnum
const (
	MONDAY    WeekdayEnum = "MONDAY"
	TUESDAY   WeekdayEnum = "TUESDAY"
	WEDNESDAY WeekdayEnum = "WEDNESDAY"
	THURSDAY  WeekdayEnum = "THURSDAY"
	FRIDAY    WeekdayEnum = "FRIDAY"
	SATURDAY  WeekdayEnum = "SATURDAY"
	SUNDAY    WeekdayEnum = "SUNDAY"
)

// All allowed values of WeekdayEnum enum
var AllowedWeekdayEnumEnumValues = []WeekdayEnum{
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"SUNDAY",
}

func (v *WeekdayEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WeekdayEnum(value)
	for _, existing := range AllowedWeekdayEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WeekdayEnum", value)
}

// NewWeekdayEnumFromValue returns a pointer to a valid WeekdayEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWeekdayEnumFromValue(v string) (*WeekdayEnum, error) {
	ev := WeekdayEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WeekdayEnum: valid values are %v", v, AllowedWeekdayEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WeekdayEnum) IsValid() bool {
	for _, existing := range AllowedWeekdayEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WeekdayEnum value
func (v WeekdayEnum) Ptr() *WeekdayEnum {
	return &v
}

type NullableWeekdayEnum struct {
	value *WeekdayEnum
	isSet bool
}

func (v NullableWeekdayEnum) Get() *WeekdayEnum {
	return v.value
}

func (v *NullableWeekdayEnum) Set(val *WeekdayEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWeekdayEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWeekdayEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeekdayEnum(val *WeekdayEnum) *NullableWeekdayEnum {
	return &NullableWeekdayEnum{value: val, isSet: true}
}

func (v NullableWeekdayEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeekdayEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}