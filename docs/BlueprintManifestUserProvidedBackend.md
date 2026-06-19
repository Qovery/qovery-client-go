# BlueprintManifestUserProvidedBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Terraform backend type (e.g. \&quot;s3\&quot;, \&quot;gcs\&quot;, \&quot;azurerm\&quot;). | 
**Config** | Pointer to **map[string]string** | Static backend config (bucket, region, etc.). | [optional] [default to {}]

## Methods

### NewBlueprintManifestUserProvidedBackend

`func NewBlueprintManifestUserProvidedBackend(type_ string, ) *BlueprintManifestUserProvidedBackend`

NewBlueprintManifestUserProvidedBackend instantiates a new BlueprintManifestUserProvidedBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestUserProvidedBackendWithDefaults

`func NewBlueprintManifestUserProvidedBackendWithDefaults() *BlueprintManifestUserProvidedBackend`

NewBlueprintManifestUserProvidedBackendWithDefaults instantiates a new BlueprintManifestUserProvidedBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintManifestUserProvidedBackend) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestUserProvidedBackend) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestUserProvidedBackend) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *BlueprintManifestUserProvidedBackend) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintManifestUserProvidedBackend) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintManifestUserProvidedBackend) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintManifestUserProvidedBackend) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


