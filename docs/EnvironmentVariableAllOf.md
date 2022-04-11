# EnvironmentVariableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverriddenVariable** | Pointer to [**EnvironmentVariableAllOfOverriddenVariable**](EnvironmentVariableAllOfOverriddenVariable.md) |  | [optional] 
**AliasedVariable** | Pointer to [**EnvironmentVariableAllOfAliasedVariable**](EnvironmentVariableAllOfAliasedVariable.md) |  | [optional] 
**Scope** | [**EnvironmentVariableScopeEnum**](EnvironmentVariableScopeEnum.md) |  | 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentVariableAllOf

`func NewEnvironmentVariableAllOf(scope EnvironmentVariableScopeEnum, ) *EnvironmentVariableAllOf`

NewEnvironmentVariableAllOf instantiates a new EnvironmentVariableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableAllOfWithDefaults

`func NewEnvironmentVariableAllOfWithDefaults() *EnvironmentVariableAllOf`

NewEnvironmentVariableAllOfWithDefaults instantiates a new EnvironmentVariableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverriddenVariable

`func (o *EnvironmentVariableAllOf) GetOverriddenVariable() EnvironmentVariableAllOfOverriddenVariable`

GetOverriddenVariable returns the OverriddenVariable field if non-nil, zero value otherwise.

### GetOverriddenVariableOk

`func (o *EnvironmentVariableAllOf) GetOverriddenVariableOk() (*EnvironmentVariableAllOfOverriddenVariable, bool)`

GetOverriddenVariableOk returns a tuple with the OverriddenVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenVariable

`func (o *EnvironmentVariableAllOf) SetOverriddenVariable(v EnvironmentVariableAllOfOverriddenVariable)`

SetOverriddenVariable sets OverriddenVariable field to given value.

### HasOverriddenVariable

`func (o *EnvironmentVariableAllOf) HasOverriddenVariable() bool`

HasOverriddenVariable returns a boolean if a field has been set.

### GetAliasedVariable

`func (o *EnvironmentVariableAllOf) GetAliasedVariable() EnvironmentVariableAllOfAliasedVariable`

GetAliasedVariable returns the AliasedVariable field if non-nil, zero value otherwise.

### GetAliasedVariableOk

`func (o *EnvironmentVariableAllOf) GetAliasedVariableOk() (*EnvironmentVariableAllOfAliasedVariable, bool)`

GetAliasedVariableOk returns a tuple with the AliasedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedVariable

`func (o *EnvironmentVariableAllOf) SetAliasedVariable(v EnvironmentVariableAllOfAliasedVariable)`

SetAliasedVariable sets AliasedVariable field to given value.

### HasAliasedVariable

`func (o *EnvironmentVariableAllOf) HasAliasedVariable() bool`

HasAliasedVariable returns a boolean if a field has been set.

### GetScope

`func (o *EnvironmentVariableAllOf) GetScope() EnvironmentVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentVariableAllOf) GetScopeOk() (*EnvironmentVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentVariableAllOf) SetScope(v EnvironmentVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetServiceName

`func (o *EnvironmentVariableAllOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *EnvironmentVariableAllOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *EnvironmentVariableAllOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *EnvironmentVariableAllOf) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


