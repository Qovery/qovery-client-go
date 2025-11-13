# TerraformVariablesReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | [**[]TerraformVarKeyValue**](TerraformVarKeyValue.md) | Complete list of variables to set (replaces all existing variables) | 

## Methods

### NewTerraformVariablesReplaceRequest

`func NewTerraformVariablesReplaceRequest(variables []TerraformVarKeyValue, ) *TerraformVariablesReplaceRequest`

NewTerraformVariablesReplaceRequest instantiates a new TerraformVariablesReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformVariablesReplaceRequestWithDefaults

`func NewTerraformVariablesReplaceRequestWithDefaults() *TerraformVariablesReplaceRequest`

NewTerraformVariablesReplaceRequestWithDefaults instantiates a new TerraformVariablesReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *TerraformVariablesReplaceRequest) GetVariables() []TerraformVarKeyValue`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TerraformVariablesReplaceRequest) GetVariablesOk() (*[]TerraformVarKeyValue, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TerraformVariablesReplaceRequest) SetVariables(v []TerraformVarKeyValue)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


