# VariableImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalVariablesToImport** | **float32** |  | 
**SuccessfulImportedVariables** | [**[]VariableImportResponseSuccessfulImportedVariables**](VariableImportResponseSuccessfulImportedVariables.md) |  | 

## Methods

### NewVariableImportResponse

`func NewVariableImportResponse(totalVariablesToImport float32, successfulImportedVariables []VariableImportResponseSuccessfulImportedVariables, ) *VariableImportResponse`

NewVariableImportResponse instantiates a new VariableImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableImportResponseWithDefaults

`func NewVariableImportResponseWithDefaults() *VariableImportResponse`

NewVariableImportResponseWithDefaults instantiates a new VariableImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalVariablesToImport

`func (o *VariableImportResponse) GetTotalVariablesToImport() float32`

GetTotalVariablesToImport returns the TotalVariablesToImport field if non-nil, zero value otherwise.

### GetTotalVariablesToImportOk

`func (o *VariableImportResponse) GetTotalVariablesToImportOk() (*float32, bool)`

GetTotalVariablesToImportOk returns a tuple with the TotalVariablesToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVariablesToImport

`func (o *VariableImportResponse) SetTotalVariablesToImport(v float32)`

SetTotalVariablesToImport sets TotalVariablesToImport field to given value.


### GetSuccessfulImportedVariables

`func (o *VariableImportResponse) GetSuccessfulImportedVariables() []VariableImportResponseSuccessfulImportedVariables`

GetSuccessfulImportedVariables returns the SuccessfulImportedVariables field if non-nil, zero value otherwise.

### GetSuccessfulImportedVariablesOk

`func (o *VariableImportResponse) GetSuccessfulImportedVariablesOk() (*[]VariableImportResponseSuccessfulImportedVariables, bool)`

GetSuccessfulImportedVariablesOk returns a tuple with the SuccessfulImportedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulImportedVariables

`func (o *VariableImportResponse) SetSuccessfulImportedVariables(v []VariableImportResponseSuccessfulImportedVariables)`

SetSuccessfulImportedVariables sets SuccessfulImportedVariables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


