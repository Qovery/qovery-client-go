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

// JobDeploymentRestrictionRequest struct for JobDeploymentRestrictionRequest
type JobDeploymentRestrictionRequest struct {
	Mode DeploymentRestrictionModeEnum `json:"mode"`
	Type DeploymentRestrictionTypeEnum `json:"type"`
	// For `PATH` restrictions, the value must not start with `/`
	Value string `json:"value"`
}

// NewJobDeploymentRestrictionRequest instantiates a new JobDeploymentRestrictionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDeploymentRestrictionRequest(mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string) *JobDeploymentRestrictionRequest {
	this := JobDeploymentRestrictionRequest{}
	this.Mode = mode
	this.Type = type_
	this.Value = value
	return &this
}

// NewJobDeploymentRestrictionRequestWithDefaults instantiates a new JobDeploymentRestrictionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDeploymentRestrictionRequestWithDefaults() *JobDeploymentRestrictionRequest {
	this := JobDeploymentRestrictionRequest{}
	return &this
}

// GetMode returns the Mode field value
func (o *JobDeploymentRestrictionRequest) GetMode() DeploymentRestrictionModeEnum {
	if o == nil {
		var ret DeploymentRestrictionModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionRequest) GetModeOk() (*DeploymentRestrictionModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *JobDeploymentRestrictionRequest) SetMode(v DeploymentRestrictionModeEnum) {
	o.Mode = v
}

// GetType returns the Type field value
func (o *JobDeploymentRestrictionRequest) GetType() DeploymentRestrictionTypeEnum {
	if o == nil {
		var ret DeploymentRestrictionTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionRequest) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobDeploymentRestrictionRequest) SetType(v DeploymentRestrictionTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *JobDeploymentRestrictionRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *JobDeploymentRestrictionRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *JobDeploymentRestrictionRequest) SetValue(v string) {
	o.Value = v
}

func (o JobDeploymentRestrictionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableJobDeploymentRestrictionRequest struct {
	value *JobDeploymentRestrictionRequest
	isSet bool
}

func (v NullableJobDeploymentRestrictionRequest) Get() *JobDeploymentRestrictionRequest {
	return v.value
}

func (v *NullableJobDeploymentRestrictionRequest) Set(val *JobDeploymentRestrictionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDeploymentRestrictionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDeploymentRestrictionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDeploymentRestrictionRequest(val *JobDeploymentRestrictionRequest) *NullableJobDeploymentRestrictionRequest {
	return &NullableJobDeploymentRestrictionRequest{value: val, isSet: true}
}

func (v NullableJobDeploymentRestrictionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDeploymentRestrictionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
