# ScalarArrayItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Constraints** | [**FieldSchemaConstraintsResponse**](FieldSchemaConstraintsResponse.md) |  | 

## Methods

### NewScalarArrayItemResponse

`func NewScalarArrayItemResponse(type_ string, constraints FieldSchemaConstraintsResponse, ) *ScalarArrayItemResponse`

NewScalarArrayItemResponse instantiates a new ScalarArrayItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalarArrayItemResponseWithDefaults

`func NewScalarArrayItemResponseWithDefaults() *ScalarArrayItemResponse`

NewScalarArrayItemResponseWithDefaults instantiates a new ScalarArrayItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScalarArrayItemResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScalarArrayItemResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScalarArrayItemResponse) SetType(v string)`

SetType sets Type field to given value.


### GetConstraints

`func (o *ScalarArrayItemResponse) GetConstraints() FieldSchemaConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ScalarArrayItemResponse) GetConstraintsOk() (*FieldSchemaConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ScalarArrayItemResponse) SetConstraints(v FieldSchemaConstraintsResponse)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


