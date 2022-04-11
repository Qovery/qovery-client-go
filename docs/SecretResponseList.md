# SecretResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Secret**](Secret.md) |  | [optional] 

## Methods

### NewSecretResponseList

`func NewSecretResponseList() *SecretResponseList`

NewSecretResponseList instantiates a new SecretResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseListWithDefaults

`func NewSecretResponseListWithDefaults() *SecretResponseList`

NewSecretResponseListWithDefaults instantiates a new SecretResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SecretResponseList) GetResults() []Secret`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SecretResponseList) GetResultsOk() (*[]Secret, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SecretResponseList) SetResults(v []Secret)`

SetResults sets Results field to given value.

### HasResults

`func (o *SecretResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


