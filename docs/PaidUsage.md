# PaidUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDeploymentsPerMonth** | Pointer to **int32** |  | [optional] 
**ConsumedDeployments** | Pointer to **int32** |  | [optional] 
**MonthlyPlanCost** | Pointer to **float32** |  | [optional] 
**MonthlyPlanCostInCents** | Pointer to **int32** |  | [optional] 
**RemainingDeployments** | Pointer to **int32** |  | [optional] 
**DeploymentsExceeded** | Pointer to **bool** |  | [optional] 
**RenewalAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPaidUsage

`func NewPaidUsage() *PaidUsage`

NewPaidUsage instantiates a new PaidUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidUsageWithDefaults

`func NewPaidUsageWithDefaults() *PaidUsage`

NewPaidUsageWithDefaults instantiates a new PaidUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxDeploymentsPerMonth

`func (o *PaidUsage) GetMaxDeploymentsPerMonth() int32`

GetMaxDeploymentsPerMonth returns the MaxDeploymentsPerMonth field if non-nil, zero value otherwise.

### GetMaxDeploymentsPerMonthOk

`func (o *PaidUsage) GetMaxDeploymentsPerMonthOk() (*int32, bool)`

GetMaxDeploymentsPerMonthOk returns a tuple with the MaxDeploymentsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeploymentsPerMonth

`func (o *PaidUsage) SetMaxDeploymentsPerMonth(v int32)`

SetMaxDeploymentsPerMonth sets MaxDeploymentsPerMonth field to given value.

### HasMaxDeploymentsPerMonth

`func (o *PaidUsage) HasMaxDeploymentsPerMonth() bool`

HasMaxDeploymentsPerMonth returns a boolean if a field has been set.

### GetConsumedDeployments

`func (o *PaidUsage) GetConsumedDeployments() int32`

GetConsumedDeployments returns the ConsumedDeployments field if non-nil, zero value otherwise.

### GetConsumedDeploymentsOk

`func (o *PaidUsage) GetConsumedDeploymentsOk() (*int32, bool)`

GetConsumedDeploymentsOk returns a tuple with the ConsumedDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedDeployments

`func (o *PaidUsage) SetConsumedDeployments(v int32)`

SetConsumedDeployments sets ConsumedDeployments field to given value.

### HasConsumedDeployments

`func (o *PaidUsage) HasConsumedDeployments() bool`

HasConsumedDeployments returns a boolean if a field has been set.

### GetMonthlyPlanCost

`func (o *PaidUsage) GetMonthlyPlanCost() float32`

GetMonthlyPlanCost returns the MonthlyPlanCost field if non-nil, zero value otherwise.

### GetMonthlyPlanCostOk

`func (o *PaidUsage) GetMonthlyPlanCostOk() (*float32, bool)`

GetMonthlyPlanCostOk returns a tuple with the MonthlyPlanCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPlanCost

`func (o *PaidUsage) SetMonthlyPlanCost(v float32)`

SetMonthlyPlanCost sets MonthlyPlanCost field to given value.

### HasMonthlyPlanCost

`func (o *PaidUsage) HasMonthlyPlanCost() bool`

HasMonthlyPlanCost returns a boolean if a field has been set.

### GetMonthlyPlanCostInCents

`func (o *PaidUsage) GetMonthlyPlanCostInCents() int32`

GetMonthlyPlanCostInCents returns the MonthlyPlanCostInCents field if non-nil, zero value otherwise.

### GetMonthlyPlanCostInCentsOk

`func (o *PaidUsage) GetMonthlyPlanCostInCentsOk() (*int32, bool)`

GetMonthlyPlanCostInCentsOk returns a tuple with the MonthlyPlanCostInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPlanCostInCents

`func (o *PaidUsage) SetMonthlyPlanCostInCents(v int32)`

SetMonthlyPlanCostInCents sets MonthlyPlanCostInCents field to given value.

### HasMonthlyPlanCostInCents

`func (o *PaidUsage) HasMonthlyPlanCostInCents() bool`

HasMonthlyPlanCostInCents returns a boolean if a field has been set.

### GetRemainingDeployments

`func (o *PaidUsage) GetRemainingDeployments() int32`

GetRemainingDeployments returns the RemainingDeployments field if non-nil, zero value otherwise.

### GetRemainingDeploymentsOk

`func (o *PaidUsage) GetRemainingDeploymentsOk() (*int32, bool)`

GetRemainingDeploymentsOk returns a tuple with the RemainingDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDeployments

`func (o *PaidUsage) SetRemainingDeployments(v int32)`

SetRemainingDeployments sets RemainingDeployments field to given value.

### HasRemainingDeployments

`func (o *PaidUsage) HasRemainingDeployments() bool`

HasRemainingDeployments returns a boolean if a field has been set.

### GetDeploymentsExceeded

`func (o *PaidUsage) GetDeploymentsExceeded() bool`

GetDeploymentsExceeded returns the DeploymentsExceeded field if non-nil, zero value otherwise.

### GetDeploymentsExceededOk

`func (o *PaidUsage) GetDeploymentsExceededOk() (*bool, bool)`

GetDeploymentsExceededOk returns a tuple with the DeploymentsExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsExceeded

`func (o *PaidUsage) SetDeploymentsExceeded(v bool)`

SetDeploymentsExceeded sets DeploymentsExceeded field to given value.

### HasDeploymentsExceeded

`func (o *PaidUsage) HasDeploymentsExceeded() bool`

HasDeploymentsExceeded returns a boolean if a field has been set.

### GetRenewalAt

`func (o *PaidUsage) GetRenewalAt() time.Time`

GetRenewalAt returns the RenewalAt field if non-nil, zero value otherwise.

### GetRenewalAtOk

`func (o *PaidUsage) GetRenewalAtOk() (*time.Time, bool)`

GetRenewalAtOk returns a tuple with the RenewalAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAt

`func (o *PaidUsage) SetRenewalAt(v time.Time)`

SetRenewalAt sets RenewalAt field to given value.

### HasRenewalAt

`func (o *PaidUsage) HasRenewalAt() bool`

HasRenewalAt returns a boolean if a field has been set.

### SetRenewalAtNil

`func (o *PaidUsage) SetRenewalAtNil(b bool)`

 SetRenewalAtNil sets the value for RenewalAt to be an explicit nil

### UnsetRenewalAt
`func (o *PaidUsage) UnsetRenewalAt()`

UnsetRenewalAt ensures that no value is present for RenewalAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


