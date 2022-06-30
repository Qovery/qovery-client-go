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

// ClusterRequest struct for ClusterRequest
type ClusterRequest struct {
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
	// specific flag to indicate that this cluster is a production one
	Production *bool `json:"production,omitempty"`
	// Indicate your public ssh_key to remotely connect to your EC2 instance.
	SshKeys  []string                      `json:"ssh_keys,omitempty"`
	Features []ClusterRequestFeaturesInner `json:"features,omitempty"`
}

// NewClusterRequest instantiates a new ClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRequest(name string, region string, cloudProvider CloudProviderEnum) *ClusterRequest {
	this := ClusterRequest{}
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

// NewClusterRequestWithDefaults instantiates a new ClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRequestWithDefaults() *ClusterRequest {
	this := ClusterRequest{}
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var diskSize int32 = 20
	this.DiskSize = &diskSize
	return &this
}

// GetName returns the Name field value
func (o *ClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRegion returns the Region field value
func (o *ClusterRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ClusterRequest) SetRegion(v string) {
	o.Region = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ClusterRequest) GetCloudProvider() CloudProviderEnum {
	if o == nil {
		var ret CloudProviderEnum
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetCloudProviderOk() (*CloudProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ClusterRequest) SetCloudProvider(v CloudProviderEnum) {
	o.CloudProvider = v
}

// GetMinRunningNodes returns the MinRunningNodes field value if set, zero value otherwise.
func (o *ClusterRequest) GetMinRunningNodes() int32 {
	if o == nil || o.MinRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningNodes
}

// GetMinRunningNodesOk returns a tuple with the MinRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetMinRunningNodesOk() (*int32, bool) {
	if o == nil || o.MinRunningNodes == nil {
		return nil, false
	}
	return o.MinRunningNodes, true
}

// HasMinRunningNodes returns a boolean if a field has been set.
func (o *ClusterRequest) HasMinRunningNodes() bool {
	if o != nil && o.MinRunningNodes != nil {
		return true
	}

	return false
}

// SetMinRunningNodes gets a reference to the given int32 and assigns it to the MinRunningNodes field.
func (o *ClusterRequest) SetMinRunningNodes(v int32) {
	o.MinRunningNodes = &v
}

// GetMaxRunningNodes returns the MaxRunningNodes field value if set, zero value otherwise.
func (o *ClusterRequest) GetMaxRunningNodes() int32 {
	if o == nil || o.MaxRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningNodes
}

// GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetMaxRunningNodesOk() (*int32, bool) {
	if o == nil || o.MaxRunningNodes == nil {
		return nil, false
	}
	return o.MaxRunningNodes, true
}

// HasMaxRunningNodes returns a boolean if a field has been set.
func (o *ClusterRequest) HasMaxRunningNodes() bool {
	if o != nil && o.MaxRunningNodes != nil {
		return true
	}

	return false
}

// SetMaxRunningNodes gets a reference to the given int32 and assigns it to the MaxRunningNodes field.
func (o *ClusterRequest) SetMaxRunningNodes(v int32) {
	o.MaxRunningNodes = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *ClusterRequest) GetDiskSize() int32 {
	if o == nil || o.DiskSize == nil {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetDiskSizeOk() (*int32, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *ClusterRequest) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *ClusterRequest) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *ClusterRequest) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *ClusterRequest) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *ClusterRequest) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetKubernetes returns the Kubernetes field value if set, zero value otherwise.
func (o *ClusterRequest) GetKubernetes() KubernetesEnum {
	if o == nil || o.Kubernetes == nil {
		var ret KubernetesEnum
		return ret
	}
	return *o.Kubernetes
}

// GetKubernetesOk returns a tuple with the Kubernetes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetKubernetesOk() (*KubernetesEnum, bool) {
	if o == nil || o.Kubernetes == nil {
		return nil, false
	}
	return o.Kubernetes, true
}

// HasKubernetes returns a boolean if a field has been set.
func (o *ClusterRequest) HasKubernetes() bool {
	if o != nil && o.Kubernetes != nil {
		return true
	}

	return false
}

// SetKubernetes gets a reference to the given KubernetesEnum and assigns it to the Kubernetes field.
func (o *ClusterRequest) SetKubernetes(v KubernetesEnum) {
	o.Kubernetes = &v
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *ClusterRequest) GetProduction() bool {
	if o == nil || o.Production == nil {
		var ret bool
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetProductionOk() (*bool, bool) {
	if o == nil || o.Production == nil {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *ClusterRequest) HasProduction() bool {
	if o != nil && o.Production != nil {
		return true
	}

	return false
}

// SetProduction gets a reference to the given bool and assigns it to the Production field.
func (o *ClusterRequest) SetProduction(v bool) {
	o.Production = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ClusterRequest) GetSshKeys() []string {
	if o == nil || o.SshKeys == nil {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetSshKeysOk() ([]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ClusterRequest) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *ClusterRequest) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClusterRequest) GetFeatures() []ClusterRequestFeaturesInner {
	if o == nil || o.Features == nil {
		var ret []ClusterRequestFeaturesInner
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetFeaturesOk() ([]ClusterRequestFeaturesInner, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClusterRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterRequestFeaturesInner and assigns it to the Features field.
func (o *ClusterRequest) SetFeatures(v []ClusterRequestFeaturesInner) {
	o.Features = v
}

func (o ClusterRequest) MarshalJSON() ([]byte, error) {
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

type NullableClusterRequest struct {
	value *ClusterRequest
	isSet bool
}

func (v NullableClusterRequest) Get() *ClusterRequest {
	return v.value
}

func (v *NullableClusterRequest) Set(val *ClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRequest(val *ClusterRequest) *NullableClusterRequest {
	return &NullableClusterRequest{value: val, isSet: true}
}

func (v NullableClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
