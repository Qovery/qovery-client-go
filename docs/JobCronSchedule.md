# JobCronSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | **[]string** |  | 
**Entrypoint** | Pointer to **string** |  | [optional] 
**ScheduledAt** | **string** |  | 
**Timezone** | **string** |  | 

## Methods

### NewJobCronSchedule

`func NewJobCronSchedule(arguments []string, scheduledAt string, timezone string, ) *JobCronSchedule`

NewJobCronSchedule instantiates a new JobCronSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobCronScheduleWithDefaults

`func NewJobCronScheduleWithDefaults() *JobCronSchedule`

NewJobCronScheduleWithDefaults instantiates a new JobCronSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *JobCronSchedule) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JobCronSchedule) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JobCronSchedule) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetEntrypoint

`func (o *JobCronSchedule) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *JobCronSchedule) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *JobCronSchedule) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *JobCronSchedule) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetScheduledAt

`func (o *JobCronSchedule) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *JobCronSchedule) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *JobCronSchedule) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.


### GetTimezone

`func (o *JobCronSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *JobCronSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *JobCronSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


