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

// ManagedDatabaseTypeResponseList struct for ManagedDatabaseTypeResponseList
type ManagedDatabaseTypeResponseList struct {
	Results []ManagedDatabaseTypeResponse `json:"results,omitempty"`
}

// NewManagedDatabaseTypeResponseList instantiates a new ManagedDatabaseTypeResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedDatabaseTypeResponseList() *ManagedDatabaseTypeResponseList {
	this := ManagedDatabaseTypeResponseList{}
	return &this
}

// NewManagedDatabaseTypeResponseListWithDefaults instantiates a new ManagedDatabaseTypeResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedDatabaseTypeResponseListWithDefaults() *ManagedDatabaseTypeResponseList {
	this := ManagedDatabaseTypeResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ManagedDatabaseTypeResponseList) GetResults() []ManagedDatabaseTypeResponse {
	if o == nil || o.Results == nil {
		var ret []ManagedDatabaseTypeResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDatabaseTypeResponseList) GetResultsOk() ([]ManagedDatabaseTypeResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ManagedDatabaseTypeResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ManagedDatabaseTypeResponse and assigns it to the Results field.
func (o *ManagedDatabaseTypeResponseList) SetResults(v []ManagedDatabaseTypeResponse) {
	o.Results = v
}

func (o ManagedDatabaseTypeResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableManagedDatabaseTypeResponseList struct {
	value *ManagedDatabaseTypeResponseList
	isSet bool
}

func (v NullableManagedDatabaseTypeResponseList) Get() *ManagedDatabaseTypeResponseList {
	return v.value
}

func (v *NullableManagedDatabaseTypeResponseList) Set(val *ManagedDatabaseTypeResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedDatabaseTypeResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedDatabaseTypeResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedDatabaseTypeResponseList(val *ManagedDatabaseTypeResponseList) *NullableManagedDatabaseTypeResponseList {
	return &NullableManagedDatabaseTypeResponseList{value: val, isSet: true}
}

func (v NullableManagedDatabaseTypeResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedDatabaseTypeResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
