/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.4
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"fmt"
)

// JobForceEvent the model 'JobForceEvent'
type JobForceEvent string

// List of JobForceEvent
const (
	JOBFORCEEVENT_START  JobForceEvent = "START"
	JOBFORCEEVENT_STOP   JobForceEvent = "STOP"
	JOBFORCEEVENT_DELETE JobForceEvent = "DELETE"
	JOBFORCEEVENT_CRON   JobForceEvent = "CRON"
)

// All allowed values of JobForceEvent enum
var AllowedJobForceEventEnumValues = []JobForceEvent{
	"START",
	"STOP",
	"DELETE",
	"CRON",
}

func (v *JobForceEvent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobForceEvent(value)
	for _, existing := range AllowedJobForceEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobForceEvent", value)
}

// NewJobForceEventFromValue returns a pointer to a valid JobForceEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobForceEventFromValue(v string) (*JobForceEvent, error) {
	ev := JobForceEvent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobForceEvent: valid values are %v", v, AllowedJobForceEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobForceEvent) IsValid() bool {
	for _, existing := range AllowedJobForceEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobForceEvent value
func (v JobForceEvent) Ptr() *JobForceEvent {
	return &v
}

type NullableJobForceEvent struct {
	value *JobForceEvent
	isSet bool
}

func (v NullableJobForceEvent) Get() *JobForceEvent {
	return v.value
}

func (v *NullableJobForceEvent) Set(val *JobForceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableJobForceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableJobForceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobForceEvent(val *JobForceEvent) *NullableJobForceEvent {
	return &NullableJobForceEvent{value: val, isSet: true}
}

func (v NullableJobForceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobForceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
