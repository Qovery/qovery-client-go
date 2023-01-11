# ContainerAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCustomDomainCheckEnabled** | Pointer to **bool** | disable custom domain check when deploying an application | [optional] [default to true]
**DeploymentTerminationGracePeriodSeconds** | Pointer to **int32** | define how long in seconds an application is supposed to be stopped gracefully | [optional] [default to 60]
**NetworkIngressProxyBodySizeMb** | Pointer to **int32** |  | [optional] [default to 100]
**NetworkIngressEnableCors** | Pointer to **bool** |  | [optional] [default to false]
**NetworkIngressCorsAllowOrigin** | Pointer to **string** |  | [optional] [default to "*"]
**NetworkIngressCorsAllowMethods** | Pointer to **string** |  | [optional] [default to "GET, PUT, POST, DELETE, PATCH, OPTIONS"]
**NetworkIngressCorsAllowHeaders** | Pointer to **string** |  | [optional] [default to "DNT,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization"]
**NetworkIngressProxyBufferSizeKb** | Pointer to **int32** | header buffer size used while reading response header from upstream | [optional] [default to 4]
**NetworkIngressKeepaliveTimeSeconds** | Pointer to **int32** | Limits the maximum time (in seconds) during which requests can be processed through one keepalive connection | [optional] [default to 3600]
**NetworkIngressKeepaliveTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) during which an idle keepalive connection to an upstream server will stay open. | [optional] [default to 60]
**NetworkIngressSendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a response to the client | [optional] [default to 60]
**NetworkIngressProxyConnectTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for establishing a connection to a proxied server | [optional] [default to 60]
**NetworkIngressProxySendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the proxied server | [optional] [default to 60]
**NetworkIngressProxyReadTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for reading a response from the proxied server | [optional] [default to 60]
**NetworkIngressWhitelistSourceRange** | Pointer to **string** | list of source ranges to allow access to ingress proxy.  This property can be used to whitelist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1 To allow all source ranges, set 0.0.0.0/0.  | [optional] [default to "0.0.0.0/0"]
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

### NewContainerAdvancedSettings

`func NewContainerAdvancedSettings() *ContainerAdvancedSettings`

NewContainerAdvancedSettings instantiates a new ContainerAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerAdvancedSettingsWithDefaults

`func NewContainerAdvancedSettingsWithDefaults() *ContainerAdvancedSettings`

NewContainerAdvancedSettingsWithDefaults instantiates a new ContainerAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCustomDomainCheckEnabled

`func (o *ContainerAdvancedSettings) GetDeploymentCustomDomainCheckEnabled() bool`

GetDeploymentCustomDomainCheckEnabled returns the DeploymentCustomDomainCheckEnabled field if non-nil, zero value otherwise.

### GetDeploymentCustomDomainCheckEnabledOk

`func (o *ContainerAdvancedSettings) GetDeploymentCustomDomainCheckEnabledOk() (*bool, bool)`

GetDeploymentCustomDomainCheckEnabledOk returns a tuple with the DeploymentCustomDomainCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCustomDomainCheckEnabled

`func (o *ContainerAdvancedSettings) SetDeploymentCustomDomainCheckEnabled(v bool)`

SetDeploymentCustomDomainCheckEnabled sets DeploymentCustomDomainCheckEnabled field to given value.

### HasDeploymentCustomDomainCheckEnabled

`func (o *ContainerAdvancedSettings) HasDeploymentCustomDomainCheckEnabled() bool`

HasDeploymentCustomDomainCheckEnabled returns a boolean if a field has been set.

### GetDeploymentTerminationGracePeriodSeconds

`func (o *ContainerAdvancedSettings) GetDeploymentTerminationGracePeriodSeconds() int32`

GetDeploymentTerminationGracePeriodSeconds returns the DeploymentTerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeploymentTerminationGracePeriodSecondsOk

