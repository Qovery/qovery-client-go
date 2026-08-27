# AgenticWorkflowRequestAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** | Five-field cron expression, the same syntax Kubernetes cron jobs accept. Rejected if it fires more than once every 5 minutes, since each run deploys a full environment. | 
**Timezone** | **string** | tz database identifier the expression is evaluated in. | 

## Methods

### NewAgenticWorkflowRequestAllOfSchedule

`func NewAgenticWorkflowRequestAllOfSchedule(cronExpression string, timezone string, ) *AgenticWorkflowRequestAllOfSchedule`

NewAgenticWorkflowRequestAllOfSchedule instantiates a new AgenticWorkflowRequestAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowRequestAllOfScheduleWithDefaults

`func NewAgenticWorkflowRequestAllOfScheduleWithDefaults() *AgenticWorkflowRequestAllOfSchedule`

NewAgenticWorkflowRequestAllOfScheduleWithDefaults instantiates a new AgenticWorkflowRequestAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *AgenticWorkflowRequestAllOfSchedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *AgenticWorkflowRequestAllOfSchedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *AgenticWorkflowRequestAllOfSchedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetTimezone

`func (o *AgenticWorkflowRequestAllOfSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AgenticWorkflowRequestAllOfSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AgenticWorkflowRequestAllOfSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


