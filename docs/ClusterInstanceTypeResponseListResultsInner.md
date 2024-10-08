# ClusterInstanceTypeResponseListResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Cpu** | **int32** |  | 
**RamInGb** | **int32** |  | 
**BandwidthInGbps** | **string** |  | 
**BandwidthGuarantee** | **string** |  | 
**Architecture** | Pointer to **string** |  | [optional] 
**GpuInfo** | Pointer to [**ClusterInstanceGpuInfo**](ClusterInstanceGpuInfo.md) |  | [optional] 
**Attributes** | Pointer to [**ClusterInstanceAttributes**](ClusterInstanceAttributes.md) |  | [optional] 

## Methods

### NewClusterInstanceTypeResponseListResultsInner

`func NewClusterInstanceTypeResponseListResultsInner(type_ string, name string, cpu int32, ramInGb int32, bandwidthInGbps string, bandwidthGuarantee string, ) *ClusterInstanceTypeResponseListResultsInner`

NewClusterInstanceTypeResponseListResultsInner instantiates a new ClusterInstanceTypeResponseListResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstanceTypeResponseListResultsInnerWithDefaults

`func NewClusterInstanceTypeResponseListResultsInnerWithDefaults() *ClusterInstanceTypeResponseListResultsInner`

NewClusterInstanceTypeResponseListResultsInnerWithDefaults instantiates a new ClusterInstanceTypeResponseListResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterInstanceTypeResponseListResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterInstanceTypeResponseListResultsInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ClusterInstanceTypeResponseListResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInstanceTypeResponseListResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetCpu

`func (o *ClusterInstanceTypeResponseListResultsInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterInstanceTypeResponseListResultsInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRamInGb

`func (o *ClusterInstanceTypeResponseListResultsInner) GetRamInGb() int32`

GetRamInGb returns the RamInGb field if non-nil, zero value otherwise.

### GetRamInGbOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetRamInGbOk() (*int32, bool)`

GetRamInGbOk returns a tuple with the RamInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamInGb

`func (o *ClusterInstanceTypeResponseListResultsInner) SetRamInGb(v int32)`

SetRamInGb sets RamInGb field to given value.


### GetBandwidthInGbps

`func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthInGbps() string`

GetBandwidthInGbps returns the BandwidthInGbps field if non-nil, zero value otherwise.

### GetBandwidthInGbpsOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthInGbpsOk() (*string, bool)`

GetBandwidthInGbpsOk returns a tuple with the BandwidthInGbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthInGbps

`func (o *ClusterInstanceTypeResponseListResultsInner) SetBandwidthInGbps(v string)`

SetBandwidthInGbps sets BandwidthInGbps field to given value.


### GetBandwidthGuarantee

`func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthGuarantee() string`

GetBandwidthGuarantee returns the BandwidthGuarantee field if non-nil, zero value otherwise.

### GetBandwidthGuaranteeOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthGuaranteeOk() (*string, bool)`

GetBandwidthGuaranteeOk returns a tuple with the BandwidthGuarantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthGuarantee

`func (o *ClusterInstanceTypeResponseListResultsInner) SetBandwidthGuarantee(v string)`

SetBandwidthGuarantee sets BandwidthGuarantee field to given value.


### GetArchitecture

`func (o *ClusterInstanceTypeResponseListResultsInner) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ClusterInstanceTypeResponseListResultsInner) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ClusterInstanceTypeResponseListResultsInner) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetGpuInfo

`func (o *ClusterInstanceTypeResponseListResultsInner) GetGpuInfo() ClusterInstanceGpuInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetGpuInfoOk() (*ClusterInstanceGpuInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *ClusterInstanceTypeResponseListResultsInner) SetGpuInfo(v ClusterInstanceGpuInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *ClusterInstanceTypeResponseListResultsInner) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetAttributes

`func (o *ClusterInstanceTypeResponseListResultsInner) GetAttributes() ClusterInstanceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ClusterInstanceTypeResponseListResultsInner) GetAttributesOk() (*ClusterInstanceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ClusterInstanceTypeResponseListResultsInner) SetAttributes(v ClusterInstanceAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ClusterInstanceTypeResponseListResultsInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


