# MetricsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Configuration** | Pointer to [**MetricsParametersConfiguration**](MetricsParametersConfiguration.md) |  | [optional] 

## Methods

### NewMetricsParameters

`func NewMetricsParameters() *MetricsParameters`

NewMetricsParameters instantiates a new MetricsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsParametersWithDefaults

`func NewMetricsParametersWithDefaults() *MetricsParameters`

NewMetricsParametersWithDefaults instantiates a new MetricsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MetricsParameters) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsParameters) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsParameters) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsParameters) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfiguration

`func (o *MetricsParameters) GetConfiguration() MetricsParametersConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MetricsParameters) GetConfigurationOk() (*MetricsParametersConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MetricsParameters) SetConfiguration(v MetricsParametersConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MetricsParameters) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


