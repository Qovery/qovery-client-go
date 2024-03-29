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

// OrganizationEventType Type of the organization event
type OrganizationEventType string

// List of OrganizationEventType
const (
	ORGANIZATIONEVENTTYPE_CREATE           OrganizationEventType = "CREATE"
	ORGANIZATIONEVENTTYPE_UPDATE           OrganizationEventType = "UPDATE"
	ORGANIZATIONEVENTTYPE_DELETE           OrganizationEventType = "DELETE"
	ORGANIZATIONEVENTTYPE_ACCEPT           OrganizationEventType = "ACCEPT"
	ORGANIZATIONEVENTTYPE_EXPORT           OrganizationEventType = "EXPORT"
	ORGANIZATIONEVENTTYPE_TRIGGER_DEPLOY   OrganizationEventType = "TRIGGER_DEPLOY"
	ORGANIZATIONEVENTTYPE_TRIGGER_REDEPLOY OrganizationEventType = "TRIGGER_REDEPLOY"
	ORGANIZATIONEVENTTYPE_TRIGGER_STOP     OrganizationEventType = "TRIGGER_STOP"
	ORGANIZATIONEVENTTYPE_TRIGGER_CANCEL   OrganizationEventType = "TRIGGER_CANCEL"
	ORGANIZATIONEVENTTYPE_TRIGGER_RESTART  OrganizationEventType = "TRIGGER_RESTART"
	ORGANIZATIONEVENTTYPE_TRIGGER_DELETE   OrganizationEventType = "TRIGGER_DELETE"
)

// All allowed values of OrganizationEventType enum
var AllowedOrganizationEventTypeEnumValues = []OrganizationEventType{
	"CREATE",
	"UPDATE",
	"DELETE",
	"ACCEPT",
	"EXPORT",
	"TRIGGER_DEPLOY",
	"TRIGGER_REDEPLOY",
	"TRIGGER_STOP",
	"TRIGGER_CANCEL",
	"TRIGGER_RESTART",
	"TRIGGER_DELETE",
}

func (v *OrganizationEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationEventType(value)
	for _, existing := range AllowedOrganizationEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationEventType", value)
}

// NewOrganizationEventTypeFromValue returns a pointer to a valid OrganizationEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationEventTypeFromValue(v string) (*OrganizationEventType, error) {
	ev := OrganizationEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationEventType: valid values are %v", v, AllowedOrganizationEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationEventType) IsValid() bool {
	for _, existing := range AllowedOrganizationEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationEventType value
func (v OrganizationEventType) Ptr() *OrganizationEventType {
	return &v
}

type NullableOrganizationEventType struct {
	value *OrganizationEventType
	isSet bool
}

func (v NullableOrganizationEventType) Get() *OrganizationEventType {
	return v.value
}

func (v *NullableOrganizationEventType) Set(val *OrganizationEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventType(val *OrganizationEventType) *NullableOrganizationEventType {
	return &NullableOrganizationEventType{value: val, isSet: true}
}

func (v NullableOrganizationEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
