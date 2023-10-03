# JobResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Registry** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**MaximumCpu** | **int32** | Maximum cpu that can be allocated to the job based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | 
**MaximumMemory** | **int32** | Maximum memory that can be allocated to the job based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Cpu** | **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | 
**Memory** | **int32** | unit is MB. 1024 MB &#x3D; 1GB | 
**MaxNbRestart** | Pointer to **int32** | Maximum number of restart allowed before the job is considered as failed 0 means that no restart/crash of the job is allowed  | [optional] 
**MaxDurationSeconds** | Pointer to **int32** | Maximum number of seconds allowed for the job to run before killing it and mark it as failed  | [optional] 
**AutoPreview** | **bool** | Indicates if the &#39;environment preview option&#39; is enabled for this container.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | 
**Port** | Pointer to **NullableInt32** | Port where to run readiness and liveliness probes checks. The port will not be exposed externally | [optional] 
**Source** | Pointer to [**JobResponseAllOfSource**](JobResponseAllOfSource.md) |  | [optional] 
**Healthchecks** | [**Healthcheck**](Healthcheck.md) |  | 
**Schedule** | Pointer to [**JobResponseAllOfSchedule**](JobResponseAllOfSchedule.md) |  | [optional] 
**AutoDeploy** | Pointer to **bool** | Specify if the job will be automatically updated after receiving a new image tag or a new commit according to the source type.  The new image tag shall be communicated via the \&quot;Auto Deploy job\&quot; endpoint https://api-doc.qovery.com/#tag/Jobs/operation/autoDeployJobEnvironments  | [optional] 

## Methods

### NewJobResponseAllOf

`func NewJobResponseAllOf(environment ReferenceObject, maximumCpu int32, maximumMemory int32, name string, cpu int32, memory int32, autoPreview bool, healthchecks Healthcheck, ) *JobResponseAllOf`

NewJobResponseAllOf instantiates a new JobResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResponseAllOfWithDefaults

`func NewJobResponseAllOfWithDefaults() *JobResponseAllOf`

NewJobResponseAllOfWithDefaults instantiates a new JobResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *JobResponseAllOf) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *JobResponseAllOf) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *JobResponseAllOf) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetRegistry

`func (o *JobResponseAllOf) GetRegistry() ReferenceObject`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *JobResponseAllOf) GetRegistryOk() (*ReferenceObject, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *JobResponseAllOf) SetRegistry(v ReferenceObject)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *JobResponseAllOf) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetMaximumCpu

`func (o *JobResponseAllOf) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *JobResponseAllOf) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *JobResponseAllOf) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.


### GetMaximumMemory

`func (o *JobResponseAllOf) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *JobResponseAllOf) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *JobResponseAllOf) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.


### GetName

`func (o *JobResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResponseAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCpu

`func (o *JobResponseAllOf) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *JobResponseAllOf) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *JobResponseAllOf) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *JobResponseAllOf) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JobResponseAllOf) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JobResponseAllOf) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetMaxNbRestart

`func (o *JobResponseAllOf) GetMaxNbRestart() int32`

GetMaxNbRestart returns the MaxNbRestart field if non-nil, zero value otherwise.

### GetMaxNbRestartOk

`func (o *JobResponseAllOf) GetMaxNbRestartOk() (*int32, bool)`

GetMaxNbRestartOk returns a tuple with the MaxNbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNbRestart

`func (o *JobResponseAllOf) SetMaxNbRestart(v int32)`

SetMaxNbRestart sets MaxNbRestart field to given value.

### HasMaxNbRestart

`func (o *JobResponseAllOf) HasMaxNbRestart() bool`

HasMaxNbRestart returns a boolean if a field has been set.

### GetMaxDurationSeconds

`func (o *JobResponseAllOf) GetMaxDurationSeconds() int32`

GetMaxDurationSeconds returns the MaxDurationSeconds field if non-nil, zero value otherwise.

### GetMaxDurationSecondsOk

`func (o *JobResponseAllOf) GetMaxDurationSecondsOk() (*int32, bool)`

GetMaxDurationSecondsOk returns a tuple with the MaxDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationSeconds

`func (o *JobResponseAllOf) SetMaxDurationSeconds(v int32)`

SetMaxDurationSeconds sets MaxDurationSeconds field to given value.

### HasMaxDurationSeconds

`func (o *JobResponseAllOf) HasMaxDurationSeconds() bool`

HasMaxDurationSeconds returns a boolean if a field has been set.

### GetAutoPreview

`func (o *JobResponseAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *JobResponseAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *JobResponseAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.


### GetPort

`func (o *JobResponseAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *JobResponseAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *JobResponseAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *JobResponseAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *JobResponseAllOf) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *JobResponseAllOf) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSource

`func (o *JobResponseAllOf) GetSource() JobResponseAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JobResponseAllOf) GetSourceOk() (*JobResponseAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JobResponseAllOf) SetSource(v JobResponseAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *JobResponseAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetHealthchecks

`func (o *JobResponseAllOf) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *JobResponseAllOf) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *JobResponseAllOf) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.


### GetSchedule

`func (o *JobResponseAllOf) GetSchedule() JobResponseAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobResponseAllOf) GetScheduleOk() (*JobResponseAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobResponseAllOf) SetSchedule(v JobResponseAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *JobResponseAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *JobResponseAllOf) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *JobResponseAllOf) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *JobResponseAllOf) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *JobResponseAllOf) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


