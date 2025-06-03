# OrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Plan** | [**PlanEnum**](PlanEnum.md) |  | 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**AdminEmails** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrganizationRequest

`func NewOrganizationRequest(name string, plan PlanEnum, ) *OrganizationRequest`

NewOrganizationRequest instantiates a new OrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationRequestWithDefaults

`func NewOrganizationRequestWithDefaults() *OrganizationRequest`

NewOrganizationRequestWithDefaults instantiates a new OrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlan

`func (o *OrganizationRequest) GetPlan() PlanEnum`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationRequest) GetPlanOk() (*PlanEnum, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationRequest) SetPlan(v PlanEnum)`

SetPlan sets Plan field to given value.


### GetWebsiteUrl

`func (o *OrganizationRequest) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *OrganizationRequest) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *OrganizationRequest) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *OrganizationRequest) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetRepository

`func (o *OrganizationRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OrganizationRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OrganizationRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *OrganizationRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetLogoUrl

`func (o *OrganizationRequest) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *OrganizationRequest) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *OrganizationRequest) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *OrganizationRequest) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetIconUrl

`func (o *OrganizationRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *OrganizationRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *OrganizationRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *OrganizationRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAdminEmails

`func (o *OrganizationRequest) GetAdminEmails() []string`

GetAdminEmails returns the AdminEmails field if non-nil, zero value otherwise.

### GetAdminEmailsOk

`func (o *OrganizationRequest) GetAdminEmailsOk() (*[]string, bool)`

GetAdminEmailsOk returns a tuple with the AdminEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmails

`func (o *OrganizationRequest) SetAdminEmails(v []string)`

SetAdminEmails sets AdminEmails field to given value.

### HasAdminEmails

`func (o *OrganizationRequest) HasAdminEmails() bool`

HasAdminEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


