# ContainerDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageTag** | **string** | Image tag to deploy | 

## Methods

### NewContainerDeployRequest

`func NewContainerDeployRequest(imageTag string, ) *ContainerDeployRequest`

NewContainerDeployRequest instantiates a new ContainerDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerDeployRequestWithDefaults

`func NewContainerDeployRequestWithDefaults() *ContainerDeployRequest`

NewContainerDeployRequestWithDefaults instantiates a new ContainerDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageTag

`func (o *ContainerDeployRequest) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ContainerDeployRequest) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ContainerDeployRequest) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


