# DeploymentHistoryJobResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the job | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to [**NullableCommit**](Commit.md) |  | [optional] 
**Schedule** | Pointer to [**DeploymentHistoryJobResponseAllOfSchedule**](DeploymentHistoryJobResponseAllOfSchedule.md) |  | [optional] 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentHistoryJobResponseAllOf

`func NewDeploymentHistoryJobResponseAllOf() *DeploymentHistoryJobResponseAllOf`

NewDeploymentHistoryJobResponseAllOf instantiates a new DeploymentHistoryJobResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryJobResponseAllOfWithDefaults

`func NewDeploymentHistoryJobResponseAllOfWithDefaults() *DeploymentHistoryJobResponseAllOf`

NewDeploymentHistoryJobResponseAllOfWithDefaults instantiates a new DeploymentHistoryJobResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentHistoryJobResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryJobResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryJobResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryJobResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryJobResponseAllOf) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryJobResponseAllOf) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryJobResponseAllOf) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryJobResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageName

`func (o *DeploymentHistoryJobResponseAllOf) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DeploymentHistoryJobResponseAllOf) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DeploymentHistoryJobResponseAllOf) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *DeploymentHistoryJobResponseAllOf) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *DeploymentHistoryJobResponseAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeploymentHistoryJobResponseAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeploymentHistoryJobResponseAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DeploymentHistoryJobResponseAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentHistoryJobResponseAllOf) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryJobResponseAllOf) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryJobResponseAllOf) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryJobResponseAllOf) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### SetCommitNil

`func (o *DeploymentHistoryJobResponseAllOf) SetCommitNil(b bool)`

 SetCommitNil sets the value for Commit to be an explicit nil

### UnsetCommit
`func (o *DeploymentHistoryJobResponseAllOf) UnsetCommit()`

UnsetCommit ensures that no value is present for Commit, not even an explicit nil
### GetSchedule

`func (o *DeploymentHistoryJobResponseAllOf) GetSchedule() DeploymentHistoryJobResponseAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DeploymentHistoryJobResponseAllOf) GetScheduleOk() (*DeploymentHistoryJobResponseAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DeploymentHistoryJobResponseAllOf) SetSchedule(v DeploymentHistoryJobResponseAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DeploymentHistoryJobResponseAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetArguments

`func (o *DeploymentHistoryJobResponseAllOf) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *DeploymentHistoryJobResponseAllOf) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *DeploymentHistoryJobResponseAllOf) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *DeploymentHistoryJobResponseAllOf) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *DeploymentHistoryJobResponseAllOf) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DeploymentHistoryJobResponseAllOf) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DeploymentHistoryJobResponseAllOf) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DeploymentHistoryJobResponseAllOf) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


