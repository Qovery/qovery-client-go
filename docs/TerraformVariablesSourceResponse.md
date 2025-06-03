# TerraformVariablesSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TfVarFilePaths** | **[]string** |  | 
**TfVars** | **[][]string** | The input is in json array format: [ [$KEY,$VALUE], [...] ] | 

## Methods

### NewTerraformVariablesSourceResponse

`func NewTerraformVariablesSourceResponse(tfVarFilePaths []string, tfVars [][]string, ) *TerraformVariablesSourceResponse`

NewTerraformVariablesSourceResponse instantiates a new TerraformVariablesSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformVariablesSourceResponseWithDefaults

`func NewTerraformVariablesSourceResponseWithDefaults() *TerraformVariablesSourceResponse`

NewTerraformVariablesSourceResponseWithDefaults instantiates a new TerraformVariablesSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTfVarFilePaths

`func (o *TerraformVariablesSourceResponse) GetTfVarFilePaths() []string`

GetTfVarFilePaths returns the TfVarFilePaths field if non-nil, zero value otherwise.

### GetTfVarFilePathsOk

`func (o *TerraformVariablesSourceResponse) GetTfVarFilePathsOk() (*[]string, bool)`

GetTfVarFilePathsOk returns a tuple with the TfVarFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVarFilePaths

`func (o *TerraformVariablesSourceResponse) SetTfVarFilePaths(v []string)`

SetTfVarFilePaths sets TfVarFilePaths field to given value.


### GetTfVars

`func (o *TerraformVariablesSourceResponse) GetTfVars() [][]string`

GetTfVars returns the TfVars field if non-nil, zero value otherwise.

### GetTfVarsOk

`func (o *TerraformVariablesSourceResponse) GetTfVarsOk() (*[][]string, bool)`

GetTfVarsOk returns a tuple with the TfVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVars

`func (o *TerraformVariablesSourceResponse) SetTfVars(v [][]string)`

SetTfVars sets TfVars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


