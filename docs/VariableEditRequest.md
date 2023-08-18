# VariableEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | the key of the environment variable | 
**Value** | **string** | the value of the environment variable | 

## Methods

### NewVariableEditRequest

`func NewVariableEditRequest(key string, value string, ) *VariableEditRequest`

NewVariableEditRequest instantiates a new VariableEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableEditRequestWithDefaults

`func NewVariableEditRequestWithDefaults() *VariableEditRequest`

NewVariableEditRequestWithDefaults instantiates a new VariableEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableEditRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableEditRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableEditRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableEditRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableEditRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableEditRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


