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

// ProjectCurrentCostResponseList struct for ProjectCurrentCostResponseList
type ProjectCurrentCostResponseList struct {
	Projects *[]ProjectCurrentCostResponse `json:"projects,omitempty"`
}

// NewProjectCurrentCostResponseList instantiates a new ProjectCurrentCostResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCurrentCostResponseList() *ProjectCurrentCostResponseList {
	this := ProjectCurrentCostResponseList{}
	return &this
}

// NewProjectCurrentCostResponseListWithDefaults instantiates a new ProjectCurrentCostResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCurrentCostResponseListWithDefaults() *ProjectCurrentCostResponseList {
	this := ProjectCurrentCostResponseList{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *ProjectCurrentCostResponseList) GetProjects() []ProjectCurrentCostResponse {
	if o == nil || o.Projects == nil {
		var ret []ProjectCurrentCostResponse
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCurrentCostResponseList) GetProjectsOk() (*[]ProjectCurrentCostResponse, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ProjectCurrentCostResponseList) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectCurrentCostResponse and assigns it to the Projects field.
func (o *ProjectCurrentCostResponseList) SetProjects(v []ProjectCurrentCostResponse) {
	o.Projects = &v
}

func (o ProjectCurrentCostResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCurrentCostResponseList struct {
	value *ProjectCurrentCostResponseList
	isSet bool
}

func (v NullableProjectCurrentCostResponseList) Get() *ProjectCurrentCostResponseList {
	return v.value
}

func (v *NullableProjectCurrentCostResponseList) Set(val *ProjectCurrentCostResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCurrentCostResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCurrentCostResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCurrentCostResponseList(val *ProjectCurrentCostResponseList) *NullableProjectCurrentCostResponseList {
	return &NullableProjectCurrentCostResponseList{value: val, isSet: true}
}

func (v NullableProjectCurrentCostResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCurrentCostResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
