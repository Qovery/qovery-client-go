# DeployAllRequestContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the container to be updated. | 
**ImageTag** | **string** | new tag for the container. | 

## Methods

### NewDeployAllRequestContainersInner

`func NewDeployAllRequestContainersInner(id string, imageTag string, ) *DeployAllRequestContainersInner`

NewDeployAllRequestContainersInner instantiates a new DeployAllRequestContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestContainersInnerWithDefaults

`func NewDeployAllRequestContainersInnerWithDefaults() *DeployAllRequestContainersInner`

NewDeployAllRequestContainersInnerWithDefaults instantiates a new DeployAllRequestContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeployAllRequestContainersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeployAllRequestContainersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeployAllRequestContainersInner) SetId(v string)`

SetId sets Id field to given value.


### GetImageTag

`func (o *DeployAllRequestContainersInner) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *DeployAllRequestContainersInner) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *DeployAllRequestContainersInner) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


