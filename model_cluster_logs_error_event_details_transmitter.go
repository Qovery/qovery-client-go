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

// ClusterLogsErrorEventDetailsTransmitter struct for ClusterLogsErrorEventDetailsTransmitter
type ClusterLogsErrorEventDetailsTransmitter struct {
	Type *string `json:"type,omitempty"`
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewClusterLogsErrorEventDetailsTransmitter instantiates a new ClusterLogsErrorEventDetailsTransmitter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogsErrorEventDetailsTransmitter() *ClusterLogsErrorEventDetailsTransmitter {
	this := ClusterLogsErrorEventDetailsTransmitter{}
	return &this
}

// NewClusterLogsErrorEventDetailsTransmitterWithDefaults instantiates a new ClusterLogsErrorEventDetailsTransmitter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsErrorEventDetailsTransmitterWithDefaults() *ClusterLogsErrorEventDetailsTransmitter {
	this := ClusterLogsErrorEventDetailsTransmitter{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetailsTransmitter) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetailsTransmitter) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetailsTransmitter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterLogsErrorEventDetailsTransmitter) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetailsTransmitter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetailsTransmitter) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetailsTransmitter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterLogsErrorEventDetailsTransmitter) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetailsTransmitter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetailsTransmitter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetailsTransmitter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterLogsErrorEventDetailsTransmitter) SetName(v string) {
	o.Name = &v
}

func (o ClusterLogsErrorEventDetailsTransmitter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableClusterLogsErrorEventDetailsTransmitter struct {
	value *ClusterLogsErrorEventDetailsTransmitter
	isSet bool
}

func (v NullableClusterLogsErrorEventDetailsTransmitter) Get() *ClusterLogsErrorEventDetailsTransmitter {
	return v.value
}

func (v *NullableClusterLogsErrorEventDetailsTransmitter) Set(val *ClusterLogsErrorEventDetailsTransmitter) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogsErrorEventDetailsTransmitter) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogsErrorEventDetailsTransmitter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogsErrorEventDetailsTransmitter(val *ClusterLogsErrorEventDetailsTransmitter) *NullableClusterLogsErrorEventDetailsTransmitter {
	return &NullableClusterLogsErrorEventDetailsTransmitter{value: val, isSet: true}
}

func (v NullableClusterLogsErrorEventDetailsTransmitter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogsErrorEventDetailsTransmitter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
