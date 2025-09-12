# TerraformDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitCommitId** | Pointer to **string** | Commit to deploy for chart source.  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**ForceUnlockState** | Pointer to **bool** |  | [optional] 

## Methods

### NewTerraformDeployRequest

`func NewTerraformDeployRequest() *TerraformDeployRequest`

NewTerraformDeployRequest instantiates a new TerraformDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformDeployRequestWithDefaults

`func NewTerraformDeployRequestWithDefaults() *TerraformDeployRequest`

NewTerraformDeployRequestWithDefaults instantiates a new TerraformDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitCommitId

`func (o *TerraformDeployRequest) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *TerraformDeployRequest) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *TerraformDeployRequest) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *TerraformDeployRequest) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.

### GetDryRun

`func (o *TerraformDeployRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *TerraformDeployRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *TerraformDeployRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *TerraformDeployRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForceUnlockState

`func (o *TerraformDeployRequest) GetForceUnlockState() bool`

GetForceUnlockState returns the ForceUnlockState field if non-nil, zero value otherwise.

### GetForceUnlockStateOk

`func (o *TerraformDeployRequest) GetForceUnlockStateOk() (*bool, bool)`

GetForceUnlockStateOk returns a tuple with the ForceUnlockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnlockState

`func (o *TerraformDeployRequest) SetForceUnlockState(v bool)`

SetForceUnlockState sets ForceUnlockState field to given value.

### HasForceUnlockState

`func (o *TerraformDeployRequest) HasForceUnlockState() bool`

HasForceUnlockState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


