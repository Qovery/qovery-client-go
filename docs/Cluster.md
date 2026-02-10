# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Organization** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**CloudProvider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**DiskSize** | Pointer to **int32** | Unit is in GB. The disk size to be used for the node configuration | [optional] [default to 20]
**DiskIops** | Pointer to **int32** | Unit is operation/seconds. The disk IOPS to be used for the node configuration | [optional] 
**DiskThroughput** | Pointer to **int32** | Unit is in MB/s. The disk thoughput to be used for the node configuration | [optional] 
**InstanceType** | Pointer to **string** | the instance type to be used for this cluster. The list of values can be retrieved via the endpoint /{CloudProvider}/instanceType | [optional] 
**Kubernetes** | Pointer to [**KubernetesEnum**](KubernetesEnum.md) |  | [optional] [default to KUBERNETESENUM_MANAGED]
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] 
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] 
**EstimatedCloudProviderCost** | Pointer to **int32** | This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration | [optional] 
**Status** | Pointer to [**ClusterStateEnum**](ClusterStateEnum.md) |  | [optional] 
**HasAccess** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsDemo** | Pointer to **bool** | specific flag to indicate that this cluster is a demo one | [optional] 
**Production** | Pointer to **bool** | specific flag to indicate that this cluster is a production one | [optional] 
**SshKeys** | Pointer to **[]string** | Indicate your public ssh_key to remotely connect to your EC2 instance. | [optional] 
**Features** | Pointer to [**[]ClusterFeatureResponse**](ClusterFeatureResponse.md) |  | [optional] 
**DeploymentStatus** | Pointer to [**ClusterDeploymentStatusEnum**](ClusterDeploymentStatusEnum.md) |  | [optional] 
**MetricsParameters** | Pointer to [**MetricsParameters**](MetricsParameters.md) |  | [optional] 
**InfrastructureOutputs** | Pointer to [**InfrastructureOutputs**](InfrastructureOutputs.md) |  | [optional] 
**InfrastructureChartsParameters** | Pointer to [**ClusterInfrastructureChartsParameters**](ClusterInfrastructureChartsParameters.md) |  | [optional] 
**Keda** | Pointer to [**ClusterKeda**](ClusterKeda.md) |  | [optional] 
**LabelsGroups** | Pointer to [**ClusterLabelsGroupList**](ClusterLabelsGroupList.md) |  | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, createdAt time.Time, organization ReferenceObject, name string, region string, cloudProvider CloudVendorEnum, ) *Cluster`

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

### GetOrganization

`func (o *Cluster) GetOrganization() ReferenceObject`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Cluster) GetOrganizationOk() (*ReferenceObject, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Cluster) SetOrganization(v ReferenceObject)`

SetOrganization sets Organization field to given value.


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


### GetCloudProvider

`func (o *Cluster) GetCloudProvider() CloudVendorEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *Cluster) GetCloudProviderOk() (*CloudVendorEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *Cluster) SetCloudProvider(v CloudVendorEnum)`

SetCloudProvider sets CloudProvider field to given value.


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

### GetDiskIops

`func (o *Cluster) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *Cluster) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *Cluster) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *Cluster) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.

### GetDiskThroughput

`func (o *Cluster) GetDiskThroughput() int32`

GetDiskThroughput returns the DiskThroughput field if non-nil, zero value otherwise.

### GetDiskThroughputOk

`func (o *Cluster) GetDiskThroughputOk() (*int32, bool)`

GetDiskThroughputOk returns a tuple with the DiskThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskThroughput

`func (o *Cluster) SetDiskThroughput(v int32)`

SetDiskThroughput sets DiskThroughput field to given value.

### HasDiskThroughput

`func (o *Cluster) HasDiskThroughput() bool`

HasDiskThroughput returns a boolean if a field has been set.

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

### HasKubernetes

`func (o *Cluster) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

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

`func (o *Cluster) GetStatus() ClusterStateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*ClusterStateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v ClusterStateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### GetIsDemo

`func (o *Cluster) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *Cluster) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *Cluster) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *Cluster) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetProduction

`func (o *Cluster) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *Cluster) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *Cluster) SetProduction(v bool)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *Cluster) HasProduction() bool`

HasProduction returns a boolean if a field has been set.

### GetSshKeys

