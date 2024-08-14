# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewKey

`func NewKey(key string, ) *Key`

NewKey instantiates a new Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyWithDefaults

`func NewKeyWithDefaults() *Key`

NewKeyWithDefaults instantiates a new Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Key) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Key) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Key) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *Key) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Key) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Key) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Key) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Key) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Key) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *Key) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *Key) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *Key) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *Key) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *Key) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *Key) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


