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

// OrganizationWebhookKindEnum Define the type of the webhook. `SLACK` is a special webhook type to push notifications directly to slack. The `target_url` must be a Slack compatible endpoint.
type OrganizationWebhookKindEnum string

// List of OrganizationWebhookKindEnum
const (
	ORGANIZATIONWEBHOOKKINDENUM_STANDARD OrganizationWebhookKindEnum = "STANDARD"
	ORGANIZATIONWEBHOOKKINDENUM_SLACK    OrganizationWebhookKindEnum = "SLACK"
)

// All allowed values of OrganizationWebhookKindEnum enum
var AllowedOrganizationWebhookKindEnumEnumValues = []OrganizationWebhookKindEnum{
	"STANDARD",
	"SLACK",
}

func (v *OrganizationWebhookKindEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationWebhookKindEnum(value)
	for _, existing := range AllowedOrganizationWebhookKindEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationWebhookKindEnum", value)
}

// NewOrganizationWebhookKindEnumFromValue returns a pointer to a valid OrganizationWebhookKindEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationWebhookKindEnumFromValue(v string) (*OrganizationWebhookKindEnum, error) {
	ev := OrganizationWebhookKindEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationWebhookKindEnum: valid values are %v", v, AllowedOrganizationWebhookKindEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationWebhookKindEnum) IsValid() bool {
	for _, existing := range AllowedOrganizationWebhookKindEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationWebhookKindEnum value
func (v OrganizationWebhookKindEnum) Ptr() *OrganizationWebhookKindEnum {
	return &v
}

type NullableOrganizationWebhookKindEnum struct {
	value *OrganizationWebhookKindEnum
	isSet bool
}

func (v NullableOrganizationWebhookKindEnum) Get() *OrganizationWebhookKindEnum {
	return v.value
}

func (v *NullableOrganizationWebhookKindEnum) Set(val *OrganizationWebhookKindEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWebhookKindEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWebhookKindEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWebhookKindEnum(val *OrganizationWebhookKindEnum) *NullableOrganizationWebhookKindEnum {
	return &NullableOrganizationWebhookKindEnum{value: val, isSet: true}
}

func (v NullableOrganizationWebhookKindEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWebhookKindEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}