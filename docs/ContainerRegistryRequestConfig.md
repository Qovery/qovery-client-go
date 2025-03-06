# ContainerRegistryRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | Required if kind is &#x60;ECR&#x60; or &#x60;PUBLIC_ECR&#x60; | [optional] 
**SecretAccessKey** | Pointer to **string** | Required if kind is &#x60;ECR&#x60; or &#x60;PUBLIC_ECR&#x60; | [optional] 
**Region** | Pointer to **string** | Required if kind is &#x60;ECR&#x60; or &#x60;SCALEWAY_CR&#x60; | [optional] 
**ScalewayAccessKey** | Pointer to **string** | Required if kind is &#x60;SCALEWAY_CR&#x60; | [optional] 
**ScalewaySecretKey** | Pointer to **string** | Required if kind is &#x60;SCALEWAY_CR&#x60; | [optional] 
**ScalewayProjectId** | Pointer to **string** | Required if kind is &#x60;SCALEWAY_CR&#x60; | [optional] 
**JsonCredentials** | Pointer to **string** | Required if kind is &#x60;GCP_ARTIFACT_REGISTRY&#x60; | [optional] 
**Username** | Pointer to **string** | optional, for kind &#x60;DOCKER_HUB&#x60;   We encourage you to set credentials for Docker Hub due to the limits on the pull rate  | [optional] 
**Password** | Pointer to **string** | optional, for kind &#x60;DOCKER_HUB&#x60;   We encourage you to set credentials for Docker Hub due to the limits on the pull rate  | [optional] 
**RoleArn** | Pointer to **string** | For ECR, you can either set a static access_key or use a role arn that we are going to assume | [optional] 

## Methods

### NewContainerRegistryRequestConfig

`func NewContainerRegistryRequestConfig() *ContainerRegistryRequestConfig`

NewContainerRegistryRequestConfig instantiates a new ContainerRegistryRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryRequestConfigWithDefaults

`func NewContainerRegistryRequestConfigWithDefaults() *ContainerRegistryRequestConfig`

NewContainerRegistryRequestConfigWithDefaults instantiates a new ContainerRegistryRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *ContainerRegistryRequestConfig) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ContainerRegistryRequestConfig) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ContainerRegistryRequestConfig) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *ContainerRegistryRequestConfig) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *ContainerRegistryRequestConfig) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ContainerRegistryRequestConfig) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ContainerRegistryRequestConfig) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *ContainerRegistryRequestConfig) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetRegion

`func (o *ContainerRegistryRequestConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ContainerRegistryRequestConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ContainerRegistryRequestConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ContainerRegistryRequestConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScalewayAccessKey

`func (o *ContainerRegistryRequestConfig) GetScalewayAccessKey() string`

GetScalewayAccessKey returns the ScalewayAccessKey field if non-nil, zero value otherwise.

### GetScalewayAccessKeyOk

`func (o *ContainerRegistryRequestConfig) GetScalewayAccessKeyOk() (*string, bool)`

GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayAccessKey

`func (o *ContainerRegistryRequestConfig) SetScalewayAccessKey(v string)`

SetScalewayAccessKey sets ScalewayAccessKey field to given value.

### HasScalewayAccessKey

`func (o *ContainerRegistryRequestConfig) HasScalewayAccessKey() bool`

HasScalewayAccessKey returns a boolean if a field has been set.

### GetScalewaySecretKey

`func (o *ContainerRegistryRequestConfig) GetScalewaySecretKey() string`

GetScalewaySecretKey returns the ScalewaySecretKey field if non-nil, zero value otherwise.

### GetScalewaySecretKeyOk

`func (o *ContainerRegistryRequestConfig) GetScalewaySecretKeyOk() (*string, bool)`

GetScalewaySecretKeyOk returns a tuple with the ScalewaySecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewaySecretKey

`func (o *ContainerRegistryRequestConfig) SetScalewaySecretKey(v string)`

SetScalewaySecretKey sets ScalewaySecretKey field to given value.

### HasScalewaySecretKey

`func (o *ContainerRegistryRequestConfig) HasScalewaySecretKey() bool`

HasScalewaySecretKey returns a boolean if a field has been set.

### GetScalewayProjectId

`func (o *ContainerRegistryRequestConfig) GetScalewayProjectId() string`

GetScalewayProjectId returns the ScalewayProjectId field if non-nil, zero value otherwise.

### GetScalewayProjectIdOk

`func (o *ContainerRegistryRequestConfig) GetScalewayProjectIdOk() (*string, bool)`

GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayProjectId

`func (o *ContainerRegistryRequestConfig) SetScalewayProjectId(v string)`

SetScalewayProjectId sets ScalewayProjectId field to given value.

### HasScalewayProjectId

`func (o *ContainerRegistryRequestConfig) HasScalewayProjectId() bool`

HasScalewayProjectId returns a boolean if a field has been set.

### GetJsonCredentials

`func (o *ContainerRegistryRequestConfig) GetJsonCredentials() string`

GetJsonCredentials returns the JsonCredentials field if non-nil, zero value otherwise.

### GetJsonCredentialsOk

`func (o *ContainerRegistryRequestConfig) GetJsonCredentialsOk() (*string, bool)`

GetJsonCredentialsOk returns a tuple with the JsonCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonCredentials

`func (o *ContainerRegistryRequestConfig) SetJsonCredentials(v string)`

SetJsonCredentials sets JsonCredentials field to given value.

### HasJsonCredentials

`func (o *ContainerRegistryRequestConfig) HasJsonCredentials() bool`

HasJsonCredentials returns a boolean if a field has been set.

### GetUsername

`func (o *ContainerRegistryRequestConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ContainerRegistryRequestConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ContainerRegistryRequestConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ContainerRegistryRequestConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ContainerRegistryRequestConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ContainerRegistryRequestConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ContainerRegistryRequestConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ContainerRegistryRequestConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoleArn

`func (o *ContainerRegistryRequestConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ContainerRegistryRequestConfig) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ContainerRegistryRequestConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *ContainerRegistryRequestConfig) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


