# DatabaseCurrentMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**DatabaseCurrentMetricCpu**](DatabaseCurrentMetricCpu.md) |  | [optional] 
**Memory** | Pointer to [**DatabaseCurrentMetricMemory**](DatabaseCurrentMetricMemory.md) |  | [optional] 
**Storage** | Pointer to [**DatabaseCurrentMetricStorage**](DatabaseCurrentMetricStorage.md) |  | [optional] 

## Methods

### NewDatabaseCurrentMetric

`func NewDatabaseCurrentMetric() *DatabaseCurrentMetric`

NewDatabaseCurrentMetric instantiates a new DatabaseCurrentMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCurrentMetricWithDefaults

`func NewDatabaseCurrentMetricWithDefaults() *DatabaseCurrentMetric`

NewDatabaseCurrentMetricWithDefaults instantiates a new DatabaseCurrentMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *DatabaseCurrentMetric) GetCpu() DatabaseCurrentMetricCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseCurrentMetric) GetCpuOk() (*DatabaseCurrentMetricCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseCurrentMetric) SetCpu(v DatabaseCurrentMetricCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DatabaseCurrentMetric) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *DatabaseCurrentMetric) GetMemory() DatabaseCurrentMetricMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DatabaseCurrentMetric) GetMemoryOk() (*DatabaseCurrentMetricMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DatabaseCurrentMetric) SetMemory(v DatabaseCurrentMetricMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DatabaseCurrentMetric) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *DatabaseCurrentMetric) GetStorage() DatabaseCurrentMetricStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DatabaseCurrentMetric) GetStorageOk() (*DatabaseCurrentMetricStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DatabaseCurrentMetric) SetStorage(v DatabaseCurrentMetricStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DatabaseCurrentMetric) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


