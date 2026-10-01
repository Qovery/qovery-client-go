# SelfManagedClusterPlatformSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateKey** | Pointer to **string** | Key of the platform template release, set together with templateVersion. | [optional] 
**TemplateVersion** | Pointer to **string** | Version of the platform template release, set together with templateKey. | [optional] 
**LayerSelections** | Pointer to **map[string]bool** |  | [optional] 
**ManagedConfig** | Pointer to **map[string]map[string]interface{}** | Component configuration values keyed by component key | [optional] 

## Methods

### NewSelfManagedClusterPlatformSelection

`func NewSelfManagedClusterPlatformSelection() *SelfManagedClusterPlatformSelection`

NewSelfManagedClusterPlatformSelection instantiates a new SelfManagedClusterPlatformSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfManagedClusterPlatformSelectionWithDefaults

`func NewSelfManagedClusterPlatformSelectionWithDefaults() *SelfManagedClusterPlatformSelection`

NewSelfManagedClusterPlatformSelectionWithDefaults instantiates a new SelfManagedClusterPlatformSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateKey

`func (o *SelfManagedClusterPlatformSelection) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *SelfManagedClusterPlatformSelection) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *SelfManagedClusterPlatformSelection) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *SelfManagedClusterPlatformSelection) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *SelfManagedClusterPlatformSelection) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *SelfManagedClusterPlatformSelection) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *SelfManagedClusterPlatformSelection) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *SelfManagedClusterPlatformSelection) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetLayerSelections

`func (o *SelfManagedClusterPlatformSelection) GetLayerSelections() map[string]bool`

GetLayerSelections returns the LayerSelections field if non-nil, zero value otherwise.

### GetLayerSelectionsOk

`func (o *SelfManagedClusterPlatformSelection) GetLayerSelectionsOk() (*map[string]bool, bool)`

GetLayerSelectionsOk returns a tuple with the LayerSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerSelections

`func (o *SelfManagedClusterPlatformSelection) SetLayerSelections(v map[string]bool)`

SetLayerSelections sets LayerSelections field to given value.

### HasLayerSelections

`func (o *SelfManagedClusterPlatformSelection) HasLayerSelections() bool`

HasLayerSelections returns a boolean if a field has been set.

### GetManagedConfig

`func (o *SelfManagedClusterPlatformSelection) GetManagedConfig() map[string]map[string]interface{}`

GetManagedConfig returns the ManagedConfig field if non-nil, zero value otherwise.

### GetManagedConfigOk

`func (o *SelfManagedClusterPlatformSelection) GetManagedConfigOk() (*map[string]map[string]interface{}, bool)`

GetManagedConfigOk returns a tuple with the ManagedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedConfig

`func (o *SelfManagedClusterPlatformSelection) SetManagedConfig(v map[string]map[string]interface{})`

SetManagedConfig sets ManagedConfig field to given value.

### HasManagedConfig

`func (o *SelfManagedClusterPlatformSelection) HasManagedConfig() bool`

HasManagedConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


