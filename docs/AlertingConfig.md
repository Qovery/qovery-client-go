# AlertingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**DefaultRuleLabels** | Pointer to **string** |  | [optional] 
**SpecConfigSecret** | Pointer to **string** |  | [optional] 
**SpecExternalUrl** | Pointer to **string** |  | [optional] 
**ConfigName** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertingConfig

`func NewAlertingConfig(enabled bool, ) *AlertingConfig`

NewAlertingConfig instantiates a new AlertingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingConfigWithDefaults

`func NewAlertingConfigWithDefaults() *AlertingConfig`

NewAlertingConfigWithDefaults instantiates a new AlertingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AlertingConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertingConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertingConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefaultRuleLabels

`func (o *AlertingConfig) GetDefaultRuleLabels() string`

GetDefaultRuleLabels returns the DefaultRuleLabels field if non-nil, zero value otherwise.

### GetDefaultRuleLabelsOk

`func (o *AlertingConfig) GetDefaultRuleLabelsOk() (*string, bool)`

GetDefaultRuleLabelsOk returns a tuple with the DefaultRuleLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRuleLabels

`func (o *AlertingConfig) SetDefaultRuleLabels(v string)`

SetDefaultRuleLabels sets DefaultRuleLabels field to given value.

### HasDefaultRuleLabels

`func (o *AlertingConfig) HasDefaultRuleLabels() bool`

HasDefaultRuleLabels returns a boolean if a field has been set.

### GetSpecConfigSecret

`func (o *AlertingConfig) GetSpecConfigSecret() string`

GetSpecConfigSecret returns the SpecConfigSecret field if non-nil, zero value otherwise.

### GetSpecConfigSecretOk

`func (o *AlertingConfig) GetSpecConfigSecretOk() (*string, bool)`

GetSpecConfigSecretOk returns a tuple with the SpecConfigSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecConfigSecret

`func (o *AlertingConfig) SetSpecConfigSecret(v string)`

SetSpecConfigSecret sets SpecConfigSecret field to given value.

### HasSpecConfigSecret

`func (o *AlertingConfig) HasSpecConfigSecret() bool`

HasSpecConfigSecret returns a boolean if a field has been set.

### GetSpecExternalUrl

`func (o *AlertingConfig) GetSpecExternalUrl() string`

GetSpecExternalUrl returns the SpecExternalUrl field if non-nil, zero value otherwise.

### GetSpecExternalUrlOk

`func (o *AlertingConfig) GetSpecExternalUrlOk() (*string, bool)`

GetSpecExternalUrlOk returns a tuple with the SpecExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecExternalUrl

`func (o *AlertingConfig) SetSpecExternalUrl(v string)`

SetSpecExternalUrl sets SpecExternalUrl field to given value.

### HasSpecExternalUrl

`func (o *AlertingConfig) HasSpecExternalUrl() bool`

HasSpecExternalUrl returns a boolean if a field has been set.

### GetConfigName

`func (o *AlertingConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *AlertingConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *AlertingConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *AlertingConfig) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


