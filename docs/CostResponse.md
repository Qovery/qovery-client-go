# CostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalInCents** | **int32** |  | 
**Total** | **float32** |  | 
**CurrencyCode** | **string** |  | 

## Methods

### NewCostResponse

`func NewCostResponse(totalInCents int32, total float32, currencyCode string, ) *CostResponse`

NewCostResponse instantiates a new CostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostResponseWithDefaults

`func NewCostResponseWithDefaults() *CostResponse`

NewCostResponseWithDefaults instantiates a new CostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInCents

`func (o *CostResponse) GetTotalInCents() int32`

GetTotalInCents returns the TotalInCents field if non-nil, zero value otherwise.

### GetTotalInCentsOk

`func (o *CostResponse) GetTotalInCentsOk() (*int32, bool)`

GetTotalInCentsOk returns a tuple with the TotalInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInCents

`func (o *CostResponse) SetTotalInCents(v int32)`

SetTotalInCents sets TotalInCents field to given value.


### GetTotal

`func (o *CostResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CostResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CostResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetCurrencyCode

`func (o *CostResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CostResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CostResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


