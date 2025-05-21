# TerraformAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildTimeoutMaxSec** | Pointer to **int32** | define the max timeout for the build | [optional] 
**BuildCpuMaxInMilli** | Pointer to **int32** | define the max cpu resources (in milli) | [optional] 
**BuildRamMaxInGib** | Pointer to **int32** | define the max ram resources (in gib) | [optional] 
**BuildEphemeralStorageInGib** | Pointer to **int32** |  | [optional] 
**DeploymentTerminationGracePeriodSeconds** | Pointer to **int32** | define how long in seconds an application is supposed to be stopped gracefully | [optional] 
**DeploymentAffinityNodeRequired** | Pointer to **map[string]string** | Set pod placement on specific Kubernetes nodes labels | [optional] 
**SecurityServiceAccountName** | Pointer to **string** | Allows you to set an existing Kubernetes service account name  | [optional] 
**SecurityReadOnlyRootFilesystem** | Pointer to **bool** | Mounts the container&#39;s root filesystem as read-only  | [optional] 

## Methods

### NewTerraformAdvancedSettings

`func NewTerraformAdvancedSettings() *TerraformAdvancedSettings`

NewTerraformAdvancedSettings instantiates a new TerraformAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformAdvancedSettingsWithDefaults

`func NewTerraformAdvancedSettingsWithDefaults() *TerraformAdvancedSettings`

NewTerraformAdvancedSettingsWithDefaults instantiates a new TerraformAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildTimeoutMaxSec

`func (o *TerraformAdvancedSettings) GetBuildTimeoutMaxSec() int32`

GetBuildTimeoutMaxSec returns the BuildTimeoutMaxSec field if non-nil, zero value otherwise.

### GetBuildTimeoutMaxSecOk

`func (o *TerraformAdvancedSettings) GetBuildTimeoutMaxSecOk() (*int32, bool)`

GetBuildTimeoutMaxSecOk returns a tuple with the BuildTimeoutMaxSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimeoutMaxSec

`func (o *TerraformAdvancedSettings) SetBuildTimeoutMaxSec(v int32)`

SetBuildTimeoutMaxSec sets BuildTimeoutMaxSec field to given value.

### HasBuildTimeoutMaxSec

`func (o *TerraformAdvancedSettings) HasBuildTimeoutMaxSec() bool`

HasBuildTimeoutMaxSec returns a boolean if a field has been set.

### GetBuildCpuMaxInMilli

`func (o *TerraformAdvancedSettings) GetBuildCpuMaxInMilli() int32`

GetBuildCpuMaxInMilli returns the BuildCpuMaxInMilli field if non-nil, zero value otherwise.

### GetBuildCpuMaxInMilliOk

`func (o *TerraformAdvancedSettings) GetBuildCpuMaxInMilliOk() (*int32, bool)`

GetBuildCpuMaxInMilliOk returns a tuple with the BuildCpuMaxInMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCpuMaxInMilli

`func (o *TerraformAdvancedSettings) SetBuildCpuMaxInMilli(v int32)`

SetBuildCpuMaxInMilli sets BuildCpuMaxInMilli field to given value.

### HasBuildCpuMaxInMilli

`func (o *TerraformAdvancedSettings) HasBuildCpuMaxInMilli() bool`

HasBuildCpuMaxInMilli returns a boolean if a field has been set.

### GetBuildRamMaxInGib

`func (o *TerraformAdvancedSettings) GetBuildRamMaxInGib() int32`

GetBuildRamMaxInGib returns the BuildRamMaxInGib field if non-nil, zero value otherwise.

### GetBuildRamMaxInGibOk

`func (o *TerraformAdvancedSettings) GetBuildRamMaxInGibOk() (*int32, bool)`

GetBuildRamMaxInGibOk returns a tuple with the BuildRamMaxInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRamMaxInGib

`func (o *TerraformAdvancedSettings) SetBuildRamMaxInGib(v int32)`

SetBuildRamMaxInGib sets BuildRamMaxInGib field to given value.

### HasBuildRamMaxInGib

`func (o *TerraformAdvancedSettings) HasBuildRamMaxInGib() bool`

HasBuildRamMaxInGib returns a boolean if a field has been set.

### GetBuildEphemeralStorageInGib

`func (o *TerraformAdvancedSettings) GetBuildEphemeralStorageInGib() int32`

GetBuildEphemeralStorageInGib returns the BuildEphemeralStorageInGib field if non-nil, zero value otherwise.

### GetBuildEphemeralStorageInGibOk

