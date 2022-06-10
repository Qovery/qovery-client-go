# ContainerAdvancedSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentDelayStartTimeSec** | Pointer to **int32** | please use &#x60;readiness_probe.initial_delay_seconds&#x60; and &#x60;liveness_probe.initial_delay_seconds&#x60; instead | [optional] [default to 30]
**DeploymentCustomDomainCheckEnabled** | Pointer to **bool** | disable custom domain check when deploying an application | [optional] [default to true]
**BuildTimeoutMaxSec** | Pointer to **int32** |  | [optional] [default to 1800]
**NetworkIngressProxyBodySizeMb** | Pointer to **int32** |  | [optional] [default to 100]
**NetworkIngressEnableCors** | Pointer to **bool** |  | [optional] [default to false]
**NetworkIngressCorsAllowOrigin** | Pointer to **string** |  | [optional] [default to "*"]
**NetworkIngressCorsAllowMethods** | Pointer to **string** |  | [optional] [default to "GET, PUT, POST, DELETE, PATCH, OPTIONS"]
**NetworkIngressCorsAllowHeaders** | Pointer to **string** |  | [optional] [default to "DNT,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization"]
**ReadinessProbeType** | Pointer to **string** | &#x60;NONE&#x60; disable readiness probe &#x60;TCP&#x60; enable TCP readiness probe &#x60;HTTP&#x60; enable HTTP readiness probe  | [optional] [default to "TCP"]
**ReadinessProbeHttpGetPath** | Pointer to **string** | HTTP GET path to check status (must returns 2xx E.g \&quot;/healtz\&quot;) - only usable with TYPE &#x3D; HTTP | [optional] [default to "/"]
**ReadinessProbeInitialDelaySeconds** | Pointer to **int32** | Delay before liveness probe is initiated | [optional] [default to 30]
**ReadinessProbePeriodSeconds** | Pointer to **int32** | How often to perform the probe | [optional] [default to 10]
**ReadinessProbeTimeoutSeconds** | Pointer to **int32** | When the probe times out | [optional] [default to 1]
**ReadinessProbeSuccessThreshold** | Pointer to **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. | [optional] [default to 1]
**ReadinessProbeFailureThreshold** | Pointer to **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. | [optional] [default to 3]
**LivenessProbeType** | Pointer to **string** | &#x60;NONE&#x60; disable liveness probe &#x60;TCP&#x60; enable TCP liveness probe &#x60;HTTP&#x60; enable HTTP liveness probe  | [optional] [default to "TCP"]
**LivenessProbeHttpGetPath** | Pointer to **string** | HTTP GET path to check status (must returns 2xx E.g \&quot;/healtz\&quot;) - only usable with TYPE &#x3D; HTTP | [optional] [default to "/"]
**LivenessProbeInitialDelaySeconds** | Pointer to **int32** | Delay before liveness probe is initiated | [optional] [default to 30]
**LivenessProbePeriodSeconds** | Pointer to **int32** | How often to perform the probe | [optional] [default to 10]
**LivenessProbeTimeoutSeconds** | Pointer to **int32** | When the probe times out | [optional] [default to 5]
**LivenessProbeSuccessThreshold** | Pointer to **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. | [optional] [default to 1]
**LivenessProbeFailureThreshold** | Pointer to **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. | [optional] [default to 3]

## Methods

### NewContainerAdvancedSettingsResponse

`func NewContainerAdvancedSettingsResponse() *ContainerAdvancedSettingsResponse`

NewContainerAdvancedSettingsResponse instantiates a new ContainerAdvancedSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerAdvancedSettingsResponseWithDefaults

`func NewContainerAdvancedSettingsResponseWithDefaults() *ContainerAdvancedSettingsResponse`

NewContainerAdvancedSettingsResponseWithDefaults instantiates a new ContainerAdvancedSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentDelayStartTimeSec

`func (o *ContainerAdvancedSettingsResponse) GetDeploymentDelayStartTimeSec() int32`

GetDeploymentDelayStartTimeSec returns the DeploymentDelayStartTimeSec field if non-nil, zero value otherwise.

### GetDeploymentDelayStartTimeSecOk

`func (o *ContainerAdvancedSettingsResponse) GetDeploymentDelayStartTimeSecOk() (*int32, bool)`

