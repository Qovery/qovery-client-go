# InvoiceResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Invoice**](Invoice.md) |  | [optional] 

## Methods

### NewInvoiceResponseList

`func NewInvoiceResponseList() *InvoiceResponseList`

NewInvoiceResponseList instantiates a new InvoiceResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseListWithDefaults

`func NewInvoiceResponseListWithDefaults() *InvoiceResponseList`

NewInvoiceResponseListWithDefaults instantiates a new InvoiceResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *InvoiceResponseList) GetResults() []Invoice`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InvoiceResponseList) GetResultsOk() (*[]Invoice, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InvoiceResponseList) SetResults(v []Invoice)`

SetResults sets Results field to given value.

### HasResults

`func (o *InvoiceResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


