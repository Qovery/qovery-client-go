/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.4
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"fmt"
)

// EnvironmentDeploymentStatusEnum the model 'EnvironmentDeploymentStatusEnum'
type EnvironmentDeploymentStatusEnum string

// List of EnvironmentDeploymentStatusEnum
const (
	ENVIRONMENTDEPLOYMENTSTATUSENUM_NEVER_DEPLOYED EnvironmentDeploymentStatusEnum = "NEVER_DEPLOYED"
	ENVIRONMENTDEPLOYMENTSTATUSENUM_UP_TO_DATE     EnvironmentDeploymentStatusEnum = "UP_TO_DATE"
	ENVIRONMENTDEPLOYMENTSTATUSENUM_OUT_OF_DATE    EnvironmentDeploymentStatusEnum = "OUT_OF_DATE"
)

// All allowed values of EnvironmentDeploymentStatusEnum enum
var AllowedEnvironmentDeploymentStatusEnumEnumValues = []EnvironmentDeploymentStatusEnum{
	"NEVER_DEPLOYED",
	"UP_TO_DATE",
	"OUT_OF_DATE",
}

func (v *EnvironmentDeploymentStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnvironmentDeploymentStatusEnum(value)
	for _, existing := range AllowedEnvironmentDeploymentStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnvironmentDeploymentStatusEnum", value)
}

// NewEnvironmentDeploymentStatusEnumFromValue returns a pointer to a valid EnvironmentDeploymentStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnvironmentDeploymentStatusEnumFromValue(v string) (*EnvironmentDeploymentStatusEnum, error) {
	ev := EnvironmentDeploymentStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnvironmentDeploymentStatusEnum: valid values are %v", v, AllowedEnvironmentDeploymentStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnvironmentDeploymentStatusEnum) IsValid() bool {
	for _, existing := range AllowedEnvironmentDeploymentStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentDeploymentStatusEnum value
func (v EnvironmentDeploymentStatusEnum) Ptr() *EnvironmentDeploymentStatusEnum {
	return &v
}

type NullableEnvironmentDeploymentStatusEnum struct {
	value *EnvironmentDeploymentStatusEnum
	isSet bool
}

func (v NullableEnvironmentDeploymentStatusEnum) Get() *EnvironmentDeploymentStatusEnum {
	return v.value
}

func (v *NullableEnvironmentDeploymentStatusEnum) Set(val *EnvironmentDeploymentStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDeploymentStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDeploymentStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDeploymentStatusEnum(val *EnvironmentDeploymentStatusEnum) *NullableEnvironmentDeploymentStatusEnum {
	return &NullableEnvironmentDeploymentStatusEnum{value: val, isSet: true}
}

func (v NullableEnvironmentDeploymentStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDeploymentStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
