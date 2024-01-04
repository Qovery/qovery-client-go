# JobRequestAllOfScheduleCronjob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**Timezone** | Pointer to **string** | Specify a timezone identifier to run the schedule at. By default Etc/UTC | [optional] 
**ScheduledAt** | **string** | Can only be set if the event is CRON.   Represent the cron format for the job schedule without seconds.   For example: &#x60;* * * * *&#x60; represent the cron to launch the job every minute.   See https://crontab.guru/ to WISIWIG interface.   Timezone is UTC  | 

## Methods

### NewJobRequestAllOfScheduleCronjob

`func NewJobRequestAllOfScheduleCronjob(scheduledAt string, ) *JobRequestAllOfScheduleCronjob`

NewJobRequestAllOfScheduleCronjob instantiates a new JobRequestAllOfScheduleCronjob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestAllOfScheduleCronjobWithDefaults

`func NewJobRequestAllOfScheduleCronjobWithDefaults() *JobRequestAllOfScheduleCronjob`

NewJobRequestAllOfScheduleCronjobWithDefaults instantiates a new JobRequestAllOfScheduleCronjob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *JobRequestAllOfScheduleCronjob) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JobRequestAllOfScheduleCronjob) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JobRequestAllOfScheduleCronjob) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *JobRequestAllOfScheduleCronjob) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *JobRequestAllOfScheduleCronjob) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *JobRequestAllOfScheduleCronjob) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *JobRequestAllOfScheduleCronjob) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *JobRequestAllOfScheduleCronjob) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetTimezone

`func (o *JobRequestAllOfScheduleCronjob) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *JobRequestAllOfScheduleCronjob) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *JobRequestAllOfScheduleCronjob) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *JobRequestAllOfScheduleCronjob) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetScheduledAt

`func (o *JobRequestAllOfScheduleCronjob) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *JobRequestAllOfScheduleCronjob) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *JobRequestAllOfScheduleCronjob) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


