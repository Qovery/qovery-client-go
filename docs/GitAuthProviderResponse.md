# GitAuthProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Owner** | **string** |  | 

## Methods

### NewGitAuthProviderResponse

`func NewGitAuthProviderResponse(name string, owner string, ) *GitAuthProviderResponse`

NewGitAuthProviderResponse instantiates a new GitAuthProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitAuthProviderResponseWithDefaults

`func NewGitAuthProviderResponseWithDefaults() *GitAuthProviderResponse`

NewGitAuthProviderResponseWithDefaults instantiates a new GitAuthProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitAuthProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitAuthProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitAuthProviderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GitAuthProviderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GitAuthProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitAuthProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitAuthProviderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *GitAuthProviderResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GitAuthProviderResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GitAuthProviderResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


