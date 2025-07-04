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
)

// checks if the OrganizationCustomRoleUpdateRequestProjectPermissionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCustomRoleUpdateRequestProjectPermissionsInner{}

// OrganizationCustomRoleUpdateRequestProjectPermissionsInner struct for OrganizationCustomRoleUpdateRequestProjectPermissionsInner
type OrganizationCustomRoleUpdateRequestProjectPermissionsInner struct {
	ProjectId *string `json:"project_id,omitempty"`
	// If `is_admin` is `true`, the user is: - automatically `MANAGER` for each environment type - allowed to manage project deployment rules - able to delete the project    Note that `permissions` can then be ignored for this project
	IsAdmin *bool `json:"is_admin,omitempty"`
	// Mandatory if `is_admin` is `false`   Should contain an entry for every environment type: - `DEVELOPMENT` - `PREVIEW` - `STAGING` - `PRODUCTION`
	Permissions          []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationCustomRoleUpdateRequestProjectPermissionsInner OrganizationCustomRoleUpdateRequestProjectPermissionsInner

// NewOrganizationCustomRoleUpdateRequestProjectPermissionsInner instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInner() *OrganizationCustomRoleUpdateRequestProjectPermissionsInner {
	this := OrganizationCustomRoleUpdateRequestProjectPermissionsInner{}
	var isAdmin bool = false
	this.IsAdmin = &isAdmin
	return &this
}

// NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerWithDefaults() *OrganizationCustomRoleUpdateRequestProjectPermissionsInner {
	this := OrganizationCustomRoleUpdateRequestProjectPermissionsInner{}
	var isAdmin bool = false
	this.IsAdmin = &isAdmin
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetPermissions() []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner {
	if o == nil || IsNil(o.Permissions) {
		var ret []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetPermissionsOk() ([]OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner and assigns it to the Permissions field.
func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) SetPermissions(v []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) {
	o.Permissions = v
}

func (o OrganizationCustomRoleUpdateRequestProjectPermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCustomRoleUpdateRequestProjectPermissionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["is_admin"] = o.IsAdmin
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) UnmarshalJSON(data []byte) (err error) {
	varOrganizationCustomRoleUpdateRequestProjectPermissionsInner := _OrganizationCustomRoleUpdateRequestProjectPermissionsInner{}

	err = json.Unmarshal(data, &varOrganizationCustomRoleUpdateRequestProjectPermissionsInner)

	if err != nil {
		return err
	}

	*o = OrganizationCustomRoleUpdateRequestProjectPermissionsInner(varOrganizationCustomRoleUpdateRequestProjectPermissionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "is_admin")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner struct {
	value *OrganizationCustomRoleUpdateRequestProjectPermissionsInner
	isSet bool
}

func (v NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner) Get() *OrganizationCustomRoleUpdateRequestProjectPermissionsInner {
	return v.value
}

func (v *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner) Set(val *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner(val *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner {
	return &NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner{value: val, isSet: true}
}

func (v NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomRoleUpdateRequestProjectPermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
