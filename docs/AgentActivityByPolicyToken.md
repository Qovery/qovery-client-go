# AgentActivityByPolicyToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyTokenId** | Pointer to **NullableString** | The token&#39;s id. Null groups older role refusals recorded before refusals named the token: a policy token, but not which one. | [optional] 
**TokenName** | Pointer to **NullableString** | The token&#39;s current name, or its last audited name when &#x60;deleted&#x60; is true. Null when unknown. | [optional] 
**Deleted** | **bool** | Whether the token no longer exists. | 
**RoleId** | Pointer to **NullableString** | The token&#39;s current role. Null when the token is deleted or unknown. | [optional] 
**RoleName** | Pointer to **NullableString** | The name of the token&#39;s current role. Null when the token is deleted or unknown. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | When the token expires. Null when it never expires, or when it is deleted or unknown. | [optional] 
**SkillLoads** | **int64** |  | 
**Actions** | **int64** |  | 
**Denials** | **int64** |  | 
**LastActivityAt** | **time.Time** |  | 
**TopActions** | [**[]AgentActionsByType**](AgentActionsByType.md) | The token&#39;s most frequent actions, at most 3, most frequent first. | 
**TopSkills** | **[]string** | The token&#39;s most used skills, at most 3, most used first. | 
**Timeline** | [**[]AgentActivityTokenPoint**](AgentActivityTokenPoint.md) | The token&#39;s actions per bucket of the requested size, oldest first, only for buckets where it acted. Filled when &#x60;bucket&#x60; is DAY, or when &#x60;policyTokenId&#x60; is set (so HOUR buckets need &#x60;policyTokenId&#x60;). Empty otherwise. | 

## Methods

### NewAgentActivityByPolicyToken

`func NewAgentActivityByPolicyToken(deleted bool, skillLoads int64, actions int64, denials int64, lastActivityAt time.Time, topActions []AgentActionsByType, topSkills []string, timeline []AgentActivityTokenPoint, ) *AgentActivityByPolicyToken`

NewAgentActivityByPolicyToken instantiates a new AgentActivityByPolicyToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityByPolicyTokenWithDefaults

`func NewAgentActivityByPolicyTokenWithDefaults() *AgentActivityByPolicyToken`

NewAgentActivityByPolicyTokenWithDefaults instantiates a new AgentActivityByPolicyToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyTokenId

`func (o *AgentActivityByPolicyToken) GetPolicyTokenId() string`

GetPolicyTokenId returns the PolicyTokenId field if non-nil, zero value otherwise.

### GetPolicyTokenIdOk

`func (o *AgentActivityByPolicyToken) GetPolicyTokenIdOk() (*string, bool)`

GetPolicyTokenIdOk returns a tuple with the PolicyTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTokenId

`func (o *AgentActivityByPolicyToken) SetPolicyTokenId(v string)`

SetPolicyTokenId sets PolicyTokenId field to given value.

### HasPolicyTokenId

`func (o *AgentActivityByPolicyToken) HasPolicyTokenId() bool`

HasPolicyTokenId returns a boolean if a field has been set.

### SetPolicyTokenIdNil

`func (o *AgentActivityByPolicyToken) SetPolicyTokenIdNil(b bool)`

 SetPolicyTokenIdNil sets the value for PolicyTokenId to be an explicit nil

### UnsetPolicyTokenId
`func (o *AgentActivityByPolicyToken) UnsetPolicyTokenId()`

UnsetPolicyTokenId ensures that no value is present for PolicyTokenId, not even an explicit nil
### GetTokenName

`func (o *AgentActivityByPolicyToken) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *AgentActivityByPolicyToken) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *AgentActivityByPolicyToken) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *AgentActivityByPolicyToken) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### SetTokenNameNil

`func (o *AgentActivityByPolicyToken) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *AgentActivityByPolicyToken) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetDeleted

`func (o *AgentActivityByPolicyToken) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *AgentActivityByPolicyToken) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *AgentActivityByPolicyToken) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetRoleId

`func (o *AgentActivityByPolicyToken) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AgentActivityByPolicyToken) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AgentActivityByPolicyToken) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *AgentActivityByPolicyToken) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### SetRoleIdNil

`func (o *AgentActivityByPolicyToken) SetRoleIdNil(b bool)`

 SetRoleIdNil sets the value for RoleId to be an explicit nil

### UnsetRoleId
`func (o *AgentActivityByPolicyToken) UnsetRoleId()`

UnsetRoleId ensures that no value is present for RoleId, not even an explicit nil
### GetRoleName

`func (o *AgentActivityByPolicyToken) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AgentActivityByPolicyToken) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AgentActivityByPolicyToken) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AgentActivityByPolicyToken) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *AgentActivityByPolicyToken) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *AgentActivityByPolicyToken) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetExpiresAt

`func (o *AgentActivityByPolicyToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AgentActivityByPolicyToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AgentActivityByPolicyToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AgentActivityByPolicyToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *AgentActivityByPolicyToken) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AgentActivityByPolicyToken) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSkillLoads

`func (o *AgentActivityByPolicyToken) GetSkillLoads() int64`

GetSkillLoads returns the SkillLoads field if non-nil, zero value otherwise.

### GetSkillLoadsOk

`func (o *AgentActivityByPolicyToken) GetSkillLoadsOk() (*int64, bool)`

GetSkillLoadsOk returns a tuple with the SkillLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLoads

`func (o *AgentActivityByPolicyToken) SetSkillLoads(v int64)`

SetSkillLoads sets SkillLoads field to given value.


### GetActions

`func (o *AgentActivityByPolicyToken) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AgentActivityByPolicyToken) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AgentActivityByPolicyToken) SetActions(v int64)`

SetActions sets Actions field to given value.


### GetDenials

`func (o *AgentActivityByPolicyToken) GetDenials() int64`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *AgentActivityByPolicyToken) GetDenialsOk() (*int64, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *AgentActivityByPolicyToken) SetDenials(v int64)`

SetDenials sets Denials field to given value.


### GetLastActivityAt

`func (o *AgentActivityByPolicyToken) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *AgentActivityByPolicyToken) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *AgentActivityByPolicyToken) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.


### GetTopActions

`func (o *AgentActivityByPolicyToken) GetTopActions() []AgentActionsByType`

GetTopActions returns the TopActions field if non-nil, zero value otherwise.

### GetTopActionsOk

`func (o *AgentActivityByPolicyToken) GetTopActionsOk() (*[]AgentActionsByType, bool)`

GetTopActionsOk returns a tuple with the TopActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopActions

`func (o *AgentActivityByPolicyToken) SetTopActions(v []AgentActionsByType)`

SetTopActions sets TopActions field to given value.


### GetTopSkills

`func (o *AgentActivityByPolicyToken) GetTopSkills() []string`

GetTopSkills returns the TopSkills field if non-nil, zero value otherwise.

### GetTopSkillsOk

`func (o *AgentActivityByPolicyToken) GetTopSkillsOk() (*[]string, bool)`

GetTopSkillsOk returns a tuple with the TopSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSkills

`func (o *AgentActivityByPolicyToken) SetTopSkills(v []string)`

SetTopSkills sets TopSkills field to given value.


### GetTimeline

`func (o *AgentActivityByPolicyToken) GetTimeline() []AgentActivityTokenPoint`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *AgentActivityByPolicyToken) GetTimelineOk() (*[]AgentActivityTokenPoint, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *AgentActivityByPolicyToken) SetTimeline(v []AgentActivityTokenPoint)`

SetTimeline sets Timeline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


