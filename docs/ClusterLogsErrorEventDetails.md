# ClusterLogsErrorEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderKind** | Pointer to **string** | cloud provider used | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Transmitter** | Pointer to [**ClusterLogsErrorEventDetailsTransmitter**](ClusterLogsErrorEventDetailsTransmitter.md) |  | [optional] 

## Methods

### NewClusterLogsErrorEventDetails

`func NewClusterLogsErrorEventDetails() *ClusterLogsErrorEventDetails`

NewClusterLogsErrorEventDetails instantiates a new ClusterLogsErrorEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsErrorEventDetailsWithDefaults

`func NewClusterLogsErrorEventDetailsWithDefaults() *ClusterLogsErrorEventDetails`

NewClusterLogsErrorEventDetailsWithDefaults instantiates a new ClusterLogsErrorEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderKind

`func (o *ClusterLogsErrorEventDetails) GetProviderKind() string`

GetProviderKind returns the ProviderKind field if non-nil, zero value otherwise.

### GetProviderKindOk

`func (o *ClusterLogsErrorEventDetails) GetProviderKindOk() (*string, bool)`

GetProviderKindOk returns a tuple with the ProviderKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKind

`func (o *ClusterLogsErrorEventDetails) SetProviderKind(v string)`

SetProviderKind sets ProviderKind field to given value.

### HasProviderKind

`func (o *ClusterLogsErrorEventDetails) HasProviderKind() bool`

HasProviderKind returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterLogsErrorEventDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterLogsErrorEventDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterLogsErrorEventDetails) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterLogsErrorEventDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTransmitter

`func (o *ClusterLogsErrorEventDetails) GetTransmitter() ClusterLogsErrorEventDetailsTransmitter`

GetTransmitter returns the Transmitter field if non-nil, zero value otherwise.

### GetTransmitterOk

`func (o *ClusterLogsErrorEventDetails) GetTransmitterOk() (*ClusterLogsErrorEventDetailsTransmitter, bool)`

GetTransmitterOk returns a tuple with the Transmitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitter

`func (o *ClusterLogsErrorEventDetails) SetTransmitter(v ClusterLogsErrorEventDetailsTransmitter)`

SetTransmitter sets Transmitter field to given value.

### HasTransmitter

`func (o *ClusterLogsErrorEventDetails) HasTransmitter() bool`

HasTransmitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


