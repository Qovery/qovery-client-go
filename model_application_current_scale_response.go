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
	"time"
)

// ApplicationCurrentScaleResponse struct for ApplicationCurrentScaleResponse
type ApplicationCurrentScaleResponse struct {
	Min                       *int32     `json:"min,omitempty"`
	Max                       *int32     `json:"max,omitempty"`
	Running                   *int32     `json:"running,omitempty"`
	RunningInPercent          *float32   `json:"running_in_percent,omitempty"`
	WarningThresholdInPercent *float32   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *string    `json:"status,omitempty"`
	UpdatedAt                 *time.Time `json:"updated_at,omitempty"`
}

// NewApplicationCurrentScaleResponse instantiates a new ApplicationCurrentScaleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCurrentScaleResponse() *ApplicationCurrentScaleResponse {
	this := ApplicationCurrentScaleResponse{}
	return &this
}

// NewApplicationCurrentScaleResponseWithDefaults instantiates a new ApplicationCurrentScaleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCurrentScaleResponseWithDefaults() *ApplicationCurrentScaleResponse {
	this := ApplicationCurrentScaleResponse{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *ApplicationCurrentScaleResponse) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *ApplicationCurrentScaleResponse) SetMax(v int32) {
	o.Max = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetRunning() int32 {
	if o == nil || o.Running == nil {
		var ret int32
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetRunningOk() (*int32, bool) {
	if o == nil || o.Running == nil {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasRunning() bool {
	if o != nil && o.Running != nil {
		return true
	}

	return false
}

// SetRunning gets a reference to the given int32 and assigns it to the Running field.
func (o *ApplicationCurrentScaleResponse) SetRunning(v int32) {
	o.Running = &v
}

// GetRunningInPercent returns the RunningInPercent field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetRunningInPercent() float32 {
	if o == nil || o.RunningInPercent == nil {
		var ret float32
		return ret
	}
	return *o.RunningInPercent
}

// GetRunningInPercentOk returns a tuple with the RunningInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetRunningInPercentOk() (*float32, bool) {
	if o == nil || o.RunningInPercent == nil {
		return nil, false
	}
	return o.RunningInPercent, true
}

// HasRunningInPercent returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasRunningInPercent() bool {
	if o != nil && o.RunningInPercent != nil {
		return true
	}

	return false
}

// SetRunningInPercent gets a reference to the given float32 and assigns it to the RunningInPercent field.
func (o *ApplicationCurrentScaleResponse) SetRunningInPercent(v float32) {
	o.RunningInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetWarningThresholdInPercent() float32 {
	if o == nil || o.WarningThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.WarningThresholdInPercent == nil {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasWarningThresholdInPercent() bool {
	if o != nil && o.WarningThresholdInPercent != nil {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *ApplicationCurrentScaleResponse) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetAlertThresholdInPercent() float32 {
	if o == nil || o.AlertThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.AlertThresholdInPercent == nil {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasAlertThresholdInPercent() bool {
	if o != nil && o.AlertThresholdInPercent != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *ApplicationCurrentScaleResponse) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplicationCurrentScaleResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationCurrentScaleResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCurrentScaleResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationCurrentScaleResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationCurrentScaleResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ApplicationCurrentScaleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Running != nil {
		toSerialize["running"] = o.Running
	}
	if o.RunningInPercent != nil {
		toSerialize["running_in_percent"] = o.RunningInPercent
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
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationCurrentScaleResponse struct {
	value *ApplicationCurrentScaleResponse
	isSet bool
}

func (v NullableApplicationCurrentScaleResponse) Get() *ApplicationCurrentScaleResponse {
	return v.value
}

func (v *NullableApplicationCurrentScaleResponse) Set(val *ApplicationCurrentScaleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCurrentScaleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCurrentScaleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCurrentScaleResponse(val *ApplicationCurrentScaleResponse) *NullableApplicationCurrentScaleResponse {
	return &NullableApplicationCurrentScaleResponse{value: val, isSet: true}
}

func (v NullableApplicationCurrentScaleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCurrentScaleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
