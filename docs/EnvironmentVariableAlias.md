# EnvironmentVariableAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Value** | **string** |  | 
**MountPath** | **string** |  | 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableType** | [**APIVariableTypeEnum**](APIVariableTypeEnum.md) |  | 

## Methods

### NewEnvironmentVariableAlias

`func NewEnvironmentVariableAlias(id string, key string, value string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum, ) *EnvironmentVariableAlias`

NewEnvironmentVariableAlias instantiates a new EnvironmentVariableAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableAliasWithDefaults

`func NewEnvironmentVariableAliasWithDefaults() *EnvironmentVariableAlias`

NewEnvironmentVariableAliasWithDefaults instantiates a new EnvironmentVariableAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentVariableAlias) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentVariableAlias) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentVariableAlias) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *EnvironmentVariableAlias) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentVariableAlias) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentVariableAlias) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EnvironmentVariableAlias) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariableAlias) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariableAlias) SetValue(v string)`

SetValue sets Value field to given value.


### GetMountPath

`func (o *EnvironmentVariableAlias) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *EnvironmentVariableAlias) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *EnvironmentVariableAlias) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetScope

`func (o *EnvironmentVariableAlias) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentVariableAlias) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentVariableAlias) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *EnvironmentVariableAlias) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *EnvironmentVariableAlias) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *EnvironmentVariableAlias) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


