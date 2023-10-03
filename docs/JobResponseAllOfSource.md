# JobResponseAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**NullableJobResponseAllOfSourceImage**](JobResponseAllOfSourceImage.md) |  | [optional] 
**Docker** | Pointer to [**NullableJobResponseAllOfSourceDocker**](JobResponseAllOfSourceDocker.md) |  | [optional] 

## Methods

### NewJobResponseAllOfSource

`func NewJobResponseAllOfSource() *JobResponseAllOfSource`

NewJobResponseAllOfSource instantiates a new JobResponseAllOfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResponseAllOfSourceWithDefaults

`func NewJobResponseAllOfSourceWithDefaults() *JobResponseAllOfSource`

NewJobResponseAllOfSourceWithDefaults instantiates a new JobResponseAllOfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *JobResponseAllOfSource) GetImage() JobResponseAllOfSourceImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *JobResponseAllOfSource) GetImageOk() (*JobResponseAllOfSourceImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *JobResponseAllOfSource) SetImage(v JobResponseAllOfSourceImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *JobResponseAllOfSource) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *JobResponseAllOfSource) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *JobResponseAllOfSource) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetDocker

`func (o *JobResponseAllOfSource) GetDocker() JobResponseAllOfSourceDocker`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *JobResponseAllOfSource) GetDockerOk() (*JobResponseAllOfSourceDocker, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *JobResponseAllOfSource) SetDocker(v JobResponseAllOfSourceDocker)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *JobResponseAllOfSource) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### SetDockerNil

`func (o *JobResponseAllOfSource) SetDockerNil(b bool)`

 SetDockerNil sets the value for Docker to be an explicit nil

### UnsetDocker
`func (o *JobResponseAllOfSource) UnsetDocker()`

UnsetDocker ensures that no value is present for Docker, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


