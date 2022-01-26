/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// EnvironmentRestartRequest struct for EnvironmentRestartRequest
type EnvironmentRestartRequest struct {
	RestartDb *bool `json:"restart_db,omitempty"`
}

// NewEnvironmentRestartRequest instantiates a new EnvironmentRestartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentRestartRequest() *EnvironmentRestartRequest {
	this := EnvironmentRestartRequest{}
	var restartDb bool = false
	this.RestartDb = &restartDb
	return &this
}

// NewEnvironmentRestartRequestWithDefaults instantiates a new EnvironmentRestartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentRestartRequestWithDefaults() *EnvironmentRestartRequest {
	this := EnvironmentRestartRequest{}
	var restartDb bool = false
	this.RestartDb = &restartDb
	return &this
}

// GetRestartDb returns the RestartDb field value if set, zero value otherwise.
func (o *EnvironmentRestartRequest) GetRestartDb() bool {
	if o == nil || o.RestartDb == nil {
		var ret bool
		return ret
	}
	return *o.RestartDb
}

// GetRestartDbOk returns a tuple with the RestartDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRestartRequest) GetRestartDbOk() (*bool, bool) {
	if o == nil || o.RestartDb == nil {
		return nil, false
	}
	return o.RestartDb, true
}

// HasRestartDb returns a boolean if a field has been set.
func (o *EnvironmentRestartRequest) HasRestartDb() bool {
	if o != nil && o.RestartDb != nil {
		return true
	}

	return false
}

// SetRestartDb gets a reference to the given bool and assigns it to the RestartDb field.
func (o *EnvironmentRestartRequest) SetRestartDb(v bool) {
	o.RestartDb = &v
}

func (o EnvironmentRestartRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RestartDb != nil {
		toSerialize["restart_db"] = o.RestartDb
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentRestartRequest struct {
	value *EnvironmentRestartRequest
	isSet bool
}

func (v NullableEnvironmentRestartRequest) Get() *EnvironmentRestartRequest {
	return v.value
}

func (v *NullableEnvironmentRestartRequest) Set(val *EnvironmentRestartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRestartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRestartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRestartRequest(val *EnvironmentRestartRequest) *NullableEnvironmentRestartRequest {
	return &NullableEnvironmentRestartRequest{value: val, isSet: true}
}

func (v NullableEnvironmentRestartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRestartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