GetDeploymentDelayStartTimeSecOk returns a tuple with the DeploymentDelayStartTimeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentDelayStartTimeSec

`func (o *ContainerAdvancedSettingsResponse) SetDeploymentDelayStartTimeSec(v int32)`

SetDeploymentDelayStartTimeSec sets DeploymentDelayStartTimeSec field to given value.

### HasDeploymentDelayStartTimeSec

`func (o *ContainerAdvancedSettingsResponse) HasDeploymentDelayStartTimeSec() bool`

HasDeploymentDelayStartTimeSec returns a boolean if a field has been set.

### GetDeploymentCustomDomainCheckEnabled

`func (o *ContainerAdvancedSettingsResponse) GetDeploymentCustomDomainCheckEnabled() bool`

GetDeploymentCustomDomainCheckEnabled returns the DeploymentCustomDomainCheckEnabled field if non-nil, zero value otherwise.

### GetDeploymentCustomDomainCheckEnabledOk

`func (o *ContainerAdvancedSettingsResponse) GetDeploymentCustomDomainCheckEnabledOk() (*bool, bool)`

GetDeploymentCustomDomainCheckEnabledOk returns a tuple with the DeploymentCustomDomainCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCustomDomainCheckEnabled

`func (o *ContainerAdvancedSettingsResponse) SetDeploymentCustomDomainCheckEnabled(v bool)`

SetDeploymentCustomDomainCheckEnabled sets DeploymentCustomDomainCheckEnabled field to given value.

### HasDeploymentCustomDomainCheckEnabled

`func (o *ContainerAdvancedSettingsResponse) HasDeploymentCustomDomainCheckEnabled() bool`

HasDeploymentCustomDomainCheckEnabled returns a boolean if a field has been set.

### GetBuildTimeoutMaxSec

`func (o *ContainerAdvancedSettingsResponse) GetBuildTimeoutMaxSec() int32`

GetBuildTimeoutMaxSec returns the BuildTimeoutMaxSec field if non-nil, zero value otherwise.

### GetBuildTimeoutMaxSecOk

`func (o *ContainerAdvancedSettingsResponse) GetBuildTimeoutMaxSecOk() (*int32, bool)`

GetBuildTimeoutMaxSecOk returns a tuple with the BuildTimeoutMaxSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimeoutMaxSec

`func (o *ContainerAdvancedSettingsResponse) SetBuildTimeoutMaxSec(v int32)`

SetBuildTimeoutMaxSec sets BuildTimeoutMaxSec field to given value.

### HasBuildTimeoutMaxSec

`func (o *ContainerAdvancedSettingsResponse) HasBuildTimeoutMaxSec() bool`

HasBuildTimeoutMaxSec returns a boolean if a field has been set.

### GetNetworkIngressProxyBodySizeMb

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressProxyBodySizeMb() int32`

GetNetworkIngressProxyBodySizeMb returns the NetworkIngressProxyBodySizeMb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBodySizeMbOk

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressProxyBodySizeMbOk() (*int32, bool)`

GetNetworkIngressProxyBodySizeMbOk returns a tuple with the NetworkIngressProxyBodySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBodySizeMb

`func (o *ContainerAdvancedSettingsResponse) SetNetworkIngressProxyBodySizeMb(v int32)`

SetNetworkIngressProxyBodySizeMb sets NetworkIngressProxyBodySizeMb field to given value.

### HasNetworkIngressProxyBodySizeMb

`func (o *ContainerAdvancedSettingsResponse) HasNetworkIngressProxyBodySizeMb() bool`

HasNetworkIngressProxyBodySizeMb returns a boolean if a field has been set.

### GetNetworkIngressEnableCors

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressEnableCors() bool`

GetNetworkIngressEnableCors returns the NetworkIngressEnableCors field if non-nil, zero value otherwise.

### GetNetworkIngressEnableCorsOk

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressEnableCorsOk() (*bool, bool)`

GetNetworkIngressEnableCorsOk returns a tuple with the NetworkIngressEnableCors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableCors

`func (o *ContainerAdvancedSettingsResponse) SetNetworkIngressEnableCors(v bool)`

SetNetworkIngressEnableCors sets NetworkIngressEnableCors field to given value.

