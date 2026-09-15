# LlmProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**LlmProviderType**](LlmProviderType.md) |  | 
**HasCredential** | **bool** | Whether a credential is stored. The credential itself is never returned. | 
**Scope** | [**LlmProviderScope**](LlmProviderScope.md) |  | 
**OwnerUserSub** | Pointer to **NullableString** | Identity of the owning member. Null for an ORGANIZATION provider. | [optional] 
**OwnerName** | Pointer to **NullableString** | Display name of the owning member. Null for an ORGANIZATION provider. | [optional] 

## Methods

### NewLlmProviderResponse

`func NewLlmProviderResponse(id string, createdAt time.Time, updatedAt time.Time, name string, description string, type_ LlmProviderType, hasCredential bool, scope LlmProviderScope, ) *LlmProviderResponse`

NewLlmProviderResponse instantiates a new LlmProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmProviderResponseWithDefaults

`func NewLlmProviderResponseWithDefaults() *LlmProviderResponse`

NewLlmProviderResponseWithDefaults instantiates a new LlmProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *LlmProviderResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmProviderResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmProviderResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LlmProviderResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmProviderResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmProviderResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *LlmProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LlmProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LlmProviderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LlmProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LlmProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LlmProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *LlmProviderResponse) GetType() LlmProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LlmProviderResponse) GetTypeOk() (*LlmProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LlmProviderResponse) SetType(v LlmProviderType)`

SetType sets Type field to given value.


### GetHasCredential

`func (o *LlmProviderResponse) GetHasCredential() bool`

GetHasCredential returns the HasCredential field if non-nil, zero value otherwise.

### GetHasCredentialOk

`func (o *LlmProviderResponse) GetHasCredentialOk() (*bool, bool)`

GetHasCredentialOk returns a tuple with the HasCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCredential

`func (o *LlmProviderResponse) SetHasCredential(v bool)`

SetHasCredential sets HasCredential field to given value.


### GetScope

`func (o *LlmProviderResponse) GetScope() LlmProviderScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LlmProviderResponse) GetScopeOk() (*LlmProviderScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LlmProviderResponse) SetScope(v LlmProviderScope)`

SetScope sets Scope field to given value.


### GetOwnerUserSub

`func (o *LlmProviderResponse) GetOwnerUserSub() string`

GetOwnerUserSub returns the OwnerUserSub field if non-nil, zero value otherwise.

### GetOwnerUserSubOk

`func (o *LlmProviderResponse) GetOwnerUserSubOk() (*string, bool)`

GetOwnerUserSubOk returns a tuple with the OwnerUserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserSub

`func (o *LlmProviderResponse) SetOwnerUserSub(v string)`

SetOwnerUserSub sets OwnerUserSub field to given value.

### HasOwnerUserSub

`func (o *LlmProviderResponse) HasOwnerUserSub() bool`

HasOwnerUserSub returns a boolean if a field has been set.

### SetOwnerUserSubNil

`func (o *LlmProviderResponse) SetOwnerUserSubNil(b bool)`

 SetOwnerUserSubNil sets the value for OwnerUserSub to be an explicit nil

### UnsetOwnerUserSub
`func (o *LlmProviderResponse) UnsetOwnerUserSub()`

UnsetOwnerUserSub ensures that no value is present for OwnerUserSub, not even an explicit nil
### GetOwnerName

`func (o *LlmProviderResponse) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *LlmProviderResponse) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *LlmProviderResponse) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *LlmProviderResponse) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *LlmProviderResponse) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *LlmProviderResponse) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


