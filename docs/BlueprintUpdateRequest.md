# BlueprintUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the service | 
**Tag** | **string** | Catalog tag identifying the target blueprint version | 
**Icon** | **string** | Icon URL for the service | 
**Variables** | Pointer to [**map[string]BlueprintUpdateVariableValue**](BlueprintUpdateVariableValue.md) | RFC 7396 patch map keyed by variable name. Non-null value upserts the variable; null value removes it. Absent keys are left untouched. Omitting the field entirely is equivalent to an empty map — no variables are modified. | [optional] 
**SpecOverrides** | Pointer to [**NullableBlueprintUpdateRequestSpecOverrides**](BlueprintUpdateRequestSpecOverrides.md) |  | [optional] 

## Methods

### NewBlueprintUpdateRequest

`func NewBlueprintUpdateRequest(name string, tag string, icon string, ) *BlueprintUpdateRequest`

NewBlueprintUpdateRequest instantiates a new BlueprintUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateRequestWithDefaults

`func NewBlueprintUpdateRequestWithDefaults() *BlueprintUpdateRequest`

NewBlueprintUpdateRequestWithDefaults instantiates a new BlueprintUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *BlueprintUpdateRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BlueprintUpdateRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BlueprintUpdateRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetIcon

`func (o *BlueprintUpdateRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BlueprintUpdateRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BlueprintUpdateRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetVariables

`func (o *BlueprintUpdateRequest) GetVariables() map[string]BlueprintUpdateVariableValue`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *BlueprintUpdateRequest) GetVariablesOk() (*map[string]BlueprintUpdateVariableValue, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *BlueprintUpdateRequest) SetVariables(v map[string]BlueprintUpdateVariableValue)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *BlueprintUpdateRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSpecOverrides

`func (o *BlueprintUpdateRequest) GetSpecOverrides() BlueprintUpdateRequestSpecOverrides`

GetSpecOverrides returns the SpecOverrides field if non-nil, zero value otherwise.

### GetSpecOverridesOk

`func (o *BlueprintUpdateRequest) GetSpecOverridesOk() (*BlueprintUpdateRequestSpecOverrides, bool)`

GetSpecOverridesOk returns a tuple with the SpecOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecOverrides

`func (o *BlueprintUpdateRequest) SetSpecOverrides(v BlueprintUpdateRequestSpecOverrides)`

SetSpecOverrides sets SpecOverrides field to given value.

### HasSpecOverrides

`func (o *BlueprintUpdateRequest) HasSpecOverrides() bool`

HasSpecOverrides returns a boolean if a field has been set.

### SetSpecOverridesNil

`func (o *BlueprintUpdateRequest) SetSpecOverridesNil(b bool)`

 SetSpecOverridesNil sets the value for SpecOverrides to be an explicit nil

### UnsetSpecOverrides
`func (o *BlueprintUpdateRequest) UnsetSpecOverrides()`

UnsetSpecOverrides ensures that no value is present for SpecOverrides, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


