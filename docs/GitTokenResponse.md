# GitTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**GitProviderEnum**](GitProviderEnum.md) |  | 
**ExpiredAt** | Pointer to **string** |  | [optional] 
**Workspace** | Pointer to **string** | Mandatory only for BITBUCKET git provider | [optional] 
**AssociatedServicesCount** | **float32** | The number of services using this git token | 
**GitApiUrl** | **string** |  | 

## Methods

### NewGitTokenResponse

`func NewGitTokenResponse(id string, createdAt time.Time, name string, type_ GitProviderEnum, associatedServicesCount float32, gitApiUrl string, ) *GitTokenResponse`

NewGitTokenResponse instantiates a new GitTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTokenResponseWithDefaults

`func NewGitTokenResponseWithDefaults() *GitTokenResponse`

NewGitTokenResponseWithDefaults instantiates a new GitTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitTokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitTokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitTokenResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GitTokenResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitTokenResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitTokenResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GitTokenResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GitTokenResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GitTokenResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GitTokenResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *GitTokenResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitTokenResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitTokenResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GitTokenResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitTokenResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitTokenResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitTokenResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GitTokenResponse) GetType() GitProviderEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitTokenResponse) GetTypeOk() (*GitProviderEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitTokenResponse) SetType(v GitProviderEnum)`

SetType sets Type field to given value.


### GetExpiredAt

`func (o *GitTokenResponse) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *GitTokenResponse) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *GitTokenResponse) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *GitTokenResponse) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetWorkspace

`func (o *GitTokenResponse) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *GitTokenResponse) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *GitTokenResponse) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *GitTokenResponse) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetAssociatedServicesCount

`func (o *GitTokenResponse) GetAssociatedServicesCount() float32`

GetAssociatedServicesCount returns the AssociatedServicesCount field if non-nil, zero value otherwise.

### GetAssociatedServicesCountOk

`func (o *GitTokenResponse) GetAssociatedServicesCountOk() (*float32, bool)`

GetAssociatedServicesCountOk returns a tuple with the AssociatedServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServicesCount

`func (o *GitTokenResponse) SetAssociatedServicesCount(v float32)`

SetAssociatedServicesCount sets AssociatedServicesCount field to given value.


### GetGitApiUrl

`func (o *GitTokenResponse) GetGitApiUrl() string`

GetGitApiUrl returns the GitApiUrl field if non-nil, zero value otherwise.

### GetGitApiUrlOk

`func (o *GitTokenResponse) GetGitApiUrlOk() (*string, bool)`

GetGitApiUrlOk returns a tuple with the GitApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitApiUrl

`func (o *GitTokenResponse) SetGitApiUrl(v string)`

SetGitApiUrl sets GitApiUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


