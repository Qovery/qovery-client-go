# EnvironmentDatabasesCurrentMetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseCpu**](EnvironmentDatabasesCurrentMetricResponseCpu.md) |  | [optional] 
**Memory** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseMemory**](EnvironmentDatabasesCurrentMetricResponseMemory.md) |  | [optional] 
**Storage** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseStorage**](EnvironmentDatabasesCurrentMetricResponseStorage.md) |  | [optional] 

## Methods

### NewEnvironmentDatabasesCurrentMetricResponse

`func NewEnvironmentDatabasesCurrentMetricResponse() *EnvironmentDatabasesCurrentMetricResponse`

NewEnvironmentDatabasesCurrentMetricResponse instantiates a new EnvironmentDatabasesCurrentMetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDatabasesCurrentMetricResponseWithDefaults

`func NewEnvironmentDatabasesCurrentMetricResponseWithDefaults() *EnvironmentDatabasesCurrentMetricResponse`

NewEnvironmentDatabasesCurrentMetricResponseWithDefaults instantiates a new EnvironmentDatabasesCurrentMetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *EnvironmentDatabasesCurrentMetricResponse) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *EnvironmentDatabasesCurrentMetricResponse) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetCpu

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetCpu() EnvironmentDatabasesCurrentMetricResponseCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetCpuOk() (*EnvironmentDatabasesCurrentMetricResponseCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentDatabasesCurrentMetricResponse) SetCpu(v EnvironmentDatabasesCurrentMetricResponseCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentDatabasesCurrentMetricResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetMemory() EnvironmentDatabasesCurrentMetricResponseMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetMemoryOk() (*EnvironmentDatabasesCurrentMetricResponseMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentDatabasesCurrentMetricResponse) SetMemory(v EnvironmentDatabasesCurrentMetricResponseMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentDatabasesCurrentMetricResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetStorage() EnvironmentDatabasesCurrentMetricResponseStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentDatabasesCurrentMetricResponse) GetStorageOk() (*EnvironmentDatabasesCurrentMetricResponseStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentDatabasesCurrentMetricResponse) SetStorage(v EnvironmentDatabasesCurrentMetricResponseStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentDatabasesCurrentMetricResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


