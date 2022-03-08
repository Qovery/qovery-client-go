# VariableImportResponseSuccessfulImportedVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **string** | Optional if the variable is secret | [optional] 
**Scope** | [**EnvironmentVariableScopeEnum**](EnvironmentVariableScopeEnum.md) |  | 
**IsSecret** | **bool** |  | 

## Methods

### NewVariableImportResponseSuccessfulImportedVariables

`func NewVariableImportResponseSuccessfulImportedVariables(name string, scope EnvironmentVariableScopeEnum, isSecret bool, ) *VariableImportResponseSuccessfulImportedVariables`

NewVariableImportResponseSuccessfulImportedVariables instantiates a new VariableImportResponseSuccessfulImportedVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableImportResponseSuccessfulImportedVariablesWithDefaults

`func NewVariableImportResponseSuccessfulImportedVariablesWithDefaults() *VariableImportResponseSuccessfulImportedVariables`

NewVariableImportResponseSuccessfulImportedVariablesWithDefaults instantiates a new VariableImportResponseSuccessfulImportedVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariableImportResponseSuccessfulImportedVariables) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableImportResponseSuccessfulImportedVariables) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableImportResponseSuccessfulImportedVariables) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *VariableImportResponseSuccessfulImportedVariables) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableImportResponseSuccessfulImportedVariables) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableImportResponseSuccessfulImportedVariables) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableImportResponseSuccessfulImportedVariables) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetScope

`func (o *VariableImportResponseSuccessfulImportedVariables) GetScope() EnvironmentVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VariableImportResponseSuccessfulImportedVariables) GetScopeOk() (*EnvironmentVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VariableImportResponseSuccessfulImportedVariables) SetScope(v EnvironmentVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetIsSecret

`func (o *VariableImportResponseSuccessfulImportedVariables) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *VariableImportResponseSuccessfulImportedVariables) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *VariableImportResponseSuccessfulImportedVariables) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


