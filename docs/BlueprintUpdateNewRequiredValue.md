# BlueprintUpdateNewRequiredValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**BlueprintManifestFieldType**](BlueprintManifestFieldType.md) |  | 
**AllowedValues** | Pointer to **[]string** | Values the user may choose from. Null &#x3D; unrestricted. | [optional] 
**IsSecret** | **bool** |  | 

## Methods

### NewBlueprintUpdateNewRequiredValue

`func NewBlueprintUpdateNewRequiredValue(name string, type_ BlueprintManifestFieldType, isSecret bool, ) *BlueprintUpdateNewRequiredValue`

NewBlueprintUpdateNewRequiredValue instantiates a new BlueprintUpdateNewRequiredValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateNewRequiredValueWithDefaults

`func NewBlueprintUpdateNewRequiredValueWithDefaults() *BlueprintUpdateNewRequiredValue`

NewBlueprintUpdateNewRequiredValueWithDefaults instantiates a new BlueprintUpdateNewRequiredValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintUpdateNewRequiredValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintUpdateNewRequiredValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintUpdateNewRequiredValue) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlueprintUpdateNewRequiredValue) GetType() BlueprintManifestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintUpdateNewRequiredValue) GetTypeOk() (*BlueprintManifestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintUpdateNewRequiredValue) SetType(v BlueprintManifestFieldType)`

SetType sets Type field to given value.


### GetAllowedValues

`func (o *BlueprintUpdateNewRequiredValue) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BlueprintUpdateNewRequiredValue) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BlueprintUpdateNewRequiredValue) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BlueprintUpdateNewRequiredValue) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *BlueprintUpdateNewRequiredValue) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *BlueprintUpdateNewRequiredValue) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetIsSecret

`func (o *BlueprintUpdateNewRequiredValue) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BlueprintUpdateNewRequiredValue) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BlueprintUpdateNewRequiredValue) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


