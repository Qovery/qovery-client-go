# BlueprintManifestEngineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Provider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**Terraform** | [**BlueprintManifestVersionBlock**](BlueprintManifestVersionBlock.md) |  | 
**Credentials** | Pointer to [**BlueprintManifestCredentialsConfig**](BlueprintManifestCredentialsConfig.md) |  | [optional] 
**Backend** | Pointer to [**BlueprintManifestBackendConfig**](BlueprintManifestBackendConfig.md) |  | [optional] 
**Timeout** | Pointer to **NullableInt64** |  | [optional] 
**Resources** | Pointer to [**BlueprintManifestResourcesConfig**](BlueprintManifestResourcesConfig.md) |  | [optional] 
**Opentofu** | [**BlueprintManifestVersionBlock**](BlueprintManifestVersionBlock.md) |  | 
**Chart** | [**BlueprintManifestChartConfig**](BlueprintManifestChartConfig.md) |  | 

## Methods

### NewBlueprintManifestEngineSpec

`func NewBlueprintManifestEngineSpec(type_ string, provider CloudVendorEnum, terraform BlueprintManifestVersionBlock, opentofu BlueprintManifestVersionBlock, chart BlueprintManifestChartConfig, ) *BlueprintManifestEngineSpec`

NewBlueprintManifestEngineSpec instantiates a new BlueprintManifestEngineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestEngineSpecWithDefaults

`func NewBlueprintManifestEngineSpecWithDefaults() *BlueprintManifestEngineSpec`

NewBlueprintManifestEngineSpecWithDefaults instantiates a new BlueprintManifestEngineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintManifestEngineSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestEngineSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestEngineSpec) SetType(v string)`

SetType sets Type field to given value.


### GetProvider

`func (o *BlueprintManifestEngineSpec) GetProvider() CloudVendorEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BlueprintManifestEngineSpec) GetProviderOk() (*CloudVendorEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BlueprintManifestEngineSpec) SetProvider(v CloudVendorEnum)`

SetProvider sets Provider field to given value.


### GetTerraform

`func (o *BlueprintManifestEngineSpec) GetTerraform() BlueprintManifestVersionBlock`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *BlueprintManifestEngineSpec) GetTerraformOk() (*BlueprintManifestVersionBlock, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *BlueprintManifestEngineSpec) SetTerraform(v BlueprintManifestVersionBlock)`

SetTerraform sets Terraform field to given value.


### GetCredentials

`func (o *BlueprintManifestEngineSpec) GetCredentials() BlueprintManifestCredentialsConfig`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BlueprintManifestEngineSpec) GetCredentialsOk() (*BlueprintManifestCredentialsConfig, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BlueprintManifestEngineSpec) SetCredentials(v BlueprintManifestCredentialsConfig)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BlueprintManifestEngineSpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBackend

`func (o *BlueprintManifestEngineSpec) GetBackend() BlueprintManifestBackendConfig`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *BlueprintManifestEngineSpec) GetBackendOk() (*BlueprintManifestBackendConfig, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *BlueprintManifestEngineSpec) SetBackend(v BlueprintManifestBackendConfig)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *BlueprintManifestEngineSpec) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetTimeout

`func (o *BlueprintManifestEngineSpec) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BlueprintManifestEngineSpec) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BlueprintManifestEngineSpec) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BlueprintManifestEngineSpec) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BlueprintManifestEngineSpec) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BlueprintManifestEngineSpec) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetResources

`func (o *BlueprintManifestEngineSpec) GetResources() BlueprintManifestResourcesConfig`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintManifestEngineSpec) GetResourcesOk() (*BlueprintManifestResourcesConfig, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintManifestEngineSpec) SetResources(v BlueprintManifestResourcesConfig)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintManifestEngineSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetOpentofu

`func (o *BlueprintManifestEngineSpec) GetOpentofu() BlueprintManifestVersionBlock`

GetOpentofu returns the Opentofu field if non-nil, zero value otherwise.

### GetOpentofuOk

`func (o *BlueprintManifestEngineSpec) GetOpentofuOk() (*BlueprintManifestVersionBlock, bool)`

GetOpentofuOk returns a tuple with the Opentofu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentofu

`func (o *BlueprintManifestEngineSpec) SetOpentofu(v BlueprintManifestVersionBlock)`

SetOpentofu sets Opentofu field to given value.


### GetChart

`func (o *BlueprintManifestEngineSpec) GetChart() BlueprintManifestChartConfig`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *BlueprintManifestEngineSpec) GetChartOk() (*BlueprintManifestChartConfig, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *BlueprintManifestEngineSpec) SetChart(v BlueprintManifestChartConfig)`

SetChart sets Chart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


