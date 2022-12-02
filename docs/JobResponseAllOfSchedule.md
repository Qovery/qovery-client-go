# JobResponseAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnStart** | Pointer to [**JobRequestAllOfScheduleOnStart**](JobRequestAllOfScheduleOnStart.md) |  | [optional] 
**OnStop** | Pointer to [**JobRequestAllOfScheduleOnStart**](JobRequestAllOfScheduleOnStart.md) |  | [optional] 
**OnDelete** | Pointer to [**JobRequestAllOfScheduleOnStart**](JobRequestAllOfScheduleOnStart.md) |  | [optional] 
**Cronjob** | Pointer to [**JobResponseAllOfScheduleCronjob**](JobResponseAllOfScheduleCronjob.md) |  | [optional] 

## Methods

### NewJobResponseAllOfSchedule

`func NewJobResponseAllOfSchedule() *JobResponseAllOfSchedule`

NewJobResponseAllOfSchedule instantiates a new JobResponseAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResponseAllOfScheduleWithDefaults

`func NewJobResponseAllOfScheduleWithDefaults() *JobResponseAllOfSchedule`

NewJobResponseAllOfScheduleWithDefaults instantiates a new JobResponseAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnStart

`func (o *JobResponseAllOfSchedule) GetOnStart() JobRequestAllOfScheduleOnStart`

GetOnStart returns the OnStart field if non-nil, zero value otherwise.

### GetOnStartOk

`func (o *JobResponseAllOfSchedule) GetOnStartOk() (*JobRequestAllOfScheduleOnStart, bool)`

GetOnStartOk returns a tuple with the OnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStart

`func (o *JobResponseAllOfSchedule) SetOnStart(v JobRequestAllOfScheduleOnStart)`

SetOnStart sets OnStart field to given value.

### HasOnStart

`func (o *JobResponseAllOfSchedule) HasOnStart() bool`

HasOnStart returns a boolean if a field has been set.

### GetOnStop

`func (o *JobResponseAllOfSchedule) GetOnStop() JobRequestAllOfScheduleOnStart`

GetOnStop returns the OnStop field if non-nil, zero value otherwise.

### GetOnStopOk

`func (o *JobResponseAllOfSchedule) GetOnStopOk() (*JobRequestAllOfScheduleOnStart, bool)`

GetOnStopOk returns a tuple with the OnStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStop

`func (o *JobResponseAllOfSchedule) SetOnStop(v JobRequestAllOfScheduleOnStart)`

SetOnStop sets OnStop field to given value.

### HasOnStop

`func (o *JobResponseAllOfSchedule) HasOnStop() bool`

HasOnStop returns a boolean if a field has been set.

### GetOnDelete

`func (o *JobResponseAllOfSchedule) GetOnDelete() JobRequestAllOfScheduleOnStart`

GetOnDelete returns the OnDelete field if non-nil, zero value otherwise.

### GetOnDeleteOk

`func (o *JobResponseAllOfSchedule) GetOnDeleteOk() (*JobRequestAllOfScheduleOnStart, bool)`

GetOnDeleteOk returns a tuple with the OnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelete

`func (o *JobResponseAllOfSchedule) SetOnDelete(v JobRequestAllOfScheduleOnStart)`

SetOnDelete sets OnDelete field to given value.

### HasOnDelete

`func (o *JobResponseAllOfSchedule) HasOnDelete() bool`

HasOnDelete returns a boolean if a field has been set.

### GetCronjob

`func (o *JobResponseAllOfSchedule) GetCronjob() JobResponseAllOfScheduleCronjob`

GetCronjob returns the Cronjob field if non-nil, zero value otherwise.

### GetCronjobOk

`func (o *JobResponseAllOfSchedule) GetCronjobOk() (*JobResponseAllOfScheduleCronjob, bool)`

GetCronjobOk returns a tuple with the Cronjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronjob

`func (o *JobResponseAllOfSchedule) SetCronjob(v JobResponseAllOfScheduleCronjob)`

SetCronjob sets Cronjob field to given value.

### HasCronjob

`func (o *JobResponseAllOfSchedule) HasCronjob() bool`

HasCronjob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


