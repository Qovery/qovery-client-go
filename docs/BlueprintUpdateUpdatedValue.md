# BlueprintUpdateUpdatedValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CurrentDefaultValue** | Pointer to **NullableString** |  | [optional] 
**NewDefaultValue** | Pointer to **NullableString** |  | [optional] 
**CurrentValue** | Pointer to **NullableString** |  | [optional] 
**Type** | [**BlueprintManifestFieldType**](BlueprintManifestFieldType.md) |  | 
**AllowedValues** | Pointer to **[]string** | Values the user may choose from. Null &#x3D; unrestricted. | [optional] 
**IsSecret** | **bool** |  | 

## Methods

### NewBlueprintUpdateUpdatedValue

`func NewBlueprintUpdateUpdatedValue(name string, type_ BlueprintManifestFieldType, isSecret bool, ) *BlueprintUpdateUpdatedValue`

NewBlueprintUpdateUpdatedValue instantiates a new BlueprintUpdateUpdatedValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateUpdatedValueWithDefaults

`func NewBlueprintUpdateUpdatedValueWithDefaults() *BlueprintUpdateUpdatedValue`

NewBlueprintUpdateUpdatedValueWithDefaults instantiates a new BlueprintUpdateUpdatedValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintUpdateUpdatedValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintUpdateUpdatedValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintUpdateUpdatedValue) SetName(v string)`

SetName sets Name field to given value.


### GetCurrentDefaultValue

`func (o *BlueprintUpdateUpdatedValue) GetCurrentDefaultValue() string`

GetCurrentDefaultValue returns the CurrentDefaultValue field if non-nil, zero value otherwise.

### GetCurrentDefaultValueOk

`func (o *BlueprintUpdateUpdatedValue) GetCurrentDefaultValueOk() (*string, bool)`

GetCurrentDefaultValueOk returns a tuple with the CurrentDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDefaultValue

`func (o *BlueprintUpdateUpdatedValue) SetCurrentDefaultValue(v string)`

SetCurrentDefaultValue sets CurrentDefaultValue field to given value.

### HasCurrentDefaultValue

`func (o *BlueprintUpdateUpdatedValue) HasCurrentDefaultValue() bool`

HasCurrentDefaultValue returns a boolean if a field has been set.

### SetCurrentDefaultValueNil

`func (o *BlueprintUpdateUpdatedValue) SetCurrentDefaultValueNil(b bool)`

 SetCurrentDefaultValueNil sets the value for CurrentDefaultValue to be an explicit nil

### UnsetCurrentDefaultValue
`func (o *BlueprintUpdateUpdatedValue) UnsetCurrentDefaultValue()`

UnsetCurrentDefaultValue ensures that no value is present for CurrentDefaultValue, not even an explicit nil
### GetNewDefaultValue

`func (o *BlueprintUpdateUpdatedValue) GetNewDefaultValue() string`

GetNewDefaultValue returns the NewDefaultValue field if non-nil, zero value otherwise.

### GetNewDefaultValueOk

`func (o *BlueprintUpdateUpdatedValue) GetNewDefaultValueOk() (*string, bool)`

GetNewDefaultValueOk returns a tuple with the NewDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDefaultValue

`func (o *BlueprintUpdateUpdatedValue) SetNewDefaultValue(v string)`

SetNewDefaultValue sets NewDefaultValue field to given value.

### HasNewDefaultValue

`func (o *BlueprintUpdateUpdatedValue) HasNewDefaultValue() bool`

HasNewDefaultValue returns a boolean if a field has been set.

### SetNewDefaultValueNil

`func (o *BlueprintUpdateUpdatedValue) SetNewDefaultValueNil(b bool)`

 SetNewDefaultValueNil sets the value for NewDefaultValue to be an explicit nil

### UnsetNewDefaultValue
`func (o *BlueprintUpdateUpdatedValue) UnsetNewDefaultValue()`

UnsetNewDefaultValue ensures that no value is present for NewDefaultValue, not even an explicit nil
### GetCurrentValue

`func (o *BlueprintUpdateUpdatedValue) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *BlueprintUpdateUpdatedValue) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *BlueprintUpdateUpdatedValue) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *BlueprintUpdateUpdatedValue) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *BlueprintUpdateUpdatedValue) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *BlueprintUpdateUpdatedValue) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
### GetType

`func (o *BlueprintUpdateUpdatedValue) GetType() BlueprintManifestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintUpdateUpdatedValue) GetTypeOk() (*BlueprintManifestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintUpdateUpdatedValue) SetType(v BlueprintManifestFieldType)`

SetType sets Type field to given value.


### GetAllowedValues

`func (o *BlueprintUpdateUpdatedValue) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintUpdateUpdatedValue) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintUpdateUpdatedValue) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintUpdateUpdatedValue) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintUpdateUpdatedValue) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintUpdateUpdatedValue) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetIsSecret

`func (o *BlueprintUpdateUpdatedValue) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintUpdateUpdatedValue) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintUpdateUpdatedValue) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


