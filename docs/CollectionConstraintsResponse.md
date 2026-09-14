# CollectionConstraintsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinItems** | Pointer to **NullableInt64** |  | [optional] 
**MaxItems** | Pointer to **NullableInt64** |  | [optional] 
**UniqueItems** | **bool** |  | 

## Methods

### NewCollectionConstraintsResponse

`func NewCollectionConstraintsResponse(uniqueItems bool, ) *CollectionConstraintsResponse`

NewCollectionConstraintsResponse instantiates a new CollectionConstraintsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionConstraintsResponseWithDefaults

`func NewCollectionConstraintsResponseWithDefaults() *CollectionConstraintsResponse`

NewCollectionConstraintsResponseWithDefaults instantiates a new CollectionConstraintsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinItems

`func (o *CollectionConstraintsResponse) GetMinItems() int64`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *CollectionConstraintsResponse) GetMinItemsOk() (*int64, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *CollectionConstraintsResponse) SetMinItems(v int64)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *CollectionConstraintsResponse) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *CollectionConstraintsResponse) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *CollectionConstraintsResponse) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *CollectionConstraintsResponse) GetMaxItems() int64`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *CollectionConstraintsResponse) GetMaxItemsOk() (*int64, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *CollectionConstraintsResponse) SetMaxItems(v int64)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *CollectionConstraintsResponse) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *CollectionConstraintsResponse) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *CollectionConstraintsResponse) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetUniqueItems

`func (o *CollectionConstraintsResponse) GetUniqueItems() bool`

GetUniqueItems returns the UniqueItems field if non-nil, zero value otherwise.

### GetUniqueItemsOk

`func (o *CollectionConstraintsResponse) GetUniqueItemsOk() (*bool, bool)`

GetUniqueItemsOk returns a tuple with the UniqueItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueItems

`func (o *CollectionConstraintsResponse) SetUniqueItems(v bool)`

SetUniqueItems sets UniqueItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


