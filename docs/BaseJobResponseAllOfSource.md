# BaseJobResponseAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**ContainerSource**](ContainerSource.md) |  | 
**Docker** | [**JobSourceDockerResponse**](JobSourceDockerResponse.md) |  | 

## Methods

### NewBaseJobResponseAllOfSource

`func NewBaseJobResponseAllOfSource(image ContainerSource, docker JobSourceDockerResponse, ) *BaseJobResponseAllOfSource`

NewBaseJobResponseAllOfSource instantiates a new BaseJobResponseAllOfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseJobResponseAllOfSourceWithDefaults

`func NewBaseJobResponseAllOfSourceWithDefaults() *BaseJobResponseAllOfSource`

NewBaseJobResponseAllOfSourceWithDefaults instantiates a new BaseJobResponseAllOfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *BaseJobResponseAllOfSource) GetImage() ContainerSource`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BaseJobResponseAllOfSource) GetImageOk() (*ContainerSource, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BaseJobResponseAllOfSource) SetImage(v ContainerSource)`

SetImage sets Image field to given value.


### GetDocker

`func (o *BaseJobResponseAllOfSource) GetDocker() JobSourceDockerResponse`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *BaseJobResponseAllOfSource) GetDockerOk() (*JobSourceDockerResponse, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *BaseJobResponseAllOfSource) SetDocker(v JobSourceDockerResponse)`

SetDocker sets Docker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


