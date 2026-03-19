# DeploymentHistoryServiceDetailsOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | [**NullableCommit**](Commit.md) |  | 
**BuildPodName** | **string** | The build pod name prefix for monitoring build runner usage. Format build-{execution_id}-0 | 

## Methods

### NewDeploymentHistoryServiceDetailsOneOf

`func NewDeploymentHistoryServiceDetailsOneOf(commit NullableCommit, buildPodName string, ) *DeploymentHistoryServiceDetailsOneOf`

NewDeploymentHistoryServiceDetailsOneOf instantiates a new DeploymentHistoryServiceDetailsOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServiceDetailsOneOfWithDefaults

`func NewDeploymentHistoryServiceDetailsOneOfWithDefaults() *DeploymentHistoryServiceDetailsOneOf`

NewDeploymentHistoryServiceDetailsOneOfWithDefaults instantiates a new DeploymentHistoryServiceDetailsOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *DeploymentHistoryServiceDetailsOneOf) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryServiceDetailsOneOf) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryServiceDetailsOneOf) SetCommit(v Commit)`

SetCommit sets Commit field to given value.


### SetCommitNil

`func (o *DeploymentHistoryServiceDetailsOneOf) SetCommitNil(b bool)`

 SetCommitNil sets the value for Commit to be an explicit nil

### UnsetCommit
`func (o *DeploymentHistoryServiceDetailsOneOf) UnsetCommit()`

UnsetCommit ensures that no value is present for Commit, not even an explicit nil
### GetBuildPodName

`func (o *DeploymentHistoryServiceDetailsOneOf) GetBuildPodName() string`

GetBuildPodName returns the BuildPodName field if non-nil, zero value otherwise.

### GetBuildPodNameOk

`func (o *DeploymentHistoryServiceDetailsOneOf) GetBuildPodNameOk() (*string, bool)`

GetBuildPodNameOk returns a tuple with the BuildPodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildPodName

`func (o *DeploymentHistoryServiceDetailsOneOf) SetBuildPodName(v string)`

SetBuildPodName sets BuildPodName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


