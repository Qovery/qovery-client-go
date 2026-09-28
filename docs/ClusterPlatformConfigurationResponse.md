# ClusterPlatformConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**OrganizationId** | **string** |  | 
**Platform** | [**PlatformSelection**](PlatformSelection.md) |  | 
**ClusterInputs** | **map[string]map[string]string** | String values keyed first by component key and then by input key | 
**Layers** | [**[]ClusterPlatformBindingLayerResponse**](ClusterPlatformBindingLayerResponse.md) |  | 

## Methods

### NewClusterPlatformConfigurationResponse

`func NewClusterPlatformConfigurationResponse(clusterId string, organizationId string, platform PlatformSelection, clusterInputs map[string]map[string]string, layers []ClusterPlatformBindingLayerResponse, ) *ClusterPlatformConfigurationResponse`

NewClusterPlatformConfigurationResponse instantiates a new ClusterPlatformConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPlatformConfigurationResponseWithDefaults

`func NewClusterPlatformConfigurationResponseWithDefaults() *ClusterPlatformConfigurationResponse`

NewClusterPlatformConfigurationResponseWithDefaults instantiates a new ClusterPlatformConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterPlatformConfigurationResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterPlatformConfigurationResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterPlatformConfigurationResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetOrganizationId

`func (o *ClusterPlatformConfigurationResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterPlatformConfigurationResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterPlatformConfigurationResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetPlatform

`func (o *ClusterPlatformConfigurationResponse) GetPlatform() PlatformSelection`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ClusterPlatformConfigurationResponse) GetPlatformOk() (*PlatformSelection, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ClusterPlatformConfigurationResponse) SetPlatform(v PlatformSelection)`

SetPlatform sets Platform field to given value.


### GetClusterInputs

`func (o *ClusterPlatformConfigurationResponse) GetClusterInputs() map[string]map[string]string`

GetClusterInputs returns the ClusterInputs field if non-nil, zero value otherwise.

### GetClusterInputsOk

`func (o *ClusterPlatformConfigurationResponse) GetClusterInputsOk() (*map[string]map[string]string, bool)`

GetClusterInputsOk returns a tuple with the ClusterInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInputs

`func (o *ClusterPlatformConfigurationResponse) SetClusterInputs(v map[string]map[string]string)`

SetClusterInputs sets ClusterInputs field to given value.


### GetLayers

`func (o *ClusterPlatformConfigurationResponse) GetLayers() []ClusterPlatformBindingLayerResponse`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *ClusterPlatformConfigurationResponse) GetLayersOk() (*[]ClusterPlatformBindingLayerResponse, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *ClusterPlatformConfigurationResponse) SetLayers(v []ClusterPlatformBindingLayerResponse)`

SetLayers sets Layers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


