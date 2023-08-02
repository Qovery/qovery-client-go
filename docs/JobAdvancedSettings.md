# JobAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildTimeoutMaxSec** | Pointer to **int32** | define the max timeout for the build | [optional] [default to 1800]
**BuildCpuMaxInMilli** | Pointer to **int32** | define the max cpu resources (in milli) | [optional] [default to 4000]
**BuildRamMaxInGib** | Pointer to **int32** | define the max ram resources (in gib) | [optional] [default to 8]
**DeploymentTerminationGracePeriodSeconds** | Pointer to **int32** | define how long in seconds an application is supposed to be stopped gracefully | [optional] [default to 60]
**DeploymentAffinityNodeRequired** | Pointer to **map[string]string** | Set pod placement on specific Kubernetes nodes labels | [optional] 
**JobDeleteTtlSecondsAfterFinished** | Pointer to **NullableInt32** |  | [optional] 
**CronjobConcurrencyPolicy** | Pointer to **string** |  | [optional] [default to "Forbid"]
**CronjobFailedJobsHistoryLimit** | Pointer to **int32** |  | [optional] [default to 1]
**CronjobSuccessJobsHistoryLimit** | Pointer to **int32** |  | [optional] [default to 1]
**SecurityServiceAccountName** | Pointer to **string** | Allows you to set an existing Kubernetes service account name  | [optional] [default to ""]

## Methods

### NewJobAdvancedSettings

`func NewJobAdvancedSettings() *JobAdvancedSettings`

NewJobAdvancedSettings instantiates a new JobAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobAdvancedSettingsWithDefaults

`func NewJobAdvancedSettingsWithDefaults() *JobAdvancedSettings`

NewJobAdvancedSettingsWithDefaults instantiates a new JobAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildTimeoutMaxSec

`func (o *JobAdvancedSettings) GetBuildTimeoutMaxSec() int32`

GetBuildTimeoutMaxSec returns the BuildTimeoutMaxSec field if non-nil, zero value otherwise.

### GetBuildTimeoutMaxSecOk

`func (o *JobAdvancedSettings) GetBuildTimeoutMaxSecOk() (*int32, bool)`

GetBuildTimeoutMaxSecOk returns a tuple with the BuildTimeoutMaxSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimeoutMaxSec

`func (o *JobAdvancedSettings) SetBuildTimeoutMaxSec(v int32)`

SetBuildTimeoutMaxSec sets BuildTimeoutMaxSec field to given value.

### HasBuildTimeoutMaxSec

`func (o *JobAdvancedSettings) HasBuildTimeoutMaxSec() bool`

HasBuildTimeoutMaxSec returns a boolean if a field has been set.

### GetBuildCpuMaxInMilli

`func (o *JobAdvancedSettings) GetBuildCpuMaxInMilli() int32`

GetBuildCpuMaxInMilli returns the BuildCpuMaxInMilli field if non-nil, zero value otherwise.

### GetBuildCpuMaxInMilliOk

`func (o *JobAdvancedSettings) GetBuildCpuMaxInMilliOk() (*int32, bool)`

GetBuildCpuMaxInMilliOk returns a tuple with the BuildCpuMaxInMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCpuMaxInMilli

`func (o *JobAdvancedSettings) SetBuildCpuMaxInMilli(v int32)`

SetBuildCpuMaxInMilli sets BuildCpuMaxInMilli field to given value.

### HasBuildCpuMaxInMilli

`func (o *JobAdvancedSettings) HasBuildCpuMaxInMilli() bool`

HasBuildCpuMaxInMilli returns a boolean if a field has been set.

### GetBuildRamMaxInGib

`func (o *JobAdvancedSettings) GetBuildRamMaxInGib() int32`

GetBuildRamMaxInGib returns the BuildRamMaxInGib field if non-nil, zero value otherwise.

### GetBuildRamMaxInGibOk

`func (o *JobAdvancedSettings) GetBuildRamMaxInGibOk() (*int32, bool)`

GetBuildRamMaxInGibOk returns a tuple with the BuildRamMaxInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRamMaxInGib

