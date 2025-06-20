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
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'CRON'
	if jsonDict["job_type"] == "CRON" {
		// try to unmarshal JSON data into CronJobResponse
		err = json.Unmarshal(data, &dst.CronJobResponse)
		if err == nil {
			return nil // data stored in dst.CronJobResponse, return on the first match
		} else {
			dst.CronJobResponse = nil
			return fmt.Errorf("failed to unmarshal JobResponse as CronJobResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LIFECYCLE'
	if jsonDict["job_type"] == "LIFECYCLE" {
		// try to unmarshal JSON data into LifecycleJobResponse
		err = json.Unmarshal(data, &dst.LifecycleJobResponse)
		if err == nil {
			return nil // data stored in dst.LifecycleJobResponse, return on the first match
		} else {
			dst.LifecycleJobResponse = nil
			return fmt.Errorf("failed to unmarshal JobResponse as LifecycleJobResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CronJobResponse'
	if jsonDict["job_type"] == "CronJobResponse" {
		// try to unmarshal JSON data into CronJobResponse
		err = json.Unmarshal(data, &dst.CronJobResponse)
		if err == nil {
			return nil // data stored in dst.CronJobResponse, return on the first match
		} else {
			dst.CronJobResponse = nil
			return fmt.Errorf("failed to unmarshal JobResponse as CronJobResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LifecycleJobResponse'
	if jsonDict["job_type"] == "LifecycleJobResponse" {
		// try to unmarshal JSON data into LifecycleJobResponse
		err = json.Unmarshal(data, &dst.LifecycleJobResponse)
		if err == nil {
			return nil // data stored in dst.LifecycleJobResponse, return on the first match
		} else {
			dst.LifecycleJobResponse = nil
			return fmt.Errorf("failed to unmarshal JobResponse as LifecycleJobResponse: %s", err.Error())
		}
	}

	return nil
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
