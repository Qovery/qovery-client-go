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

// ContainerEditRequest struct for ContainerEditRequest
type ContainerEditRequest struct {
	Storage []ApplicationStorageStorageInner `json:"storage,omitempty"`
	Ports   []ServicePortPortsInner          `json:"ports,omitempty"`
	// name is case insensitive
	Name string `json:"name"`
	// id of the linked registry
	RegistryId string `json:"registry_id"`
	// The image name pattern differs according to chosen container registry provider:   * `ECR`: `repository` * `SCALEWAY_CR`: `namespace/image` * `DOCKER_HUB`: `namespace/image` * `PUBLIC_ECR`: `registry_alias/repository`
	ImageName string `json:"image_name"`
	// tag of the image container
	Tag       string   `json:"tag"`
	Arguments []string `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32 `json:"max_running_instances,omitempty"`
}

// NewContainerEditRequest instantiates a new ContainerEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerEditRequest(name string, registryId string, imageName string, tag string) *ContainerEditRequest {
	this := ContainerEditRequest{}
	this.Name = name
	this.RegistryId = registryId
	this.ImageName = imageName
	this.Tag = tag
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	return &this
}

// NewContainerEditRequestWithDefaults instantiates a new ContainerEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerEditRequestWithDefaults() *ContainerEditRequest {
	this := ContainerEditRequest{}
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetStorage() []ApplicationStorageStorageInner {
	if o == nil || o.Storage == nil {
		var ret []ApplicationStorageStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetStorageOk() ([]ApplicationStorageStorageInner, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ApplicationStorageStorageInner and assigns it to the Storage field.
func (o *ContainerEditRequest) SetStorage(v []ApplicationStorageStorageInner) {
	o.Storage = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetPorts() []ServicePortPortsInner {
	if o == nil || o.Ports == nil {
		var ret []ServicePortPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetPortsOk() ([]ServicePortPortsInner, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePortPortsInner and assigns it to the Ports field.
func (o *ContainerEditRequest) SetPorts(v []ServicePortPortsInner) {
	o.Ports = v
}

// GetName returns the Name field value
func (o *ContainerEditRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerEditRequest) SetName(v string) {
	o.Name = v
}

// GetRegistryId returns the RegistryId field value
func (o *ContainerEditRequest) GetRegistryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetRegistryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryId, true
}

// SetRegistryId sets field value
func (o *ContainerEditRequest) SetRegistryId(v string) {
	o.RegistryId = v
}

// GetImageName returns the ImageName field value
func (o *ContainerEditRequest) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *ContainerEditRequest) SetImageName(v string) {
	o.ImageName = v
}

// GetTag returns the Tag field value
func (o *ContainerEditRequest) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ContainerEditRequest) SetTag(v string) {
	o.Tag = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetArguments() []string {
	if o == nil || o.Arguments == nil {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetArgumentsOk() ([]string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ContainerEditRequest) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetEntrypoint() string {
	if o == nil || o.Entrypoint == nil {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetEntrypointOk() (*string, bool) {
	if o == nil || o.Entrypoint == nil {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasEntrypoint() bool {
	if o != nil && o.Entrypoint != nil {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *ContainerEditRequest) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ContainerEditRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ContainerEditRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ContainerEditRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ContainerEditRequest) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEditRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ContainerEditRequest) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ContainerEditRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

func (o ContainerEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["registry_id"] = o.RegistryId
	}
	if true {
		toSerialize["image_name"] = o.ImageName
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Entrypoint != nil {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningInstances != nil {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if o.MaxRunningInstances != nil {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	return json.Marshal(toSerialize)
}

type NullableContainerEditRequest struct {
	value *ContainerEditRequest
	isSet bool
}

func (v NullableContainerEditRequest) Get() *ContainerEditRequest {
	return v.value
}

func (v *NullableContainerEditRequest) Set(val *ContainerEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerEditRequest(val *ContainerEditRequest) *NullableContainerEditRequest {
	return &NullableContainerEditRequest{value: val, isSet: true}
}

func (v NullableContainerEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
