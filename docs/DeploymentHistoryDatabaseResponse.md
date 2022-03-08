# DeploymentHistoryDatabaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | [optional] 

## Methods

### NewDeploymentHistoryDatabaseResponse

`func NewDeploymentHistoryDatabaseResponse(id string, createdAt time.Time, ) *DeploymentHistoryDatabaseResponse`

NewDeploymentHistoryDatabaseResponse instantiates a new DeploymentHistoryDatabaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryDatabaseResponseWithDefaults

`func NewDeploymentHistoryDatabaseResponseWithDefaults() *DeploymentHistoryDatabaseResponse`

NewDeploymentHistoryDatabaseResponseWithDefaults instantiates a new DeploymentHistoryDatabaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentHistoryDatabaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryDatabaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryDatabaseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryDatabaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryDatabaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryDatabaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryDatabaseResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryDatabaseResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryDatabaseResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryDatabaseResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *DeploymentHistoryDatabaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryDatabaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryDatabaseResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryDatabaseResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryDatabaseResponse) GetStatus() GlobalDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryDatabaseResponse) GetStatusOk() (*GlobalDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryDatabaseResponse) SetStatus(v GlobalDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryDatabaseResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