`func (o *Cluster) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Cluster) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Cluster) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Cluster) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetFeatures

`func (o *Cluster) GetFeatures() []ClusterFeatureResponse`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Cluster) GetFeaturesOk() (*[]ClusterFeatureResponse, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Cluster) SetFeatures(v []ClusterFeatureResponse)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Cluster) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDeploymentStatus

`func (o *Cluster) GetDeploymentStatus() ClusterDeploymentStatusEnum`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *Cluster) GetDeploymentStatusOk() (*ClusterDeploymentStatusEnum, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *Cluster) SetDeploymentStatus(v ClusterDeploymentStatusEnum)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *Cluster) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.

### GetMetricsParameters

`func (o *Cluster) GetMetricsParameters() MetricsParameters`

GetMetricsParameters returns the MetricsParameters field if non-nil, zero value otherwise.

### GetMetricsParametersOk

`func (o *Cluster) GetMetricsParametersOk() (*MetricsParameters, bool)`

GetMetricsParametersOk returns a tuple with the MetricsParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsParameters

`func (o *Cluster) SetMetricsParameters(v MetricsParameters)`

SetMetricsParameters sets MetricsParameters field to given value.

### HasMetricsParameters

`func (o *Cluster) HasMetricsParameters() bool`

HasMetricsParameters returns a boolean if a field has been set.

### GetInfrastructureOutputs

`func (o *Cluster) GetInfrastructureOutputs() InfrastructureOutputs`

GetInfrastructureOutputs returns the InfrastructureOutputs field if non-nil, zero value otherwise.

### GetInfrastructureOutputsOk

`func (o *Cluster) GetInfrastructureOutputsOk() (*InfrastructureOutputs, bool)`

GetInfrastructureOutputsOk returns a tuple with the InfrastructureOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureOutputs

`func (o *Cluster) SetInfrastructureOutputs(v InfrastructureOutputs)`

SetInfrastructureOutputs sets InfrastructureOutputs field to given value.

### HasInfrastructureOutputs

`func (o *Cluster) HasInfrastructureOutputs() bool`

HasInfrastructureOutputs returns a boolean if a field has been set.

### GetInfrastructureChartsParameters

`func (o *Cluster) GetInfrastructureChartsParameters() ClusterInfrastructureChartsParameters`

GetInfrastructureChartsParameters returns the InfrastructureChartsParameters field if non-nil, zero value otherwise.

### GetInfrastructureChartsParametersOk

`func (o *Cluster) GetInfrastructureChartsParametersOk() (*ClusterInfrastructureChartsParameters, bool)`

GetInfrastructureChartsParametersOk returns a tuple with the InfrastructureChartsParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureChartsParameters

`func (o *Cluster) SetInfrastructureChartsParameters(v ClusterInfrastructureChartsParameters)`

SetInfrastructureChartsParameters sets InfrastructureChartsParameters field to given value.

### HasInfrastructureChartsParameters

`func (o *Cluster) HasInfrastructureChartsParameters() bool`

HasInfrastructureChartsParameters returns a boolean if a field has been set.

### GetKeda

`func (o *Cluster) GetKeda() ClusterKeda`

GetKeda returns the Keda field if non-nil, zero value otherwise.

### GetKedaOk

`func (o *Cluster) GetKedaOk() (*ClusterKeda, bool)`

GetKedaOk returns a tuple with the Keda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeda

`func (o *Cluster) SetKeda(v ClusterKeda)`

SetKeda sets Keda field to given value.

### HasKeda

`func (o *Cluster) HasKeda() bool`

HasKeda returns a boolean if a field has been set.

### GetLabelsGroups

`func (o *Cluster) GetLabelsGroups() ClusterLabelsGroupList`

GetLabelsGroups returns the LabelsGroups field if non-nil, zero value otherwise.

### GetLabelsGroupsOk

`func (o *Cluster) GetLabelsGroupsOk() (*ClusterLabelsGroupList, bool)`

GetLabelsGroupsOk returns a tuple with the LabelsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsGroups

`func (o *Cluster) SetLabelsGroups(v ClusterLabelsGroupList)`

SetLabelsGroups sets LabelsGroups field to given value.

### HasLabelsGroups

`func (o *Cluster) HasLabelsGroups() bool`

HasLabelsGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


