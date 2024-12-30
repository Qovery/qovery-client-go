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
	"fmt"
)

// checks if the KarpenterNodePoolConsolidation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KarpenterNodePoolConsolidation{}

// KarpenterNodePoolConsolidation struct for KarpenterNodePoolConsolidation
type KarpenterNodePoolConsolidation struct {
	Enabled bool          `json:"enabled"`
	Days    []WeekdayEnum `json:"days"`
	// The start date of the consolidation. The format should follow ISO-8601 convention: \"PThh:mm\"
	StartTime string `json:"start_time"`
	// The duration during the consolidation will be active. The format should follow ISO-8601 convention: \"PThhHmmM\"
	Duration             string `json:"duration"`
	AdditionalProperties map[string]interface{}
}

type _KarpenterNodePoolConsolidation KarpenterNodePoolConsolidation

// NewKarpenterNodePoolConsolidation instantiates a new KarpenterNodePoolConsolidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKarpenterNodePoolConsolidation(enabled bool, days []WeekdayEnum, startTime string, duration string) *KarpenterNodePoolConsolidation {
	this := KarpenterNodePoolConsolidation{}
	this.Enabled = enabled
	this.Days = days
	this.StartTime = startTime
	this.Duration = duration
	return &this
}

// NewKarpenterNodePoolConsolidationWithDefaults instantiates a new KarpenterNodePoolConsolidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKarpenterNodePoolConsolidationWithDefaults() *KarpenterNodePoolConsolidation {
	this := KarpenterNodePoolConsolidation{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *KarpenterNodePoolConsolidation) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *KarpenterNodePoolConsolidation) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *KarpenterNodePoolConsolidation) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDays returns the Days field value
func (o *KarpenterNodePoolConsolidation) GetDays() []WeekdayEnum {
	if o == nil {
		var ret []WeekdayEnum
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *KarpenterNodePoolConsolidation) GetDaysOk() ([]WeekdayEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days, true
}

// SetDays sets field value
func (o *KarpenterNodePoolConsolidation) SetDays(v []WeekdayEnum) {
	o.Days = v
}

// GetStartTime returns the StartTime field value
func (o *KarpenterNodePoolConsolidation) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *KarpenterNodePoolConsolidation) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *KarpenterNodePoolConsolidation) SetStartTime(v string) {
	o.StartTime = v
}

// GetDuration returns the Duration field value
func (o *KarpenterNodePoolConsolidation) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *KarpenterNodePoolConsolidation) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *KarpenterNodePoolConsolidation) SetDuration(v string) {
	o.Duration = v
}

func (o KarpenterNodePoolConsolidation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KarpenterNodePoolConsolidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["days"] = o.Days
	toSerialize["start_time"] = o.StartTime
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KarpenterNodePoolConsolidation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"days",
		"start_time",
		"duration",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKarpenterNodePoolConsolidation := _KarpenterNodePoolConsolidation{}

	err = json.Unmarshal(data, &varKarpenterNodePoolConsolidation)

	if err != nil {
		return err
	}

	*o = KarpenterNodePoolConsolidation(varKarpenterNodePoolConsolidation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "days")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "duration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKarpenterNodePoolConsolidation struct {
	value *KarpenterNodePoolConsolidation
	isSet bool
}

func (v NullableKarpenterNodePoolConsolidation) Get() *KarpenterNodePoolConsolidation {
	return v.value
}

func (v *NullableKarpenterNodePoolConsolidation) Set(val *KarpenterNodePoolConsolidation) {
	v.value = val
	v.isSet = true
}

func (v NullableKarpenterNodePoolConsolidation) IsSet() bool {
	return v.isSet
}

func (v *NullableKarpenterNodePoolConsolidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKarpenterNodePoolConsolidation(val *KarpenterNodePoolConsolidation) *NullableKarpenterNodePoolConsolidation {
	return &NullableKarpenterNodePoolConsolidation{value: val, isSet: true}
}

func (v NullableKarpenterNodePoolConsolidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKarpenterNodePoolConsolidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
