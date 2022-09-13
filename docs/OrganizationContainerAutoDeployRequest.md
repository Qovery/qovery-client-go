# OrganizationContainerAutoDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** | the container image name to deploy | [optional] 
**Tag** | Pointer to **string** | the new tag to deploy | [optional] 

## Methods

### NewOrganizationContainerAutoDeployRequest

`func NewOrganizationContainerAutoDeployRequest() *OrganizationContainerAutoDeployRequest`

NewOrganizationContainerAutoDeployRequest instantiates a new OrganizationContainerAutoDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationContainerAutoDeployRequestWithDefaults

`func NewOrganizationContainerAutoDeployRequestWithDefaults() *OrganizationContainerAutoDeployRequest`

NewOrganizationContainerAutoDeployRequestWithDefaults instantiates a new OrganizationContainerAutoDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *OrganizationContainerAutoDeployRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *OrganizationContainerAutoDeployRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *OrganizationContainerAutoDeployRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *OrganizationContainerAutoDeployRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *OrganizationContainerAutoDeployRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OrganizationContainerAutoDeployRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OrganizationContainerAutoDeployRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *OrganizationContainerAutoDeployRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


