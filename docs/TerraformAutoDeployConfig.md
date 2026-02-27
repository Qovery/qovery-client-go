# TerraformAutoDeployConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | **bool** |  | 
**AutoDeployAction** | **string** | Action to force a specific Terraform behavior on autodeploy. &#x60;DEFAULT&#x60;: The action is resolved based on the deployment type:   - Start/Restart -&gt; PLAN_AND_APPLY   - Delete -&gt; DESTROY   - Pause -&gt; PLAN_ONLY  | 

## Methods

### NewTerraformAutoDeployConfig

`func NewTerraformAutoDeployConfig(autoDeploy bool, autoDeployAction string, ) *TerraformAutoDeployConfig`

NewTerraformAutoDeployConfig instantiates a new TerraformAutoDeployConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformAutoDeployConfigWithDefaults

`func NewTerraformAutoDeployConfigWithDefaults() *TerraformAutoDeployConfig`

NewTerraformAutoDeployConfigWithDefaults instantiates a new TerraformAutoDeployConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeploy

`func (o *TerraformAutoDeployConfig) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *TerraformAutoDeployConfig) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *TerraformAutoDeployConfig) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.


### GetAutoDeployAction

`func (o *TerraformAutoDeployConfig) GetAutoDeployAction() string`

GetAutoDeployAction returns the AutoDeployAction field if non-nil, zero value otherwise.

### GetAutoDeployActionOk

`func (o *TerraformAutoDeployConfig) GetAutoDeployActionOk() (*string, bool)`

GetAutoDeployActionOk returns a tuple with the AutoDeployAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployAction

`func (o *TerraformAutoDeployConfig) SetAutoDeployAction(v string)`

SetAutoDeployAction sets AutoDeployAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


