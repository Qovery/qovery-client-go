# CustomDomainResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationDomain** | Pointer to **string** | URL provided by Qovery. You must create a CNAME on your DNS provider using that URL | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomDomainResponseAllOf

`func NewCustomDomainResponseAllOf() *CustomDomainResponseAllOf`

NewCustomDomainResponseAllOf instantiates a new CustomDomainResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainResponseAllOfWithDefaults

`func NewCustomDomainResponseAllOfWithDefaults() *CustomDomainResponseAllOf`

NewCustomDomainResponseAllOfWithDefaults instantiates a new CustomDomainResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationDomain

`func (o *CustomDomainResponseAllOf) GetValidationDomain() string`

GetValidationDomain returns the ValidationDomain field if non-nil, zero value otherwise.

### GetValidationDomainOk

`func (o *CustomDomainResponseAllOf) GetValidationDomainOk() (*string, bool)`

GetValidationDomainOk returns a tuple with the ValidationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDomain

`func (o *CustomDomainResponseAllOf) SetValidationDomain(v string)`

SetValidationDomain sets ValidationDomain field to given value.

### HasValidationDomain

`func (o *CustomDomainResponseAllOf) HasValidationDomain() bool`

HasValidationDomain returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDomainResponseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomainResponseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomainResponseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDomainResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


