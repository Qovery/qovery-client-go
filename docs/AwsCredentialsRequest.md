# AwsCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**RoleArn** | **string** |  | 
**AccessKeyId** | **string** |  | 
**SecretAccessKey** | **string** |  | 
**VsphereUser** | **string** |  | 
**VspherePassword** | **string** |  | 

## Methods

### NewAwsCredentialsRequest

`func NewAwsCredentialsRequest(type_ string, name string, roleArn string, accessKeyId string, secretAccessKey string, vsphereUser string, vspherePassword string, ) *AwsCredentialsRequest`

NewAwsCredentialsRequest instantiates a new AwsCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCredentialsRequestWithDefaults

`func NewAwsCredentialsRequestWithDefaults() *AwsCredentialsRequest`

NewAwsCredentialsRequestWithDefaults instantiates a new AwsCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwsCredentialsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsCredentialsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsCredentialsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AwsCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRoleArn

`func (o *AwsCredentialsRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsCredentialsRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsCredentialsRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetAccessKeyId

`func (o *AwsCredentialsRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AwsCredentialsRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AwsCredentialsRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetSecretAccessKey

`func (o *AwsCredentialsRequest) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsCredentialsRequest) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsCredentialsRequest) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### GetVsphereUser

`func (o *AwsCredentialsRequest) GetVsphereUser() string`

GetVsphereUser returns the VsphereUser field if non-nil, zero value otherwise.

### GetVsphereUserOk

`func (o *AwsCredentialsRequest) GetVsphereUserOk() (*string, bool)`

GetVsphereUserOk returns a tuple with the VsphereUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereUser

`func (o *AwsCredentialsRequest) SetVsphereUser(v string)`

SetVsphereUser sets VsphereUser field to given value.


### GetVspherePassword

`func (o *AwsCredentialsRequest) GetVspherePassword() string`

GetVspherePassword returns the VspherePassword field if non-nil, zero value otherwise.

### GetVspherePasswordOk

`func (o *AwsCredentialsRequest) GetVspherePasswordOk() (*string, bool)`

GetVspherePasswordOk returns a tuple with the VspherePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVspherePassword

`func (o *AwsCredentialsRequest) SetVspherePassword(v string)`

SetVspherePassword sets VspherePassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


