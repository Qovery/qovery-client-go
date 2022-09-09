# VariableImportSuccessfulImportedVariablesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **string** | Optional if the variable is secret | [optional] 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**IsSecret** | **bool** |  | 

## Methods

### NewVariableImportSuccessfulImportedVariablesInner

`func NewVariableImportSuccessfulImportedVariablesInner(name string, scope APIVariableScopeEnum, isSecret bool, ) *VariableImportSuccessfulImportedVariablesInner`

NewVariableImportSuccessfulImportedVariablesInner instantiates a new VariableImportSuccessfulImportedVariablesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableImportSuccessfulImportedVariablesInnerWithDefaults

`func NewVariableImportSuccessfulImportedVariablesInnerWithDefaults() *VariableImportSuccessfulImportedVariablesInner`

NewVariableImportSuccessfulImportedVariablesInnerWithDefaults instantiates a new VariableImportSuccessfulImportedVariablesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariableImportSuccessfulImportedVariablesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableImportSuccessfulImportedVariablesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableImportSuccessfulImportedVariablesInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *VariableImportSuccessfulImportedVariablesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableImportSuccessfulImportedVariablesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableImportSuccessfulImportedVariablesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableImportSuccessfulImportedVariablesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetScope

`func (o *VariableImportSuccessfulImportedVariablesInner) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VariableImportSuccessfulImportedVariablesInner) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VariableImportSuccessfulImportedVariablesInner) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetIsSecret

`func (o *VariableImportSuccessfulImportedVariablesInner) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *VariableImportSuccessfulImportedVariablesInner) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *VariableImportSuccessfulImportedVariablesInner) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


