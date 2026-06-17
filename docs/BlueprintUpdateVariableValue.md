# BlueprintUpdateVariableValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**IsSecret** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBlueprintUpdateVariableValue

`func NewBlueprintUpdateVariableValue(value string, ) *BlueprintUpdateVariableValue`

NewBlueprintUpdateVariableValue instantiates a new BlueprintUpdateVariableValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateVariableValueWithDefaults

`func NewBlueprintUpdateVariableValueWithDefaults() *BlueprintUpdateVariableValue`

NewBlueprintUpdateVariableValueWithDefaults instantiates a new BlueprintUpdateVariableValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *BlueprintUpdateVariableValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BlueprintUpdateVariableValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BlueprintUpdateVariableValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecret

`func (o *BlueprintUpdateVariableValue) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintUpdateVariableValue) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintUpdateVariableValue) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *BlueprintUpdateVariableValue) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


