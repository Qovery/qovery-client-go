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

// GitTokenAssociatedServiceType the model 'GitTokenAssociatedServiceType'
type GitTokenAssociatedServiceType string

// List of GitTokenAssociatedServiceType
const (
	GITTOKENASSOCIATEDSERVICETYPE_APPLICATION GitTokenAssociatedServiceType = "APPLICATION"
	GITTOKENASSOCIATEDSERVICETYPE_JOB         GitTokenAssociatedServiceType = "JOB"
	GITTOKENASSOCIATEDSERVICETYPE_HELM        GitTokenAssociatedServiceType = "HELM"
)

// All allowed values of GitTokenAssociatedServiceType enum
var AllowedGitTokenAssociatedServiceTypeEnumValues = []GitTokenAssociatedServiceType{
	"APPLICATION",
	"JOB",
	"HELM",
}

func (v *GitTokenAssociatedServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GitTokenAssociatedServiceType(value)
	for _, existing := range AllowedGitTokenAssociatedServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GitTokenAssociatedServiceType", value)
}

// NewGitTokenAssociatedServiceTypeFromValue returns a pointer to a valid GitTokenAssociatedServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGitTokenAssociatedServiceTypeFromValue(v string) (*GitTokenAssociatedServiceType, error) {
	ev := GitTokenAssociatedServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GitTokenAssociatedServiceType: valid values are %v", v, AllowedGitTokenAssociatedServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GitTokenAssociatedServiceType) IsValid() bool {
	for _, existing := range AllowedGitTokenAssociatedServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitTokenAssociatedServiceType value
func (v GitTokenAssociatedServiceType) Ptr() *GitTokenAssociatedServiceType {
	return &v
}

type NullableGitTokenAssociatedServiceType struct {
	value *GitTokenAssociatedServiceType
	isSet bool
}

func (v NullableGitTokenAssociatedServiceType) Get() *GitTokenAssociatedServiceType {
	return v.value
}

func (v *NullableGitTokenAssociatedServiceType) Set(val *GitTokenAssociatedServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTokenAssociatedServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTokenAssociatedServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTokenAssociatedServiceType(val *GitTokenAssociatedServiceType) *NullableGitTokenAssociatedServiceType {
	return &NullableGitTokenAssociatedServiceType{value: val, isSet: true}
}

func (v NullableGitTokenAssociatedServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTokenAssociatedServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}