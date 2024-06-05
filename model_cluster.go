/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cluster{}

// Cluster struct for Cluster
type Cluster struct {
	Id           string          `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty"`
	Organization ReferenceObject `json:"organization"`
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
	EstimatedCloudProviderCost *int32            `json:"estimated_cloud_provider_cost,omitempty"`
	Status                     *ClusterStateEnum `json:"status,omitempty"`
	HasAccess                  *bool             `json:"has_access,omitempty"`
	Version                    *string           `json:"version,omitempty"`
	IsDefault                  *bool             `json:"is_default,omitempty"`
	// specific flag to indicate that this cluster is a demo one
	IsDemo *bool `json:"is_demo,omitempty"`
	// specific flag to indicate that this cluster is a production one
	Production *bool `json:"production,omitempty"`
	// Indicate your public ssh_key to remotely connect to your EC2 instance.
	SshKeys          []string                     `json:"ssh_keys,omitempty"`
	Features         []ClusterFeatureResponse     `json:"features,omitempty"`
	DeploymentStatus *ClusterDeploymentStatusEnum `json:"deployment_status,omitempty"`
}

type _Cluster Cluster

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(id string, createdAt time.Time, organization ReferenceObject, name string, region string, cloudProvider CloudProviderEnum) *Cluster {
	this := Cluster{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Organization = organization
	this.Name = name
	this.Region = region
	this.CloudProvider = cloudProvider
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var diskSize int32 = 20
	this.DiskSize = &diskSize
	var kubernetes KubernetesEnum = KUBERNETESENUM_MANAGED
	this.Kubernetes = &kubernetes
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	var diskSize int32 = 20
	this.DiskSize = &diskSize
	var kubernetes KubernetesEnum = KUBERNETESENUM_MANAGED
	this.Kubernetes = &kubernetes
	return &this
}

// GetId returns the Id field value
func (o *Cluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cluster) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Cluster) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Cluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Cluster) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Cluster) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Cluster) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOrganization returns the Organization field value
func (o *Cluster) GetOrganization() ReferenceObject {
	if o == nil {
		var ret ReferenceObject
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetOrganizationOk() (*ReferenceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *Cluster) SetOrganization(v ReferenceObject) {
	o.Organization = v
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cluster) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cluster) SetDescription(v string) {
	o.Description = &v
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

// GetCloudProvider returns the CloudProvider field value
func (o *Cluster) GetCloudProvider() CloudProviderEnum {
	if o == nil {
		var ret CloudProviderEnum
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCloudProviderOk() (*CloudProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *Cluster) SetCloudProvider(v CloudProviderEnum) {
	o.CloudProvider = v
}

// GetMinRunningNodes returns the MinRunningNodes field value if set, zero value otherwise.
func (o *Cluster) GetMinRunningNodes() int32 {
	if o == nil || IsNil(o.MinRunningNodes) {
		var ret int32
		return ret
	}
	return *o.MinRunningNodes
}

// GetMinRunningNodesOk returns a tuple with the MinRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMinRunningNodesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinRunningNodes) {
		return nil, false
	}
	return o.MinRunningNodes, true
}

// HasMinRunningNodes returns a boolean if a field has been set.
func (o *Cluster) HasMinRunningNodes() bool {
	if o != nil && !IsNil(o.MinRunningNodes) {
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
	if o == nil || IsNil(o.MaxRunningNodes) {
		var ret int32
		return ret
	}
	return *o.MaxRunningNodes
}

// GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMaxRunningNodesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRunningNodes) {
		return nil, false
	}
	return o.MaxRunningNodes, true
}

// HasMaxRunningNodes returns a boolean if a field has been set.
func (o *Cluster) HasMaxRunningNodes() bool {
	if o != nil && !IsNil(o.MaxRunningNodes) {
		return true
	}

	return false
}

// SetMaxRunningNodes gets a reference to the given int32 and assigns it to the MaxRunningNodes field.
func (o *Cluster) SetMaxRunningNodes(v int32) {
	o.MaxRunningNodes = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *Cluster) GetDiskSize() int32 {
	if o == nil || IsNil(o.DiskSize) {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *Cluster) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *Cluster) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *Cluster) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *Cluster) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *Cluster) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetKubernetes returns the Kubernetes field value if set, zero value otherwise.
func (o *Cluster) GetKubernetes() KubernetesEnum {
	if o == nil || IsNil(o.Kubernetes) {
		var ret KubernetesEnum
		return ret
	}
	return *o.Kubernetes
}

// GetKubernetesOk returns a tuple with the Kubernetes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetKubernetesOk() (*KubernetesEnum, bool) {
	if o == nil || IsNil(o.Kubernetes) {
		return nil, false
	}
	return o.Kubernetes, true
}

// HasKubernetes returns a boolean if a field has been set.
func (o *Cluster) HasKubernetes() bool {
	if o != nil && !IsNil(o.Kubernetes) {
		return true
	}

	return false
}

// SetKubernetes gets a reference to the given KubernetesEnum and assigns it to the Kubernetes field.
func (o *Cluster) SetKubernetes(v KubernetesEnum) {
	o.Kubernetes = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Cluster) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Cluster) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *Cluster) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Cluster) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Cluster) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *Cluster) SetMemory(v int32) {
	o.Memory = &v
}

// GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field value if set, zero value otherwise.
func (o *Cluster) GetEstimatedCloudProviderCost() int32 {
	if o == nil || IsNil(o.EstimatedCloudProviderCost) {
		var ret int32
		return ret
	}
	return *o.EstimatedCloudProviderCost
}

// GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetEstimatedCloudProviderCostOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedCloudProviderCost) {
		return nil, false
	}
	return o.EstimatedCloudProviderCost, true
}

// HasEstimatedCloudProviderCost returns a boolean if a field has been set.
func (o *Cluster) HasEstimatedCloudProviderCost() bool {
	if o != nil && !IsNil(o.EstimatedCloudProviderCost) {
		return true
	}

	return false
}

// SetEstimatedCloudProviderCost gets a reference to the given int32 and assigns it to the EstimatedCloudProviderCost field.
func (o *Cluster) SetEstimatedCloudProviderCost(v int32) {
	o.EstimatedCloudProviderCost = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Cluster) GetStatus() ClusterStateEnum {
	if o == nil || IsNil(o.Status) {
		var ret ClusterStateEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetStatusOk() (*ClusterStateEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Cluster) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterStateEnum and assigns it to the Status field.
func (o *Cluster) SetStatus(v ClusterStateEnum) {
	o.Status = &v
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *Cluster) GetHasAccess() bool {
	if o == nil || IsNil(o.HasAccess) {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetHasAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccess) {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *Cluster) HasHasAccess() bool {
	if o != nil && !IsNil(o.HasAccess) {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *Cluster) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Cluster) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Cluster) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Cluster) SetVersion(v string) {
	o.Version = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *Cluster) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *Cluster) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *Cluster) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsDemo returns the IsDemo field value if set, zero value otherwise.
func (o *Cluster) GetIsDemo() bool {
	if o == nil || IsNil(o.IsDemo) {
		var ret bool
		return ret
	}
	return *o.IsDemo
}

// GetIsDemoOk returns a tuple with the IsDemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIsDemoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDemo) {
		return nil, false
	}
	return o.IsDemo, true
}

// HasIsDemo returns a boolean if a field has been set.
func (o *Cluster) HasIsDemo() bool {
	if o != nil && !IsNil(o.IsDemo) {
		return true
	}

	return false
}

// SetIsDemo gets a reference to the given bool and assigns it to the IsDemo field.
func (o *Cluster) SetIsDemo(v bool) {
	o.IsDemo = &v
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *Cluster) GetProduction() bool {
	if o == nil || IsNil(o.Production) {
		var ret bool
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetProductionOk() (*bool, bool) {
	if o == nil || IsNil(o.Production) {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *Cluster) HasProduction() bool {
	if o != nil && !IsNil(o.Production) {
		return true
	}

	return false
}

// SetProduction gets a reference to the given bool and assigns it to the Production field.
func (o *Cluster) SetProduction(v bool) {
	o.Production = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Cluster) GetSshKeys() []string {
	if o == nil || IsNil(o.SshKeys) {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetSshKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Cluster) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *Cluster) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Cluster) GetFeatures() []ClusterFeatureResponse {
	if o == nil || IsNil(o.Features) {
		var ret []ClusterFeatureResponse
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetFeaturesOk() ([]ClusterFeatureResponse, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Cluster) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterFeatureResponse and assigns it to the Features field.
func (o *Cluster) SetFeatures(v []ClusterFeatureResponse) {
	o.Features = v
}

// GetDeploymentStatus returns the DeploymentStatus field value if set, zero value otherwise.
func (o *Cluster) GetDeploymentStatus() ClusterDeploymentStatusEnum {
	if o == nil || IsNil(o.DeploymentStatus) {
		var ret ClusterDeploymentStatusEnum
		return ret
	}
	return *o.DeploymentStatus
}

// GetDeploymentStatusOk returns a tuple with the DeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDeploymentStatusOk() (*ClusterDeploymentStatusEnum, bool) {
	if o == nil || IsNil(o.DeploymentStatus) {
		return nil, false
	}
	return o.DeploymentStatus, true
}

// HasDeploymentStatus returns a boolean if a field has been set.
func (o *Cluster) HasDeploymentStatus() bool {
	if o != nil && !IsNil(o.DeploymentStatus) {
		return true
	}

	return false
}

// SetDeploymentStatus gets a reference to the given ClusterDeploymentStatusEnum and assigns it to the DeploymentStatus field.
func (o *Cluster) SetDeploymentStatus(v ClusterDeploymentStatusEnum) {
	o.DeploymentStatus = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["organization"] = o.Organization
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["region"] = o.Region
	toSerialize["cloud_provider"] = o.CloudProvider
	if !IsNil(o.MinRunningNodes) {
		toSerialize["min_running_nodes"] = o.MinRunningNodes
	}
	if !IsNil(o.MaxRunningNodes) {
		toSerialize["max_running_nodes"] = o.MaxRunningNodes
	}
	if !IsNil(o.DiskSize) {
		toSerialize["disk_size"] = o.DiskSize
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !IsNil(o.Kubernetes) {
		toSerialize["kubernetes"] = o.Kubernetes
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.EstimatedCloudProviderCost) {
		toSerialize["estimated_cloud_provider_cost"] = o.EstimatedCloudProviderCost
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.HasAccess) {
		toSerialize["has_access"] = o.HasAccess
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.IsDemo) {
		toSerialize["is_demo"] = o.IsDemo
	}
	if !IsNil(o.Production) {
		toSerialize["production"] = o.Production
	}
	if !IsNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.DeploymentStatus) {
		toSerialize["deployment_status"] = o.DeploymentStatus
	}
	return toSerialize, nil
}

func (o *Cluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"organization",
		"name",
		"region",
		"cloud_provider",
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

	varCluster := _Cluster{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCluster)

	if err != nil {
		return err
	}

	*o = Cluster(varCluster)

	return err
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
