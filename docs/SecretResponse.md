# SecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Key** | Pointer to **string** | key is case sensitive | [optional] 
**OverriddenSecret** | Pointer to [**OverriddenSecret**](OverriddenSecret.md) |  | [optional] 
**AliasedSecret** | Pointer to [**AliasedSecret**](AliasedSecret.md) |  | [optional] 
**Scope** | [**EnvironmentVariableScopeEnum**](EnvironmentVariableScopeEnum.md) |  | 

## Methods

### NewSecretResponse

`func NewSecretResponse(id string, createdAt time.Time, scope EnvironmentVariableScopeEnum, ) *SecretResponse`

NewSecretResponse instantiates a new SecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseWithDefaults

`func NewSecretResponseWithDefaults() *SecretResponse`

NewSecretResponseWithDefaults instantiates a new SecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SecretResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecretResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecretResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SecretResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecretResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecretResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SecretResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetKey

`func (o *SecretResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SecretResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOverriddenSecret

`func (o *SecretResponse) GetOverriddenSecret() OverriddenSecret`

GetOverriddenSecret returns the OverriddenSecret field if non-nil, zero value otherwise.

### GetOverriddenSecretOk

`func (o *SecretResponse) GetOverriddenSecretOk() (*OverriddenSecret, bool)`

GetOverriddenSecretOk returns a tuple with the OverriddenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenSecret

`func (o *SecretResponse) SetOverriddenSecret(v OverriddenSecret)`

SetOverriddenSecret sets OverriddenSecret field to given value.

### HasOverriddenSecret

`func (o *SecretResponse) HasOverriddenSecret() bool`

HasOverriddenSecret returns a boolean if a field has been set.

### GetAliasedSecret

`func (o *SecretResponse) GetAliasedSecret() AliasedSecret`

GetAliasedSecret returns the AliasedSecret field if non-nil, zero value otherwise.

### GetAliasedSecretOk

`func (o *SecretResponse) GetAliasedSecretOk() (*AliasedSecret, bool)`

GetAliasedSecretOk returns a tuple with the AliasedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedSecret

`func (o *SecretResponse) SetAliasedSecret(v AliasedSecret)`

SetAliasedSecret sets AliasedSecret field to given value.

### HasAliasedSecret

`func (o *SecretResponse) HasAliasedSecret() bool`

HasAliasedSecret returns a boolean if a field has been set.

### GetScope

`func (o *SecretResponse) GetScope() EnvironmentVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecretResponse) GetScopeOk() (*EnvironmentVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecretResponse) SetScope(v EnvironmentVariableScopeEnum)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


