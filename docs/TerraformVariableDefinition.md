# TerraformVariableDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The name of the variable | 
**Sensitive** | **bool** | Whether the variable is marked as sensitive | 
**Default** | Pointer to **NullableString** | The default value of the variable, or null if no default is provided | [optional] 

## Methods

### NewTerraformVariableDefinition

`func NewTerraformVariableDefinition(key string, sensitive bool, ) *TerraformVariableDefinition`

NewTerraformVariableDefinition instantiates a new TerraformVariableDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformVariableDefinitionWithDefaults

`func NewTerraformVariableDefinitionWithDefaults() *TerraformVariableDefinition`

NewTerraformVariableDefinitionWithDefaults instantiates a new TerraformVariableDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TerraformVariableDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TerraformVariableDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TerraformVariableDefinition) SetKey(v string)`

SetKey sets Key field to given value.


### GetSensitive

`func (o *TerraformVariableDefinition) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *TerraformVariableDefinition) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *TerraformVariableDefinition) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetDefault

`func (o *TerraformVariableDefinition) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TerraformVariableDefinition) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TerraformVariableDefinition) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TerraformVariableDefinition) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *TerraformVariableDefinition) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *TerraformVariableDefinition) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