`func (o *JobAdvancedSettings) SetBuildRamMaxInGib(v int32)`

SetBuildRamMaxInGib sets BuildRamMaxInGib field to given value.

### HasBuildRamMaxInGib

`func (o *JobAdvancedSettings) HasBuildRamMaxInGib() bool`

HasBuildRamMaxInGib returns a boolean if a field has been set.

### GetDeploymentTerminationGracePeriodSeconds

`func (o *JobAdvancedSettings) GetDeploymentTerminationGracePeriodSeconds() int32`

GetDeploymentTerminationGracePeriodSeconds returns the DeploymentTerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeploymentTerminationGracePeriodSecondsOk

`func (o *JobAdvancedSettings) GetDeploymentTerminationGracePeriodSecondsOk() (*int32, bool)`

GetDeploymentTerminationGracePeriodSecondsOk returns a tuple with the DeploymentTerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTerminationGracePeriodSeconds

`func (o *JobAdvancedSettings) SetDeploymentTerminationGracePeriodSeconds(v int32)`

SetDeploymentTerminationGracePeriodSeconds sets DeploymentTerminationGracePeriodSeconds field to given value.

### HasDeploymentTerminationGracePeriodSeconds

`func (o *JobAdvancedSettings) HasDeploymentTerminationGracePeriodSeconds() bool`

HasDeploymentTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetDeploymentAffinityNodeRequired

`func (o *JobAdvancedSettings) GetDeploymentAffinityNodeRequired() map[string]string`

GetDeploymentAffinityNodeRequired returns the DeploymentAffinityNodeRequired field if non-nil, zero value otherwise.

### GetDeploymentAffinityNodeRequiredOk

`func (o *JobAdvancedSettings) GetDeploymentAffinityNodeRequiredOk() (*map[string]string, bool)`

GetDeploymentAffinityNodeRequiredOk returns a tuple with the DeploymentAffinityNodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentAffinityNodeRequired

`func (o *JobAdvancedSettings) SetDeploymentAffinityNodeRequired(v map[string]string)`

SetDeploymentAffinityNodeRequired sets DeploymentAffinityNodeRequired field to given value.

### HasDeploymentAffinityNodeRequired

`func (o *JobAdvancedSettings) HasDeploymentAffinityNodeRequired() bool`

HasDeploymentAffinityNodeRequired returns a boolean if a field has been set.

### GetJobDeleteTtlSecondsAfterFinished

`func (o *JobAdvancedSettings) GetJobDeleteTtlSecondsAfterFinished() int32`

GetJobDeleteTtlSecondsAfterFinished returns the JobDeleteTtlSecondsAfterFinished field if non-nil, zero value otherwise.

### GetJobDeleteTtlSecondsAfterFinishedOk

`func (o *JobAdvancedSettings) GetJobDeleteTtlSecondsAfterFinishedOk() (*int32, bool)`

GetJobDeleteTtlSecondsAfterFinishedOk returns a tuple with the JobDeleteTtlSecondsAfterFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDeleteTtlSecondsAfterFinished

`func (o *JobAdvancedSettings) SetJobDeleteTtlSecondsAfterFinished(v int32)`

SetJobDeleteTtlSecondsAfterFinished sets JobDeleteTtlSecondsAfterFinished field to given value.

### HasJobDeleteTtlSecondsAfterFinished

`func (o *JobAdvancedSettings) HasJobDeleteTtlSecondsAfterFinished() bool`

HasJobDeleteTtlSecondsAfterFinished returns a boolean if a field has been set.

### SetJobDeleteTtlSecondsAfterFinishedNil

`func (o *JobAdvancedSettings) SetJobDeleteTtlSecondsAfterFinishedNil(b bool)`

 SetJobDeleteTtlSecondsAfterFinishedNil sets the value for JobDeleteTtlSecondsAfterFinished to be an explicit nil

### UnsetJobDeleteTtlSecondsAfterFinished
`func (o *JobAdvancedSettings) UnsetJobDeleteTtlSecondsAfterFinished()`

