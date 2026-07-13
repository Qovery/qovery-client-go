# ClusterPlatformBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateKey** | **string** |  | 
**TemplateVersion** | **string** |  | 
**LayerSelections** | Pointer to **map[string]bool** |  | [optional] 
**ManagedConfig** | Pointer to **map[string]map[string]interface{}** | Component configuration values keyed by component key | [optional] 
**CustomerProvidedInputs** | Pointer to **map[string]map[string]string** | String values keyed first by component key and then by input key | [optional] 

## Methods

### NewClusterPlatformBindingRequest

`func NewClusterPlatformBindingRequest(templateKey string, templateVersion string, ) *ClusterPlatformBindingRequest`

NewClusterPlatformBindingRequest instantiates a new ClusterPlatformBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPlatformBindingRequestWithDefaults

`func NewClusterPlatformBindingRequestWithDefaults() *ClusterPlatformBindingRequest`

NewClusterPlatformBindingRequestWithDefaults instantiates a new ClusterPlatformBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateKey

`func (o *ClusterPlatformBindingRequest) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *ClusterPlatformBindingRequest) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *ClusterPlatformBindingRequest) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.


### GetTemplateVersion

`func (o *ClusterPlatformBindingRequest) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *ClusterPlatformBindingRequest) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *ClusterPlatformBindingRequest) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.


### GetLayerSelections

`func (o *ClusterPlatformBindingRequest) GetLayerSelections() map[string]bool`

GetLayerSelections returns the LayerSelections field if non-nil, zero value otherwise.

### GetLayerSelectionsOk

`func (o *ClusterPlatformBindingRequest) GetLayerSelectionsOk() (*map[string]bool, bool)`

GetLayerSelectionsOk returns a tuple with the LayerSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerSelections

`func (o *ClusterPlatformBindingRequest) SetLayerSelections(v map[string]bool)`

SetLayerSelections sets LayerSelections field to given value.

### HasLayerSelections

`func (o *ClusterPlatformBindingRequest) HasLayerSelections() bool`

HasLayerSelections returns a boolean if a field has been set.

### GetManagedConfig

`func (o *ClusterPlatformBindingRequest) GetManagedConfig() map[string]map[string]interface{}`

GetManagedConfig returns the ManagedConfig field if non-nil, zero value otherwise.

### GetManagedConfigOk

`func (o *ClusterPlatformBindingRequest) GetManagedConfigOk() (*map[string]map[string]interface{}, bool)`

GetManagedConfigOk returns a tuple with the ManagedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedConfig

`func (o *ClusterPlatformBindingRequest) SetManagedConfig(v map[string]map[string]interface{})`

SetManagedConfig sets ManagedConfig field to given value.

### HasManagedConfig

`func (o *ClusterPlatformBindingRequest) HasManagedConfig() bool`

HasManagedConfig returns a boolean if a field has been set.

### GetCustomerProvidedInputs

`func (o *ClusterPlatformBindingRequest) GetCustomerProvidedInputs() map[string]map[string]string`

GetCustomerProvidedInputs returns the CustomerProvidedInputs field if non-nil, zero value otherwise.

### GetCustomerProvidedInputsOk

`func (o *ClusterPlatformBindingRequest) GetCustomerProvidedInputsOk() (*map[string]map[string]string, bool)`

GetCustomerProvidedInputsOk returns a tuple with the CustomerProvidedInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProvidedInputs

`func (o *ClusterPlatformBindingRequest) SetCustomerProvidedInputs(v map[string]map[string]string)`

SetCustomerProvidedInputs sets CustomerProvidedInputs field to given value.

### HasCustomerProvidedInputs

`func (o *ClusterPlatformBindingRequest) HasCustomerProvidedInputs() bool`

HasCustomerProvidedInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


