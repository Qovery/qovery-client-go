# CreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ExpiryMonth** | **int32** |  | 
**ExpiryYear** | **int32** |  | 
**LastDigit** | **string** |  | 
**IsExpired** | **bool** |  | 
**Brand** | **string** |  | 

## Methods

### NewCreditCard

`func NewCreditCard(id string, createdAt time.Time, expiryMonth int32, expiryYear int32, lastDigit string, isExpired bool, brand string, ) *CreditCard`

NewCreditCard instantiates a new CreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardWithDefaults

`func NewCreditCardWithDefaults() *CreditCard`

NewCreditCardWithDefaults instantiates a new CreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCard) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreditCard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditCard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditCard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiryMonth

`func (o *CreditCard) GetExpiryMonth() int32`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CreditCard) GetExpiryMonthOk() (*int32, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CreditCard) SetExpiryMonth(v int32)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetExpiryYear

`func (o *CreditCard) GetExpiryYear() int32`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CreditCard) GetExpiryYearOk() (*int32, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CreditCard) SetExpiryYear(v int32)`

SetExpiryYear sets ExpiryYear field to given value.


### GetLastDigit

`func (o *CreditCard) GetLastDigit() string`

GetLastDigit returns the LastDigit field if non-nil, zero value otherwise.

### GetLastDigitOk

`func (o *CreditCard) GetLastDigitOk() (*string, bool)`

GetLastDigitOk returns a tuple with the LastDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDigit

`func (o *CreditCard) SetLastDigit(v string)`

SetLastDigit sets LastDigit field to given value.


### GetIsExpired

`func (o *CreditCard) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *CreditCard) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *CreditCard) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.


### GetBrand

`func (o *CreditCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreditCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreditCard) SetBrand(v string)`

SetBrand sets Brand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


