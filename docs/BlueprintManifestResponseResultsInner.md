# BlueprintManifestResponseResultsInner

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

### NewBlueprintManifestResponseResultsInner

`func NewBlueprintManifestResponseResultsInner(kind string, name string, type_ BlueprintManifestFieldType, required bool, isSecret bool, ) *BlueprintManifestResponseResultsInner`

NewBlueprintManifestResponseResultsInner instantiates a new BlueprintManifestResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestResponseResultsInnerWithDefaults

`func NewBlueprintManifestResponseResultsInnerWithDefaults() *BlueprintManifestResponseResultsInner`

NewBlueprintManifestResponseResultsInnerWithDefaults instantiates a new BlueprintManifestResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BlueprintManifestResponseResultsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlueprintManifestResponseResultsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlueprintManifestResponseResultsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *BlueprintManifestResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintManifestResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintManifestResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlueprintManifestResponseResultsInner) GetType() BlueprintManifestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestResponseResultsInner) GetTypeOk() (*BlueprintManifestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestResponseResultsInner) SetType(v BlueprintManifestFieldType)`

SetType sets Type field to given value.


### GetRequired

`func (o *BlueprintManifestResponseResultsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BlueprintManifestResponseResultsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BlueprintManifestResponseResultsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetIsSecret

`func (o *BlueprintManifestResponseResultsInner) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintManifestResponseResultsInner) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintManifestResponseResultsInner) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetDescription

`func (o *BlueprintManifestResponseResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintManifestResponseResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintManifestResponseResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintManifestResponseResultsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BlueprintManifestResponseResultsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BlueprintManifestResponseResultsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowedValues

`func (o *BlueprintManifestResponseResultsInner) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintManifestResponseResultsInner) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintManifestResponseResultsInner) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintManifestResponseResultsInner) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintManifestResponseResultsInner) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintManifestResponseResultsInner) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetDefaultValue

`func (o *BlueprintManifestResponseResultsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BlueprintManifestResponseResultsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BlueprintManifestResponseResultsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BlueprintManifestResponseResultsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *BlueprintManifestResponseResultsInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *BlueprintManifestResponseResultsInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetSource

`func (o *BlueprintManifestResponseResultsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BlueprintManifestResponseResultsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BlueprintManifestResponseResultsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BlueprintManifestResponseResultsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *BlueprintManifestResponseResultsInner) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *BlueprintManifestResponseResultsInner) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetValue

`func (o *BlueprintManifestResponseResultsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BlueprintManifestResponseResultsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BlueprintManifestResponseResultsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BlueprintManifestResponseResultsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BlueprintManifestResponseResultsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BlueprintManifestResponseResultsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


