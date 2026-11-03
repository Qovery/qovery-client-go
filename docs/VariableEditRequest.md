# VariableEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | the key of the environment variable | 
**Value** | Pointer to **NullableString** | the value of the environment variable | [optional] 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 
**SecretManagerAccessId** | Pointer to **NullableString** |  | [optional] 
**JsonTargetKey** | Pointer to **NullableString** | Optional JSON member/index path, only supported for &#x60;EXTERNAL_SECRET&#x60; variables (not external secret files). Use dots for nested members and numeric indexes for arrays, for example &#x60;accounts.0.password&#x60;. Escape literal dots with &#x60;\\.&#x60; and backslashes with &#x60;\\\\&#x60;; other query characters are treated literally. Members must be non-empty, with no leading or trailing whitespace, control characters, or unescaped trailing backslash. Omit or set to null to clear JSON selection and read the whole secret. Send the current selector on unrelated edits to preserve it. | [optional] 

## Methods

### NewVariableEditRequest

`func NewVariableEditRequest(key string, ) *VariableEditRequest`

NewVariableEditRequest instantiates a new VariableEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableEditRequestWithDefaults

`func NewVariableEditRequestWithDefaults() *VariableEditRequest`

NewVariableEditRequestWithDefaults instantiates a new VariableEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableEditRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableEditRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableEditRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableEditRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableEditRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableEditRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableEditRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableEditRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableEditRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *VariableEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableEditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VariableEditRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VariableEditRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *VariableEditRequest) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *VariableEditRequest) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *VariableEditRequest) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *VariableEditRequest) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *VariableEditRequest) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *VariableEditRequest) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil
### GetSecretManagerAccessId

`func (o *VariableEditRequest) GetSecretManagerAccessId() string`

GetSecretManagerAccessId returns the SecretManagerAccessId field if non-nil, zero value otherwise.

### GetSecretManagerAccessIdOk

`func (o *VariableEditRequest) GetSecretManagerAccessIdOk() (*string, bool)`

GetSecretManagerAccessIdOk returns a tuple with the SecretManagerAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerAccessId

`func (o *VariableEditRequest) SetSecretManagerAccessId(v string)`

SetSecretManagerAccessId sets SecretManagerAccessId field to given value.

### HasSecretManagerAccessId

`func (o *VariableEditRequest) HasSecretManagerAccessId() bool`

HasSecretManagerAccessId returns a boolean if a field has been set.

### SetSecretManagerAccessIdNil

`func (o *VariableEditRequest) SetSecretManagerAccessIdNil(b bool)`

 SetSecretManagerAccessIdNil sets the value for SecretManagerAccessId to be an explicit nil

### UnsetSecretManagerAccessId
`func (o *VariableEditRequest) UnsetSecretManagerAccessId()`

UnsetSecretManagerAccessId ensures that no value is present for SecretManagerAccessId, not even an explicit nil
### GetJsonTargetKey

`func (o *VariableEditRequest) GetJsonTargetKey() string`

GetJsonTargetKey returns the JsonTargetKey field if non-nil, zero value otherwise.

### GetJsonTargetKeyOk

`func (o *VariableEditRequest) GetJsonTargetKeyOk() (*string, bool)`

GetJsonTargetKeyOk returns a tuple with the JsonTargetKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonTargetKey

`func (o *VariableEditRequest) SetJsonTargetKey(v string)`

SetJsonTargetKey sets JsonTargetKey field to given value.

### HasJsonTargetKey

`func (o *VariableEditRequest) HasJsonTargetKey() bool`

HasJsonTargetKey returns a boolean if a field has been set.

### SetJsonTargetKeyNil

`func (o *VariableEditRequest) SetJsonTargetKeyNil(b bool)`

 SetJsonTargetKeyNil sets the value for JsonTargetKey to be an explicit nil

### UnsetJsonTargetKey
`func (o *VariableEditRequest) UnsetJsonTargetKey()`

UnsetJsonTargetKey ensures that no value is present for JsonTargetKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


