# OrganizationEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**AdminEmails** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrganizationEditRequest

`func NewOrganizationEditRequest(name string, ) *OrganizationEditRequest`

NewOrganizationEditRequest instantiates a new OrganizationEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationEditRequestWithDefaults

`func NewOrganizationEditRequestWithDefaults() *OrganizationEditRequest`

NewOrganizationEditRequestWithDefaults instantiates a new OrganizationEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationEditRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationEditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *OrganizationEditRequest) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *OrganizationEditRequest) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *OrganizationEditRequest) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *OrganizationEditRequest) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetRepository

`func (o *OrganizationEditRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OrganizationEditRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OrganizationEditRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *OrganizationEditRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetLogoUrl

`func (o *OrganizationEditRequest) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *OrganizationEditRequest) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *OrganizationEditRequest) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *OrganizationEditRequest) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetIconUrl

`func (o *OrganizationEditRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *OrganizationEditRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *OrganizationEditRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *OrganizationEditRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAdminEmails

`func (o *OrganizationEditRequest) GetAdminEmails() []string`

GetAdminEmails returns the AdminEmails field if non-nil, zero value otherwise.

### GetAdminEmailsOk

`func (o *OrganizationEditRequest) GetAdminEmailsOk() (*[]string, bool)`

GetAdminEmailsOk returns a tuple with the AdminEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmails

`func (o *OrganizationEditRequest) SetAdminEmails(v []string)`

SetAdminEmails sets AdminEmails field to given value.

### HasAdminEmails

`func (o *OrganizationEditRequest) HasAdminEmails() bool`

HasAdminEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


