# MetricsConfigurationManagedByQovery

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

### NewMetricsConfigurationManagedByQovery

`func NewMetricsConfigurationManagedByQovery() *MetricsConfigurationManagedByQovery`

NewMetricsConfigurationManagedByQovery instantiates a new MetricsConfigurationManagedByQovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsConfigurationManagedByQoveryWithDefaults

`func NewMetricsConfigurationManagedByQoveryWithDefaults() *MetricsConfigurationManagedByQovery`

NewMetricsConfigurationManagedByQoveryWithDefaults instantiates a new MetricsConfigurationManagedByQovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *MetricsConfigurationManagedByQovery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricsConfigurationManagedByQovery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricsConfigurationManagedByQovery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetricsConfigurationManagedByQovery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetResourceProfile

`func (o *MetricsConfigurationManagedByQovery) GetResourceProfile() ObservabilityResourceProfile`

GetResourceProfile returns the ResourceProfile field if non-nil, zero value otherwise.

### GetResourceProfileOk

`func (o *MetricsConfigurationManagedByQovery) GetResourceProfileOk() (*ObservabilityResourceProfile, bool)`

GetResourceProfileOk returns a tuple with the ResourceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProfile

`func (o *MetricsConfigurationManagedByQovery) SetResourceProfile(v ObservabilityResourceProfile)`

SetResourceProfile sets ResourceProfile field to given value.

### HasResourceProfile

`func (o *MetricsConfigurationManagedByQovery) HasResourceProfile() bool`

HasResourceProfile returns a boolean if a field has been set.

### GetCloudWatchExportConfig

`func (o *MetricsConfigurationManagedByQovery) GetCloudWatchExportConfig() CloudWatchExportConfig`

GetCloudWatchExportConfig returns the CloudWatchExportConfig field if non-nil, zero value otherwise.

### GetCloudWatchExportConfigOk

`func (o *MetricsConfigurationManagedByQovery) GetCloudWatchExportConfigOk() (*CloudWatchExportConfig, bool)`

GetCloudWatchExportConfigOk returns a tuple with the CloudWatchExportConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWatchExportConfig

`func (o *MetricsConfigurationManagedByQovery) SetCloudWatchExportConfig(v CloudWatchExportConfig)`

SetCloudWatchExportConfig sets CloudWatchExportConfig field to given value.

### HasCloudWatchExportConfig

`func (o *MetricsConfigurationManagedByQovery) HasCloudWatchExportConfig() bool`

HasCloudWatchExportConfig returns a boolean if a field has been set.

### GetHighAvailability

`func (o *MetricsConfigurationManagedByQovery) GetHighAvailability() bool`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *MetricsConfigurationManagedByQovery) GetHighAvailabilityOk() (*bool, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *MetricsConfigurationManagedByQovery) SetHighAvailability(v bool)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *MetricsConfigurationManagedByQovery) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### GetInternalNetworkMonitoring

`func (o *MetricsConfigurationManagedByQovery) GetInternalNetworkMonitoring() InternalNetworkMonitoring`

GetInternalNetworkMonitoring returns the InternalNetworkMonitoring field if non-nil, zero value otherwise.

### GetInternalNetworkMonitoringOk

`func (o *MetricsConfigurationManagedByQovery) GetInternalNetworkMonitoringOk() (*InternalNetworkMonitoring, bool)`

GetInternalNetworkMonitoringOk returns a tuple with the InternalNetworkMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNetworkMonitoring

`func (o *MetricsConfigurationManagedByQovery) SetInternalNetworkMonitoring(v InternalNetworkMonitoring)`

SetInternalNetworkMonitoring sets InternalNetworkMonitoring field to given value.

### HasInternalNetworkMonitoring

`func (o *MetricsConfigurationManagedByQovery) HasInternalNetworkMonitoring() bool`

HasInternalNetworkMonitoring returns a boolean if a field has been set.

### GetAlerting

`func (o *MetricsConfigurationManagedByQovery) GetAlerting() AlertingConfig`

GetAlerting returns the Alerting field if non-nil, zero value otherwise.

### GetAlertingOk

`func (o *MetricsConfigurationManagedByQovery) GetAlertingOk() (*AlertingConfig, bool)`

GetAlertingOk returns a tuple with the Alerting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerting

`func (o *MetricsConfigurationManagedByQovery) SetAlerting(v AlertingConfig)`

SetAlerting sets Alerting field to given value.

### HasAlerting

`func (o *MetricsConfigurationManagedByQovery) HasAlerting() bool`

HasAlerting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