`func (o *ContainerAdvancedSettings) GetDeploymentTerminationGracePeriodSecondsOk() (*int32, bool)`

GetDeploymentTerminationGracePeriodSecondsOk returns a tuple with the DeploymentTerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTerminationGracePeriodSeconds

`func (o *ContainerAdvancedSettings) SetDeploymentTerminationGracePeriodSeconds(v int32)`

SetDeploymentTerminationGracePeriodSeconds sets DeploymentTerminationGracePeriodSeconds field to given value.

### HasDeploymentTerminationGracePeriodSeconds

`func (o *ContainerAdvancedSettings) HasDeploymentTerminationGracePeriodSeconds() bool`

HasDeploymentTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyBodySizeMb

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyBodySizeMb() int32`

GetNetworkIngressProxyBodySizeMb returns the NetworkIngressProxyBodySizeMb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBodySizeMbOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyBodySizeMbOk() (*int32, bool)`

GetNetworkIngressProxyBodySizeMbOk returns a tuple with the NetworkIngressProxyBodySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBodySizeMb

`func (o *ContainerAdvancedSettings) SetNetworkIngressProxyBodySizeMb(v int32)`

SetNetworkIngressProxyBodySizeMb sets NetworkIngressProxyBodySizeMb field to given value.

### HasNetworkIngressProxyBodySizeMb

`func (o *ContainerAdvancedSettings) HasNetworkIngressProxyBodySizeMb() bool`

HasNetworkIngressProxyBodySizeMb returns a boolean if a field has been set.

### GetNetworkIngressEnableCors

`func (o *ContainerAdvancedSettings) GetNetworkIngressEnableCors() bool`

GetNetworkIngressEnableCors returns the NetworkIngressEnableCors field if non-nil, zero value otherwise.

### GetNetworkIngressEnableCorsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressEnableCorsOk() (*bool, bool)`

GetNetworkIngressEnableCorsOk returns a tuple with the NetworkIngressEnableCors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableCors

`func (o *ContainerAdvancedSettings) SetNetworkIngressEnableCors(v bool)`

SetNetworkIngressEnableCors sets NetworkIngressEnableCors field to given value.

### HasNetworkIngressEnableCors

`func (o *ContainerAdvancedSettings) HasNetworkIngressEnableCors() bool`

HasNetworkIngressEnableCors returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowOrigin

`func (o *ContainerAdvancedSettings) GetNetworkIngressCorsAllowOrigin() string`

GetNetworkIngressCorsAllowOrigin returns the NetworkIngressCorsAllowOrigin field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowOriginOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressCorsAllowOriginOk() (*string, bool)`

GetNetworkIngressCorsAllowOriginOk returns a tuple with the NetworkIngressCorsAllowOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowOrigin

`func (o *ContainerAdvancedSettings) SetNetworkIngressCorsAllowOrigin(v string)`

SetNetworkIngressCorsAllowOrigin sets NetworkIngressCorsAllowOrigin field to given value.

### HasNetworkIngressCorsAllowOrigin

`func (o *ContainerAdvancedSettings) HasNetworkIngressCorsAllowOrigin() bool`

HasNetworkIngressCorsAllowOrigin returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowMethods

`func (o *ContainerAdvancedSettings) GetNetworkIngressCorsAllowMethods() string`

GetNetworkIngressCorsAllowMethods returns the NetworkIngressCorsAllowMethods field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowMethodsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressCorsAllowMethodsOk() (*string, bool)`

GetNetworkIngressCorsAllowMethodsOk returns a tuple with the NetworkIngressCorsAllowMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowMethods

`func (o *ContainerAdvancedSettings) SetNetworkIngressCorsAllowMethods(v string)`

SetNetworkIngressCorsAllowMethods sets NetworkIngressCorsAllowMethods field to given value.

### HasNetworkIngressCorsAllowMethods

`func (o *ContainerAdvancedSettings) HasNetworkIngressCorsAllowMethods() bool`

