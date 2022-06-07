# VariableImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overwrite** | **bool** |  | [default to false]
**Vars** | [**[]VariableImportRequestVarsInner**](VariableImportRequestVarsInner.md) |  | 

## Methods

### NewVariableImportRequest

`func NewVariableImportRequest(overwrite bool, vars []VariableImportRequestVarsInner, ) *VariableImportRequest`

NewVariableImportRequest instantiates a new VariableImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableImportRequestWithDefaults

`func NewVariableImportRequestWithDefaults() *VariableImportRequest`

NewVariableImportRequestWithDefaults instantiates a new VariableImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverwrite

`func (o *VariableImportRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *VariableImportRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *VariableImportRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.


### GetVars

`func (o *VariableImportRequest) GetVars() []VariableImportRequestVarsInner`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *VariableImportRequest) GetVarsOk() (*[]VariableImportRequestVarsInner, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *VariableImportRequest) SetVars(v []VariableImportRequestVarsInner)`

SetVars sets Vars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


