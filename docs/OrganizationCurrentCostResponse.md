# OrganizationCurrentCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalInCents** | **int32** |  | 
**Total** | **float32** |  | 
**CurrencyCode** | **string** |  | 
**BillingStartedOn** | Pointer to **time.Time** |  | [optional] 
**BillingEndedOn** | Pointer to **time.Time** |  | [optional] 
**BudgetExceeded** | Pointer to **bool** |  | [optional] 
**Projects** | Pointer to [**[]ProjectCurrentCostResponse**](ProjectCurrentCostResponse.md) |  | [optional] 

## Methods

### NewOrganizationCurrentCostResponse

`func NewOrganizationCurrentCostResponse(totalInCents int32, total float32, currencyCode string, ) *OrganizationCurrentCostResponse`

NewOrganizationCurrentCostResponse instantiates a new OrganizationCurrentCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCurrentCostResponseWithDefaults

`func NewOrganizationCurrentCostResponseWithDefaults() *OrganizationCurrentCostResponse`

NewOrganizationCurrentCostResponseWithDefaults instantiates a new OrganizationCurrentCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInCents

`func (o *OrganizationCurrentCostResponse) GetTotalInCents() int32`

GetTotalInCents returns the TotalInCents field if non-nil, zero value otherwise.

### GetTotalInCentsOk

`func (o *OrganizationCurrentCostResponse) GetTotalInCentsOk() (*int32, bool)`

GetTotalInCentsOk returns a tuple with the TotalInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInCents

`func (o *OrganizationCurrentCostResponse) SetTotalInCents(v int32)`

SetTotalInCents sets TotalInCents field to given value.


### GetTotal

`func (o *OrganizationCurrentCostResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrganizationCurrentCostResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrganizationCurrentCostResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetCurrencyCode

`func (o *OrganizationCurrentCostResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrganizationCurrentCostResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrganizationCurrentCostResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetBillingStartedOn

`func (o *OrganizationCurrentCostResponse) GetBillingStartedOn() time.Time`

GetBillingStartedOn returns the BillingStartedOn field if non-nil, zero value otherwise.

### GetBillingStartedOnOk

`func (o *OrganizationCurrentCostResponse) GetBillingStartedOnOk() (*time.Time, bool)`

GetBillingStartedOnOk returns a tuple with the BillingStartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartedOn

`func (o *OrganizationCurrentCostResponse) SetBillingStartedOn(v time.Time)`

SetBillingStartedOn sets BillingStartedOn field to given value.

### HasBillingStartedOn

`func (o *OrganizationCurrentCostResponse) HasBillingStartedOn() bool`

HasBillingStartedOn returns a boolean if a field has been set.

### GetBillingEndedOn

`func (o *OrganizationCurrentCostResponse) GetBillingEndedOn() time.Time`

GetBillingEndedOn returns the BillingEndedOn field if non-nil, zero value otherwise.

### GetBillingEndedOnOk

`func (o *OrganizationCurrentCostResponse) GetBillingEndedOnOk() (*time.Time, bool)`

GetBillingEndedOnOk returns a tuple with the BillingEndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndedOn

`func (o *OrganizationCurrentCostResponse) SetBillingEndedOn(v time.Time)`

SetBillingEndedOn sets BillingEndedOn field to given value.

### HasBillingEndedOn

`func (o *OrganizationCurrentCostResponse) HasBillingEndedOn() bool`

HasBillingEndedOn returns a boolean if a field has been set.

### GetBudgetExceeded

`func (o *OrganizationCurrentCostResponse) GetBudgetExceeded() bool`

GetBudgetExceeded returns the BudgetExceeded field if non-nil, zero value otherwise.

### GetBudgetExceededOk

`func (o *OrganizationCurrentCostResponse) GetBudgetExceededOk() (*bool, bool)`

GetBudgetExceededOk returns a tuple with the BudgetExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetExceeded

`func (o *OrganizationCurrentCostResponse) SetBudgetExceeded(v bool)`

SetBudgetExceeded sets BudgetExceeded field to given value.

### HasBudgetExceeded

`func (o *OrganizationCurrentCostResponse) HasBudgetExceeded() bool`

HasBudgetExceeded returns a boolean if a field has been set.

### GetProjects

`func (o *OrganizationCurrentCostResponse) GetProjects() []ProjectCurrentCostResponse`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationCurrentCostResponse) GetProjectsOk() (*[]ProjectCurrentCostResponse, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationCurrentCostResponse) SetProjects(v []ProjectCurrentCostResponse)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OrganizationCurrentCostResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


