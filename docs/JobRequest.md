# JobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 500]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 512]
**MaxNbRestart** | Pointer to **int32** | Maximum number of restart allowed before the job is considered as failed 0 means that no restart/crash of the job is allowed  | [optional] [default to 0]
**MaxDurationSeconds** | Pointer to **int32** | Maximum number of seconds allowed for the job to run before killing it and mark it as failed  | [optional] 
**AutoPreview** | Pointer to **bool** | Indicates if the &#39;environment preview option&#39; is enabled for this container.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | [optional] 
**Port** | Pointer to **NullableInt32** | Port where to run readiness and liveliness probes checks. The port will not be exposed externally | [optional] 
**Source** | Pointer to [**JobRequestAllOfSource**](JobRequestAllOfSource.md) |  | [optional] 
**Healthchecks** | Pointer to [**Healthcheck**](Healthcheck.md) |  | [optional] 
**Schedule** | Pointer to [**JobRequestAllOfSchedule**](JobRequestAllOfSchedule.md) |  | [optional] 
**AutoDeploy** | Pointer to **NullableBool** | Specify if the job will be automatically updated after receiving a new image tag or a new commit according to the source type.  The new image tag shall be communicated via the \&quot;Auto Deploy job\&quot; endpoint https://api-doc.qovery.com/#tag/Jobs/operation/autoDeployJobEnvironments  | [optional] 

## Methods

### NewJobRequest

`func NewJobRequest(name string, ) *JobRequest`

NewJobRequest instantiates a new JobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestWithDefaults

`func NewJobRequestWithDefaults() *JobRequest`

NewJobRequestWithDefaults instantiates a new JobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCpu

`func (o *JobRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *JobRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *JobRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *JobRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *JobRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JobRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JobRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *JobRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMaxNbRestart

`func (o *JobRequest) GetMaxNbRestart() int32`

GetMaxNbRestart returns the MaxNbRestart field if non-nil, zero value otherwise.

### GetMaxNbRestartOk

`func (o *JobRequest) GetMaxNbRestartOk() (*int32, bool)`

GetMaxNbRestartOk returns a tuple with the MaxNbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNbRestart

`func (o *JobRequest) SetMaxNbRestart(v int32)`

SetMaxNbRestart sets MaxNbRestart field to given value.

### HasMaxNbRestart

`func (o *JobRequest) HasMaxNbRestart() bool`

HasMaxNbRestart returns a boolean if a field has been set.

### GetMaxDurationSeconds

`func (o *JobRequest) GetMaxDurationSeconds() int32`

GetMaxDurationSeconds returns the MaxDurationSeconds field if non-nil, zero value otherwise.

### GetMaxDurationSecondsOk

`func (o *JobRequest) GetMaxDurationSecondsOk() (*int32, bool)`

GetMaxDurationSecondsOk returns a tuple with the MaxDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationSeconds

`func (o *JobRequest) SetMaxDurationSeconds(v int32)`

SetMaxDurationSeconds sets MaxDurationSeconds field to given value.

### HasMaxDurationSeconds

`func (o *JobRequest) HasMaxDurationSeconds() bool`

HasMaxDurationSeconds returns a boolean if a field has been set.

### GetAutoPreview

`func (o *JobRequest) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *JobRequest) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *JobRequest) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *JobRequest) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetPort

`func (o *JobRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *JobRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *JobRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *JobRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *JobRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *JobRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSource

`func (o *JobRequest) GetSource() JobRequestAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JobRequest) GetSourceOk() (*JobRequestAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JobRequest) SetSource(v JobRequestAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *JobRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetHealthchecks

`func (o *JobRequest) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *JobRequest) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *JobRequest) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.

### HasHealthchecks

`func (o *JobRequest) HasHealthchecks() bool`

HasHealthchecks returns a boolean if a field has been set.

### GetSchedule

`func (o *JobRequest) GetSchedule() JobRequestAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobRequest) GetScheduleOk() (*JobRequestAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobRequest) SetSchedule(v JobRequestAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *JobRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *JobRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *JobRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *JobRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *JobRequest) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### SetAutoDeployNil

`func (o *JobRequest) SetAutoDeployNil(b bool)`

 SetAutoDeployNil sets the value for AutoDeploy to be an explicit nil

### UnsetAutoDeploy
`func (o *JobRequest) UnsetAutoDeploy()`

UnsetAutoDeploy ensures that no value is present for AutoDeploy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


