# AgentActivityBySkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillName** | Pointer to **NullableString** | Null groups the agent actions that name no skill: MCP tool calls and API calls made without a skill user agent, such as by a policy token. | [optional] 
**Loads** | **int64** |  | 
**Actions** | **int64** |  | 
**Denials** | **int64** |  | 
**Members** | **int64** | Distinct members with activity on this skill. | 

## Methods

### NewAgentActivityBySkill

`func NewAgentActivityBySkill(loads int64, actions int64, denials int64, members int64, ) *AgentActivityBySkill`

NewAgentActivityBySkill instantiates a new AgentActivityBySkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityBySkillWithDefaults

`func NewAgentActivityBySkillWithDefaults() *AgentActivityBySkill`

NewAgentActivityBySkillWithDefaults instantiates a new AgentActivityBySkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillName

`func (o *AgentActivityBySkill) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *AgentActivityBySkill) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *AgentActivityBySkill) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.

### HasSkillName

`func (o *AgentActivityBySkill) HasSkillName() bool`

HasSkillName returns a boolean if a field has been set.

### SetSkillNameNil

`func (o *AgentActivityBySkill) SetSkillNameNil(b bool)`

 SetSkillNameNil sets the value for SkillName to be an explicit nil

### UnsetSkillName
`func (o *AgentActivityBySkill) UnsetSkillName()`

UnsetSkillName ensures that no value is present for SkillName, not even an explicit nil
### GetLoads

`func (o *AgentActivityBySkill) GetLoads() int64`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *AgentActivityBySkill) GetLoadsOk() (*int64, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *AgentActivityBySkill) SetLoads(v int64)`

SetLoads sets Loads field to given value.


### GetActions

`func (o *AgentActivityBySkill) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AgentActivityBySkill) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AgentActivityBySkill) SetActions(v int64)`

SetActions sets Actions field to given value.


### GetDenials

`func (o *AgentActivityBySkill) GetDenials() int64`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *AgentActivityBySkill) GetDenialsOk() (*int64, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *AgentActivityBySkill) SetDenials(v int64)`

SetDenials sets Denials field to given value.


### GetMembers

`func (o *AgentActivityBySkill) GetMembers() int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AgentActivityBySkill) GetMembersOk() (*int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AgentActivityBySkill) SetMembers(v int64)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


