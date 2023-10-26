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

## Methods

### NewGitTokenResponse

`func NewGitTokenResponse(id string, createdAt time.Time, name string, type_ GitProviderEnum, ) *GitTokenResponse`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


