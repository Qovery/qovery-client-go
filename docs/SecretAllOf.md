# SecretAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | key is case sensitive | 
**OverriddenSecret** | Pointer to [**SecretOverride**](SecretOverride.md) |  | [optional] 
**AliasedSecret** | Pointer to [**SecretAlias**](SecretAlias.md) |  | [optional] 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**ServiceId** | Pointer to **string** | present only for &#x60;BUILT_IN&#x60; variable | [optional] 
**ServiceName** | Pointer to **string** | present only for &#x60;BUILT_IN&#x60; variable | [optional] 
**ServiceType** | Pointer to [**LinkedServiceTypeEnum**](LinkedServiceTypeEnum.md) |  | [optional] 

## Methods

### NewSecretAllOf

`func NewSecretAllOf(key string, scope APIVariableScopeEnum, ) *SecretAllOf`

NewSecretAllOf instantiates a new SecretAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretAllOfWithDefaults

`func NewSecretAllOfWithDefaults() *SecretAllOf`

NewSecretAllOfWithDefaults instantiates a new SecretAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SecretAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretAllOf) SetKey(v string)`

SetKey sets Key field to given value.


### GetOverriddenSecret

`func (o *SecretAllOf) GetOverriddenSecret() SecretOverride`

GetOverriddenSecret returns the OverriddenSecret field if non-nil, zero value otherwise.

### GetOverriddenSecretOk

`func (o *SecretAllOf) GetOverriddenSecretOk() (*SecretOverride, bool)`

GetOverriddenSecretOk returns a tuple with the OverriddenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenSecret

`func (o *SecretAllOf) SetOverriddenSecret(v SecretOverride)`

SetOverriddenSecret sets OverriddenSecret field to given value.

### HasOverriddenSecret

`func (o *SecretAllOf) HasOverriddenSecret() bool`

HasOverriddenSecret returns a boolean if a field has been set.

### GetAliasedSecret

`func (o *SecretAllOf) GetAliasedSecret() SecretAlias`

GetAliasedSecret returns the AliasedSecret field if non-nil, zero value otherwise.

### GetAliasedSecretOk

`func (o *SecretAllOf) GetAliasedSecretOk() (*SecretAlias, bool)`

GetAliasedSecretOk returns a tuple with the AliasedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedSecret

`func (o *SecretAllOf) SetAliasedSecret(v SecretAlias)`

SetAliasedSecret sets AliasedSecret field to given value.

### HasAliasedSecret

`func (o *SecretAllOf) HasAliasedSecret() bool`

HasAliasedSecret returns a boolean if a field has been set.

### GetScope

`func (o *SecretAllOf) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecretAllOf) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecretAllOf) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetServiceId

`func (o *SecretAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SecretAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SecretAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SecretAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *SecretAllOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SecretAllOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SecretAllOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *SecretAllOf) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *SecretAllOf) GetServiceType() LinkedServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *SecretAllOf) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *SecretAllOf) SetServiceType(v LinkedServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *SecretAllOf) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


