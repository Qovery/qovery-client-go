# ClusterPlatformConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | [**PlatformSelection**](PlatformSelection.md) |  | 
**ClusterInputs** | **map[string]map[string]string** | String values keyed first by component key and then by input key | 

## Methods

### NewClusterPlatformConfigurationRequest

`func NewClusterPlatformConfigurationRequest(platform PlatformSelection, clusterInputs map[string]map[string]string, ) *ClusterPlatformConfigurationRequest`

NewClusterPlatformConfigurationRequest instantiates a new ClusterPlatformConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPlatformConfigurationRequestWithDefaults

`func NewClusterPlatformConfigurationRequestWithDefaults() *ClusterPlatformConfigurationRequest`

NewClusterPlatformConfigurationRequestWithDefaults instantiates a new ClusterPlatformConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *ClusterPlatformConfigurationRequest) GetPlatform() PlatformSelection`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ClusterPlatformConfigurationRequest) GetPlatformOk() (*PlatformSelection, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ClusterPlatformConfigurationRequest) SetPlatform(v PlatformSelection)`

SetPlatform sets Platform field to given value.


### GetClusterInputs

`func (o *ClusterPlatformConfigurationRequest) GetClusterInputs() map[string]map[string]string`

GetClusterInputs returns the ClusterInputs field if non-nil, zero value otherwise.

### GetClusterInputsOk

`func (o *ClusterPlatformConfigurationRequest) GetClusterInputsOk() (*map[string]map[string]string, bool)`

GetClusterInputsOk returns a tuple with the ClusterInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInputs

`func (o *ClusterPlatformConfigurationRequest) SetClusterInputs(v map[string]map[string]string)`

SetClusterInputs sets ClusterInputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


