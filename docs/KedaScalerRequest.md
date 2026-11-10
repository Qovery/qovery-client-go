# KedaScalerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScalerType** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Role** | [**KedaScalerRole**](KedaScalerRole.md) |  | 
**ConfigJson** | Pointer to **map[string]interface{}** | Scaler &#x60;metadata&#x60; only, as an object of string values, e.g. &#x60;{\&quot;query\&quot;: \&quot;...\&quot;, \&quot;threshold\&quot;: \&quot;10\&quot;}&#x60;. Other trigger fields (&#x60;name&#x60;, &#x60;metricType&#x60;, &#x60;useCachedMetrics&#x60;) cannot be set here: use &#x60;config_yaml&#x60;. Exactly one of &#x60;config_json&#x60; and &#x60;config_yaml&#x60; must be set. | [optional] 
**ConfigYaml** | Pointer to **string** | Body of the KEDA trigger as YAML, without &#x60;type&#x60; (set from &#x60;scaler_type&#x60;) and &#x60;authenticationRef&#x60; (set from &#x60;trigger_authentication&#x60;). Scaler settings go under &#x60;metadata:&#x60;; &#x60;name&#x60;, &#x60;metricType&#x60; and &#x60;useCachedMetrics&#x60; can be set next to it. Exactly one of &#x60;config_json&#x60; and &#x60;config_yaml&#x60; must be set. | [optional] 
**TriggerAuthentication** | Pointer to [**KedaTriggerAuthenticationRequest**](KedaTriggerAuthenticationRequest.md) |  | [optional] 

## Methods

### NewKedaScalerRequest

`func NewKedaScalerRequest(scalerType string, role KedaScalerRole, ) *KedaScalerRequest`

NewKedaScalerRequest instantiates a new KedaScalerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKedaScalerRequestWithDefaults

`func NewKedaScalerRequestWithDefaults() *KedaScalerRequest`

NewKedaScalerRequestWithDefaults instantiates a new KedaScalerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScalerType

`func (o *KedaScalerRequest) GetScalerType() string`

GetScalerType returns the ScalerType field if non-nil, zero value otherwise.

### GetScalerTypeOk

`func (o *KedaScalerRequest) GetScalerTypeOk() (*string, bool)`

GetScalerTypeOk returns a tuple with the ScalerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalerType

`func (o *KedaScalerRequest) SetScalerType(v string)`

SetScalerType sets ScalerType field to given value.


### GetEnabled

`func (o *KedaScalerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KedaScalerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KedaScalerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KedaScalerRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRole

`func (o *KedaScalerRequest) GetRole() KedaScalerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *KedaScalerRequest) GetRoleOk() (*KedaScalerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *KedaScalerRequest) SetRole(v KedaScalerRole)`

SetRole sets Role field to given value.


### GetConfigJson

`func (o *KedaScalerRequest) GetConfigJson() map[string]interface{}`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *KedaScalerRequest) GetConfigJsonOk() (*map[string]interface{}, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *KedaScalerRequest) SetConfigJson(v map[string]interface{})`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *KedaScalerRequest) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### GetConfigYaml

`func (o *KedaScalerRequest) GetConfigYaml() string`

GetConfigYaml returns the ConfigYaml field if non-nil, zero value otherwise.

### GetConfigYamlOk

`func (o *KedaScalerRequest) GetConfigYamlOk() (*string, bool)`

GetConfigYamlOk returns a tuple with the ConfigYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigYaml

`func (o *KedaScalerRequest) SetConfigYaml(v string)`

SetConfigYaml sets ConfigYaml field to given value.

### HasConfigYaml

`func (o *KedaScalerRequest) HasConfigYaml() bool`

HasConfigYaml returns a boolean if a field has been set.

### GetTriggerAuthentication

`func (o *KedaScalerRequest) GetTriggerAuthentication() KedaTriggerAuthenticationRequest`

GetTriggerAuthentication returns the TriggerAuthentication field if non-nil, zero value otherwise.

### GetTriggerAuthenticationOk

`func (o *KedaScalerRequest) GetTriggerAuthenticationOk() (*KedaTriggerAuthenticationRequest, bool)`

GetTriggerAuthenticationOk returns a tuple with the TriggerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAuthentication

`func (o *KedaScalerRequest) SetTriggerAuthentication(v KedaTriggerAuthenticationRequest)`

SetTriggerAuthentication sets TriggerAuthentication field to given value.

### HasTriggerAuthentication

`func (o *KedaScalerRequest) HasTriggerAuthentication() bool`

HasTriggerAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