HasNetworkIngressCorsAllowMethods returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowHeaders

`func (o *ContainerAdvancedSettings) GetNetworkIngressCorsAllowHeaders() string`

GetNetworkIngressCorsAllowHeaders returns the NetworkIngressCorsAllowHeaders field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowHeadersOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressCorsAllowHeadersOk() (*string, bool)`

GetNetworkIngressCorsAllowHeadersOk returns a tuple with the NetworkIngressCorsAllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowHeaders

`func (o *ContainerAdvancedSettings) SetNetworkIngressCorsAllowHeaders(v string)`

SetNetworkIngressCorsAllowHeaders sets NetworkIngressCorsAllowHeaders field to given value.

### HasNetworkIngressCorsAllowHeaders

`func (o *ContainerAdvancedSettings) HasNetworkIngressCorsAllowHeaders() bool`

HasNetworkIngressCorsAllowHeaders returns a boolean if a field has been set.

### GetNetworkIngressProxyBufferSizeKb

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyBufferSizeKb() int32`

GetNetworkIngressProxyBufferSizeKb returns the NetworkIngressProxyBufferSizeKb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBufferSizeKbOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyBufferSizeKbOk() (*int32, bool)`

GetNetworkIngressProxyBufferSizeKbOk returns a tuple with the NetworkIngressProxyBufferSizeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBufferSizeKb

`func (o *ContainerAdvancedSettings) SetNetworkIngressProxyBufferSizeKb(v int32)`

SetNetworkIngressProxyBufferSizeKb sets NetworkIngressProxyBufferSizeKb field to given value.

### HasNetworkIngressProxyBufferSizeKb

`func (o *ContainerAdvancedSettings) HasNetworkIngressProxyBufferSizeKb() bool`

HasNetworkIngressProxyBufferSizeKb returns a boolean if a field has been set.

### GetNetworkIngressKeepaliveTimeSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressKeepaliveTimeSeconds() int32`

GetNetworkIngressKeepaliveTimeSeconds returns the NetworkIngressKeepaliveTimeSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressKeepaliveTimeSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressKeepaliveTimeSecondsOk() (*int32, bool)`

GetNetworkIngressKeepaliveTimeSecondsOk returns a tuple with the NetworkIngressKeepaliveTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressKeepaliveTimeSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressKeepaliveTimeSeconds(v int32)`

SetNetworkIngressKeepaliveTimeSeconds sets NetworkIngressKeepaliveTimeSeconds field to given value.

### HasNetworkIngressKeepaliveTimeSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressKeepaliveTimeSeconds() bool`

HasNetworkIngressKeepaliveTimeSeconds returns a boolean if a field has been set.

### GetNetworkIngressKeepaliveTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressKeepaliveTimeoutSeconds() int32`

GetNetworkIngressKeepaliveTimeoutSeconds returns the NetworkIngressKeepaliveTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressKeepaliveTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressKeepaliveTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressKeepaliveTimeoutSecondsOk returns a tuple with the NetworkIngressKeepaliveTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressKeepaliveTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressKeepaliveTimeoutSeconds(v int32)`

SetNetworkIngressKeepaliveTimeoutSeconds sets NetworkIngressKeepaliveTimeoutSeconds field to given value.

### HasNetworkIngressKeepaliveTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressKeepaliveTimeoutSeconds() bool`

HasNetworkIngressKeepaliveTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressSendTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressSendTimeoutSeconds() int32`

GetNetworkIngressSendTimeoutSeconds returns the NetworkIngressSendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressSendTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressSendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressSendTimeoutSecondsOk returns a tuple with the NetworkIngressSendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressSendTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressSendTimeoutSeconds(v int32)`

SetNetworkIngressSendTimeoutSeconds sets NetworkIngressSendTimeoutSeconds field to given value.

### HasNetworkIngressSendTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressSendTimeoutSeconds() bool`

HasNetworkIngressSendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyConnectTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyConnectTimeoutSeconds() int32`

GetNetworkIngressProxyConnectTimeoutSeconds returns the NetworkIngressProxyConnectTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxyConnectTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyConnectTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxyConnectTimeoutSecondsOk returns a tuple with the NetworkIngressProxyConnectTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyConnectTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressProxyConnectTimeoutSeconds(v int32)`

SetNetworkIngressProxyConnectTimeoutSeconds sets NetworkIngressProxyConnectTimeoutSeconds field to given value.

### HasNetworkIngressProxyConnectTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressProxyConnectTimeoutSeconds() bool`

HasNetworkIngressProxyConnectTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxySendTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxySendTimeoutSeconds() int32`

GetNetworkIngressProxySendTimeoutSeconds returns the NetworkIngressProxySendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxySendTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxySendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxySendTimeoutSecondsOk returns a tuple with the NetworkIngressProxySendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxySendTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressProxySendTimeoutSeconds(v int32)`

SetNetworkIngressProxySendTimeoutSeconds sets NetworkIngressProxySendTimeoutSeconds field to given value.

### HasNetworkIngressProxySendTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressProxySendTimeoutSeconds() bool`

HasNetworkIngressProxySendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyReadTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyReadTimeoutSeconds() int32`

GetNetworkIngressProxyReadTimeoutSeconds returns the NetworkIngressProxyReadTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxyReadTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressProxyReadTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxyReadTimeoutSecondsOk returns a tuple with the NetworkIngressProxyReadTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyReadTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressProxyReadTimeoutSeconds(v int32)`

SetNetworkIngressProxyReadTimeoutSeconds sets NetworkIngressProxyReadTimeoutSeconds field to given value.

### HasNetworkIngressProxyReadTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressProxyReadTimeoutSeconds() bool`

HasNetworkIngressProxyReadTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressWhitelistSourceRange

`func (o *ContainerAdvancedSettings) GetNetworkIngressWhitelistSourceRange() string`

GetNetworkIngressWhitelistSourceRange returns the NetworkIngressWhitelistSourceRange field if non-nil, zero value otherwise.

### GetNetworkIngressWhitelistSourceRangeOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressWhitelistSourceRangeOk() (*string, bool)`

GetNetworkIngressWhitelistSourceRangeOk returns a tuple with the NetworkIngressWhitelistSourceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressWhitelistSourceRange

`func (o *ContainerAdvancedSettings) SetNetworkIngressWhitelistSourceRange(v string)`

SetNetworkIngressWhitelistSourceRange sets NetworkIngressWhitelistSourceRange field to given value.

### HasNetworkIngressWhitelistSourceRange

`func (o *ContainerAdvancedSettings) HasNetworkIngressWhitelistSourceRange() bool`

HasNetworkIngressWhitelistSourceRange returns a boolean if a field has been set.

### GetReadinessProbeType

`func (o *ContainerAdvancedSettings) GetReadinessProbeType() string`

GetReadinessProbeType returns the ReadinessProbeType field if non-nil, zero value otherwise.

### GetReadinessProbeTypeOk

`func (o *ContainerAdvancedSettings) GetReadinessProbeTypeOk() (*string, bool)`

GetReadinessProbeTypeOk returns a tuple with the ReadinessProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeType

`func (o *ContainerAdvancedSettings) SetReadinessProbeType(v string)`

SetReadinessProbeType sets ReadinessProbeType field to given value.

### HasReadinessProbeType

`func (o *ContainerAdvancedSettings) HasReadinessProbeType() bool`

HasReadinessProbeType returns a boolean if a field has been set.

### GetReadinessProbeHttpGetPath

`func (o *ContainerAdvancedSettings) GetReadinessProbeHttpGetPath() string`

