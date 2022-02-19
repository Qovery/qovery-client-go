/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.2
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ProjectRequest struct for ProjectRequest
type ProjectRequest struct {
	// name is case insensitive
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewProjectRequest instantiates a new ProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRequest(name string) *ProjectRequest {
	this := ProjectRequest{}
	this.Name = name
	return &this
}

// NewProjectRequestWithDefaults instantiates a new ProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRequestWithDefaults() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableProjectRequest struct {
	value *ProjectRequest
	isSet bool
}

func (v NullableProjectRequest) Get() *ProjectRequest {
	return v.value
}

func (v *NullableProjectRequest) Set(val *ProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRequest(val *ProjectRequest) *NullableProjectRequest {
	return &NullableProjectRequest{value: val, isSet: true}
}

func (v NullableProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
