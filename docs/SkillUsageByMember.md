# SkillUsageByMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSub** | **string** | The member&#39;s id, as returned by \&quot;List organization members\&quot;. | 
**UserName** | **string** | The member&#39;s name at their most recent load in the range. | 
**Loads** | **int64** |  | 
**Skills** | **int64** | Distinct skills this member&#39;s agent loaded. | 
**LastLoadedAt** | **time.Time** |  | 

## Methods

### NewSkillUsageByMember

`func NewSkillUsageByMember(userSub string, userName string, loads int64, skills int64, lastLoadedAt time.Time, ) *SkillUsageByMember`

NewSkillUsageByMember instantiates a new SkillUsageByMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillUsageByMemberWithDefaults

`func NewSkillUsageByMemberWithDefaults() *SkillUsageByMember`

NewSkillUsageByMemberWithDefaults instantiates a new SkillUsageByMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSub

`func (o *SkillUsageByMember) GetUserSub() string`

GetUserSub returns the UserSub field if non-nil, zero value otherwise.

### GetUserSubOk

`func (o *SkillUsageByMember) GetUserSubOk() (*string, bool)`

GetUserSubOk returns a tuple with the UserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSub

`func (o *SkillUsageByMember) SetUserSub(v string)`

SetUserSub sets UserSub field to given value.


### GetUserName

`func (o *SkillUsageByMember) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SkillUsageByMember) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SkillUsageByMember) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetLoads

`func (o *SkillUsageByMember) GetLoads() int64`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *SkillUsageByMember) GetLoadsOk() (*int64, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *SkillUsageByMember) SetLoads(v int64)`

SetLoads sets Loads field to given value.


### GetSkills

`func (o *SkillUsageByMember) GetSkills() int64`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillUsageByMember) GetSkillsOk() (*int64, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillUsageByMember) SetSkills(v int64)`

SetSkills sets Skills field to given value.


### GetLastLoadedAt

`func (o *SkillUsageByMember) GetLastLoadedAt() time.Time`

GetLastLoadedAt returns the LastLoadedAt field if non-nil, zero value otherwise.

### GetLastLoadedAtOk

`func (o *SkillUsageByMember) GetLastLoadedAtOk() (*time.Time, bool)`

GetLastLoadedAtOk returns a tuple with the LastLoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoadedAt

`func (o *SkillUsageByMember) SetLastLoadedAt(v time.Time)`

SetLastLoadedAt sets LastLoadedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


