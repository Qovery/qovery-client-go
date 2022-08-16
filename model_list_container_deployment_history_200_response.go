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

// ListContainerDeploymentHistory200Response struct for ListContainerDeploymentHistory200Response
type ListContainerDeploymentHistory200Response struct {
	Page     float32                                                      `json:"page"`
	PageSize float32                                                      `json:"page_size"`
	Results  []ListContainerDeploymentHistory200ResponseAllOfResultsInner `json:"results,omitempty"`
}

// NewListContainerDeploymentHistory200Response instantiates a new ListContainerDeploymentHistory200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContainerDeploymentHistory200Response(page float32, pageSize float32) *ListContainerDeploymentHistory200Response {
	this := ListContainerDeploymentHistory200Response{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewListContainerDeploymentHistory200ResponseWithDefaults instantiates a new ListContainerDeploymentHistory200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContainerDeploymentHistory200ResponseWithDefaults() *ListContainerDeploymentHistory200Response {
	this := ListContainerDeploymentHistory200Response{}
	return &this
}

// GetPage returns the Page field value
func (o *ListContainerDeploymentHistory200Response) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ListContainerDeploymentHistory200Response) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ListContainerDeploymentHistory200Response) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *ListContainerDeploymentHistory200Response) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListContainerDeploymentHistory200Response) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListContainerDeploymentHistory200Response) SetPageSize(v float32) {
	o.PageSize = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListContainerDeploymentHistory200Response) GetResults() []ListContainerDeploymentHistory200ResponseAllOfResultsInner {
	if o == nil || o.Results == nil {
		var ret []ListContainerDeploymentHistory200ResponseAllOfResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainerDeploymentHistory200Response) GetResultsOk() ([]ListContainerDeploymentHistory200ResponseAllOfResultsInner, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListContainerDeploymentHistory200Response) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ListContainerDeploymentHistory200ResponseAllOfResultsInner and assigns it to the Results field.
func (o *ListContainerDeploymentHistory200Response) SetResults(v []ListContainerDeploymentHistory200ResponseAllOfResultsInner) {
	o.Results = v
}

func (o ListContainerDeploymentHistory200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableListContainerDeploymentHistory200Response struct {
	value *ListContainerDeploymentHistory200Response
	isSet bool
}

func (v NullableListContainerDeploymentHistory200Response) Get() *ListContainerDeploymentHistory200Response {
	return v.value
}

func (v *NullableListContainerDeploymentHistory200Response) Set(val *ListContainerDeploymentHistory200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListContainerDeploymentHistory200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListContainerDeploymentHistory200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContainerDeploymentHistory200Response(val *ListContainerDeploymentHistory200Response) *NullableListContainerDeploymentHistory200Response {
	return &NullableListContainerDeploymentHistory200Response{value: val, isSet: true}
}

func (v NullableListContainerDeploymentHistory200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContainerDeploymentHistory200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
