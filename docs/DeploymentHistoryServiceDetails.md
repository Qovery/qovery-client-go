# DeploymentHistoryServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | [**Commit**](Commit.md) |  | 
**ImageName** | **string** |  | 
**Tag** | **string** |  | 
**Arguments** | **[]string** |  | 
**Entrypoint** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**DeploymentHistoryServiceDetailsOneOf2Schedule**](DeploymentHistoryServiceDetailsOneOf2Schedule.md) |  | [optional] 
**JobType** | **string** |  | 
**Repository** | Pointer to [**DeploymentHistoryServiceDetailsOneOf3Repository**](DeploymentHistoryServiceDetailsOneOf3Repository.md) |  | [optional] 

## Methods

### NewDeploymentHistoryServiceDetails

`func NewDeploymentHistoryServiceDetails(commit Commit, imageName string, tag string, arguments []string, jobType string, ) *DeploymentHistoryServiceDetails`

NewDeploymentHistoryServiceDetails instantiates a new DeploymentHistoryServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServiceDetailsWithDefaults

`func NewDeploymentHistoryServiceDetailsWithDefaults() *DeploymentHistoryServiceDetails`

NewDeploymentHistoryServiceDetailsWithDefaults instantiates a new DeploymentHistoryServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *DeploymentHistoryServiceDetails) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryServiceDetails) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryServiceDetails) SetCommit(v Commit)`

SetCommit sets Commit field to given value.


### GetImageName

`func (o *DeploymentHistoryServiceDetails) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DeploymentHistoryServiceDetails) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DeploymentHistoryServiceDetails) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *DeploymentHistoryServiceDetails) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeploymentHistoryServiceDetails) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeploymentHistoryServiceDetails) SetTag(v string)`

SetTag sets Tag field to given value.


### GetArguments

`func (o *DeploymentHistoryServiceDetails) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *DeploymentHistoryServiceDetails) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *DeploymentHistoryServiceDetails) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetEntrypoint

`func (o *DeploymentHistoryServiceDetails) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DeploymentHistoryServiceDetails) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DeploymentHistoryServiceDetails) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DeploymentHistoryServiceDetails) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetSchedule

`func (o *DeploymentHistoryServiceDetails) GetSchedule() DeploymentHistoryServiceDetailsOneOf2Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DeploymentHistoryServiceDetails) GetScheduleOk() (*DeploymentHistoryServiceDetailsOneOf2Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DeploymentHistoryServiceDetails) SetSchedule(v DeploymentHistoryServiceDetailsOneOf2Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DeploymentHistoryServiceDetails) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobType

`func (o *DeploymentHistoryServiceDetails) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *DeploymentHistoryServiceDetails) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *DeploymentHistoryServiceDetails) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetRepository

`func (o *DeploymentHistoryServiceDetails) GetRepository() DeploymentHistoryServiceDetailsOneOf3Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DeploymentHistoryServiceDetails) GetRepositoryOk() (*DeploymentHistoryServiceDetailsOneOf3Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DeploymentHistoryServiceDetails) SetRepository(v DeploymentHistoryServiceDetailsOneOf3Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DeploymentHistoryServiceDetails) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


