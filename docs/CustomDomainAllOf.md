# CustomDomainAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationDomain** | Pointer to **string** | URL provided by Qovery. You must create a CNAME on your DNS provider using that URL | [optional] 
**Status** | Pointer to [**CustomDomainStatusEnum**](CustomDomainStatusEnum.md) |  | [optional] 

## Methods

### NewCustomDomainAllOf

`func NewCustomDomainAllOf() *CustomDomainAllOf`

NewCustomDomainAllOf instantiates a new CustomDomainAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainAllOfWithDefaults

`func NewCustomDomainAllOfWithDefaults() *CustomDomainAllOf`

NewCustomDomainAllOfWithDefaults instantiates a new CustomDomainAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationDomain

`func (o *CustomDomainAllOf) GetValidationDomain() string`

GetValidationDomain returns the ValidationDomain field if non-nil, zero value otherwise.

### GetValidationDomainOk

`func (o *CustomDomainAllOf) GetValidationDomainOk() (*string, bool)`

GetValidationDomainOk returns a tuple with the ValidationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDomain

`func (o *CustomDomainAllOf) SetValidationDomain(v string)`

SetValidationDomain sets ValidationDomain field to given value.

### HasValidationDomain

`func (o *CustomDomainAllOf) HasValidationDomain() bool`

HasValidationDomain returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDomainAllOf) GetStatus() CustomDomainStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomainAllOf) GetStatusOk() (*CustomDomainStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomainAllOf) SetStatus(v CustomDomainStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDomainAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


