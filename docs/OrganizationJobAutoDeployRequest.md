# OrganizationJobAutoDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** | the job image name to deploy | [optional] 
**Tag** | Pointer to **string** | the new tag to deploy | [optional] 

## Methods

### NewOrganizationJobAutoDeployRequest

`func NewOrganizationJobAutoDeployRequest() *OrganizationJobAutoDeployRequest`

NewOrganizationJobAutoDeployRequest instantiates a new OrganizationJobAutoDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationJobAutoDeployRequestWithDefaults

`func NewOrganizationJobAutoDeployRequestWithDefaults() *OrganizationJobAutoDeployRequest`

NewOrganizationJobAutoDeployRequestWithDefaults instantiates a new OrganizationJobAutoDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *OrganizationJobAutoDeployRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *OrganizationJobAutoDeployRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *OrganizationJobAutoDeployRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *OrganizationJobAutoDeployRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *OrganizationJobAutoDeployRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OrganizationJobAutoDeployRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OrganizationJobAutoDeployRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *OrganizationJobAutoDeployRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


