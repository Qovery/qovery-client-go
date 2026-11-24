# SkillUsageTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loads** | **int64** | Skill loads in the range, cached ones included. | 
**CachedLoads** | **int64** | Loads answered with a 304, where the agent reused its cached copy. | 
**Members** | **int64** | Distinct members whose agents loaded at least one skill. | 
**ApproxInjectedTokens** | **int64** | Estimated tokens of organization content and team context sent to agents, the number of characters divided by 4. Cached loads send nothing and add zero. The text of the Qovery skills themselves is not counted. | 

## Methods

### NewSkillUsageTotals

`func NewSkillUsageTotals(loads int64, cachedLoads int64, members int64, approxInjectedTokens int64, ) *SkillUsageTotals`

NewSkillUsageTotals instantiates a new SkillUsageTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillUsageTotalsWithDefaults

`func NewSkillUsageTotalsWithDefaults() *SkillUsageTotals`

NewSkillUsageTotalsWithDefaults instantiates a new SkillUsageTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoads

`func (o *SkillUsageTotals) GetLoads() int64`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *SkillUsageTotals) GetLoadsOk() (*int64, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *SkillUsageTotals) SetLoads(v int64)`

SetLoads sets Loads field to given value.


### GetCachedLoads

`func (o *SkillUsageTotals) GetCachedLoads() int64`

GetCachedLoads returns the CachedLoads field if non-nil, zero value otherwise.

### GetCachedLoadsOk

`func (o *SkillUsageTotals) GetCachedLoadsOk() (*int64, bool)`

GetCachedLoadsOk returns a tuple with the CachedLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedLoads

`func (o *SkillUsageTotals) SetCachedLoads(v int64)`

SetCachedLoads sets CachedLoads field to given value.


### GetMembers

`func (o *SkillUsageTotals) GetMembers() int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SkillUsageTotals) GetMembersOk() (*int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SkillUsageTotals) SetMembers(v int64)`

SetMembers sets Members field to given value.


### GetApproxInjectedTokens

`func (o *SkillUsageTotals) GetApproxInjectedTokens() int64`

GetApproxInjectedTokens returns the ApproxInjectedTokens field if non-nil, zero value otherwise.

### GetApproxInjectedTokensOk

`func (o *SkillUsageTotals) GetApproxInjectedTokensOk() (*int64, bool)`

GetApproxInjectedTokensOk returns a tuple with the ApproxInjectedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproxInjectedTokens

`func (o *SkillUsageTotals) SetApproxInjectedTokens(v int64)`

SetApproxInjectedTokens sets ApproxInjectedTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


