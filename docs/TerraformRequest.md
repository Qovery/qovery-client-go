# TerraformRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**AutoDeploy** | **bool** |  | 
**TerraformFilesSource** | [**TerraformRequestTerraformFilesSource**](TerraformRequestTerraformFilesSource.md) |  | 
**TerraformVariablesSource** | [**TerraformVariablesSourceRequest**](TerraformVariablesSourceRequest.md) |  | 
**Backend** | [**TerraformBackend**](TerraformBackend.md) |  | 
**Provider** | **string** |  | 
**ProviderVersion** | [**TerraformProviderVersion**](TerraformProviderVersion.md) |  | 
**TimeoutSec** | Pointer to **int32** |  | [optional] 
**IconUri** | Pointer to **string** |  | [optional] 
**JobResources** | [**TerraformRequestJobResources**](TerraformRequestJobResources.md) |  | 
**UseClusterCredentials** | Pointer to **bool** |  | [optional] 
**ActionExtraArguments** | Pointer to **map[string][]string** | The key represent the action command name i.e: \&quot;plan\&quot; The value represent the extra arguments to pass to this command  i.e: {\&quot;apply\&quot;, [\&quot;-lock&#x3D;false\&quot;]} is going to prepend &#x60;-lock&#x3D;false&#x60; to terraform apply commands | [optional] 

## Methods

### NewTerraformRequest

`func NewTerraformRequest(name string, description string, autoDeploy bool, terraformFilesSource TerraformRequestTerraformFilesSource, terraformVariablesSource TerraformVariablesSourceRequest, backend TerraformBackend, provider string, providerVersion TerraformProviderVersion, jobResources TerraformRequestJobResources, ) *TerraformRequest`

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


### GetTerraformFilesSource

`func (o *TerraformRequest) GetTerraformFilesSource() TerraformRequestTerraformFilesSource`

GetTerraformFilesSource returns the TerraformFilesSource field if non-nil, zero value otherwise.

### GetTerraformFilesSourceOk

`func (o *TerraformRequest) GetTerraformFilesSourceOk() (*TerraformRequestTerraformFilesSource, bool)`

GetTerraformFilesSourceOk returns a tuple with the TerraformFilesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformFilesSource

`func (o *TerraformRequest) SetTerraformFilesSource(v TerraformRequestTerraformFilesSource)`

SetTerraformFilesSource sets TerraformFilesSource field to given value.


### GetTerraformVariablesSource

`func (o *TerraformRequest) GetTerraformVariablesSource() TerraformVariablesSourceRequest`

GetTerraformVariablesSource returns the TerraformVariablesSource field if non-nil, zero value otherwise.

### GetTerraformVariablesSourceOk

`func (o *TerraformRequest) GetTerraformVariablesSourceOk() (*TerraformVariablesSourceRequest, bool)`

GetTerraformVariablesSourceOk returns a tuple with the TerraformVariablesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformVariablesSource

`func (o *TerraformRequest) SetTerraformVariablesSource(v TerraformVariablesSourceRequest)`

SetTerraformVariablesSource sets TerraformVariablesSource field to given value.


### GetBackend

`func (o *TerraformRequest) GetBackend() TerraformBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *TerraformRequest) GetBackendOk() (*TerraformBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *TerraformRequest) SetBackend(v TerraformBackend)`

SetBackend sets Backend field to given value.


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


### GetProviderVersion

`func (o *TerraformRequest) GetProviderVersion() TerraformProviderVersion`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *TerraformRequest) GetProviderVersionOk() (*TerraformProviderVersion, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *TerraformRequest) SetProviderVersion(v TerraformProviderVersion)`

SetProviderVersion sets ProviderVersion field to given value.


### GetTimeoutSec

`func (o *TerraformRequest) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *TerraformRequest) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *TerraformRequest) SetTimeoutSec(v int32)`

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

`func (o *TerraformRequest) GetJobResources() TerraformRequestJobResources`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *TerraformRequest) GetJobResourcesOk() (*TerraformRequestJobResources, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *TerraformRequest) SetJobResources(v TerraformRequestJobResources)`

SetJobResources sets JobResources field to given value.


### GetUseClusterCredentials

`func (o *TerraformRequest) GetUseClusterCredentials() bool`

GetUseClusterCredentials returns the UseClusterCredentials field if non-nil, zero value otherwise.

### GetUseClusterCredentialsOk

`func (o *TerraformRequest) GetUseClusterCredentialsOk() (*bool, bool)`

GetUseClusterCredentialsOk returns a tuple with the UseClusterCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseClusterCredentials

`func (o *TerraformRequest) SetUseClusterCredentials(v bool)`

SetUseClusterCredentials sets UseClusterCredentials field to given value.

### HasUseClusterCredentials

`func (o *TerraformRequest) HasUseClusterCredentials() bool`

HasUseClusterCredentials returns a boolean if a field has been set.

### GetActionExtraArguments

`func (o *TerraformRequest) GetActionExtraArguments() map[string][]string`

GetActionExtraArguments returns the ActionExtraArguments field if non-nil, zero value otherwise.

### GetActionExtraArgumentsOk

`func (o *TerraformRequest) GetActionExtraArgumentsOk() (*map[string][]string, bool)`

GetActionExtraArgumentsOk returns a tuple with the ActionExtraArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionExtraArguments

`func (o *TerraformRequest) SetActionExtraArguments(v map[string][]string)`

SetActionExtraArguments sets ActionExtraArguments field to given value.

### HasActionExtraArguments

`func (o *TerraformRequest) HasActionExtraArguments() bool`

HasActionExtraArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


