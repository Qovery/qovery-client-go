# EnvironmentVariableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | key is case sensitive. | 
**Value** | Pointer to **string** | value of the env variable. | [optional] 
**MountPath** | Pointer to **NullableString** | should be set for file only. variable mount path makes variable a file (where file should be mounted). | [optional] 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewEnvironmentVariableRequest

`func NewEnvironmentVariableRequest(key string, ) *EnvironmentVariableRequest`

NewEnvironmentVariableRequest instantiates a new EnvironmentVariableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableRequestWithDefaults

`func NewEnvironmentVariableRequestWithDefaults() *EnvironmentVariableRequest`

NewEnvironmentVariableRequestWithDefaults instantiates a new EnvironmentVariableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EnvironmentVariableRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentVariableRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentVariableRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EnvironmentVariableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariableRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironmentVariableRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMountPath

`func (o *EnvironmentVariableRequest) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *EnvironmentVariableRequest) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *EnvironmentVariableRequest) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *EnvironmentVariableRequest) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *EnvironmentVariableRequest) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *EnvironmentVariableRequest) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetDescription

`func (o *EnvironmentVariableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentVariableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentVariableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentVariableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnvironmentVariableRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnvironmentVariableRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *EnvironmentVariableRequest) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *EnvironmentVariableRequest) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *EnvironmentVariableRequest) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *EnvironmentVariableRequest) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *EnvironmentVariableRequest) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *EnvironmentVariableRequest) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


