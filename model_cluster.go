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
)

// Cluster struct for Cluster
type Cluster struct {
	// name is case-insensitive
	Name          string  `json:"name"`
	Description   *string `json:"description,omitempty"`
	CloudProvider string  `json:"cloud_provider"`
	Region        string  `json:"region"`
	AutoUpdate    *bool   `json:"auto_update,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *float32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory          *float32 `json:"memory,omitempty"`
	MinRunningNodes *int32   `json:"min_running_nodes,omitempty"`
	MaxRunningNodes *int32   `json:"max_running_nodes,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(name string, cloudProvider string, region string) *Cluster {
	this := Cluster{}
	this.Name = name
	this.CloudProvider = cloudProvider
	this.Region = region
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	return &this
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Cluster) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cluster) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cluster) SetDescription(v string) {
	o.Description = &v
}

// GetCloudProvider returns the CloudProvider field value
func (o *Cluster) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *Cluster) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegion returns the Region field value
func (o *Cluster) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Cluster) SetRegion(v string) {
	o.Region = v
}

// GetAutoUpdate returns the AutoUpdate field value if set, zero value otherwise.
func (o *Cluster) GetAutoUpdate() bool {
	if o == nil || o.AutoUpdate == nil {
		var ret bool
		return ret
	}
	return *o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetAutoUpdateOk() (*bool, bool) {
	if o == nil || o.AutoUpdate == nil {
		return nil, false
	}
	return o.AutoUpdate, true
}

// HasAutoUpdate returns a boolean if a field has been set.
func (o *Cluster) HasAutoUpdate() bool {
	if o != nil && o.AutoUpdate != nil {
		return true
	}

	return false
}

// SetAutoUpdate gets a reference to the given bool and assigns it to the AutoUpdate field.
func (o *Cluster) SetAutoUpdate(v bool) {
	o.AutoUpdate = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Cluster) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Cluster) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *Cluster) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Cluster) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Cluster) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *Cluster) SetMemory(v float32) {
	o.Memory = &v
}

// GetMinRunningNodes returns the MinRunningNodes field value if set, zero value otherwise.
func (o *Cluster) GetMinRunningNodes() int32 {
	if o == nil || o.MinRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningNodes
}

// GetMinRunningNodesOk returns a tuple with the MinRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMinRunningNodesOk() (*int32, bool) {
	if o == nil || o.MinRunningNodes == nil {
		return nil, false
	}
	return o.MinRunningNodes, true
}

// HasMinRunningNodes returns a boolean if a field has been set.
func (o *Cluster) HasMinRunningNodes() bool {
	if o != nil && o.MinRunningNodes != nil {
		return true
	}

	return false
}

// SetMinRunningNodes gets a reference to the given int32 and assigns it to the MinRunningNodes field.
func (o *Cluster) SetMinRunningNodes(v int32) {
	o.MinRunningNodes = &v
}

// GetMaxRunningNodes returns the MaxRunningNodes field value if set, zero value otherwise.
func (o *Cluster) GetMaxRunningNodes() int32 {
	if o == nil || o.MaxRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningNodes
}

// GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMaxRunningNodesOk() (*int32, bool) {
	if o == nil || o.MaxRunningNodes == nil {
		return nil, false
	}
	return o.MaxRunningNodes, true
}

// HasMaxRunningNodes returns a boolean if a field has been set.
func (o *Cluster) HasMaxRunningNodes() bool {
	if o != nil && o.MaxRunningNodes != nil {
		return true
	}

	return false
}

// SetMaxRunningNodes gets a reference to the given int32 and assigns it to the MaxRunningNodes field.
func (o *Cluster) SetMaxRunningNodes(v int32) {
	o.MaxRunningNodes = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.AutoUpdate != nil {
		toSerialize["auto_update"] = o.AutoUpdate
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningNodes != nil {
		toSerialize["min_running_nodes"] = o.MinRunningNodes
	}
	if o.MaxRunningNodes != nil {
		toSerialize["max_running_nodes"] = o.MaxRunningNodes
	}
	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
