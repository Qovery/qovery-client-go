# AwsCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsCredentialsRequest

`func NewAwsCredentialsRequest(name string, ) *AwsCredentialsRequest`

NewAwsCredentialsRequest instantiates a new AwsCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCredentialsRequestWithDefaults

`func NewAwsCredentialsRequestWithDefaults() *AwsCredentialsRequest`

NewAwsCredentialsRequestWithDefaults instantiates a new AwsCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasAccessKeyId

`func (o *AwsCredentialsRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

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

### HasSecretAccessKey

`func (o *AwsCredentialsRequest) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


