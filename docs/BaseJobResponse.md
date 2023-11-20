# BaseJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
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
**Source** | [**BaseJobResponseAllOfSource**](BaseJobResponseAllOfSource.md) |  | 
**Healthchecks** | [**Healthcheck**](Healthcheck.md) |  | 
**AutoDeploy** | Pointer to **bool** | Specify if the job will be automatically updated after receiving a new image tag or a new commit according to the source type.  The new image tag shall be communicated via the \&quot;Auto Deploy job\&quot; endpoint https://api-doc.qovery.com/#tag/Jobs/operation/autoDeployJobEnvironments  | [optional] 

## Methods

### NewBaseJobResponse

`func NewBaseJobResponse(id string, createdAt time.Time, environment ReferenceObject, maximumCpu int32, maximumMemory int32, name string, cpu int32, memory int32, autoPreview bool, source BaseJobResponseAllOfSource, healthchecks Healthcheck, ) *BaseJobResponse`

NewBaseJobResponse instantiates a new BaseJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseJobResponseWithDefaults

`func NewBaseJobResponseWithDefaults() *BaseJobResponse`

NewBaseJobResponseWithDefaults instantiates a new BaseJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseJobResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BaseJobResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseJobResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseJobResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseJobResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseJobResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseJobResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BaseJobResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *BaseJobResponse) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BaseJobResponse) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BaseJobResponse) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetMaximumCpu

`func (o *BaseJobResponse) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *BaseJobResponse) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *BaseJobResponse) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.


### GetMaximumMemory

`func (o *BaseJobResponse) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *BaseJobResponse) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *BaseJobResponse) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.


### GetName

`func (o *BaseJobResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseJobResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseJobResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BaseJobResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseJobResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseJobResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseJobResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCpu

`func (o *BaseJobResponse) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BaseJobResponse) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BaseJobResponse) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *BaseJobResponse) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *BaseJobResponse) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *BaseJobResponse) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetMaxNbRestart

`func (o *BaseJobResponse) GetMaxNbRestart() int32`

GetMaxNbRestart returns the MaxNbRestart field if non-nil, zero value otherwise.

### GetMaxNbRestartOk

`func (o *BaseJobResponse) GetMaxNbRestartOk() (*int32, bool)`

GetMaxNbRestartOk returns a tuple with the MaxNbRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNbRestart

`func (o *BaseJobResponse) SetMaxNbRestart(v int32)`

SetMaxNbRestart sets MaxNbRestart field to given value.

### HasMaxNbRestart

`func (o *BaseJobResponse) HasMaxNbRestart() bool`

HasMaxNbRestart returns a boolean if a field has been set.

### GetMaxDurationSeconds

`func (o *BaseJobResponse) GetMaxDurationSeconds() int32`

GetMaxDurationSeconds returns the MaxDurationSeconds field if non-nil, zero value otherwise.

### GetMaxDurationSecondsOk

`func (o *BaseJobResponse) GetMaxDurationSecondsOk() (*int32, bool)`

GetMaxDurationSecondsOk returns a tuple with the MaxDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationSeconds

`func (o *BaseJobResponse) SetMaxDurationSeconds(v int32)`

SetMaxDurationSeconds sets MaxDurationSeconds field to given value.

### HasMaxDurationSeconds

`func (o *BaseJobResponse) HasMaxDurationSeconds() bool`

HasMaxDurationSeconds returns a boolean if a field has been set.

### GetAutoPreview

`func (o *BaseJobResponse) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *BaseJobResponse) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *BaseJobResponse) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.


### GetPort

`func (o *BaseJobResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BaseJobResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BaseJobResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BaseJobResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *BaseJobResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *BaseJobResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSource

`func (o *BaseJobResponse) GetSource() BaseJobResponseAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BaseJobResponse) GetSourceOk() (*BaseJobResponseAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BaseJobResponse) SetSource(v BaseJobResponseAllOfSource)`

SetSource sets Source field to given value.


### GetHealthchecks

`func (o *BaseJobResponse) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *BaseJobResponse) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *BaseJobResponse) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.


### GetAutoDeploy

`func (o *BaseJobResponse) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *BaseJobResponse) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *BaseJobResponse) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *BaseJobResponse) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


