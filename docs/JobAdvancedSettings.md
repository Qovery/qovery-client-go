# JobAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobDeleteTtlSecondsAfterFinished** | Pointer to **NullableInt32** |  | [optional] 
**CronjobConcurrencyPolicy** | Pointer to **string** |  | [optional] [default to "Forbid"]
**CronjobFailedJobsHistoryLimit** | Pointer to **int32** |  | [optional] [default to 1]
**CronjobSuccessJobsHistoryLimit** | Pointer to **int32** |  | [optional] [default to 1]
**ReadinessProbeType** | Pointer to **string** | &#x60;NONE&#x60; disable readiness probe &#x60;TCP&#x60; enable TCP readiness probe &#x60;HTTP&#x60; enable HTTP readiness probe  | [optional] [default to "NONE"]
**ReadinessProbeHttpGetPath** | Pointer to **string** | HTTP GET path to check status (must returns 2xx E.g \&quot;/healtz\&quot;) - only usable with TYPE &#x3D; HTTP | [optional] [default to ""]
**ReadinessProbeInitialDelaySeconds** | Pointer to **int32** | Delay before liveness probe is initiated | [optional] [default to 0]
**ReadinessProbePeriodSeconds** | Pointer to **int32** | How often to perform the probe | [optional] [default to 0]
**ReadinessProbeTimeoutSeconds** | Pointer to **int32** | When the probe times out | [optional] [default to 0]
**ReadinessProbeSuccessThreshold** | Pointer to **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. | [optional] [default to 0]
**ReadinessProbeFailureThreshold** | Pointer to **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. | [optional] [default to 0]
**LivenessProbeType** | Pointer to **string** | &#x60;NONE&#x60; disable liveness probe &#x60;TCP&#x60; enable TCP liveness probe &#x60;HTTP&#x60; enable HTTP liveness probe  | [optional] [default to "NONE"]
**LivenessProbeHttpGetPath** | Pointer to **string** | HTTP GET path to check status (must returns 2xx E.g \&quot;/healtz\&quot;) - only usable with TYPE &#x3D; HTTP | [optional] [default to ""]
**LivenessProbeInitialDelaySeconds** | Pointer to **int32** | Delay before liveness probe is initiated | [optional] [default to 0]
**LivenessProbePeriodSeconds** | Pointer to **int32** | How often to perform the probe | [optional] [default to 0]
**LivenessProbeTimeoutSeconds** | Pointer to **int32** | When the probe times out | [optional] [default to 0]
**LivenessProbeSuccessThreshold** | Pointer to **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. | [optional] [default to 0]
**LivenessProbeFailureThreshold** | Pointer to **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. | [optional] [default to 0]

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

### GetReadinessProbeType

`func (o *JobAdvancedSettings) GetReadinessProbeType() string`

GetReadinessProbeType returns the ReadinessProbeType field if non-nil, zero value otherwise.

### GetReadinessProbeTypeOk

`func (o *JobAdvancedSettings) GetReadinessProbeTypeOk() (*string, bool)`

GetReadinessProbeTypeOk returns a tuple with the ReadinessProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeType

`func (o *JobAdvancedSettings) SetReadinessProbeType(v string)`

SetReadinessProbeType sets ReadinessProbeType field to given value.

### HasReadinessProbeType

`func (o *JobAdvancedSettings) HasReadinessProbeType() bool`

HasReadinessProbeType returns a boolean if a field has been set.

### GetReadinessProbeHttpGetPath

`func (o *JobAdvancedSettings) GetReadinessProbeHttpGetPath() string`

GetReadinessProbeHttpGetPath returns the ReadinessProbeHttpGetPath field if non-nil, zero value otherwise.

### GetReadinessProbeHttpGetPathOk

`func (o *JobAdvancedSettings) GetReadinessProbeHttpGetPathOk() (*string, bool)`

GetReadinessProbeHttpGetPathOk returns a tuple with the ReadinessProbeHttpGetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeHttpGetPath

