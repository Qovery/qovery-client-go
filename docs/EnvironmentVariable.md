# EnvironmentVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Key** | **string** | key is case sensitive. | 
**Value** | **string** | value of the env variable. | 
**MountPath** | Pointer to **NullableString** | should be set for file only. variable mount path makes variable a file (where file should be mounted). | [optional] 
**OverriddenVariable** | Pointer to [**EnvironmentVariableOverride**](EnvironmentVariableOverride.md) |  | [optional] 
**AliasedVariable** | Pointer to [**EnvironmentVariableAlias**](EnvironmentVariableAlias.md) |  | [optional] 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableType** | Pointer to [**APIVariableTypeEnum**](APIVariableTypeEnum.md) |  | [optional] 
**ServiceId** | Pointer to **string** | present only for &#x60;BUILT_IN&#x60; variable | [optional] 
**ServiceName** | Pointer to **string** | present only for &#x60;BUILT_IN&#x60; variable | [optional] 
**ServiceType** | Pointer to [**LinkedServiceTypeEnum**](LinkedServiceTypeEnum.md) |  | [optional] 

## Methods

### NewEnvironmentVariable

`func NewEnvironmentVariable(id string, createdAt time.Time, key string, value string, scope APIVariableScopeEnum, ) *EnvironmentVariable`

NewEnvironmentVariable instantiates a new EnvironmentVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableWithDefaults

`func NewEnvironmentVariableWithDefaults() *EnvironmentVariable`

NewEnvironmentVariableWithDefaults instantiates a new EnvironmentVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentVariable) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentVariable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentVariable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentVariable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentVariable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentVariable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentVariable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentVariable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetKey

`func (o *EnvironmentVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentVariable) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EnvironmentVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariable) SetValue(v string)`

SetValue sets Value field to given value.


### GetMountPath

`func (o *EnvironmentVariable) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *EnvironmentVariable) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *EnvironmentVariable) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *EnvironmentVariable) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *EnvironmentVariable) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *EnvironmentVariable) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetOverriddenVariable

`func (o *EnvironmentVariable) GetOverriddenVariable() EnvironmentVariableOverride`

GetOverriddenVariable returns the OverriddenVariable field if non-nil, zero value otherwise.

### GetOverriddenVariableOk

`func (o *EnvironmentVariable) GetOverriddenVariableOk() (*EnvironmentVariableOverride, bool)`

GetOverriddenVariableOk returns a tuple with the OverriddenVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenVariable

`func (o *EnvironmentVariable) SetOverriddenVariable(v EnvironmentVariableOverride)`

SetOverriddenVariable sets OverriddenVariable field to given value.

### HasOverriddenVariable

`func (o *EnvironmentVariable) HasOverriddenVariable() bool`

HasOverriddenVariable returns a boolean if a field has been set.

### GetAliasedVariable

`func (o *EnvironmentVariable) GetAliasedVariable() EnvironmentVariableAlias`

GetAliasedVariable returns the AliasedVariable field if non-nil, zero value otherwise.

### GetAliasedVariableOk

`func (o *EnvironmentVariable) GetAliasedVariableOk() (*EnvironmentVariableAlias, bool)`

GetAliasedVariableOk returns a tuple with the AliasedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedVariable

`func (o *EnvironmentVariable) SetAliasedVariable(v EnvironmentVariableAlias)`

SetAliasedVariable sets AliasedVariable field to given value.

### HasAliasedVariable

`func (o *EnvironmentVariable) HasAliasedVariable() bool`

HasAliasedVariable returns a boolean if a field has been set.

### GetScope

`func (o *EnvironmentVariable) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentVariable) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentVariable) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *EnvironmentVariable) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *EnvironmentVariable) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *EnvironmentVariable) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.

### HasVariableType

`func (o *EnvironmentVariable) HasVariableType() bool`

HasVariableType returns a boolean if a field has been set.

### GetServiceId

`func (o *EnvironmentVariable) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *EnvironmentVariable) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *EnvironmentVariable) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *EnvironmentVariable) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *EnvironmentVariable) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *EnvironmentVariable) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *EnvironmentVariable) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *EnvironmentVariable) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *EnvironmentVariable) GetServiceType() LinkedServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *EnvironmentVariable) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *EnvironmentVariable) SetServiceType(v LinkedServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *EnvironmentVariable) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


