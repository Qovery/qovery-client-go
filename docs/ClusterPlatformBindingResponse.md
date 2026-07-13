# ClusterPlatformBindingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**OrganizationId** | **string** |  | 
**TemplateKey** | **string** |  | 
**TemplateVersion** | **string** |  | 
**LayerSelections** | **map[string]bool** |  | 
**ManagedConfig** | **map[string]map[string]interface{}** | Component configuration values keyed by component key | 
**CustomerProvidedInputs** | **map[string]map[string]string** | String values keyed first by component key and then by input key | 
**Layers** | [**[]ClusterPlatformBindingLayerResponse**](ClusterPlatformBindingLayerResponse.md) |  | 

## Methods

### NewClusterPlatformBindingResponse

`func NewClusterPlatformBindingResponse(clusterId string, organizationId string, templateKey string, templateVersion string, layerSelections map[string]bool, managedConfig map[string]map[string]interface{}, customerProvidedInputs map[string]map[string]string, layers []ClusterPlatformBindingLayerResponse, ) *ClusterPlatformBindingResponse`

NewClusterPlatformBindingResponse instantiates a new ClusterPlatformBindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPlatformBindingResponseWithDefaults

`func NewClusterPlatformBindingResponseWithDefaults() *ClusterPlatformBindingResponse`

NewClusterPlatformBindingResponseWithDefaults instantiates a new ClusterPlatformBindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterPlatformBindingResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterPlatformBindingResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterPlatformBindingResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetOrganizationId

`func (o *ClusterPlatformBindingResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterPlatformBindingResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterPlatformBindingResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetTemplateKey

`func (o *ClusterPlatformBindingResponse) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *ClusterPlatformBindingResponse) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *ClusterPlatformBindingResponse) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.


### GetTemplateVersion

`func (o *ClusterPlatformBindingResponse) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *ClusterPlatformBindingResponse) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *ClusterPlatformBindingResponse) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.


### GetLayerSelections

`func (o *ClusterPlatformBindingResponse) GetLayerSelections() map[string]bool`

GetLayerSelections returns the LayerSelections field if non-nil, zero value otherwise.

### GetLayerSelectionsOk

`func (o *ClusterPlatformBindingResponse) GetLayerSelectionsOk() (*map[string]bool, bool)`

GetLayerSelectionsOk returns a tuple with the LayerSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerSelections

`func (o *ClusterPlatformBindingResponse) SetLayerSelections(v map[string]bool)`

SetLayerSelections sets LayerSelections field to given value.


### GetManagedConfig

`func (o *ClusterPlatformBindingResponse) GetManagedConfig() map[string]map[string]interface{}`

GetManagedConfig returns the ManagedConfig field if non-nil, zero value otherwise.

### GetManagedConfigOk

`func (o *ClusterPlatformBindingResponse) GetManagedConfigOk() (*map[string]map[string]interface{}, bool)`

GetManagedConfigOk returns a tuple with the ManagedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedConfig

`func (o *ClusterPlatformBindingResponse) SetManagedConfig(v map[string]map[string]interface{})`

SetManagedConfig sets ManagedConfig field to given value.


### GetCustomerProvidedInputs

`func (o *ClusterPlatformBindingResponse) GetCustomerProvidedInputs() map[string]map[string]string`

GetCustomerProvidedInputs returns the CustomerProvidedInputs field if non-nil, zero value otherwise.

### GetCustomerProvidedInputsOk

`func (o *ClusterPlatformBindingResponse) GetCustomerProvidedInputsOk() (*map[string]map[string]string, bool)`

GetCustomerProvidedInputsOk returns a tuple with the CustomerProvidedInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProvidedInputs

`func (o *ClusterPlatformBindingResponse) SetCustomerProvidedInputs(v map[string]map[string]string)`

SetCustomerProvidedInputs sets CustomerProvidedInputs field to given value.


### GetLayers

`func (o *ClusterPlatformBindingResponse) GetLayers() []ClusterPlatformBindingLayerResponse`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *ClusterPlatformBindingResponse) GetLayersOk() (*[]ClusterPlatformBindingLayerResponse, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *ClusterPlatformBindingResponse) SetLayers(v []ClusterPlatformBindingLayerResponse)`

SetLayers sets Layers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


