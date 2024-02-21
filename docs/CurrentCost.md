# CurrentCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**PlanEnum**](PlanEnum.md) |  | [optional] 
**RemainingTrialDay** | Pointer to **int32** | number of days remaining before the end of the trial period | [optional] 
**RenewalAt** | Pointer to **NullableTime** | date when the current plan will be renewed | [optional] [readonly] 
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

### GetRenewalAt

`func (o *CurrentCost) GetRenewalAt() time.Time`

GetRenewalAt returns the RenewalAt field if non-nil, zero value otherwise.

### GetRenewalAtOk

`func (o *CurrentCost) GetRenewalAtOk() (*time.Time, bool)`

GetRenewalAtOk returns a tuple with the RenewalAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAt

`func (o *CurrentCost) SetRenewalAt(v time.Time)`

SetRenewalAt sets RenewalAt field to given value.

### HasRenewalAt

`func (o *CurrentCost) HasRenewalAt() bool`

HasRenewalAt returns a boolean if a field has been set.

### SetRenewalAtNil

`func (o *CurrentCost) SetRenewalAtNil(b bool)`

 SetRenewalAtNil sets the value for RenewalAt to be an explicit nil

### UnsetRenewalAt
`func (o *CurrentCost) UnsetRenewalAt()`

UnsetRenewalAt ensures that no value is present for RenewalAt, not even an explicit nil
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


