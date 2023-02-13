# EnvironmentVariableOverride

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

### NewEnvironmentVariableOverride

`func NewEnvironmentVariableOverride(id string, key string, value string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum, ) *EnvironmentVariableOverride`

NewEnvironmentVariableOverride instantiates a new EnvironmentVariableOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableOverrideWithDefaults

`func NewEnvironmentVariableOverrideWithDefaults() *EnvironmentVariableOverride`

NewEnvironmentVariableOverrideWithDefaults instantiates a new EnvironmentVariableOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentVariableOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentVariableOverride) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentVariableOverride) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *EnvironmentVariableOverride) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentVariableOverride) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentVariableOverride) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EnvironmentVariableOverride) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariableOverride) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariableOverride) SetValue(v string)`

SetValue sets Value field to given value.


### GetMountPath

`func (o *EnvironmentVariableOverride) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *EnvironmentVariableOverride) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *EnvironmentVariableOverride) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetScope

`func (o *EnvironmentVariableOverride) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentVariableOverride) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentVariableOverride) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *EnvironmentVariableOverride) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *EnvironmentVariableOverride) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *EnvironmentVariableOverride) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


