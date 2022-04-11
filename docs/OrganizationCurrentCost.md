# OrganizationCurrentCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**PlanEnum**](PlanEnum.md) |  | [optional] 
**RemainingTrialDay** | Pointer to **int32** | number of days remaining before the end of the trial period | [optional] 
**RemainingCredits** | Pointer to [**RemainingCredits**](RemainingCredits.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**PaidUsage** | Pointer to [**PaidUsage**](PaidUsage.md) |  | [optional] 
**CommunityUsage** | Pointer to [**CommunityUsage**](CommunityUsage.md) |  | [optional] 

## Methods

### NewOrganizationCurrentCost

`func NewOrganizationCurrentCost() *OrganizationCurrentCost`

NewOrganizationCurrentCost instantiates a new OrganizationCurrentCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCurrentCostWithDefaults

`func NewOrganizationCurrentCostWithDefaults() *OrganizationCurrentCost`

NewOrganizationCurrentCostWithDefaults instantiates a new OrganizationCurrentCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *OrganizationCurrentCost) GetPlan() PlanEnum`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationCurrentCost) GetPlanOk() (*PlanEnum, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationCurrentCost) SetPlan(v PlanEnum)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrganizationCurrentCost) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRemainingTrialDay

`func (o *OrganizationCurrentCost) GetRemainingTrialDay() int32`

GetRemainingTrialDay returns the RemainingTrialDay field if non-nil, zero value otherwise.

### GetRemainingTrialDayOk

`func (o *OrganizationCurrentCost) GetRemainingTrialDayOk() (*int32, bool)`

GetRemainingTrialDayOk returns a tuple with the RemainingTrialDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTrialDay

`func (o *OrganizationCurrentCost) SetRemainingTrialDay(v int32)`

SetRemainingTrialDay sets RemainingTrialDay field to given value.

### HasRemainingTrialDay

`func (o *OrganizationCurrentCost) HasRemainingTrialDay() bool`

HasRemainingTrialDay returns a boolean if a field has been set.

### GetRemainingCredits

`func (o *OrganizationCurrentCost) GetRemainingCredits() RemainingCredits`

GetRemainingCredits returns the RemainingCredits field if non-nil, zero value otherwise.

### GetRemainingCreditsOk

`func (o *OrganizationCurrentCost) GetRemainingCreditsOk() (*RemainingCredits, bool)`

GetRemainingCreditsOk returns a tuple with the RemainingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredits

`func (o *OrganizationCurrentCost) SetRemainingCredits(v RemainingCredits)`

SetRemainingCredits sets RemainingCredits field to given value.

### HasRemainingCredits

`func (o *OrganizationCurrentCost) HasRemainingCredits() bool`

HasRemainingCredits returns a boolean if a field has been set.

### GetCost

`func (o *OrganizationCurrentCost) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OrganizationCurrentCost) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OrganizationCurrentCost) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OrganizationCurrentCost) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPaidUsage

`func (o *OrganizationCurrentCost) GetPaidUsage() PaidUsage`

GetPaidUsage returns the PaidUsage field if non-nil, zero value otherwise.

### GetPaidUsageOk

`func (o *OrganizationCurrentCost) GetPaidUsageOk() (*PaidUsage, bool)`

GetPaidUsageOk returns a tuple with the PaidUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidUsage

`func (o *OrganizationCurrentCost) SetPaidUsage(v PaidUsage)`

SetPaidUsage sets PaidUsage field to given value.

### HasPaidUsage

`func (o *OrganizationCurrentCost) HasPaidUsage() bool`

HasPaidUsage returns a boolean if a field has been set.

### GetCommunityUsage

`func (o *OrganizationCurrentCost) GetCommunityUsage() CommunityUsage`

GetCommunityUsage returns the CommunityUsage field if non-nil, zero value otherwise.

### GetCommunityUsageOk

`func (o *OrganizationCurrentCost) GetCommunityUsageOk() (*CommunityUsage, bool)`

GetCommunityUsageOk returns a tuple with the CommunityUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityUsage

`func (o *OrganizationCurrentCost) SetCommunityUsage(v CommunityUsage)`

SetCommunityUsage sets CommunityUsage field to given value.

### HasCommunityUsage

`func (o *OrganizationCurrentCost) HasCommunityUsage() bool`

HasCommunityUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


