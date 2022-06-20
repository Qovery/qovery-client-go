# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**CloudProvider** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | 
**Region** | **string** |  | 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**Kubernetes** | [**KubernetesEnum**](KubernetesEnum.md) |  | 
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**InstanceType** | Pointer to **string** | the instance type to be used for this cluster. The list of values can be retrieved via the endpoint /{CloudProvider}/instanceType | [optional] 
**DiskSize** | Pointer to **int32** | Unit is in GB. The disk size to be used for the node configuration | [optional] [default to 20]
**SshKey** | Pointer to [**ClusterBaseSshKey**](ClusterBaseSshKey.md) |  | [optional] 
**EstimatedCloudProviderCost** | Pointer to **int32** | This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**Features** | Pointer to [**[]ClusterFeature**](ClusterFeature.md) |  | [optional] 
**HasAccess** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, createdAt time.Time, name string, cloudProvider CloudProviderEnum, region string, kubernetes KubernetesEnum, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCloudProvider

`func (o *Cluster) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *Cluster) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *Cluster) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *Cluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Cluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Cluster) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAutoUpdate

`func (o *Cluster) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *Cluster) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *Cluster) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *Cluster) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetCpu

`func (o *Cluster) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Cluster) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Cluster) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Cluster) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *Cluster) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Cluster) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Cluster) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Cluster) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetKubernetes

`func (o *Cluster) GetKubernetes() KubernetesEnum`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *Cluster) GetKubernetesOk() (*KubernetesEnum, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *Cluster) SetKubernetes(v KubernetesEnum)`

SetKubernetes sets Kubernetes field to given value.


### GetMinRunningNodes

`func (o *Cluster) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *Cluster) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *Cluster) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *Cluster) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *Cluster) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *Cluster) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *Cluster) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *Cluster) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.

### GetInstanceType

`func (o *Cluster) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Cluster) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Cluster) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Cluster) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDiskSize

`func (o *Cluster) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *Cluster) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *Cluster) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *Cluster) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetSshKey

`func (o *Cluster) GetSshKey() ClusterBaseSshKey`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *Cluster) GetSshKeyOk() (*ClusterBaseSshKey, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *Cluster) SetSshKey(v ClusterBaseSshKey)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *Cluster) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetEstimatedCloudProviderCost

`func (o *Cluster) GetEstimatedCloudProviderCost() int32`

GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field if non-nil, zero value otherwise.

### GetEstimatedCloudProviderCostOk

`func (o *Cluster) GetEstimatedCloudProviderCostOk() (*int32, bool)`

GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCloudProviderCost

`func (o *Cluster) SetEstimatedCloudProviderCost(v int32)`

SetEstimatedCloudProviderCost sets EstimatedCloudProviderCost field to given value.

### HasEstimatedCloudProviderCost

`func (o *Cluster) HasEstimatedCloudProviderCost() bool`

HasEstimatedCloudProviderCost returns a boolean if a field has been set.

### GetStatus

`func (o *Cluster) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFeatures

`func (o *Cluster) GetFeatures() []ClusterFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Cluster) GetFeaturesOk() (*[]ClusterFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Cluster) SetFeatures(v []ClusterFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Cluster) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHasAccess

`func (o *Cluster) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *Cluster) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *Cluster) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *Cluster) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsDefault

`func (o *Cluster) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Cluster) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Cluster) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Cluster) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


