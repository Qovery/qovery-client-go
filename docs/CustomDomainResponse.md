# CustomDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationDomain** | Pointer to **string** | URL provided by Qovery. You must create a CNAME on your DNS provider using that URL | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Domain** | **string** | your custom domain | 

## Methods

### NewCustomDomainResponse

`func NewCustomDomainResponse(id string, createdAt time.Time, domain string, ) *CustomDomainResponse`

NewCustomDomainResponse instantiates a new CustomDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainResponseWithDefaults

`func NewCustomDomainResponseWithDefaults() *CustomDomainResponse`

NewCustomDomainResponseWithDefaults instantiates a new CustomDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationDomain

`func (o *CustomDomainResponse) GetValidationDomain() string`

GetValidationDomain returns the ValidationDomain field if non-nil, zero value otherwise.

### GetValidationDomainOk

`func (o *CustomDomainResponse) GetValidationDomainOk() (*string, bool)`

GetValidationDomainOk returns a tuple with the ValidationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDomain

`func (o *CustomDomainResponse) SetValidationDomain(v string)`

SetValidationDomain sets ValidationDomain field to given value.

### HasValidationDomain

`func (o *CustomDomainResponse) HasValidationDomain() bool`

HasValidationDomain returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDomainResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomainResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomainResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDomainResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *CustomDomainResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CustomDomainResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomDomainResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomDomainResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomDomainResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomDomainResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomDomainResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomDomainResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDomain

`func (o *CustomDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustomDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustomDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


