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

// TagRequest struct for TagRequest
type TagRequest struct {
	Name string `json:"name"`
}

// NewTagRequest instantiates a new TagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagRequest(name string) *TagRequest {
	this := TagRequest{}
	this.Name = name
	return &this
}

// NewTagRequestWithDefaults instantiates a new TagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagRequestWithDefaults() *TagRequest {
	this := TagRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TagRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagRequest) SetName(v string) {
	o.Name = v
}

func (o TagRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTagRequest struct {
	value *TagRequest
	isSet bool
}

func (v NullableTagRequest) Get() *TagRequest {
	return v.value
}

func (v *NullableTagRequest) Set(val *TagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRequest(val *TagRequest) *NullableTagRequest {
	return &NullableTagRequest{value: val, isSet: true}
}

func (v NullableTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
