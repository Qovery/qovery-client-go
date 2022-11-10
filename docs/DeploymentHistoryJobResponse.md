# DeploymentHistoryJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name of the job | [optional] 
**Status** | Pointer to [**DeploymentHistoryStatusEnum**](DeploymentHistoryStatusEnum.md) |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentHistoryJobResponse

`func NewDeploymentHistoryJobResponse(id string, createdAt time.Time, ) *DeploymentHistoryJobResponse`

NewDeploymentHistoryJobResponse instantiates a new DeploymentHistoryJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryJobResponseWithDefaults

`func NewDeploymentHistoryJobResponseWithDefaults() *DeploymentHistoryJobResponse`

NewDeploymentHistoryJobResponseWithDefaults instantiates a new DeploymentHistoryJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentHistoryJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryJobResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryJobResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryJobResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryJobResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryJobResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryJobResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryJobResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryJobResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *DeploymentHistoryJobResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryJobResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryJobResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryJobResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryJobResponse) GetStatus() DeploymentHistoryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryJobResponse) GetStatusOk() (*DeploymentHistoryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryJobResponse) SetStatus(v DeploymentHistoryStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryJobResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageName

`func (o *DeploymentHistoryJobResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DeploymentHistoryJobResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DeploymentHistoryJobResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *DeploymentHistoryJobResponse) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *DeploymentHistoryJobResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeploymentHistoryJobResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeploymentHistoryJobResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DeploymentHistoryJobResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetArguments

`func (o *DeploymentHistoryJobResponse) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *DeploymentHistoryJobResponse) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *DeploymentHistoryJobResponse) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *DeploymentHistoryJobResponse) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *DeploymentHistoryJobResponse) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DeploymentHistoryJobResponse) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DeploymentHistoryJobResponse) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DeploymentHistoryJobResponse) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


