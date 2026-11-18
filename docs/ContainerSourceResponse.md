# ContainerSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**ContainerSource**](ContainerSource.md) |  | 
**GitRepository** | [**ContainerGitRepositorySourceResponse**](ContainerGitRepositorySourceResponse.md) |  | 

## Methods

### NewContainerSourceResponse

`func NewContainerSourceResponse(image ContainerSource, gitRepository ContainerGitRepositorySourceResponse, ) *ContainerSourceResponse`

NewContainerSourceResponse instantiates a new ContainerSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerSourceResponseWithDefaults

`func NewContainerSourceResponseWithDefaults() *ContainerSourceResponse`

NewContainerSourceResponseWithDefaults instantiates a new ContainerSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ContainerSourceResponse) GetImage() ContainerSource`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerSourceResponse) GetImageOk() (*ContainerSource, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerSourceResponse) SetImage(v ContainerSource)`

SetImage sets Image field to given value.


### GetGitRepository

`func (o *ContainerSourceResponse) GetGitRepository() ContainerGitRepositorySourceResponse`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ContainerSourceResponse) GetGitRepositoryOk() (*ContainerGitRepositorySourceResponse, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ContainerSourceResponse) SetGitRepository(v ContainerGitRepositorySourceResponse)`

SetGitRepository sets GitRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


