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

// OrganizationAnnotationsGroupScopeEnum Annotations Group Scope
type OrganizationAnnotationsGroupScopeEnum string

// List of OrganizationAnnotationsGroupScopeEnum
const (
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_DEPLOYMENTS   OrganizationAnnotationsGroupScopeEnum = "DEPLOYMENTS"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_STATEFUL_SETS OrganizationAnnotationsGroupScopeEnum = "STATEFUL_SETS"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_SERVICES      OrganizationAnnotationsGroupScopeEnum = "SERVICES"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_INGRESS       OrganizationAnnotationsGroupScopeEnum = "INGRESS"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_HPA           OrganizationAnnotationsGroupScopeEnum = "HPA"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_PODS          OrganizationAnnotationsGroupScopeEnum = "PODS"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_SECRETS       OrganizationAnnotationsGroupScopeEnum = "SECRETS"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_JOBS          OrganizationAnnotationsGroupScopeEnum = "JOBS"
	ORGANIZATIONANNOTATIONSGROUPSCOPEENUM_CRON_JOBS     OrganizationAnnotationsGroupScopeEnum = "CRON_JOBS"
)

// All allowed values of OrganizationAnnotationsGroupScopeEnum enum
var AllowedOrganizationAnnotationsGroupScopeEnumEnumValues = []OrganizationAnnotationsGroupScopeEnum{
	"DEPLOYMENTS",
	"STATEFUL_SETS",
	"SERVICES",
	"INGRESS",
	"HPA",
	"PODS",
	"SECRETS",
	"JOBS",
	"CRON_JOBS",
}

func (v *OrganizationAnnotationsGroupScopeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationAnnotationsGroupScopeEnum(value)
	for _, existing := range AllowedOrganizationAnnotationsGroupScopeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationAnnotationsGroupScopeEnum", value)
}

// NewOrganizationAnnotationsGroupScopeEnumFromValue returns a pointer to a valid OrganizationAnnotationsGroupScopeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationAnnotationsGroupScopeEnumFromValue(v string) (*OrganizationAnnotationsGroupScopeEnum, error) {
	ev := OrganizationAnnotationsGroupScopeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationAnnotationsGroupScopeEnum: valid values are %v", v, AllowedOrganizationAnnotationsGroupScopeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationAnnotationsGroupScopeEnum) IsValid() bool {
	for _, existing := range AllowedOrganizationAnnotationsGroupScopeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationAnnotationsGroupScopeEnum value
func (v OrganizationAnnotationsGroupScopeEnum) Ptr() *OrganizationAnnotationsGroupScopeEnum {
	return &v
}

type NullableOrganizationAnnotationsGroupScopeEnum struct {
	value *OrganizationAnnotationsGroupScopeEnum
	isSet bool
}

func (v NullableOrganizationAnnotationsGroupScopeEnum) Get() *OrganizationAnnotationsGroupScopeEnum {
	return v.value
}

func (v *NullableOrganizationAnnotationsGroupScopeEnum) Set(val *OrganizationAnnotationsGroupScopeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAnnotationsGroupScopeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAnnotationsGroupScopeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAnnotationsGroupScopeEnum(val *OrganizationAnnotationsGroupScopeEnum) *NullableOrganizationAnnotationsGroupScopeEnum {
	return &NullableOrganizationAnnotationsGroupScopeEnum{value: val, isSet: true}
}

func (v NullableOrganizationAnnotationsGroupScopeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAnnotationsGroupScopeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
