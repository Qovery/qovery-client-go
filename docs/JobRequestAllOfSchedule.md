# JobRequestAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnStart** | Pointer to [**JobRequestAllOfScheduleOnStart**](JobRequestAllOfScheduleOnStart.md) |  | [optional] 
**OnStop** | Pointer to [**JobRequestAllOfScheduleOnStart**](JobRequestAllOfScheduleOnStart.md) |  | [optional] 
**OnDelete** | Pointer to [**JobRequestAllOfScheduleOnStart**](JobRequestAllOfScheduleOnStart.md) |  | [optional] 
**Cronjob** | Pointer to [**JobRequestAllOfScheduleCronjob**](JobRequestAllOfScheduleCronjob.md) |  | [optional] 
**LifecycleType** | Pointer to [**JobLifecycleTypeEnum**](JobLifecycleTypeEnum.md) |  | [optional] 

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

### GetOnStart

`func (o *JobRequestAllOfSchedule) GetOnStart() JobRequestAllOfScheduleOnStart`

GetOnStart returns the OnStart field if non-nil, zero value otherwise.

### GetOnStartOk

`func (o *JobRequestAllOfSchedule) GetOnStartOk() (*JobRequestAllOfScheduleOnStart, bool)`

GetOnStartOk returns a tuple with the OnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStart

`func (o *JobRequestAllOfSchedule) SetOnStart(v JobRequestAllOfScheduleOnStart)`

SetOnStart sets OnStart field to given value.

### HasOnStart

`func (o *JobRequestAllOfSchedule) HasOnStart() bool`

HasOnStart returns a boolean if a field has been set.

### GetOnStop

`func (o *JobRequestAllOfSchedule) GetOnStop() JobRequestAllOfScheduleOnStart`

GetOnStop returns the OnStop field if non-nil, zero value otherwise.

### GetOnStopOk

`func (o *JobRequestAllOfSchedule) GetOnStopOk() (*JobRequestAllOfScheduleOnStart, bool)`

GetOnStopOk returns a tuple with the OnStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStop

`func (o *JobRequestAllOfSchedule) SetOnStop(v JobRequestAllOfScheduleOnStart)`

SetOnStop sets OnStop field to given value.

### HasOnStop

`func (o *JobRequestAllOfSchedule) HasOnStop() bool`

HasOnStop returns a boolean if a field has been set.

### GetOnDelete

`func (o *JobRequestAllOfSchedule) GetOnDelete() JobRequestAllOfScheduleOnStart`

GetOnDelete returns the OnDelete field if non-nil, zero value otherwise.

### GetOnDeleteOk

`func (o *JobRequestAllOfSchedule) GetOnDeleteOk() (*JobRequestAllOfScheduleOnStart, bool)`

GetOnDeleteOk returns a tuple with the OnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelete

`func (o *JobRequestAllOfSchedule) SetOnDelete(v JobRequestAllOfScheduleOnStart)`

SetOnDelete sets OnDelete field to given value.

### HasOnDelete

`func (o *JobRequestAllOfSchedule) HasOnDelete() bool`

HasOnDelete returns a boolean if a field has been set.

### GetCronjob

`func (o *JobRequestAllOfSchedule) GetCronjob() JobRequestAllOfScheduleCronjob`

GetCronjob returns the Cronjob field if non-nil, zero value otherwise.

### GetCronjobOk

`func (o *JobRequestAllOfSchedule) GetCronjobOk() (*JobRequestAllOfScheduleCronjob, bool)`

GetCronjobOk returns a tuple with the Cronjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronjob

`func (o *JobRequestAllOfSchedule) SetCronjob(v JobRequestAllOfScheduleCronjob)`

SetCronjob sets Cronjob field to given value.

### HasCronjob

`func (o *JobRequestAllOfSchedule) HasCronjob() bool`

HasCronjob returns a boolean if a field has been set.

### GetLifecycleType

`func (o *JobRequestAllOfSchedule) GetLifecycleType() JobLifecycleTypeEnum`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *JobRequestAllOfSchedule) GetLifecycleTypeOk() (*JobLifecycleTypeEnum, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *JobRequestAllOfSchedule) SetLifecycleType(v JobLifecycleTypeEnum)`

SetLifecycleType sets LifecycleType field to given value.

### HasLifecycleType

`func (o *JobRequestAllOfSchedule) HasLifecycleType() bool`

HasLifecycleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


