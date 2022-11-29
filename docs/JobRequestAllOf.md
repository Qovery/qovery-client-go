# JobRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 500]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 512]
**MaxNbRestart** | Pointer to **int32** | Maximum number of restart allowed before the job is considered as failed 0 means that no restart/crash of the job is allowed  | [optional] [default to 0]
**MaxDurationSeconds** | Pointer to **int32** | Maximum number of seconds allowed for the job to run before killing it and mark it as failed  | [optional] 
**AutoPreview** | Pointer to **bool** | Indicates if the &#39;environment preview option&#39; is enabled for this container.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | [optional] 
**Port** | Pointer to **NullableInt32** | Port where to run readiness and liveliness probes checks. The port will not be exposed externally | [optional] 
**Source** | Pointer to [**JobRequestAllOfSource**](JobRequestAllOfSource.md) |  | [optional] 
**Schedule** | Pointer to [**JobRequestAllOfSchedule**](JobRequestAllOfSchedule.md) |  | [optional] 

## Methods

### NewJobRequestAllOf

`func NewJobRequestAllOf(name string, ) *JobRequestAllOf`

NewJobRequestAllOf instantiates a new JobRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestAllOfWithDefaults

`func NewJobRequestAllOfWithDefaults() *JobRequestAllOf`

NewJobRequestAllOfWithDefaults instantiates a new JobRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobRequestAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *JobRequestAllOf) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JobRequestAllOf) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JobRequestAllOf) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *JobRequestAllOf) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *JobRequestAllOf) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *JobRequestAllOf) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *JobRequestAllOf) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *JobRequestAllOf) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetCpu

`func (o *JobRequestAllOf) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *JobRequestAllOf) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *JobRequestAllOf) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *JobRequestAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *JobRequestAllOf) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JobRequestAllOf) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JobRequestAllOf) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *JobRequestAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMaxNbRestart

`func (o *JobRequestAllOf) GetMaxNbRestart() int32`

GetMaxNbRestart returns the MaxNbRestart field if non-nil, zero value otherwise.

### GetMaxNbRestartOk

`func (o *JobRequestAllOf) GetMaxNbRestartOk() (*int32, bool)`

GetMaxNbRestartOk returns a tuple with the MaxNbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNbRestart

`func (o *JobRequestAllOf) SetMaxNbRestart(v int32)`

SetMaxNbRestart sets MaxNbRestart field to given value.

### HasMaxNbRestart

`func (o *JobRequestAllOf) HasMaxNbRestart() bool`

HasMaxNbRestart returns a boolean if a field has been set.

### GetMaxDurationSeconds

`func (o *JobRequestAllOf) GetMaxDurationSeconds() int32`

GetMaxDurationSeconds returns the MaxDurationSeconds field if non-nil, zero value otherwise.

### GetMaxDurationSecondsOk

`func (o *JobRequestAllOf) GetMaxDurationSecondsOk() (*int32, bool)`

GetMaxDurationSecondsOk returns a tuple with the MaxDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationSeconds

`func (o *JobRequestAllOf) SetMaxDurationSeconds(v int32)`

SetMaxDurationSeconds sets MaxDurationSeconds field to given value.

### HasMaxDurationSeconds

`func (o *JobRequestAllOf) HasMaxDurationSeconds() bool`

HasMaxDurationSeconds returns a boolean if a field has been set.

### GetAutoPreview

`func (o *JobRequestAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *JobRequestAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *JobRequestAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *JobRequestAllOf) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetPort

`func (o *JobRequestAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *JobRequestAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *JobRequestAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *JobRequestAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *JobRequestAllOf) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *JobRequestAllOf) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSource

`func (o *JobRequestAllOf) GetSource() JobRequestAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JobRequestAllOf) GetSourceOk() (*JobRequestAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JobRequestAllOf) SetSource(v JobRequestAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *JobRequestAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSchedule

`func (o *JobRequestAllOf) GetSchedule() JobRequestAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobRequestAllOf) GetScheduleOk() (*JobRequestAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobRequestAllOf) SetSchedule(v JobRequestAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *JobRequestAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


