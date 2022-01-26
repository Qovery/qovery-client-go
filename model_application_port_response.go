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

// ApplicationPortResponse struct for ApplicationPortResponse
type ApplicationPortResponse struct {
	Ports *[]ApplicationPortResponsePorts `json:"ports,omitempty"`
}

// NewApplicationPortResponse instantiates a new ApplicationPortResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPortResponse() *ApplicationPortResponse {
	this := ApplicationPortResponse{}
	return &this
}

// NewApplicationPortResponseWithDefaults instantiates a new ApplicationPortResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPortResponseWithDefaults() *ApplicationPortResponse {
	this := ApplicationPortResponse{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationPortResponse) GetPorts() []ApplicationPortResponsePorts {
	if o == nil || o.Ports == nil {
		var ret []ApplicationPortResponsePorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortResponse) GetPortsOk() (*[]ApplicationPortResponsePorts, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationPortResponse) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApplicationPortResponsePorts and assigns it to the Ports field.
func (o *ApplicationPortResponse) SetPorts(v []ApplicationPortResponsePorts) {
	o.Ports = &v
}

func (o ApplicationPortResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPortResponse struct {
	value *ApplicationPortResponse
	isSet bool
}

func (v NullableApplicationPortResponse) Get() *ApplicationPortResponse {
	return v.value
}

func (v *NullableApplicationPortResponse) Set(val *ApplicationPortResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPortResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPortResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPortResponse(val *ApplicationPortResponse) *NullableApplicationPortResponse {
	return &NullableApplicationPortResponse{value: val, isSet: true}
}

func (v NullableApplicationPortResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPortResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
