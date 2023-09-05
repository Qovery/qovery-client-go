# ApplicationAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCustomDomainCheckEnabled** | Pointer to **bool** | disable custom domain check when deploying an application | [optional] [default to true]
**DeploymentTerminationGracePeriodSeconds** | Pointer to **int32** | define how long in seconds an application is supposed to be stopped gracefully | [optional] [default to 60]
**DeploymentAffinityNodeRequired** | Pointer to **map[string]string** | Set pod placement on specific Kubernetes nodes labels | [optional] 
**DeploymentAntiaffinityPod** | Pointer to **string** | Define how you want pods affinity to behave: * &#x60;Preferred&#x60; allows, but does not require, pods of a given service are not co-located (or co-hosted) on a single node * &#x60;Requirred&#x60; ensures that the pods of a given service are not co-located (or co-hosted) on a single node (safer in term of availability but can be expensive depending on the number of replicas)  | [optional] [default to "Preferred"]
**DeploymentUpdateStrategyType** | Pointer to **string** | * &#x60;RollingUpdate&#x60; gracefully rollout new versions, and automatically rollback if the new version fails to start * &#x60;Recreate&#x60; stop all current versions and create new ones once all old ones have been shutdown  | [optional] [default to "RollingUpdate"]
**DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent** | Pointer to **int32** | Define the percentage of a maximum number of pods that can be unavailable during the update process | [optional] [default to 25]
**DeploymentUpdateStrategyRollingUpdateMaxSurgePercent** | Pointer to **int32** | Define the percentage of the maximum number of pods that can be created over the desired number of pods | [optional] [default to 25]
**BuildTimeoutMaxSec** | Pointer to **int32** |  | [optional] [default to 1800]
**BuildCpuMaxInMilli** | Pointer to **int32** | define the max cpu resources (in milli) | [optional] [default to 4000]
**BuildRamMaxInGib** | Pointer to **int32** | define the max ram resources (in gib) | [optional] [default to 8]
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
**NetworkIngressProxyBuffering** | Pointer to **string** | Allows to enable or disable nginx &#x60;proxy-buffering&#x60; | [optional] [default to "on"]
**NetworkIngressProxyRequestBuffering** | Pointer to **string** | Allows to enable or disable nginx &#x60;proxy-request-buffering&#x60; | [optional] [default to "on"]
**NetworkIngressWhitelistSourceRange** | Pointer to **string** | list of source ranges to allow access to ingress proxy.  This property can be used to whitelist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1 To allow all source ranges, set 0.0.0.0/0.  | [optional] [default to "0.0.0.0/0"]
**NetworkIngressDenylistSourceRange** | Pointer to **string** | list of source ranges to deny access to ingress proxy.  This property can be used to blacklist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1  | [optional] [default to ""]
**NetworkIngressBasicAuthEnvVar** | Pointer to **string** | Set the name of an environment variable to use as a basic authentication (&#x60;login:crypted_password&#x60;) from &#x60;htpasswd&#x60; command.  | [optional] [default to ""]
**NetworkIngressEnableStickySession** | Pointer to **bool** | Enable the load balancer to bind a user&#39;s session to a specific target. This ensures that all requests from the user during the session are sent to the same target  | [optional] [default to false]
**NetworkIngressGrpcSendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the grpc server | [optional] [default to 60]
**NetworkIngressGrpcReadTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the grpc server | [optional] [default to 60]
**NetworkIngressExtraHeaders** | Pointer to **string** | Allows to define response headers | [optional] [default to "{}"]
**HpaCpuAverageUtilizationPercent** | Pointer to **int32** | Percentage value of cpu usage at which point pods should scale up. | [optional] [default to 60]
**SecurityServiceAccountName** | Pointer to **string** | Allows you to set an existing Kubernetes service account name  | [optional] [default to ""]

## Methods

### NewApplicationAdvancedSettings

`func NewApplicationAdvancedSettings() *ApplicationAdvancedSettings`

NewApplicationAdvancedSettings instantiates a new ApplicationAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAdvancedSettingsWithDefaults

`func NewApplicationAdvancedSettingsWithDefaults() *ApplicationAdvancedSettings`

NewApplicationAdvancedSettingsWithDefaults instantiates a new ApplicationAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCustomDomainCheckEnabled

`func (o *ApplicationAdvancedSettings) GetDeploymentCustomDomainCheckEnabled() bool`

GetDeploymentCustomDomainCheckEnabled returns the DeploymentCustomDomainCheckEnabled field if non-nil, zero value otherwise.

### GetDeploymentCustomDomainCheckEnabledOk

`func (o *ApplicationAdvancedSettings) GetDeploymentCustomDomainCheckEnabledOk() (*bool, bool)`

GetDeploymentCustomDomainCheckEnabledOk returns a tuple with the DeploymentCustomDomainCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCustomDomainCheckEnabled

`func (o *ApplicationAdvancedSettings) SetDeploymentCustomDomainCheckEnabled(v bool)`

SetDeploymentCustomDomainCheckEnabled sets DeploymentCustomDomainCheckEnabled field to given value.

### HasDeploymentCustomDomainCheckEnabled

`func (o *ApplicationAdvancedSettings) HasDeploymentCustomDomainCheckEnabled() bool`

HasDeploymentCustomDomainCheckEnabled returns a boolean if a field has been set.

### GetDeploymentTerminationGracePeriodSeconds

`func (o *ApplicationAdvancedSettings) GetDeploymentTerminationGracePeriodSeconds() int32`

GetDeploymentTerminationGracePeriodSeconds returns the DeploymentTerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeploymentTerminationGracePeriodSecondsOk

`func (o *ApplicationAdvancedSettings) GetDeploymentTerminationGracePeriodSecondsOk() (*int32, bool)`

GetDeploymentTerminationGracePeriodSecondsOk returns a tuple with the DeploymentTerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTerminationGracePeriodSeconds

`func (o *ApplicationAdvancedSettings) SetDeploymentTerminationGracePeriodSeconds(v int32)`

SetDeploymentTerminationGracePeriodSeconds sets DeploymentTerminationGracePeriodSeconds field to given value.

### HasDeploymentTerminationGracePeriodSeconds

`func (o *ApplicationAdvancedSettings) HasDeploymentTerminationGracePeriodSeconds() bool`

HasDeploymentTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetDeploymentAffinityNodeRequired

`func (o *ApplicationAdvancedSettings) GetDeploymentAffinityNodeRequired() map[string]string`

GetDeploymentAffinityNodeRequired returns the DeploymentAffinityNodeRequired field if non-nil, zero value otherwise.

### GetDeploymentAffinityNodeRequiredOk

`func (o *ApplicationAdvancedSettings) GetDeploymentAffinityNodeRequiredOk() (*map[string]string, bool)`

GetDeploymentAffinityNodeRequiredOk returns a tuple with the DeploymentAffinityNodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentAffinityNodeRequired

`func (o *ApplicationAdvancedSettings) SetDeploymentAffinityNodeRequired(v map[string]string)`

SetDeploymentAffinityNodeRequired sets DeploymentAffinityNodeRequired field to given value.

### HasDeploymentAffinityNodeRequired

`func (o *ApplicationAdvancedSettings) HasDeploymentAffinityNodeRequired() bool`

HasDeploymentAffinityNodeRequired returns a boolean if a field has been set.

### GetDeploymentAntiaffinityPod

`func (o *ApplicationAdvancedSettings) GetDeploymentAntiaffinityPod() string`

GetDeploymentAntiaffinityPod returns the DeploymentAntiaffinityPod field if non-nil, zero value otherwise.

### GetDeploymentAntiaffinityPodOk

`func (o *ApplicationAdvancedSettings) GetDeploymentAntiaffinityPodOk() (*string, bool)`

GetDeploymentAntiaffinityPodOk returns a tuple with the DeploymentAntiaffinityPod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentAntiaffinityPod

`func (o *ApplicationAdvancedSettings) SetDeploymentAntiaffinityPod(v string)`

SetDeploymentAntiaffinityPod sets DeploymentAntiaffinityPod field to given value.

### HasDeploymentAntiaffinityPod

`func (o *ApplicationAdvancedSettings) HasDeploymentAntiaffinityPod() bool`

HasDeploymentAntiaffinityPod returns a boolean if a field has been set.

### GetDeploymentUpdateStrategyType

`func (o *ApplicationAdvancedSettings) GetDeploymentUpdateStrategyType() string`

GetDeploymentUpdateStrategyType returns the DeploymentUpdateStrategyType field if non-nil, zero value otherwise.

### GetDeploymentUpdateStrategyTypeOk

`func (o *ApplicationAdvancedSettings) GetDeploymentUpdateStrategyTypeOk() (*string, bool)`

GetDeploymentUpdateStrategyTypeOk returns a tuple with the DeploymentUpdateStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStrategyType

`func (o *ApplicationAdvancedSettings) SetDeploymentUpdateStrategyType(v string)`

SetDeploymentUpdateStrategyType sets DeploymentUpdateStrategyType field to given value.

### HasDeploymentUpdateStrategyType

`func (o *ApplicationAdvancedSettings) HasDeploymentUpdateStrategyType() bool`

HasDeploymentUpdateStrategyType returns a boolean if a field has been set.

### GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent

`func (o *ApplicationAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent() int32`

GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent returns the DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent field if non-nil, zero value otherwise.

### GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercentOk

`func (o *ApplicationAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercentOk() (*int32, bool)`

GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercentOk returns a tuple with the DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent

`func (o *ApplicationAdvancedSettings) SetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent(v int32)`

SetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent sets DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent field to given value.

### HasDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent

`func (o *ApplicationAdvancedSettings) HasDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent() bool`

HasDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent returns a boolean if a field has been set.

### GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent

`func (o *ApplicationAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent() int32`

GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent returns the DeploymentUpdateStrategyRollingUpdateMaxSurgePercent field if non-nil, zero value otherwise.

### GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercentOk

`func (o *ApplicationAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercentOk() (*int32, bool)`

GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercentOk returns a tuple with the DeploymentUpdateStrategyRollingUpdateMaxSurgePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent

`func (o *ApplicationAdvancedSettings) SetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent(v int32)`

SetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent sets DeploymentUpdateStrategyRollingUpdateMaxSurgePercent field to given value.

### HasDeploymentUpdateStrategyRollingUpdateMaxSurgePercent

`func (o *ApplicationAdvancedSettings) HasDeploymentUpdateStrategyRollingUpdateMaxSurgePercent() bool`

HasDeploymentUpdateStrategyRollingUpdateMaxSurgePercent returns a boolean if a field has been set.

### GetBuildTimeoutMaxSec

`func (o *ApplicationAdvancedSettings) GetBuildTimeoutMaxSec() int32`

GetBuildTimeoutMaxSec returns the BuildTimeoutMaxSec field if non-nil, zero value otherwise.

### GetBuildTimeoutMaxSecOk

`func (o *ApplicationAdvancedSettings) GetBuildTimeoutMaxSecOk() (*int32, bool)`

GetBuildTimeoutMaxSecOk returns a tuple with the BuildTimeoutMaxSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimeoutMaxSec

`func (o *ApplicationAdvancedSettings) SetBuildTimeoutMaxSec(v int32)`

SetBuildTimeoutMaxSec sets BuildTimeoutMaxSec field to given value.

### HasBuildTimeoutMaxSec

`func (o *ApplicationAdvancedSettings) HasBuildTimeoutMaxSec() bool`

HasBuildTimeoutMaxSec returns a boolean if a field has been set.

### GetBuildCpuMaxInMilli

`func (o *ApplicationAdvancedSettings) GetBuildCpuMaxInMilli() int32`

GetBuildCpuMaxInMilli returns the BuildCpuMaxInMilli field if non-nil, zero value otherwise.

### GetBuildCpuMaxInMilliOk

`func (o *ApplicationAdvancedSettings) GetBuildCpuMaxInMilliOk() (*int32, bool)`

GetBuildCpuMaxInMilliOk returns a tuple with the BuildCpuMaxInMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCpuMaxInMilli

`func (o *ApplicationAdvancedSettings) SetBuildCpuMaxInMilli(v int32)`

SetBuildCpuMaxInMilli sets BuildCpuMaxInMilli field to given value.

### HasBuildCpuMaxInMilli

`func (o *ApplicationAdvancedSettings) HasBuildCpuMaxInMilli() bool`

HasBuildCpuMaxInMilli returns a boolean if a field has been set.

### GetBuildRamMaxInGib

`func (o *ApplicationAdvancedSettings) GetBuildRamMaxInGib() int32`

GetBuildRamMaxInGib returns the BuildRamMaxInGib field if non-nil, zero value otherwise.

### GetBuildRamMaxInGibOk

`func (o *ApplicationAdvancedSettings) GetBuildRamMaxInGibOk() (*int32, bool)`

GetBuildRamMaxInGibOk returns a tuple with the BuildRamMaxInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRamMaxInGib

`func (o *ApplicationAdvancedSettings) SetBuildRamMaxInGib(v int32)`

SetBuildRamMaxInGib sets BuildRamMaxInGib field to given value.

### HasBuildRamMaxInGib

`func (o *ApplicationAdvancedSettings) HasBuildRamMaxInGib() bool`

HasBuildRamMaxInGib returns a boolean if a field has been set.

### GetNetworkIngressProxyBodySizeMb

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyBodySizeMb() int32`

GetNetworkIngressProxyBodySizeMb returns the NetworkIngressProxyBodySizeMb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBodySizeMbOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyBodySizeMbOk() (*int32, bool)`

GetNetworkIngressProxyBodySizeMbOk returns a tuple with the NetworkIngressProxyBodySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBodySizeMb

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxyBodySizeMb(v int32)`

SetNetworkIngressProxyBodySizeMb sets NetworkIngressProxyBodySizeMb field to given value.

### HasNetworkIngressProxyBodySizeMb

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxyBodySizeMb() bool`

HasNetworkIngressProxyBodySizeMb returns a boolean if a field has been set.

### GetNetworkIngressEnableCors

`func (o *ApplicationAdvancedSettings) GetNetworkIngressEnableCors() bool`

GetNetworkIngressEnableCors returns the NetworkIngressEnableCors field if non-nil, zero value otherwise.

### GetNetworkIngressEnableCorsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressEnableCorsOk() (*bool, bool)`

GetNetworkIngressEnableCorsOk returns a tuple with the NetworkIngressEnableCors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableCors

`func (o *ApplicationAdvancedSettings) SetNetworkIngressEnableCors(v bool)`

SetNetworkIngressEnableCors sets NetworkIngressEnableCors field to given value.

### HasNetworkIngressEnableCors

`func (o *ApplicationAdvancedSettings) HasNetworkIngressEnableCors() bool`

HasNetworkIngressEnableCors returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowOrigin

`func (o *ApplicationAdvancedSettings) GetNetworkIngressCorsAllowOrigin() string`

GetNetworkIngressCorsAllowOrigin returns the NetworkIngressCorsAllowOrigin field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowOriginOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressCorsAllowOriginOk() (*string, bool)`

GetNetworkIngressCorsAllowOriginOk returns a tuple with the NetworkIngressCorsAllowOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowOrigin

`func (o *ApplicationAdvancedSettings) SetNetworkIngressCorsAllowOrigin(v string)`

SetNetworkIngressCorsAllowOrigin sets NetworkIngressCorsAllowOrigin field to given value.

### HasNetworkIngressCorsAllowOrigin

`func (o *ApplicationAdvancedSettings) HasNetworkIngressCorsAllowOrigin() bool`

HasNetworkIngressCorsAllowOrigin returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowMethods

`func (o *ApplicationAdvancedSettings) GetNetworkIngressCorsAllowMethods() string`

GetNetworkIngressCorsAllowMethods returns the NetworkIngressCorsAllowMethods field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowMethodsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressCorsAllowMethodsOk() (*string, bool)`

GetNetworkIngressCorsAllowMethodsOk returns a tuple with the NetworkIngressCorsAllowMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowMethods

`func (o *ApplicationAdvancedSettings) SetNetworkIngressCorsAllowMethods(v string)`

SetNetworkIngressCorsAllowMethods sets NetworkIngressCorsAllowMethods field to given value.

### HasNetworkIngressCorsAllowMethods

`func (o *ApplicationAdvancedSettings) HasNetworkIngressCorsAllowMethods() bool`

HasNetworkIngressCorsAllowMethods returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowHeaders

`func (o *ApplicationAdvancedSettings) GetNetworkIngressCorsAllowHeaders() string`

GetNetworkIngressCorsAllowHeaders returns the NetworkIngressCorsAllowHeaders field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowHeadersOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressCorsAllowHeadersOk() (*string, bool)`

GetNetworkIngressCorsAllowHeadersOk returns a tuple with the NetworkIngressCorsAllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowHeaders

`func (o *ApplicationAdvancedSettings) SetNetworkIngressCorsAllowHeaders(v string)`

SetNetworkIngressCorsAllowHeaders sets NetworkIngressCorsAllowHeaders field to given value.

### HasNetworkIngressCorsAllowHeaders

`func (o *ApplicationAdvancedSettings) HasNetworkIngressCorsAllowHeaders() bool`

HasNetworkIngressCorsAllowHeaders returns a boolean if a field has been set.

### GetNetworkIngressProxyBufferSizeKb

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyBufferSizeKb() int32`

GetNetworkIngressProxyBufferSizeKb returns the NetworkIngressProxyBufferSizeKb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBufferSizeKbOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyBufferSizeKbOk() (*int32, bool)`

GetNetworkIngressProxyBufferSizeKbOk returns a tuple with the NetworkIngressProxyBufferSizeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBufferSizeKb

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxyBufferSizeKb(v int32)`

SetNetworkIngressProxyBufferSizeKb sets NetworkIngressProxyBufferSizeKb field to given value.

### HasNetworkIngressProxyBufferSizeKb

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxyBufferSizeKb() bool`

HasNetworkIngressProxyBufferSizeKb returns a boolean if a field has been set.

### GetNetworkIngressKeepaliveTimeSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressKeepaliveTimeSeconds() int32`

GetNetworkIngressKeepaliveTimeSeconds returns the NetworkIngressKeepaliveTimeSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressKeepaliveTimeSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressKeepaliveTimeSecondsOk() (*int32, bool)`

GetNetworkIngressKeepaliveTimeSecondsOk returns a tuple with the NetworkIngressKeepaliveTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressKeepaliveTimeSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressKeepaliveTimeSeconds(v int32)`

SetNetworkIngressKeepaliveTimeSeconds sets NetworkIngressKeepaliveTimeSeconds field to given value.

### HasNetworkIngressKeepaliveTimeSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressKeepaliveTimeSeconds() bool`

HasNetworkIngressKeepaliveTimeSeconds returns a boolean if a field has been set.

### GetNetworkIngressKeepaliveTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressKeepaliveTimeoutSeconds() int32`

GetNetworkIngressKeepaliveTimeoutSeconds returns the NetworkIngressKeepaliveTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressKeepaliveTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressKeepaliveTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressKeepaliveTimeoutSecondsOk returns a tuple with the NetworkIngressKeepaliveTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressKeepaliveTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressKeepaliveTimeoutSeconds(v int32)`

SetNetworkIngressKeepaliveTimeoutSeconds sets NetworkIngressKeepaliveTimeoutSeconds field to given value.

### HasNetworkIngressKeepaliveTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressKeepaliveTimeoutSeconds() bool`

HasNetworkIngressKeepaliveTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressSendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressSendTimeoutSeconds() int32`

GetNetworkIngressSendTimeoutSeconds returns the NetworkIngressSendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressSendTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressSendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressSendTimeoutSecondsOk returns a tuple with the NetworkIngressSendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressSendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressSendTimeoutSeconds(v int32)`

SetNetworkIngressSendTimeoutSeconds sets NetworkIngressSendTimeoutSeconds field to given value.

### HasNetworkIngressSendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressSendTimeoutSeconds() bool`

HasNetworkIngressSendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyConnectTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyConnectTimeoutSeconds() int32`

GetNetworkIngressProxyConnectTimeoutSeconds returns the NetworkIngressProxyConnectTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxyConnectTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyConnectTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxyConnectTimeoutSecondsOk returns a tuple with the NetworkIngressProxyConnectTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyConnectTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxyConnectTimeoutSeconds(v int32)`

SetNetworkIngressProxyConnectTimeoutSeconds sets NetworkIngressProxyConnectTimeoutSeconds field to given value.

### HasNetworkIngressProxyConnectTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxyConnectTimeoutSeconds() bool`

HasNetworkIngressProxyConnectTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxySendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxySendTimeoutSeconds() int32`

GetNetworkIngressProxySendTimeoutSeconds returns the NetworkIngressProxySendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxySendTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxySendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxySendTimeoutSecondsOk returns a tuple with the NetworkIngressProxySendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxySendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxySendTimeoutSeconds(v int32)`

SetNetworkIngressProxySendTimeoutSeconds sets NetworkIngressProxySendTimeoutSeconds field to given value.

### HasNetworkIngressProxySendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxySendTimeoutSeconds() bool`

HasNetworkIngressProxySendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyReadTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyReadTimeoutSeconds() int32`

GetNetworkIngressProxyReadTimeoutSeconds returns the NetworkIngressProxyReadTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxyReadTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyReadTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxyReadTimeoutSecondsOk returns a tuple with the NetworkIngressProxyReadTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyReadTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxyReadTimeoutSeconds(v int32)`

SetNetworkIngressProxyReadTimeoutSeconds sets NetworkIngressProxyReadTimeoutSeconds field to given value.

### HasNetworkIngressProxyReadTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxyReadTimeoutSeconds() bool`

HasNetworkIngressProxyReadTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyBuffering

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyBuffering() string`

GetNetworkIngressProxyBuffering returns the NetworkIngressProxyBuffering field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBufferingOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyBufferingOk() (*string, bool)`

GetNetworkIngressProxyBufferingOk returns a tuple with the NetworkIngressProxyBuffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBuffering

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxyBuffering(v string)`

SetNetworkIngressProxyBuffering sets NetworkIngressProxyBuffering field to given value.

### HasNetworkIngressProxyBuffering

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxyBuffering() bool`

HasNetworkIngressProxyBuffering returns a boolean if a field has been set.

### GetNetworkIngressProxyRequestBuffering

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyRequestBuffering() string`

GetNetworkIngressProxyRequestBuffering returns the NetworkIngressProxyRequestBuffering field if non-nil, zero value otherwise.

### GetNetworkIngressProxyRequestBufferingOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressProxyRequestBufferingOk() (*string, bool)`

GetNetworkIngressProxyRequestBufferingOk returns a tuple with the NetworkIngressProxyRequestBuffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyRequestBuffering

`func (o *ApplicationAdvancedSettings) SetNetworkIngressProxyRequestBuffering(v string)`

SetNetworkIngressProxyRequestBuffering sets NetworkIngressProxyRequestBuffering field to given value.

### HasNetworkIngressProxyRequestBuffering

`func (o *ApplicationAdvancedSettings) HasNetworkIngressProxyRequestBuffering() bool`

HasNetworkIngressProxyRequestBuffering returns a boolean if a field has been set.

### GetNetworkIngressWhitelistSourceRange

`func (o *ApplicationAdvancedSettings) GetNetworkIngressWhitelistSourceRange() string`

GetNetworkIngressWhitelistSourceRange returns the NetworkIngressWhitelistSourceRange field if non-nil, zero value otherwise.

### GetNetworkIngressWhitelistSourceRangeOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressWhitelistSourceRangeOk() (*string, bool)`

GetNetworkIngressWhitelistSourceRangeOk returns a tuple with the NetworkIngressWhitelistSourceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressWhitelistSourceRange

`func (o *ApplicationAdvancedSettings) SetNetworkIngressWhitelistSourceRange(v string)`

SetNetworkIngressWhitelistSourceRange sets NetworkIngressWhitelistSourceRange field to given value.

### HasNetworkIngressWhitelistSourceRange

`func (o *ApplicationAdvancedSettings) HasNetworkIngressWhitelistSourceRange() bool`

HasNetworkIngressWhitelistSourceRange returns a boolean if a field has been set.

### GetNetworkIngressDenylistSourceRange

`func (o *ApplicationAdvancedSettings) GetNetworkIngressDenylistSourceRange() string`

GetNetworkIngressDenylistSourceRange returns the NetworkIngressDenylistSourceRange field if non-nil, zero value otherwise.

### GetNetworkIngressDenylistSourceRangeOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressDenylistSourceRangeOk() (*string, bool)`

GetNetworkIngressDenylistSourceRangeOk returns a tuple with the NetworkIngressDenylistSourceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressDenylistSourceRange

`func (o *ApplicationAdvancedSettings) SetNetworkIngressDenylistSourceRange(v string)`

SetNetworkIngressDenylistSourceRange sets NetworkIngressDenylistSourceRange field to given value.

### HasNetworkIngressDenylistSourceRange

`func (o *ApplicationAdvancedSettings) HasNetworkIngressDenylistSourceRange() bool`

HasNetworkIngressDenylistSourceRange returns a boolean if a field has been set.

### GetNetworkIngressBasicAuthEnvVar

`func (o *ApplicationAdvancedSettings) GetNetworkIngressBasicAuthEnvVar() string`

GetNetworkIngressBasicAuthEnvVar returns the NetworkIngressBasicAuthEnvVar field if non-nil, zero value otherwise.

### GetNetworkIngressBasicAuthEnvVarOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressBasicAuthEnvVarOk() (*string, bool)`

GetNetworkIngressBasicAuthEnvVarOk returns a tuple with the NetworkIngressBasicAuthEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressBasicAuthEnvVar

`func (o *ApplicationAdvancedSettings) SetNetworkIngressBasicAuthEnvVar(v string)`

SetNetworkIngressBasicAuthEnvVar sets NetworkIngressBasicAuthEnvVar field to given value.

### HasNetworkIngressBasicAuthEnvVar

`func (o *ApplicationAdvancedSettings) HasNetworkIngressBasicAuthEnvVar() bool`

HasNetworkIngressBasicAuthEnvVar returns a boolean if a field has been set.

### GetNetworkIngressEnableStickySession

`func (o *ApplicationAdvancedSettings) GetNetworkIngressEnableStickySession() bool`

GetNetworkIngressEnableStickySession returns the NetworkIngressEnableStickySession field if non-nil, zero value otherwise.

### GetNetworkIngressEnableStickySessionOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressEnableStickySessionOk() (*bool, bool)`

GetNetworkIngressEnableStickySessionOk returns a tuple with the NetworkIngressEnableStickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableStickySession

`func (o *ApplicationAdvancedSettings) SetNetworkIngressEnableStickySession(v bool)`

SetNetworkIngressEnableStickySession sets NetworkIngressEnableStickySession field to given value.

### HasNetworkIngressEnableStickySession

`func (o *ApplicationAdvancedSettings) HasNetworkIngressEnableStickySession() bool`

HasNetworkIngressEnableStickySession returns a boolean if a field has been set.

### GetNetworkIngressGrpcSendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressGrpcSendTimeoutSeconds() int32`

GetNetworkIngressGrpcSendTimeoutSeconds returns the NetworkIngressGrpcSendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressGrpcSendTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressGrpcSendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressGrpcSendTimeoutSecondsOk returns a tuple with the NetworkIngressGrpcSendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressGrpcSendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressGrpcSendTimeoutSeconds(v int32)`

SetNetworkIngressGrpcSendTimeoutSeconds sets NetworkIngressGrpcSendTimeoutSeconds field to given value.

### HasNetworkIngressGrpcSendTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressGrpcSendTimeoutSeconds() bool`

HasNetworkIngressGrpcSendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressGrpcReadTimeoutSeconds

`func (o *ApplicationAdvancedSettings) GetNetworkIngressGrpcReadTimeoutSeconds() int32`

GetNetworkIngressGrpcReadTimeoutSeconds returns the NetworkIngressGrpcReadTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressGrpcReadTimeoutSecondsOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressGrpcReadTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressGrpcReadTimeoutSecondsOk returns a tuple with the NetworkIngressGrpcReadTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressGrpcReadTimeoutSeconds

`func (o *ApplicationAdvancedSettings) SetNetworkIngressGrpcReadTimeoutSeconds(v int32)`

SetNetworkIngressGrpcReadTimeoutSeconds sets NetworkIngressGrpcReadTimeoutSeconds field to given value.

### HasNetworkIngressGrpcReadTimeoutSeconds

`func (o *ApplicationAdvancedSettings) HasNetworkIngressGrpcReadTimeoutSeconds() bool`

HasNetworkIngressGrpcReadTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressExtraHeaders

`func (o *ApplicationAdvancedSettings) GetNetworkIngressExtraHeaders() string`

GetNetworkIngressExtraHeaders returns the NetworkIngressExtraHeaders field if non-nil, zero value otherwise.

### GetNetworkIngressExtraHeadersOk

`func (o *ApplicationAdvancedSettings) GetNetworkIngressExtraHeadersOk() (*string, bool)`

GetNetworkIngressExtraHeadersOk returns a tuple with the NetworkIngressExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressExtraHeaders

`func (o *ApplicationAdvancedSettings) SetNetworkIngressExtraHeaders(v string)`

SetNetworkIngressExtraHeaders sets NetworkIngressExtraHeaders field to given value.

### HasNetworkIngressExtraHeaders

`func (o *ApplicationAdvancedSettings) HasNetworkIngressExtraHeaders() bool`

HasNetworkIngressExtraHeaders returns a boolean if a field has been set.

### GetHpaCpuAverageUtilizationPercent

`func (o *ApplicationAdvancedSettings) GetHpaCpuAverageUtilizationPercent() int32`

GetHpaCpuAverageUtilizationPercent returns the HpaCpuAverageUtilizationPercent field if non-nil, zero value otherwise.

### GetHpaCpuAverageUtilizationPercentOk

`func (o *ApplicationAdvancedSettings) GetHpaCpuAverageUtilizationPercentOk() (*int32, bool)`

GetHpaCpuAverageUtilizationPercentOk returns a tuple with the HpaCpuAverageUtilizationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpaCpuAverageUtilizationPercent

`func (o *ApplicationAdvancedSettings) SetHpaCpuAverageUtilizationPercent(v int32)`

SetHpaCpuAverageUtilizationPercent sets HpaCpuAverageUtilizationPercent field to given value.

### HasHpaCpuAverageUtilizationPercent

`func (o *ApplicationAdvancedSettings) HasHpaCpuAverageUtilizationPercent() bool`

HasHpaCpuAverageUtilizationPercent returns a boolean if a field has been set.

### GetSecurityServiceAccountName

`func (o *ApplicationAdvancedSettings) GetSecurityServiceAccountName() string`

GetSecurityServiceAccountName returns the SecurityServiceAccountName field if non-nil, zero value otherwise.

### GetSecurityServiceAccountNameOk

`func (o *ApplicationAdvancedSettings) GetSecurityServiceAccountNameOk() (*string, bool)`

GetSecurityServiceAccountNameOk returns a tuple with the SecurityServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServiceAccountName

`func (o *ApplicationAdvancedSettings) SetSecurityServiceAccountName(v string)`

SetSecurityServiceAccountName sets SecurityServiceAccountName field to given value.

### HasSecurityServiceAccountName

`func (o *ApplicationAdvancedSettings) HasSecurityServiceAccountName() bool`

HasSecurityServiceAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


