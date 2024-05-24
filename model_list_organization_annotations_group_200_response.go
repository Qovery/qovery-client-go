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

// checks if the ListOrganizationAnnotationsGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationAnnotationsGroup200Response{}

// ListOrganizationAnnotationsGroup200Response struct for ListOrganizationAnnotationsGroup200Response
type ListOrganizationAnnotationsGroup200Response struct {
	Results []OrganizationAnnotationsGroupEnrichedResponse `json:"results,omitempty"`
}

// NewListOrganizationAnnotationsGroup200Response instantiates a new ListOrganizationAnnotationsGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationAnnotationsGroup200Response() *ListOrganizationAnnotationsGroup200Response {
	this := ListOrganizationAnnotationsGroup200Response{}
	return &this
}

// NewListOrganizationAnnotationsGroup200ResponseWithDefaults instantiates a new ListOrganizationAnnotationsGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationAnnotationsGroup200ResponseWithDefaults() *ListOrganizationAnnotationsGroup200Response {
	this := ListOrganizationAnnotationsGroup200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListOrganizationAnnotationsGroup200Response) GetResults() []OrganizationAnnotationsGroupEnrichedResponse {
	if o == nil || IsNil(o.Results) {
		var ret []OrganizationAnnotationsGroupEnrichedResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationAnnotationsGroup200Response) GetResultsOk() ([]OrganizationAnnotationsGroupEnrichedResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListOrganizationAnnotationsGroup200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OrganizationAnnotationsGroupEnrichedResponse and assigns it to the Results field.
func (o *ListOrganizationAnnotationsGroup200Response) SetResults(v []OrganizationAnnotationsGroupEnrichedResponse) {
	o.Results = v
}

func (o ListOrganizationAnnotationsGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationAnnotationsGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableListOrganizationAnnotationsGroup200Response struct {
	value *ListOrganizationAnnotationsGroup200Response
	isSet bool
}

func (v NullableListOrganizationAnnotationsGroup200Response) Get() *ListOrganizationAnnotationsGroup200Response {
	return v.value
}

func (v *NullableListOrganizationAnnotationsGroup200Response) Set(val *ListOrganizationAnnotationsGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationAnnotationsGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationAnnotationsGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationAnnotationsGroup200Response(val *ListOrganizationAnnotationsGroup200Response) *NullableListOrganizationAnnotationsGroup200Response {
	return &NullableListOrganizationAnnotationsGroup200Response{value: val, isSet: true}
}

func (v NullableListOrganizationAnnotationsGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationAnnotationsGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}