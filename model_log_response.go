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
	"time"
)

// LogResponse struct for LogResponse
type LogResponse struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
}

// NewLogResponse instantiates a new LogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResponse(id string, createdAt time.Time, message string) *LogResponse {
	this := LogResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Message = message
	return &this
}

// NewLogResponseWithDefaults instantiates a new LogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResponseWithDefaults() *LogResponse {
	this := LogResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LogResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LogResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LogResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LogResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMessage returns the Message field value
func (o *LogResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LogResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *LogResponse) SetMessage(v string) {
	o.Message = v
}

func (o LogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableLogResponse struct {
	value *LogResponse
	isSet bool
}

func (v NullableLogResponse) Get() *LogResponse {
	return v.value
}

func (v *NullableLogResponse) Set(val *LogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResponse(val *LogResponse) *NullableLogResponse {
	return &NullableLogResponse{value: val, isSet: true}
}

func (v NullableLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
