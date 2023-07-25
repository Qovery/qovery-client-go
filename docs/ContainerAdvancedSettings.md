# ContainerAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCustomDomainCheckEnabled** | Pointer to **bool** | disable custom domain check when deploying an application | [optional] [default to true]
**DeploymentTerminationGracePeriodSeconds** | Pointer to **int32** | define how long in seconds an application is supposed to be stopped gracefully | [optional] [default to 60]
**DeploymentUpdateStrategyType** | Pointer to **string** | * &#x60;RollingUpdate&#x60; gracefully rollout new versions, and automatically rollback if the new version fails to start * &#x60;Recreate&#x60; stop all current versions and create new ones once all old ones have been shutdown  | [optional] [default to "RollingUpdate"]
**DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent** | Pointer to **int32** | Define the percentage of a maximum number of pods that can be unavailable during the update process | [optional] [default to 25]
**DeploymentUpdateStrategyRollingUpdateMaxSurgePercent** | Pointer to **int32** | Define the percentage of the maximum number of pods that can be created over the desired number of pods | [optional] [default to 25]
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
**NetworkIngressGrpcSendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the grpc server | [optional] [default to 60]
**NetworkIngressGrpcReadTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the grpc server | [optional] [default to 60]
**NetworkIngressWhitelistSourceRange** | Pointer to **string** | list of source ranges to allow access to ingress proxy.  This property can be used to whitelist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1 To allow all source ranges, set 0.0.0.0/0.  | [optional] [default to "0.0.0.0/0"]
**NetworkIngressDenylistSourceRange** | Pointer to **string** | list of source ranges to deny access to ingress proxy.  This property can be used to blacklist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1  | [optional] [default to ""]
**NetworkIngressBasicAuthEnvVar** | Pointer to **string** | Set the name of an environment variable to use as a basic authentication (&#x60;login:crypted_password&#x60;) from &#x60;htpasswd&#x60; command. You can add multiples comma separated values.  | [optional] [default to ""]
**NetworkIngressEnableStickySession** | Pointer to **bool** | Enable the load balancer to bind a user&#39;s session to a specific target. This ensures that all requests from the user during the session are sent to the same target  | [optional] [default to false]
**SecurityServiceAccountName** | Pointer to **string** | Allows you to set an existing Kubernetes service account name  | [optional] [default to ""]
**HpaCpuAverageUtilizationPercent** | Pointer to **int32** | Percentage value of cpu usage at which point pods should scale up. | [optional] [default to 60]

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

### GetDeploymentUpdateStrategyType

`func (o *ContainerAdvancedSettings) GetDeploymentUpdateStrategyType() string`

GetDeploymentUpdateStrategyType returns the DeploymentUpdateStrategyType field if non-nil, zero value otherwise.

### GetDeploymentUpdateStrategyTypeOk

`func (o *ContainerAdvancedSettings) GetDeploymentUpdateStrategyTypeOk() (*string, bool)`

GetDeploymentUpdateStrategyTypeOk returns a tuple with the DeploymentUpdateStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStrategyType

`func (o *ContainerAdvancedSettings) SetDeploymentUpdateStrategyType(v string)`

SetDeploymentUpdateStrategyType sets DeploymentUpdateStrategyType field to given value.

### HasDeploymentUpdateStrategyType

`func (o *ContainerAdvancedSettings) HasDeploymentUpdateStrategyType() bool`

HasDeploymentUpdateStrategyType returns a boolean if a field has been set.

### GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent

`func (o *ContainerAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent() int32`

GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent returns the DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent field if non-nil, zero value otherwise.

### GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercentOk

`func (o *ContainerAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercentOk() (*int32, bool)`

GetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercentOk returns a tuple with the DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent

`func (o *ContainerAdvancedSettings) SetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent(v int32)`

SetDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent sets DeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent field to given value.

### HasDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent

`func (o *ContainerAdvancedSettings) HasDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent() bool`

HasDeploymentUpdateStrategyRollingUpdateMaxUnavailablePercent returns a boolean if a field has been set.

### GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent

`func (o *ContainerAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent() int32`

GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent returns the DeploymentUpdateStrategyRollingUpdateMaxSurgePercent field if non-nil, zero value otherwise.

### GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercentOk

`func (o *ContainerAdvancedSettings) GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercentOk() (*int32, bool)`

GetDeploymentUpdateStrategyRollingUpdateMaxSurgePercentOk returns a tuple with the DeploymentUpdateStrategyRollingUpdateMaxSurgePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent

`func (o *ContainerAdvancedSettings) SetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent(v int32)`

SetDeploymentUpdateStrategyRollingUpdateMaxSurgePercent sets DeploymentUpdateStrategyRollingUpdateMaxSurgePercent field to given value.

### HasDeploymentUpdateStrategyRollingUpdateMaxSurgePercent

`func (o *ContainerAdvancedSettings) HasDeploymentUpdateStrategyRollingUpdateMaxSurgePercent() bool`

HasDeploymentUpdateStrategyRollingUpdateMaxSurgePercent returns a boolean if a field has been set.

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

### GetNetworkIngressGrpcSendTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressGrpcSendTimeoutSeconds() int32`

GetNetworkIngressGrpcSendTimeoutSeconds returns the NetworkIngressGrpcSendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressGrpcSendTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressGrpcSendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressGrpcSendTimeoutSecondsOk returns a tuple with the NetworkIngressGrpcSendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressGrpcSendTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressGrpcSendTimeoutSeconds(v int32)`

SetNetworkIngressGrpcSendTimeoutSeconds sets NetworkIngressGrpcSendTimeoutSeconds field to given value.

### HasNetworkIngressGrpcSendTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressGrpcSendTimeoutSeconds() bool`

HasNetworkIngressGrpcSendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressGrpcReadTimeoutSeconds

`func (o *ContainerAdvancedSettings) GetNetworkIngressGrpcReadTimeoutSeconds() int32`

GetNetworkIngressGrpcReadTimeoutSeconds returns the NetworkIngressGrpcReadTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressGrpcReadTimeoutSecondsOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressGrpcReadTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressGrpcReadTimeoutSecondsOk returns a tuple with the NetworkIngressGrpcReadTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressGrpcReadTimeoutSeconds

`func (o *ContainerAdvancedSettings) SetNetworkIngressGrpcReadTimeoutSeconds(v int32)`

SetNetworkIngressGrpcReadTimeoutSeconds sets NetworkIngressGrpcReadTimeoutSeconds field to given value.

### HasNetworkIngressGrpcReadTimeoutSeconds

`func (o *ContainerAdvancedSettings) HasNetworkIngressGrpcReadTimeoutSeconds() bool`

HasNetworkIngressGrpcReadTimeoutSeconds returns a boolean if a field has been set.

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

### GetNetworkIngressDenylistSourceRange

`func (o *ContainerAdvancedSettings) GetNetworkIngressDenylistSourceRange() string`

GetNetworkIngressDenylistSourceRange returns the NetworkIngressDenylistSourceRange field if non-nil, zero value otherwise.

### GetNetworkIngressDenylistSourceRangeOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressDenylistSourceRangeOk() (*string, bool)`

GetNetworkIngressDenylistSourceRangeOk returns a tuple with the NetworkIngressDenylistSourceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressDenylistSourceRange

`func (o *ContainerAdvancedSettings) SetNetworkIngressDenylistSourceRange(v string)`

SetNetworkIngressDenylistSourceRange sets NetworkIngressDenylistSourceRange field to given value.

### HasNetworkIngressDenylistSourceRange

`func (o *ContainerAdvancedSettings) HasNetworkIngressDenylistSourceRange() bool`

HasNetworkIngressDenylistSourceRange returns a boolean if a field has been set.

### GetNetworkIngressBasicAuthEnvVar

`func (o *ContainerAdvancedSettings) GetNetworkIngressBasicAuthEnvVar() string`

GetNetworkIngressBasicAuthEnvVar returns the NetworkIngressBasicAuthEnvVar field if non-nil, zero value otherwise.

### GetNetworkIngressBasicAuthEnvVarOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressBasicAuthEnvVarOk() (*string, bool)`

GetNetworkIngressBasicAuthEnvVarOk returns a tuple with the NetworkIngressBasicAuthEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressBasicAuthEnvVar

`func (o *ContainerAdvancedSettings) SetNetworkIngressBasicAuthEnvVar(v string)`

SetNetworkIngressBasicAuthEnvVar sets NetworkIngressBasicAuthEnvVar field to given value.

### HasNetworkIngressBasicAuthEnvVar

`func (o *ContainerAdvancedSettings) HasNetworkIngressBasicAuthEnvVar() bool`

HasNetworkIngressBasicAuthEnvVar returns a boolean if a field has been set.

### GetNetworkIngressEnableStickySession

`func (o *ContainerAdvancedSettings) GetNetworkIngressEnableStickySession() bool`

GetNetworkIngressEnableStickySession returns the NetworkIngressEnableStickySession field if non-nil, zero value otherwise.

### GetNetworkIngressEnableStickySessionOk

`func (o *ContainerAdvancedSettings) GetNetworkIngressEnableStickySessionOk() (*bool, bool)`

GetNetworkIngressEnableStickySessionOk returns a tuple with the NetworkIngressEnableStickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableStickySession

`func (o *ContainerAdvancedSettings) SetNetworkIngressEnableStickySession(v bool)`

SetNetworkIngressEnableStickySession sets NetworkIngressEnableStickySession field to given value.

### HasNetworkIngressEnableStickySession

`func (o *ContainerAdvancedSettings) HasNetworkIngressEnableStickySession() bool`

HasNetworkIngressEnableStickySession returns a boolean if a field has been set.

### GetSecurityServiceAccountName

`func (o *ContainerAdvancedSettings) GetSecurityServiceAccountName() string`

GetSecurityServiceAccountName returns the SecurityServiceAccountName field if non-nil, zero value otherwise.

### GetSecurityServiceAccountNameOk

`func (o *ContainerAdvancedSettings) GetSecurityServiceAccountNameOk() (*string, bool)`

GetSecurityServiceAccountNameOk returns a tuple with the SecurityServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServiceAccountName

`func (o *ContainerAdvancedSettings) SetSecurityServiceAccountName(v string)`

SetSecurityServiceAccountName sets SecurityServiceAccountName field to given value.

### HasSecurityServiceAccountName

`func (o *ContainerAdvancedSettings) HasSecurityServiceAccountName() bool`

HasSecurityServiceAccountName returns a boolean if a field has been set.

### GetHpaCpuAverageUtilizationPercent

`func (o *ContainerAdvancedSettings) GetHpaCpuAverageUtilizationPercent() int32`

GetHpaCpuAverageUtilizationPercent returns the HpaCpuAverageUtilizationPercent field if non-nil, zero value otherwise.

### GetHpaCpuAverageUtilizationPercentOk

`func (o *ContainerAdvancedSettings) GetHpaCpuAverageUtilizationPercentOk() (*int32, bool)`

GetHpaCpuAverageUtilizationPercentOk returns a tuple with the HpaCpuAverageUtilizationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpaCpuAverageUtilizationPercent

`func (o *ContainerAdvancedSettings) SetHpaCpuAverageUtilizationPercent(v int32)`

SetHpaCpuAverageUtilizationPercent sets HpaCpuAverageUtilizationPercent field to given value.

### HasHpaCpuAverageUtilizationPercent

`func (o *ContainerAdvancedSettings) HasHpaCpuAverageUtilizationPercent() bool`

HasHpaCpuAverageUtilizationPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


