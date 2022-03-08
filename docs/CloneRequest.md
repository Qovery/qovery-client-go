# CloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**ClusterId** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 

## Methods

### NewCloneRequest

`func NewCloneRequest(name string, ) *CloneRequest`

NewCloneRequest instantiates a new CloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneRequestWithDefaults

`func NewCloneRequestWithDefaults() *CloneRequest`

NewCloneRequestWithDefaults instantiates a new CloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneRequest) SetName(v string)`

SetName sets Name field to given value.


### GetClusterId

`func (o *CloneRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloneRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloneRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloneRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetMode

`func (o *CloneRequest) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CloneRequest) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CloneRequest) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CloneRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


