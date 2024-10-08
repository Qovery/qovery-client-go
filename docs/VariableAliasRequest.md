# VariableAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | the value to be used as Alias of the targeted environment variable. | 
**AliasScope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**AliasParentId** | **string** | the id of the variable that is aliased. | 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewVariableAliasRequest

`func NewVariableAliasRequest(key string, aliasScope APIVariableScopeEnum, aliasParentId string, ) *VariableAliasRequest`

NewVariableAliasRequest instantiates a new VariableAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableAliasRequestWithDefaults

`func NewVariableAliasRequestWithDefaults() *VariableAliasRequest`

NewVariableAliasRequestWithDefaults instantiates a new VariableAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableAliasRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableAliasRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableAliasRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetAliasScope

`func (o *VariableAliasRequest) GetAliasScope() APIVariableScopeEnum`

GetAliasScope returns the AliasScope field if non-nil, zero value otherwise.

### GetAliasScopeOk

`func (o *VariableAliasRequest) GetAliasScopeOk() (*APIVariableScopeEnum, bool)`

GetAliasScopeOk returns a tuple with the AliasScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasScope

`func (o *VariableAliasRequest) SetAliasScope(v APIVariableScopeEnum)`

SetAliasScope sets AliasScope field to given value.


### GetAliasParentId

`func (o *VariableAliasRequest) GetAliasParentId() string`

GetAliasParentId returns the AliasParentId field if non-nil, zero value otherwise.

### GetAliasParentIdOk

`func (o *VariableAliasRequest) GetAliasParentIdOk() (*string, bool)`

GetAliasParentIdOk returns a tuple with the AliasParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasParentId

`func (o *VariableAliasRequest) SetAliasParentId(v string)`

SetAliasParentId sets AliasParentId field to given value.


### GetDescription

`func (o *VariableAliasRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableAliasRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableAliasRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableAliasRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VariableAliasRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VariableAliasRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *VariableAliasRequest) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *VariableAliasRequest) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *VariableAliasRequest) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *VariableAliasRequest) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *VariableAliasRequest) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *VariableAliasRequest) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


