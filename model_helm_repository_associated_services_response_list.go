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

// checks if the HelmRepositoryAssociatedServicesResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRepositoryAssociatedServicesResponseList{}

// HelmRepositoryAssociatedServicesResponseList struct for HelmRepositoryAssociatedServicesResponseList
type HelmRepositoryAssociatedServicesResponseList struct {
	Results              []HelmRepositoryAssociatedServicesResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmRepositoryAssociatedServicesResponseList HelmRepositoryAssociatedServicesResponseList

// NewHelmRepositoryAssociatedServicesResponseList instantiates a new HelmRepositoryAssociatedServicesResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRepositoryAssociatedServicesResponseList() *HelmRepositoryAssociatedServicesResponseList {
	this := HelmRepositoryAssociatedServicesResponseList{}
	return &this
}

// NewHelmRepositoryAssociatedServicesResponseListWithDefaults instantiates a new HelmRepositoryAssociatedServicesResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRepositoryAssociatedServicesResponseListWithDefaults() *HelmRepositoryAssociatedServicesResponseList {
	this := HelmRepositoryAssociatedServicesResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HelmRepositoryAssociatedServicesResponseList) GetResults() []HelmRepositoryAssociatedServicesResponse {
	if o == nil || IsNil(o.Results) {
		var ret []HelmRepositoryAssociatedServicesResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRepositoryAssociatedServicesResponseList) GetResultsOk() ([]HelmRepositoryAssociatedServicesResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HelmRepositoryAssociatedServicesResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HelmRepositoryAssociatedServicesResponse and assigns it to the Results field.
func (o *HelmRepositoryAssociatedServicesResponseList) SetResults(v []HelmRepositoryAssociatedServicesResponse) {
	o.Results = v
}

func (o HelmRepositoryAssociatedServicesResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRepositoryAssociatedServicesResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRepositoryAssociatedServicesResponseList) UnmarshalJSON(data []byte) (err error) {
	varHelmRepositoryAssociatedServicesResponseList := _HelmRepositoryAssociatedServicesResponseList{}

	err = json.Unmarshal(data, &varHelmRepositoryAssociatedServicesResponseList)

	if err != nil {
		return err
	}

	*o = HelmRepositoryAssociatedServicesResponseList(varHelmRepositoryAssociatedServicesResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRepositoryAssociatedServicesResponseList struct {
	value *HelmRepositoryAssociatedServicesResponseList
	isSet bool
}

func (v NullableHelmRepositoryAssociatedServicesResponseList) Get() *HelmRepositoryAssociatedServicesResponseList {
	return v.value
}

func (v *NullableHelmRepositoryAssociatedServicesResponseList) Set(val *HelmRepositoryAssociatedServicesResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRepositoryAssociatedServicesResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRepositoryAssociatedServicesResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRepositoryAssociatedServicesResponseList(val *HelmRepositoryAssociatedServicesResponseList) *NullableHelmRepositoryAssociatedServicesResponseList {
	return &NullableHelmRepositoryAssociatedServicesResponseList{value: val, isSet: true}
}

func (v NullableHelmRepositoryAssociatedServicesResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRepositoryAssociatedServicesResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
