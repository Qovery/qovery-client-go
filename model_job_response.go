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

// JobResponse - struct for JobResponse
type JobResponse struct {
	CronJobResponse      *CronJobResponse
	LifecycleJobResponse *LifecycleJobResponse
}

// CronJobResponseAsJobResponse is a convenience function that returns CronJobResponse wrapped in JobResponse
func CronJobResponseAsJobResponse(v *CronJobResponse) JobResponse {
	return JobResponse{
		CronJobResponse: v,
	}
}

// LifecycleJobResponseAsJobResponse is a convenience function that returns LifecycleJobResponse wrapped in JobResponse
func LifecycleJobResponseAsJobResponse(v *LifecycleJobResponse) JobResponse {
	return JobResponse{
		LifecycleJobResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *JobResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CronJobResponse
	err = newStrictDecoder(data).Decode(&dst.CronJobResponse)
	if err == nil {
		jsonCronJobResponse, _ := json.Marshal(dst.CronJobResponse)
		if string(jsonCronJobResponse) == "{}" { // empty struct
			dst.CronJobResponse = nil
		} else {
			match++
		}
	} else {
		dst.CronJobResponse = nil
	}

	// try to unmarshal data into LifecycleJobResponse
	err = newStrictDecoder(data).Decode(&dst.LifecycleJobResponse)
	if err == nil {
		jsonLifecycleJobResponse, _ := json.Marshal(dst.LifecycleJobResponse)
		if string(jsonLifecycleJobResponse) == "{}" { // empty struct
			dst.LifecycleJobResponse = nil
		} else {
			match++
		}
	} else {
		dst.LifecycleJobResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CronJobResponse = nil
		dst.LifecycleJobResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(JobResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(JobResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src JobResponse) MarshalJSON() ([]byte, error) {
	if src.CronJobResponse != nil {
		return json.Marshal(&src.CronJobResponse)
	}

	if src.LifecycleJobResponse != nil {
		return json.Marshal(&src.LifecycleJobResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *JobResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CronJobResponse != nil {
		return obj.CronJobResponse
	}

	if obj.LifecycleJobResponse != nil {
		return obj.LifecycleJobResponse
	}

	// all schemas are nil
	return nil
}

type NullableJobResponse struct {
	value *JobResponse
	isSet bool
}

func (v NullableJobResponse) Get() *JobResponse {
	return v.value
}

func (v *NullableJobResponse) Set(val *JobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResponse(val *JobResponse) *NullableJobResponse {
	return &NullableJobResponse{value: val, isSet: true}
}

func (v NullableJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
