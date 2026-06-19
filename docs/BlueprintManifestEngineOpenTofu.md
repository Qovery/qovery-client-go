# BlueprintManifestEngineOpenTofu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Provider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**Opentofu** | [**BlueprintManifestVersionBlock**](BlueprintManifestVersionBlock.md) |  | 
**Credentials** | Pointer to [**BlueprintManifestCredentialsConfig**](BlueprintManifestCredentialsConfig.md) |  | [optional] 
**Backend** | Pointer to [**BlueprintManifestBackendConfig**](BlueprintManifestBackendConfig.md) |  | [optional] 
**Timeout** | Pointer to **NullableInt64** |  | [optional] 
**Resources** | Pointer to [**BlueprintManifestResourcesConfig**](BlueprintManifestResourcesConfig.md) |  | [optional] 

## Methods

### NewBlueprintManifestEngineOpenTofu

`func NewBlueprintManifestEngineOpenTofu(type_ string, provider CloudVendorEnum, opentofu BlueprintManifestVersionBlock, ) *BlueprintManifestEngineOpenTofu`

NewBlueprintManifestEngineOpenTofu instantiates a new BlueprintManifestEngineOpenTofu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestEngineOpenTofuWithDefaults

`func NewBlueprintManifestEngineOpenTofuWithDefaults() *BlueprintManifestEngineOpenTofu`

NewBlueprintManifestEngineOpenTofuWithDefaults instantiates a new BlueprintManifestEngineOpenTofu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintManifestEngineOpenTofu) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestEngineOpenTofu) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestEngineOpenTofu) SetType(v string)`

SetType sets Type field to given value.


### GetProvider

`func (o *BlueprintManifestEngineOpenTofu) GetProvider() CloudVendorEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BlueprintManifestEngineOpenTofu) GetProviderOk() (*CloudVendorEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BlueprintManifestEngineOpenTofu) SetProvider(v CloudVendorEnum)`

SetProvider sets Provider field to given value.


### GetOpentofu

`func (o *BlueprintManifestEngineOpenTofu) GetOpentofu() BlueprintManifestVersionBlock`

GetOpentofu returns the Opentofu field if non-nil, zero value otherwise.

### GetOpentofuOk

`func (o *BlueprintManifestEngineOpenTofu) GetOpentofuOk() (*BlueprintManifestVersionBlock, bool)`

GetOpentofuOk returns a tuple with the Opentofu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentofu

`func (o *BlueprintManifestEngineOpenTofu) SetOpentofu(v BlueprintManifestVersionBlock)`

SetOpentofu sets Opentofu field to given value.


### GetCredentials

`func (o *BlueprintManifestEngineOpenTofu) GetCredentials() BlueprintManifestCredentialsConfig`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BlueprintManifestEngineOpenTofu) GetCredentialsOk() (*BlueprintManifestCredentialsConfig, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BlueprintManifestEngineOpenTofu) SetCredentials(v BlueprintManifestCredentialsConfig)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BlueprintManifestEngineOpenTofu) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBackend

`func (o *BlueprintManifestEngineOpenTofu) GetBackend() BlueprintManifestBackendConfig`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *BlueprintManifestEngineOpenTofu) GetBackendOk() (*BlueprintManifestBackendConfig, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *BlueprintManifestEngineOpenTofu) SetBackend(v BlueprintManifestBackendConfig)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *BlueprintManifestEngineOpenTofu) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetTimeout

`func (o *BlueprintManifestEngineOpenTofu) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BlueprintManifestEngineOpenTofu) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BlueprintManifestEngineOpenTofu) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BlueprintManifestEngineOpenTofu) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BlueprintManifestEngineOpenTofu) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BlueprintManifestEngineOpenTofu) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetResources

`func (o *BlueprintManifestEngineOpenTofu) GetResources() BlueprintManifestResourcesConfig`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintManifestEngineOpenTofu) GetResourcesOk() (*BlueprintManifestResourcesConfig, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintManifestEngineOpenTofu) SetResources(v BlueprintManifestResourcesConfig)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintManifestEngineOpenTofu) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


