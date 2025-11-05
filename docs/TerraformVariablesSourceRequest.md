# TerraformVariablesSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TfVarFilePaths** | **[]string** |  | 
**TfVars** | [**[]TerraformVarKeyValue**](TerraformVarKeyValue.md) |  | 

## Methods

### NewTerraformVariablesSourceRequest

`func NewTerraformVariablesSourceRequest(tfVarFilePaths []string, tfVars []TerraformVarKeyValue, ) *TerraformVariablesSourceRequest`

NewTerraformVariablesSourceRequest instantiates a new TerraformVariablesSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformVariablesSourceRequestWithDefaults

`func NewTerraformVariablesSourceRequestWithDefaults() *TerraformVariablesSourceRequest`

NewTerraformVariablesSourceRequestWithDefaults instantiates a new TerraformVariablesSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTfVarFilePaths

`func (o *TerraformVariablesSourceRequest) GetTfVarFilePaths() []string`

GetTfVarFilePaths returns the TfVarFilePaths field if non-nil, zero value otherwise.

### GetTfVarFilePathsOk

`func (o *TerraformVariablesSourceRequest) GetTfVarFilePathsOk() (*[]string, bool)`

GetTfVarFilePathsOk returns a tuple with the TfVarFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVarFilePaths

`func (o *TerraformVariablesSourceRequest) SetTfVarFilePaths(v []string)`

SetTfVarFilePaths sets TfVarFilePaths field to given value.


### GetTfVars

`func (o *TerraformVariablesSourceRequest) GetTfVars() []TerraformVarKeyValue`

GetTfVars returns the TfVars field if non-nil, zero value otherwise.

### GetTfVarsOk

`func (o *TerraformVariablesSourceRequest) GetTfVarsOk() (*[]TerraformVarKeyValue, bool)`

GetTfVarsOk returns a tuple with the TfVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVars

`func (o *TerraformVariablesSourceRequest) SetTfVars(v []TerraformVarKeyValue)`

SetTfVars sets TfVars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


