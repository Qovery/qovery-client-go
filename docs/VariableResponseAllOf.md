# VariableResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverriddenVariable** | Pointer to [**VariableOverride**](VariableOverride.md) |  | [optional] 
**AliasedVariable** | Pointer to [**VariableAlias**](VariableAlias.md) |  | [optional] 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableType** | Pointer to [**APIVariableTypeEnum**](APIVariableTypeEnum.md) |  | [optional] 
**ServiceId** | Pointer to **string** | The id of the service referenced by this variable. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service referenced by this variable. | [optional] 
**ServiceType** | Pointer to [**LinkedServiceTypeEnum**](LinkedServiceTypeEnum.md) |  | [optional] 
**OwnedBy** | Pointer to **string** | Entity that created/own the variable (i.e: Qovery, Doppler) | [optional] 

## Methods

### NewVariableResponseAllOf

`func NewVariableResponseAllOf(scope APIVariableScopeEnum, ) *VariableResponseAllOf`

NewVariableResponseAllOf instantiates a new VariableResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableResponseAllOfWithDefaults

`func NewVariableResponseAllOfWithDefaults() *VariableResponseAllOf`

NewVariableResponseAllOfWithDefaults instantiates a new VariableResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverriddenVariable

`func (o *VariableResponseAllOf) GetOverriddenVariable() VariableOverride`

GetOverriddenVariable returns the OverriddenVariable field if non-nil, zero value otherwise.

### GetOverriddenVariableOk

`func (o *VariableResponseAllOf) GetOverriddenVariableOk() (*VariableOverride, bool)`

GetOverriddenVariableOk returns a tuple with the OverriddenVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenVariable

`func (o *VariableResponseAllOf) SetOverriddenVariable(v VariableOverride)`

SetOverriddenVariable sets OverriddenVariable field to given value.

### HasOverriddenVariable

`func (o *VariableResponseAllOf) HasOverriddenVariable() bool`

HasOverriddenVariable returns a boolean if a field has been set.

### GetAliasedVariable

`func (o *VariableResponseAllOf) GetAliasedVariable() VariableAlias`

GetAliasedVariable returns the AliasedVariable field if non-nil, zero value otherwise.

### GetAliasedVariableOk

`func (o *VariableResponseAllOf) GetAliasedVariableOk() (*VariableAlias, bool)`

GetAliasedVariableOk returns a tuple with the AliasedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedVariable

`func (o *VariableResponseAllOf) SetAliasedVariable(v VariableAlias)`

SetAliasedVariable sets AliasedVariable field to given value.

### HasAliasedVariable

`func (o *VariableResponseAllOf) HasAliasedVariable() bool`

HasAliasedVariable returns a boolean if a field has been set.

### GetScope

`func (o *VariableResponseAllOf) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VariableResponseAllOf) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VariableResponseAllOf) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *VariableResponseAllOf) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *VariableResponseAllOf) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *VariableResponseAllOf) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.

### HasVariableType

`func (o *VariableResponseAllOf) HasVariableType() bool`

HasVariableType returns a boolean if a field has been set.

### GetServiceId

`func (o *VariableResponseAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *VariableResponseAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *VariableResponseAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *VariableResponseAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *VariableResponseAllOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *VariableResponseAllOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *VariableResponseAllOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *VariableResponseAllOf) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *VariableResponseAllOf) GetServiceType() LinkedServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VariableResponseAllOf) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VariableResponseAllOf) SetServiceType(v LinkedServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *VariableResponseAllOf) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetOwnedBy

`func (o *VariableResponseAllOf) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *VariableResponseAllOf) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *VariableResponseAllOf) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *VariableResponseAllOf) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


