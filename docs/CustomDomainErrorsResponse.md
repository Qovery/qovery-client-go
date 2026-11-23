# CustomDomainErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]CustomDomainErrorResponse**](CustomDomainErrorResponse.md) | Every conflict of the request. Empty when the domain can be saved. | 

## Methods

### NewCustomDomainErrorsResponse

`func NewCustomDomainErrorsResponse(errors []CustomDomainErrorResponse, ) *CustomDomainErrorsResponse`

NewCustomDomainErrorsResponse instantiates a new CustomDomainErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainErrorsResponseWithDefaults

`func NewCustomDomainErrorsResponseWithDefaults() *CustomDomainErrorsResponse`

NewCustomDomainErrorsResponseWithDefaults instantiates a new CustomDomainErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CustomDomainErrorsResponse) GetErrors() []CustomDomainErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CustomDomainErrorsResponse) GetErrorsOk() (*[]CustomDomainErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CustomDomainErrorsResponse) SetErrors(v []CustomDomainErrorResponse)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


