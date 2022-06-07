# VariableImportRequestVarsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Scope** | [**EnvironmentVariableScopeEnum**](EnvironmentVariableScopeEnum.md) |  | 
**IsSecret** | **bool** |  | 

## Methods

### NewVariableImportRequestVarsInner

`func NewVariableImportRequestVarsInner(name string, value string, scope EnvironmentVariableScopeEnum, isSecret bool, ) *VariableImportRequestVarsInner`

NewVariableImportRequestVarsInner instantiates a new VariableImportRequestVarsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableImportRequestVarsInnerWithDefaults

`func NewVariableImportRequestVarsInnerWithDefaults() *VariableImportRequestVarsInner`

NewVariableImportRequestVarsInnerWithDefaults instantiates a new VariableImportRequestVarsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariableImportRequestVarsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableImportRequestVarsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableImportRequestVarsInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *VariableImportRequestVarsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableImportRequestVarsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableImportRequestVarsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetScope

`func (o *VariableImportRequestVarsInner) GetScope() EnvironmentVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VariableImportRequestVarsInner) GetScopeOk() (*EnvironmentVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VariableImportRequestVarsInner) SetScope(v EnvironmentVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetIsSecret

`func (o *VariableImportRequestVarsInner) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *VariableImportRequestVarsInner) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *VariableImportRequestVarsInner) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


