# DeploymentHistoryJobResponseAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**JobScheduleEvent**](JobScheduleEvent.md) |  | [optional] 
**ScheduleAt** | Pointer to **string** | Can only be set if the event is CRON. Represent the cron format for the job schedule without seconds. For example: &#x60;* * * * *&#x60; represent the cron to launch the job every minute. See https://crontab.guru/ to WISIWIG interface. Timezone is UTC  | [optional] 

## Methods

### NewDeploymentHistoryJobResponseAllOfSchedule

`func NewDeploymentHistoryJobResponseAllOfSchedule() *DeploymentHistoryJobResponseAllOfSchedule`

NewDeploymentHistoryJobResponseAllOfSchedule instantiates a new DeploymentHistoryJobResponseAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryJobResponseAllOfScheduleWithDefaults

`func NewDeploymentHistoryJobResponseAllOfScheduleWithDefaults() *DeploymentHistoryJobResponseAllOfSchedule`

NewDeploymentHistoryJobResponseAllOfScheduleWithDefaults instantiates a new DeploymentHistoryJobResponseAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *DeploymentHistoryJobResponseAllOfSchedule) GetEvent() JobScheduleEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DeploymentHistoryJobResponseAllOfSchedule) GetEventOk() (*JobScheduleEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DeploymentHistoryJobResponseAllOfSchedule) SetEvent(v JobScheduleEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *DeploymentHistoryJobResponseAllOfSchedule) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetScheduleAt

`func (o *DeploymentHistoryJobResponseAllOfSchedule) GetScheduleAt() string`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *DeploymentHistoryJobResponseAllOfSchedule) GetScheduleAtOk() (*string, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *DeploymentHistoryJobResponseAllOfSchedule) SetScheduleAt(v string)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *DeploymentHistoryJobResponseAllOfSchedule) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


