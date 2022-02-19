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

// OrganizationResponseList struct for OrganizationResponseList
type OrganizationResponseList struct {
	Results []OrganizationResponse `json:"results,omitempty"`
}

// NewOrganizationResponseList instantiates a new OrganizationResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationResponseList() *OrganizationResponseList {
	this := OrganizationResponseList{}
	return &this
}

// NewOrganizationResponseListWithDefaults instantiates a new OrganizationResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationResponseListWithDefaults() *OrganizationResponseList {
	this := OrganizationResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OrganizationResponseList) GetResults() []OrganizationResponse {
	if o == nil || o.Results == nil {
		var ret []OrganizationResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationResponseList) GetResultsOk() ([]OrganizationResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OrganizationResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OrganizationResponse and assigns it to the Results field.
func (o *OrganizationResponseList) SetResults(v []OrganizationResponse) {
	o.Results = v
}

func (o OrganizationResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationResponseList struct {
	value *OrganizationResponseList
	isSet bool
}

func (v NullableOrganizationResponseList) Get() *OrganizationResponseList {
	return v.value
}

func (v *NullableOrganizationResponseList) Set(val *OrganizationResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationResponseList(val *OrganizationResponseList) *NullableOrganizationResponseList {
	return &NullableOrganizationResponseList{value: val, isSet: true}
}

func (v NullableOrganizationResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
