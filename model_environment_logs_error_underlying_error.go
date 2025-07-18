/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.4
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// checks if the EnvironmentLogsErrorUnderlyingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentLogsErrorUnderlyingError{}

// EnvironmentLogsErrorUnderlyingError struct for EnvironmentLogsErrorUnderlyingError
type EnvironmentLogsErrorUnderlyingError struct {
	Message              *string `json:"message,omitempty"`
	FullDetails          *string `json:"full_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentLogsErrorUnderlyingError EnvironmentLogsErrorUnderlyingError

// NewEnvironmentLogsErrorUnderlyingError instantiates a new EnvironmentLogsErrorUnderlyingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogsErrorUnderlyingError() *EnvironmentLogsErrorUnderlyingError {
	this := EnvironmentLogsErrorUnderlyingError{}
	return &this
}

// NewEnvironmentLogsErrorUnderlyingErrorWithDefaults instantiates a new EnvironmentLogsErrorUnderlyingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogsErrorUnderlyingErrorWithDefaults() *EnvironmentLogsErrorUnderlyingError {
	this := EnvironmentLogsErrorUnderlyingError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EnvironmentLogsErrorUnderlyingError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsErrorUnderlyingError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EnvironmentLogsErrorUnderlyingError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EnvironmentLogsErrorUnderlyingError) SetMessage(v string) {
	o.Message = &v
}

// GetFullDetails returns the FullDetails field value if set, zero value otherwise.
func (o *EnvironmentLogsErrorUnderlyingError) GetFullDetails() string {
	if o == nil || IsNil(o.FullDetails) {
		var ret string
		return ret
	}
	return *o.FullDetails
}

// GetFullDetailsOk returns a tuple with the FullDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsErrorUnderlyingError) GetFullDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.FullDetails) {
		return nil, false
	}
	return o.FullDetails, true
}

// HasFullDetails returns a boolean if a field has been set.
func (o *EnvironmentLogsErrorUnderlyingError) HasFullDetails() bool {
	if o != nil && !IsNil(o.FullDetails) {
		return true
	}

	return false
}

// SetFullDetails gets a reference to the given string and assigns it to the FullDetails field.
func (o *EnvironmentLogsErrorUnderlyingError) SetFullDetails(v string) {
	o.FullDetails = &v
}

func (o EnvironmentLogsErrorUnderlyingError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentLogsErrorUnderlyingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.FullDetails) {
		toSerialize["full_details"] = o.FullDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentLogsErrorUnderlyingError) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentLogsErrorUnderlyingError := _EnvironmentLogsErrorUnderlyingError{}

	err = json.Unmarshal(data, &varEnvironmentLogsErrorUnderlyingError)

	if err != nil {
		return err
	}

	*o = EnvironmentLogsErrorUnderlyingError(varEnvironmentLogsErrorUnderlyingError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "full_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentLogsErrorUnderlyingError struct {
	value *EnvironmentLogsErrorUnderlyingError
	isSet bool
}

func (v NullableEnvironmentLogsErrorUnderlyingError) Get() *EnvironmentLogsErrorUnderlyingError {
	return v.value
}

func (v *NullableEnvironmentLogsErrorUnderlyingError) Set(val *EnvironmentLogsErrorUnderlyingError) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogsErrorUnderlyingError) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogsErrorUnderlyingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogsErrorUnderlyingError(val *EnvironmentLogsErrorUnderlyingError) *NullableEnvironmentLogsErrorUnderlyingError {
	return &NullableEnvironmentLogsErrorUnderlyingError{value: val, isSet: true}
}

func (v NullableEnvironmentLogsErrorUnderlyingError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogsErrorUnderlyingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
