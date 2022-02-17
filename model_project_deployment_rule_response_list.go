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

// ProjectDeploymentRuleResponseList struct for ProjectDeploymentRuleResponseList
type ProjectDeploymentRuleResponseList struct {
	Results []ProjectDeploymentRuleResponse `json:"results,omitempty"`
}

// NewProjectDeploymentRuleResponseList instantiates a new ProjectDeploymentRuleResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDeploymentRuleResponseList() *ProjectDeploymentRuleResponseList {
	this := ProjectDeploymentRuleResponseList{}
	return &this
}

// NewProjectDeploymentRuleResponseListWithDefaults instantiates a new ProjectDeploymentRuleResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDeploymentRuleResponseListWithDefaults() *ProjectDeploymentRuleResponseList {
	this := ProjectDeploymentRuleResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProjectDeploymentRuleResponseList) GetResults() []ProjectDeploymentRuleResponse {
	if o == nil || o.Results == nil {
		var ret []ProjectDeploymentRuleResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDeploymentRuleResponseList) GetResultsOk() ([]ProjectDeploymentRuleResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProjectDeploymentRuleResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProjectDeploymentRuleResponse and assigns it to the Results field.
func (o *ProjectDeploymentRuleResponseList) SetResults(v []ProjectDeploymentRuleResponse) {
	o.Results = v
}

func (o ProjectDeploymentRuleResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDeploymentRuleResponseList struct {
	value *ProjectDeploymentRuleResponseList
	isSet bool
}

func (v NullableProjectDeploymentRuleResponseList) Get() *ProjectDeploymentRuleResponseList {
	return v.value
}

func (v *NullableProjectDeploymentRuleResponseList) Set(val *ProjectDeploymentRuleResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDeploymentRuleResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDeploymentRuleResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDeploymentRuleResponseList(val *ProjectDeploymentRuleResponseList) *NullableProjectDeploymentRuleResponseList {
	return &NullableProjectDeploymentRuleResponseList{value: val, isSet: true}
}

func (v NullableProjectDeploymentRuleResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDeploymentRuleResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
