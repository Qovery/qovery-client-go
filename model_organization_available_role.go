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
)

// OrganizationAvailableRole struct for OrganizationAvailableRole
type OrganizationAvailableRole struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewOrganizationAvailableRole instantiates a new OrganizationAvailableRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAvailableRole() *OrganizationAvailableRole {
	this := OrganizationAvailableRole{}
	return &this
}

// NewOrganizationAvailableRoleWithDefaults instantiates a new OrganizationAvailableRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAvailableRoleWithDefaults() *OrganizationAvailableRole {
	this := OrganizationAvailableRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationAvailableRole) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAvailableRole) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationAvailableRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationAvailableRole) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationAvailableRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAvailableRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationAvailableRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationAvailableRole) SetName(v string) {
	o.Name = &v
}

func (o OrganizationAvailableRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationAvailableRole struct {
	value *OrganizationAvailableRole
	isSet bool
}

func (v NullableOrganizationAvailableRole) Get() *OrganizationAvailableRole {
	return v.value
}

func (v *NullableOrganizationAvailableRole) Set(val *OrganizationAvailableRole) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAvailableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAvailableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAvailableRole(val *OrganizationAvailableRole) *NullableOrganizationAvailableRole {
	return &NullableOrganizationAvailableRole{value: val, isSet: true}
}

func (v NullableOrganizationAvailableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAvailableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}