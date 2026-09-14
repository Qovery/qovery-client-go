# ArrayItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Constraints** | [**FieldSchemaConstraintsResponse**](FieldSchemaConstraintsResponse.md) |  | 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) |  | 

## Methods

### NewArrayItemResponse

`func NewArrayItemResponse(type_ string, constraints FieldSchemaConstraintsResponse, fields []FieldSchemaResponse, ) *ArrayItemResponse`

NewArrayItemResponse instantiates a new ArrayItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayItemResponseWithDefaults

`func NewArrayItemResponseWithDefaults() *ArrayItemResponse`

NewArrayItemResponseWithDefaults instantiates a new ArrayItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArrayItemResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArrayItemResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArrayItemResponse) SetType(v string)`

SetType sets Type field to given value.


### GetConstraints

`func (o *ArrayItemResponse) GetConstraints() FieldSchemaConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ArrayItemResponse) GetConstraintsOk() (*FieldSchemaConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ArrayItemResponse) SetConstraints(v FieldSchemaConstraintsResponse)`

SetConstraints sets Constraints field to given value.


### GetFields

`func (o *ArrayItemResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ArrayItemResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ArrayItemResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


