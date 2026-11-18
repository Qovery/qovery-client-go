# ContainerSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**ContainerSourceRequestOneOfImage**](ContainerSourceRequestOneOfImage.md) |  | 
**GitRepository** | [**ContainerGitRepositorySourceRequest**](ContainerGitRepositorySourceRequest.md) |  | 

## Methods

### NewContainerSourceRequest

`func NewContainerSourceRequest(image ContainerSourceRequestOneOfImage, gitRepository ContainerGitRepositorySourceRequest, ) *ContainerSourceRequest`

NewContainerSourceRequest instantiates a new ContainerSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerSourceRequestWithDefaults

`func NewContainerSourceRequestWithDefaults() *ContainerSourceRequest`

NewContainerSourceRequestWithDefaults instantiates a new ContainerSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ContainerSourceRequest) GetImage() ContainerSourceRequestOneOfImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerSourceRequest) GetImageOk() (*ContainerSourceRequestOneOfImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerSourceRequest) SetImage(v ContainerSourceRequestOneOfImage)`

SetImage sets Image field to given value.


### GetGitRepository

`func (o *ContainerSourceRequest) GetGitRepository() ContainerGitRepositorySourceRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ContainerSourceRequest) GetGitRepositoryOk() (*ContainerGitRepositorySourceRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ContainerSourceRequest) SetGitRepository(v ContainerGitRepositorySourceRequest)`

SetGitRepository sets GitRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


