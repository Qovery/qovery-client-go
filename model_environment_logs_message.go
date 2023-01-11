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

// EnvironmentLogsMessage struct for EnvironmentLogsMessage
type EnvironmentLogsMessage struct {
	SafeMessage *string `json:"safe_message,omitempty"`
	FullDetails *string `json:"full_details,omitempty"`
}

// NewEnvironmentLogsMessage instantiates a new EnvironmentLogsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogsMessage() *EnvironmentLogsMessage {
	this := EnvironmentLogsMessage{}
	return &this
}

// NewEnvironmentLogsMessageWithDefaults instantiates a new EnvironmentLogsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogsMessageWithDefaults() *EnvironmentLogsMessage {
	this := EnvironmentLogsMessage{}
	return &this
}

// GetSafeMessage returns the SafeMessage field value if set, zero value otherwise.
func (o *EnvironmentLogsMessage) GetSafeMessage() string {
	if o == nil || o.SafeMessage == nil {
		var ret string
		return ret
	}
	return *o.SafeMessage
}

// GetSafeMessageOk returns a tuple with the SafeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsMessage) GetSafeMessageOk() (*string, bool) {
	if o == nil || o.SafeMessage == nil {
		return nil, false
	}
	return o.SafeMessage, true
}

// HasSafeMessage returns a boolean if a field has been set.
func (o *EnvironmentLogsMessage) HasSafeMessage() bool {
	if o != nil && o.SafeMessage != nil {
		return true
	}

	return false
}

// SetSafeMessage gets a reference to the given string and assigns it to the SafeMessage field.
func (o *EnvironmentLogsMessage) SetSafeMessage(v string) {
	o.SafeMessage = &v
}

// GetFullDetails returns the FullDetails field value if set, zero value otherwise.
func (o *EnvironmentLogsMessage) GetFullDetails() string {
	if o == nil || o.FullDetails == nil {
		var ret string
		return ret
	}
	return *o.FullDetails
}

// GetFullDetailsOk returns a tuple with the FullDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogsMessage) GetFullDetailsOk() (*string, bool) {
	if o == nil || o.FullDetails == nil {
		return nil, false
	}
	return o.FullDetails, true
}

// HasFullDetails returns a boolean if a field has been set.
func (o *EnvironmentLogsMessage) HasFullDetails() bool {
	if o != nil && o.FullDetails != nil {
		return true
	}

	return false
}

// SetFullDetails gets a reference to the given string and assigns it to the FullDetails field.
func (o *EnvironmentLogsMessage) SetFullDetails(v string) {
	o.FullDetails = &v
}

func (o EnvironmentLogsMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafeMessage != nil {
		toSerialize["safe_message"] = o.SafeMessage
	}
	if o.FullDetails != nil {
		toSerialize["full_details"] = o.FullDetails
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentLogsMessage struct {
	value *EnvironmentLogsMessage
	isSet bool
}

func (v NullableEnvironmentLogsMessage) Get() *EnvironmentLogsMessage {
	return v.value
}

func (v *NullableEnvironmentLogsMessage) Set(val *EnvironmentLogsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogsMessage(val *EnvironmentLogsMessage) *NullableEnvironmentLogsMessage {
	return &NullableEnvironmentLogsMessage{value: val, isSet: true}
}

func (v NullableEnvironmentLogsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}