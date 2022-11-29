# JobRequestAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**JobScheduleEvent**](JobScheduleEvent.md) |  | [optional] 
**ScheduledAt** | Pointer to **NullableString** | Can only be set if the event is CRON. Represent the cron format for the job schedule without seconds. For example: &#x60;* * * * *&#x60; represent the cron to launch the job every minute. See https://crontab.guru/ to WISIWIG interface. Timezone is UTC  | [optional] 

## Methods

### NewJobRequestAllOfSchedule

`func NewJobRequestAllOfSchedule() *JobRequestAllOfSchedule`

NewJobRequestAllOfSchedule instantiates a new JobRequestAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestAllOfScheduleWithDefaults

`func NewJobRequestAllOfScheduleWithDefaults() *JobRequestAllOfSchedule`

NewJobRequestAllOfScheduleWithDefaults instantiates a new JobRequestAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *JobRequestAllOfSchedule) GetEvent() JobScheduleEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *JobRequestAllOfSchedule) GetEventOk() (*JobScheduleEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *JobRequestAllOfSchedule) SetEvent(v JobScheduleEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *JobRequestAllOfSchedule) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetScheduledAt

`func (o *JobRequestAllOfSchedule) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *JobRequestAllOfSchedule) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *JobRequestAllOfSchedule) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *JobRequestAllOfSchedule) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### SetScheduledAtNil

`func (o *JobRequestAllOfSchedule) SetScheduledAtNil(b bool)`

 SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil

### UnsetScheduledAt
`func (o *JobRequestAllOfSchedule) UnsetScheduledAt()`

UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


