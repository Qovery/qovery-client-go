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

// OrganizationEventTargetType Type of the organization event
type OrganizationEventTargetType string

// List of OrganizationEventTargetType
const (
	ORGANIZATIONEVENTTARGETTYPE_APPLICATION        OrganizationEventTargetType = "APPLICATION"
	ORGANIZATIONEVENTTARGETTYPE_CLUSTER            OrganizationEventTargetType = "CLUSTER"
	ORGANIZATIONEVENTTARGETTYPE_CONTAINER          OrganizationEventTargetType = "CONTAINER"
	ORGANIZATIONEVENTTARGETTYPE_CONTAINER_REGISTRY OrganizationEventTargetType = "CONTAINER_REGISTRY"
	ORGANIZATIONEVENTTARGETTYPE_DATABASE           OrganizationEventTargetType = "DATABASE"
	ORGANIZATIONEVENTTARGETTYPE_ENVIRONMENT        OrganizationEventTargetType = "ENVIRONMENT"
	ORGANIZATIONEVENTTARGETTYPE_JOB                OrganizationEventTargetType = "JOB"
	ORGANIZATIONEVENTTARGETTYPE_MEMBERS_AND_ROLES  OrganizationEventTargetType = "MEMBERS_AND_ROLES"
	ORGANIZATIONEVENTTARGETTYPE_ORGANIZATION       OrganizationEventTargetType = "ORGANIZATION"
	ORGANIZATIONEVENTTARGETTYPE_PROJECT            OrganizationEventTargetType = "PROJECT"
	ORGANIZATIONEVENTTARGETTYPE_WEBHOOK            OrganizationEventTargetType = "WEBHOOK"
)

// All allowed values of OrganizationEventTargetType enum
var AllowedOrganizationEventTargetTypeEnumValues = []OrganizationEventTargetType{
	"APPLICATION",
	"CLUSTER",
	"CONTAINER",
	"CONTAINER_REGISTRY",
	"DATABASE",
	"ENVIRONMENT",
	"JOB",
	"MEMBERS_AND_ROLES",
	"ORGANIZATION",
	"PROJECT",
	"WEBHOOK",
}

func (v *OrganizationEventTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationEventTargetType(value)
	for _, existing := range AllowedOrganizationEventTargetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationEventTargetType", value)
}

// NewOrganizationEventTargetTypeFromValue returns a pointer to a valid OrganizationEventTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationEventTargetTypeFromValue(v string) (*OrganizationEventTargetType, error) {
	ev := OrganizationEventTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationEventTargetType: valid values are %v", v, AllowedOrganizationEventTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationEventTargetType) IsValid() bool {
	for _, existing := range AllowedOrganizationEventTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationEventTargetType value
func (v OrganizationEventTargetType) Ptr() *OrganizationEventTargetType {
	return &v
}

type NullableOrganizationEventTargetType struct {
	value *OrganizationEventTargetType
	isSet bool
}

func (v NullableOrganizationEventTargetType) Get() *OrganizationEventTargetType {
	return v.value
}

func (v *NullableOrganizationEventTargetType) Set(val *OrganizationEventTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventTargetType(val *OrganizationEventTargetType) *NullableOrganizationEventTargetType {
	return &NullableOrganizationEventTargetType{value: val, isSet: true}
}

func (v NullableOrganizationEventTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}