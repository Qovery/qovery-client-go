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

// DatabaseCurrentMetricStorage struct for DatabaseCurrentMetricStorage
type DatabaseCurrentMetricStorage struct {
	RequestedInGb             *int32                     `json:"requested_in_gb,omitempty"`
	ConsumedInGb              *int32                     `json:"consumed_in_gb,omitempty"`
	ConsumedInPercent         *float32                   `json:"consumed_in_percent,omitempty"`
	WarningThresholdInPercent *float32                   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32                   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *ThresholdMetricStatusEnum `json:"status,omitempty"`
}

// NewDatabaseCurrentMetricStorage instantiates a new DatabaseCurrentMetricStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseCurrentMetricStorage() *DatabaseCurrentMetricStorage {
	this := DatabaseCurrentMetricStorage{}
	return &this
}

// NewDatabaseCurrentMetricStorageWithDefaults instantiates a new DatabaseCurrentMetricStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseCurrentMetricStorageWithDefaults() *DatabaseCurrentMetricStorage {
	this := DatabaseCurrentMetricStorage{}
	return &this
}

// GetRequestedInGb returns the RequestedInGb field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricStorage) GetRequestedInGb() int32 {
	if o == nil || o.RequestedInGb == nil {
		var ret int32
		return ret
	}
	return *o.RequestedInGb
}

// GetRequestedInGbOk returns a tuple with the RequestedInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricStorage) GetRequestedInGbOk() (*int32, bool) {
	if o == nil || o.RequestedInGb == nil {
		return nil, false
	}
	return o.RequestedInGb, true
}

// HasRequestedInGb returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricStorage) HasRequestedInGb() bool {
	if o != nil && o.RequestedInGb != nil {
		return true
	}

	return false
}

// SetRequestedInGb gets a reference to the given int32 and assigns it to the RequestedInGb field.
func (o *DatabaseCurrentMetricStorage) SetRequestedInGb(v int32) {
	o.RequestedInGb = &v
}

// GetConsumedInGb returns the ConsumedInGb field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricStorage) GetConsumedInGb() int32 {
	if o == nil || o.ConsumedInGb == nil {
		var ret int32
		return ret
	}
	return *o.ConsumedInGb
}

// GetConsumedInGbOk returns a tuple with the ConsumedInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricStorage) GetConsumedInGbOk() (*int32, bool) {
	if o == nil || o.ConsumedInGb == nil {
		return nil, false
	}
	return o.ConsumedInGb, true
}

// HasConsumedInGb returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricStorage) HasConsumedInGb() bool {
	if o != nil && o.ConsumedInGb != nil {
		return true
	}

	return false
}

// SetConsumedInGb gets a reference to the given int32 and assigns it to the ConsumedInGb field.
func (o *DatabaseCurrentMetricStorage) SetConsumedInGb(v int32) {
	o.ConsumedInGb = &v
}

// GetConsumedInPercent returns the ConsumedInPercent field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricStorage) GetConsumedInPercent() float32 {
	if o == nil || o.ConsumedInPercent == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricStorage) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil || o.ConsumedInPercent == nil {
		return nil, false
	}
	return o.ConsumedInPercent, true
}

// HasConsumedInPercent returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricStorage) HasConsumedInPercent() bool {
	if o != nil && o.ConsumedInPercent != nil {
		return true
	}

	return false
}

// SetConsumedInPercent gets a reference to the given float32 and assigns it to the ConsumedInPercent field.
func (o *DatabaseCurrentMetricStorage) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricStorage) GetWarningThresholdInPercent() float32 {
	if o == nil || o.WarningThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricStorage) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.WarningThresholdInPercent == nil {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricStorage) HasWarningThresholdInPercent() bool {
	if o != nil && o.WarningThresholdInPercent != nil {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *DatabaseCurrentMetricStorage) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricStorage) GetAlertThresholdInPercent() float32 {
	if o == nil || o.AlertThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricStorage) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.AlertThresholdInPercent == nil {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricStorage) HasAlertThresholdInPercent() bool {
	if o != nil && o.AlertThresholdInPercent != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *DatabaseCurrentMetricStorage) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricStorage) GetStatus() ThresholdMetricStatusEnum {
	if o == nil || o.Status == nil {
		var ret ThresholdMetricStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricStorage) GetStatusOk() (*ThresholdMetricStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricStorage) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ThresholdMetricStatusEnum and assigns it to the Status field.
func (o *DatabaseCurrentMetricStorage) SetStatus(v ThresholdMetricStatusEnum) {
	o.Status = &v
}

func (o DatabaseCurrentMetricStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestedInGb != nil {
		toSerialize["requested_in_gb"] = o.RequestedInGb
	}
	if o.ConsumedInGb != nil {
		toSerialize["consumed_in_gb"] = o.ConsumedInGb
	}
	if o.ConsumedInPercent != nil {
		toSerialize["consumed_in_percent"] = o.ConsumedInPercent
	}
	if o.WarningThresholdInPercent != nil {
		toSerialize["warning_threshold_in_percent"] = o.WarningThresholdInPercent
	}
	if o.AlertThresholdInPercent != nil {
		toSerialize["alert_threshold_in_percent"] = o.AlertThresholdInPercent
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseCurrentMetricStorage struct {
	value *DatabaseCurrentMetricStorage
	isSet bool
}

func (v NullableDatabaseCurrentMetricStorage) Get() *DatabaseCurrentMetricStorage {
	return v.value
}

func (v *NullableDatabaseCurrentMetricStorage) Set(val *DatabaseCurrentMetricStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseCurrentMetricStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseCurrentMetricStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseCurrentMetricStorage(val *DatabaseCurrentMetricStorage) *NullableDatabaseCurrentMetricStorage {
	return &NullableDatabaseCurrentMetricStorage{value: val, isSet: true}
}

func (v NullableDatabaseCurrentMetricStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseCurrentMetricStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
