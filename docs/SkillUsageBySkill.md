# SkillUsageBySkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillName** | **string** |  | 
**Loads** | **int64** |  | 
**CachedLoads** | **int64** |  | 
**Members** | **int64** | Distinct members who loaded this skill. | 
**LastLoadedAt** | **time.Time** |  | 

## Methods

### NewSkillUsageBySkill

`func NewSkillUsageBySkill(skillName string, loads int64, cachedLoads int64, members int64, lastLoadedAt time.Time, ) *SkillUsageBySkill`

NewSkillUsageBySkill instantiates a new SkillUsageBySkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillUsageBySkillWithDefaults

`func NewSkillUsageBySkillWithDefaults() *SkillUsageBySkill`

NewSkillUsageBySkillWithDefaults instantiates a new SkillUsageBySkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillName

`func (o *SkillUsageBySkill) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *SkillUsageBySkill) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *SkillUsageBySkill) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.


### GetLoads

`func (o *SkillUsageBySkill) GetLoads() int64`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *SkillUsageBySkill) GetLoadsOk() (*int64, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *SkillUsageBySkill) SetLoads(v int64)`

SetLoads sets Loads field to given value.


### GetCachedLoads

`func (o *SkillUsageBySkill) GetCachedLoads() int64`

GetCachedLoads returns the CachedLoads field if non-nil, zero value otherwise.

### GetCachedLoadsOk

`func (o *SkillUsageBySkill) GetCachedLoadsOk() (*int64, bool)`

GetCachedLoadsOk returns a tuple with the CachedLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedLoads

`func (o *SkillUsageBySkill) SetCachedLoads(v int64)`

SetCachedLoads sets CachedLoads field to given value.


### GetMembers

`func (o *SkillUsageBySkill) GetMembers() int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SkillUsageBySkill) GetMembersOk() (*int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SkillUsageBySkill) SetMembers(v int64)`

SetMembers sets Members field to given value.


### GetLastLoadedAt

`func (o *SkillUsageBySkill) GetLastLoadedAt() time.Time`

GetLastLoadedAt returns the LastLoadedAt field if non-nil, zero value otherwise.

### GetLastLoadedAtOk

`func (o *SkillUsageBySkill) GetLastLoadedAtOk() (*time.Time, bool)`

GetLastLoadedAtOk returns a tuple with the LastLoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoadedAt

`func (o *SkillUsageBySkill) SetLastLoadedAt(v time.Time)`

SetLastLoadedAt sets LastLoadedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


