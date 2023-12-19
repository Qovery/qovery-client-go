# DeployAllRequestHelmsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the helm to be updated. | [optional] 
**ChartVersion** | Pointer to **string** | The new chart version for the Helm source. Use this only if the helm has a Helm repository source. | [optional] 
**GitCommitId** | Pointer to **string** | The commit Id to deploy. Use this only if the helm has a Git repository source. | [optional] 
**ValuesOverrideGitCommitId** | Pointer to **string** | The commit Id of the override values to deploy. Use only if the helm has a Git override values repository. | [optional] 

## Methods

### NewDeployAllRequestHelmsInner

`func NewDeployAllRequestHelmsInner() *DeployAllRequestHelmsInner`

NewDeployAllRequestHelmsInner instantiates a new DeployAllRequestHelmsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestHelmsInnerWithDefaults

`func NewDeployAllRequestHelmsInnerWithDefaults() *DeployAllRequestHelmsInner`

NewDeployAllRequestHelmsInnerWithDefaults instantiates a new DeployAllRequestHelmsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeployAllRequestHelmsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeployAllRequestHelmsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeployAllRequestHelmsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeployAllRequestHelmsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChartVersion

`func (o *DeployAllRequestHelmsInner) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *DeployAllRequestHelmsInner) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *DeployAllRequestHelmsInner) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.

### HasChartVersion

`func (o *DeployAllRequestHelmsInner) HasChartVersion() bool`

HasChartVersion returns a boolean if a field has been set.

### GetGitCommitId

`func (o *DeployAllRequestHelmsInner) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployAllRequestHelmsInner) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployAllRequestHelmsInner) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *DeployAllRequestHelmsInner) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.

### GetValuesOverrideGitCommitId

`func (o *DeployAllRequestHelmsInner) GetValuesOverrideGitCommitId() string`

GetValuesOverrideGitCommitId returns the ValuesOverrideGitCommitId field if non-nil, zero value otherwise.

### GetValuesOverrideGitCommitIdOk

`func (o *DeployAllRequestHelmsInner) GetValuesOverrideGitCommitIdOk() (*string, bool)`

GetValuesOverrideGitCommitIdOk returns a tuple with the ValuesOverrideGitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesOverrideGitCommitId

`func (o *DeployAllRequestHelmsInner) SetValuesOverrideGitCommitId(v string)`

SetValuesOverrideGitCommitId sets ValuesOverrideGitCommitId field to given value.

### HasValuesOverrideGitCommitId

`func (o *DeployAllRequestHelmsInner) HasValuesOverrideGitCommitId() bool`

HasValuesOverrideGitCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


