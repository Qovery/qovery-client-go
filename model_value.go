/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// Value struct for Value
type Value struct {
	Value string `json:"value"`
}

// NewValue instantiates a new Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValue(value string) *Value {
	this := Value{}
	this.Value = value
	return &this
}

// NewValueWithDefaults instantiates a new Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueWithDefaults() *Value {
	this := Value{}
	return &this
}

// GetValue returns the Value field value
func (o *Value) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Value) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Value) SetValue(v string) {
	o.Value = v
}

func (o Value) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableValue struct {
	value *Value
	isSet bool
}

func (v NullableValue) Get() *Value {
	return v.value
}

func (v *NullableValue) Set(val *Value) {
	v.value = val
	v.isSet = true
}

func (v NullableValue) IsSet() bool {
	return v.isSet
}

func (v *NullableValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValue(val *Value) *NullableValue {
	return &NullableValue{value: val, isSet: true}
}

func (v NullableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
