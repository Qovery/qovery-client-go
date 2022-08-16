# ListContainerDeploymentHistory200ResponseAllOfResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name of the container | [optional] 
**Status** | Pointer to [**DeploymentHistoryStatusEnum**](DeploymentHistoryStatusEnum.md) |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** |  | [optional] 

## Methods

### NewListContainerDeploymentHistory200ResponseAllOfResultsInner

`func NewListContainerDeploymentHistory200ResponseAllOfResultsInner(id string, createdAt time.Time, ) *ListContainerDeploymentHistory200ResponseAllOfResultsInner`

NewListContainerDeploymentHistory200ResponseAllOfResultsInner instantiates a new ListContainerDeploymentHistory200ResponseAllOfResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContainerDeploymentHistory200ResponseAllOfResultsInnerWithDefaults

`func NewListContainerDeploymentHistory200ResponseAllOfResultsInnerWithDefaults() *ListContainerDeploymentHistory200ResponseAllOfResultsInner`

NewListContainerDeploymentHistory200ResponseAllOfResultsInnerWithDefaults instantiates a new ListContainerDeploymentHistory200ResponseAllOfResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetStatus() DeploymentHistoryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetStatusOk() (*DeploymentHistoryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetStatus(v DeploymentHistoryStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageName

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetArguments

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ListContainerDeploymentHistory200ResponseAllOfResultsInner) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


