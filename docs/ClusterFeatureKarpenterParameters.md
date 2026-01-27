# ClusterFeatureKarpenterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpotEnabled** | **bool** |  | 
**DiskSizeInGib** | **int32** |  | 
**DiskIops** | Pointer to **int32** |  | [optional] 
**DiskThrouput** | Pointer to **int32** |  | [optional] 
**DefaultServiceArchitecture** | [**CpuArchitectureEnum**](CpuArchitectureEnum.md) |  | 
**QoveryNodePools** | [**KarpenterNodePool**](KarpenterNodePool.md) |  | 

## Methods

### NewClusterFeatureKarpenterParameters

`func NewClusterFeatureKarpenterParameters(spotEnabled bool, diskSizeInGib int32, defaultServiceArchitecture CpuArchitectureEnum, qoveryNodePools KarpenterNodePool, ) *ClusterFeatureKarpenterParameters`

NewClusterFeatureKarpenterParameters instantiates a new ClusterFeatureKarpenterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureKarpenterParametersWithDefaults

`func NewClusterFeatureKarpenterParametersWithDefaults() *ClusterFeatureKarpenterParameters`

NewClusterFeatureKarpenterParametersWithDefaults instantiates a new ClusterFeatureKarpenterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpotEnabled

`func (o *ClusterFeatureKarpenterParameters) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *ClusterFeatureKarpenterParameters) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *ClusterFeatureKarpenterParameters) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.


### GetDiskSizeInGib

`func (o *ClusterFeatureKarpenterParameters) GetDiskSizeInGib() int32`

GetDiskSizeInGib returns the DiskSizeInGib field if non-nil, zero value otherwise.

### GetDiskSizeInGibOk

`func (o *ClusterFeatureKarpenterParameters) GetDiskSizeInGibOk() (*int32, bool)`

GetDiskSizeInGibOk returns a tuple with the DiskSizeInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInGib

`func (o *ClusterFeatureKarpenterParameters) SetDiskSizeInGib(v int32)`

SetDiskSizeInGib sets DiskSizeInGib field to given value.


### GetDiskIops

`func (o *ClusterFeatureKarpenterParameters) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *ClusterFeatureKarpenterParameters) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *ClusterFeatureKarpenterParameters) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *ClusterFeatureKarpenterParameters) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.

### GetDiskThrouput

`func (o *ClusterFeatureKarpenterParameters) GetDiskThrouput() int32`

GetDiskThrouput returns the DiskThrouput field if non-nil, zero value otherwise.

### GetDiskThrouputOk

`func (o *ClusterFeatureKarpenterParameters) GetDiskThrouputOk() (*int32, bool)`

GetDiskThrouputOk returns a tuple with the DiskThrouput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskThrouput

`func (o *ClusterFeatureKarpenterParameters) SetDiskThrouput(v int32)`

SetDiskThrouput sets DiskThrouput field to given value.

### HasDiskThrouput

`func (o *ClusterFeatureKarpenterParameters) HasDiskThrouput() bool`

HasDiskThrouput returns a boolean if a field has been set.

### GetDefaultServiceArchitecture

`func (o *ClusterFeatureKarpenterParameters) GetDefaultServiceArchitecture() CpuArchitectureEnum`

GetDefaultServiceArchitecture returns the DefaultServiceArchitecture field if non-nil, zero value otherwise.

### GetDefaultServiceArchitectureOk

`func (o *ClusterFeatureKarpenterParameters) GetDefaultServiceArchitectureOk() (*CpuArchitectureEnum, bool)`

GetDefaultServiceArchitectureOk returns a tuple with the DefaultServiceArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceArchitecture

`func (o *ClusterFeatureKarpenterParameters) SetDefaultServiceArchitecture(v CpuArchitectureEnum)`

SetDefaultServiceArchitecture sets DefaultServiceArchitecture field to given value.


### GetQoveryNodePools

`func (o *ClusterFeatureKarpenterParameters) GetQoveryNodePools() KarpenterNodePool`

GetQoveryNodePools returns the QoveryNodePools field if non-nil, zero value otherwise.

### GetQoveryNodePoolsOk

`func (o *ClusterFeatureKarpenterParameters) GetQoveryNodePoolsOk() (*KarpenterNodePool, bool)`

GetQoveryNodePoolsOk returns a tuple with the QoveryNodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryNodePools

`func (o *ClusterFeatureKarpenterParameters) SetQoveryNodePools(v KarpenterNodePool)`

SetQoveryNodePools sets QoveryNodePools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


