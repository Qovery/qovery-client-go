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

// DatabaseRequest struct for DatabaseRequest
type DatabaseRequest struct {
	// name is case insensitive
	Name string `json:"name"`
	// give a description to this database
	Description   *string                    `json:"description,omitempty"`
	Type          DatabaseTypeEnum           `json:"type"`
	Version       string                     `json:"version"`
	Mode          DatabaseModeEnum           `json:"mode"`
	Accessibility *DatabaseAccessibilityEnum `json:"accessibility,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB   Default value is linked to the database type: - MANAGED: `100` - CONTAINER   - POSTGRES: `100`   - REDIS: `100`   - MYSQL: `512`   - MONGODB: `256`
	Memory *int32 `json:"memory,omitempty"`
	// unit is GB
	Storage *int32 `json:"storage,omitempty"`
}

// NewDatabaseRequest instantiates a new DatabaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseRequest(name string, type_ DatabaseTypeEnum, version string, mode DatabaseModeEnum) *DatabaseRequest {
	this := DatabaseRequest{}
	this.Name = name
	this.Type = type_
	this.Version = version
	this.Mode = mode
	var accessibility DatabaseAccessibilityEnum = DATABASEACCESSIBILITYENUM_PRIVATE
	this.Accessibility = &accessibility
	var cpu int32 = 250
	this.Cpu = &cpu
	var storage int32 = 10
	this.Storage = &storage
	return &this
}

// NewDatabaseRequestWithDefaults instantiates a new DatabaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseRequestWithDefaults() *DatabaseRequest {
	this := DatabaseRequest{}
	var accessibility DatabaseAccessibilityEnum = DATABASEACCESSIBILITYENUM_PRIVATE
	this.Accessibility = &accessibility
	var cpu int32 = 250
	this.Cpu = &cpu
	var storage int32 = 10
	this.Storage = &storage
	return &this
}

// GetName returns the Name field value
func (o *DatabaseRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatabaseRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatabaseRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatabaseRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *DatabaseRequest) GetType() DatabaseTypeEnum {
	if o == nil {
		var ret DatabaseTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetTypeOk() (*DatabaseTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatabaseRequest) SetType(v DatabaseTypeEnum) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *DatabaseRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DatabaseRequest) SetVersion(v string) {
	o.Version = v
}

// GetMode returns the Mode field value
func (o *DatabaseRequest) GetMode() DatabaseModeEnum {
	if o == nil {
		var ret DatabaseModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetModeOk() (*DatabaseModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *DatabaseRequest) SetMode(v DatabaseModeEnum) {
	o.Mode = v
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *DatabaseRequest) GetAccessibility() DatabaseAccessibilityEnum {
	if o == nil || o.Accessibility == nil {
		var ret DatabaseAccessibilityEnum
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetAccessibilityOk() (*DatabaseAccessibilityEnum, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *DatabaseRequest) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given DatabaseAccessibilityEnum and assigns it to the Accessibility field.
func (o *DatabaseRequest) SetAccessibility(v DatabaseAccessibilityEnum) {
	o.Accessibility = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DatabaseRequest) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DatabaseRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *DatabaseRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DatabaseRequest) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DatabaseRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *DatabaseRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *DatabaseRequest) GetStorage() int32 {
	if o == nil || o.Storage == nil {
		var ret int32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetStorageOk() (*int32, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *DatabaseRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given int32 and assigns it to the Storage field.
func (o *DatabaseRequest) SetStorage(v int32) {
	o.Storage = &v
}

func (o DatabaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseRequest struct {
	value *DatabaseRequest
	isSet bool
}

func (v NullableDatabaseRequest) Get() *DatabaseRequest {
	return v.value
}

func (v *NullableDatabaseRequest) Set(val *DatabaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseRequest(val *DatabaseRequest) *NullableDatabaseRequest {
	return &NullableDatabaseRequest{value: val, isSet: true}
}

func (v NullableDatabaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
