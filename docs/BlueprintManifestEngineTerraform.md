# BlueprintManifestEngineTerraform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Provider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**Terraform** | [**BlueprintManifestVersionBlock**](BlueprintManifestVersionBlock.md) |  | 
**Credentials** | Pointer to [**BlueprintManifestCredentialsConfig**](BlueprintManifestCredentialsConfig.md) |  | [optional] 
**Backend** | Pointer to [**BlueprintManifestBackendConfig**](BlueprintManifestBackendConfig.md) |  | [optional] 
**Timeout** | Pointer to **NullableInt64** | Default execution timeout in seconds. | [optional] 
**Resources** | Pointer to [**BlueprintManifestResourcesConfig**](BlueprintManifestResourcesConfig.md) |  | [optional] 

## Methods

### NewBlueprintManifestEngineTerraform

`func NewBlueprintManifestEngineTerraform(type_ string, provider CloudVendorEnum, terraform BlueprintManifestVersionBlock, ) *BlueprintManifestEngineTerraform`

NewBlueprintManifestEngineTerraform instantiates a new BlueprintManifestEngineTerraform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestEngineTerraformWithDefaults

`func NewBlueprintManifestEngineTerraformWithDefaults() *BlueprintManifestEngineTerraform`

NewBlueprintManifestEngineTerraformWithDefaults instantiates a new BlueprintManifestEngineTerraform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintManifestEngineTerraform) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestEngineTerraform) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestEngineTerraform) SetType(v string)`

SetType sets Type field to given value.


### GetProvider

`func (o *BlueprintManifestEngineTerraform) GetProvider() CloudVendorEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BlueprintManifestEngineTerraform) GetProviderOk() (*CloudVendorEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BlueprintManifestEngineTerraform) SetProvider(v CloudVendorEnum)`

SetProvider sets Provider field to given value.


### GetTerraform

`func (o *BlueprintManifestEngineTerraform) GetTerraform() BlueprintManifestVersionBlock`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *BlueprintManifestEngineTerraform) GetTerraformOk() (*BlueprintManifestVersionBlock, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *BlueprintManifestEngineTerraform) SetTerraform(v BlueprintManifestVersionBlock)`

SetTerraform sets Terraform field to given value.


### GetCredentials

`func (o *BlueprintManifestEngineTerraform) GetCredentials() BlueprintManifestCredentialsConfig`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BlueprintManifestEngineTerraform) GetCredentialsOk() (*BlueprintManifestCredentialsConfig, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BlueprintManifestEngineTerraform) SetCredentials(v BlueprintManifestCredentialsConfig)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BlueprintManifestEngineTerraform) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBackend

`func (o *BlueprintManifestEngineTerraform) GetBackend() BlueprintManifestBackendConfig`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *BlueprintManifestEngineTerraform) GetBackendOk() (*BlueprintManifestBackendConfig, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *BlueprintManifestEngineTerraform) SetBackend(v BlueprintManifestBackendConfig)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *BlueprintManifestEngineTerraform) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetTimeout

`func (o *BlueprintManifestEngineTerraform) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BlueprintManifestEngineTerraform) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BlueprintManifestEngineTerraform) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BlueprintManifestEngineTerraform) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BlueprintManifestEngineTerraform) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BlueprintManifestEngineTerraform) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetResources

`func (o *BlueprintManifestEngineTerraform) GetResources() BlueprintManifestResourcesConfig`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintManifestEngineTerraform) GetResourcesOk() (*BlueprintManifestResourcesConfig, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintManifestEngineTerraform) SetResources(v BlueprintManifestResourcesConfig)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintManifestEngineTerraform) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


