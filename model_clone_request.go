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

// CloneRequest struct for CloneRequest
type CloneRequest struct {
	// name is case insensitive
	Name      string               `json:"name"`
	ClusterId *string              `json:"cluster_id,omitempty"`
	Mode      *EnvironmentModeEnum `json:"mode,omitempty"`
}

// NewCloneRequest instantiates a new CloneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneRequest(name string) *CloneRequest {
	this := CloneRequest{}
	this.Name = name
	return &this
}

// NewCloneRequestWithDefaults instantiates a new CloneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneRequestWithDefaults() *CloneRequest {
	this := CloneRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CloneRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneRequest) SetName(v string) {
	o.Name = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *CloneRequest) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneRequest) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *CloneRequest) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *CloneRequest) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CloneRequest) GetMode() EnvironmentModeEnum {
	if o == nil || o.Mode == nil {
		var ret EnvironmentModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneRequest) GetModeOk() (*EnvironmentModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CloneRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given EnvironmentModeEnum and assigns it to the Mode field.
func (o *CloneRequest) SetMode(v EnvironmentModeEnum) {
	o.Mode = &v
}

func (o CloneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableCloneRequest struct {
	value *CloneRequest
	isSet bool
}

func (v NullableCloneRequest) Get() *CloneRequest {
	return v.value
}

func (v *NullableCloneRequest) Set(val *CloneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneRequest(val *CloneRequest) *NullableCloneRequest {
	return &NullableCloneRequest{value: val, isSet: true}
}

func (v NullableCloneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
