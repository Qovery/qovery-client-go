# CurrentCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**PlanEnum**](PlanEnum.md) |  | [optional] 
**RemainingTrialDay** | Pointer to **int32** | number of days remaining before the end of the trial period | [optional] 
**RemainingCredits** | Pointer to [**RemainingCredits**](RemainingCredits.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 

## Methods

### NewCurrentCost

`func NewCurrentCost() *CurrentCost`

NewCurrentCost instantiates a new CurrentCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentCostWithDefaults

`func NewCurrentCostWithDefaults() *CurrentCost`

NewCurrentCostWithDefaults instantiates a new CurrentCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *CurrentCost) GetPlan() PlanEnum`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CurrentCost) GetPlanOk() (*PlanEnum, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CurrentCost) SetPlan(v PlanEnum)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CurrentCost) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRemainingTrialDay

`func (o *CurrentCost) GetRemainingTrialDay() int32`

GetRemainingTrialDay returns the RemainingTrialDay field if non-nil, zero value otherwise.

### GetRemainingTrialDayOk

`func (o *CurrentCost) GetRemainingTrialDayOk() (*int32, bool)`

GetRemainingTrialDayOk returns a tuple with the RemainingTrialDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTrialDay

`func (o *CurrentCost) SetRemainingTrialDay(v int32)`

SetRemainingTrialDay sets RemainingTrialDay field to given value.

### HasRemainingTrialDay

`func (o *CurrentCost) HasRemainingTrialDay() bool`

HasRemainingTrialDay returns a boolean if a field has been set.

### GetRemainingCredits

`func (o *CurrentCost) GetRemainingCredits() RemainingCredits`

GetRemainingCredits returns the RemainingCredits field if non-nil, zero value otherwise.

### GetRemainingCreditsOk

`func (o *CurrentCost) GetRemainingCreditsOk() (*RemainingCredits, bool)`

GetRemainingCreditsOk returns a tuple with the RemainingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredits

`func (o *CurrentCost) SetRemainingCredits(v RemainingCredits)`

SetRemainingCredits sets RemainingCredits field to given value.

### HasRemainingCredits

`func (o *CurrentCost) HasRemainingCredits() bool`

HasRemainingCredits returns a boolean if a field has been set.

### GetCost

`func (o *CurrentCost) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CurrentCost) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CurrentCost) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CurrentCost) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


