# SecretResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | key is case sensitive | [optional] 
**OverriddenSecret** | Pointer to [**OverriddenSecret**](OverriddenSecret.md) |  | [optional] 
**AliasedSecret** | Pointer to [**AliasedSecret**](AliasedSecret.md) |  | [optional] 
**Scope** | [**EnvironmentVariableScopeEnum**](EnvironmentVariableScopeEnum.md) |  | 

## Methods

### NewSecretResponseAllOf

`func NewSecretResponseAllOf(scope EnvironmentVariableScopeEnum, ) *SecretResponseAllOf`

NewSecretResponseAllOf instantiates a new SecretResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseAllOfWithDefaults

`func NewSecretResponseAllOfWithDefaults() *SecretResponseAllOf`

NewSecretResponseAllOfWithDefaults instantiates a new SecretResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SecretResponseAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretResponseAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretResponseAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SecretResponseAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOverriddenSecret

`func (o *SecretResponseAllOf) GetOverriddenSecret() OverriddenSecret`

GetOverriddenSecret returns the OverriddenSecret field if non-nil, zero value otherwise.

### GetOverriddenSecretOk

`func (o *SecretResponseAllOf) GetOverriddenSecretOk() (*OverriddenSecret, bool)`

GetOverriddenSecretOk returns a tuple with the OverriddenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenSecret

`func (o *SecretResponseAllOf) SetOverriddenSecret(v OverriddenSecret)`

SetOverriddenSecret sets OverriddenSecret field to given value.

### HasOverriddenSecret

`func (o *SecretResponseAllOf) HasOverriddenSecret() bool`

HasOverriddenSecret returns a boolean if a field has been set.

### GetAliasedSecret

`func (o *SecretResponseAllOf) GetAliasedSecret() AliasedSecret`

GetAliasedSecret returns the AliasedSecret field if non-nil, zero value otherwise.

### GetAliasedSecretOk

`func (o *SecretResponseAllOf) GetAliasedSecretOk() (*AliasedSecret, bool)`

GetAliasedSecretOk returns a tuple with the AliasedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedSecret

`func (o *SecretResponseAllOf) SetAliasedSecret(v AliasedSecret)`

SetAliasedSecret sets AliasedSecret field to given value.

### HasAliasedSecret

`func (o *SecretResponseAllOf) HasAliasedSecret() bool`

HasAliasedSecret returns a boolean if a field has been set.

### GetScope

`func (o *SecretResponseAllOf) GetScope() EnvironmentVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecretResponseAllOf) GetScopeOk() (*EnvironmentVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecretResponseAllOf) SetScope(v EnvironmentVariableScopeEnum)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


