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

// EnvironmentTotalNumber struct for EnvironmentTotalNumber
type EnvironmentTotalNumber struct {
	EnvironmentTotalNumber *float32 `json:"environment_total_number,omitempty"`
}

// NewEnvironmentTotalNumber instantiates a new EnvironmentTotalNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentTotalNumber() *EnvironmentTotalNumber {
	this := EnvironmentTotalNumber{}
	return &this
}

// NewEnvironmentTotalNumberWithDefaults instantiates a new EnvironmentTotalNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentTotalNumberWithDefaults() *EnvironmentTotalNumber {
	this := EnvironmentTotalNumber{}
	return &this
}

// GetEnvironmentTotalNumber returns the EnvironmentTotalNumber field value if set, zero value otherwise.
func (o *EnvironmentTotalNumber) GetEnvironmentTotalNumber() float32 {
	if o == nil || o.EnvironmentTotalNumber == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentTotalNumber
}

// GetEnvironmentTotalNumberOk returns a tuple with the EnvironmentTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTotalNumber) GetEnvironmentTotalNumberOk() (*float32, bool) {
	if o == nil || o.EnvironmentTotalNumber == nil {
		return nil, false
	}
	return o.EnvironmentTotalNumber, true
}

// HasEnvironmentTotalNumber returns a boolean if a field has been set.
func (o *EnvironmentTotalNumber) HasEnvironmentTotalNumber() bool {
	if o != nil && o.EnvironmentTotalNumber != nil {
		return true
	}

	return false
}

// SetEnvironmentTotalNumber gets a reference to the given float32 and assigns it to the EnvironmentTotalNumber field.
func (o *EnvironmentTotalNumber) SetEnvironmentTotalNumber(v float32) {
	o.EnvironmentTotalNumber = &v
}

func (o EnvironmentTotalNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnvironmentTotalNumber != nil {
		toSerialize["environment_total_number"] = o.EnvironmentTotalNumber
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentTotalNumber struct {
	value *EnvironmentTotalNumber
	isSet bool
}

func (v NullableEnvironmentTotalNumber) Get() *EnvironmentTotalNumber {
	return v.value
}

func (v *NullableEnvironmentTotalNumber) Set(val *EnvironmentTotalNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentTotalNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentTotalNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentTotalNumber(val *EnvironmentTotalNumber) *NullableEnvironmentTotalNumber {
	return &NullableEnvironmentTotalNumber{value: val, isSet: true}
}

func (v NullableEnvironmentTotalNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentTotalNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
