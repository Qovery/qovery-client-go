# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Plan** | [**PlanEnum**](PlanEnum.md) |  | 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**AdminEmails** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **string** | uuid of the user owning the organization | [optional] 
**OrganizationPlan** | Pointer to [**OrganizationAllOfOrganizationPlan**](OrganizationAllOfOrganizationPlan.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization(id string, createdAt time.Time, name string, plan PlanEnum, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Organization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Organization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Organization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Organization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Organization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Organization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPlan

`func (o *Organization) GetPlan() PlanEnum`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Organization) GetPlanOk() (*PlanEnum, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Organization) SetPlan(v PlanEnum)`

SetPlan sets Plan field to given value.


### GetWebsiteUrl

`func (o *Organization) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Organization) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Organization) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Organization) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *Organization) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *Organization) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetRepository

`func (o *Organization) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Organization) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Organization) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Organization) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *Organization) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *Organization) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetLogoUrl

`func (o *Organization) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Organization) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Organization) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Organization) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *Organization) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *Organization) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetIconUrl

`func (o *Organization) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Organization) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Organization) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Organization) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *Organization) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *Organization) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetAdminEmails

`func (o *Organization) GetAdminEmails() []string`

GetAdminEmails returns the AdminEmails field if non-nil, zero value otherwise.

### GetAdminEmailsOk

`func (o *Organization) GetAdminEmailsOk() (*[]string, bool)`

GetAdminEmailsOk returns a tuple with the AdminEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmails

`func (o *Organization) SetAdminEmails(v []string)`

SetAdminEmails sets AdminEmails field to given value.

### HasAdminEmails

`func (o *Organization) HasAdminEmails() bool`

HasAdminEmails returns a boolean if a field has been set.

### SetAdminEmailsNil

`func (o *Organization) SetAdminEmailsNil(b bool)`

 SetAdminEmailsNil sets the value for AdminEmails to be an explicit nil

### UnsetAdminEmails
`func (o *Organization) UnsetAdminEmails()`

UnsetAdminEmails ensures that no value is present for AdminEmails, not even an explicit nil
### GetOwner

`func (o *Organization) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Organization) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Organization) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Organization) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOrganizationPlan

`func (o *Organization) GetOrganizationPlan() OrganizationAllOfOrganizationPlan`

GetOrganizationPlan returns the OrganizationPlan field if non-nil, zero value otherwise.

### GetOrganizationPlanOk

`func (o *Organization) GetOrganizationPlanOk() (*OrganizationAllOfOrganizationPlan, bool)`

GetOrganizationPlanOk returns a tuple with the OrganizationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPlan

`func (o *Organization) SetOrganizationPlan(v OrganizationAllOfOrganizationPlan)`

SetOrganizationPlan sets OrganizationPlan field to given value.

### HasOrganizationPlan

`func (o *Organization) HasOrganizationPlan() bool`

HasOrganizationPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


