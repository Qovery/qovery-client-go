# AgentActivityTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillLoads** | **int64** | Skills loaded by agents. | 
**Actions** | **int64** | Agent actions, allowed and denied. | 
**Denials** | **int64** | Agent actions that were refused. | 
**Members** | **int64** | Distinct actors with any activity, policy tokens included. | 
**PolicyTokens** | **int64** | Distinct policy tokens with any activity. | 

## Methods

### NewAgentActivityTotals

`func NewAgentActivityTotals(skillLoads int64, actions int64, denials int64, members int64, policyTokens int64, ) *AgentActivityTotals`

NewAgentActivityTotals instantiates a new AgentActivityTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityTotalsWithDefaults

`func NewAgentActivityTotalsWithDefaults() *AgentActivityTotals`

NewAgentActivityTotalsWithDefaults instantiates a new AgentActivityTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillLoads

`func (o *AgentActivityTotals) GetSkillLoads() int64`

GetSkillLoads returns the SkillLoads field if non-nil, zero value otherwise.

### GetSkillLoadsOk

`func (o *AgentActivityTotals) GetSkillLoadsOk() (*int64, bool)`

GetSkillLoadsOk returns a tuple with the SkillLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLoads

`func (o *AgentActivityTotals) SetSkillLoads(v int64)`

SetSkillLoads sets SkillLoads field to given value.


### GetActions

`func (o *AgentActivityTotals) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AgentActivityTotals) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AgentActivityTotals) SetActions(v int64)`

SetActions sets Actions field to given value.


### GetDenials

`func (o *AgentActivityTotals) GetDenials() int64`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *AgentActivityTotals) GetDenialsOk() (*int64, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *AgentActivityTotals) SetDenials(v int64)`

SetDenials sets Denials field to given value.


### GetMembers

`func (o *AgentActivityTotals) GetMembers() int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AgentActivityTotals) GetMembersOk() (*int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AgentActivityTotals) SetMembers(v int64)`

SetMembers sets Members field to given value.


### GetPolicyTokens

`func (o *AgentActivityTotals) GetPolicyTokens() int64`

GetPolicyTokens returns the PolicyTokens field if non-nil, zero value otherwise.

### GetPolicyTokensOk

`func (o *AgentActivityTotals) GetPolicyTokensOk() (*int64, bool)`

GetPolicyTokensOk returns a tuple with the PolicyTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTokens

`func (o *AgentActivityTotals) SetPolicyTokens(v int64)`

SetPolicyTokens sets PolicyTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


