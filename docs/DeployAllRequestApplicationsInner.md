# DeployAllRequestApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | id of the application to be deployed. | 
**GitCommitId** | **string** | Commit ID to deploy. | 

## Methods

### NewDeployAllRequestApplicationsInner

`func NewDeployAllRequestApplicationsInner(applicationId string, gitCommitId string, ) *DeployAllRequestApplicationsInner`

NewDeployAllRequestApplicationsInner instantiates a new DeployAllRequestApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestApplicationsInnerWithDefaults

`func NewDeployAllRequestApplicationsInnerWithDefaults() *DeployAllRequestApplicationsInner`

NewDeployAllRequestApplicationsInnerWithDefaults instantiates a new DeployAllRequestApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *DeployAllRequestApplicationsInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DeployAllRequestApplicationsInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DeployAllRequestApplicationsInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetGitCommitId

`func (o *DeployAllRequestApplicationsInner) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployAllRequestApplicationsInner) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployAllRequestApplicationsInner) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


