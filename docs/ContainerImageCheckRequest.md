# ContainerImageCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryId** | Pointer to **string** |  | [optional] 
**ImageName** | **string** |  | 
**Tag** | **string** |  | 

## Methods

### NewContainerImageCheckRequest

`func NewContainerImageCheckRequest(imageName string, tag string, ) *ContainerImageCheckRequest`

NewContainerImageCheckRequest instantiates a new ContainerImageCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageCheckRequestWithDefaults

`func NewContainerImageCheckRequestWithDefaults() *ContainerImageCheckRequest`

NewContainerImageCheckRequestWithDefaults instantiates a new ContainerImageCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryId

`func (o *ContainerImageCheckRequest) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *ContainerImageCheckRequest) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *ContainerImageCheckRequest) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *ContainerImageCheckRequest) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetImageName

`func (o *ContainerImageCheckRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerImageCheckRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerImageCheckRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *ContainerImageCheckRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ContainerImageCheckRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ContainerImageCheckRequest) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


