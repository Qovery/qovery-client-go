# CommitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**GitCommitId** | **string** |  | 
**Tag** | **string** |  | 
**Message** | **string** |  | 
**AuthorName** | **string** |  | 
**AuthorAvatarUrl** | Pointer to **string** |  | [optional] 
**CommitPageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCommitResponse

`func NewCommitResponse(createdAt time.Time, gitCommitId string, tag string, message string, authorName string, ) *CommitResponse`

NewCommitResponse instantiates a new CommitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitResponseWithDefaults

`func NewCommitResponseWithDefaults() *CommitResponse`

NewCommitResponseWithDefaults instantiates a new CommitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CommitResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommitResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommitResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGitCommitId

`func (o *CommitResponse) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *CommitResponse) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *CommitResponse) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.


### GetTag

`func (o *CommitResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CommitResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CommitResponse) SetTag(v string)`

SetTag sets Tag field to given value.


### GetMessage

`func (o *CommitResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommitResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommitResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthorName

`func (o *CommitResponse) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *CommitResponse) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *CommitResponse) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetAuthorAvatarUrl

`func (o *CommitResponse) GetAuthorAvatarUrl() string`

GetAuthorAvatarUrl returns the AuthorAvatarUrl field if non-nil, zero value otherwise.

### GetAuthorAvatarUrlOk

`func (o *CommitResponse) GetAuthorAvatarUrlOk() (*string, bool)`

GetAuthorAvatarUrlOk returns a tuple with the AuthorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAvatarUrl

`func (o *CommitResponse) SetAuthorAvatarUrl(v string)`

SetAuthorAvatarUrl sets AuthorAvatarUrl field to given value.

### HasAuthorAvatarUrl

`func (o *CommitResponse) HasAuthorAvatarUrl() bool`

HasAuthorAvatarUrl returns a boolean if a field has been set.

### GetCommitPageUrl

`func (o *CommitResponse) GetCommitPageUrl() string`

GetCommitPageUrl returns the CommitPageUrl field if non-nil, zero value otherwise.

### GetCommitPageUrlOk

`func (o *CommitResponse) GetCommitPageUrlOk() (*string, bool)`

GetCommitPageUrlOk returns a tuple with the CommitPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitPageUrl

`func (o *CommitResponse) SetCommitPageUrl(v string)`

SetCommitPageUrl sets CommitPageUrl field to given value.

### HasCommitPageUrl

`func (o *CommitResponse) HasCommitPageUrl() bool`

HasCommitPageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


