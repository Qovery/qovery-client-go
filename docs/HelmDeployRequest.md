# HelmDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartVersion** | Pointer to **string** | version of the chart to deploy. Cannot be set if &#x60;git_commit_id&#x60; is defined  | [optional] 
**GitCommitId** | Pointer to **string** | Commit to deploy for chart source. Cannot be set if &#x60;version&#x60; is defined  | [optional] 
**ValuesOverrideGitCommitId** | Pointer to **string** | Commit to deploy for values override  | [optional] 

## Methods

### NewHelmDeployRequest

`func NewHelmDeployRequest() *HelmDeployRequest`

NewHelmDeployRequest instantiates a new HelmDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmDeployRequestWithDefaults

`func NewHelmDeployRequestWithDefaults() *HelmDeployRequest`

NewHelmDeployRequestWithDefaults instantiates a new HelmDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartVersion

`func (o *HelmDeployRequest) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmDeployRequest) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmDeployRequest) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.

### HasChartVersion

`func (o *HelmDeployRequest) HasChartVersion() bool`

HasChartVersion returns a boolean if a field has been set.

### GetGitCommitId

`func (o *HelmDeployRequest) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *HelmDeployRequest) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *HelmDeployRequest) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *HelmDeployRequest) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.

### GetValuesOverrideGitCommitId

`func (o *HelmDeployRequest) GetValuesOverrideGitCommitId() string`

GetValuesOverrideGitCommitId returns the ValuesOverrideGitCommitId field if non-nil, zero value otherwise.

### GetValuesOverrideGitCommitIdOk

`func (o *HelmDeployRequest) GetValuesOverrideGitCommitIdOk() (*string, bool)`

GetValuesOverrideGitCommitIdOk returns a tuple with the ValuesOverrideGitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesOverrideGitCommitId

`func (o *HelmDeployRequest) SetValuesOverrideGitCommitId(v string)`

SetValuesOverrideGitCommitId sets ValuesOverrideGitCommitId field to given value.

### HasValuesOverrideGitCommitId

`func (o *HelmDeployRequest) HasValuesOverrideGitCommitId() bool`

HasValuesOverrideGitCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


