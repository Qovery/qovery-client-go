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
	"time"
)

// ClusterLogs struct for ClusterLogs
type ClusterLogs struct {
	// log level
	Type *string `json:"type,omitempty"`
	// log date creation
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// log step
	Step    *string             `json:"step,omitempty"`
	Message *ClusterLogsMessage `json:"message,omitempty"`
	Error   *ClusterLogsError   `json:"error,omitempty"`
	Details *ClusterLogsDetails `json:"details,omitempty"`
}

// NewClusterLogs instantiates a new ClusterLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogs() *ClusterLogs {
	this := ClusterLogs{}
	return &this
}

// NewClusterLogsWithDefaults instantiates a new ClusterLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsWithDefaults() *ClusterLogs {
	this := ClusterLogs{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterLogs) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterLogs) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterLogs) SetType(v string) {
	o.Type = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClusterLogs) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClusterLogs) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ClusterLogs) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *ClusterLogs) GetStep() string {
	if o == nil || o.Step == nil {
		var ret string
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetStepOk() (*string, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *ClusterLogs) HasStep() bool {
	if o != nil && o.Step != nil {
		return true
	}

	return false
}

// SetStep gets a reference to the given string and assigns it to the Step field.
func (o *ClusterLogs) SetStep(v string) {
	o.Step = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterLogs) GetMessage() ClusterLogsMessage {
	if o == nil || o.Message == nil {
		var ret ClusterLogsMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetMessageOk() (*ClusterLogsMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterLogs) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ClusterLogsMessage and assigns it to the Message field.
func (o *ClusterLogs) SetMessage(v ClusterLogsMessage) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ClusterLogs) GetError() ClusterLogsError {
	if o == nil || o.Error == nil {
		var ret ClusterLogsError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetErrorOk() (*ClusterLogsError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ClusterLogs) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ClusterLogsError and assigns it to the Error field.
func (o *ClusterLogs) SetError(v ClusterLogsError) {
	o.Error = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ClusterLogs) GetDetails() ClusterLogsDetails {
	if o == nil || o.Details == nil {
		var ret ClusterLogsDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetDetailsOk() (*ClusterLogsDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ClusterLogs) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ClusterLogsDetails and assigns it to the Details field.
func (o *ClusterLogs) SetDetails(v ClusterLogsDetails) {
	o.Details = &v
}

func (o ClusterLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableClusterLogs struct {
	value *ClusterLogs
	isSet bool
}

func (v NullableClusterLogs) Get() *ClusterLogs {
	return v.value
}

func (v *NullableClusterLogs) Set(val *ClusterLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogs(val *ClusterLogs) *NullableClusterLogs {
	return &NullableClusterLogs{value: val, isSet: true}
}

func (v NullableClusterLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}