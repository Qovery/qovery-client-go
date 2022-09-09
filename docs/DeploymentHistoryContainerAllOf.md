# DeploymentHistoryContainerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the container | [optional] 
**Status** | Pointer to [**DeploymentHistoryStatusEnum**](DeploymentHistoryStatusEnum.md) |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentHistoryContainerAllOf

`func NewDeploymentHistoryContainerAllOf() *DeploymentHistoryContainerAllOf`

NewDeploymentHistoryContainerAllOf instantiates a new DeploymentHistoryContainerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryContainerAllOfWithDefaults

`func NewDeploymentHistoryContainerAllOfWithDefaults() *DeploymentHistoryContainerAllOf`

NewDeploymentHistoryContainerAllOfWithDefaults instantiates a new DeploymentHistoryContainerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentHistoryContainerAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryContainerAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryContainerAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryContainerAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryContainerAllOf) GetStatus() DeploymentHistoryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryContainerAllOf) GetStatusOk() (*DeploymentHistoryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryContainerAllOf) SetStatus(v DeploymentHistoryStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryContainerAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageName

`func (o *DeploymentHistoryContainerAllOf) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DeploymentHistoryContainerAllOf) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DeploymentHistoryContainerAllOf) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *DeploymentHistoryContainerAllOf) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *DeploymentHistoryContainerAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeploymentHistoryContainerAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeploymentHistoryContainerAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DeploymentHistoryContainerAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetArguments

`func (o *DeploymentHistoryContainerAllOf) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *DeploymentHistoryContainerAllOf) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *DeploymentHistoryContainerAllOf) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *DeploymentHistoryContainerAllOf) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *DeploymentHistoryContainerAllOf) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DeploymentHistoryContainerAllOf) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DeploymentHistoryContainerAllOf) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DeploymentHistoryContainerAllOf) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


