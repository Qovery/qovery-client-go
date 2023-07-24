# VariableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | **string** |  | 
**MountPath** | Pointer to **NullableString** |  | [optional] 
**IsSecret** | **bool** |  | 
**VariableScope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableParentId** | **string** |  | 

## Methods

### NewVariableRequest

`func NewVariableRequest(key string, value string, isSecret bool, variableScope APIVariableScopeEnum, variableParentId string, ) *VariableRequest`

NewVariableRequest instantiates a new VariableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableRequestWithDefaults

`func NewVariableRequestWithDefaults() *VariableRequest`

NewVariableRequestWithDefaults instantiates a new VariableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetMountPath

`func (o *VariableRequest) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *VariableRequest) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *VariableRequest) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *VariableRequest) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *VariableRequest) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *VariableRequest) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetIsSecret

`func (o *VariableRequest) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *VariableRequest) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *VariableRequest) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetVariableScope

`func (o *VariableRequest) GetVariableScope() APIVariableScopeEnum`

GetVariableScope returns the VariableScope field if non-nil, zero value otherwise.

### GetVariableScopeOk

`func (o *VariableRequest) GetVariableScopeOk() (*APIVariableScopeEnum, bool)`

GetVariableScopeOk returns a tuple with the VariableScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableScope

`func (o *VariableRequest) SetVariableScope(v APIVariableScopeEnum)`

SetVariableScope sets VariableScope field to given value.


### GetVariableParentId

`func (o *VariableRequest) GetVariableParentId() string`

GetVariableParentId returns the VariableParentId field if non-nil, zero value otherwise.

### GetVariableParentIdOk

`func (o *VariableRequest) GetVariableParentIdOk() (*string, bool)`

GetVariableParentIdOk returns a tuple with the VariableParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableParentId

`func (o *VariableRequest) SetVariableParentId(v string)`

SetVariableParentId sets VariableParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


