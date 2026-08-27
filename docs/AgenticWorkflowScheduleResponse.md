# AgenticWorkflowScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** |  | 
**Timezone** | **string** | tz database identifier the expression is evaluated in. | 
**NextRunAt** | **NullableTime** | When the schedule fires next. Null while the workflow is disabled, since a disabled workflow accumulates no occurrences. | [readonly] 

## Methods

### NewAgenticWorkflowScheduleResponse

`func NewAgenticWorkflowScheduleResponse(cronExpression string, timezone string, nextRunAt NullableTime, ) *AgenticWorkflowScheduleResponse`

NewAgenticWorkflowScheduleResponse instantiates a new AgenticWorkflowScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowScheduleResponseWithDefaults

`func NewAgenticWorkflowScheduleResponseWithDefaults() *AgenticWorkflowScheduleResponse`

NewAgenticWorkflowScheduleResponseWithDefaults instantiates a new AgenticWorkflowScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *AgenticWorkflowScheduleResponse) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *AgenticWorkflowScheduleResponse) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *AgenticWorkflowScheduleResponse) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetTimezone

`func (o *AgenticWorkflowScheduleResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AgenticWorkflowScheduleResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AgenticWorkflowScheduleResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetNextRunAt

`func (o *AgenticWorkflowScheduleResponse) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *AgenticWorkflowScheduleResponse) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *AgenticWorkflowScheduleResponse) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.


### SetNextRunAtNil

`func (o *AgenticWorkflowScheduleResponse) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *AgenticWorkflowScheduleResponse) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


