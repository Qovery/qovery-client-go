# ObjectArrayItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) |  | 

## Methods

### NewObjectArrayItemResponse

`func NewObjectArrayItemResponse(type_ string, fields []FieldSchemaResponse, ) *ObjectArrayItemResponse`

NewObjectArrayItemResponse instantiates a new ObjectArrayItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectArrayItemResponseWithDefaults

`func NewObjectArrayItemResponseWithDefaults() *ObjectArrayItemResponse`

NewObjectArrayItemResponseWithDefaults instantiates a new ObjectArrayItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ObjectArrayItemResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectArrayItemResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectArrayItemResponse) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *ObjectArrayItemResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ObjectArrayItemResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ObjectArrayItemResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


