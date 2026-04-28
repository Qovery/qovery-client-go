# AwsStaticCredentialsAuthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Region** | **string** |  | 
**AccessKey** | **string** |  | 
**SecretKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAwsStaticCredentialsAuthDto

`func NewAwsStaticCredentialsAuthDto(mode string, region string, accessKey string, ) *AwsStaticCredentialsAuthDto`

NewAwsStaticCredentialsAuthDto instantiates a new AwsStaticCredentialsAuthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsStaticCredentialsAuthDtoWithDefaults

`func NewAwsStaticCredentialsAuthDtoWithDefaults() *AwsStaticCredentialsAuthDto`

NewAwsStaticCredentialsAuthDtoWithDefaults instantiates a new AwsStaticCredentialsAuthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AwsStaticCredentialsAuthDto) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AwsStaticCredentialsAuthDto) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AwsStaticCredentialsAuthDto) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRegion

`func (o *AwsStaticCredentialsAuthDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsStaticCredentialsAuthDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsStaticCredentialsAuthDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessKey

`func (o *AwsStaticCredentialsAuthDto) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AwsStaticCredentialsAuthDto) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AwsStaticCredentialsAuthDto) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *AwsStaticCredentialsAuthDto) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AwsStaticCredentialsAuthDto) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AwsStaticCredentialsAuthDto) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AwsStaticCredentialsAuthDto) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *AwsStaticCredentialsAuthDto) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *AwsStaticCredentialsAuthDto) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


