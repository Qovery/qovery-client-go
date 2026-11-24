# AgentActivityEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorKind** | [**AgentActorKind**](AgentActorKind.md) |  | 
**PolicyTokenId** | Pointer to **NullableString** | The policy token&#39;s id when &#x60;actor_kind&#x60; is POLICY_TOKEN and the token is known. Null otherwise. | [optional] 
**At** | **time.Time** |  | 
**UserSub** | Pointer to **NullableString** | The member&#39;s id. Null when the actor loaded no skill in the range, such as a policy token. | [optional] 
**UserName** | **string** | The username the audit log records for the actor. | 
**Source** | [**AgentActivitySource**](AgentActivitySource.md) | Where the action comes from: SKILL_API_ACTION, MCP_TOOL_CALL, POLICY_DENIAL or POLICY_TOKEN_API_ACTION. Skill loads are counted, never listed here, so SKILL_LOAD does not appear. | 
**SkillName** | Pointer to **NullableString** | The skill named in the user agent. Null when the action names none. | [optional] 
**Event** | **string** | The audit event type, as in &#x60;OrganizationEventType&#x60;. | 
**TargetType** | **string** | The audit target type, as in &#x60;OrganizationEventTargetType&#x60;. | 
**TargetName** | **string** |  | 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**EnvironmentName** | Pointer to **NullableString** |  | [optional] 
**Outcome** | [**AgentActionOutcome**](AgentActionOutcome.md) |  | 

## Methods

### NewAgentActivityEntry

`func NewAgentActivityEntry(actorKind AgentActorKind, at time.Time, userName string, source AgentActivitySource, event string, targetType string, targetName string, outcome AgentActionOutcome, ) *AgentActivityEntry`

NewAgentActivityEntry instantiates a new AgentActivityEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityEntryWithDefaults

`func NewAgentActivityEntryWithDefaults() *AgentActivityEntry`

NewAgentActivityEntryWithDefaults instantiates a new AgentActivityEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorKind

`func (o *AgentActivityEntry) GetActorKind() AgentActorKind`

GetActorKind returns the ActorKind field if non-nil, zero value otherwise.

### GetActorKindOk

`func (o *AgentActivityEntry) GetActorKindOk() (*AgentActorKind, bool)`

GetActorKindOk returns a tuple with the ActorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorKind

`func (o *AgentActivityEntry) SetActorKind(v AgentActorKind)`

SetActorKind sets ActorKind field to given value.


### GetPolicyTokenId

`func (o *AgentActivityEntry) GetPolicyTokenId() string`

GetPolicyTokenId returns the PolicyTokenId field if non-nil, zero value otherwise.

### GetPolicyTokenIdOk

`func (o *AgentActivityEntry) GetPolicyTokenIdOk() (*string, bool)`

GetPolicyTokenIdOk returns a tuple with the PolicyTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTokenId

`func (o *AgentActivityEntry) SetPolicyTokenId(v string)`

SetPolicyTokenId sets PolicyTokenId field to given value.

### HasPolicyTokenId

`func (o *AgentActivityEntry) HasPolicyTokenId() bool`

HasPolicyTokenId returns a boolean if a field has been set.

### SetPolicyTokenIdNil

`func (o *AgentActivityEntry) SetPolicyTokenIdNil(b bool)`

 SetPolicyTokenIdNil sets the value for PolicyTokenId to be an explicit nil

### UnsetPolicyTokenId
`func (o *AgentActivityEntry) UnsetPolicyTokenId()`

UnsetPolicyTokenId ensures that no value is present for PolicyTokenId, not even an explicit nil
### GetAt

`func (o *AgentActivityEntry) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentActivityEntry) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentActivityEntry) SetAt(v time.Time)`

SetAt sets At field to given value.


### GetUserSub

`func (o *AgentActivityEntry) GetUserSub() string`

GetUserSub returns the UserSub field if non-nil, zero value otherwise.

### GetUserSubOk

`func (o *AgentActivityEntry) GetUserSubOk() (*string, bool)`

GetUserSubOk returns a tuple with the UserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSub

`func (o *AgentActivityEntry) SetUserSub(v string)`

SetUserSub sets UserSub field to given value.

### HasUserSub

`func (o *AgentActivityEntry) HasUserSub() bool`

HasUserSub returns a boolean if a field has been set.

### SetUserSubNil

`func (o *AgentActivityEntry) SetUserSubNil(b bool)`

 SetUserSubNil sets the value for UserSub to be an explicit nil

### UnsetUserSub
`func (o *AgentActivityEntry) UnsetUserSub()`

UnsetUserSub ensures that no value is present for UserSub, not even an explicit nil
### GetUserName

`func (o *AgentActivityEntry) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AgentActivityEntry) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AgentActivityEntry) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetSource

`func (o *AgentActivityEntry) GetSource() AgentActivitySource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AgentActivityEntry) GetSourceOk() (*AgentActivitySource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AgentActivityEntry) SetSource(v AgentActivitySource)`

SetSource sets Source field to given value.


### GetSkillName

`func (o *AgentActivityEntry) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *AgentActivityEntry) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *AgentActivityEntry) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.

### HasSkillName

`func (o *AgentActivityEntry) HasSkillName() bool`

HasSkillName returns a boolean if a field has been set.

### SetSkillNameNil

`func (o *AgentActivityEntry) SetSkillNameNil(b bool)`

 SetSkillNameNil sets the value for SkillName to be an explicit nil

### UnsetSkillName
`func (o *AgentActivityEntry) UnsetSkillName()`

UnsetSkillName ensures that no value is present for SkillName, not even an explicit nil
### GetEvent

`func (o *AgentActivityEntry) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AgentActivityEntry) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AgentActivityEntry) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTargetType

`func (o *AgentActivityEntry) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AgentActivityEntry) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AgentActivityEntry) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTargetName

`func (o *AgentActivityEntry) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AgentActivityEntry) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AgentActivityEntry) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetProjectName

`func (o *AgentActivityEntry) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *AgentActivityEntry) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *AgentActivityEntry) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *AgentActivityEntry) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *AgentActivityEntry) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *AgentActivityEntry) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetEnvironmentName

`func (o *AgentActivityEntry) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *AgentActivityEntry) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *AgentActivityEntry) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *AgentActivityEntry) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### SetEnvironmentNameNil

`func (o *AgentActivityEntry) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *AgentActivityEntry) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil
### GetOutcome

`func (o *AgentActivityEntry) GetOutcome() AgentActionOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AgentActivityEntry) GetOutcomeOk() (*AgentActionOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AgentActivityEntry) SetOutcome(v AgentActionOutcome)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


