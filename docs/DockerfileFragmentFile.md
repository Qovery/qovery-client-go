# DockerfileFragmentFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Fragment type discriminator | 
**Path** | **string** | Absolute path to the fragment file. | 

## Methods

### NewDockerfileFragmentFile

`func NewDockerfileFragmentFile(type_ string, path string, ) *DockerfileFragmentFile`

NewDockerfileFragmentFile instantiates a new DockerfileFragmentFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerfileFragmentFileWithDefaults

`func NewDockerfileFragmentFileWithDefaults() *DockerfileFragmentFile`

NewDockerfileFragmentFileWithDefaults instantiates a new DockerfileFragmentFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DockerfileFragmentFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DockerfileFragmentFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DockerfileFragmentFile) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *DockerfileFragmentFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DockerfileFragmentFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DockerfileFragmentFile) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


