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

// JobLifecycleTypeEnum the model 'JobLifecycleTypeEnum'
type JobLifecycleTypeEnum string

// List of JobLifecycleTypeEnum
const (
	JOBLIFECYCLETYPEENUM_GENERIC        JobLifecycleTypeEnum = "GENERIC"
	JOBLIFECYCLETYPEENUM_TERRAFORM      JobLifecycleTypeEnum = "TERRAFORM"
	JOBLIFECYCLETYPEENUM_CLOUDFORMATION JobLifecycleTypeEnum = "CLOUDFORMATION"
)

// All allowed values of JobLifecycleTypeEnum enum
var AllowedJobLifecycleTypeEnumEnumValues = []JobLifecycleTypeEnum{
	"GENERIC",
	"TERRAFORM",
	"CLOUDFORMATION",
}

func (v *JobLifecycleTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobLifecycleTypeEnum(value)
	for _, existing := range AllowedJobLifecycleTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobLifecycleTypeEnum", value)
}

// NewJobLifecycleTypeEnumFromValue returns a pointer to a valid JobLifecycleTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobLifecycleTypeEnumFromValue(v string) (*JobLifecycleTypeEnum, error) {
	ev := JobLifecycleTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobLifecycleTypeEnum: valid values are %v", v, AllowedJobLifecycleTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobLifecycleTypeEnum) IsValid() bool {
	for _, existing := range AllowedJobLifecycleTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobLifecycleTypeEnum value
func (v JobLifecycleTypeEnum) Ptr() *JobLifecycleTypeEnum {
	return &v
}

type NullableJobLifecycleTypeEnum struct {
	value *JobLifecycleTypeEnum
	isSet bool
}

func (v NullableJobLifecycleTypeEnum) Get() *JobLifecycleTypeEnum {
	return v.value
}

func (v *NullableJobLifecycleTypeEnum) Set(val *JobLifecycleTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableJobLifecycleTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableJobLifecycleTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobLifecycleTypeEnum(val *JobLifecycleTypeEnum) *NullableJobLifecycleTypeEnum {
	return &NullableJobLifecycleTypeEnum{value: val, isSet: true}
}

func (v NullableJobLifecycleTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobLifecycleTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}