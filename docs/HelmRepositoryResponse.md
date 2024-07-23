# HelmRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Kind** | Pointer to [**HelmRepositoryKindEnum**](HelmRepositoryKindEnum.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the helm repository | [optional] 
**SkipTlsVerification** | Pointer to **bool** | Bypass tls certificate verification when connecting to repository | [optional] 
**AssociatedServicesCount** | Pointer to **int32** | The number of services using this helm repository | [optional] 

## Methods

### NewHelmRepositoryResponse

`func NewHelmRepositoryResponse(id string, createdAt time.Time, name string, ) *HelmRepositoryResponse`

NewHelmRepositoryResponse instantiates a new HelmRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepositoryResponseWithDefaults

`func NewHelmRepositoryResponseWithDefaults() *HelmRepositoryResponse`

NewHelmRepositoryResponseWithDefaults instantiates a new HelmRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HelmRepositoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelmRepositoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelmRepositoryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HelmRepositoryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HelmRepositoryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HelmRepositoryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HelmRepositoryResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HelmRepositoryResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HelmRepositoryResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HelmRepositoryResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *HelmRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *HelmRepositoryResponse) GetKind() HelmRepositoryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HelmRepositoryResponse) GetKindOk() (*HelmRepositoryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HelmRepositoryResponse) SetKind(v HelmRepositoryKindEnum)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HelmRepositoryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDescription

`func (o *HelmRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelmRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelmRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelmRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *HelmRepositoryResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HelmRepositoryResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HelmRepositoryResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HelmRepositoryResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSkipTlsVerification

`func (o *HelmRepositoryResponse) GetSkipTlsVerification() bool`

GetSkipTlsVerification returns the SkipTlsVerification field if non-nil, zero value otherwise.

### GetSkipTlsVerificationOk

`func (o *HelmRepositoryResponse) GetSkipTlsVerificationOk() (*bool, bool)`

GetSkipTlsVerificationOk returns a tuple with the SkipTlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTlsVerification

`func (o *HelmRepositoryResponse) SetSkipTlsVerification(v bool)`

SetSkipTlsVerification sets SkipTlsVerification field to given value.

### HasSkipTlsVerification

`func (o *HelmRepositoryResponse) HasSkipTlsVerification() bool`

HasSkipTlsVerification returns a boolean if a field has been set.

### GetAssociatedServicesCount

`func (o *HelmRepositoryResponse) GetAssociatedServicesCount() int32`

GetAssociatedServicesCount returns the AssociatedServicesCount field if non-nil, zero value otherwise.

### GetAssociatedServicesCountOk

`func (o *HelmRepositoryResponse) GetAssociatedServicesCountOk() (*int32, bool)`

GetAssociatedServicesCountOk returns a tuple with the AssociatedServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServicesCount

`func (o *HelmRepositoryResponse) SetAssociatedServicesCount(v int32)`

SetAssociatedServicesCount sets AssociatedServicesCount field to given value.

### HasAssociatedServicesCount

`func (o *HelmRepositoryResponse) HasAssociatedServicesCount() bool`

HasAssociatedServicesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