GetReadinessProbeHttpGetPath returns the ReadinessProbeHttpGetPath field if non-nil, zero value otherwise.

### GetReadinessProbeHttpGetPathOk

`func (o *ContainerAdvancedSettings) GetReadinessProbeHttpGetPathOk() (*string, bool)`

GetReadinessProbeHttpGetPathOk returns a tuple with the ReadinessProbeHttpGetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeHttpGetPath

`func (o *ContainerAdvancedSettings) SetReadinessProbeHttpGetPath(v string)`

SetReadinessProbeHttpGetPath sets ReadinessProbeHttpGetPath field to given value.

### HasReadinessProbeHttpGetPath

`func (o *ContainerAdvancedSettings) HasReadinessProbeHttpGetPath() bool`

HasReadinessProbeHttpGetPath returns a boolean if a field has been set.

### GetReadinessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettings) GetReadinessProbeInitialDelaySeconds() int32`

GetReadinessProbeInitialDelaySeconds returns the ReadinessProbeInitialDelaySeconds field if non-nil, zero value otherwise.

### GetReadinessProbeInitialDelaySecondsOk

`func (o *ContainerAdvancedSettings) GetReadinessProbeInitialDelaySecondsOk() (*int32, bool)`

GetReadinessProbeInitialDelaySecondsOk returns a tuple with the ReadinessProbeInitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettings) SetReadinessProbeInitialDelaySeconds(v int32)`

SetReadinessProbeInitialDelaySeconds sets ReadinessProbeInitialDelaySeconds field to given value.

### HasReadinessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettings) HasReadinessProbeInitialDelaySeconds() bool`

HasReadinessProbeInitialDelaySeconds returns a boolean if a field has been set.

### GetReadinessProbePeriodSeconds

`func (o *ContainerAdvancedSettings) GetReadinessProbePeriodSeconds() int32`

GetReadinessProbePeriodSeconds returns the ReadinessProbePeriodSeconds field if non-nil, zero value otherwise.

### GetReadinessProbePeriodSecondsOk

`func (o *ContainerAdvancedSettings) GetReadinessProbePeriodSecondsOk() (*int32, bool)`

GetReadinessProbePeriodSecondsOk returns a tuple with the ReadinessProbePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbePeriodSeconds

`func (o *ContainerAdvancedSettings) SetReadinessProbePeriodSeconds(v int32)`

SetReadinessProbePeriodSeconds sets ReadinessProbePeriodSeconds field to given value.

### HasReadinessProbePeriodSeconds

`func (o *ContainerAdvancedSettings) HasReadinessProbePeriodSeconds() bool`

HasReadinessProbePeriodSeconds returns a boolean if a field has been set.

### GetReadinessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetReadinessProbeTimeoutSeconds() int32`

GetReadinessProbeTimeoutSeconds returns the ReadinessProbeTimeoutSeconds field if non-nil, zero value otherwise.

### GetReadinessProbeTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetReadinessProbeTimeoutSecondsOk() (*int32, bool)`

GetReadinessProbeTimeoutSecondsOk returns a tuple with the ReadinessProbeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetReadinessProbeTimeoutSeconds(v int32)`

SetReadinessProbeTimeoutSeconds sets ReadinessProbeTimeoutSeconds field to given value.

### HasReadinessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasReadinessProbeTimeoutSeconds() bool`

HasReadinessProbeTimeoutSeconds returns a boolean if a field has been set.

### GetReadinessProbeSuccessThreshold

`func (o *ContainerAdvancedSettings) GetReadinessProbeSuccessThreshold() int32`

GetReadinessProbeSuccessThreshold returns the ReadinessProbeSuccessThreshold field if non-nil, zero value otherwise.

### GetReadinessProbeSuccessThresholdOk

`func (o *ContainerAdvancedSettings) GetReadinessProbeSuccessThresholdOk() (*int32, bool)`

