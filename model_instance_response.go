/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// InstanceResponse struct for InstanceResponse
type InstanceResponse struct {
	CreatedAt *time.Time                                       `json:"created_at,omitempty"`
	Name      *string                                          `json:"name,omitempty"`
	Cpu       *EnvironmentDatabasesCurrentMetricResponseCpu    `json:"cpu,omitempty"`
	Memory    *EnvironmentDatabasesCurrentMetricResponseMemory `json:"memory,omitempty"`
}

// NewInstanceResponse instantiates a new InstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceResponse() *InstanceResponse {
	this := InstanceResponse{}
	return &this
}

// NewInstanceResponseWithDefaults instantiates a new InstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceResponseWithDefaults() *InstanceResponse {
	this := InstanceResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InstanceResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InstanceResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InstanceResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceResponse) SetName(v string) {
	o.Name = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *InstanceResponse) GetCpu() EnvironmentDatabasesCurrentMetricResponseCpu {
	if o == nil || o.Cpu == nil {
		var ret EnvironmentDatabasesCurrentMetricResponseCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponse) GetCpuOk() (*EnvironmentDatabasesCurrentMetricResponseCpu, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *InstanceResponse) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given EnvironmentDatabasesCurrentMetricResponseCpu and assigns it to the Cpu field.
func (o *InstanceResponse) SetCpu(v EnvironmentDatabasesCurrentMetricResponseCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *InstanceResponse) GetMemory() EnvironmentDatabasesCurrentMetricResponseMemory {
	if o == nil || o.Memory == nil {
		var ret EnvironmentDatabasesCurrentMetricResponseMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponse) GetMemoryOk() (*EnvironmentDatabasesCurrentMetricResponseMemory, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *InstanceResponse) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given EnvironmentDatabasesCurrentMetricResponseMemory and assigns it to the Memory field.
func (o *InstanceResponse) SetMemory(v EnvironmentDatabasesCurrentMetricResponseMemory) {
	o.Memory = &v
}

func (o InstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceResponse struct {
	value *InstanceResponse
	isSet bool
}

func (v NullableInstanceResponse) Get() *InstanceResponse {
	return v.value
}

func (v *NullableInstanceResponse) Set(val *InstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceResponse(val *InstanceResponse) *NullableInstanceResponse {
	return &NullableInstanceResponse{value: val, isSet: true}
}

func (v NullableInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
