# JobRequestAllOfSourceImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** | The image name pattern differs according to chosen container registry provider:   * &#x60;ECR&#x60;: &#x60;repository&#x60; * &#x60;SCALEWAY_CR&#x60;: &#x60;namespace/image&#x60; * &#x60;DOCKER_HUB&#x60;: &#x60;image&#x60; or &#x60;repository/image&#x60; * &#x60;PUBLIC_ECR&#x60;: &#x60;registry_alias/repository&#x60;  | [optional] 
**Tag** | Pointer to **string** | tag of the image container | [optional] 
**RegistryId** | Pointer to **string** | tag of the image container | [optional] 

## Methods

### NewJobRequestAllOfSourceImage

`func NewJobRequestAllOfSourceImage() *JobRequestAllOfSourceImage`

NewJobRequestAllOfSourceImage instantiates a new JobRequestAllOfSourceImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestAllOfSourceImageWithDefaults

`func NewJobRequestAllOfSourceImageWithDefaults() *JobRequestAllOfSourceImage`

NewJobRequestAllOfSourceImageWithDefaults instantiates a new JobRequestAllOfSourceImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *JobRequestAllOfSourceImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *JobRequestAllOfSourceImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *JobRequestAllOfSourceImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *JobRequestAllOfSourceImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *JobRequestAllOfSourceImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *JobRequestAllOfSourceImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *JobRequestAllOfSourceImage) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *JobRequestAllOfSourceImage) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetRegistryId

`func (o *JobRequestAllOfSourceImage) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *JobRequestAllOfSourceImage) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *JobRequestAllOfSourceImage) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *JobRequestAllOfSourceImage) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


