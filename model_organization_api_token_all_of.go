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

// OrganizationApiTokenAllOf struct for OrganizationApiTokenAllOf
type OrganizationApiTokenAllOf struct {
	Name        *string                    `json:"name,omitempty"`
	Description *string                    `json:"description,omitempty"`
	Scope       *OrganizationApiTokenScope `json:"scope,omitempty"`
}

// NewOrganizationApiTokenAllOf instantiates a new OrganizationApiTokenAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationApiTokenAllOf() *OrganizationApiTokenAllOf {
	this := OrganizationApiTokenAllOf{}
	return &this
}

// NewOrganizationApiTokenAllOfWithDefaults instantiates a new OrganizationApiTokenAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationApiTokenAllOfWithDefaults() *OrganizationApiTokenAllOf {
	this := OrganizationApiTokenAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationApiTokenAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationApiTokenAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationApiTokenAllOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationApiTokenAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationApiTokenAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationApiTokenAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OrganizationApiTokenAllOf) GetScope() OrganizationApiTokenScope {
	if o == nil || o.Scope == nil {
		var ret OrganizationApiTokenScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenAllOf) GetScopeOk() (*OrganizationApiTokenScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OrganizationApiTokenAllOf) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given OrganizationApiTokenScope and assigns it to the Scope field.
func (o *OrganizationApiTokenAllOf) SetScope(v OrganizationApiTokenScope) {
	o.Scope = &v
}

func (o OrganizationApiTokenAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationApiTokenAllOf struct {
	value *OrganizationApiTokenAllOf
	isSet bool
}

func (v NullableOrganizationApiTokenAllOf) Get() *OrganizationApiTokenAllOf {
	return v.value
}

func (v *NullableOrganizationApiTokenAllOf) Set(val *OrganizationApiTokenAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationApiTokenAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationApiTokenAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationApiTokenAllOf(val *OrganizationApiTokenAllOf) *NullableOrganizationApiTokenAllOf {
	return &NullableOrganizationApiTokenAllOf{value: val, isSet: true}
}

func (v NullableOrganizationApiTokenAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationApiTokenAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
