/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.2
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// LogResponseList struct for LogResponseList
type LogResponseList struct {
	Results []LogResponse `json:"results,omitempty"`
}

// NewLogResponseList instantiates a new LogResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResponseList() *LogResponseList {
	this := LogResponseList{}
	return &this
}

// NewLogResponseListWithDefaults instantiates a new LogResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResponseListWithDefaults() *LogResponseList {
	this := LogResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LogResponseList) GetResults() []LogResponse {
	if o == nil || o.Results == nil {
		var ret []LogResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponseList) GetResultsOk() ([]LogResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LogResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LogResponse and assigns it to the Results field.
func (o *LogResponseList) SetResults(v []LogResponse) {
	o.Results = v
}

func (o LogResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableLogResponseList struct {
	value *LogResponseList
	isSet bool
}

func (v NullableLogResponseList) Get() *LogResponseList {
	return v.value
}

func (v *NullableLogResponseList) Set(val *LogResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResponseList(val *LogResponseList) *NullableLogResponseList {
	return &NullableLogResponseList{value: val, isSet: true}
}

func (v NullableLogResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