`func (o *JobAdvancedSettings) SetReadinessProbeHttpGetPath(v string)`

SetReadinessProbeHttpGetPath sets ReadinessProbeHttpGetPath field to given value.

### HasReadinessProbeHttpGetPath

`func (o *JobAdvancedSettings) HasReadinessProbeHttpGetPath() bool`

HasReadinessProbeHttpGetPath returns a boolean if a field has been set.

### GetReadinessProbeInitialDelaySeconds

`func (o *JobAdvancedSettings) GetReadinessProbeInitialDelaySeconds() int32`

GetReadinessProbeInitialDelaySeconds returns the ReadinessProbeInitialDelaySeconds field if non-nil, zero value otherwise.

### GetReadinessProbeInitialDelaySecondsOk

`func (o *JobAdvancedSettings) GetReadinessProbeInitialDelaySecondsOk() (*int32, bool)`

GetReadinessProbeInitialDelaySecondsOk returns a tuple with the ReadinessProbeInitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeInitialDelaySeconds

`func (o *JobAdvancedSettings) SetReadinessProbeInitialDelaySeconds(v int32)`

SetReadinessProbeInitialDelaySeconds sets ReadinessProbeInitialDelaySeconds field to given value.

### HasReadinessProbeInitialDelaySeconds

`func (o *JobAdvancedSettings) HasReadinessProbeInitialDelaySeconds() bool`

HasReadinessProbeInitialDelaySeconds returns a boolean if a field has been set.

### GetReadinessProbePeriodSeconds

`func (o *JobAdvancedSettings) GetReadinessProbePeriodSeconds() int32`

GetReadinessProbePeriodSeconds returns the ReadinessProbePeriodSeconds field if non-nil, zero value otherwise.

### GetReadinessProbePeriodSecondsOk

`func (o *JobAdvancedSettings) GetReadinessProbePeriodSecondsOk() (*int32, bool)`

GetReadinessProbePeriodSecondsOk returns a tuple with the ReadinessProbePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbePeriodSeconds

`func (o *JobAdvancedSettings) SetReadinessProbePeriodSeconds(v int32)`

SetReadinessProbePeriodSeconds sets ReadinessProbePeriodSeconds field to given value.

### HasReadinessProbePeriodSeconds

`func (o *JobAdvancedSettings) HasReadinessProbePeriodSeconds() bool`

HasReadinessProbePeriodSeconds returns a boolean if a field has been set.

### GetReadinessProbeTimeoutSeconds

`func (o *JobAdvancedSettings) GetReadinessProbeTimeoutSeconds() int32`

GetReadinessProbeTimeoutSeconds returns the ReadinessProbeTimeoutSeconds field if non-nil, zero value otherwise.

### GetReadinessProbeTimeoutSecondsOk

`func (o *JobAdvancedSettings) GetReadinessProbeTimeoutSecondsOk() (*int32, bool)`

GetReadinessProbeTimeoutSecondsOk returns a tuple with the ReadinessProbeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeTimeoutSeconds

`func (o *JobAdvancedSettings) SetReadinessProbeTimeoutSeconds(v int32)`

SetReadinessProbeTimeoutSeconds sets ReadinessProbeTimeoutSeconds field to given value.

### HasReadinessProbeTimeoutSeconds

`func (o *JobAdvancedSettings) HasReadinessProbeTimeoutSeconds() bool`

HasReadinessProbeTimeoutSeconds returns a boolean if a field has been set.

### GetReadinessProbeSuccessThreshold

`func (o *JobAdvancedSettings) GetReadinessProbeSuccessThreshold() int32`

GetReadinessProbeSuccessThreshold returns the ReadinessProbeSuccessThreshold field if non-nil, zero value otherwise.

### GetReadinessProbeSuccessThresholdOk

`func (o *JobAdvancedSettings) GetReadinessProbeSuccessThresholdOk() (*int32, bool)`

GetReadinessProbeSuccessThresholdOk returns a tuple with the ReadinessProbeSuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeSuccessThreshold

