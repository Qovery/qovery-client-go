# EnvironmentRestartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartDb** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEnvironmentRestartRequest

`func NewEnvironmentRestartRequest() *EnvironmentRestartRequest`

NewEnvironmentRestartRequest instantiates a new EnvironmentRestartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentRestartRequestWithDefaults

`func NewEnvironmentRestartRequestWithDefaults() *EnvironmentRestartRequest`

NewEnvironmentRestartRequestWithDefaults instantiates a new EnvironmentRestartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestartDb

`func (o *EnvironmentRestartRequest) GetRestartDb() bool`

GetRestartDb returns the RestartDb field if non-nil, zero value otherwise.

### GetRestartDbOk

`func (o *EnvironmentRestartRequest) GetRestartDbOk() (*bool, bool)`

GetRestartDbOk returns a tuple with the RestartDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartDb

`func (o *EnvironmentRestartRequest) SetRestartDb(v bool)`

SetRestartDb sets RestartDb field to given value.

### HasRestartDb

`func (o *EnvironmentRestartRequest) HasRestartDb() bool`

HasRestartDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