`func (o *TerraformAdvancedSettings) GetBuildEphemeralStorageInGibOk() (*int32, bool)`

GetBuildEphemeralStorageInGibOk returns a tuple with the BuildEphemeralStorageInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEphemeralStorageInGib

`func (o *TerraformAdvancedSettings) SetBuildEphemeralStorageInGib(v int32)`

SetBuildEphemeralStorageInGib sets BuildEphemeralStorageInGib field to given value.

### HasBuildEphemeralStorageInGib

`func (o *TerraformAdvancedSettings) HasBuildEphemeralStorageInGib() bool`

HasBuildEphemeralStorageInGib returns a boolean if a field has been set.

### GetDeploymentTerminationGracePeriodSeconds

`func (o *TerraformAdvancedSettings) GetDeploymentTerminationGracePeriodSeconds() int32`

GetDeploymentTerminationGracePeriodSeconds returns the DeploymentTerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeploymentTerminationGracePeriodSecondsOk

`func (o *TerraformAdvancedSettings) GetDeploymentTerminationGracePeriodSecondsOk() (*int32, bool)`

GetDeploymentTerminationGracePeriodSecondsOk returns a tuple with the DeploymentTerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTerminationGracePeriodSeconds

`func (o *TerraformAdvancedSettings) SetDeploymentTerminationGracePeriodSeconds(v int32)`

SetDeploymentTerminationGracePeriodSeconds sets DeploymentTerminationGracePeriodSeconds field to given value.

### HasDeploymentTerminationGracePeriodSeconds

`func (o *TerraformAdvancedSettings) HasDeploymentTerminationGracePeriodSeconds() bool`

HasDeploymentTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetDeploymentAffinityNodeRequired

`func (o *TerraformAdvancedSettings) GetDeploymentAffinityNodeRequired() map[string]string`

GetDeploymentAffinityNodeRequired returns the DeploymentAffinityNodeRequired field if non-nil, zero value otherwise.

### GetDeploymentAffinityNodeRequiredOk

`func (o *TerraformAdvancedSettings) GetDeploymentAffinityNodeRequiredOk() (*map[string]string, bool)`

GetDeploymentAffinityNodeRequiredOk returns a tuple with the DeploymentAffinityNodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentAffinityNodeRequired

`func (o *TerraformAdvancedSettings) SetDeploymentAffinityNodeRequired(v map[string]string)`

SetDeploymentAffinityNodeRequired sets DeploymentAffinityNodeRequired field to given value.

### HasDeploymentAffinityNodeRequired

`func (o *TerraformAdvancedSettings) HasDeploymentAffinityNodeRequired() bool`

HasDeploymentAffinityNodeRequired returns a boolean if a field has been set.

### GetSecurityServiceAccountName

`func (o *TerraformAdvancedSettings) GetSecurityServiceAccountName() string`

GetSecurityServiceAccountName returns the SecurityServiceAccountName field if non-nil, zero value otherwise.

### GetSecurityServiceAccountNameOk

`func (o *TerraformAdvancedSettings) GetSecurityServiceAccountNameOk() (*string, bool)`

GetSecurityServiceAccountNameOk returns a tuple with the SecurityServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServiceAccountName

`func (o *TerraformAdvancedSettings) SetSecurityServiceAccountName(v string)`

SetSecurityServiceAccountName sets SecurityServiceAccountName field to given value.

### HasSecurityServiceAccountName

`func (o *TerraformAdvancedSettings) HasSecurityServiceAccountName() bool`

HasSecurityServiceAccountName returns a boolean if a field has been set.

### GetSecurityReadOnlyRootFilesystem

`func (o *TerraformAdvancedSettings) GetSecurityReadOnlyRootFilesystem() bool`

GetSecurityReadOnlyRootFilesystem returns the SecurityReadOnlyRootFilesystem field if non-nil, zero value otherwise.

### GetSecurityReadOnlyRootFilesystemOk

`func (o *TerraformAdvancedSettings) GetSecurityReadOnlyRootFilesystemOk() (*bool, bool)`

GetSecurityReadOnlyRootFilesystemOk returns a tuple with the SecurityReadOnlyRootFilesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReadOnlyRootFilesystem

`func (o *TerraformAdvancedSettings) SetSecurityReadOnlyRootFilesystem(v bool)`

SetSecurityReadOnlyRootFilesystem sets SecurityReadOnlyRootFilesystem field to given value.

### HasSecurityReadOnlyRootFilesystem

`func (o *TerraformAdvancedSettings) HasSecurityReadOnlyRootFilesystem() bool`

HasSecurityReadOnlyRootFilesystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