`func (o *JobAdvancedSettings) SetReadinessProbeSuccessThreshold(v int32)`

SetReadinessProbeSuccessThreshold sets ReadinessProbeSuccessThreshold field to given value.

### HasReadinessProbeSuccessThreshold

`func (o *JobAdvancedSettings) HasReadinessProbeSuccessThreshold() bool`

HasReadinessProbeSuccessThreshold returns a boolean if a field has been set.

### GetReadinessProbeFailureThreshold

`func (o *JobAdvancedSettings) GetReadinessProbeFailureThreshold() int32`

GetReadinessProbeFailureThreshold returns the ReadinessProbeFailureThreshold field if non-nil, zero value otherwise.

### GetReadinessProbeFailureThresholdOk

`func (o *JobAdvancedSettings) GetReadinessProbeFailureThresholdOk() (*int32, bool)`

GetReadinessProbeFailureThresholdOk returns a tuple with the ReadinessProbeFailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeFailureThreshold

`func (o *JobAdvancedSettings) SetReadinessProbeFailureThreshold(v int32)`

SetReadinessProbeFailureThreshold sets ReadinessProbeFailureThreshold field to given value.

### HasReadinessProbeFailureThreshold

`func (o *JobAdvancedSettings) HasReadinessProbeFailureThreshold() bool`

HasReadinessProbeFailureThreshold returns a boolean if a field has been set.

### GetLivenessProbeType

`func (o *JobAdvancedSettings) GetLivenessProbeType() string`

GetLivenessProbeType returns the LivenessProbeType field if non-nil, zero value otherwise.

### GetLivenessProbeTypeOk

`func (o *JobAdvancedSettings) GetLivenessProbeTypeOk() (*string, bool)`

GetLivenessProbeTypeOk returns a tuple with the LivenessProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeType

`func (o *JobAdvancedSettings) SetLivenessProbeType(v string)`

SetLivenessProbeType sets LivenessProbeType field to given value.

### HasLivenessProbeType

`func (o *JobAdvancedSettings) HasLivenessProbeType() bool`

HasLivenessProbeType returns a boolean if a field has been set.

### GetLivenessProbeHttpGetPath

`func (o *JobAdvancedSettings) GetLivenessProbeHttpGetPath() string`

GetLivenessProbeHttpGetPath returns the LivenessProbeHttpGetPath field if non-nil, zero value otherwise.

### GetLivenessProbeHttpGetPathOk

`func (o *JobAdvancedSettings) GetLivenessProbeHttpGetPathOk() (*string, bool)`

GetLivenessProbeHttpGetPathOk returns a tuple with the LivenessProbeHttpGetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeHttpGetPath

`func (o *JobAdvancedSettings) SetLivenessProbeHttpGetPath(v string)`

SetLivenessProbeHttpGetPath sets LivenessProbeHttpGetPath field to given value.

### HasLivenessProbeHttpGetPath

`func (o *JobAdvancedSettings) HasLivenessProbeHttpGetPath() bool`

HasLivenessProbeHttpGetPath returns a boolean if a field has been set.

### GetLivenessProbeInitialDelaySeconds

`func (o *JobAdvancedSettings) GetLivenessProbeInitialDelaySeconds() int32`

GetLivenessProbeInitialDelaySeconds returns the LivenessProbeInitialDelaySeconds field if non-nil, zero value otherwise.

### GetLivenessProbeInitialDelaySecondsOk

`func (o *JobAdvancedSettings) GetLivenessProbeInitialDelaySecondsOk() (*int32, bool)`

GetLivenessProbeInitialDelaySecondsOk returns a tuple with the LivenessProbeInitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeInitialDelaySeconds

`func (o *JobAdvancedSettings) SetLivenessProbeInitialDelaySeconds(v int32)`

SetLivenessProbeInitialDelaySeconds sets LivenessProbeInitialDelaySeconds field to given value.

### HasLivenessProbeInitialDelaySeconds

