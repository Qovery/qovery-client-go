# ContainerRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Kind** | [**ContainerRegistryKindEnum**](ContainerRegistryKindEnum.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the container registry: * For &#x60;DOCKER_HUB&#x60;: it must be &#x60;https://docker.io&#x60; (default with &#39;https://docker.io&#39; if no url provided for &#x60;DOCKER_HUB&#x60;) * For &#x60;GITHUB_CR&#x60;: it must be &#x60;https://ghcr.io&#x60; (default with &#39;https://ghcr.io&#39; if no url provided for &#x60;GITHUB_CR&#x60;) * For &#x60;GITLAB_CR&#x60;: it must be &#x60;https://registry.gitlab.com&#x60; (default with &#39;https://registry.gitlab.com&#39; if no url provided for &#x60;GITLAB_CR&#x60;) * For others: it&#39;s required and must start by &#x60;https://&#x60;  | [optional] 
**Config** | [**ContainerRegistryRequestConfig**](ContainerRegistryRequestConfig.md) |  | 

## Methods

### NewContainerRegistryRequest

`func NewContainerRegistryRequest(name string, kind ContainerRegistryKindEnum, config ContainerRegistryRequestConfig, ) *ContainerRegistryRequest`

NewContainerRegistryRequest instantiates a new ContainerRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryRequestWithDefaults

`func NewContainerRegistryRequestWithDefaults() *ContainerRegistryRequest`

NewContainerRegistryRequestWithDefaults instantiates a new ContainerRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerRegistryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRegistryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRegistryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *ContainerRegistryRequest) GetKind() ContainerRegistryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContainerRegistryRequest) GetKindOk() (*ContainerRegistryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContainerRegistryRequest) SetKind(v ContainerRegistryKindEnum)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *ContainerRegistryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerRegistryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerRegistryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerRegistryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ContainerRegistryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerRegistryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerRegistryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ContainerRegistryRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetConfig

`func (o *ContainerRegistryRequest) GetConfig() ContainerRegistryRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerRegistryRequest) GetConfigOk() (*ContainerRegistryRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerRegistryRequest) SetConfig(v ContainerRegistryRequestConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


