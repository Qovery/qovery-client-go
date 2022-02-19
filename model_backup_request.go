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

// BackupRequest struct for BackupRequest
type BackupRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// NewBackupRequest instantiates a new BackupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRequest(name string, message string) *BackupRequest {
	this := BackupRequest{}
	this.Name = name
	this.Message = message
	return &this
}

// NewBackupRequestWithDefaults instantiates a new BackupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRequestWithDefaults() *BackupRequest {
	this := BackupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BackupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BackupRequest) SetName(v string) {
	o.Name = v
}

// GetMessage returns the Message field value
func (o *BackupRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BackupRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BackupRequest) SetMessage(v string) {
	o.Message = v
}

func (o BackupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableBackupRequest struct {
	value *BackupRequest
	isSet bool
}

func (v NullableBackupRequest) Get() *BackupRequest {
	return v.value
}

func (v *NullableBackupRequest) Set(val *BackupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRequest(val *BackupRequest) *NullableBackupRequest {
	return &NullableBackupRequest{value: val, isSet: true}
}

func (v NullableBackupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
