# ContainerSourceRequestOneOfImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryId** | **string** | id of the linked registry | 
**ImageName** | **string** | The image name pattern differs according to chosen container registry provider: * &#x60;ECR&#x60;: &#x60;repository&#x60; * &#x60;SCALEWAY_CR&#x60;: &#x60;namespace/image&#x60; * &#x60;DOCKER_HUB&#x60;: &#x60;image&#x60; or &#x60;repository/image&#x60; * &#x60;PUBLIC_ECR&#x60;: &#x60;registry_alias/repository&#x60;  | 
**Tag** | **string** | tag of the image container | 

## Methods

### NewContainerSourceRequestOneOfImage

`func NewContainerSourceRequestOneOfImage(registryId string, imageName string, tag string, ) *ContainerSourceRequestOneOfImage`

NewContainerSourceRequestOneOfImage instantiates a new ContainerSourceRequestOneOfImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerSourceRequestOneOfImageWithDefaults

`func NewContainerSourceRequestOneOfImageWithDefaults() *ContainerSourceRequestOneOfImage`

NewContainerSourceRequestOneOfImageWithDefaults instantiates a new ContainerSourceRequestOneOfImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryId

`func (o *ContainerSourceRequestOneOfImage) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *ContainerSourceRequestOneOfImage) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *ContainerSourceRequestOneOfImage) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.


### GetImageName

`func (o *ContainerSourceRequestOneOfImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerSourceRequestOneOfImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerSourceRequestOneOfImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *ContainerSourceRequestOneOfImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ContainerSourceRequestOneOfImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ContainerSourceRequestOneOfImage) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


