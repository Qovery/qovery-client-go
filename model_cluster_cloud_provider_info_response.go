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

// ClusterCloudProviderInfoResponse struct for ClusterCloudProviderInfoResponse
type ClusterCloudProviderInfoResponse struct {
	CloudProvider *string                                     `json:"cloud_provider,omitempty"`
	Credentials   *ClusterCloudProviderInfoRequestCredentials `json:"credentials,omitempty"`
	Region        *string                                     `json:"region,omitempty"`
}

// NewClusterCloudProviderInfoResponse instantiates a new ClusterCloudProviderInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCloudProviderInfoResponse() *ClusterCloudProviderInfoResponse {
	this := ClusterCloudProviderInfoResponse{}
	return &this
}

// NewClusterCloudProviderInfoResponseWithDefaults instantiates a new ClusterCloudProviderInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCloudProviderInfoResponseWithDefaults() *ClusterCloudProviderInfoResponse {
	this := ClusterCloudProviderInfoResponse{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ClusterCloudProviderInfoResponse) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCloudProviderInfoResponse) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ClusterCloudProviderInfoResponse) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ClusterCloudProviderInfoResponse) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ClusterCloudProviderInfoResponse) GetCredentials() ClusterCloudProviderInfoRequestCredentials {
	if o == nil || o.Credentials == nil {
		var ret ClusterCloudProviderInfoRequestCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCloudProviderInfoResponse) GetCredentialsOk() (*ClusterCloudProviderInfoRequestCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ClusterCloudProviderInfoResponse) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ClusterCloudProviderInfoRequestCredentials and assigns it to the Credentials field.
func (o *ClusterCloudProviderInfoResponse) SetCredentials(v ClusterCloudProviderInfoRequestCredentials) {
	o.Credentials = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClusterCloudProviderInfoResponse) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCloudProviderInfoResponse) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterCloudProviderInfoResponse) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterCloudProviderInfoResponse) SetRegion(v string) {
	o.Region = &v
}

func (o ClusterCloudProviderInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableClusterCloudProviderInfoResponse struct {
	value *ClusterCloudProviderInfoResponse
	isSet bool
}

func (v NullableClusterCloudProviderInfoResponse) Get() *ClusterCloudProviderInfoResponse {
	return v.value
}

func (v *NullableClusterCloudProviderInfoResponse) Set(val *ClusterCloudProviderInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCloudProviderInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCloudProviderInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCloudProviderInfoResponse(val *ClusterCloudProviderInfoResponse) *NullableClusterCloudProviderInfoResponse {
	return &NullableClusterCloudProviderInfoResponse{value: val, isSet: true}
}

func (v NullableClusterCloudProviderInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCloudProviderInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
