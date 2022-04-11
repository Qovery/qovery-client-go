# UserResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewUserResponseList

`func NewUserResponseList() *UserResponseList`

NewUserResponseList instantiates a new UserResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseListWithDefaults

`func NewUserResponseListWithDefaults() *UserResponseList`

NewUserResponseListWithDefaults instantiates a new UserResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *UserResponseList) GetResults() []User`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserResponseList) GetResultsOk() (*[]User, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserResponseList) SetResults(v []User)`

SetResults sets Results field to given value.

### HasResults

`func (o *UserResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


