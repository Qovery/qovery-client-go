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

// ApplicationNetworkResponse struct for ApplicationNetworkResponse
type ApplicationNetworkResponse struct {
	// Specify if the sticky session option (also called persistant session) is activated or not for this application. If activated, user will be redirected by the load balancer to the same instance each time he access to the application.
	StickySession *bool `json:"sticky_session,omitempty"`
}

// NewApplicationNetworkResponse instantiates a new ApplicationNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationNetworkResponse() *ApplicationNetworkResponse {
	this := ApplicationNetworkResponse{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// NewApplicationNetworkResponseWithDefaults instantiates a new ApplicationNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationNetworkResponseWithDefaults() *ApplicationNetworkResponse {
	this := ApplicationNetworkResponse{}
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// GetStickySession returns the StickySession field value if set, zero value otherwise.
func (o *ApplicationNetworkResponse) GetStickySession() bool {
	if o == nil || o.StickySession == nil {
		var ret bool
		return ret
	}
	return *o.StickySession
}

// GetStickySessionOk returns a tuple with the StickySession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationNetworkResponse) GetStickySessionOk() (*bool, bool) {
	if o == nil || o.StickySession == nil {
		return nil, false
	}
	return o.StickySession, true
}

// HasStickySession returns a boolean if a field has been set.
func (o *ApplicationNetworkResponse) HasStickySession() bool {
	if o != nil && o.StickySession != nil {
		return true
	}

	return false
}

// SetStickySession gets a reference to the given bool and assigns it to the StickySession field.
func (o *ApplicationNetworkResponse) SetStickySession(v bool) {
	o.StickySession = &v
}

func (o ApplicationNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StickySession != nil {
		toSerialize["sticky_session"] = o.StickySession
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationNetworkResponse struct {
	value *ApplicationNetworkResponse
	isSet bool
}

func (v NullableApplicationNetworkResponse) Get() *ApplicationNetworkResponse {
	return v.value
}

func (v *NullableApplicationNetworkResponse) Set(val *ApplicationNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationNetworkResponse(val *ApplicationNetworkResponse) *NullableApplicationNetworkResponse {
	return &NullableApplicationNetworkResponse{value: val, isSet: true}
}

func (v NullableApplicationNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
