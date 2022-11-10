# JobRequestAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**NullableJobRequestAllOfSourceImage**](JobRequestAllOfSourceImage.md) |  | [optional] 
**Docker** | Pointer to [**NullableJobRequestAllOfSourceDocker**](JobRequestAllOfSourceDocker.md) |  | [optional] 

## Methods

### NewJobRequestAllOfSource

`func NewJobRequestAllOfSource() *JobRequestAllOfSource`

NewJobRequestAllOfSource instantiates a new JobRequestAllOfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestAllOfSourceWithDefaults

`func NewJobRequestAllOfSourceWithDefaults() *JobRequestAllOfSource`

NewJobRequestAllOfSourceWithDefaults instantiates a new JobRequestAllOfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *JobRequestAllOfSource) GetImage() JobRequestAllOfSourceImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *JobRequestAllOfSource) GetImageOk() (*JobRequestAllOfSourceImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *JobRequestAllOfSource) SetImage(v JobRequestAllOfSourceImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *JobRequestAllOfSource) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *JobRequestAllOfSource) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *JobRequestAllOfSource) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetDocker

`func (o *JobRequestAllOfSource) GetDocker() JobRequestAllOfSourceDocker`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *JobRequestAllOfSource) GetDockerOk() (*JobRequestAllOfSourceDocker, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *JobRequestAllOfSource) SetDocker(v JobRequestAllOfSourceDocker)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *JobRequestAllOfSource) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### SetDockerNil

`func (o *JobRequestAllOfSource) SetDockerNil(b bool)`

 SetDockerNil sets the value for Docker to be an explicit nil

### UnsetDocker
`func (o *JobRequestAllOfSource) UnsetDocker()`

UnsetDocker ensures that no value is present for Docker, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


