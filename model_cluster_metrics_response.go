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

// checks if the ClusterMetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterMetricsResponse{}

// ClusterMetricsResponse struct for ClusterMetricsResponse
type ClusterMetricsResponse struct {
	Metrics              *string `json:"metrics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClusterMetricsResponse ClusterMetricsResponse

// NewClusterMetricsResponse instantiates a new ClusterMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterMetricsResponse() *ClusterMetricsResponse {
	this := ClusterMetricsResponse{}
	return &this
}

// NewClusterMetricsResponseWithDefaults instantiates a new ClusterMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterMetricsResponseWithDefaults() *ClusterMetricsResponse {
	this := ClusterMetricsResponse{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ClusterMetricsResponse) GetMetrics() string {
	if o == nil || IsNil(o.Metrics) {
		var ret string
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMetricsResponse) GetMetricsOk() (*string, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ClusterMetricsResponse) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given string and assigns it to the Metrics field.
func (o *ClusterMetricsResponse) SetMetrics(v string) {
	o.Metrics = &v
}

func (o ClusterMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterMetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	varClusterMetricsResponse := _ClusterMetricsResponse{}

	err = json.Unmarshal(data, &varClusterMetricsResponse)

	if err != nil {
		return err
	}

	*o = ClusterMetricsResponse(varClusterMetricsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metrics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterMetricsResponse struct {
	value *ClusterMetricsResponse
	isSet bool
}

func (v NullableClusterMetricsResponse) Get() *ClusterMetricsResponse {
	return v.value
}

func (v *NullableClusterMetricsResponse) Set(val *ClusterMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterMetricsResponse(val *ClusterMetricsResponse) *NullableClusterMetricsResponse {
	return &NullableClusterMetricsResponse{value: val, isSet: true}
}

func (v NullableClusterMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
