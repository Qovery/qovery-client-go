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

// DeploymentHistoryResponseList struct for DeploymentHistoryResponseList
type DeploymentHistoryResponseList struct {
	Results *[]DeploymentHistoryResponse `json:"results,omitempty"`
}

// NewDeploymentHistoryResponseList instantiates a new DeploymentHistoryResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryResponseList() *DeploymentHistoryResponseList {
	this := DeploymentHistoryResponseList{}
	return &this
}

// NewDeploymentHistoryResponseListWithDefaults instantiates a new DeploymentHistoryResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryResponseListWithDefaults() *DeploymentHistoryResponseList {
	this := DeploymentHistoryResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeploymentHistoryResponseList) GetResults() []DeploymentHistoryResponse {
	if o == nil || o.Results == nil {
		var ret []DeploymentHistoryResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryResponseList) GetResultsOk() (*[]DeploymentHistoryResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeploymentHistoryResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeploymentHistoryResponse and assigns it to the Results field.
func (o *DeploymentHistoryResponseList) SetResults(v []DeploymentHistoryResponse) {
	o.Results = &v
}

func (o DeploymentHistoryResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryResponseList struct {
	value *DeploymentHistoryResponseList
	isSet bool
}

func (v NullableDeploymentHistoryResponseList) Get() *DeploymentHistoryResponseList {
	return v.value
}

func (v *NullableDeploymentHistoryResponseList) Set(val *DeploymentHistoryResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryResponseList(val *DeploymentHistoryResponseList) *NullableDeploymentHistoryResponseList {
	return &NullableDeploymentHistoryResponseList{value: val, isSet: true}
}

func (v NullableDeploymentHistoryResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
