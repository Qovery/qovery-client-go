# SkillUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | Start of the range actually used, after defaults. | 
**To** | **time.Time** | End of the range actually used, after defaults. | 
**Bucket** | [**SkillUsageBucket**](SkillUsageBucket.md) |  | 
**Totals** | [**SkillUsageTotals**](SkillUsageTotals.md) |  | 
**Timeline** | [**[]SkillUsagePoint**](SkillUsagePoint.md) | Loads per hour or per day, oldest first. Buckets are cut in UTC, and buckets with no load are left out rather than returned as zero. | 
**BySkill** | [**[]SkillUsageBySkill**](SkillUsageBySkill.md) | Loads per skill, most loaded first, at most 100 rows. | 
**ByMember** | [**[]SkillUsageByMember**](SkillUsageByMember.md) | Loads per member, most active first, at most 500 rows. | 
**ByPreset** | [**[]SkillUsageByPreset**](SkillUsageByPreset.md) | Loads per preset, most loaded first. A row with a null preset groups the loads made with no preset. | 

## Methods

### NewSkillUsageResponse

`func NewSkillUsageResponse(from time.Time, to time.Time, bucket SkillUsageBucket, totals SkillUsageTotals, timeline []SkillUsagePoint, bySkill []SkillUsageBySkill, byMember []SkillUsageByMember, byPreset []SkillUsageByPreset, ) *SkillUsageResponse`

NewSkillUsageResponse instantiates a new SkillUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillUsageResponseWithDefaults

`func NewSkillUsageResponseWithDefaults() *SkillUsageResponse`

NewSkillUsageResponseWithDefaults instantiates a new SkillUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SkillUsageResponse) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SkillUsageResponse) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SkillUsageResponse) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *SkillUsageResponse) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SkillUsageResponse) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SkillUsageResponse) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetBucket

`func (o *SkillUsageResponse) GetBucket() SkillUsageBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *SkillUsageResponse) GetBucketOk() (*SkillUsageBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *SkillUsageResponse) SetBucket(v SkillUsageBucket)`

SetBucket sets Bucket field to given value.


### GetTotals

`func (o *SkillUsageResponse) GetTotals() SkillUsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *SkillUsageResponse) GetTotalsOk() (*SkillUsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *SkillUsageResponse) SetTotals(v SkillUsageTotals)`

SetTotals sets Totals field to given value.


### GetTimeline

`func (o *SkillUsageResponse) GetTimeline() []SkillUsagePoint`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *SkillUsageResponse) GetTimelineOk() (*[]SkillUsagePoint, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *SkillUsageResponse) SetTimeline(v []SkillUsagePoint)`

SetTimeline sets Timeline field to given value.


### GetBySkill

`func (o *SkillUsageResponse) GetBySkill() []SkillUsageBySkill`

GetBySkill returns the BySkill field if non-nil, zero value otherwise.

### GetBySkillOk

`func (o *SkillUsageResponse) GetBySkillOk() (*[]SkillUsageBySkill, bool)`

GetBySkillOk returns a tuple with the BySkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySkill

`func (o *SkillUsageResponse) SetBySkill(v []SkillUsageBySkill)`

SetBySkill sets BySkill field to given value.


### GetByMember

`func (o *SkillUsageResponse) GetByMember() []SkillUsageByMember`

GetByMember returns the ByMember field if non-nil, zero value otherwise.

### GetByMemberOk

`func (o *SkillUsageResponse) GetByMemberOk() (*[]SkillUsageByMember, bool)`

GetByMemberOk returns a tuple with the ByMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMember

`func (o *SkillUsageResponse) SetByMember(v []SkillUsageByMember)`

SetByMember sets ByMember field to given value.


### GetByPreset

`func (o *SkillUsageResponse) GetByPreset() []SkillUsageByPreset`

GetByPreset returns the ByPreset field if non-nil, zero value otherwise.

### GetByPresetOk

`func (o *SkillUsageResponse) GetByPresetOk() (*[]SkillUsageByPreset, bool)`

GetByPresetOk returns a tuple with the ByPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByPreset

`func (o *SkillUsageResponse) SetByPreset(v []SkillUsageByPreset)`

SetByPreset sets ByPreset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


