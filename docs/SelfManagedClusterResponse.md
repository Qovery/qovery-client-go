# SelfManagedClusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OrganizationId** | **string** |  | 
**Name** | **string** |  | 
**Production** | **bool** |  | 
**Provider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**Region** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Credentials** | [**SelfManagedClusterCredentials**](SelfManagedClusterCredentials.md) |  | 
**Registry** | [**SelfManagedClusterRegistryResponse**](SelfManagedClusterRegistryResponse.md) |  | 
**Platform** | [**PlatformSelection**](PlatformSelection.md) |  | 
**ClusterInputs** | **map[string]map[string]string** | String values keyed first by component key and then by input key | 
**Layers** | [**[]ClusterPlatformBindingLayerResponse**](ClusterPlatformBindingLayerResponse.md) |  | 

## Methods

### NewSelfManagedClusterResponse

`func NewSelfManagedClusterResponse(id string, organizationId string, name string, production bool, provider CloudVendorEnum, region string, createdAt time.Time, credentials SelfManagedClusterCredentials, registry SelfManagedClusterRegistryResponse, platform PlatformSelection, clusterInputs map[string]map[string]string, layers []ClusterPlatformBindingLayerResponse, ) *SelfManagedClusterResponse`

NewSelfManagedClusterResponse instantiates a new SelfManagedClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfManagedClusterResponseWithDefaults

`func NewSelfManagedClusterResponseWithDefaults() *SelfManagedClusterResponse`

NewSelfManagedClusterResponseWithDefaults instantiates a new SelfManagedClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelfManagedClusterResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelfManagedClusterResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelfManagedClusterResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *SelfManagedClusterResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SelfManagedClusterResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SelfManagedClusterResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *SelfManagedClusterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelfManagedClusterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelfManagedClusterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProduction

`func (o *SelfManagedClusterResponse) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *SelfManagedClusterResponse) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *SelfManagedClusterResponse) SetProduction(v bool)`

SetProduction sets Production field to given value.


### GetProvider

`func (o *SelfManagedClusterResponse) GetProvider() CloudVendorEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SelfManagedClusterResponse) GetProviderOk() (*CloudVendorEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SelfManagedClusterResponse) SetProvider(v CloudVendorEnum)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *SelfManagedClusterResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SelfManagedClusterResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SelfManagedClusterResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCreatedAt

`func (o *SelfManagedClusterResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SelfManagedClusterResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SelfManagedClusterResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCredentials

`func (o *SelfManagedClusterResponse) GetCredentials() SelfManagedClusterCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SelfManagedClusterResponse) GetCredentialsOk() (*SelfManagedClusterCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SelfManagedClusterResponse) SetCredentials(v SelfManagedClusterCredentials)`

SetCredentials sets Credentials field to given value.


### GetRegistry

`func (o *SelfManagedClusterResponse) GetRegistry() SelfManagedClusterRegistryResponse`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SelfManagedClusterResponse) GetRegistryOk() (*SelfManagedClusterRegistryResponse, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SelfManagedClusterResponse) SetRegistry(v SelfManagedClusterRegistryResponse)`

SetRegistry sets Registry field to given value.


### GetPlatform

`func (o *SelfManagedClusterResponse) GetPlatform() PlatformSelection`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SelfManagedClusterResponse) GetPlatformOk() (*PlatformSelection, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SelfManagedClusterResponse) SetPlatform(v PlatformSelection)`

SetPlatform sets Platform field to given value.


### GetClusterInputs

`func (o *SelfManagedClusterResponse) GetClusterInputs() map[string]map[string]string`

GetClusterInputs returns the ClusterInputs field if non-nil, zero value otherwise.

### GetClusterInputsOk

`func (o *SelfManagedClusterResponse) GetClusterInputsOk() (*map[string]map[string]string, bool)`

GetClusterInputsOk returns a tuple with the ClusterInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInputs

`func (o *SelfManagedClusterResponse) SetClusterInputs(v map[string]map[string]string)`

SetClusterInputs sets ClusterInputs field to given value.


### GetLayers

`func (o *SelfManagedClusterResponse) GetLayers() []ClusterPlatformBindingLayerResponse`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *SelfManagedClusterResponse) GetLayersOk() (*[]ClusterPlatformBindingLayerResponse, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *SelfManagedClusterResponse) SetLayers(v []ClusterPlatformBindingLayerResponse)`

SetLayers sets Layers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


