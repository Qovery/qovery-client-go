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

// OrganizationCustomRoleClusterPermission Indicates the permission for a target cluster, from the lowest to the highest: - `VIEWER` user has only read access on target cluster - `ENV_CREATOR` user can deploy on the cluster - `ADMIN` user can modify the cluster settings
type OrganizationCustomRoleClusterPermission string

// List of OrganizationCustomRoleClusterPermission
const (
	ORGANIZATIONCUSTOMROLECLUSTERPERMISSION_VIEWER      OrganizationCustomRoleClusterPermission = "VIEWER"
	ORGANIZATIONCUSTOMROLECLUSTERPERMISSION_ENV_CREATOR OrganizationCustomRoleClusterPermission = "ENV_CREATOR"
	ORGANIZATIONCUSTOMROLECLUSTERPERMISSION_ADMIN       OrganizationCustomRoleClusterPermission = "ADMIN"
)

// All allowed values of OrganizationCustomRoleClusterPermission enum
var AllowedOrganizationCustomRoleClusterPermissionEnumValues = []OrganizationCustomRoleClusterPermission{
	"VIEWER",
	"ENV_CREATOR",
	"ADMIN",
}

func (v *OrganizationCustomRoleClusterPermission) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationCustomRoleClusterPermission(value)
	for _, existing := range AllowedOrganizationCustomRoleClusterPermissionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationCustomRoleClusterPermission", value)
}

// NewOrganizationCustomRoleClusterPermissionFromValue returns a pointer to a valid OrganizationCustomRoleClusterPermission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationCustomRoleClusterPermissionFromValue(v string) (*OrganizationCustomRoleClusterPermission, error) {
	ev := OrganizationCustomRoleClusterPermission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationCustomRoleClusterPermission: valid values are %v", v, AllowedOrganizationCustomRoleClusterPermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationCustomRoleClusterPermission) IsValid() bool {
	for _, existing := range AllowedOrganizationCustomRoleClusterPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationCustomRoleClusterPermission value
func (v OrganizationCustomRoleClusterPermission) Ptr() *OrganizationCustomRoleClusterPermission {
	return &v
}

type NullableOrganizationCustomRoleClusterPermission struct {
	value *OrganizationCustomRoleClusterPermission
	isSet bool
}

func (v NullableOrganizationCustomRoleClusterPermission) Get() *OrganizationCustomRoleClusterPermission {
	return v.value
}

func (v *NullableOrganizationCustomRoleClusterPermission) Set(val *OrganizationCustomRoleClusterPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomRoleClusterPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomRoleClusterPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomRoleClusterPermission(val *OrganizationCustomRoleClusterPermission) *NullableOrganizationCustomRoleClusterPermission {
	return &NullableOrganizationCustomRoleClusterPermission{value: val, isSet: true}
}

func (v NullableOrganizationCustomRoleClusterPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomRoleClusterPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
