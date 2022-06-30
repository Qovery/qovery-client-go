# ClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**CloudProvider** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | 
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**DiskSize** | Pointer to **int32** | Unit is in GB. The disk size to be used for the node configuration | [optional] [default to 20]
**InstanceType** | Pointer to **string** | the instance type to be used for this cluster. The list of values can be retrieved via the endpoint /{CloudProvider}/instanceType | [optional] 
**Kubernetes** | Pointer to [**KubernetesEnum**](KubernetesEnum.md) |  | [optional] 
**Production** | Pointer to **bool** | specific flag to indicate that this cluster is a production one | [optional] 
**SshKeys** | Pointer to **[]string** | Indicate your public ssh_key to remotely connect to your EC2 instance. | [optional] 
**Features** | Pointer to [**[]ClusterRequestFeaturesInner**](ClusterRequestFeaturesInner.md) |  | [optional] 

## Methods

### NewClusterRequest

`func NewClusterRequest(name string, region string, cloudProvider CloudProviderEnum, ) *ClusterRequest`

NewClusterRequest instantiates a new ClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRequestWithDefaults

`func NewClusterRequestWithDefaults() *ClusterRequest`

NewClusterRequestWithDefaults instantiates a new ClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCloudProvider

`func (o *ClusterRequest) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterRequest) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterRequest) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.


### GetMinRunningNodes

`func (o *ClusterRequest) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *ClusterRequest) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *ClusterRequest) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *ClusterRequest) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *ClusterRequest) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *ClusterRequest) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *ClusterRequest) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *ClusterRequest) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.

### GetDiskSize

`func (o *ClusterRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ClusterRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ClusterRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ClusterRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetInstanceType

`func (o *ClusterRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetKubernetes

`func (o *ClusterRequest) GetKubernetes() KubernetesEnum`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *ClusterRequest) GetKubernetesOk() (*KubernetesEnum, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *ClusterRequest) SetKubernetes(v KubernetesEnum)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *ClusterRequest) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetProduction

`func (o *ClusterRequest) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *ClusterRequest) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *ClusterRequest) SetProduction(v bool)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *ClusterRequest) HasProduction() bool`

HasProduction returns a boolean if a field has been set.

### GetSshKeys

`func (o *ClusterRequest) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ClusterRequest) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ClusterRequest) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ClusterRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetFeatures

`func (o *ClusterRequest) GetFeatures() []ClusterRequestFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterRequest) GetFeaturesOk() (*[]ClusterRequestFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterRequest) SetFeatures(v []ClusterRequestFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


