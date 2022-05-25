# EnvironmentDatabasesCurrentMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to [**DatabaseCurrentMetricCpu**](DatabaseCurrentMetricCpu.md) |  | [optional] 
**Memory** | Pointer to [**DatabaseCurrentMetricMemory**](DatabaseCurrentMetricMemory.md) |  | [optional] 
**Storage** | Pointer to [**DatabaseCurrentMetricStorage**](DatabaseCurrentMetricStorage.md) |  | [optional] 

## Methods

### NewEnvironmentDatabasesCurrentMetric

`func NewEnvironmentDatabasesCurrentMetric() *EnvironmentDatabasesCurrentMetric`

NewEnvironmentDatabasesCurrentMetric instantiates a new EnvironmentDatabasesCurrentMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDatabasesCurrentMetricWithDefaults

`func NewEnvironmentDatabasesCurrentMetricWithDefaults() *EnvironmentDatabasesCurrentMetric`

NewEnvironmentDatabasesCurrentMetricWithDefaults instantiates a new EnvironmentDatabasesCurrentMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *EnvironmentDatabasesCurrentMetric) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *EnvironmentDatabasesCurrentMetric) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *EnvironmentDatabasesCurrentMetric) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *EnvironmentDatabasesCurrentMetric) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetCpu

`func (o *EnvironmentDatabasesCurrentMetric) GetCpu() DatabaseCurrentMetricCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentDatabasesCurrentMetric) GetCpuOk() (*DatabaseCurrentMetricCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentDatabasesCurrentMetric) SetCpu(v DatabaseCurrentMetricCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentDatabasesCurrentMetric) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *EnvironmentDatabasesCurrentMetric) GetMemory() DatabaseCurrentMetricMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentDatabasesCurrentMetric) GetMemoryOk() (*DatabaseCurrentMetricMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentDatabasesCurrentMetric) SetMemory(v DatabaseCurrentMetricMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentDatabasesCurrentMetric) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *EnvironmentDatabasesCurrentMetric) GetStorage() DatabaseCurrentMetricStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentDatabasesCurrentMetric) GetStorageOk() (*DatabaseCurrentMetricStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentDatabasesCurrentMetric) SetStorage(v DatabaseCurrentMetricStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentDatabasesCurrentMetric) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


