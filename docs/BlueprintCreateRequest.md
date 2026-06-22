# BlueprintCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the service | 
**Tag** | **string** | Catalog tag identifying the blueprint version | 
**Icon** | **string** | Icon URL for the service | 
**Variables** | Pointer to [**[]BlueprintVariableRequest**](BlueprintVariableRequest.md) | Variable overrides for the blueprint | [optional] [default to []]
**SpecOverrides** | Pointer to [**BlueprintSpecOverrides**](BlueprintSpecOverrides.md) |  | [optional] 

## Methods

### NewBlueprintCreateRequest

`func NewBlueprintCreateRequest(name string, tag string, icon string, ) *BlueprintCreateRequest`

NewBlueprintCreateRequest instantiates a new BlueprintCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCreateRequestWithDefaults

`func NewBlueprintCreateRequestWithDefaults() *BlueprintCreateRequest`

NewBlueprintCreateRequestWithDefaults instantiates a new BlueprintCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *BlueprintCreateRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BlueprintCreateRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BlueprintCreateRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetIcon

`func (o *BlueprintCreateRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BlueprintCreateRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BlueprintCreateRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetVariables

`func (o *BlueprintCreateRequest) GetVariables() []BlueprintVariableRequest`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *BlueprintCreateRequest) GetVariablesOk() (*[]BlueprintVariableRequest, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *BlueprintCreateRequest) SetVariables(v []BlueprintVariableRequest)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *BlueprintCreateRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSpecOverrides

`func (o *BlueprintCreateRequest) GetSpecOverrides() BlueprintSpecOverrides`

GetSpecOverrides returns the SpecOverrides field if non-nil, zero value otherwise.

### GetSpecOverridesOk

`func (o *BlueprintCreateRequest) GetSpecOverridesOk() (*BlueprintSpecOverrides, bool)`

GetSpecOverridesOk returns a tuple with the SpecOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecOverrides

`func (o *BlueprintCreateRequest) SetSpecOverrides(v BlueprintSpecOverrides)`

SetSpecOverrides sets SpecOverrides field to given value.

### HasSpecOverrides

`func (o *BlueprintCreateRequest) HasSpecOverrides() bool`

HasSpecOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


