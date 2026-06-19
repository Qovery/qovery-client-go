# BlueprintManifestVersionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Catalog-declared default version for this engine. | 
**AllowedValues** | Pointer to **[]string** | Versions the user may choose from. Null &#x3D; unrestricted. | [optional] 

## Methods

### NewBlueprintManifestVersionBlock

`func NewBlueprintManifestVersionBlock(version string, ) *BlueprintManifestVersionBlock`

NewBlueprintManifestVersionBlock instantiates a new BlueprintManifestVersionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestVersionBlockWithDefaults

`func NewBlueprintManifestVersionBlockWithDefaults() *BlueprintManifestVersionBlock`

NewBlueprintManifestVersionBlockWithDefaults instantiates a new BlueprintManifestVersionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *BlueprintManifestVersionBlock) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintManifestVersionBlock) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintManifestVersionBlock) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAllowedValues

`func (o *BlueprintManifestVersionBlock) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintManifestVersionBlock) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintManifestVersionBlock) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintManifestVersionBlock) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintManifestVersionBlock) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintManifestVersionBlock) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


