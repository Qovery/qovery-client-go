# AgentActivityByMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorKind** | [**AgentActorKind**](AgentActorKind.md) |  | 
**PolicyTokenId** | Pointer to **NullableString** | The policy token&#39;s id when &#x60;actor_kind&#x60; is POLICY_TOKEN and the token is known. Null otherwise. | [optional] 
**UserSub** | Pointer to **NullableString** | The member&#39;s id, as returned by \&quot;List organization members\&quot;. Null for an actor that loaded no skill in the range, such as a policy token. | [optional] 
**UserName** | **string** | The username the audit log records for the actor. | 
**SkillLoads** | **int64** |  | 
**Actions** | **int64** |  | 
**Denials** | **int64** |  | 
**TopSkills** | **[]string** | The member&#39;s most used skills, at most 3, most used first. | 
**LastActivityAt** | **time.Time** |  | 

## Methods

### NewAgentActivityByMember

`func NewAgentActivityByMember(actorKind AgentActorKind, userName string, skillLoads int64, actions int64, denials int64, topSkills []string, lastActivityAt time.Time, ) *AgentActivityByMember`

NewAgentActivityByMember instantiates a new AgentActivityByMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityByMemberWithDefaults

`func NewAgentActivityByMemberWithDefaults() *AgentActivityByMember`

NewAgentActivityByMemberWithDefaults instantiates a new AgentActivityByMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorKind

`func (o *AgentActivityByMember) GetActorKind() AgentActorKind`

GetActorKind returns the ActorKind field if non-nil, zero value otherwise.

### GetActorKindOk

`func (o *AgentActivityByMember) GetActorKindOk() (*AgentActorKind, bool)`

GetActorKindOk returns a tuple with the ActorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorKind

`func (o *AgentActivityByMember) SetActorKind(v AgentActorKind)`

SetActorKind sets ActorKind field to given value.


### GetPolicyTokenId

`func (o *AgentActivityByMember) GetPolicyTokenId() string`

GetPolicyTokenId returns the PolicyTokenId field if non-nil, zero value otherwise.

### GetPolicyTokenIdOk

`func (o *AgentActivityByMember) GetPolicyTokenIdOk() (*string, bool)`

GetPolicyTokenIdOk returns a tuple with the PolicyTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTokenId

`func (o *AgentActivityByMember) SetPolicyTokenId(v string)`

SetPolicyTokenId sets PolicyTokenId field to given value.

### HasPolicyTokenId

`func (o *AgentActivityByMember) HasPolicyTokenId() bool`

HasPolicyTokenId returns a boolean if a field has been set.

### SetPolicyTokenIdNil

`func (o *AgentActivityByMember) SetPolicyTokenIdNil(b bool)`

 SetPolicyTokenIdNil sets the value for PolicyTokenId to be an explicit nil

### UnsetPolicyTokenId
`func (o *AgentActivityByMember) UnsetPolicyTokenId()`

UnsetPolicyTokenId ensures that no value is present for PolicyTokenId, not even an explicit nil
### GetUserSub

`func (o *AgentActivityByMember) GetUserSub() string`

GetUserSub returns the UserSub field if non-nil, zero value otherwise.

### GetUserSubOk

`func (o *AgentActivityByMember) GetUserSubOk() (*string, bool)`

GetUserSubOk returns a tuple with the UserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSub

`func (o *AgentActivityByMember) SetUserSub(v string)`

SetUserSub sets UserSub field to given value.

### HasUserSub

`func (o *AgentActivityByMember) HasUserSub() bool`

HasUserSub returns a boolean if a field has been set.

### SetUserSubNil

`func (o *AgentActivityByMember) SetUserSubNil(b bool)`

 SetUserSubNil sets the value for UserSub to be an explicit nil

### UnsetUserSub
`func (o *AgentActivityByMember) UnsetUserSub()`

UnsetUserSub ensures that no value is present for UserSub, not even an explicit nil
### GetUserName

`func (o *AgentActivityByMember) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AgentActivityByMember) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AgentActivityByMember) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetSkillLoads

`func (o *AgentActivityByMember) GetSkillLoads() int64`

GetSkillLoads returns the SkillLoads field if non-nil, zero value otherwise.

### GetSkillLoadsOk

`func (o *AgentActivityByMember) GetSkillLoadsOk() (*int64, bool)`

GetSkillLoadsOk returns a tuple with the SkillLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLoads

`func (o *AgentActivityByMember) SetSkillLoads(v int64)`

SetSkillLoads sets SkillLoads field to given value.


### GetActions

`func (o *AgentActivityByMember) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AgentActivityByMember) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AgentActivityByMember) SetActions(v int64)`

SetActions sets Actions field to given value.


### GetDenials

`func (o *AgentActivityByMember) GetDenials() int64`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *AgentActivityByMember) GetDenialsOk() (*int64, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *AgentActivityByMember) SetDenials(v int64)`

SetDenials sets Denials field to given value.


### GetTopSkills

`func (o *AgentActivityByMember) GetTopSkills() []string`

GetTopSkills returns the TopSkills field if non-nil, zero value otherwise.

### GetTopSkillsOk

`func (o *AgentActivityByMember) GetTopSkillsOk() (*[]string, bool)`

GetTopSkillsOk returns a tuple with the TopSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSkills

`func (o *AgentActivityByMember) SetTopSkills(v []string)`

SetTopSkills sets TopSkills field to given value.


### GetLastActivityAt

`func (o *AgentActivityByMember) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *AgentActivityByMember) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *AgentActivityByMember) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


