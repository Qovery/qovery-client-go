# CreditCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ExpiryMonth** | **int32** |  | 
**ExpiryYear** | **int32** |  | 
**LastDigit** | **string** |  | 
**IsExpired** | **bool** |  | 

## Methods

### NewCreditCardResponse

`func NewCreditCardResponse(id string, createdAt time.Time, expiryMonth int32, expiryYear int32, lastDigit string, isExpired bool, ) *CreditCardResponse`

NewCreditCardResponse instantiates a new CreditCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardResponseWithDefaults

`func NewCreditCardResponseWithDefaults() *CreditCardResponse`

NewCreditCardResponseWithDefaults instantiates a new CreditCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreditCardResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditCardResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditCardResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiryMonth

`func (o *CreditCardResponse) GetExpiryMonth() int32`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CreditCardResponse) GetExpiryMonthOk() (*int32, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CreditCardResponse) SetExpiryMonth(v int32)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetExpiryYear

`func (o *CreditCardResponse) GetExpiryYear() int32`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CreditCardResponse) GetExpiryYearOk() (*int32, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CreditCardResponse) SetExpiryYear(v int32)`

SetExpiryYear sets ExpiryYear field to given value.


### GetLastDigit

`func (o *CreditCardResponse) GetLastDigit() string`

GetLastDigit returns the LastDigit field if non-nil, zero value otherwise.

### GetLastDigitOk

`func (o *CreditCardResponse) GetLastDigitOk() (*string, bool)`

GetLastDigitOk returns a tuple with the LastDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDigit

`func (o *CreditCardResponse) SetLastDigit(v string)`

SetLastDigit sets LastDigit field to given value.


### GetIsExpired

`func (o *CreditCardResponse) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *CreditCardResponse) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *CreditCardResponse) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