### HasNetworkIngressEnableCors

`func (o *ContainerAdvancedSettingsResponse) HasNetworkIngressEnableCors() bool`

HasNetworkIngressEnableCors returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowOrigin

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressCorsAllowOrigin() string`

GetNetworkIngressCorsAllowOrigin returns the NetworkIngressCorsAllowOrigin field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowOriginOk

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressCorsAllowOriginOk() (*string, bool)`

GetNetworkIngressCorsAllowOriginOk returns a tuple with the NetworkIngressCorsAllowOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowOrigin

`func (o *ContainerAdvancedSettingsResponse) SetNetworkIngressCorsAllowOrigin(v string)`

SetNetworkIngressCorsAllowOrigin sets NetworkIngressCorsAllowOrigin field to given value.

### HasNetworkIngressCorsAllowOrigin

`func (o *ContainerAdvancedSettingsResponse) HasNetworkIngressCorsAllowOrigin() bool`

HasNetworkIngressCorsAllowOrigin returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowMethods

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressCorsAllowMethods() string`

GetNetworkIngressCorsAllowMethods returns the NetworkIngressCorsAllowMethods field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowMethodsOk

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressCorsAllowMethodsOk() (*string, bool)`

GetNetworkIngressCorsAllowMethodsOk returns a tuple with the NetworkIngressCorsAllowMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowMethods

`func (o *ContainerAdvancedSettingsResponse) SetNetworkIngressCorsAllowMethods(v string)`

SetNetworkIngressCorsAllowMethods sets NetworkIngressCorsAllowMethods field to given value.

### HasNetworkIngressCorsAllowMethods

`func (o *ContainerAdvancedSettingsResponse) HasNetworkIngressCorsAllowMethods() bool`

HasNetworkIngressCorsAllowMethods returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowHeaders

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressCorsAllowHeaders() string`

GetNetworkIngressCorsAllowHeaders returns the NetworkIngressCorsAllowHeaders field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowHeadersOk

`func (o *ContainerAdvancedSettingsResponse) GetNetworkIngressCorsAllowHeadersOk() (*string, bool)`

GetNetworkIngressCorsAllowHeadersOk returns a tuple with the NetworkIngressCorsAllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowHeaders

`func (o *ContainerAdvancedSettingsResponse) SetNetworkIngressCorsAllowHeaders(v string)`

SetNetworkIngressCorsAllowHeaders sets NetworkIngressCorsAllowHeaders field to given value.

### HasNetworkIngressCorsAllowHeaders

`func (o *ContainerAdvancedSettingsResponse) HasNetworkIngressCorsAllowHeaders() bool`

HasNetworkIngressCorsAllowHeaders returns a boolean if a field has been set.

### GetReadinessProbeType

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeType() string`

GetReadinessProbeType returns the ReadinessProbeType field if non-nil, zero value otherwise.

### GetReadinessProbeTypeOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeTypeOk() (*string, bool)`

GetReadinessProbeTypeOk returns a tuple with the ReadinessProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeType

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbeType(v string)`

SetReadinessProbeType sets ReadinessProbeType field to given value.

### HasReadinessProbeType

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbeType() bool`

HasReadinessProbeType returns a boolean if a field has been set.

### GetReadinessProbeHttpGetPath

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeHttpGetPath() string`

GetReadinessProbeHttpGetPath returns the ReadinessProbeHttpGetPath field if non-nil, zero value otherwise.

### GetReadinessProbeHttpGetPathOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeHttpGetPathOk() (*string, bool)`

GetReadinessProbeHttpGetPathOk returns a tuple with the ReadinessProbeHttpGetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeHttpGetPath

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbeHttpGetPath(v string)`

SetReadinessProbeHttpGetPath sets ReadinessProbeHttpGetPath field to given value.

### HasReadinessProbeHttpGetPath

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbeHttpGetPath() bool`

HasReadinessProbeHttpGetPath returns a boolean if a field has been set.

### GetReadinessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeInitialDelaySeconds() int32`

GetReadinessProbeInitialDelaySeconds returns the ReadinessProbeInitialDelaySeconds field if non-nil, zero value otherwise.

### GetReadinessProbeInitialDelaySecondsOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeInitialDelaySecondsOk() (*int32, bool)`