GetReadinessProbeSuccessThresholdOk returns a tuple with the ReadinessProbeSuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeSuccessThreshold

`func (o *ContainerAdvancedSettings) SetReadinessProbeSuccessThreshold(v int32)`

SetReadinessProbeSuccessThreshold sets ReadinessProbeSuccessThreshold field to given value.

### HasReadinessProbeSuccessThreshold

`func (o *ContainerAdvancedSettings) HasReadinessProbeSuccessThreshold() bool`

HasReadinessProbeSuccessThreshold returns a boolean if a field has been set.

### GetReadinessProbeFailureThreshold

`func (o *ContainerAdvancedSettings) GetReadinessProbeFailureThreshold() int32`

GetReadinessProbeFailureThreshold returns the ReadinessProbeFailureThreshold field if non-nil, zero value otherwise.

### GetReadinessProbeFailureThresholdOk

`func (o *ContainerAdvancedSettings) GetReadinessProbeFailureThresholdOk() (*int32, bool)`

GetReadinessProbeFailureThresholdOk returns a tuple with the ReadinessProbeFailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbeFailureThreshold

`func (o *ContainerAdvancedSettings) SetReadinessProbeFailureThreshold(v int32)`

SetReadinessProbeFailureThreshold sets ReadinessProbeFailureThreshold field to given value.

### HasReadinessProbeFailureThreshold

`func (o *ContainerAdvancedSettings) HasReadinessProbeFailureThreshold() bool`

HasReadinessProbeFailureThreshold returns a boolean if a field has been set.

### GetLivenessProbeType

`func (o *ContainerAdvancedSettings) GetLivenessProbeType() string`

GetLivenessProbeType returns the LivenessProbeType field if non-nil, zero value otherwise.

### GetLivenessProbeTypeOk

`func (o *ContainerAdvancedSettings) GetLivenessProbeTypeOk() (*string, bool)`

GetLivenessProbeTypeOk returns a tuple with the LivenessProbeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeType

`func (o *ContainerAdvancedSettings) SetLivenessProbeType(v string)`

SetLivenessProbeType sets LivenessProbeType field to given value.

### HasLivenessProbeType

`func (o *ContainerAdvancedSettings) HasLivenessProbeType() bool`

HasLivenessProbeType returns a boolean if a field has been set.

### GetLivenessProbeHttpGetPath

`func (o *ContainerAdvancedSettings) GetLivenessProbeHttpGetPath() string`

GetLivenessProbeHttpGetPath returns the LivenessProbeHttpGetPath field if non-nil, zero value otherwise.

### GetLivenessProbeHttpGetPathOk

`func (o *ContainerAdvancedSettings) GetLivenessProbeHttpGetPathOk() (*string, bool)`

GetLivenessProbeHttpGetPathOk returns a tuple with the LivenessProbeHttpGetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeHttpGetPath

`func (o *ContainerAdvancedSettings) SetLivenessProbeHttpGetPath(v string)`

SetLivenessProbeHttpGetPath sets LivenessProbeHttpGetPath field to given value.

### HasLivenessProbeHttpGetPath

`func (o *ContainerAdvancedSettings) HasLivenessProbeHttpGetPath() bool`

HasLivenessProbeHttpGetPath returns a boolean if a field has been set.

### GetLivenessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettings) GetLivenessProbeInitialDelaySeconds() int32`

GetLivenessProbeInitialDelaySeconds returns the LivenessProbeInitialDelaySeconds field if non-nil, zero value otherwise.

### GetLivenessProbeInitialDelaySecondsOk

`func (o *ContainerAdvancedSettings) GetLivenessProbeInitialDelaySecondsOk() (*int32, bool)`

