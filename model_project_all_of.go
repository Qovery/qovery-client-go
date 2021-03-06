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

// ProjectAllOf struct for ProjectAllOf
type ProjectAllOf struct {
	Name         string           `json:"name"`
	Description  *string          `json:"description,omitempty"`
	Organization *ReferenceObject `json:"organization,omitempty"`
}

// NewProjectAllOf instantiates a new ProjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAllOf(name string) *ProjectAllOf {
	this := ProjectAllOf{}
	this.Name = name
	return &this
}

// NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAllOfWithDefaults() *ProjectAllOf {
	this := ProjectAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectAllOf) GetOrganization() ReferenceObject {
	if o == nil || o.Organization == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetOrganizationOk() (*ReferenceObject, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given ReferenceObject and assigns it to the Organization field.
func (o *ProjectAllOf) SetOrganization(v ReferenceObject) {
	o.Organization = &v
}

func (o ProjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableProjectAllOf struct {
	value *ProjectAllOf
	isSet bool
}

func (v NullableProjectAllOf) Get() *ProjectAllOf {
	return v.value
}

func (v *NullableProjectAllOf) Set(val *ProjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAllOf(val *ProjectAllOf) *NullableProjectAllOf {
	return &NullableProjectAllOf{value: val, isSet: true}
}

func (v NullableProjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
