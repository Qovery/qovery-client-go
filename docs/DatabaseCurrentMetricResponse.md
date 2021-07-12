# DatabaseCurrentMetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseCpu**](EnvironmentDatabasesCurrentMetricResponseCpu.md) |  | [optional] 
**Memory** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseMemory**](EnvironmentDatabasesCurrentMetricResponseMemory.md) |  | [optional] 
**Storage** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseStorage**](EnvironmentDatabasesCurrentMetricResponseStorage.md) |  | [optional] 

## Methods

### NewDatabaseCurrentMetricResponse

`func NewDatabaseCurrentMetricResponse() *DatabaseCurrentMetricResponse`

NewDatabaseCurrentMetricResponse instantiates a new DatabaseCurrentMetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCurrentMetricResponseWithDefaults

`func NewDatabaseCurrentMetricResponseWithDefaults() *DatabaseCurrentMetricResponse`

NewDatabaseCurrentMetricResponseWithDefaults instantiates a new DatabaseCurrentMetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *DatabaseCurrentMetricResponse) GetCpu() EnvironmentDatabasesCurrentMetricResponseCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseCurrentMetricResponse) GetCpuOk() (*EnvironmentDatabasesCurrentMetricResponseCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseCurrentMetricResponse) SetCpu(v EnvironmentDatabasesCurrentMetricResponseCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DatabaseCurrentMetricResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *DatabaseCurrentMetricResponse) GetMemory() EnvironmentDatabasesCurrentMetricResponseMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DatabaseCurrentMetricResponse) GetMemoryOk() (*EnvironmentDatabasesCurrentMetricResponseMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DatabaseCurrentMetricResponse) SetMemory(v EnvironmentDatabasesCurrentMetricResponseMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DatabaseCurrentMetricResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *DatabaseCurrentMetricResponse) GetStorage() EnvironmentDatabasesCurrentMetricResponseStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DatabaseCurrentMetricResponse) GetStorageOk() (*EnvironmentDatabasesCurrentMetricResponseStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DatabaseCurrentMetricResponse) SetStorage(v EnvironmentDatabasesCurrentMetricResponseStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DatabaseCurrentMetricResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


