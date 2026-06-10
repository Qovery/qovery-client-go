# BlueprintManifestVariableField

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

## Methods

### NewBlueprintManifestVariableField

`func NewBlueprintManifestVariableField(kind string, name string, type_ BlueprintManifestFieldType, required bool, isSecret bool, ) *BlueprintManifestVariableField`

NewBlueprintManifestVariableField instantiates a new BlueprintManifestVariableField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestVariableFieldWithDefaults

`func NewBlueprintManifestVariableFieldWithDefaults() *BlueprintManifestVariableField`

NewBlueprintManifestVariableFieldWithDefaults instantiates a new BlueprintManifestVariableField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BlueprintManifestVariableField) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlueprintManifestVariableField) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlueprintManifestVariableField) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *BlueprintManifestVariableField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintManifestVariableField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintManifestVariableField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlueprintManifestVariableField) GetType() BlueprintManifestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestVariableField) GetTypeOk() (*BlueprintManifestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestVariableField) SetType(v BlueprintManifestFieldType)`

SetType sets Type field to given value.


### GetRequired

`func (o *BlueprintManifestVariableField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BlueprintManifestVariableField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BlueprintManifestVariableField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetIsSecret

`func (o *BlueprintManifestVariableField) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintManifestVariableField) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintManifestVariableField) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetDescription

`func (o *BlueprintManifestVariableField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintManifestVariableField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintManifestVariableField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintManifestVariableField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BlueprintManifestVariableField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BlueprintManifestVariableField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowedValues

`func (o *BlueprintManifestVariableField) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintManifestVariableField) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintManifestVariableField) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintManifestVariableField) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintManifestVariableField) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintManifestVariableField) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetDefaultValue

`func (o *BlueprintManifestVariableField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BlueprintManifestVariableField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BlueprintManifestVariableField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BlueprintManifestVariableField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *BlueprintManifestVariableField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *BlueprintManifestVariableField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


