# InstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseCpu**](EnvironmentDatabasesCurrentMetricResponseCpu.md) |  | [optional] 
**Memory** | Pointer to [**EnvironmentDatabasesCurrentMetricResponseMemory**](EnvironmentDatabasesCurrentMetricResponseMemory.md) |  | [optional] 

## Methods

### NewInstanceResponse

`func NewInstanceResponse() *InstanceResponse`

NewInstanceResponse instantiates a new InstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResponseWithDefaults

`func NewInstanceResponseWithDefaults() *InstanceResponse`

NewInstanceResponseWithDefaults instantiates a new InstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InstanceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *InstanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpu

`func (o *InstanceResponse) GetCpu() EnvironmentDatabasesCurrentMetricResponseCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceResponse) GetCpuOk() (*EnvironmentDatabasesCurrentMetricResponseCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceResponse) SetCpu(v EnvironmentDatabasesCurrentMetricResponseCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InstanceResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *InstanceResponse) GetMemory() EnvironmentDatabasesCurrentMetricResponseMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InstanceResponse) GetMemoryOk() (*EnvironmentDatabasesCurrentMetricResponseMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InstanceResponse) SetMemory(v EnvironmentDatabasesCurrentMetricResponseMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InstanceResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


