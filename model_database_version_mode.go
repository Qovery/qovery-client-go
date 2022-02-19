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

// DatabaseVersionMode struct for DatabaseVersionMode
type DatabaseVersionMode struct {
	Name          *string `json:"name,omitempty"`
	SupportedMode *string `json:"supported_mode,omitempty"`
}

// NewDatabaseVersionMode instantiates a new DatabaseVersionMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseVersionMode() *DatabaseVersionMode {
	this := DatabaseVersionMode{}
	return &this
}

// NewDatabaseVersionModeWithDefaults instantiates a new DatabaseVersionMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseVersionModeWithDefaults() *DatabaseVersionMode {
	this := DatabaseVersionMode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatabaseVersionMode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseVersionMode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseVersionMode) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseVersionMode) SetName(v string) {
	o.Name = &v
}

// GetSupportedMode returns the SupportedMode field value if set, zero value otherwise.
func (o *DatabaseVersionMode) GetSupportedMode() string {
	if o == nil || o.SupportedMode == nil {
		var ret string
		return ret
	}
	return *o.SupportedMode
}

// GetSupportedModeOk returns a tuple with the SupportedMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseVersionMode) GetSupportedModeOk() (*string, bool) {
	if o == nil || o.SupportedMode == nil {
		return nil, false
	}
	return o.SupportedMode, true
}

// HasSupportedMode returns a boolean if a field has been set.
func (o *DatabaseVersionMode) HasSupportedMode() bool {
	if o != nil && o.SupportedMode != nil {
		return true
	}

	return false
}

// SetSupportedMode gets a reference to the given string and assigns it to the SupportedMode field.
func (o *DatabaseVersionMode) SetSupportedMode(v string) {
	o.SupportedMode = &v
}

func (o DatabaseVersionMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SupportedMode != nil {
		toSerialize["supported_mode"] = o.SupportedMode
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseVersionMode struct {
	value *DatabaseVersionMode
	isSet bool
}

func (v NullableDatabaseVersionMode) Get() *DatabaseVersionMode {
	return v.value
}

func (v *NullableDatabaseVersionMode) Set(val *DatabaseVersionMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseVersionMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseVersionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseVersionMode(val *DatabaseVersionMode) *NullableDatabaseVersionMode {
	return &NullableDatabaseVersionMode{value: val, isSet: true}
}

func (v NullableDatabaseVersionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseVersionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
