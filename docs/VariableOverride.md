# VariableOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the overriden variable | 
**Key** | **string** | The key of the overriden variable | 
**Value** | Pointer to **NullableString** | The value of the overriden variable | [optional] 
**MountPath** | **string** | The mounth path of the overriden variable (only if environment variable type is &#39;file&#39;) | 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableType** | [**APIVariableTypeEnum**](APIVariableTypeEnum.md) |  | 

## Methods

### NewVariableOverride

`func NewVariableOverride(id string, key string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum, ) *VariableOverride`

NewVariableOverride instantiates a new VariableOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableOverrideWithDefaults

`func NewVariableOverrideWithDefaults() *VariableOverride`

NewVariableOverrideWithDefaults instantiates a new VariableOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariableOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableOverride) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableOverride) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *VariableOverride) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableOverride) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableOverride) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableOverride) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableOverride) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableOverride) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableOverride) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableOverride) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableOverride) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMountPath

`func (o *VariableOverride) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *VariableOverride) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *VariableOverride) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetScope

`func (o *VariableOverride) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VariableOverride) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VariableOverride) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *VariableOverride) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *VariableOverride) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *VariableOverride) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


