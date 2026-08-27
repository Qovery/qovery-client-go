# AgenticWorkflowResponseAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** |  | 
**Timezone** | **string** | tz database identifier the expression is evaluated in. | 
**NextRunAt** | **NullableTime** | When the schedule fires next. Null while the workflow is disabled, since a disabled workflow accumulates no occurrences. | [readonly] 

## Methods

### NewAgenticWorkflowResponseAllOfSchedule

`func NewAgenticWorkflowResponseAllOfSchedule(cronExpression string, timezone string, nextRunAt NullableTime, ) *AgenticWorkflowResponseAllOfSchedule`

NewAgenticWorkflowResponseAllOfSchedule instantiates a new AgenticWorkflowResponseAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowResponseAllOfScheduleWithDefaults

`func NewAgenticWorkflowResponseAllOfScheduleWithDefaults() *AgenticWorkflowResponseAllOfSchedule`

NewAgenticWorkflowResponseAllOfScheduleWithDefaults instantiates a new AgenticWorkflowResponseAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *AgenticWorkflowResponseAllOfSchedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *AgenticWorkflowResponseAllOfSchedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *AgenticWorkflowResponseAllOfSchedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetTimezone

`func (o *AgenticWorkflowResponseAllOfSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AgenticWorkflowResponseAllOfSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AgenticWorkflowResponseAllOfSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetNextRunAt

`func (o *AgenticWorkflowResponseAllOfSchedule) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *AgenticWorkflowResponseAllOfSchedule) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *AgenticWorkflowResponseAllOfSchedule) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.


### SetNextRunAtNil

`func (o *AgenticWorkflowResponseAllOfSchedule) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *AgenticWorkflowResponseAllOfSchedule) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


