# CheckedCustomDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | **string** | domain name checked | 
**Status** | [**CheckedCustomDomainStatus**](CheckedCustomDomainStatus.md) |  | 
**ErrorDetails** | Pointer to **string** | optional field containing information about failure check | [optional] 

## Methods

### NewCheckedCustomDomainResponse

`func NewCheckedCustomDomainResponse(domainName string, status CheckedCustomDomainStatus, ) *CheckedCustomDomainResponse`

NewCheckedCustomDomainResponse instantiates a new CheckedCustomDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckedCustomDomainResponseWithDefaults

`func NewCheckedCustomDomainResponseWithDefaults() *CheckedCustomDomainResponse`

NewCheckedCustomDomainResponseWithDefaults instantiates a new CheckedCustomDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *CheckedCustomDomainResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CheckedCustomDomainResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CheckedCustomDomainResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetStatus

`func (o *CheckedCustomDomainResponse) GetStatus() CheckedCustomDomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckedCustomDomainResponse) GetStatusOk() (*CheckedCustomDomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckedCustomDomainResponse) SetStatus(v CheckedCustomDomainStatus)`

SetStatus sets Status field to given value.


### GetErrorDetails

`func (o *CheckedCustomDomainResponse) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *CheckedCustomDomainResponse) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *CheckedCustomDomainResponse) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *CheckedCustomDomainResponse) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


