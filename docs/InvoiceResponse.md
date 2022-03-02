# InvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalInCents** | **int32** |  | 
**Total** | **float32** |  | 
**CurrencyCode** | **string** |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Status** | **string** |  | 

## Methods

### NewInvoiceResponse

`func NewInvoiceResponse(totalInCents int32, total float32, currencyCode string, id string, createdAt time.Time, status string, ) *InvoiceResponse`

NewInvoiceResponse instantiates a new InvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseWithDefaults

`func NewInvoiceResponseWithDefaults() *InvoiceResponse`

NewInvoiceResponseWithDefaults instantiates a new InvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInCents

`func (o *InvoiceResponse) GetTotalInCents() int32`

GetTotalInCents returns the TotalInCents field if non-nil, zero value otherwise.

### GetTotalInCentsOk

`func (o *InvoiceResponse) GetTotalInCentsOk() (*int32, bool)`

GetTotalInCentsOk returns a tuple with the TotalInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInCents

`func (o *InvoiceResponse) SetTotalInCents(v int32)`

SetTotalInCents sets TotalInCents field to given value.


### GetTotal

`func (o *InvoiceResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetCurrencyCode

`func (o *InvoiceResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetId

`func (o *InvoiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *InvoiceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *InvoiceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


