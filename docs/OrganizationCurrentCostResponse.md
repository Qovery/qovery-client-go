# OrganizationCurrentCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to **string** |  | [optional] 
**RemainingTrialDay** | Pointer to **int32** | number of days remaining before the end of the trial period | [optional] 
**RemainingCredits** | Pointer to [**RemainingCredits**](RemainingCredits.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**PaidUsage** | Pointer to [**PaidUsageResponse**](PaidUsageResponse.md) |  | [optional] 
**CommunityUsage** | Pointer to [**CommunityUsageResponse**](CommunityUsageResponse.md) |  | [optional] 

## Methods

### NewOrganizationCurrentCostResponse

`func NewOrganizationCurrentCostResponse() *OrganizationCurrentCostResponse`

NewOrganizationCurrentCostResponse instantiates a new OrganizationCurrentCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCurrentCostResponseWithDefaults

`func NewOrganizationCurrentCostResponseWithDefaults() *OrganizationCurrentCostResponse`

NewOrganizationCurrentCostResponseWithDefaults instantiates a new OrganizationCurrentCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *OrganizationCurrentCostResponse) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationCurrentCostResponse) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationCurrentCostResponse) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrganizationCurrentCostResponse) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRemainingTrialDay

`func (o *OrganizationCurrentCostResponse) GetRemainingTrialDay() int32`

GetRemainingTrialDay returns the RemainingTrialDay field if non-nil, zero value otherwise.

### GetRemainingTrialDayOk

`func (o *OrganizationCurrentCostResponse) GetRemainingTrialDayOk() (*int32, bool)`

GetRemainingTrialDayOk returns a tuple with the RemainingTrialDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTrialDay

`func (o *OrganizationCurrentCostResponse) SetRemainingTrialDay(v int32)`

SetRemainingTrialDay sets RemainingTrialDay field to given value.

### HasRemainingTrialDay

`func (o *OrganizationCurrentCostResponse) HasRemainingTrialDay() bool`

HasRemainingTrialDay returns a boolean if a field has been set.

### GetRemainingCredits

`func (o *OrganizationCurrentCostResponse) GetRemainingCredits() RemainingCredits`

GetRemainingCredits returns the RemainingCredits field if non-nil, zero value otherwise.

### GetRemainingCreditsOk

`func (o *OrganizationCurrentCostResponse) GetRemainingCreditsOk() (*RemainingCredits, bool)`

GetRemainingCreditsOk returns a tuple with the RemainingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredits

`func (o *OrganizationCurrentCostResponse) SetRemainingCredits(v RemainingCredits)`

SetRemainingCredits sets RemainingCredits field to given value.

### HasRemainingCredits

`func (o *OrganizationCurrentCostResponse) HasRemainingCredits() bool`

HasRemainingCredits returns a boolean if a field has been set.

### GetCost

`func (o *OrganizationCurrentCostResponse) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OrganizationCurrentCostResponse) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OrganizationCurrentCostResponse) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OrganizationCurrentCostResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPaidUsage

`func (o *OrganizationCurrentCostResponse) GetPaidUsage() PaidUsageResponse`

GetPaidUsage returns the PaidUsage field if non-nil, zero value otherwise.

### GetPaidUsageOk

`func (o *OrganizationCurrentCostResponse) GetPaidUsageOk() (*PaidUsageResponse, bool)`

GetPaidUsageOk returns a tuple with the PaidUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidUsage

`func (o *OrganizationCurrentCostResponse) SetPaidUsage(v PaidUsageResponse)`

SetPaidUsage sets PaidUsage field to given value.

### HasPaidUsage

`func (o *OrganizationCurrentCostResponse) HasPaidUsage() bool`

HasPaidUsage returns a boolean if a field has been set.

### GetCommunityUsage

`func (o *OrganizationCurrentCostResponse) GetCommunityUsage() CommunityUsageResponse`

GetCommunityUsage returns the CommunityUsage field if non-nil, zero value otherwise.

### GetCommunityUsageOk

`func (o *OrganizationCurrentCostResponse) GetCommunityUsageOk() (*CommunityUsageResponse, bool)`

GetCommunityUsageOk returns a tuple with the CommunityUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityUsage

`func (o *OrganizationCurrentCostResponse) SetCommunityUsage(v CommunityUsageResponse)`

SetCommunityUsage sets CommunityUsage field to given value.

### HasCommunityUsage

`func (o *OrganizationCurrentCostResponse) HasCommunityUsage() bool`

HasCommunityUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


