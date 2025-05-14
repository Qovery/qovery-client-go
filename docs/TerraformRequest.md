# TerraformRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AutoApprove** | Pointer to **bool** |  | [optional] 
**AutoDeploy** | Pointer to **bool** |  | [optional] 
**TerraformFilesSource** | Pointer to **string** |  | [optional] 
**TerraformVariablesSource** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**ProviderVersion** | Pointer to **string** |  | [optional] 
**TimeoutSec** | Pointer to **string** |  | [optional] 
**IconUri** | Pointer to **string** |  | [optional] 
**JobResources** | Pointer to **string** |  | [optional] 

## Methods

### NewTerraformRequest

`func NewTerraformRequest() *TerraformRequest`

NewTerraformRequest instantiates a new TerraformRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformRequestWithDefaults

`func NewTerraformRequestWithDefaults() *TerraformRequest`

NewTerraformRequestWithDefaults instantiates a new TerraformRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TerraformRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerraformRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerraformRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TerraformRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TerraformRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TerraformRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TerraformRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TerraformRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoApprove

`func (o *TerraformRequest) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *TerraformRequest) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *TerraformRequest) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *TerraformRequest) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *TerraformRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *TerraformRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *TerraformRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *TerraformRequest) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetTerraformFilesSource

`func (o *TerraformRequest) GetTerraformFilesSource() string`

GetTerraformFilesSource returns the TerraformFilesSource field if non-nil, zero value otherwise.

### GetTerraformFilesSourceOk

`func (o *TerraformRequest) GetTerraformFilesSourceOk() (*string, bool)`

GetTerraformFilesSourceOk returns a tuple with the TerraformFilesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformFilesSource

`func (o *TerraformRequest) SetTerraformFilesSource(v string)`

SetTerraformFilesSource sets TerraformFilesSource field to given value.

### HasTerraformFilesSource

`func (o *TerraformRequest) HasTerraformFilesSource() bool`

HasTerraformFilesSource returns a boolean if a field has been set.

### GetTerraformVariablesSource

`func (o *TerraformRequest) GetTerraformVariablesSource() string`

GetTerraformVariablesSource returns the TerraformVariablesSource field if non-nil, zero value otherwise.

### GetTerraformVariablesSourceOk

`func (o *TerraformRequest) GetTerraformVariablesSourceOk() (*string, bool)`

GetTerraformVariablesSourceOk returns a tuple with the TerraformVariablesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformVariablesSource

`func (o *TerraformRequest) SetTerraformVariablesSource(v string)`

SetTerraformVariablesSource sets TerraformVariablesSource field to given value.

### HasTerraformVariablesSource

`func (o *TerraformRequest) HasTerraformVariablesSource() bool`

HasTerraformVariablesSource returns a boolean if a field has been set.

### GetProvider

`func (o *TerraformRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TerraformRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TerraformRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TerraformRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderVersion

`func (o *TerraformRequest) GetProviderVersion() string`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *TerraformRequest) GetProviderVersionOk() (*string, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *TerraformRequest) SetProviderVersion(v string)`

SetProviderVersion sets ProviderVersion field to given value.

### HasProviderVersion

`func (o *TerraformRequest) HasProviderVersion() bool`

HasProviderVersion returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *TerraformRequest) GetTimeoutSec() string`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *TerraformRequest) GetTimeoutSecOk() (*string, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *TerraformRequest) SetTimeoutSec(v string)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *TerraformRequest) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetIconUri

`func (o *TerraformRequest) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *TerraformRequest) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *TerraformRequest) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *TerraformRequest) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.

### GetJobResources

`func (o *TerraformRequest) GetJobResources() string`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *TerraformRequest) GetJobResourcesOk() (*string, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *TerraformRequest) SetJobResources(v string)`

SetJobResources sets JobResources field to given value.

### HasJobResources

`func (o *TerraformRequest) HasJobResources() bool`

HasJobResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


