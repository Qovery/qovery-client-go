# TerraformDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Terraform service identifier | [optional] 
**GitCommitId** | Pointer to **string** | Commit to deploy for chart source. | [optional] 
**Action** | Pointer to **NullableString** | Terraform action to execute. | [optional] 

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

### GetId

`func (o *TerraformDeployRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerraformDeployRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerraformDeployRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TerraformDeployRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TerraformDeployRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TerraformDeployRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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

### GetAction

`func (o *TerraformDeployRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TerraformDeployRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TerraformDeployRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TerraformDeployRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *TerraformDeployRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *TerraformDeployRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


