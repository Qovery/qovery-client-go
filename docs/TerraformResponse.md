# TerraformResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**TimeoutSec** | **int32** |  | [default to 600]
**AutoApprove** | **bool** |  | 
**AutoDeploy** | **bool** |  | 
**TerraformFilesSource** | Pointer to [**TerraformResponseAllOfTerraformFilesSource**](TerraformResponseAllOfTerraformFilesSource.md) |  | [optional] 
**IconUri** | **string** | Icon URI representing the terraform service. | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**TerraformVariablesSource** | [**TerraformVariablesSourceResponse**](TerraformVariablesSourceResponse.md) |  | 
**Provider** | **string** |  | 
**ProviderVersion** | [**TerraformProviderVersion**](TerraformProviderVersion.md) |  | 
**JobResources** | [**TerraformJobResourcesResponse**](TerraformJobResourcesResponse.md) |  | 
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**UseClusterCredentials** | **bool** |  | 

## Methods

### NewTerraformResponse

`func NewTerraformResponse(id string, createdAt time.Time, name string, timeoutSec int32, autoApprove bool, autoDeploy bool, iconUri string, serviceType ServiceTypeEnum, terraformVariablesSource TerraformVariablesSourceResponse, provider string, providerVersion TerraformProviderVersion, jobResources TerraformJobResourcesResponse, environment ReferenceObject, useClusterCredentials bool, ) *TerraformResponse`

NewTerraformResponse instantiates a new TerraformResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformResponseWithDefaults

`func NewTerraformResponseWithDefaults() *TerraformResponse`

NewTerraformResponseWithDefaults instantiates a new TerraformResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerraformResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerraformResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerraformResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TerraformResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TerraformResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TerraformResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TerraformResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TerraformResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TerraformResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TerraformResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *TerraformResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerraformResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerraformResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TerraformResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TerraformResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TerraformResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TerraformResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *TerraformResponse) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *TerraformResponse) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *TerraformResponse) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.


### GetAutoApprove

`func (o *TerraformResponse) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *TerraformResponse) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *TerraformResponse) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.


### GetAutoDeploy

`func (o *TerraformResponse) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *TerraformResponse) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *TerraformResponse) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.


### GetTerraformFilesSource

`func (o *TerraformResponse) GetTerraformFilesSource() TerraformResponseAllOfTerraformFilesSource`

GetTerraformFilesSource returns the TerraformFilesSource field if non-nil, zero value otherwise.

### GetTerraformFilesSourceOk

`func (o *TerraformResponse) GetTerraformFilesSourceOk() (*TerraformResponseAllOfTerraformFilesSource, bool)`

GetTerraformFilesSourceOk returns a tuple with the TerraformFilesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformFilesSource

`func (o *TerraformResponse) SetTerraformFilesSource(v TerraformResponseAllOfTerraformFilesSource)`

SetTerraformFilesSource sets TerraformFilesSource field to given value.

### HasTerraformFilesSource

`func (o *TerraformResponse) HasTerraformFilesSource() bool`

HasTerraformFilesSource returns a boolean if a field has been set.

### GetIconUri

`func (o *TerraformResponse) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *TerraformResponse) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *TerraformResponse) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.


### GetServiceType

`func (o *TerraformResponse) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *TerraformResponse) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *TerraformResponse) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetTerraformVariablesSource

`func (o *TerraformResponse) GetTerraformVariablesSource() TerraformVariablesSourceResponse`

GetTerraformVariablesSource returns the TerraformVariablesSource field if non-nil, zero value otherwise.

### GetTerraformVariablesSourceOk

`func (o *TerraformResponse) GetTerraformVariablesSourceOk() (*TerraformVariablesSourceResponse, bool)`

GetTerraformVariablesSourceOk returns a tuple with the TerraformVariablesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformVariablesSource

`func (o *TerraformResponse) SetTerraformVariablesSource(v TerraformVariablesSourceResponse)`

SetTerraformVariablesSource sets TerraformVariablesSource field to given value.


### GetProvider

`func (o *TerraformResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TerraformResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TerraformResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderVersion

`func (o *TerraformResponse) GetProviderVersion() TerraformProviderVersion`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *TerraformResponse) GetProviderVersionOk() (*TerraformProviderVersion, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *TerraformResponse) SetProviderVersion(v TerraformProviderVersion)`

SetProviderVersion sets ProviderVersion field to given value.


### GetJobResources

`func (o *TerraformResponse) GetJobResources() TerraformJobResourcesResponse`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *TerraformResponse) GetJobResourcesOk() (*TerraformJobResourcesResponse, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *TerraformResponse) SetJobResources(v TerraformJobResourcesResponse)`

SetJobResources sets JobResources field to given value.


### GetEnvironment

`func (o *TerraformResponse) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TerraformResponse) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TerraformResponse) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetUseClusterCredentials

`func (o *TerraformResponse) GetUseClusterCredentials() bool`

GetUseClusterCredentials returns the UseClusterCredentials field if non-nil, zero value otherwise.

### GetUseClusterCredentialsOk

`func (o *TerraformResponse) GetUseClusterCredentialsOk() (*bool, bool)`

GetUseClusterCredentialsOk returns a tuple with the UseClusterCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseClusterCredentials

`func (o *TerraformResponse) SetUseClusterCredentials(v bool)`

SetUseClusterCredentials sets UseClusterCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


