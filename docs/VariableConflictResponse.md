# VariableConflictResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | **int32** |  | 
**Detail** | **string** | Says which scope holds the existing variable | 
**Instance** | Pointer to **string** |  | [optional] 
**ExistingVariable** | Pointer to [**NullableExistingVariableResponse**](ExistingVariableResponse.md) |  | [optional] 

## Methods

### NewVariableConflictResponse

`func NewVariableConflictResponse(status int32, detail string, ) *VariableConflictResponse`

NewVariableConflictResponse instantiates a new VariableConflictResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableConflictResponseWithDefaults

`func NewVariableConflictResponseWithDefaults() *VariableConflictResponse`

NewVariableConflictResponseWithDefaults instantiates a new VariableConflictResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VariableConflictResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariableConflictResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariableConflictResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VariableConflictResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *VariableConflictResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VariableConflictResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VariableConflictResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VariableConflictResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *VariableConflictResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VariableConflictResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VariableConflictResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *VariableConflictResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *VariableConflictResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *VariableConflictResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *VariableConflictResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *VariableConflictResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *VariableConflictResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *VariableConflictResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetExistingVariable

`func (o *VariableConflictResponse) GetExistingVariable() ExistingVariableResponse`

GetExistingVariable returns the ExistingVariable field if non-nil, zero value otherwise.

### GetExistingVariableOk

`func (o *VariableConflictResponse) GetExistingVariableOk() (*ExistingVariableResponse, bool)`

GetExistingVariableOk returns a tuple with the ExistingVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingVariable

`func (o *VariableConflictResponse) SetExistingVariable(v ExistingVariableResponse)`

SetExistingVariable sets ExistingVariable field to given value.

### HasExistingVariable

`func (o *VariableConflictResponse) HasExistingVariable() bool`

HasExistingVariable returns a boolean if a field has been set.

### SetExistingVariableNil

`func (o *VariableConflictResponse) SetExistingVariableNil(b bool)`

 SetExistingVariableNil sets the value for ExistingVariable to be an explicit nil

### UnsetExistingVariable
`func (o *VariableConflictResponse) UnsetExistingVariable()`

UnsetExistingVariable ensures that no value is present for ExistingVariable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


