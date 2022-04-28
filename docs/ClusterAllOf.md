# ClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedCloudProviderCost** | Pointer to **int32** | This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**Features** | Pointer to [**[]ClusterFeature**](ClusterFeature.md) |  | [optional] 
**HasAccess** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterAllOf

`func NewClusterAllOf() *ClusterAllOf`

NewClusterAllOf instantiates a new ClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAllOfWithDefaults

`func NewClusterAllOfWithDefaults() *ClusterAllOf`

NewClusterAllOfWithDefaults instantiates a new ClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedCloudProviderCost

`func (o *ClusterAllOf) GetEstimatedCloudProviderCost() int32`

GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field if non-nil, zero value otherwise.

### GetEstimatedCloudProviderCostOk

`func (o *ClusterAllOf) GetEstimatedCloudProviderCostOk() (*int32, bool)`

GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCloudProviderCost

`func (o *ClusterAllOf) SetEstimatedCloudProviderCost(v int32)`

SetEstimatedCloudProviderCost sets EstimatedCloudProviderCost field to given value.

### HasEstimatedCloudProviderCost

`func (o *ClusterAllOf) HasEstimatedCloudProviderCost() bool`

HasEstimatedCloudProviderCost returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterAllOf) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterAllOf) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterAllOf) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFeatures

`func (o *ClusterAllOf) GetFeatures() []ClusterFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterAllOf) GetFeaturesOk() (*[]ClusterFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterAllOf) SetFeatures(v []ClusterFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterAllOf) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHasAccess

`func (o *ClusterAllOf) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *ClusterAllOf) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *ClusterAllOf) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *ClusterAllOf) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsDefault

`func (o *ClusterAllOf) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ClusterAllOf) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ClusterAllOf) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ClusterAllOf) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


