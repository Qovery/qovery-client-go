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

// checks if the HelmDeploymentRestrictionResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmDeploymentRestrictionResponseList{}

// HelmDeploymentRestrictionResponseList struct for HelmDeploymentRestrictionResponseList
type HelmDeploymentRestrictionResponseList struct {
	Results []HelmDeploymentRestrictionResponse `json:"results,omitempty"`
}

// NewHelmDeploymentRestrictionResponseList instantiates a new HelmDeploymentRestrictionResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmDeploymentRestrictionResponseList() *HelmDeploymentRestrictionResponseList {
	this := HelmDeploymentRestrictionResponseList{}
	return &this
}

// NewHelmDeploymentRestrictionResponseListWithDefaults instantiates a new HelmDeploymentRestrictionResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmDeploymentRestrictionResponseListWithDefaults() *HelmDeploymentRestrictionResponseList {
	this := HelmDeploymentRestrictionResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HelmDeploymentRestrictionResponseList) GetResults() []HelmDeploymentRestrictionResponse {
	if o == nil || IsNil(o.Results) {
		var ret []HelmDeploymentRestrictionResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeploymentRestrictionResponseList) GetResultsOk() ([]HelmDeploymentRestrictionResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HelmDeploymentRestrictionResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HelmDeploymentRestrictionResponse and assigns it to the Results field.
func (o *HelmDeploymentRestrictionResponseList) SetResults(v []HelmDeploymentRestrictionResponse) {
	o.Results = v
}

func (o HelmDeploymentRestrictionResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmDeploymentRestrictionResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableHelmDeploymentRestrictionResponseList struct {
	value *HelmDeploymentRestrictionResponseList
	isSet bool
}

func (v NullableHelmDeploymentRestrictionResponseList) Get() *HelmDeploymentRestrictionResponseList {
	return v.value
}

func (v *NullableHelmDeploymentRestrictionResponseList) Set(val *HelmDeploymentRestrictionResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmDeploymentRestrictionResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmDeploymentRestrictionResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmDeploymentRestrictionResponseList(val *HelmDeploymentRestrictionResponseList) *NullableHelmDeploymentRestrictionResponseList {
	return &NullableHelmDeploymentRestrictionResponseList{value: val, isSet: true}
}

func (v NullableHelmDeploymentRestrictionResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmDeploymentRestrictionResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
