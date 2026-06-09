# BlueprintVariableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsSecret** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBlueprintVariableRequest

`func NewBlueprintVariableRequest(name string, value string, ) *BlueprintVariableRequest`

NewBlueprintVariableRequest instantiates a new BlueprintVariableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintVariableRequestWithDefaults

`func NewBlueprintVariableRequestWithDefaults() *BlueprintVariableRequest`

NewBlueprintVariableRequestWithDefaults instantiates a new BlueprintVariableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintVariableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintVariableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintVariableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *BlueprintVariableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BlueprintVariableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BlueprintVariableRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecret

`func (o *BlueprintVariableRequest) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintVariableRequest) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintVariableRequest) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *BlueprintVariableRequest) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


