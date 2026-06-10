# BlueprintManifestFieldType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Min** | Pointer to **NullableInt64** | Lower bound, only present for number fields | [optional] 
**Max** | Pointer to **NullableInt64** | Upper bound, only present for number fields | [optional] 

## Methods

### NewBlueprintManifestFieldType

`func NewBlueprintManifestFieldType(type_ string, ) *BlueprintManifestFieldType`

NewBlueprintManifestFieldType instantiates a new BlueprintManifestFieldType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestFieldTypeWithDefaults

`func NewBlueprintManifestFieldTypeWithDefaults() *BlueprintManifestFieldType`

NewBlueprintManifestFieldTypeWithDefaults instantiates a new BlueprintManifestFieldType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintManifestFieldType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestFieldType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestFieldType) SetType(v string)`

SetType sets Type field to given value.


### GetMin

`func (o *BlueprintManifestFieldType) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *BlueprintManifestFieldType) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *BlueprintManifestFieldType) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *BlueprintManifestFieldType) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *BlueprintManifestFieldType) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *BlueprintManifestFieldType) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *BlueprintManifestFieldType) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *BlueprintManifestFieldType) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *BlueprintManifestFieldType) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *BlueprintManifestFieldType) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *BlueprintManifestFieldType) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *BlueprintManifestFieldType) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


