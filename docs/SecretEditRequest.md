# SecretEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Description** | Pointer to **string** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecretEditRequest

`func NewSecretEditRequest(key string, ) *SecretEditRequest`

NewSecretEditRequest instantiates a new SecretEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretEditRequestWithDefaults

`func NewSecretEditRequestWithDefaults() *SecretEditRequest`

NewSecretEditRequestWithDefaults instantiates a new SecretEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SecretEditRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecretEditRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecretEditRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SecretEditRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *SecretEditRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretEditRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretEditRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *SecretEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecretEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecretEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecretEditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableInterpolationInFile

`func (o *SecretEditRequest) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *SecretEditRequest) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *SecretEditRequest) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *SecretEditRequest) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


