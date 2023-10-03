# CustomDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Domain** | **string** | your custom domain | 
**GenerateCertificate** | Pointer to **bool** | to control if a certificate has to be generated for this custom domain by Qovery. The default value is &#x60;true&#x60;. This flag should be set to &#x60;false&#x60; if a CDN or other entities are managing the certificate for the specified domain and the traffic is proxied by the CDN to Qovery. | [optional] 
**ValidationDomain** | Pointer to **string** | URL provided by Qovery. You must create a CNAME on your DNS provider using that URL | [optional] 
**Status** | Pointer to [**CustomDomainStatusEnum**](CustomDomainStatusEnum.md) |  | [optional] 

## Methods

### NewCustomDomain

`func NewCustomDomain(id string, createdAt time.Time, domain string, ) *CustomDomain`

NewCustomDomain instantiates a new CustomDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainWithDefaults

`func NewCustomDomainWithDefaults() *CustomDomain`

NewCustomDomainWithDefaults instantiates a new CustomDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomain) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CustomDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomDomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomDomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomDomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomDomain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDomain

`func (o *CustomDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustomDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustomDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGenerateCertificate

`func (o *CustomDomain) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *CustomDomain) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *CustomDomain) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.

### HasGenerateCertificate

`func (o *CustomDomain) HasGenerateCertificate() bool`

HasGenerateCertificate returns a boolean if a field has been set.

### GetValidationDomain

`func (o *CustomDomain) GetValidationDomain() string`

GetValidationDomain returns the ValidationDomain field if non-nil, zero value otherwise.

### GetValidationDomainOk

`func (o *CustomDomain) GetValidationDomainOk() (*string, bool)`

GetValidationDomainOk returns a tuple with the ValidationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDomain

`func (o *CustomDomain) SetValidationDomain(v string)`

SetValidationDomain sets ValidationDomain field to given value.

### HasValidationDomain

`func (o *CustomDomain) HasValidationDomain() bool`

HasValidationDomain returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDomain) GetStatus() CustomDomainStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomain) GetStatusOk() (*CustomDomainStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomain) SetStatus(v CustomDomainStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


