# BlueprintManifestBackendConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **NullableString** |  | [optional] 
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**Overridable** | Pointer to **bool** |  | [optional] [default to true]
**UserProvided** | Pointer to [**BlueprintManifestUserProvidedBackend**](BlueprintManifestUserProvidedBackend.md) |  | [optional] 

## Methods

### NewBlueprintManifestBackendConfig

`func NewBlueprintManifestBackendConfig() *BlueprintManifestBackendConfig`

NewBlueprintManifestBackendConfig instantiates a new BlueprintManifestBackendConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestBackendConfigWithDefaults

`func NewBlueprintManifestBackendConfigWithDefaults() *BlueprintManifestBackendConfig`

NewBlueprintManifestBackendConfigWithDefaults instantiates a new BlueprintManifestBackendConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *BlueprintManifestBackendConfig) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BlueprintManifestBackendConfig) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BlueprintManifestBackendConfig) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *BlueprintManifestBackendConfig) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *BlueprintManifestBackendConfig) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *BlueprintManifestBackendConfig) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetAllowedValues

`func (o *BlueprintManifestBackendConfig) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintManifestBackendConfig) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintManifestBackendConfig) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintManifestBackendConfig) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintManifestBackendConfig) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintManifestBackendConfig) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetOverridable

`func (o *BlueprintManifestBackendConfig) GetOverridable() bool`

GetOverridable returns the Overridable field if non-nil, zero value otherwise.

### GetOverridableOk

`func (o *BlueprintManifestBackendConfig) GetOverridableOk() (*bool, bool)`

GetOverridableOk returns a tuple with the Overridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridable

`func (o *BlueprintManifestBackendConfig) SetOverridable(v bool)`

SetOverridable sets Overridable field to given value.

### HasOverridable

`func (o *BlueprintManifestBackendConfig) HasOverridable() bool`

HasOverridable returns a boolean if a field has been set.

### GetUserProvided

`func (o *BlueprintManifestBackendConfig) GetUserProvided() BlueprintManifestUserProvidedBackend`

GetUserProvided returns the UserProvided field if non-nil, zero value otherwise.

### GetUserProvidedOk

`func (o *BlueprintManifestBackendConfig) GetUserProvidedOk() (*BlueprintManifestUserProvidedBackend, bool)`

GetUserProvidedOk returns a tuple with the UserProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvided

`func (o *BlueprintManifestBackendConfig) SetUserProvided(v BlueprintManifestUserProvidedBackend)`

SetUserProvided sets UserProvided field to given value.

### HasUserProvided

`func (o *BlueprintManifestBackendConfig) HasUserProvided() bool`

HasUserProvided returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


