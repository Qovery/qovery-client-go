# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalInCents** | **int32** |  | 
**Total** | **float32** |  | 
**CurrencyCode** | **string** |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Status** | [**InvoiceStatusEnum**](InvoiceStatusEnum.md) |  | 

## Methods

### NewInvoice

`func NewInvoice(totalInCents int32, total float32, currencyCode string, id string, createdAt time.Time, status InvoiceStatusEnum, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInCents

`func (o *Invoice) GetTotalInCents() int32`

GetTotalInCents returns the TotalInCents field if non-nil, zero value otherwise.

### GetTotalInCentsOk

`func (o *Invoice) GetTotalInCentsOk() (*int32, bool)`

GetTotalInCentsOk returns a tuple with the TotalInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInCents

`func (o *Invoice) SetTotalInCents(v int32)`

SetTotalInCents sets TotalInCents field to given value.


### GetTotal

`func (o *Invoice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetCurrencyCode

`func (o *Invoice) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Invoice) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Invoice) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Invoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *Invoice) GetStatus() InvoiceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*InvoiceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v InvoiceStatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


