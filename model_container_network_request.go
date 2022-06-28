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

// ContainerNetworkRequest struct for ContainerNetworkRequest
type ContainerNetworkRequest struct {
	// Specify if the sticky session option (also called persistant session) is activated or not for this container. If activated, user will be redirected by the load balancer to the same instance each time he access to the container.
	StickySession *bool `json:"sticky_session,omitempty"`
}

// NewContainerNetworkRequest instantiates a new ContainerNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerNetworkRequest() *ContainerNetworkRequest {
	this := ContainerNetworkRequest{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// NewContainerNetworkRequestWithDefaults instantiates a new ContainerNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerNetworkRequestWithDefaults() *ContainerNetworkRequest {
	this := ContainerNetworkRequest{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// GetStickySession returns the StickySession field value if set, zero value otherwise.
func (o *ContainerNetworkRequest) GetStickySession() bool {
	if o == nil || o.StickySession == nil {
		var ret bool
		return ret
	}
	return *o.StickySession
}

// GetStickySessionOk returns a tuple with the StickySession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerNetworkRequest) GetStickySessionOk() (*bool, bool) {
	if o == nil || o.StickySession == nil {
		return nil, false
	}
	return o.StickySession, true
}

// HasStickySession returns a boolean if a field has been set.
func (o *ContainerNetworkRequest) HasStickySession() bool {
	if o != nil && o.StickySession != nil {
		return true
	}

	return false
}

// SetStickySession gets a reference to the given bool and assigns it to the StickySession field.
func (o *ContainerNetworkRequest) SetStickySession(v bool) {
	o.StickySession = &v
}

func (o ContainerNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StickySession != nil {
		toSerialize["sticky_session"] = o.StickySession
	}
	return json.Marshal(toSerialize)
}

type NullableContainerNetworkRequest struct {
	value *ContainerNetworkRequest
	isSet bool
}

func (v NullableContainerNetworkRequest) Get() *ContainerNetworkRequest {
	return v.value
}

func (v *NullableContainerNetworkRequest) Set(val *ContainerNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerNetworkRequest(val *ContainerNetworkRequest) *NullableContainerNetworkRequest {
	return &NullableContainerNetworkRequest{value: val, isSet: true}
}

func (v NullableContainerNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}