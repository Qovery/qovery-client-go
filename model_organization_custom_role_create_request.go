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

// checks if the OrganizationCustomRoleCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCustomRoleCreateRequest{}

// OrganizationCustomRoleCreateRequest struct for OrganizationCustomRoleCreateRequest
type OrganizationCustomRoleCreateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewOrganizationCustomRoleCreateRequest instantiates a new OrganizationCustomRoleCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomRoleCreateRequest(name string) *OrganizationCustomRoleCreateRequest {
	this := OrganizationCustomRoleCreateRequest{}
	this.Name = name
	return &this
}

// NewOrganizationCustomRoleCreateRequestWithDefaults instantiates a new OrganizationCustomRoleCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomRoleCreateRequestWithDefaults() *OrganizationCustomRoleCreateRequest {
	this := OrganizationCustomRoleCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationCustomRoleCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationCustomRoleCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationCustomRoleCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationCustomRoleCreateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationCustomRoleCreateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o OrganizationCustomRoleCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCustomRoleCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableOrganizationCustomRoleCreateRequest struct {
	value *OrganizationCustomRoleCreateRequest
	isSet bool
}

func (v NullableOrganizationCustomRoleCreateRequest) Get() *OrganizationCustomRoleCreateRequest {
	return v.value
}

func (v *NullableOrganizationCustomRoleCreateRequest) Set(val *OrganizationCustomRoleCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomRoleCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomRoleCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomRoleCreateRequest(val *OrganizationCustomRoleCreateRequest) *NullableOrganizationCustomRoleCreateRequest {
	return &NullableOrganizationCustomRoleCreateRequest{value: val, isSet: true}
}

func (v NullableOrganizationCustomRoleCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomRoleCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
