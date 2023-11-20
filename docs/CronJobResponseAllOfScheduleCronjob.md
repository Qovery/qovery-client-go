# CronJobResponseAllOfScheduleCronjob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**ScheduledAt** | **string** | Can only be set if the event is CRON.   Represent the cron format for the job schedule without seconds.   For example: &#x60;* * * * *&#x60; represent the cron to launch the job every minute.   See https://crontab.guru/ to WISIWIG interface.   Timezone is UT  | 

## Methods

### NewCronJobResponseAllOfScheduleCronjob

`func NewCronJobResponseAllOfScheduleCronjob(scheduledAt string, ) *CronJobResponseAllOfScheduleCronjob`

NewCronJobResponseAllOfScheduleCronjob instantiates a new CronJobResponseAllOfScheduleCronjob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobResponseAllOfScheduleCronjobWithDefaults

`func NewCronJobResponseAllOfScheduleCronjobWithDefaults() *CronJobResponseAllOfScheduleCronjob`

NewCronJobResponseAllOfScheduleCronjobWithDefaults instantiates a new CronJobResponseAllOfScheduleCronjob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *CronJobResponseAllOfScheduleCronjob) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *CronJobResponseAllOfScheduleCronjob) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *CronJobResponseAllOfScheduleCronjob) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *CronJobResponseAllOfScheduleCronjob) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *CronJobResponseAllOfScheduleCronjob) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *CronJobResponseAllOfScheduleCronjob) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *CronJobResponseAllOfScheduleCronjob) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *CronJobResponseAllOfScheduleCronjob) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetScheduledAt

`func (o *CronJobResponseAllOfScheduleCronjob) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *CronJobResponseAllOfScheduleCronjob) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *CronJobResponseAllOfScheduleCronjob) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


