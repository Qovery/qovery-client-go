# BaseJobResponseAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**ContainerSource**](ContainerSource.md) |  | [optional] 
**Docker** | Pointer to [**BaseJobResponseAllOfSourceOneOf1Docker**](BaseJobResponseAllOfSourceOneOf1Docker.md) |  | [optional] 

## Methods

### NewBaseJobResponseAllOfSource

`func NewBaseJobResponseAllOfSource() *BaseJobResponseAllOfSource`

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

### HasImage

`func (o *BaseJobResponseAllOfSource) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDocker

`func (o *BaseJobResponseAllOfSource) GetDocker() BaseJobResponseAllOfSourceOneOf1Docker`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *BaseJobResponseAllOfSource) GetDockerOk() (*BaseJobResponseAllOfSourceOneOf1Docker, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *BaseJobResponseAllOfSource) SetDocker(v BaseJobResponseAllOfSourceOneOf1Docker)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *BaseJobResponseAllOfSource) HasDocker() bool`

HasDocker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