`func (o *JobAdvancedSettings) HasLivenessProbeInitialDelaySeconds() bool`

HasLivenessProbeInitialDelaySeconds returns a boolean if a field has been set.

### GetLivenessProbePeriodSeconds

`func (o *JobAdvancedSettings) GetLivenessProbePeriodSeconds() int32`

GetLivenessProbePeriodSeconds returns the LivenessProbePeriodSeconds field if non-nil, zero value otherwise.

### GetLivenessProbePeriodSecondsOk

`func (o *JobAdvancedSettings) GetLivenessProbePeriodSecondsOk() (*int32, bool)`

GetLivenessProbePeriodSecondsOk returns a tuple with the LivenessProbePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbePeriodSeconds

`func (o *JobAdvancedSettings) SetLivenessProbePeriodSeconds(v int32)`

SetLivenessProbePeriodSeconds sets LivenessProbePeriodSeconds field to given value.

### HasLivenessProbePeriodSeconds

`func (o *JobAdvancedSettings) HasLivenessProbePeriodSeconds() bool`

HasLivenessProbePeriodSeconds returns a boolean if a field has been set.

### GetLivenessProbeTimeoutSeconds

`func (o *JobAdvancedSettings) GetLivenessProbeTimeoutSeconds() int32`

GetLivenessProbeTimeoutSeconds returns the LivenessProbeTimeoutSeconds field if non-nil, zero value otherwise.

### GetLivenessProbeTimeoutSecondsOk

`func (o *JobAdvancedSettings) GetLivenessProbeTimeoutSecondsOk() (*int32, bool)`

GetLivenessProbeTimeoutSecondsOk returns a tuple with the LivenessProbeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeTimeoutSeconds

`func (o *JobAdvancedSettings) SetLivenessProbeTimeoutSeconds(v int32)`

SetLivenessProbeTimeoutSeconds sets LivenessProbeTimeoutSeconds field to given value.

### HasLivenessProbeTimeoutSeconds

`func (o *JobAdvancedSettings) HasLivenessProbeTimeoutSeconds() bool`

HasLivenessProbeTimeoutSeconds returns a boolean if a field has been set.

### GetLivenessProbeSuccessThreshold

`func (o *JobAdvancedSettings) GetLivenessProbeSuccessThreshold() int32`

GetLivenessProbeSuccessThreshold returns the LivenessProbeSuccessThreshold field if non-nil, zero value otherwise.

### GetLivenessProbeSuccessThresholdOk

`func (o *JobAdvancedSettings) GetLivenessProbeSuccessThresholdOk() (*int32, bool)`

GetLivenessProbeSuccessThresholdOk returns a tuple with the LivenessProbeSuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeSuccessThreshold

`func (o *JobAdvancedSettings) SetLivenessProbeSuccessThreshold(v int32)`

SetLivenessProbeSuccessThreshold sets LivenessProbeSuccessThreshold field to given value.

### HasLivenessProbeSuccessThreshold

`func (o *JobAdvancedSettings) HasLivenessProbeSuccessThreshold() bool`

HasLivenessProbeSuccessThreshold returns a boolean if a field has been set.

### GetLivenessProbeFailureThreshold

`func (o *JobAdvancedSettings) GetLivenessProbeFailureThreshold() int32`

GetLivenessProbeFailureThreshold returns the LivenessProbeFailureThreshold field if non-nil, zero value otherwise.

### GetLivenessProbeFailureThresholdOk

`func (o *JobAdvancedSettings) GetLivenessProbeFailureThresholdOk() (*int32, bool)`

GetLivenessProbeFailureThresholdOk returns a tuple with the LivenessProbeFailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeFailureThreshold

`func (o *JobAdvancedSettings) SetLivenessProbeFailureThreshold(v int32)`

SetLivenessProbeFailureThreshold sets LivenessProbeFailureThreshold field to given value.

### HasLivenessProbeFailureThreshold

`func (o *JobAdvancedSettings) HasLivenessProbeFailureThreshold() bool`

HasLivenessProbeFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


