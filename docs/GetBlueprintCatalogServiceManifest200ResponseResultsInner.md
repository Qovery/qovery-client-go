# GetBlueprintCatalogServiceManifest200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**BlueprintManifestFieldType**](BlueprintManifestFieldType.md) |  | 
**Required** | **bool** |  | 
**IsSecret** | **bool** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** | Origin of the auto-sourced value (e.g. cluster.region) | [optional] 
**Value** | Pointer to **NullableString** | Resolved value of the context variable at the time of the request | [optional] 

## Methods

### NewGetBlueprintCatalogServiceManifest200ResponseResultsInner

`func NewGetBlueprintCatalogServiceManifest200ResponseResultsInner(kind string, name string, type_ BlueprintManifestFieldType, required bool, isSecret bool, ) *GetBlueprintCatalogServiceManifest200ResponseResultsInner`

NewGetBlueprintCatalogServiceManifest200ResponseResultsInner instantiates a new GetBlueprintCatalogServiceManifest200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlueprintCatalogServiceManifest200ResponseResultsInnerWithDefaults

`func NewGetBlueprintCatalogServiceManifest200ResponseResultsInnerWithDefaults() *GetBlueprintCatalogServiceManifest200ResponseResultsInner`

NewGetBlueprintCatalogServiceManifest200ResponseResultsInnerWithDefaults instantiates a new GetBlueprintCatalogServiceManifest200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetType() BlueprintManifestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetTypeOk() (*BlueprintManifestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetType(v BlueprintManifestFieldType)`

SetType sets Type field to given value.


### GetRequired

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetIsSecret

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetDescription

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowedValues

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetDefaultValue

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetSource

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetValue

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *GetBlueprintCatalogServiceManifest200ResponseResultsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


