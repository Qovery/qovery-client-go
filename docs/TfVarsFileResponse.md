# TfVarsFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The path to the tfvars file within the Git repository | 
**Variables** | **map[string]string** | Map of variable names to their values from the tfvars file | 

## Methods

### NewTfVarsFileResponse

`func NewTfVarsFileResponse(source string, variables map[string]string, ) *TfVarsFileResponse`

NewTfVarsFileResponse instantiates a new TfVarsFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfVarsFileResponseWithDefaults

`func NewTfVarsFileResponseWithDefaults() *TfVarsFileResponse`

NewTfVarsFileResponseWithDefaults instantiates a new TfVarsFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TfVarsFileResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TfVarsFileResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TfVarsFileResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetVariables

`func (o *TfVarsFileResponse) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TfVarsFileResponse) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TfVarsFileResponse) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


