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

// checks if the UnexpectedError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnexpectedError{}

// UnexpectedError struct for UnexpectedError
type UnexpectedError struct {
	Message string `json:"message"`
}

// NewUnexpectedError instantiates a new UnexpectedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnexpectedError(message string) *UnexpectedError {
	this := UnexpectedError{}
	this.Message = message
	return &this
}

// NewUnexpectedErrorWithDefaults instantiates a new UnexpectedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnexpectedErrorWithDefaults() *UnexpectedError {
	this := UnexpectedError{}
	return &this
}

// GetMessage returns the Message field value
func (o *UnexpectedError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UnexpectedError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UnexpectedError) SetMessage(v string) {
	o.Message = v
}

func (o UnexpectedError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnexpectedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUnexpectedError struct {
	value *UnexpectedError
	isSet bool
}

func (v NullableUnexpectedError) Get() *UnexpectedError {
	return v.value
}

func (v *NullableUnexpectedError) Set(val *UnexpectedError) {
	v.value = val
	v.isSet = true
}

func (v NullableUnexpectedError) IsSet() bool {
	return v.isSet
}

func (v *NullableUnexpectedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnexpectedError(val *UnexpectedError) *NullableUnexpectedError {
	return &NullableUnexpectedError{value: val, isSet: true}
}

func (v NullableUnexpectedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnexpectedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
