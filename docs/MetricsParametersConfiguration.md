# MetricsParametersConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ResourceProfile** | Pointer to [**ObservabilityResourceProfile**](ObservabilityResourceProfile.md) |  | [optional] 
**CloudWatchExportConfig** | Pointer to [**CloudWatchExportConfig**](CloudWatchExportConfig.md) |  | [optional] 
**HighAvailability** | Pointer to **bool** |  | [optional] 
**InternalNetworkMonitoring** | Pointer to [**InternalNetworkMonitoring**](InternalNetworkMonitoring.md) |  | [optional] 
**Alerting** | Pointer to [**AlertingConfig**](AlertingConfig.md) |  | [optional] 

## Methods

### NewMetricsParametersConfiguration

`func NewMetricsParametersConfiguration() *MetricsParametersConfiguration`

NewMetricsParametersConfiguration instantiates a new MetricsParametersConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsParametersConfigurationWithDefaults

`func NewMetricsParametersConfigurationWithDefaults() *MetricsParametersConfiguration`

NewMetricsParametersConfigurationWithDefaults instantiates a new MetricsParametersConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *MetricsParametersConfiguration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricsParametersConfiguration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricsParametersConfiguration) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetricsParametersConfiguration) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetResourceProfile

`func (o *MetricsParametersConfiguration) GetResourceProfile() ObservabilityResourceProfile`

GetResourceProfile returns the ResourceProfile field if non-nil, zero value otherwise.

### GetResourceProfileOk

`func (o *MetricsParametersConfiguration) GetResourceProfileOk() (*ObservabilityResourceProfile, bool)`

GetResourceProfileOk returns a tuple with the ResourceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProfile

`func (o *MetricsParametersConfiguration) SetResourceProfile(v ObservabilityResourceProfile)`

SetResourceProfile sets ResourceProfile field to given value.

### HasResourceProfile

`func (o *MetricsParametersConfiguration) HasResourceProfile() bool`

HasResourceProfile returns a boolean if a field has been set.

### GetCloudWatchExportConfig

`func (o *MetricsParametersConfiguration) GetCloudWatchExportConfig() CloudWatchExportConfig`

GetCloudWatchExportConfig returns the CloudWatchExportConfig field if non-nil, zero value otherwise.

### GetCloudWatchExportConfigOk

`func (o *MetricsParametersConfiguration) GetCloudWatchExportConfigOk() (*CloudWatchExportConfig, bool)`

GetCloudWatchExportConfigOk returns a tuple with the CloudWatchExportConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWatchExportConfig

`func (o *MetricsParametersConfiguration) SetCloudWatchExportConfig(v CloudWatchExportConfig)`

SetCloudWatchExportConfig sets CloudWatchExportConfig field to given value.

### HasCloudWatchExportConfig

`func (o *MetricsParametersConfiguration) HasCloudWatchExportConfig() bool`

HasCloudWatchExportConfig returns a boolean if a field has been set.

### GetHighAvailability

`func (o *MetricsParametersConfiguration) GetHighAvailability() bool`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *MetricsParametersConfiguration) GetHighAvailabilityOk() (*bool, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *MetricsParametersConfiguration) SetHighAvailability(v bool)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *MetricsParametersConfiguration) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### GetInternalNetworkMonitoring

`func (o *MetricsParametersConfiguration) GetInternalNetworkMonitoring() InternalNetworkMonitoring`

GetInternalNetworkMonitoring returns the InternalNetworkMonitoring field if non-nil, zero value otherwise.

### GetInternalNetworkMonitoringOk

`func (o *MetricsParametersConfiguration) GetInternalNetworkMonitoringOk() (*InternalNetworkMonitoring, bool)`

GetInternalNetworkMonitoringOk returns a tuple with the InternalNetworkMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNetworkMonitoring

`func (o *MetricsParametersConfiguration) SetInternalNetworkMonitoring(v InternalNetworkMonitoring)`

SetInternalNetworkMonitoring sets InternalNetworkMonitoring field to given value.

### HasInternalNetworkMonitoring

`func (o *MetricsParametersConfiguration) HasInternalNetworkMonitoring() bool`

HasInternalNetworkMonitoring returns a boolean if a field has been set.

### GetAlerting

`func (o *MetricsParametersConfiguration) GetAlerting() AlertingConfig`

GetAlerting returns the Alerting field if non-nil, zero value otherwise.

### GetAlertingOk

`func (o *MetricsParametersConfiguration) GetAlertingOk() (*AlertingConfig, bool)`

GetAlertingOk returns a tuple with the Alerting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerting

`func (o *MetricsParametersConfiguration) SetAlerting(v AlertingConfig)`

SetAlerting sets Alerting field to given value.

### HasAlerting

`func (o *MetricsParametersConfiguration) HasAlerting() bool`

HasAlerting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


