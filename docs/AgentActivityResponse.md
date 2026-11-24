# AgentActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | Start of the range actually used, after defaults. | 
**To** | **time.Time** | End of the range actually used, after defaults. | 
**Bucket** | [**SkillUsageBucket**](SkillUsageBucket.md) |  | 
**Coverage** | [**AgentActivityCoverage**](AgentActivityCoverage.md) |  | 
**Totals** | [**AgentActivityTotals**](AgentActivityTotals.md) |  | 
**Timeline** | [**[]AgentActivityPoint**](AgentActivityPoint.md) | Activity per hour or per day, oldest first. Buckets are cut in UTC; empty buckets are omitted. | 
**ActionsByType** | [**[]AgentActionsByType**](AgentActionsByType.md) | Agent actions per audit event and target type, most frequent first, at most 200 rows. | 
**ByMember** | [**[]AgentActivityByMember**](AgentActivityByMember.md) | Activity per member, most active first, at most 500 rows. | 
**BySkill** | [**[]AgentActivityBySkill**](AgentActivityBySkill.md) | Activity per skill, most active first, at most 100 rows. | 
**ByPolicyToken** | [**[]AgentActivityByPolicyToken**](AgentActivityByPolicyToken.md) | Activity per policy token, most active first, at most 100 rows. Always empty on &#x60;getMyAgentActivity&#x60;. | 
**Recent** | [**[]AgentActivityEntry**](AgentActivityEntry.md) | The 100 most recent agent actions, newest first. | 

## Methods

### NewAgentActivityResponse

`func NewAgentActivityResponse(from time.Time, to time.Time, bucket SkillUsageBucket, coverage AgentActivityCoverage, totals AgentActivityTotals, timeline []AgentActivityPoint, actionsByType []AgentActionsByType, byMember []AgentActivityByMember, bySkill []AgentActivityBySkill, byPolicyToken []AgentActivityByPolicyToken, recent []AgentActivityEntry, ) *AgentActivityResponse`

NewAgentActivityResponse instantiates a new AgentActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityResponseWithDefaults

`func NewAgentActivityResponseWithDefaults() *AgentActivityResponse`

NewAgentActivityResponseWithDefaults instantiates a new AgentActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *AgentActivityResponse) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AgentActivityResponse) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AgentActivityResponse) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *AgentActivityResponse) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AgentActivityResponse) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AgentActivityResponse) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetBucket

`func (o *AgentActivityResponse) GetBucket() SkillUsageBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AgentActivityResponse) GetBucketOk() (*SkillUsageBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AgentActivityResponse) SetBucket(v SkillUsageBucket)`

SetBucket sets Bucket field to given value.


### GetCoverage

`func (o *AgentActivityResponse) GetCoverage() AgentActivityCoverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *AgentActivityResponse) GetCoverageOk() (*AgentActivityCoverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *AgentActivityResponse) SetCoverage(v AgentActivityCoverage)`

SetCoverage sets Coverage field to given value.


### GetTotals

`func (o *AgentActivityResponse) GetTotals() AgentActivityTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AgentActivityResponse) GetTotalsOk() (*AgentActivityTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AgentActivityResponse) SetTotals(v AgentActivityTotals)`

SetTotals sets Totals field to given value.


### GetTimeline

`func (o *AgentActivityResponse) GetTimeline() []AgentActivityPoint`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *AgentActivityResponse) GetTimelineOk() (*[]AgentActivityPoint, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *AgentActivityResponse) SetTimeline(v []AgentActivityPoint)`

SetTimeline sets Timeline field to given value.


### GetActionsByType

`func (o *AgentActivityResponse) GetActionsByType() []AgentActionsByType`

GetActionsByType returns the ActionsByType field if non-nil, zero value otherwise.

### GetActionsByTypeOk

`func (o *AgentActivityResponse) GetActionsByTypeOk() (*[]AgentActionsByType, bool)`

GetActionsByTypeOk returns a tuple with the ActionsByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsByType

`func (o *AgentActivityResponse) SetActionsByType(v []AgentActionsByType)`

SetActionsByType sets ActionsByType field to given value.


### GetByMember

`func (o *AgentActivityResponse) GetByMember() []AgentActivityByMember`

GetByMember returns the ByMember field if non-nil, zero value otherwise.

### GetByMemberOk

`func (o *AgentActivityResponse) GetByMemberOk() (*[]AgentActivityByMember, bool)`

GetByMemberOk returns a tuple with the ByMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMember

`func (o *AgentActivityResponse) SetByMember(v []AgentActivityByMember)`

SetByMember sets ByMember field to given value.


### GetBySkill

`func (o *AgentActivityResponse) GetBySkill() []AgentActivityBySkill`

GetBySkill returns the BySkill field if non-nil, zero value otherwise.

### GetBySkillOk

`func (o *AgentActivityResponse) GetBySkillOk() (*[]AgentActivityBySkill, bool)`

GetBySkillOk returns a tuple with the BySkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySkill

`func (o *AgentActivityResponse) SetBySkill(v []AgentActivityBySkill)`

SetBySkill sets BySkill field to given value.


### GetByPolicyToken

`func (o *AgentActivityResponse) GetByPolicyToken() []AgentActivityByPolicyToken`

GetByPolicyToken returns the ByPolicyToken field if non-nil, zero value otherwise.

### GetByPolicyTokenOk

`func (o *AgentActivityResponse) GetByPolicyTokenOk() (*[]AgentActivityByPolicyToken, bool)`

GetByPolicyTokenOk returns a tuple with the ByPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByPolicyToken

`func (o *AgentActivityResponse) SetByPolicyToken(v []AgentActivityByPolicyToken)`

SetByPolicyToken sets ByPolicyToken field to given value.


### GetRecent

`func (o *AgentActivityResponse) GetRecent() []AgentActivityEntry`

GetRecent returns the Recent field if non-nil, zero value otherwise.

### GetRecentOk

`func (o *AgentActivityResponse) GetRecentOk() (*[]AgentActivityEntry, bool)`

GetRecentOk returns a tuple with the Recent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecent

`func (o *AgentActivityResponse) SetRecent(v []AgentActivityEntry)`

SetRecent sets Recent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


