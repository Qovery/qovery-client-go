/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// MetricRestartResponse struct for MetricRestartResponse
type MetricRestartResponse struct {
	Results []MetricRestartResponseResults `json:"results,omitempty"`
}

// NewMetricRestartResponse instantiates a new MetricRestartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricRestartResponse() *MetricRestartResponse {
	this := MetricRestartResponse{}
	return &this
}

// NewMetricRestartResponseWithDefaults instantiates a new MetricRestartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricRestartResponseWithDefaults() *MetricRestartResponse {
	this := MetricRestartResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MetricRestartResponse) GetResults() []MetricRestartResponseResults {
	if o == nil || o.Results == nil {
		var ret []MetricRestartResponseResults
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRestartResponse) GetResultsOk() ([]MetricRestartResponseResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MetricRestartResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MetricRestartResponseResults and assigns it to the Results field.
func (o *MetricRestartResponse) SetResults(v []MetricRestartResponseResults) {
	o.Results = v
}

func (o MetricRestartResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableMetricRestartResponse struct {
	value *MetricRestartResponse
	isSet bool
}

func (v NullableMetricRestartResponse) Get() *MetricRestartResponse {
	return v.value
}

func (v *NullableMetricRestartResponse) Set(val *MetricRestartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricRestartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricRestartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricRestartResponse(val *MetricRestartResponse) *NullableMetricRestartResponse {
	return &NullableMetricRestartResponse{value: val, isSet: true}
}

func (v NullableMetricRestartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricRestartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