UnsetJobDeleteTtlSecondsAfterFinished ensures that no value is present for JobDeleteTtlSecondsAfterFinished, not even an explicit nil
### GetCronjobConcurrencyPolicy

`func (o *JobAdvancedSettings) GetCronjobConcurrencyPolicy() string`

GetCronjobConcurrencyPolicy returns the CronjobConcurrencyPolicy field if non-nil, zero value otherwise.

### GetCronjobConcurrencyPolicyOk

`func (o *JobAdvancedSettings) GetCronjobConcurrencyPolicyOk() (*string, bool)`

GetCronjobConcurrencyPolicyOk returns a tuple with the CronjobConcurrencyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronjobConcurrencyPolicy

`func (o *JobAdvancedSettings) SetCronjobConcurrencyPolicy(v string)`

SetCronjobConcurrencyPolicy sets CronjobConcurrencyPolicy field to given value.

### HasCronjobConcurrencyPolicy

`func (o *JobAdvancedSettings) HasCronjobConcurrencyPolicy() bool`

HasCronjobConcurrencyPolicy returns a boolean if a field has been set.

### GetCronjobFailedJobsHistoryLimit

`func (o *JobAdvancedSettings) GetCronjobFailedJobsHistoryLimit() int32`

GetCronjobFailedJobsHistoryLimit returns the CronjobFailedJobsHistoryLimit field if non-nil, zero value otherwise.

### GetCronjobFailedJobsHistoryLimitOk

`func (o *JobAdvancedSettings) GetCronjobFailedJobsHistoryLimitOk() (*int32, bool)`

GetCronjobFailedJobsHistoryLimitOk returns a tuple with the CronjobFailedJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronjobFailedJobsHistoryLimit

`func (o *JobAdvancedSettings) SetCronjobFailedJobsHistoryLimit(v int32)`

SetCronjobFailedJobsHistoryLimit sets CronjobFailedJobsHistoryLimit field to given value.

### HasCronjobFailedJobsHistoryLimit

`func (o *JobAdvancedSettings) HasCronjobFailedJobsHistoryLimit() bool`

HasCronjobFailedJobsHistoryLimit returns a boolean if a field has been set.

### GetCronjobSuccessJobsHistoryLimit

`func (o *JobAdvancedSettings) GetCronjobSuccessJobsHistoryLimit() int32`

GetCronjobSuccessJobsHistoryLimit returns the CronjobSuccessJobsHistoryLimit field if non-nil, zero value otherwise.

### GetCronjobSuccessJobsHistoryLimitOk

`func (o *JobAdvancedSettings) GetCronjobSuccessJobsHistoryLimitOk() (*int32, bool)`

GetCronjobSuccessJobsHistoryLimitOk returns a tuple with the CronjobSuccessJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronjobSuccessJobsHistoryLimit

`func (o *JobAdvancedSettings) SetCronjobSuccessJobsHistoryLimit(v int32)`

SetCronjobSuccessJobsHistoryLimit sets CronjobSuccessJobsHistoryLimit field to given value.

### HasCronjobSuccessJobsHistoryLimit

`func (o *JobAdvancedSettings) HasCronjobSuccessJobsHistoryLimit() bool`

HasCronjobSuccessJobsHistoryLimit returns a boolean if a field has been set.

### GetSecurityServiceAccountName

`func (o *JobAdvancedSettings) GetSecurityServiceAccountName() string`

GetSecurityServiceAccountName returns the SecurityServiceAccountName field if non-nil, zero value otherwise.

### GetSecurityServiceAccountNameOk

`func (o *JobAdvancedSettings) GetSecurityServiceAccountNameOk() (*string, bool)`

GetSecurityServiceAccountNameOk returns a tuple with the SecurityServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServiceAccountName

`func (o *JobAdvancedSettings) SetSecurityServiceAccountName(v string)`

SetSecurityServiceAccountName sets SecurityServiceAccountName field to given value.

### HasSecurityServiceAccountName

`func (o *JobAdvancedSettings) HasSecurityServiceAccountName() bool`

HasSecurityServiceAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


