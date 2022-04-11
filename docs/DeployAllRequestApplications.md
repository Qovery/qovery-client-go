# DeployAllRequestApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | id of the application to be deployed. | 
**GitCommitId** | **string** | Commit ID to deploy. | 

## Methods

### NewDeployAllRequestApplications

`func NewDeployAllRequestApplications(applicationId string, gitCommitId string, ) *DeployAllRequestApplications`

NewDeployAllRequestApplications instantiates a new DeployAllRequestApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestApplicationsWithDefaults

`func NewDeployAllRequestApplicationsWithDefaults() *DeployAllRequestApplications`

NewDeployAllRequestApplicationsWithDefaults instantiates a new DeployAllRequestApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *DeployAllRequestApplications) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DeployAllRequestApplications) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DeployAllRequestApplications) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetGitCommitId

`func (o *DeployAllRequestApplications) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployAllRequestApplications) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployAllRequestApplications) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


