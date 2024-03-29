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

// checks if the VariableEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableEditRequest{}

// VariableEditRequest struct for VariableEditRequest
type VariableEditRequest struct {
	// the key of the environment variable
	Key string `json:"key"`
	// the value of the environment variable
	Value string `json:"value"`
}

// NewVariableEditRequest instantiates a new VariableEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableEditRequest(key string, value string) *VariableEditRequest {
	this := VariableEditRequest{}
	this.Key = key
	this.Value = value
	return &this
}

// NewVariableEditRequestWithDefaults instantiates a new VariableEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableEditRequestWithDefaults() *VariableEditRequest {
	this := VariableEditRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *VariableEditRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *VariableEditRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *VariableEditRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *VariableEditRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VariableEditRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VariableEditRequest) SetValue(v string) {
	o.Value = v
}

func (o VariableEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableVariableEditRequest struct {
	value *VariableEditRequest
	isSet bool
}

func (v NullableVariableEditRequest) Get() *VariableEditRequest {
	return v.value
}

func (v *NullableVariableEditRequest) Set(val *VariableEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableEditRequest(val *VariableEditRequest) *NullableVariableEditRequest {
	return &NullableVariableEditRequest{value: val, isSet: true}
}

func (v NullableVariableEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
