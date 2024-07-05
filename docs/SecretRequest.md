# SecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | key is case sensitive | 
**Value** | Pointer to **string** | value of the secret. Clear value will never be returned | [optional] 
**MountPath** | Pointer to **NullableString** | should be set for file only. variable mount path make secret a file (where file should be mounted). | [optional] 
**Description** | Pointer to **NullableString** | optional variable description (255 character maximum) | [optional] 

## Methods

### NewSecretRequest

`func NewSecretRequest(key string, ) *SecretRequest`

NewSecretRequest instantiates a new SecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretRequestWithDefaults

`func NewSecretRequestWithDefaults() *SecretRequest`

NewSecretRequestWithDefaults instantiates a new SecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SecretRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *SecretRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecretRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecretRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SecretRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMountPath

`func (o *SecretRequest) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *SecretRequest) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *SecretRequest) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *SecretRequest) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *SecretRequest) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *SecretRequest) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetDescription

`func (o *SecretRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecretRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecretRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecretRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecretRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecretRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


