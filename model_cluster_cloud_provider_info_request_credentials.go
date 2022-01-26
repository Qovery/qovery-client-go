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

// ClusterCloudProviderInfoRequestCredentials struct for ClusterCloudProviderInfoRequestCredentials
type ClusterCloudProviderInfoRequestCredentials struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewClusterCloudProviderInfoRequestCredentials instantiates a new ClusterCloudProviderInfoRequestCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCloudProviderInfoRequestCredentials() *ClusterCloudProviderInfoRequestCredentials {
	this := ClusterCloudProviderInfoRequestCredentials{}
	return &this
}

// NewClusterCloudProviderInfoRequestCredentialsWithDefaults instantiates a new ClusterCloudProviderInfoRequestCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCloudProviderInfoRequestCredentialsWithDefaults() *ClusterCloudProviderInfoRequestCredentials {
	this := ClusterCloudProviderInfoRequestCredentials{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterCloudProviderInfoRequestCredentials) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCloudProviderInfoRequestCredentials) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterCloudProviderInfoRequestCredentials) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterCloudProviderInfoRequestCredentials) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterCloudProviderInfoRequestCredentials) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCloudProviderInfoRequestCredentials) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterCloudProviderInfoRequestCredentials) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterCloudProviderInfoRequestCredentials) SetName(v string) {
	o.Name = &v
}

func (o ClusterCloudProviderInfoRequestCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableClusterCloudProviderInfoRequestCredentials struct {
	value *ClusterCloudProviderInfoRequestCredentials
	isSet bool
}

func (v NullableClusterCloudProviderInfoRequestCredentials) Get() *ClusterCloudProviderInfoRequestCredentials {
	return v.value
}

func (v *NullableClusterCloudProviderInfoRequestCredentials) Set(val *ClusterCloudProviderInfoRequestCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCloudProviderInfoRequestCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCloudProviderInfoRequestCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCloudProviderInfoRequestCredentials(val *ClusterCloudProviderInfoRequestCredentials) *NullableClusterCloudProviderInfoRequestCredentials {
	return &NullableClusterCloudProviderInfoRequestCredentials{value: val, isSet: true}
}

func (v NullableClusterCloudProviderInfoRequestCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCloudProviderInfoRequestCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
