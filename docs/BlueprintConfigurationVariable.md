# BlueprintConfigurationVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **string** | Omitted for secret variables. | [optional] 
**IsSecret** | **bool** |  | 

## Methods

### NewBlueprintConfigurationVariable

`func NewBlueprintConfigurationVariable(name string, isSecret bool, ) *BlueprintConfigurationVariable`

NewBlueprintConfigurationVariable instantiates a new BlueprintConfigurationVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintConfigurationVariableWithDefaults

`func NewBlueprintConfigurationVariableWithDefaults() *BlueprintConfigurationVariable`

NewBlueprintConfigurationVariableWithDefaults instantiates a new BlueprintConfigurationVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintConfigurationVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintConfigurationVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintConfigurationVariable) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *BlueprintConfigurationVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BlueprintConfigurationVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BlueprintConfigurationVariable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BlueprintConfigurationVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsSecret

`func (o *BlueprintConfigurationVariable) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintConfigurationVariable) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintConfigurationVariable) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