GetReadinessProbeInitialDelaySecondsOk returns a tuple with the ReadinessProbeInitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbeInitialDelaySeconds(v int32)`

SetReadinessProbeInitialDelaySeconds sets ReadinessProbeInitialDelaySeconds field to given value.

### HasReadinessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbeInitialDelaySeconds() bool`

HasReadinessProbeInitialDelaySeconds returns a boolean if a field has been set.

### GetReadinessProbePeriodSeconds

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbePeriodSeconds() int32`

GetReadinessProbePeriodSeconds returns the ReadinessProbePeriodSeconds field if non-nil, zero value otherwise.

### GetReadinessProbePeriodSecondsOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbePeriodSecondsOk() (*int32, bool)`

GetReadinessProbePeriodSecondsOk returns a tuple with the ReadinessProbePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbePeriodSeconds

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbePeriodSeconds(v int32)`

SetReadinessProbePeriodSeconds sets ReadinessProbePeriodSeconds field to given value.

### HasReadinessProbePeriodSeconds

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbePeriodSeconds() bool`

HasReadinessProbePeriodSeconds returns a boolean if a field has been set.

### GetReadinessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeTimeoutSeconds() int32`

GetReadinessProbeTimeoutSeconds returns the ReadinessProbeTimeoutSeconds field if non-nil, zero value otherwise.

### GetReadinessProbeTimeoutSecondsOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeTimeoutSecondsOk() (*int32, bool)`

GetReadinessProbeTimeoutSecondsOk returns a tuple with the ReadinessProbeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbeTimeoutSeconds(v int32)`

SetReadinessProbeTimeoutSeconds sets ReadinessProbeTimeoutSeconds field to given value.

### HasReadinessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbeTimeoutSeconds() bool`

HasReadinessProbeTimeoutSeconds returns a boolean if a field has been set.

### GetReadinessProbeSuccessThreshold

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeSuccessThreshold() int32`

GetReadinessProbeSuccessThreshold returns the ReadinessProbeSuccessThreshold field if non-nil, zero value otherwise.

### GetReadinessProbeSuccessThresholdOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeSuccessThresholdOk() (*int32, bool)`

GetReadinessProbeSuccessThresholdOk returns a tuple with the ReadinessProbeSuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeSuccessThreshold

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbeSuccessThreshold(v int32)`

SetReadinessProbeSuccessThreshold sets ReadinessProbeSuccessThreshold field to given value.

### HasReadinessProbeSuccessThreshold

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbeSuccessThreshold() bool`

HasReadinessProbeSuccessThreshold returns a boolean if a field has been set.

### GetReadinessProbeFailureThreshold

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeFailureThreshold() int32`

GetReadinessProbeFailureThreshold returns the ReadinessProbeFailureThreshold field if non-nil, zero value otherwise.

### GetReadinessProbeFailureThresholdOk

`func (o *ContainerAdvancedSettingsResponse) GetReadinessProbeFailureThresholdOk() (*int32, bool)`

GetReadinessProbeFailureThresholdOk returns a tuple with the ReadinessProbeFailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeFailureThreshold

`func (o *ContainerAdvancedSettingsResponse) SetReadinessProbeFailureThreshold(v int32)`

SetReadinessProbeFailureThreshold sets ReadinessProbeFailureThreshold field to given value.

### HasReadinessProbeFailureThreshold

`func (o *ContainerAdvancedSettingsResponse) HasReadinessProbeFailureThreshold() bool`

HasReadinessProbeFailureThreshold returns a boolean if a field has been set.

### GetLivenessProbeType

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeType() string`

GetLivenessProbeType returns the LivenessProbeType field if non-nil, zero value otherwise.

### GetLivenessProbeTypeOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeTypeOk() (*string, bool)`

GetLivenessProbeTypeOk returns a tuple with the LivenessProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeType

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbeType(v string)`

SetLivenessProbeType sets LivenessProbeType field to given value.

### HasLivenessProbeType

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbeType() bool`

HasLivenessProbeType returns a boolean if a field has been set.

### GetLivenessProbeHttpGetPath

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeHttpGetPath() string`

GetLivenessProbeHttpGetPath returns the LivenessProbeHttpGetPath field if non-nil, zero value otherwise.

### GetLivenessProbeHttpGetPathOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeHttpGetPathOk() (*string, bool)`

