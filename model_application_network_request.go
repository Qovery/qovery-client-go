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

// ApplicationNetworkRequest struct for ApplicationNetworkRequest
type ApplicationNetworkRequest struct {
	// Specify if the sticky session option (also called persistant session) is activated or not for this application. If activated, user will be redirected by the load balancer to the same instance each time he access to the application.
	StickySession *bool `json:"sticky-session,omitempty"`
}

// NewApplicationNetworkRequest instantiates a new ApplicationNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationNetworkRequest() *ApplicationNetworkRequest {
	this := ApplicationNetworkRequest{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// NewApplicationNetworkRequestWithDefaults instantiates a new ApplicationNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationNetworkRequestWithDefaults() *ApplicationNetworkRequest {
	this := ApplicationNetworkRequest{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// GetStickySession returns the StickySession field value if set, zero value otherwise.
func (o *ApplicationNetworkRequest) GetStickySession() bool {
	if o == nil || o.StickySession == nil {
		var ret bool
		return ret
	}
	return *o.StickySession
}

// GetStickySessionOk returns a tuple with the StickySession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationNetworkRequest) GetStickySessionOk() (*bool, bool) {
	if o == nil || o.StickySession == nil {
		return nil, false
	}
	return o.StickySession, true
}

// HasStickySession returns a boolean if a field has been set.
func (o *ApplicationNetworkRequest) HasStickySession() bool {
	if o != nil && o.StickySession != nil {
		return true
	}

	return false
}

// SetStickySession gets a reference to the given bool and assigns it to the StickySession field.
func (o *ApplicationNetworkRequest) SetStickySession(v bool) {
	o.StickySession = &v
}

func (o ApplicationNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StickySession != nil {
		toSerialize["sticky-session"] = o.StickySession
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationNetworkRequest struct {
	value *ApplicationNetworkRequest
	isSet bool
}

func (v NullableApplicationNetworkRequest) Get() *ApplicationNetworkRequest {
	return v.value
}

func (v *NullableApplicationNetworkRequest) Set(val *ApplicationNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationNetworkRequest(val *ApplicationNetworkRequest) *NullableApplicationNetworkRequest {
	return &NullableApplicationNetworkRequest{value: val, isSet: true}
}

func (v NullableApplicationNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
