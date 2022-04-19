# GitAuthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Owner** | **string** |  | 
**UseBot** | Pointer to **bool** |  | [optional] 

## Methods

### NewGitAuthProvider

`func NewGitAuthProvider(name string, owner string, ) *GitAuthProvider`

NewGitAuthProvider instantiates a new GitAuthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitAuthProviderWithDefaults

`func NewGitAuthProviderWithDefaults() *GitAuthProvider`

NewGitAuthProviderWithDefaults instantiates a new GitAuthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitAuthProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitAuthProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitAuthProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GitAuthProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GitAuthProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitAuthProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitAuthProvider) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *GitAuthProvider) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GitAuthProvider) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GitAuthProvider) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUseBot

`func (o *GitAuthProvider) GetUseBot() bool`

GetUseBot returns the UseBot field if non-nil, zero value otherwise.

### GetUseBotOk

`func (o *GitAuthProvider) GetUseBotOk() (*bool, bool)`

GetUseBotOk returns a tuple with the UseBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBot

`func (o *GitAuthProvider) SetUseBot(v bool)`

SetUseBot sets UseBot field to given value.

### HasUseBot

`func (o *GitAuthProvider) HasUseBot() bool`

HasUseBot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


