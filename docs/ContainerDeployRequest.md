# ContainerDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ImageTag** | Pointer to **NullableString** | Image tag to deploy | [optional] 

## Methods

### NewContainerDeployRequest

`func NewContainerDeployRequest() *ContainerDeployRequest`

NewContainerDeployRequest instantiates a new ContainerDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerDeployRequestWithDefaults

`func NewContainerDeployRequestWithDefaults() *ContainerDeployRequest`

NewContainerDeployRequestWithDefaults instantiates a new ContainerDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerDeployRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerDeployRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerDeployRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerDeployRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContainerDeployRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContainerDeployRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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

### HasImageTag

`func (o *ContainerDeployRequest) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### SetImageTagNil

`func (o *ContainerDeployRequest) SetImageTagNil(b bool)`

 SetImageTagNil sets the value for ImageTag to be an explicit nil

### UnsetImageTag
`func (o *ContainerDeployRequest) UnsetImageTag()`

UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


