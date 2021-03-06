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

// EnvironmentAllOfCloudProvider struct for EnvironmentAllOfCloudProvider
type EnvironmentAllOfCloudProvider struct {
	Provider *string `json:"provider,omitempty"`
	Cluster  *string `json:"cluster,omitempty"`
}

// NewEnvironmentAllOfCloudProvider instantiates a new EnvironmentAllOfCloudProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentAllOfCloudProvider() *EnvironmentAllOfCloudProvider {
	this := EnvironmentAllOfCloudProvider{}
	return &this
}

// NewEnvironmentAllOfCloudProviderWithDefaults instantiates a new EnvironmentAllOfCloudProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentAllOfCloudProviderWithDefaults() *EnvironmentAllOfCloudProvider {
	this := EnvironmentAllOfCloudProvider{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *EnvironmentAllOfCloudProvider) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAllOfCloudProvider) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *EnvironmentAllOfCloudProvider) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *EnvironmentAllOfCloudProvider) SetProvider(v string) {
	o.Provider = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *EnvironmentAllOfCloudProvider) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAllOfCloudProvider) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *EnvironmentAllOfCloudProvider) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *EnvironmentAllOfCloudProvider) SetCluster(v string) {
	o.Cluster = &v
}

func (o EnvironmentAllOfCloudProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentAllOfCloudProvider struct {
	value *EnvironmentAllOfCloudProvider
	isSet bool
}

func (v NullableEnvironmentAllOfCloudProvider) Get() *EnvironmentAllOfCloudProvider {
	return v.value
}

func (v *NullableEnvironmentAllOfCloudProvider) Set(val *EnvironmentAllOfCloudProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentAllOfCloudProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentAllOfCloudProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentAllOfCloudProvider(val *EnvironmentAllOfCloudProvider) *NullableEnvironmentAllOfCloudProvider {
	return &NullableEnvironmentAllOfCloudProvider{value: val, isSet: true}
}

func (v NullableEnvironmentAllOfCloudProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentAllOfCloudProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
