# DeploymentHistoryApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to [**CommitResponse**](CommitResponse.md) |  | [optional] 
**Status** | Pointer to [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | [optional] 

## Methods

### NewDeploymentHistoryApplicationResponse

`func NewDeploymentHistoryApplicationResponse(id string, createdAt time.Time, ) *DeploymentHistoryApplicationResponse`

NewDeploymentHistoryApplicationResponse instantiates a new DeploymentHistoryApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryApplicationResponseWithDefaults

`func NewDeploymentHistoryApplicationResponseWithDefaults() *DeploymentHistoryApplicationResponse`

NewDeploymentHistoryApplicationResponseWithDefaults instantiates a new DeploymentHistoryApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentHistoryApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryApplicationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryApplicationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryApplicationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryApplicationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryApplicationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryApplicationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryApplicationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *DeploymentHistoryApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryApplicationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryApplicationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentHistoryApplicationResponse) GetCommit() CommitResponse`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryApplicationResponse) GetCommitOk() (*CommitResponse, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryApplicationResponse) SetCommit(v CommitResponse)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryApplicationResponse) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryApplicationResponse) GetStatus() GlobalDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryApplicationResponse) GetStatusOk() (*GlobalDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryApplicationResponse) SetStatus(v GlobalDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryApplicationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


