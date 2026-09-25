# PlatformSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateKey** | **string** |  | 
**TemplateVersion** | **string** |  | 
**LayerSelections** | Pointer to **map[string]bool** |  | [optional] 
**ManagedConfig** | Pointer to **map[string]map[string]interface{}** | Component configuration values keyed by component key | [optional] 

## Methods

### NewPlatformSelection

`func NewPlatformSelection(templateKey string, templateVersion string, ) *PlatformSelection`

NewPlatformSelection instantiates a new PlatformSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformSelectionWithDefaults

`func NewPlatformSelectionWithDefaults() *PlatformSelection`

NewPlatformSelectionWithDefaults instantiates a new PlatformSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateKey

`func (o *PlatformSelection) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *PlatformSelection) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *PlatformSelection) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.


### GetTemplateVersion

`func (o *PlatformSelection) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *PlatformSelection) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *PlatformSelection) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.


### GetLayerSelections

`func (o *PlatformSelection) GetLayerSelections() map[string]bool`

GetLayerSelections returns the LayerSelections field if non-nil, zero value otherwise.

### GetLayerSelectionsOk

`func (o *PlatformSelection) GetLayerSelectionsOk() (*map[string]bool, bool)`

GetLayerSelectionsOk returns a tuple with the LayerSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerSelections

`func (o *PlatformSelection) SetLayerSelections(v map[string]bool)`

SetLayerSelections sets LayerSelections field to given value.

### HasLayerSelections

`func (o *PlatformSelection) HasLayerSelections() bool`

HasLayerSelections returns a boolean if a field has been set.

### GetManagedConfig

`func (o *PlatformSelection) GetManagedConfig() map[string]map[string]interface{}`

GetManagedConfig returns the ManagedConfig field if non-nil, zero value otherwise.

### GetManagedConfigOk

`func (o *PlatformSelection) GetManagedConfigOk() (*map[string]map[string]interface{}, bool)`

GetManagedConfigOk returns a tuple with the ManagedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedConfig

`func (o *PlatformSelection) SetManagedConfig(v map[string]map[string]interface{})`

SetManagedConfig sets ManagedConfig field to given value.

### HasManagedConfig

`func (o *PlatformSelection) HasManagedConfig() bool`

HasManagedConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


