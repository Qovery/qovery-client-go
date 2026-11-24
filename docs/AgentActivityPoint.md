# AgentActivityPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** | Start of the hour or day. Buckets are cut in UTC; empty buckets are omitted. | 
**SkillLoads** | **int64** |  | 
**Actions** | **int64** |  | 
**Denials** | **int64** |  | 

## Methods

### NewAgentActivityPoint

`func NewAgentActivityPoint(start time.Time, skillLoads int64, actions int64, denials int64, ) *AgentActivityPoint`

NewAgentActivityPoint instantiates a new AgentActivityPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityPointWithDefaults

`func NewAgentActivityPointWithDefaults() *AgentActivityPoint`

NewAgentActivityPointWithDefaults instantiates a new AgentActivityPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *AgentActivityPoint) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AgentActivityPoint) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AgentActivityPoint) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetSkillLoads

`func (o *AgentActivityPoint) GetSkillLoads() int64`

GetSkillLoads returns the SkillLoads field if non-nil, zero value otherwise.

### GetSkillLoadsOk

`func (o *AgentActivityPoint) GetSkillLoadsOk() (*int64, bool)`

GetSkillLoadsOk returns a tuple with the SkillLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLoads

`func (o *AgentActivityPoint) SetSkillLoads(v int64)`

SetSkillLoads sets SkillLoads field to given value.


### GetActions

`func (o *AgentActivityPoint) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AgentActivityPoint) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AgentActivityPoint) SetActions(v int64)`

SetActions sets Actions field to given value.


### GetDenials

`func (o *AgentActivityPoint) GetDenials() int64`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *AgentActivityPoint) GetDenialsOk() (*int64, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *AgentActivityPoint) SetDenials(v int64)`

SetDenials sets Denials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


