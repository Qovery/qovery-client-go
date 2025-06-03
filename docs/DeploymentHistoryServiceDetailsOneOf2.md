# DeploymentHistoryServiceDetailsOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | **string** |  | 
**Tag** | **string** |  | 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Schedule** | Pointer to [**DeploymentHistoryServiceDetailsOneOf2Schedule**](DeploymentHistoryServiceDetailsOneOf2Schedule.md) |  | [optional] 
**JobType** | **string** |  | 

## Methods

### NewDeploymentHistoryServiceDetailsOneOf2

`func NewDeploymentHistoryServiceDetailsOneOf2(imageName string, tag string, jobType string, ) *DeploymentHistoryServiceDetailsOneOf2`

NewDeploymentHistoryServiceDetailsOneOf2 instantiates a new DeploymentHistoryServiceDetailsOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServiceDetailsOneOf2WithDefaults

`func NewDeploymentHistoryServiceDetailsOneOf2WithDefaults() *DeploymentHistoryServiceDetailsOneOf2`

NewDeploymentHistoryServiceDetailsOneOf2WithDefaults instantiates a new DeploymentHistoryServiceDetailsOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DeploymentHistoryServiceDetailsOneOf2) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeploymentHistoryServiceDetailsOneOf2) SetTag(v string)`

SetTag sets Tag field to given value.


### GetCommit

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryServiceDetailsOneOf2) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryServiceDetailsOneOf2) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSchedule

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetSchedule() DeploymentHistoryServiceDetailsOneOf2Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetScheduleOk() (*DeploymentHistoryServiceDetailsOneOf2Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DeploymentHistoryServiceDetailsOneOf2) SetSchedule(v DeploymentHistoryServiceDetailsOneOf2Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DeploymentHistoryServiceDetailsOneOf2) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetJobType

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *DeploymentHistoryServiceDetailsOneOf2) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *DeploymentHistoryServiceDetailsOneOf2) SetJobType(v string)`

SetJobType sets JobType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