GetLivenessProbeInitialDelaySecondsOk returns a tuple with the LivenessProbeInitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettings) SetLivenessProbeInitialDelaySeconds(v int32)`

SetLivenessProbeInitialDelaySeconds sets LivenessProbeInitialDelaySeconds field to given value.

### HasLivenessProbeInitialDelaySeconds

`func (o *ContainerAdvancedSettings) HasLivenessProbeInitialDelaySeconds() bool`

HasLivenessProbeInitialDelaySeconds returns a boolean if a field has been set.

### GetLivenessProbePeriodSeconds

`func (o *ContainerAdvancedSettings) GetLivenessProbePeriodSeconds() int32`

GetLivenessProbePeriodSeconds returns the LivenessProbePeriodSeconds field if non-nil, zero value otherwise.

### GetLivenessProbePeriodSecondsOk

`func (o *ContainerAdvancedSettings) GetLivenessProbePeriodSecondsOk() (*int32, bool)`

GetLivenessProbePeriodSecondsOk returns a tuple with the LivenessProbePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbePeriodSeconds

`func (o *ContainerAdvancedSettings) SetLivenessProbePeriodSeconds(v int32)`

SetLivenessProbePeriodSeconds sets LivenessProbePeriodSeconds field to given value.

### HasLivenessProbePeriodSeconds

`func (o *ContainerAdvancedSettings) HasLivenessProbePeriodSeconds() bool`

HasLivenessProbePeriodSeconds returns a boolean if a field has been set.

### GetLivenessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetLivenessProbeTimeoutSeconds() int32`

GetLivenessProbeTimeoutSeconds returns the LivenessProbeTimeoutSeconds field if non-nil, zero value otherwise.

### GetLivenessProbeTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetLivenessProbeTimeoutSecondsOk() (*int32, bool)`

GetLivenessProbeTimeoutSecondsOk returns a tuple with the LivenessProbeTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetLivenessProbeTimeoutSeconds(v int32)`

SetLivenessProbeTimeoutSeconds sets LivenessProbeTimeoutSeconds field to given value.

### HasLivenessProbeTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasLivenessProbeTimeoutSeconds() bool`

HasLivenessProbeTimeoutSeconds returns a boolean if a field has been set.

### GetLivenessProbeSuccessThreshold

`func (o *ContainerAdvancedSettings) GetLivenessProbeSuccessThreshold() int32`

GetLivenessProbeSuccessThreshold returns the LivenessProbeSuccessThreshold field if non-nil, zero value otherwise.

### GetLivenessProbeSuccessThresholdOk

`func (o *ContainerAdvancedSettings) GetLivenessProbeSuccessThresholdOk() (*int32, bool)`

GetLivenessProbeSuccessThresholdOk returns a tuple with the LivenessProbeSuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeSuccessThreshold

`func (o *ContainerAdvancedSettings) SetLivenessProbeSuccessThreshold(v int32)`

SetLivenessProbeSuccessThreshold sets LivenessProbeSuccessThreshold field to given value.

### HasLivenessProbeSuccessThreshold

`func (o *ContainerAdvancedSettings) HasLivenessProbeSuccessThreshold() bool`

HasLivenessProbeSuccessThreshold returns a boolean if a field has been set.

### GetLivenessProbeFailureThreshold

`func (o *ContainerAdvancedSettings) GetLivenessProbeFailureThreshold() int32`

GetLivenessProbeFailureThreshold returns the LivenessProbeFailureThreshold field if non-nil, zero value otherwise.

### GetLivenessProbeFailureThresholdOk

`func (o *ContainerAdvancedSettings) GetLivenessProbeFailureThresholdOk() (*int32, bool)`

GetLivenessProbeFailureThresholdOk returns a tuple with the LivenessProbeFailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbeFailureThreshold

`func (o *ContainerAdvancedSettings) SetLivenessProbeFailureThreshold(v int32)`

SetLivenessProbeFailureThreshold sets LivenessProbeFailureThreshold field to given value.

### HasLivenessProbeFailureThreshold

`func (o *ContainerAdvancedSettings) HasLivenessProbeFailureThreshold() bool`

HasLivenessProbeFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


