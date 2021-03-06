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

// EnvironmentDeploymentRuleAllOf struct for EnvironmentDeploymentRuleAllOf
type EnvironmentDeploymentRuleAllOf struct {
	AutoDeploy  *bool         `json:"auto_deploy,omitempty"`
	AutoStop    *bool         `json:"auto_stop,omitempty"`
	AutoDelete  *bool         `json:"auto_delete,omitempty"`
	AutoPreview *bool         `json:"auto_preview,omitempty"`
	Timezone    string        `json:"timezone"`
	StartTime   time.Time     `json:"start_time"`
	StopTime    time.Time     `json:"stop_time"`
	Weekdays    []WeekdayEnum `json:"weekdays"`
}

// NewEnvironmentDeploymentRuleAllOf instantiates a new EnvironmentDeploymentRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDeploymentRuleAllOf(timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum) *EnvironmentDeploymentRuleAllOf {
	this := EnvironmentDeploymentRuleAllOf{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoStop bool = false
	this.AutoStop = &autoStop
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoPreview bool = false
	this.AutoPreview = &autoPreview
	this.Timezone = timezone
	this.StartTime = startTime
	this.StopTime = stopTime
	this.Weekdays = weekdays
	return &this
}

// NewEnvironmentDeploymentRuleAllOfWithDefaults instantiates a new EnvironmentDeploymentRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDeploymentRuleAllOfWithDefaults() *EnvironmentDeploymentRuleAllOf {
	this := EnvironmentDeploymentRuleAllOf{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoStop bool = false
	this.AutoStop = &autoStop
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoPreview bool = false
	this.AutoPreview = &autoPreview
	return &this
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleAllOf) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *EnvironmentDeploymentRuleAllOf) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoStop() bool {
	if o == nil || o.AutoStop == nil {
		var ret bool
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoStopOk() (*bool, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleAllOf) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given bool and assigns it to the AutoStop field.
func (o *EnvironmentDeploymentRuleAllOf) SetAutoStop(v bool) {
	o.AutoStop = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleAllOf) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *EnvironmentDeploymentRuleAllOf) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoPreview() bool {
	if o == nil || o.AutoPreview == nil {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || o.AutoPreview == nil {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleAllOf) HasAutoPreview() bool {
	if o != nil && o.AutoPreview != nil {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *EnvironmentDeploymentRuleAllOf) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetTimezone returns the Timezone field value
func (o *EnvironmentDeploymentRuleAllOf) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *EnvironmentDeploymentRuleAllOf) SetTimezone(v string) {
	o.Timezone = v
}

// GetStartTime returns the StartTime field value
func (o *EnvironmentDeploymentRuleAllOf) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *EnvironmentDeploymentRuleAllOf) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetStopTime returns the StopTime field value
func (o *EnvironmentDeploymentRuleAllOf) GetStopTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopTime, true
}

// SetStopTime sets field value
func (o *EnvironmentDeploymentRuleAllOf) SetStopTime(v time.Time) {
	o.StopTime = v
}

// GetWeekdays returns the Weekdays field value
func (o *EnvironmentDeploymentRuleAllOf) GetWeekdays() []WeekdayEnum {
	if o == nil {
		var ret []WeekdayEnum
		return ret
	}

	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleAllOf) GetWeekdaysOk() ([]WeekdayEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// SetWeekdays sets field value
func (o *EnvironmentDeploymentRuleAllOf) SetWeekdays(v []WeekdayEnum) {
	o.Weekdays = v
}

func (o EnvironmentDeploymentRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if o.AutoStop != nil {
		toSerialize["auto_stop"] = o.AutoStop
	}
	if o.AutoDelete != nil {
		toSerialize["auto_delete"] = o.AutoDelete
	}
	if o.AutoPreview != nil {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["stop_time"] = o.StopTime
	}
	if true {
		toSerialize["weekdays"] = o.Weekdays
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentDeploymentRuleAllOf struct {
	value *EnvironmentDeploymentRuleAllOf
	isSet bool
}

func (v NullableEnvironmentDeploymentRuleAllOf) Get() *EnvironmentDeploymentRuleAllOf {
	return v.value
}

func (v *NullableEnvironmentDeploymentRuleAllOf) Set(val *EnvironmentDeploymentRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDeploymentRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDeploymentRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDeploymentRuleAllOf(val *EnvironmentDeploymentRuleAllOf) *NullableEnvironmentDeploymentRuleAllOf {
	return &NullableEnvironmentDeploymentRuleAllOf{value: val, isSet: true}
}

func (v NullableEnvironmentDeploymentRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDeploymentRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
