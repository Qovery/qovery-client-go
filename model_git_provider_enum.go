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

// GitProviderEnum the model 'GitProviderEnum'
type GitProviderEnum string

// List of GitProviderEnum
const (
	GITPROVIDERENUM_BITBUCKET GitProviderEnum = "BITBUCKET"
	GITPROVIDERENUM_GITHUB    GitProviderEnum = "GITHUB"
	GITPROVIDERENUM_GITLAB    GitProviderEnum = "GITLAB"
)

// All allowed values of GitProviderEnum enum
var AllowedGitProviderEnumEnumValues = []GitProviderEnum{
	"BITBUCKET",
	"GITHUB",
	"GITLAB",
}

func (v *GitProviderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GitProviderEnum(value)
	for _, existing := range AllowedGitProviderEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GitProviderEnum", value)
}

// NewGitProviderEnumFromValue returns a pointer to a valid GitProviderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGitProviderEnumFromValue(v string) (*GitProviderEnum, error) {
	ev := GitProviderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GitProviderEnum: valid values are %v", v, AllowedGitProviderEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GitProviderEnum) IsValid() bool {
	for _, existing := range AllowedGitProviderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitProviderEnum value
func (v GitProviderEnum) Ptr() *GitProviderEnum {
	return &v
}

type NullableGitProviderEnum struct {
	value *GitProviderEnum
	isSet bool
}

func (v NullableGitProviderEnum) Get() *GitProviderEnum {
	return v.value
}

func (v *NullableGitProviderEnum) Set(val *GitProviderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGitProviderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGitProviderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitProviderEnum(val *GitProviderEnum) *NullableGitProviderEnum {
	return &NullableGitProviderEnum{value: val, isSet: true}
}

func (v NullableGitProviderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitProviderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
