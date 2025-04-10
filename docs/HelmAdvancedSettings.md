# HelmAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCustomDomainCheckEnabled** | Pointer to **bool** | disable custom domain check when deploying a helm | [optional] 
**NetworkIngressProxyBodySizeMb** | Pointer to **int32** |  | [optional] 
**NetworkIngressForceSslRedirect** | Pointer to **bool** | When using SSL offloading outside of cluster, you can enforce a redirect to HTTPS even when there is no TLS certificate available | [optional] 
**NetworkIngressEnableCors** | Pointer to **bool** |  | [optional] 
**NetworkIngressCorsAllowOrigin** | Pointer to **string** |  | [optional] 
**NetworkIngressCorsAllowMethods** | Pointer to **string** |  | [optional] 
**NetworkIngressCorsAllowHeaders** | Pointer to **string** |  | [optional] 
**NetworkIngressProxyBufferSizeKb** | Pointer to **int32** | header buffer size used while reading response header from upstream | [optional] 
**NetworkIngressKeepaliveTimeSeconds** | Pointer to **int32** | Limits the maximum time (in seconds) during which requests can be processed through one keepalive connection | [optional] 
**NetworkIngressKeepaliveTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) during which an idle keepalive connection to an upstream server will stay open. | [optional] 
**NetworkIngressSendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a response to the client | [optional] 
**NetworkIngressProxyConnectTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for establishing a connection to a proxied server | [optional] 
**NetworkIngressProxySendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the proxied server | [optional] 
**NetworkIngressProxyReadTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for reading a response from the proxied server | [optional] 
**NetworkIngressProxyBuffering** | Pointer to **string** | Allows to enable or disable nginx &#x60;proxy-buffering&#x60; | [optional] 
**NetworkIngressProxyRequestBuffering** | Pointer to **string** | Allows to enable or disable nginx &#x60;proxy-request-buffering&#x60; | [optional] 
**NetworkIngressGrpcSendTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the grpc server | [optional] 
**NetworkIngressGrpcReadTimeoutSeconds** | Pointer to **int32** | Sets a timeout (in seconds) for transmitting a request to the grpc server | [optional] 
**NetworkIngressWhitelistSourceRange** | Pointer to **string** | list of source ranges to allow access to ingress proxy.  This property can be used to whitelist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1 To allow all source ranges, set 0.0.0.0/0.  | [optional] 
**NetworkIngressDenylistSourceRange** | Pointer to **string** | list of source ranges to deny access to ingress proxy.  This property can be used to blacklist source IP ranges for ingress proxy. The value is a comma separated list of CIDRs, e.g. 10.0.0.0/24,172.10.0.1  | [optional] 
**NetworkIngressExtraHeaders** | Pointer to **string** | Allows to define response headers | [optional] 
**NetworkIngressBasicAuthEnvVar** | Pointer to **string** | Set the name of an environment variable to use as a basic authentication (&#x60;login:crypted_password&#x60;) from &#x60;htpasswd&#x60; command. You can add multiples comma separated values.  | [optional] 
**NetworkIngressEnableStickySession** | Pointer to **bool** | Enable the load balancer to bind a user&#39;s session to a specific target. This ensures that all requests from the user during the session are sent to the same target  | [optional] 

## Methods

### NewHelmAdvancedSettings

`func NewHelmAdvancedSettings() *HelmAdvancedSettings`

NewHelmAdvancedSettings instantiates a new HelmAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmAdvancedSettingsWithDefaults

`func NewHelmAdvancedSettingsWithDefaults() *HelmAdvancedSettings`

NewHelmAdvancedSettingsWithDefaults instantiates a new HelmAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCustomDomainCheckEnabled

`func (o *HelmAdvancedSettings) GetDeploymentCustomDomainCheckEnabled() bool`

GetDeploymentCustomDomainCheckEnabled returns the DeploymentCustomDomainCheckEnabled field if non-nil, zero value otherwise.

### GetDeploymentCustomDomainCheckEnabledOk

`func (o *HelmAdvancedSettings) GetDeploymentCustomDomainCheckEnabledOk() (*bool, bool)`

GetDeploymentCustomDomainCheckEnabledOk returns a tuple with the DeploymentCustomDomainCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCustomDomainCheckEnabled

`func (o *HelmAdvancedSettings) SetDeploymentCustomDomainCheckEnabled(v bool)`

SetDeploymentCustomDomainCheckEnabled sets DeploymentCustomDomainCheckEnabled field to given value.

### HasDeploymentCustomDomainCheckEnabled

`func (o *HelmAdvancedSettings) HasDeploymentCustomDomainCheckEnabled() bool`

HasDeploymentCustomDomainCheckEnabled returns a boolean if a field has been set.

### GetNetworkIngressProxyBodySizeMb

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyBodySizeMb() int32`

GetNetworkIngressProxyBodySizeMb returns the NetworkIngressProxyBodySizeMb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBodySizeMbOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyBodySizeMbOk() (*int32, bool)`

GetNetworkIngressProxyBodySizeMbOk returns a tuple with the NetworkIngressProxyBodySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBodySizeMb

`func (o *HelmAdvancedSettings) SetNetworkIngressProxyBodySizeMb(v int32)`

SetNetworkIngressProxyBodySizeMb sets NetworkIngressProxyBodySizeMb field to given value.

### HasNetworkIngressProxyBodySizeMb

`func (o *HelmAdvancedSettings) HasNetworkIngressProxyBodySizeMb() bool`

HasNetworkIngressProxyBodySizeMb returns a boolean if a field has been set.

### GetNetworkIngressForceSslRedirect

`func (o *HelmAdvancedSettings) GetNetworkIngressForceSslRedirect() bool`

GetNetworkIngressForceSslRedirect returns the NetworkIngressForceSslRedirect field if non-nil, zero value otherwise.

### GetNetworkIngressForceSslRedirectOk

`func (o *HelmAdvancedSettings) GetNetworkIngressForceSslRedirectOk() (*bool, bool)`

GetNetworkIngressForceSslRedirectOk returns a tuple with the NetworkIngressForceSslRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressForceSslRedirect

`func (o *HelmAdvancedSettings) SetNetworkIngressForceSslRedirect(v bool)`

SetNetworkIngressForceSslRedirect sets NetworkIngressForceSslRedirect field to given value.

### HasNetworkIngressForceSslRedirect

`func (o *HelmAdvancedSettings) HasNetworkIngressForceSslRedirect() bool`

HasNetworkIngressForceSslRedirect returns a boolean if a field has been set.

### GetNetworkIngressEnableCors

`func (o *HelmAdvancedSettings) GetNetworkIngressEnableCors() bool`

GetNetworkIngressEnableCors returns the NetworkIngressEnableCors field if non-nil, zero value otherwise.

### GetNetworkIngressEnableCorsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressEnableCorsOk() (*bool, bool)`

GetNetworkIngressEnableCorsOk returns a tuple with the NetworkIngressEnableCors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableCors

`func (o *HelmAdvancedSettings) SetNetworkIngressEnableCors(v bool)`

SetNetworkIngressEnableCors sets NetworkIngressEnableCors field to given value.

### HasNetworkIngressEnableCors

`func (o *HelmAdvancedSettings) HasNetworkIngressEnableCors() bool`

HasNetworkIngressEnableCors returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowOrigin

`func (o *HelmAdvancedSettings) GetNetworkIngressCorsAllowOrigin() string`

GetNetworkIngressCorsAllowOrigin returns the NetworkIngressCorsAllowOrigin field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowOriginOk

`func (o *HelmAdvancedSettings) GetNetworkIngressCorsAllowOriginOk() (*string, bool)`

GetNetworkIngressCorsAllowOriginOk returns a tuple with the NetworkIngressCorsAllowOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowOrigin

`func (o *HelmAdvancedSettings) SetNetworkIngressCorsAllowOrigin(v string)`

SetNetworkIngressCorsAllowOrigin sets NetworkIngressCorsAllowOrigin field to given value.

### HasNetworkIngressCorsAllowOrigin

`func (o *HelmAdvancedSettings) HasNetworkIngressCorsAllowOrigin() bool`

HasNetworkIngressCorsAllowOrigin returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowMethods

`func (o *HelmAdvancedSettings) GetNetworkIngressCorsAllowMethods() string`

GetNetworkIngressCorsAllowMethods returns the NetworkIngressCorsAllowMethods field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowMethodsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressCorsAllowMethodsOk() (*string, bool)`

GetNetworkIngressCorsAllowMethodsOk returns a tuple with the NetworkIngressCorsAllowMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowMethods

`func (o *HelmAdvancedSettings) SetNetworkIngressCorsAllowMethods(v string)`

SetNetworkIngressCorsAllowMethods sets NetworkIngressCorsAllowMethods field to given value.

### HasNetworkIngressCorsAllowMethods

`func (o *HelmAdvancedSettings) HasNetworkIngressCorsAllowMethods() bool`

HasNetworkIngressCorsAllowMethods returns a boolean if a field has been set.

### GetNetworkIngressCorsAllowHeaders

`func (o *HelmAdvancedSettings) GetNetworkIngressCorsAllowHeaders() string`

GetNetworkIngressCorsAllowHeaders returns the NetworkIngressCorsAllowHeaders field if non-nil, zero value otherwise.

### GetNetworkIngressCorsAllowHeadersOk

`func (o *HelmAdvancedSettings) GetNetworkIngressCorsAllowHeadersOk() (*string, bool)`

GetNetworkIngressCorsAllowHeadersOk returns a tuple with the NetworkIngressCorsAllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressCorsAllowHeaders

`func (o *HelmAdvancedSettings) SetNetworkIngressCorsAllowHeaders(v string)`

SetNetworkIngressCorsAllowHeaders sets NetworkIngressCorsAllowHeaders field to given value.

### HasNetworkIngressCorsAllowHeaders

`func (o *HelmAdvancedSettings) HasNetworkIngressCorsAllowHeaders() bool`

HasNetworkIngressCorsAllowHeaders returns a boolean if a field has been set.

### GetNetworkIngressProxyBufferSizeKb

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyBufferSizeKb() int32`

GetNetworkIngressProxyBufferSizeKb returns the NetworkIngressProxyBufferSizeKb field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBufferSizeKbOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyBufferSizeKbOk() (*int32, bool)`

GetNetworkIngressProxyBufferSizeKbOk returns a tuple with the NetworkIngressProxyBufferSizeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBufferSizeKb

`func (o *HelmAdvancedSettings) SetNetworkIngressProxyBufferSizeKb(v int32)`

SetNetworkIngressProxyBufferSizeKb sets NetworkIngressProxyBufferSizeKb field to given value.

### HasNetworkIngressProxyBufferSizeKb

`func (o *HelmAdvancedSettings) HasNetworkIngressProxyBufferSizeKb() bool`

HasNetworkIngressProxyBufferSizeKb returns a boolean if a field has been set.

### GetNetworkIngressKeepaliveTimeSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressKeepaliveTimeSeconds() int32`

GetNetworkIngressKeepaliveTimeSeconds returns the NetworkIngressKeepaliveTimeSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressKeepaliveTimeSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressKeepaliveTimeSecondsOk() (*int32, bool)`

GetNetworkIngressKeepaliveTimeSecondsOk returns a tuple with the NetworkIngressKeepaliveTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressKeepaliveTimeSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressKeepaliveTimeSeconds(v int32)`

SetNetworkIngressKeepaliveTimeSeconds sets NetworkIngressKeepaliveTimeSeconds field to given value.

### HasNetworkIngressKeepaliveTimeSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressKeepaliveTimeSeconds() bool`

HasNetworkIngressKeepaliveTimeSeconds returns a boolean if a field has been set.

### GetNetworkIngressKeepaliveTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressKeepaliveTimeoutSeconds() int32`

GetNetworkIngressKeepaliveTimeoutSeconds returns the NetworkIngressKeepaliveTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressKeepaliveTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressKeepaliveTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressKeepaliveTimeoutSecondsOk returns a tuple with the NetworkIngressKeepaliveTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressKeepaliveTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressKeepaliveTimeoutSeconds(v int32)`

SetNetworkIngressKeepaliveTimeoutSeconds sets NetworkIngressKeepaliveTimeoutSeconds field to given value.

### HasNetworkIngressKeepaliveTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressKeepaliveTimeoutSeconds() bool`

HasNetworkIngressKeepaliveTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressSendTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressSendTimeoutSeconds() int32`

GetNetworkIngressSendTimeoutSeconds returns the NetworkIngressSendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressSendTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressSendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressSendTimeoutSecondsOk returns a tuple with the NetworkIngressSendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressSendTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressSendTimeoutSeconds(v int32)`

SetNetworkIngressSendTimeoutSeconds sets NetworkIngressSendTimeoutSeconds field to given value.

### HasNetworkIngressSendTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressSendTimeoutSeconds() bool`

HasNetworkIngressSendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyConnectTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyConnectTimeoutSeconds() int32`

GetNetworkIngressProxyConnectTimeoutSeconds returns the NetworkIngressProxyConnectTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxyConnectTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyConnectTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxyConnectTimeoutSecondsOk returns a tuple with the NetworkIngressProxyConnectTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyConnectTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressProxyConnectTimeoutSeconds(v int32)`

SetNetworkIngressProxyConnectTimeoutSeconds sets NetworkIngressProxyConnectTimeoutSeconds field to given value.

### HasNetworkIngressProxyConnectTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressProxyConnectTimeoutSeconds() bool`

HasNetworkIngressProxyConnectTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxySendTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressProxySendTimeoutSeconds() int32`

GetNetworkIngressProxySendTimeoutSeconds returns the NetworkIngressProxySendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxySendTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxySendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxySendTimeoutSecondsOk returns a tuple with the NetworkIngressProxySendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxySendTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressProxySendTimeoutSeconds(v int32)`

SetNetworkIngressProxySendTimeoutSeconds sets NetworkIngressProxySendTimeoutSeconds field to given value.

### HasNetworkIngressProxySendTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressProxySendTimeoutSeconds() bool`

HasNetworkIngressProxySendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyReadTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyReadTimeoutSeconds() int32`

GetNetworkIngressProxyReadTimeoutSeconds returns the NetworkIngressProxyReadTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressProxyReadTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyReadTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressProxyReadTimeoutSecondsOk returns a tuple with the NetworkIngressProxyReadTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyReadTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressProxyReadTimeoutSeconds(v int32)`

SetNetworkIngressProxyReadTimeoutSeconds sets NetworkIngressProxyReadTimeoutSeconds field to given value.

### HasNetworkIngressProxyReadTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressProxyReadTimeoutSeconds() bool`

HasNetworkIngressProxyReadTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressProxyBuffering

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyBuffering() string`

GetNetworkIngressProxyBuffering returns the NetworkIngressProxyBuffering field if non-nil, zero value otherwise.

### GetNetworkIngressProxyBufferingOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyBufferingOk() (*string, bool)`

GetNetworkIngressProxyBufferingOk returns a tuple with the NetworkIngressProxyBuffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyBuffering

`func (o *HelmAdvancedSettings) SetNetworkIngressProxyBuffering(v string)`

SetNetworkIngressProxyBuffering sets NetworkIngressProxyBuffering field to given value.

### HasNetworkIngressProxyBuffering

`func (o *HelmAdvancedSettings) HasNetworkIngressProxyBuffering() bool`

HasNetworkIngressProxyBuffering returns a boolean if a field has been set.

### GetNetworkIngressProxyRequestBuffering

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyRequestBuffering() string`

GetNetworkIngressProxyRequestBuffering returns the NetworkIngressProxyRequestBuffering field if non-nil, zero value otherwise.

### GetNetworkIngressProxyRequestBufferingOk

`func (o *HelmAdvancedSettings) GetNetworkIngressProxyRequestBufferingOk() (*string, bool)`

GetNetworkIngressProxyRequestBufferingOk returns a tuple with the NetworkIngressProxyRequestBuffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressProxyRequestBuffering

`func (o *HelmAdvancedSettings) SetNetworkIngressProxyRequestBuffering(v string)`

SetNetworkIngressProxyRequestBuffering sets NetworkIngressProxyRequestBuffering field to given value.

### HasNetworkIngressProxyRequestBuffering

`func (o *HelmAdvancedSettings) HasNetworkIngressProxyRequestBuffering() bool`

HasNetworkIngressProxyRequestBuffering returns a boolean if a field has been set.

### GetNetworkIngressGrpcSendTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressGrpcSendTimeoutSeconds() int32`

GetNetworkIngressGrpcSendTimeoutSeconds returns the NetworkIngressGrpcSendTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressGrpcSendTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressGrpcSendTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressGrpcSendTimeoutSecondsOk returns a tuple with the NetworkIngressGrpcSendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressGrpcSendTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressGrpcSendTimeoutSeconds(v int32)`

SetNetworkIngressGrpcSendTimeoutSeconds sets NetworkIngressGrpcSendTimeoutSeconds field to given value.

### HasNetworkIngressGrpcSendTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressGrpcSendTimeoutSeconds() bool`

HasNetworkIngressGrpcSendTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressGrpcReadTimeoutSeconds

`func (o *HelmAdvancedSettings) GetNetworkIngressGrpcReadTimeoutSeconds() int32`

GetNetworkIngressGrpcReadTimeoutSeconds returns the NetworkIngressGrpcReadTimeoutSeconds field if non-nil, zero value otherwise.

### GetNetworkIngressGrpcReadTimeoutSecondsOk

`func (o *HelmAdvancedSettings) GetNetworkIngressGrpcReadTimeoutSecondsOk() (*int32, bool)`

GetNetworkIngressGrpcReadTimeoutSecondsOk returns a tuple with the NetworkIngressGrpcReadTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressGrpcReadTimeoutSeconds

`func (o *HelmAdvancedSettings) SetNetworkIngressGrpcReadTimeoutSeconds(v int32)`

SetNetworkIngressGrpcReadTimeoutSeconds sets NetworkIngressGrpcReadTimeoutSeconds field to given value.

### HasNetworkIngressGrpcReadTimeoutSeconds

`func (o *HelmAdvancedSettings) HasNetworkIngressGrpcReadTimeoutSeconds() bool`

HasNetworkIngressGrpcReadTimeoutSeconds returns a boolean if a field has been set.

### GetNetworkIngressWhitelistSourceRange

`func (o *HelmAdvancedSettings) GetNetworkIngressWhitelistSourceRange() string`

GetNetworkIngressWhitelistSourceRange returns the NetworkIngressWhitelistSourceRange field if non-nil, zero value otherwise.

### GetNetworkIngressWhitelistSourceRangeOk

`func (o *HelmAdvancedSettings) GetNetworkIngressWhitelistSourceRangeOk() (*string, bool)`

GetNetworkIngressWhitelistSourceRangeOk returns a tuple with the NetworkIngressWhitelistSourceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressWhitelistSourceRange

`func (o *HelmAdvancedSettings) SetNetworkIngressWhitelistSourceRange(v string)`

SetNetworkIngressWhitelistSourceRange sets NetworkIngressWhitelistSourceRange field to given value.

### HasNetworkIngressWhitelistSourceRange

`func (o *HelmAdvancedSettings) HasNetworkIngressWhitelistSourceRange() bool`

HasNetworkIngressWhitelistSourceRange returns a boolean if a field has been set.

### GetNetworkIngressDenylistSourceRange

`func (o *HelmAdvancedSettings) GetNetworkIngressDenylistSourceRange() string`

GetNetworkIngressDenylistSourceRange returns the NetworkIngressDenylistSourceRange field if non-nil, zero value otherwise.

### GetNetworkIngressDenylistSourceRangeOk

`func (o *HelmAdvancedSettings) GetNetworkIngressDenylistSourceRangeOk() (*string, bool)`

GetNetworkIngressDenylistSourceRangeOk returns a tuple with the NetworkIngressDenylistSourceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressDenylistSourceRange

`func (o *HelmAdvancedSettings) SetNetworkIngressDenylistSourceRange(v string)`

SetNetworkIngressDenylistSourceRange sets NetworkIngressDenylistSourceRange field to given value.

### HasNetworkIngressDenylistSourceRange

`func (o *HelmAdvancedSettings) HasNetworkIngressDenylistSourceRange() bool`

HasNetworkIngressDenylistSourceRange returns a boolean if a field has been set.

### GetNetworkIngressExtraHeaders

`func (o *HelmAdvancedSettings) GetNetworkIngressExtraHeaders() string`

GetNetworkIngressExtraHeaders returns the NetworkIngressExtraHeaders field if non-nil, zero value otherwise.

### GetNetworkIngressExtraHeadersOk

`func (o *HelmAdvancedSettings) GetNetworkIngressExtraHeadersOk() (*string, bool)`

GetNetworkIngressExtraHeadersOk returns a tuple with the NetworkIngressExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressExtraHeaders

`func (o *HelmAdvancedSettings) SetNetworkIngressExtraHeaders(v string)`

SetNetworkIngressExtraHeaders sets NetworkIngressExtraHeaders field to given value.

### HasNetworkIngressExtraHeaders

`func (o *HelmAdvancedSettings) HasNetworkIngressExtraHeaders() bool`

HasNetworkIngressExtraHeaders returns a boolean if a field has been set.

### GetNetworkIngressBasicAuthEnvVar

`func (o *HelmAdvancedSettings) GetNetworkIngressBasicAuthEnvVar() string`

GetNetworkIngressBasicAuthEnvVar returns the NetworkIngressBasicAuthEnvVar field if non-nil, zero value otherwise.

### GetNetworkIngressBasicAuthEnvVarOk

`func (o *HelmAdvancedSettings) GetNetworkIngressBasicAuthEnvVarOk() (*string, bool)`

GetNetworkIngressBasicAuthEnvVarOk returns a tuple with the NetworkIngressBasicAuthEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressBasicAuthEnvVar

`func (o *HelmAdvancedSettings) SetNetworkIngressBasicAuthEnvVar(v string)`

SetNetworkIngressBasicAuthEnvVar sets NetworkIngressBasicAuthEnvVar field to given value.

### HasNetworkIngressBasicAuthEnvVar

`func (o *HelmAdvancedSettings) HasNetworkIngressBasicAuthEnvVar() bool`

HasNetworkIngressBasicAuthEnvVar returns a boolean if a field has been set.

### GetNetworkIngressEnableStickySession

`func (o *HelmAdvancedSettings) GetNetworkIngressEnableStickySession() bool`

GetNetworkIngressEnableStickySession returns the NetworkIngressEnableStickySession field if non-nil, zero value otherwise.

### GetNetworkIngressEnableStickySessionOk

`func (o *HelmAdvancedSettings) GetNetworkIngressEnableStickySessionOk() (*bool, bool)`

GetNetworkIngressEnableStickySessionOk returns a tuple with the NetworkIngressEnableStickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngressEnableStickySession

`func (o *HelmAdvancedSettings) SetNetworkIngressEnableStickySession(v bool)`

SetNetworkIngressEnableStickySession sets NetworkIngressEnableStickySession field to given value.

### HasNetworkIngressEnableStickySession

`func (o *HelmAdvancedSettings) HasNetworkIngressEnableStickySession() bool`

HasNetworkIngressEnableStickySession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


