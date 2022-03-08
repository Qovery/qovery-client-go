# EnvironmentEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 

## Methods

### NewEnvironmentEditRequest

`func NewEnvironmentEditRequest() *EnvironmentEditRequest`

NewEnvironmentEditRequest instantiates a new EnvironmentEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditRequestWithDefaults

`func NewEnvironmentEditRequestWithDefaults() *EnvironmentEditRequest`

NewEnvironmentEditRequestWithDefaults instantiates a new EnvironmentEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentEditRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentEditRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *EnvironmentEditRequest) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EnvironmentEditRequest) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EnvironmentEditRequest) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EnvironmentEditRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


