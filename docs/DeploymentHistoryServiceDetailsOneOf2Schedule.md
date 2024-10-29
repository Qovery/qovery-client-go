# DeploymentHistoryServiceDetailsOneOf2Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnStart** | Pointer to [**JobLifecyleSchedule**](JobLifecyleSchedule.md) |  | [optional] 
**OnStop** | Pointer to [**JobLifecyleSchedule**](JobLifecyleSchedule.md) |  | [optional] 
**OnDelete** | Pointer to [**JobLifecyleSchedule**](JobLifecyleSchedule.md) |  | [optional] 
**CronJob** | Pointer to [**JobCronSchedule**](JobCronSchedule.md) |  | [optional] 
**LifecycleType** | Pointer to [**JobLifecycleTypeEnum**](JobLifecycleTypeEnum.md) |  | [optional] 

## Methods

### NewDeploymentHistoryServiceDetailsOneOf2Schedule

`func NewDeploymentHistoryServiceDetailsOneOf2Schedule() *DeploymentHistoryServiceDetailsOneOf2Schedule`

NewDeploymentHistoryServiceDetailsOneOf2Schedule instantiates a new DeploymentHistoryServiceDetailsOneOf2Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServiceDetailsOneOf2ScheduleWithDefaults

`func NewDeploymentHistoryServiceDetailsOneOf2ScheduleWithDefaults() *DeploymentHistoryServiceDetailsOneOf2Schedule`

NewDeploymentHistoryServiceDetailsOneOf2ScheduleWithDefaults instantiates a new DeploymentHistoryServiceDetailsOneOf2Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnStart

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetOnStart() JobLifecyleSchedule`

GetOnStart returns the OnStart field if non-nil, zero value otherwise.

### GetOnStartOk

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetOnStartOk() (*JobLifecyleSchedule, bool)`

GetOnStartOk returns a tuple with the OnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStart

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) SetOnStart(v JobLifecyleSchedule)`

SetOnStart sets OnStart field to given value.

### HasOnStart

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) HasOnStart() bool`

HasOnStart returns a boolean if a field has been set.

### GetOnStop

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetOnStop() JobLifecyleSchedule`

GetOnStop returns the OnStop field if non-nil, zero value otherwise.

### GetOnStopOk

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetOnStopOk() (*JobLifecyleSchedule, bool)`

GetOnStopOk returns a tuple with the OnStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStop

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) SetOnStop(v JobLifecyleSchedule)`

SetOnStop sets OnStop field to given value.

### HasOnStop

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) HasOnStop() bool`

HasOnStop returns a boolean if a field has been set.

### GetOnDelete

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetOnDelete() JobLifecyleSchedule`

GetOnDelete returns the OnDelete field if non-nil, zero value otherwise.

### GetOnDeleteOk

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetOnDeleteOk() (*JobLifecyleSchedule, bool)`

GetOnDeleteOk returns a tuple with the OnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelete

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) SetOnDelete(v JobLifecyleSchedule)`

SetOnDelete sets OnDelete field to given value.

### HasOnDelete

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) HasOnDelete() bool`

HasOnDelete returns a boolean if a field has been set.

### GetCronJob

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetCronJob() JobCronSchedule`

GetCronJob returns the CronJob field if non-nil, zero value otherwise.

### GetCronJobOk

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetCronJobOk() (*JobCronSchedule, bool)`

GetCronJobOk returns a tuple with the CronJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronJob

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) SetCronJob(v JobCronSchedule)`

SetCronJob sets CronJob field to given value.

### HasCronJob

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) HasCronJob() bool`

HasCronJob returns a boolean if a field has been set.

### GetLifecycleType

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetLifecycleType() JobLifecycleTypeEnum`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) GetLifecycleTypeOk() (*JobLifecycleTypeEnum, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) SetLifecycleType(v JobLifecycleTypeEnum)`

SetLifecycleType sets LifecycleType field to given value.

### HasLifecycleType

`func (o *DeploymentHistoryServiceDetailsOneOf2Schedule) HasLifecycleType() bool`

HasLifecycleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


