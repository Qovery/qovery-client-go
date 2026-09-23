# AgenticWorkflowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Run ID. | 
**SourceWorkflowId** | **string** | ID of the workflow the run was requested for. A CLONE_ENVIRONMENT run executes as a fresh clone carrying its own ID, which run history does not report, so this is never the ID of the workflow that actually executed. | 
**Trigger** | [**AgenticWorkflowRunTrigger**](AgenticWorkflowRunTrigger.md) |  | 
**Prompt** | **NullableString** | Agent prompt captured when the run was requested. It is a snapshot, so later edits to the workflow do not change it. Null when the workflow had no prompt. | 
**CreatedAt** | **time.Time** | Time the run was requested. | 
**RecordedAt** | **NullableTime** | Time the run was registered in run history, shortly after it was requested. This is not a lifecycle start time: nothing reports when the agent itself started, so this value must not be used to measure a run. Null when it is unknown. | 

## Methods

### NewAgenticWorkflowRun

`func NewAgenticWorkflowRun(id string, sourceWorkflowId string, trigger AgenticWorkflowRunTrigger, prompt NullableString, createdAt time.Time, recordedAt NullableTime, ) *AgenticWorkflowRun`

NewAgenticWorkflowRun instantiates a new AgenticWorkflowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowRunWithDefaults

`func NewAgenticWorkflowRunWithDefaults() *AgenticWorkflowRun`

NewAgenticWorkflowRunWithDefaults instantiates a new AgenticWorkflowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgenticWorkflowRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgenticWorkflowRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgenticWorkflowRun) SetId(v string)`

SetId sets Id field to given value.


### GetSourceWorkflowId

`func (o *AgenticWorkflowRun) GetSourceWorkflowId() string`

GetSourceWorkflowId returns the SourceWorkflowId field if non-nil, zero value otherwise.

### GetSourceWorkflowIdOk

`func (o *AgenticWorkflowRun) GetSourceWorkflowIdOk() (*string, bool)`

GetSourceWorkflowIdOk returns a tuple with the SourceWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWorkflowId

`func (o *AgenticWorkflowRun) SetSourceWorkflowId(v string)`

SetSourceWorkflowId sets SourceWorkflowId field to given value.


### GetTrigger

`func (o *AgenticWorkflowRun) GetTrigger() AgenticWorkflowRunTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *AgenticWorkflowRun) GetTriggerOk() (*AgenticWorkflowRunTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *AgenticWorkflowRun) SetTrigger(v AgenticWorkflowRunTrigger)`

SetTrigger sets Trigger field to given value.


### GetPrompt

`func (o *AgenticWorkflowRun) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AgenticWorkflowRun) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AgenticWorkflowRun) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### SetPromptNil

`func (o *AgenticWorkflowRun) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *AgenticWorkflowRun) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetCreatedAt

`func (o *AgenticWorkflowRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgenticWorkflowRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgenticWorkflowRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRecordedAt

`func (o *AgenticWorkflowRun) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *AgenticWorkflowRun) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *AgenticWorkflowRun) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.


### SetRecordedAtNil

`func (o *AgenticWorkflowRun) SetRecordedAtNil(b bool)`

 SetRecordedAtNil sets the value for RecordedAt to be an explicit nil

### UnsetRecordedAt
`func (o *AgenticWorkflowRun) UnsetRecordedAt()`

UnsetRecordedAt ensures that no value is present for RecordedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


