# TerraformRequestDockerfileFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Fragment type discriminator | 
**Path** | **string** | Absolute path to the fragment file. | 
**Content** | **string** | Dockerfile commands to inject (max 8KB). | 

## Methods

### NewTerraformRequestDockerfileFragment

`func NewTerraformRequestDockerfileFragment(type_ string, path string, content string, ) *TerraformRequestDockerfileFragment`

NewTerraformRequestDockerfileFragment instantiates a new TerraformRequestDockerfileFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformRequestDockerfileFragmentWithDefaults

`func NewTerraformRequestDockerfileFragmentWithDefaults() *TerraformRequestDockerfileFragment`

NewTerraformRequestDockerfileFragmentWithDefaults instantiates a new TerraformRequestDockerfileFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TerraformRequestDockerfileFragment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerraformRequestDockerfileFragment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerraformRequestDockerfileFragment) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *TerraformRequestDockerfileFragment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TerraformRequestDockerfileFragment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TerraformRequestDockerfileFragment) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *TerraformRequestDockerfileFragment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TerraformRequestDockerfileFragment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TerraformRequestDockerfileFragment) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


