# BuildSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutMaxSec** | Pointer to **int32** | Maximum build timeout in seconds | [optional] [default to 1800]
**CpuMaxInMilli** | Pointer to **int32** | Maximum CPU resources for the build (in millicores) | [optional] [default to 4000]
**RamMaxInGib** | Pointer to **int32** | Maximum RAM resources for the build (in GiB) | [optional] [default to 8]
**EphemeralStorageInGib** | Pointer to **NullableInt32** | Ephemeral storage for the build (in GiB). When null, the platform default is used. | [optional] 
**DisableBuildkitCache** | Pointer to **bool** | Disable buildkit registry cache during build | [optional] [default to false]
**SkipGitSubmodules** | Pointer to **bool** | Skip git submodules update when cloning the repository | [optional] [default to false]

## Methods

### NewBuildSettings

`func NewBuildSettings() *BuildSettings`

NewBuildSettings instantiates a new BuildSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSettingsWithDefaults

`func NewBuildSettingsWithDefaults() *BuildSettings`

NewBuildSettingsWithDefaults instantiates a new BuildSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutMaxSec

`func (o *BuildSettings) GetTimeoutMaxSec() int32`

GetTimeoutMaxSec returns the TimeoutMaxSec field if non-nil, zero value otherwise.

### GetTimeoutMaxSecOk

`func (o *BuildSettings) GetTimeoutMaxSecOk() (*int32, bool)`

GetTimeoutMaxSecOk returns a tuple with the TimeoutMaxSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMaxSec

`func (o *BuildSettings) SetTimeoutMaxSec(v int32)`

SetTimeoutMaxSec sets TimeoutMaxSec field to given value.

### HasTimeoutMaxSec

`func (o *BuildSettings) HasTimeoutMaxSec() bool`

HasTimeoutMaxSec returns a boolean if a field has been set.

### GetCpuMaxInMilli

`func (o *BuildSettings) GetCpuMaxInMilli() int32`

GetCpuMaxInMilli returns the CpuMaxInMilli field if non-nil, zero value otherwise.

### GetCpuMaxInMilliOk

`func (o *BuildSettings) GetCpuMaxInMilliOk() (*int32, bool)`

GetCpuMaxInMilliOk returns a tuple with the CpuMaxInMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMaxInMilli

`func (o *BuildSettings) SetCpuMaxInMilli(v int32)`

SetCpuMaxInMilli sets CpuMaxInMilli field to given value.

### HasCpuMaxInMilli

`func (o *BuildSettings) HasCpuMaxInMilli() bool`

HasCpuMaxInMilli returns a boolean if a field has been set.

### GetRamMaxInGib

`func (o *BuildSettings) GetRamMaxInGib() int32`

GetRamMaxInGib returns the RamMaxInGib field if non-nil, zero value otherwise.

### GetRamMaxInGibOk

`func (o *BuildSettings) GetRamMaxInGibOk() (*int32, bool)`

GetRamMaxInGibOk returns a tuple with the RamMaxInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMaxInGib

`func (o *BuildSettings) SetRamMaxInGib(v int32)`

SetRamMaxInGib sets RamMaxInGib field to given value.

### HasRamMaxInGib

`func (o *BuildSettings) HasRamMaxInGib() bool`

HasRamMaxInGib returns a boolean if a field has been set.

### GetEphemeralStorageInGib

`func (o *BuildSettings) GetEphemeralStorageInGib() int32`

GetEphemeralStorageInGib returns the EphemeralStorageInGib field if non-nil, zero value otherwise.

### GetEphemeralStorageInGibOk

`func (o *BuildSettings) GetEphemeralStorageInGibOk() (*int32, bool)`

GetEphemeralStorageInGibOk returns a tuple with the EphemeralStorageInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageInGib

`func (o *BuildSettings) SetEphemeralStorageInGib(v int32)`

SetEphemeralStorageInGib sets EphemeralStorageInGib field to given value.

### HasEphemeralStorageInGib

`func (o *BuildSettings) HasEphemeralStorageInGib() bool`

HasEphemeralStorageInGib returns a boolean if a field has been set.

### SetEphemeralStorageInGibNil

`func (o *BuildSettings) SetEphemeralStorageInGibNil(b bool)`

 SetEphemeralStorageInGibNil sets the value for EphemeralStorageInGib to be an explicit nil

### UnsetEphemeralStorageInGib
`func (o *BuildSettings) UnsetEphemeralStorageInGib()`

UnsetEphemeralStorageInGib ensures that no value is present for EphemeralStorageInGib, not even an explicit nil
### GetDisableBuildkitCache

`func (o *BuildSettings) GetDisableBuildkitCache() bool`

GetDisableBuildkitCache returns the DisableBuildkitCache field if non-nil, zero value otherwise.

### GetDisableBuildkitCacheOk

`func (o *BuildSettings) GetDisableBuildkitCacheOk() (*bool, bool)`

GetDisableBuildkitCacheOk returns a tuple with the DisableBuildkitCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBuildkitCache

`func (o *BuildSettings) SetDisableBuildkitCache(v bool)`

SetDisableBuildkitCache sets DisableBuildkitCache field to given value.

### HasDisableBuildkitCache

`func (o *BuildSettings) HasDisableBuildkitCache() bool`

HasDisableBuildkitCache returns a boolean if a field has been set.

### GetSkipGitSubmodules

`func (o *BuildSettings) GetSkipGitSubmodules() bool`

GetSkipGitSubmodules returns the SkipGitSubmodules field if non-nil, zero value otherwise.

### GetSkipGitSubmodulesOk

`func (o *BuildSettings) GetSkipGitSubmodulesOk() (*bool, bool)`

GetSkipGitSubmodulesOk returns a tuple with the SkipGitSubmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGitSubmodules

`func (o *BuildSettings) SetSkipGitSubmodules(v bool)`

SetSkipGitSubmodules sets SkipGitSubmodules field to given value.

### HasSkipGitSubmodules

`func (o *BuildSettings) HasSkipGitSubmodules() bool`

HasSkipGitSubmodules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


