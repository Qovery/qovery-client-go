# ClusterInstanceGpuInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of GPUs included in the instance. | [optional] 
**Name** | Pointer to **string** | The model name of the GPU. | [optional] 
**Manufacturer** | Pointer to **string** | The manufacturer of the GPUs in the instance. | [optional] 
**MemoryInMib** | Pointer to **string** | The total GPU memory available. | [optional] 

## Methods

### NewClusterInstanceGpuInfo

`func NewClusterInstanceGpuInfo() *ClusterInstanceGpuInfo`

NewClusterInstanceGpuInfo instantiates a new ClusterInstanceGpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstanceGpuInfoWithDefaults

`func NewClusterInstanceGpuInfoWithDefaults() *ClusterInstanceGpuInfo`

NewClusterInstanceGpuInfoWithDefaults instantiates a new ClusterInstanceGpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ClusterInstanceGpuInfo) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClusterInstanceGpuInfo) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClusterInstanceGpuInfo) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ClusterInstanceGpuInfo) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetName

`func (o *ClusterInstanceGpuInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInstanceGpuInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInstanceGpuInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterInstanceGpuInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetManufacturer

`func (o *ClusterInstanceGpuInfo) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ClusterInstanceGpuInfo) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ClusterInstanceGpuInfo) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ClusterInstanceGpuInfo) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetMemoryInMib

`func (o *ClusterInstanceGpuInfo) GetMemoryInMib() string`

GetMemoryInMib returns the MemoryInMib field if non-nil, zero value otherwise.

### GetMemoryInMibOk

`func (o *ClusterInstanceGpuInfo) GetMemoryInMibOk() (*string, bool)`

GetMemoryInMibOk returns a tuple with the MemoryInMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMib

`func (o *ClusterInstanceGpuInfo) SetMemoryInMib(v string)`

SetMemoryInMib sets MemoryInMib field to given value.

### HasMemoryInMib

`func (o *ClusterInstanceGpuInfo) HasMemoryInMib() bool`

HasMemoryInMib returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


