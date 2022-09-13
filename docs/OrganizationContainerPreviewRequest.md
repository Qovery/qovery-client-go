# OrganizationContainerPreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** | the container image name to trigger preview environment | [optional] 
**Tag** | Pointer to **string** | the tag to be used in the preview environment | [optional] 

## Methods

### NewOrganizationContainerPreviewRequest

`func NewOrganizationContainerPreviewRequest() *OrganizationContainerPreviewRequest`

NewOrganizationContainerPreviewRequest instantiates a new OrganizationContainerPreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationContainerPreviewRequestWithDefaults

`func NewOrganizationContainerPreviewRequestWithDefaults() *OrganizationContainerPreviewRequest`

NewOrganizationContainerPreviewRequestWithDefaults instantiates a new OrganizationContainerPreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *OrganizationContainerPreviewRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *OrganizationContainerPreviewRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *OrganizationContainerPreviewRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *OrganizationContainerPreviewRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *OrganizationContainerPreviewRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OrganizationContainerPreviewRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OrganizationContainerPreviewRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *OrganizationContainerPreviewRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


