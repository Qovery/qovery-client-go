# DockerfileFragmentInline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Fragment type discriminator | 
**Content** | **string** | Dockerfile commands to inject (max 8KB). | 

## Methods

### NewDockerfileFragmentInline

`func NewDockerfileFragmentInline(type_ string, content string, ) *DockerfileFragmentInline`

NewDockerfileFragmentInline instantiates a new DockerfileFragmentInline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerfileFragmentInlineWithDefaults

`func NewDockerfileFragmentInlineWithDefaults() *DockerfileFragmentInline`

NewDockerfileFragmentInlineWithDefaults instantiates a new DockerfileFragmentInline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DockerfileFragmentInline) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DockerfileFragmentInline) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DockerfileFragmentInline) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *DockerfileFragmentInline) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DockerfileFragmentInline) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DockerfileFragmentInline) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


