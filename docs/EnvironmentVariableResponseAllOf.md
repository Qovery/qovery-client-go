# EnvironmentVariableResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverriddenVariable** | Pointer to [**EnvironmentVariableResponseAllOfOverriddenVariable**](EnvironmentVariableResponseAllOfOverriddenVariable.md) |  | [optional] 
**AliasedVariable** | Pointer to [**EnvironmentVariableResponseAllOfAliasedVariable**](EnvironmentVariableResponseAllOfAliasedVariable.md) |  | [optional] 
**Scope** | **string** |  | 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentVariableResponseAllOf

`func NewEnvironmentVariableResponseAllOf(scope string, ) *EnvironmentVariableResponseAllOf`

NewEnvironmentVariableResponseAllOf instantiates a new EnvironmentVariableResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableResponseAllOfWithDefaults

`func NewEnvironmentVariableResponseAllOfWithDefaults() *EnvironmentVariableResponseAllOf`

NewEnvironmentVariableResponseAllOfWithDefaults instantiates a new EnvironmentVariableResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverriddenVariable

`func (o *EnvironmentVariableResponseAllOf) GetOverriddenVariable() EnvironmentVariableResponseAllOfOverriddenVariable`

GetOverriddenVariable returns the OverriddenVariable field if non-nil, zero value otherwise.

### GetOverriddenVariableOk

`func (o *EnvironmentVariableResponseAllOf) GetOverriddenVariableOk() (*EnvironmentVariableResponseAllOfOverriddenVariable, bool)`

GetOverriddenVariableOk returns a tuple with the OverriddenVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenVariable

`func (o *EnvironmentVariableResponseAllOf) SetOverriddenVariable(v EnvironmentVariableResponseAllOfOverriddenVariable)`

SetOverriddenVariable sets OverriddenVariable field to given value.

### HasOverriddenVariable

`func (o *EnvironmentVariableResponseAllOf) HasOverriddenVariable() bool`

HasOverriddenVariable returns a boolean if a field has been set.

### GetAliasedVariable

`func (o *EnvironmentVariableResponseAllOf) GetAliasedVariable() EnvironmentVariableResponseAllOfAliasedVariable`

GetAliasedVariable returns the AliasedVariable field if non-nil, zero value otherwise.

### GetAliasedVariableOk

`func (o *EnvironmentVariableResponseAllOf) GetAliasedVariableOk() (*EnvironmentVariableResponseAllOfAliasedVariable, bool)`

GetAliasedVariableOk returns a tuple with the AliasedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedVariable

`func (o *EnvironmentVariableResponseAllOf) SetAliasedVariable(v EnvironmentVariableResponseAllOfAliasedVariable)`

SetAliasedVariable sets AliasedVariable field to given value.

### HasAliasedVariable

`func (o *EnvironmentVariableResponseAllOf) HasAliasedVariable() bool`

HasAliasedVariable returns a boolean if a field has been set.

### GetScope

`func (o *EnvironmentVariableResponseAllOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentVariableResponseAllOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentVariableResponseAllOf) SetScope(v string)`

SetScope sets Scope field to given value.


### GetServiceName

`func (o *EnvironmentVariableResponseAllOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *EnvironmentVariableResponseAllOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *EnvironmentVariableResponseAllOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *EnvironmentVariableResponseAllOf) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


