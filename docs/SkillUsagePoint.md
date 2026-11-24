# SkillUsagePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** | Start of the hour or day, in UTC. | 
**Loads** | **int64** |  | 
**CachedLoads** | **int64** |  | 

## Methods

### NewSkillUsagePoint

`func NewSkillUsagePoint(start time.Time, loads int64, cachedLoads int64, ) *SkillUsagePoint`

NewSkillUsagePoint instantiates a new SkillUsagePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillUsagePointWithDefaults

`func NewSkillUsagePointWithDefaults() *SkillUsagePoint`

NewSkillUsagePointWithDefaults instantiates a new SkillUsagePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *SkillUsagePoint) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SkillUsagePoint) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SkillUsagePoint) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetLoads

`func (o *SkillUsagePoint) GetLoads() int64`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *SkillUsagePoint) GetLoadsOk() (*int64, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *SkillUsagePoint) SetLoads(v int64)`

SetLoads sets Loads field to given value.


### GetCachedLoads

`func (o *SkillUsagePoint) GetCachedLoads() int64`

GetCachedLoads returns the CachedLoads field if non-nil, zero value otherwise.

### GetCachedLoadsOk

`func (o *SkillUsagePoint) GetCachedLoadsOk() (*int64, bool)`

GetCachedLoadsOk returns a tuple with the CachedLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedLoads

`func (o *SkillUsagePoint) SetCachedLoads(v int64)`

SetCachedLoads sets CachedLoads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


