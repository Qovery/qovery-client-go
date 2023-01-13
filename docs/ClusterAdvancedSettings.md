# ClusterAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LokiLogRetentionInWeek** | Pointer to **int32** | For how long in week loki is going to keep logs of your applications | [optional] [default to 12]
**AwsVpcEnableS3FlowLogs** | Pointer to **bool** | Enable flow logs for on the VPC and store them in an S3 bucket | [optional] [default to false]
**RegistryImageRetentionTime** | Pointer to **int32** | Configure the number of seconds before cleaning images in the registry | [optional] [default to 31536000]
**CloudProviderContainerRegistryTags** | Pointer to [**ClusterAdvancedSettingsCloudProviderContainerRegistryTags**](ClusterAdvancedSettingsCloudProviderContainerRegistryTags.md) |  | [optional] 
**LoadBalancerSize** | Pointer to **string** | Select the size of the main load_balancer (only effective for Scaleway) | [optional] [default to "lb-s"]
**PlecoResourcesTtl** | Pointer to **int32** |  | [optional] [default to -1]

## Methods

### NewClusterAdvancedSettings

`func NewClusterAdvancedSettings() *ClusterAdvancedSettings`

NewClusterAdvancedSettings instantiates a new ClusterAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAdvancedSettingsWithDefaults

`func NewClusterAdvancedSettingsWithDefaults() *ClusterAdvancedSettings`

NewClusterAdvancedSettingsWithDefaults instantiates a new ClusterAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLokiLogRetentionInWeek

`func (o *ClusterAdvancedSettings) GetLokiLogRetentionInWeek() int32`

GetLokiLogRetentionInWeek returns the LokiLogRetentionInWeek field if non-nil, zero value otherwise.

### GetLokiLogRetentionInWeekOk

`func (o *ClusterAdvancedSettings) GetLokiLogRetentionInWeekOk() (*int32, bool)`

GetLokiLogRetentionInWeekOk returns a tuple with the LokiLogRetentionInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLokiLogRetentionInWeek

`func (o *ClusterAdvancedSettings) SetLokiLogRetentionInWeek(v int32)`

SetLokiLogRetentionInWeek sets LokiLogRetentionInWeek field to given value.

### HasLokiLogRetentionInWeek

`func (o *ClusterAdvancedSettings) HasLokiLogRetentionInWeek() bool`

HasLokiLogRetentionInWeek returns a boolean if a field has been set.

### GetAwsVpcEnableS3FlowLogs

`func (o *ClusterAdvancedSettings) GetAwsVpcEnableS3FlowLogs() bool`

GetAwsVpcEnableS3FlowLogs returns the AwsVpcEnableS3FlowLogs field if non-nil, zero value otherwise.

### GetAwsVpcEnableS3FlowLogsOk

`func (o *ClusterAdvancedSettings) GetAwsVpcEnableS3FlowLogsOk() (*bool, bool)`

GetAwsVpcEnableS3FlowLogsOk returns a tuple with the AwsVpcEnableS3FlowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpcEnableS3FlowLogs

`func (o *ClusterAdvancedSettings) SetAwsVpcEnableS3FlowLogs(v bool)`

SetAwsVpcEnableS3FlowLogs sets AwsVpcEnableS3FlowLogs field to given value.

### HasAwsVpcEnableS3FlowLogs

`func (o *ClusterAdvancedSettings) HasAwsVpcEnableS3FlowLogs() bool`

HasAwsVpcEnableS3FlowLogs returns a boolean if a field has been set.

### GetRegistryImageRetentionTime

`func (o *ClusterAdvancedSettings) GetRegistryImageRetentionTime() int32`

GetRegistryImageRetentionTime returns the RegistryImageRetentionTime field if non-nil, zero value otherwise.

### GetRegistryImageRetentionTimeOk

`func (o *ClusterAdvancedSettings) GetRegistryImageRetentionTimeOk() (*int32, bool)`

GetRegistryImageRetentionTimeOk returns a tuple with the RegistryImageRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryImageRetentionTime

`func (o *ClusterAdvancedSettings) SetRegistryImageRetentionTime(v int32)`

SetRegistryImageRetentionTime sets RegistryImageRetentionTime field to given value.

### HasRegistryImageRetentionTime

`func (o *ClusterAdvancedSettings) HasRegistryImageRetentionTime() bool`

HasRegistryImageRetentionTime returns a boolean if a field has been set.

### GetCloudProviderContainerRegistryTags

`func (o *ClusterAdvancedSettings) GetCloudProviderContainerRegistryTags() ClusterAdvancedSettingsCloudProviderContainerRegistryTags`

GetCloudProviderContainerRegistryTags returns the CloudProviderContainerRegistryTags field if non-nil, zero value otherwise.

### GetCloudProviderContainerRegistryTagsOk

`func (o *ClusterAdvancedSettings) GetCloudProviderContainerRegistryTagsOk() (*ClusterAdvancedSettingsCloudProviderContainerRegistryTags, bool)`

GetCloudProviderContainerRegistryTagsOk returns a tuple with the CloudProviderContainerRegistryTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderContainerRegistryTags

`func (o *ClusterAdvancedSettings) SetCloudProviderContainerRegistryTags(v ClusterAdvancedSettingsCloudProviderContainerRegistryTags)`

SetCloudProviderContainerRegistryTags sets CloudProviderContainerRegistryTags field to given value.

### HasCloudProviderContainerRegistryTags

`func (o *ClusterAdvancedSettings) HasCloudProviderContainerRegistryTags() bool`

HasCloudProviderContainerRegistryTags returns a boolean if a field has been set.

### GetLoadBalancerSize

`func (o *ClusterAdvancedSettings) GetLoadBalancerSize() string`

GetLoadBalancerSize returns the LoadBalancerSize field if non-nil, zero value otherwise.

### GetLoadBalancerSizeOk

`func (o *ClusterAdvancedSettings) GetLoadBalancerSizeOk() (*string, bool)`

GetLoadBalancerSizeOk returns a tuple with the LoadBalancerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSize

`func (o *ClusterAdvancedSettings) SetLoadBalancerSize(v string)`

SetLoadBalancerSize sets LoadBalancerSize field to given value.

### HasLoadBalancerSize

`func (o *ClusterAdvancedSettings) HasLoadBalancerSize() bool`

HasLoadBalancerSize returns a boolean if a field has been set.

### GetPlecoResourcesTtl

`func (o *ClusterAdvancedSettings) GetPlecoResourcesTtl() int32`

GetPlecoResourcesTtl returns the PlecoResourcesTtl field if non-nil, zero value otherwise.

### GetPlecoResourcesTtlOk

`func (o *ClusterAdvancedSettings) GetPlecoResourcesTtlOk() (*int32, bool)`

GetPlecoResourcesTtlOk returns a tuple with the PlecoResourcesTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlecoResourcesTtl

`func (o *ClusterAdvancedSettings) SetPlecoResourcesTtl(v int32)`

SetPlecoResourcesTtl sets PlecoResourcesTtl field to given value.

### HasPlecoResourcesTtl

`func (o *ClusterAdvancedSettings) HasPlecoResourcesTtl() bool`

HasPlecoResourcesTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


