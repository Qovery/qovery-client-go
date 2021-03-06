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

// EnvironmentDatabasesCurrentMetricCpu struct for EnvironmentDatabasesCurrentMetricCpu
type EnvironmentDatabasesCurrentMetricCpu struct {
	// unit is millicores (m). 1000m = 1 cpu
	RequestedInFloat *float32 `json:"requested_in_float,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	ConsumedInNumber          *float32                   `json:"consumed_in_number,omitempty"`
	ConsumedInPercent         *float32                   `json:"consumed_in_percent,omitempty"`
	WarningThresholdInPercent *float32                   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32                   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *ThresholdMetricStatusEnum `json:"status,omitempty"`
}

// NewEnvironmentDatabasesCurrentMetricCpu instantiates a new EnvironmentDatabasesCurrentMetricCpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDatabasesCurrentMetricCpu() *EnvironmentDatabasesCurrentMetricCpu {
	this := EnvironmentDatabasesCurrentMetricCpu{}
	return &this
}

// NewEnvironmentDatabasesCurrentMetricCpuWithDefaults instantiates a new EnvironmentDatabasesCurrentMetricCpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDatabasesCurrentMetricCpuWithDefaults() *EnvironmentDatabasesCurrentMetricCpu {
	this := EnvironmentDatabasesCurrentMetricCpu{}
	return &this
}

// GetRequestedInFloat returns the RequestedInFloat field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetRequestedInFloat() float32 {
	if o == nil || o.RequestedInFloat == nil {
		var ret float32
		return ret
	}
	return *o.RequestedInFloat
}

// GetRequestedInFloatOk returns a tuple with the RequestedInFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetRequestedInFloatOk() (*float32, bool) {
	if o == nil || o.RequestedInFloat == nil {
		return nil, false
	}
	return o.RequestedInFloat, true
}

// HasRequestedInFloat returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) HasRequestedInFloat() bool {
	if o != nil && o.RequestedInFloat != nil {
		return true
	}

	return false
}

// SetRequestedInFloat gets a reference to the given float32 and assigns it to the RequestedInFloat field.
func (o *EnvironmentDatabasesCurrentMetricCpu) SetRequestedInFloat(v float32) {
	o.RequestedInFloat = &v
}

// GetConsumedInNumber returns the ConsumedInNumber field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetConsumedInNumber() float32 {
	if o == nil || o.ConsumedInNumber == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInNumber
}

// GetConsumedInNumberOk returns a tuple with the ConsumedInNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetConsumedInNumberOk() (*float32, bool) {
	if o == nil || o.ConsumedInNumber == nil {
		return nil, false
	}
	return o.ConsumedInNumber, true
}

// HasConsumedInNumber returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) HasConsumedInNumber() bool {
	if o != nil && o.ConsumedInNumber != nil {
		return true
	}

	return false
}

// SetConsumedInNumber gets a reference to the given float32 and assigns it to the ConsumedInNumber field.
func (o *EnvironmentDatabasesCurrentMetricCpu) SetConsumedInNumber(v float32) {
	o.ConsumedInNumber = &v
}

// GetConsumedInPercent returns the ConsumedInPercent field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetConsumedInPercent() float32 {
	if o == nil || o.ConsumedInPercent == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil || o.ConsumedInPercent == nil {
		return nil, false
	}
	return o.ConsumedInPercent, true
}

// HasConsumedInPercent returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) HasConsumedInPercent() bool {
	if o != nil && o.ConsumedInPercent != nil {
		return true
	}

	return false
}

// SetConsumedInPercent gets a reference to the given float32 and assigns it to the ConsumedInPercent field.
func (o *EnvironmentDatabasesCurrentMetricCpu) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetWarningThresholdInPercent() float32 {
	if o == nil || o.WarningThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.WarningThresholdInPercent == nil {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) HasWarningThresholdInPercent() bool {
	if o != nil && o.WarningThresholdInPercent != nil {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *EnvironmentDatabasesCurrentMetricCpu) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetAlertThresholdInPercent() float32 {
	if o == nil || o.AlertThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.AlertThresholdInPercent == nil {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) HasAlertThresholdInPercent() bool {
	if o != nil && o.AlertThresholdInPercent != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *EnvironmentDatabasesCurrentMetricCpu) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetStatus() ThresholdMetricStatusEnum {
	if o == nil || o.Status == nil {
		var ret ThresholdMetricStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) GetStatusOk() (*ThresholdMetricStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvironmentDatabasesCurrentMetricCpu) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ThresholdMetricStatusEnum and assigns it to the Status field.
func (o *EnvironmentDatabasesCurrentMetricCpu) SetStatus(v ThresholdMetricStatusEnum) {
	o.Status = &v
}

func (o EnvironmentDatabasesCurrentMetricCpu) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestedInFloat != nil {
		toSerialize["requested_in_float"] = o.RequestedInFloat
	}
	if o.ConsumedInNumber != nil {
		toSerialize["consumed_in_number"] = o.ConsumedInNumber
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

type NullableEnvironmentDatabasesCurrentMetricCpu struct {
	value *EnvironmentDatabasesCurrentMetricCpu
	isSet bool
}

func (v NullableEnvironmentDatabasesCurrentMetricCpu) Get() *EnvironmentDatabasesCurrentMetricCpu {
	return v.value
}

func (v *NullableEnvironmentDatabasesCurrentMetricCpu) Set(val *EnvironmentDatabasesCurrentMetricCpu) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDatabasesCurrentMetricCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDatabasesCurrentMetricCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDatabasesCurrentMetricCpu(val *EnvironmentDatabasesCurrentMetricCpu) *NullableEnvironmentDatabasesCurrentMetricCpu {
	return &NullableEnvironmentDatabasesCurrentMetricCpu{value: val, isSet: true}
}

func (v NullableEnvironmentDatabasesCurrentMetricCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDatabasesCurrentMetricCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
