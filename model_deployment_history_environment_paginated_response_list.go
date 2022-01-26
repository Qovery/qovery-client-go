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

// DeploymentHistoryEnvironmentPaginatedResponseList struct for DeploymentHistoryEnvironmentPaginatedResponseList
type DeploymentHistoryEnvironmentPaginatedResponseList struct {
	Results  *[]DeploymentHistoryEnvironmentResponse `json:"results,omitempty"`
	Page     float32                                 `json:"page"`
	PageSize float32                                 `json:"page_size"`
}

// NewDeploymentHistoryEnvironmentPaginatedResponseList instantiates a new DeploymentHistoryEnvironmentPaginatedResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryEnvironmentPaginatedResponseList(page float32, pageSize float32) *DeploymentHistoryEnvironmentPaginatedResponseList {
	this := DeploymentHistoryEnvironmentPaginatedResponseList{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewDeploymentHistoryEnvironmentPaginatedResponseListWithDefaults instantiates a new DeploymentHistoryEnvironmentPaginatedResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryEnvironmentPaginatedResponseListWithDefaults() *DeploymentHistoryEnvironmentPaginatedResponseList {
	this := DeploymentHistoryEnvironmentPaginatedResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetResults() []DeploymentHistoryEnvironmentResponse {
	if o == nil || o.Results == nil {
		var ret []DeploymentHistoryEnvironmentResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetResultsOk() (*[]DeploymentHistoryEnvironmentResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeploymentHistoryEnvironmentResponse and assigns it to the Results field.
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) SetResults(v []DeploymentHistoryEnvironmentResponse) {
	o.Results = &v
}

// GetPage returns the Page field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseList) SetPageSize(v float32) {
	o.PageSize = v
}

func (o DeploymentHistoryEnvironmentPaginatedResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryEnvironmentPaginatedResponseList struct {
	value *DeploymentHistoryEnvironmentPaginatedResponseList
	isSet bool
}

func (v NullableDeploymentHistoryEnvironmentPaginatedResponseList) Get() *DeploymentHistoryEnvironmentPaginatedResponseList {
	return v.value
}

func (v *NullableDeploymentHistoryEnvironmentPaginatedResponseList) Set(val *DeploymentHistoryEnvironmentPaginatedResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryEnvironmentPaginatedResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryEnvironmentPaginatedResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryEnvironmentPaginatedResponseList(val *DeploymentHistoryEnvironmentPaginatedResponseList) *NullableDeploymentHistoryEnvironmentPaginatedResponseList {
	return &NullableDeploymentHistoryEnvironmentPaginatedResponseList{value: val, isSet: true}
}

func (v NullableDeploymentHistoryEnvironmentPaginatedResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryEnvironmentPaginatedResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
