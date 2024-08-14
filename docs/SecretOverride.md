# SecretOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**MountPath** | **string** |  | 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableType** | [**APIVariableTypeEnum**](APIVariableTypeEnum.md) |  | 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSecretOverride

`func NewSecretOverride(id string, key string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum, ) *SecretOverride`

NewSecretOverride instantiates a new SecretOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretOverrideWithDefaults

`func NewSecretOverrideWithDefaults() *SecretOverride`

NewSecretOverrideWithDefaults instantiates a new SecretOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretOverride) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretOverride) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *SecretOverride) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretOverride) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretOverride) SetKey(v string)`

SetKey sets Key field to given value.


### GetMountPath

`func (o *SecretOverride) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *SecretOverride) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *SecretOverride) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetScope

`func (o *SecretOverride) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecretOverride) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecretOverride) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *SecretOverride) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *SecretOverride) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *SecretOverride) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.


### GetDescription

`func (o *SecretOverride) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecretOverride) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecretOverride) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecretOverride) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecretOverride) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecretOverride) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *SecretOverride) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *SecretOverride) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *SecretOverride) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *SecretOverride) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *SecretOverride) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *SecretOverride) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