GetLivenessProbeHttpGetPathOk returns a tuple with the LivenessProbeHttpGetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeHttpGetPath

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbeHttpGetPath(v string)`

SetLivenessProbeHttpGetPath sets LivenessProbeHttpGetPath field to given value.

### HasLivenessProbeHttpGetPath

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbeHttpGetPath() bool`

HasLivenessProbeHttpGetPath returns a boolean if a field has been set.

### GetLivenessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeInitialDelaySeconds() int32`

GetLivenessProbeInitialDelaySeconds returns the LivenessProbeInitialDelaySeconds field if non-nil, zero value otherwise.

### GetLivenessProbeInitialDelaySecondsOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeInitialDelaySecondsOk() (*int32, bool)`

GetLivenessProbeInitialDelaySecondsOk returns a tuple with the LivenessProbeInitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbeInitialDelaySeconds(v int32)`

SetLivenessProbeInitialDelaySeconds sets LivenessProbeInitialDelaySeconds field to given value.

### HasLivenessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbeInitialDelaySeconds() bool`

HasLivenessProbeInitialDelaySeconds returns a boolean if a field has been set.

### GetLivenessProbePeriodSeconds

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbePeriodSeconds() int32`

GetLivenessProbePeriodSeconds returns the LivenessProbePeriodSeconds field if non-nil, zero value otherwise.

### GetLivenessProbePeriodSecondsOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbePeriodSecondsOk() (*int32, bool)`

GetLivenessProbePeriodSecondsOk returns a tuple with the LivenessProbePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbePeriodSeconds

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbePeriodSeconds(v int32)`

SetLivenessProbePeriodSeconds sets LivenessProbePeriodSeconds field to given value.

### HasLivenessProbePeriodSeconds

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbePeriodSeconds() bool`

HasLivenessProbePeriodSeconds returns a boolean if a field has been set.

### GetLivenessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeTimeoutSeconds() int32`

GetLivenessProbeTimeoutSeconds returns the LivenessProbeTimeoutSeconds field if non-nil, zero value otherwise.

### GetLivenessProbeTimeoutSecondsOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeTimeoutSecondsOk() (*int32, bool)`

GetLivenessProbeTimeoutSecondsOk returns a tuple with the LivenessProbeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbeTimeoutSeconds(v int32)`

SetLivenessProbeTimeoutSeconds sets LivenessProbeTimeoutSeconds field to given value.

### HasLivenessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbeTimeoutSeconds() bool`

HasLivenessProbeTimeoutSeconds returns a boolean if a field has been set.

### GetLivenessProbeSuccessThreshold

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeSuccessThreshold() int32`

GetLivenessProbeSuccessThreshold returns the LivenessProbeSuccessThreshold field if non-nil, zero value otherwise.

### GetLivenessProbeSuccessThresholdOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeSuccessThresholdOk() (*int32, bool)`

GetLivenessProbeSuccessThresholdOk returns a tuple with the LivenessProbeSuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeSuccessThreshold

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbeSuccessThreshold(v int32)`

SetLivenessProbeSuccessThreshold sets LivenessProbeSuccessThreshold field to given value.

### HasLivenessProbeSuccessThreshold

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbeSuccessThreshold() bool`

HasLivenessProbeSuccessThreshold returns a boolean if a field has been set.

### GetLivenessProbeFailureThreshold

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeFailureThreshold() int32`

GetLivenessProbeFailureThreshold returns the LivenessProbeFailureThreshold field if non-nil, zero value otherwise.

### GetLivenessProbeFailureThresholdOk

`func (o *ContainerAdvancedSettingsResponse) GetLivenessProbeFailureThresholdOk() (*int32, bool)`

GetLivenessProbeFailureThresholdOk returns a tuple with the LivenessProbeFailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeFailureThreshold

`func (o *ContainerAdvancedSettingsResponse) SetLivenessProbeFailureThreshold(v int32)`

SetLivenessProbeFailureThreshold sets LivenessProbeFailureThreshold field to given value.

### HasLivenessProbeFailureThreshold

`func (o *ContainerAdvancedSettingsResponse) HasLivenessProbeFailureThreshold() bool`

HasLivenessProbeFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


