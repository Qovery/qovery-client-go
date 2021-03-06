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

// ClusterAllOf struct for ClusterAllOf
type ClusterAllOf struct {
	// name is case-insensitive
	Name            string            `json:"name"`
	Description     *string           `json:"description,omitempty"`
	Region          string            `json:"region"`
	CloudProvider   CloudProviderEnum `json:"cloud_provider"`
	MinRunningNodes *int32            `json:"min_running_nodes,omitempty"`
	MaxRunningNodes *int32            `json:"max_running_nodes,omitempty"`
	// Unit is in GB. The disk size to be used for the node configuration
	DiskSize *int32 `json:"disk_size,omitempty"`
	// the instance type to be used for this cluster. The list of values can be retrieved via the endpoint /{CloudProvider}/instanceType
	InstanceType *string         `json:"instance_type,omitempty"`
	Kubernetes   *KubernetesEnum `json:"kubernetes,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration
	EstimatedCloudProviderCost *int32     `json:"estimated_cloud_provider_cost,omitempty"`
	Status                     *StateEnum `json:"status,omitempty"`
	HasAccess                  *bool      `json:"has_access,omitempty"`
	Version                    *string    `json:"version,omitempty"`
	IsDefault                  *bool      `json:"is_default,omitempty"`
	// specific flag to indicate that this cluster is a production one
	Production *bool `json:"production,omitempty"`
	// Indicate your public ssh_key to remotely connect to your EC2 instance.
	SshKeys  []string         `json:"ssh_keys,omitempty"`
	Features []ClusterFeature `json:"features,omitempty"`
}

// NewClusterAllOf instantiates a new ClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAllOf(name string, region string, cloudProvider CloudProviderEnum) *ClusterAllOf {
	this := ClusterAllOf{}
	this.Name = name
	this.Region = region
	this.CloudProvider = cloudProvider
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var diskSize int32 = 20
	this.DiskSize = &diskSize
	return &this
}

// NewClusterAllOfWithDefaults instantiates a new ClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAllOfWithDefaults() *ClusterAllOf {
	this := ClusterAllOf{}
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var diskSize int32 = 20
	this.DiskSize = &diskSize
	return &this
}

// GetName returns the Name field value
func (o *ClusterAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetRegion returns the Region field value
func (o *ClusterAllOf) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ClusterAllOf) SetRegion(v string) {
	o.Region = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ClusterAllOf) GetCloudProvider() CloudProviderEnum {
	if o == nil {
		var ret CloudProviderEnum
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetCloudProviderOk() (*CloudProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ClusterAllOf) SetCloudProvider(v CloudProviderEnum) {
	o.CloudProvider = v
}

// GetMinRunningNodes returns the MinRunningNodes field value if set, zero value otherwise.
func (o *ClusterAllOf) GetMinRunningNodes() int32 {
	if o == nil || o.MinRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningNodes
}

// GetMinRunningNodesOk returns a tuple with the MinRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetMinRunningNodesOk() (*int32, bool) {
	if o == nil || o.MinRunningNodes == nil {
		return nil, false
	}
	return o.MinRunningNodes, true
}

// HasMinRunningNodes returns a boolean if a field has been set.
func (o *ClusterAllOf) HasMinRunningNodes() bool {
	if o != nil && o.MinRunningNodes != nil {
		return true
	}

	return false
}

// SetMinRunningNodes gets a reference to the given int32 and assigns it to the MinRunningNodes field.
func (o *ClusterAllOf) SetMinRunningNodes(v int32) {
	o.MinRunningNodes = &v
}

// GetMaxRunningNodes returns the MaxRunningNodes field value if set, zero value otherwise.
func (o *ClusterAllOf) GetMaxRunningNodes() int32 {
	if o == nil || o.MaxRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningNodes
}

// GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetMaxRunningNodesOk() (*int32, bool) {
	if o == nil || o.MaxRunningNodes == nil {
		return nil, false
	}
	return o.MaxRunningNodes, true
}

// HasMaxRunningNodes returns a boolean if a field has been set.
func (o *ClusterAllOf) HasMaxRunningNodes() bool {
	if o != nil && o.MaxRunningNodes != nil {
		return true
	}

	return false
}

// SetMaxRunningNodes gets a reference to the given int32 and assigns it to the MaxRunningNodes field.
func (o *ClusterAllOf) SetMaxRunningNodes(v int32) {
	o.MaxRunningNodes = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *ClusterAllOf) GetDiskSize() int32 {
	if o == nil || o.DiskSize == nil {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetDiskSizeOk() (*int32, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *ClusterAllOf) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *ClusterAllOf) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *ClusterAllOf) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *ClusterAllOf) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *ClusterAllOf) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetKubernetes returns the Kubernetes field value if set, zero value otherwise.
func (o *ClusterAllOf) GetKubernetes() KubernetesEnum {
	if o == nil || o.Kubernetes == nil {
		var ret KubernetesEnum
		return ret
	}
	return *o.Kubernetes
}

// GetKubernetesOk returns a tuple with the Kubernetes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetKubernetesOk() (*KubernetesEnum, bool) {
	if o == nil || o.Kubernetes == nil {
		return nil, false
	}
	return o.Kubernetes, true
}

// HasKubernetes returns a boolean if a field has been set.
func (o *ClusterAllOf) HasKubernetes() bool {
	if o != nil && o.Kubernetes != nil {
		return true
	}

	return false
}

// SetKubernetes gets a reference to the given KubernetesEnum and assigns it to the Kubernetes field.
func (o *ClusterAllOf) SetKubernetes(v KubernetesEnum) {
	o.Kubernetes = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterAllOf) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ClusterAllOf) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ClusterAllOf) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ClusterAllOf) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ClusterAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ClusterAllOf) SetMemory(v int32) {
	o.Memory = &v
}

// GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field value if set, zero value otherwise.
func (o *ClusterAllOf) GetEstimatedCloudProviderCost() int32 {
	if o == nil || o.EstimatedCloudProviderCost == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedCloudProviderCost
}

// GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetEstimatedCloudProviderCostOk() (*int32, bool) {
	if o == nil || o.EstimatedCloudProviderCost == nil {
		return nil, false
	}
	return o.EstimatedCloudProviderCost, true
}

// HasEstimatedCloudProviderCost returns a boolean if a field has been set.
func (o *ClusterAllOf) HasEstimatedCloudProviderCost() bool {
	if o != nil && o.EstimatedCloudProviderCost != nil {
		return true
	}

	return false
}

// SetEstimatedCloudProviderCost gets a reference to the given int32 and assigns it to the EstimatedCloudProviderCost field.
func (o *ClusterAllOf) SetEstimatedCloudProviderCost(v int32) {
	o.EstimatedCloudProviderCost = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterAllOf) GetStatus() StateEnum {
	if o == nil || o.Status == nil {
		var ret StateEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetStatusOk() (*StateEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StateEnum and assigns it to the Status field.
func (o *ClusterAllOf) SetStatus(v StateEnum) {
	o.Status = &v
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *ClusterAllOf) GetHasAccess() bool {
	if o == nil || o.HasAccess == nil {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetHasAccessOk() (*bool, bool) {
	if o == nil || o.HasAccess == nil {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *ClusterAllOf) HasHasAccess() bool {
	if o != nil && o.HasAccess != nil {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *ClusterAllOf) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ClusterAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ClusterAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ClusterAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *ClusterAllOf) GetProduction() bool {
	if o == nil || o.Production == nil {
		var ret bool
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetProductionOk() (*bool, bool) {
	if o == nil || o.Production == nil {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *ClusterAllOf) HasProduction() bool {
	if o != nil && o.Production != nil {
		return true
	}

	return false
}

// SetProduction gets a reference to the given bool and assigns it to the Production field.
func (o *ClusterAllOf) SetProduction(v bool) {
	o.Production = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ClusterAllOf) GetSshKeys() []string {
	if o == nil || o.SshKeys == nil {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetSshKeysOk() ([]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ClusterAllOf) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *ClusterAllOf) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClusterAllOf) GetFeatures() []ClusterFeature {
	if o == nil || o.Features == nil {
		var ret []ClusterFeature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAllOf) GetFeaturesOk() ([]ClusterFeature, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClusterAllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterFeature and assigns it to the Features field.
func (o *ClusterAllOf) SetFeatures(v []ClusterFeature) {
	o.Features = v
}

func (o ClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.MinRunningNodes != nil {
		toSerialize["min_running_nodes"] = o.MinRunningNodes
	}
	if o.MaxRunningNodes != nil {
		toSerialize["max_running_nodes"] = o.MaxRunningNodes
	}
	if o.DiskSize != nil {
		toSerialize["disk_size"] = o.DiskSize
	}
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
	if o.Kubernetes != nil {
		toSerialize["kubernetes"] = o.Kubernetes
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.EstimatedCloudProviderCost != nil {
		toSerialize["estimated_cloud_provider_cost"] = o.EstimatedCloudProviderCost
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.HasAccess != nil {
		toSerialize["has_access"] = o.HasAccess
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.Production != nil {
		toSerialize["production"] = o.Production
	}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAllOf struct {
	value *ClusterAllOf
	isSet bool
}

func (v NullableClusterAllOf) Get() *ClusterAllOf {
	return v.value
}

func (v *NullableClusterAllOf) Set(val *ClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAllOf(val *ClusterAllOf) *NullableClusterAllOf {
	return &NullableClusterAllOf{value: val, isSet: true}
}

func (v NullableClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
