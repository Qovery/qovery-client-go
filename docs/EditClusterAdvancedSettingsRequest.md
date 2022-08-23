# EditClusterAdvancedSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryImageRetentionTime** | Pointer to **int32** | Configure the number of seconds before cleaning images in the registry | [optional] [default to 31536000]
**LoadBalancerSize** | Pointer to **string** | Select the size of the main load_balancer (only effective for Scaleway) | [optional] [default to "lb-s"]
**PlecoResourcesTtl** | Pointer to **int32** |  | [optional] [default to -1]
**LokiLogRetentionInWeek** | Pointer to **int32** | For how long in week loki is going to keep logs of your applications | [optional] [default to 12]

## Methods

### NewEditClusterAdvancedSettingsRequest

`func NewEditClusterAdvancedSettingsRequest() *EditClusterAdvancedSettingsRequest`

NewEditClusterAdvancedSettingsRequest instantiates a new EditClusterAdvancedSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditClusterAdvancedSettingsRequestWithDefaults

`func NewEditClusterAdvancedSettingsRequestWithDefaults() *EditClusterAdvancedSettingsRequest`

NewEditClusterAdvancedSettingsRequestWithDefaults instantiates a new EditClusterAdvancedSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryImageRetentionTime

`func (o *EditClusterAdvancedSettingsRequest) GetRegistryImageRetentionTime() int32`

GetRegistryImageRetentionTime returns the RegistryImageRetentionTime field if non-nil, zero value otherwise.

### GetRegistryImageRetentionTimeOk

`func (o *EditClusterAdvancedSettingsRequest) GetRegistryImageRetentionTimeOk() (*int32, bool)`

GetRegistryImageRetentionTimeOk returns a tuple with the RegistryImageRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryImageRetentionTime

`func (o *EditClusterAdvancedSettingsRequest) SetRegistryImageRetentionTime(v int32)`

SetRegistryImageRetentionTime sets RegistryImageRetentionTime field to given value.

### HasRegistryImageRetentionTime

`func (o *EditClusterAdvancedSettingsRequest) HasRegistryImageRetentionTime() bool`

HasRegistryImageRetentionTime returns a boolean if a field has been set.

### GetLoadBalancerSize

`func (o *EditClusterAdvancedSettingsRequest) GetLoadBalancerSize() string`

GetLoadBalancerSize returns the LoadBalancerSize field if non-nil, zero value otherwise.

### GetLoadBalancerSizeOk

`func (o *EditClusterAdvancedSettingsRequest) GetLoadBalancerSizeOk() (*string, bool)`

GetLoadBalancerSizeOk returns a tuple with the LoadBalancerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSize

`func (o *EditClusterAdvancedSettingsRequest) SetLoadBalancerSize(v string)`

SetLoadBalancerSize sets LoadBalancerSize field to given value.

### HasLoadBalancerSize

`func (o *EditClusterAdvancedSettingsRequest) HasLoadBalancerSize() bool`

HasLoadBalancerSize returns a boolean if a field has been set.

### GetPlecoResourcesTtl

`func (o *EditClusterAdvancedSettingsRequest) GetPlecoResourcesTtl() int32`

GetPlecoResourcesTtl returns the PlecoResourcesTtl field if non-nil, zero value otherwise.

### GetPlecoResourcesTtlOk

`func (o *EditClusterAdvancedSettingsRequest) GetPlecoResourcesTtlOk() (*int32, bool)`

GetPlecoResourcesTtlOk returns a tuple with the PlecoResourcesTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlecoResourcesTtl

`func (o *EditClusterAdvancedSettingsRequest) SetPlecoResourcesTtl(v int32)`

SetPlecoResourcesTtl sets PlecoResourcesTtl field to given value.

### HasPlecoResourcesTtl

`func (o *EditClusterAdvancedSettingsRequest) HasPlecoResourcesTtl() bool`

HasPlecoResourcesTtl returns a boolean if a field has been set.

### GetLokiLogRetentionInWeek

`func (o *EditClusterAdvancedSettingsRequest) GetLokiLogRetentionInWeek() int32`

GetLokiLogRetentionInWeek returns the LokiLogRetentionInWeek field if non-nil, zero value otherwise.

### GetLokiLogRetentionInWeekOk

`func (o *EditClusterAdvancedSettingsRequest) GetLokiLogRetentionInWeekOk() (*int32, bool)`

GetLokiLogRetentionInWeekOk returns a tuple with the LokiLogRetentionInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLokiLogRetentionInWeek

`func (o *EditClusterAdvancedSettingsRequest) SetLokiLogRetentionInWeek(v int32)`

SetLokiLogRetentionInWeek sets LokiLogRetentionInWeek field to given value.

### HasLokiLogRetentionInWeek

`func (o *EditClusterAdvancedSettingsRequest) HasLokiLogRetentionInWeek() bool`

HasLokiLogRetentionInWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


