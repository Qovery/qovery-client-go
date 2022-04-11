# Budget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalInCents** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 

## Methods

### NewBudget

`func NewBudget() *Budget`

NewBudget instantiates a new Budget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetWithDefaults

`func NewBudgetWithDefaults() *Budget`

NewBudgetWithDefaults instantiates a new Budget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInCents

`func (o *Budget) GetTotalInCents() int32`

GetTotalInCents returns the TotalInCents field if non-nil, zero value otherwise.

### GetTotalInCentsOk

`func (o *Budget) GetTotalInCentsOk() (*int32, bool)`

GetTotalInCentsOk returns a tuple with the TotalInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInCents

`func (o *Budget) SetTotalInCents(v int32)`

SetTotalInCents sets TotalInCents field to given value.

### HasTotalInCents

`func (o *Budget) HasTotalInCents() bool`

HasTotalInCents returns a boolean if a field has been set.

### GetTotal

`func (o *Budget) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Budget) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Budget) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Budget) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Budget) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Budget) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Budget) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Budget) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


