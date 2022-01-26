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

// CommunityUsage struct for CommunityUsage
type CommunityUsage struct {
	CommunityUsage *CommunityUsageResponse `json:"community_usage,omitempty"`
}

// NewCommunityUsage instantiates a new CommunityUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityUsage() *CommunityUsage {
	this := CommunityUsage{}
	return &this
}

// NewCommunityUsageWithDefaults instantiates a new CommunityUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityUsageWithDefaults() *CommunityUsage {
	this := CommunityUsage{}
	return &this
}

// GetCommunityUsage returns the CommunityUsage field value if set, zero value otherwise.
func (o *CommunityUsage) GetCommunityUsage() CommunityUsageResponse {
	if o == nil || o.CommunityUsage == nil {
		var ret CommunityUsageResponse
		return ret
	}
	return *o.CommunityUsage
}

// GetCommunityUsageOk returns a tuple with the CommunityUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityUsage) GetCommunityUsageOk() (*CommunityUsageResponse, bool) {
	if o == nil || o.CommunityUsage == nil {
		return nil, false
	}
	return o.CommunityUsage, true
}

// HasCommunityUsage returns a boolean if a field has been set.
func (o *CommunityUsage) HasCommunityUsage() bool {
	if o != nil && o.CommunityUsage != nil {
		return true
	}

	return false
}

// SetCommunityUsage gets a reference to the given CommunityUsageResponse and assigns it to the CommunityUsage field.
func (o *CommunityUsage) SetCommunityUsage(v CommunityUsageResponse) {
	o.CommunityUsage = &v
}

func (o CommunityUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommunityUsage != nil {
		toSerialize["community_usage"] = o.CommunityUsage
	}
	return json.Marshal(toSerialize)
}

type NullableCommunityUsage struct {
	value *CommunityUsage
	isSet bool
}

func (v NullableCommunityUsage) Get() *CommunityUsage {
	return v.value
}

func (v *NullableCommunityUsage) Set(val *CommunityUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityUsage(val *CommunityUsage) *NullableCommunityUsage {
	return &NullableCommunityUsage{value: val, isSet: true}
}

func (v NullableCommunityUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
