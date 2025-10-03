# TerraformJobResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilli** | **int32** |  | 
**RamMib** | **int32** |  | 
**Gpu** | **int32** |  | [default to 0]
**StorageGib** | **int32** |  | 

## Methods

### NewTerraformJobResourcesResponse

`func NewTerraformJobResourcesResponse(cpuMilli int32, ramMib int32, gpu int32, storageGib int32, ) *TerraformJobResourcesResponse`

NewTerraformJobResourcesResponse instantiates a new TerraformJobResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformJobResourcesResponseWithDefaults

`func NewTerraformJobResourcesResponseWithDefaults() *TerraformJobResourcesResponse`

NewTerraformJobResourcesResponseWithDefaults instantiates a new TerraformJobResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilli

`func (o *TerraformJobResourcesResponse) GetCpuMilli() int32`

GetCpuMilli returns the CpuMilli field if non-nil, zero value otherwise.

### GetCpuMilliOk

`func (o *TerraformJobResourcesResponse) GetCpuMilliOk() (*int32, bool)`

GetCpuMilliOk returns a tuple with the CpuMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilli

`func (o *TerraformJobResourcesResponse) SetCpuMilli(v int32)`

SetCpuMilli sets CpuMilli field to given value.


### GetRamMib

`func (o *TerraformJobResourcesResponse) GetRamMib() int32`

GetRamMib returns the RamMib field if non-nil, zero value otherwise.

### GetRamMibOk

`func (o *TerraformJobResourcesResponse) GetRamMibOk() (*int32, bool)`

GetRamMibOk returns a tuple with the RamMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMib

`func (o *TerraformJobResourcesResponse) SetRamMib(v int32)`

SetRamMib sets RamMib field to given value.


### GetGpu

`func (o *TerraformJobResourcesResponse) GetGpu() int32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *TerraformJobResourcesResponse) GetGpuOk() (*int32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *TerraformJobResourcesResponse) SetGpu(v int32)`

SetGpu sets Gpu field to given value.


### GetStorageGib

`func (o *TerraformJobResourcesResponse) GetStorageGib() int32`

GetStorageGib returns the StorageGib field if non-nil, zero value otherwise.

### GetStorageGibOk

`func (o *TerraformJobResourcesResponse) GetStorageGibOk() (*int32, bool)`

GetStorageGibOk returns a tuple with the StorageGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGib

`func (o *TerraformJobResourcesResponse) SetStorageGib(v int32)`

SetStorageGib sets StorageGib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


