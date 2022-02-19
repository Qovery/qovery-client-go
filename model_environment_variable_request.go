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

// EnvironmentVariableRequest struct for EnvironmentVariableRequest
type EnvironmentVariableRequest struct {
	// key is case sensitive
	Key string `json:"key"`
	// value of the env variable.
	Value string `json:"value"`
}

// NewEnvironmentVariableRequest instantiates a new EnvironmentVariableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableRequest(key string, value string) *EnvironmentVariableRequest {
	this := EnvironmentVariableRequest{}
	this.Key = key
	this.Value = value
	return &this
}

// NewEnvironmentVariableRequestWithDefaults instantiates a new EnvironmentVariableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableRequestWithDefaults() *EnvironmentVariableRequest {
	this := EnvironmentVariableRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *EnvironmentVariableRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariableRequest) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariableRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariableRequest) SetValue(v string) {
	o.Value = v
}

func (o EnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableRequest struct {
	value *EnvironmentVariableRequest
	isSet bool
}

func (v NullableEnvironmentVariableRequest) Get() *EnvironmentVariableRequest {
	return v.value
}

func (v *NullableEnvironmentVariableRequest) Set(val *EnvironmentVariableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableRequest(val *EnvironmentVariableRequest) *NullableEnvironmentVariableRequest {
	return &NullableEnvironmentVariableRequest{value: val, isSet: true}
}

func (v NullableEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
