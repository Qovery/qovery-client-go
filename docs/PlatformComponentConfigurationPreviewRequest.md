# PlatformComponentConfigurationPreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileConfig** | Pointer to **map[string]interface{}** | Configuration values keyed by their catalog field name | [optional] 
**ClusterInputs** | Pointer to **map[string]string** |  | [optional] 
**ComponentOutputs** | Pointer to **map[string]map[string]string** | String values keyed first by component key and then by input key | [optional] 

## Methods

### NewPlatformComponentConfigurationPreviewRequest

`func NewPlatformComponentConfigurationPreviewRequest() *PlatformComponentConfigurationPreviewRequest`

NewPlatformComponentConfigurationPreviewRequest instantiates a new PlatformComponentConfigurationPreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentConfigurationPreviewRequestWithDefaults

`func NewPlatformComponentConfigurationPreviewRequestWithDefaults() *PlatformComponentConfigurationPreviewRequest`

NewPlatformComponentConfigurationPreviewRequestWithDefaults instantiates a new PlatformComponentConfigurationPreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileConfig

`func (o *PlatformComponentConfigurationPreviewRequest) GetProfileConfig() map[string]interface{}`

GetProfileConfig returns the ProfileConfig field if non-nil, zero value otherwise.

### GetProfileConfigOk

`func (o *PlatformComponentConfigurationPreviewRequest) GetProfileConfigOk() (*map[string]interface{}, bool)`

GetProfileConfigOk returns a tuple with the ProfileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileConfig

`func (o *PlatformComponentConfigurationPreviewRequest) SetProfileConfig(v map[string]interface{})`

SetProfileConfig sets ProfileConfig field to given value.

### HasProfileConfig

`func (o *PlatformComponentConfigurationPreviewRequest) HasProfileConfig() bool`

HasProfileConfig returns a boolean if a field has been set.

### GetClusterInputs

`func (o *PlatformComponentConfigurationPreviewRequest) GetClusterInputs() map[string]string`

GetClusterInputs returns the ClusterInputs field if non-nil, zero value otherwise.

### GetClusterInputsOk

`func (o *PlatformComponentConfigurationPreviewRequest) GetClusterInputsOk() (*map[string]string, bool)`

GetClusterInputsOk returns a tuple with the ClusterInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInputs

`func (o *PlatformComponentConfigurationPreviewRequest) SetClusterInputs(v map[string]string)`

SetClusterInputs sets ClusterInputs field to given value.

### HasClusterInputs

`func (o *PlatformComponentConfigurationPreviewRequest) HasClusterInputs() bool`

HasClusterInputs returns a boolean if a field has been set.

### GetComponentOutputs

`func (o *PlatformComponentConfigurationPreviewRequest) GetComponentOutputs() map[string]map[string]string`

GetComponentOutputs returns the ComponentOutputs field if non-nil, zero value otherwise.

### GetComponentOutputsOk

`func (o *PlatformComponentConfigurationPreviewRequest) GetComponentOutputsOk() (*map[string]map[string]string, bool)`

GetComponentOutputsOk returns a tuple with the ComponentOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentOutputs

`func (o *PlatformComponentConfigurationPreviewRequest) SetComponentOutputs(v map[string]map[string]string)`

SetComponentOutputs sets ComponentOutputs field to given value.

### HasComponentOutputs

`func (o *PlatformComponentConfigurationPreviewRequest) HasComponentOutputs() bool`

HasComponentOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


