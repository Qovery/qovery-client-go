# HelmRepositoryRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Required if the repository is private | [optional] 
**Password** | Pointer to **string** | Required if the repository is private | [optional] 
**AccessKeyId** | Pointer to **string** | Required if kind is &#x60;ECR&#x60; or &#x60;PUBLIC_ECR&#x60; | [optional] 
**SecretAccessKey** | Pointer to **string** | Required if kind is &#x60;ECR&#x60; or &#x60;PUBLIC_ECR&#x60; | [optional] 
**Region** | Pointer to **string** | Required if kind is &#x60;ECR&#x60; or &#x60;SCALEWAY_CR&#x60; | [optional] 
**ScalewayAccessKey** | Pointer to **string** | Required if kind is &#x60;SCALEWAY_CR&#x60; | [optional] 
**ScalewaySecretKey** | Pointer to **string** | Required if kind is &#x60;SCALEWAY_CR&#x60; | [optional] 
**ScalewayProjectId** | Pointer to **string** | Required if kind is &#x60;SCALEWAY_CR&#x60; | [optional] 
**AzureTenantId** | Pointer to **string** | Required if kind is &#x60;AZURE_CR&#x60;. | [optional] 
**AzureSubscriptionId** | Pointer to **string** | Required if kind is &#x60;AZURE_CR&#x60;. | [optional] 

## Methods

### NewHelmRepositoryRequestConfig

`func NewHelmRepositoryRequestConfig() *HelmRepositoryRequestConfig`

NewHelmRepositoryRequestConfig instantiates a new HelmRepositoryRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepositoryRequestConfigWithDefaults

`func NewHelmRepositoryRequestConfigWithDefaults() *HelmRepositoryRequestConfig`

NewHelmRepositoryRequestConfigWithDefaults instantiates a new HelmRepositoryRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *HelmRepositoryRequestConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HelmRepositoryRequestConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HelmRepositoryRequestConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HelmRepositoryRequestConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *HelmRepositoryRequestConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HelmRepositoryRequestConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HelmRepositoryRequestConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HelmRepositoryRequestConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *HelmRepositoryRequestConfig) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *HelmRepositoryRequestConfig) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *HelmRepositoryRequestConfig) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *HelmRepositoryRequestConfig) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *HelmRepositoryRequestConfig) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *HelmRepositoryRequestConfig) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *HelmRepositoryRequestConfig) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *HelmRepositoryRequestConfig) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetRegion

`func (o *HelmRepositoryRequestConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HelmRepositoryRequestConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HelmRepositoryRequestConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HelmRepositoryRequestConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScalewayAccessKey

`func (o *HelmRepositoryRequestConfig) GetScalewayAccessKey() string`

GetScalewayAccessKey returns the ScalewayAccessKey field if non-nil, zero value otherwise.

### GetScalewayAccessKeyOk

`func (o *HelmRepositoryRequestConfig) GetScalewayAccessKeyOk() (*string, bool)`

GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayAccessKey

`func (o *HelmRepositoryRequestConfig) SetScalewayAccessKey(v string)`

SetScalewayAccessKey sets ScalewayAccessKey field to given value.

### HasScalewayAccessKey

`func (o *HelmRepositoryRequestConfig) HasScalewayAccessKey() bool`

HasScalewayAccessKey returns a boolean if a field has been set.

### GetScalewaySecretKey

`func (o *HelmRepositoryRequestConfig) GetScalewaySecretKey() string`

GetScalewaySecretKey returns the ScalewaySecretKey field if non-nil, zero value otherwise.

### GetScalewaySecretKeyOk

`func (o *HelmRepositoryRequestConfig) GetScalewaySecretKeyOk() (*string, bool)`

GetScalewaySecretKeyOk returns a tuple with the ScalewaySecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewaySecretKey

`func (o *HelmRepositoryRequestConfig) SetScalewaySecretKey(v string)`

SetScalewaySecretKey sets ScalewaySecretKey field to given value.

### HasScalewaySecretKey

`func (o *HelmRepositoryRequestConfig) HasScalewaySecretKey() bool`

HasScalewaySecretKey returns a boolean if a field has been set.

### GetScalewayProjectId

`func (o *HelmRepositoryRequestConfig) GetScalewayProjectId() string`

GetScalewayProjectId returns the ScalewayProjectId field if non-nil, zero value otherwise.

### GetScalewayProjectIdOk

`func (o *HelmRepositoryRequestConfig) GetScalewayProjectIdOk() (*string, bool)`

GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayProjectId

`func (o *HelmRepositoryRequestConfig) SetScalewayProjectId(v string)`

SetScalewayProjectId sets ScalewayProjectId field to given value.

### HasScalewayProjectId

`func (o *HelmRepositoryRequestConfig) HasScalewayProjectId() bool`

HasScalewayProjectId returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *HelmRepositoryRequestConfig) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *HelmRepositoryRequestConfig) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *HelmRepositoryRequestConfig) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *HelmRepositoryRequestConfig) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetAzureSubscriptionId

`func (o *HelmRepositoryRequestConfig) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *HelmRepositoryRequestConfig) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *HelmRepositoryRequestConfig) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *HelmRepositoryRequestConfig) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


