# BlueprintUpdateNewOptionalValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Type** | [**BlueprintManifestFieldType**](BlueprintManifestFieldType.md) |  | 
**AllowedValues** | Pointer to **[]string** | Values the user may choose from. Null &#x3D; unrestricted. | [optional] 
**IsSecret** | **bool** |  | 

## Methods

### NewBlueprintUpdateNewOptionalValue

`func NewBlueprintUpdateNewOptionalValue(name string, type_ BlueprintManifestFieldType, isSecret bool, ) *BlueprintUpdateNewOptionalValue`

NewBlueprintUpdateNewOptionalValue instantiates a new BlueprintUpdateNewOptionalValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateNewOptionalValueWithDefaults

`func NewBlueprintUpdateNewOptionalValueWithDefaults() *BlueprintUpdateNewOptionalValue`

NewBlueprintUpdateNewOptionalValueWithDefaults instantiates a new BlueprintUpdateNewOptionalValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintUpdateNewOptionalValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintUpdateNewOptionalValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintUpdateNewOptionalValue) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultValue

`func (o *BlueprintUpdateNewOptionalValue) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BlueprintUpdateNewOptionalValue) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BlueprintUpdateNewOptionalValue) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BlueprintUpdateNewOptionalValue) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *BlueprintUpdateNewOptionalValue) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *BlueprintUpdateNewOptionalValue) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetType

`func (o *BlueprintUpdateNewOptionalValue) GetType() BlueprintManifestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintUpdateNewOptionalValue) GetTypeOk() (*BlueprintManifestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintUpdateNewOptionalValue) SetType(v BlueprintManifestFieldType)`

SetType sets Type field to given value.


### GetAllowedValues

`func (o *BlueprintUpdateNewOptionalValue) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintUpdateNewOptionalValue) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintUpdateNewOptionalValue) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintUpdateNewOptionalValue) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintUpdateNewOptionalValue) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintUpdateNewOptionalValue) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetIsSecret

`func (o *BlueprintUpdateNewOptionalValue) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintUpdateNewOptionalValue) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintUpdateNewOptionalValue) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


