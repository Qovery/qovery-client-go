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

// VariableResponseList struct for VariableResponseList
type VariableResponseList struct {
	Results []VariableResponse `json:"results,omitempty"`
}

// NewVariableResponseList instantiates a new VariableResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableResponseList() *VariableResponseList {
	this := VariableResponseList{}
	return &this
}

// NewVariableResponseListWithDefaults instantiates a new VariableResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableResponseListWithDefaults() *VariableResponseList {
	this := VariableResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *VariableResponseList) GetResults() []VariableResponse {
	if o == nil || o.Results == nil {
		var ret []VariableResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseList) GetResultsOk() ([]VariableResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *VariableResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []VariableResponse and assigns it to the Results field.
func (o *VariableResponseList) SetResults(v []VariableResponse) {
	o.Results = v
}

func (o VariableResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableVariableResponseList struct {
	value *VariableResponseList
	isSet bool
}

func (v NullableVariableResponseList) Get() *VariableResponseList {
	return v.value
}

func (v *NullableVariableResponseList) Set(val *VariableResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableResponseList(val *VariableResponseList) *NullableVariableResponseList {
	return &NullableVariableResponseList{value: val, isSet: true}
}

func (v NullableVariableResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
