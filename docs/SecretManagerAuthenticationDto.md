# SecretManagerAuthenticationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**RoleArn** | **string** |  | 
**Region** | **string** |  | 
**AccessKey** | **string** |  | 
**SecretKey** | Pointer to **NullableString** |  | [optional] 
**JsonCredentials** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecretManagerAuthenticationDto

`func NewSecretManagerAuthenticationDto(mode string, roleArn string, region string, accessKey string, ) *SecretManagerAuthenticationDto`

NewSecretManagerAuthenticationDto instantiates a new SecretManagerAuthenticationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretManagerAuthenticationDtoWithDefaults

`func NewSecretManagerAuthenticationDtoWithDefaults() *SecretManagerAuthenticationDto`

NewSecretManagerAuthenticationDtoWithDefaults instantiates a new SecretManagerAuthenticationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *SecretManagerAuthenticationDto) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SecretManagerAuthenticationDto) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SecretManagerAuthenticationDto) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRoleArn

`func (o *SecretManagerAuthenticationDto) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SecretManagerAuthenticationDto) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SecretManagerAuthenticationDto) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRegion

`func (o *SecretManagerAuthenticationDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SecretManagerAuthenticationDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SecretManagerAuthenticationDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessKey

`func (o *SecretManagerAuthenticationDto) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SecretManagerAuthenticationDto) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SecretManagerAuthenticationDto) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *SecretManagerAuthenticationDto) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SecretManagerAuthenticationDto) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SecretManagerAuthenticationDto) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SecretManagerAuthenticationDto) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *SecretManagerAuthenticationDto) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *SecretManagerAuthenticationDto) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetJsonCredentials

`func (o *SecretManagerAuthenticationDto) GetJsonCredentials() string`

GetJsonCredentials returns the JsonCredentials field if non-nil, zero value otherwise.

### GetJsonCredentialsOk

`func (o *SecretManagerAuthenticationDto) GetJsonCredentialsOk() (*string, bool)`

GetJsonCredentialsOk returns a tuple with the JsonCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonCredentials

`func (o *SecretManagerAuthenticationDto) SetJsonCredentials(v string)`

SetJsonCredentials sets JsonCredentials field to given value.

### HasJsonCredentials

`func (o *SecretManagerAuthenticationDto) HasJsonCredentials() bool`

HasJsonCredentials returns a boolean if a field has been set.

### SetJsonCredentialsNil

`func (o *SecretManagerAuthenticationDto) SetJsonCredentialsNil(b bool)`

 SetJsonCredentialsNil sets the value for JsonCredentials to be an explicit nil

### UnsetJsonCredentials
`func (o *SecretManagerAuthenticationDto) UnsetJsonCredentials()`

UnsetJsonCredentials ensures that no value is present for JsonCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


