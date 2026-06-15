# BlueprintUpdatePreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the service | 
**Tag** | **string** | Catalog tag identifying the target blueprint version | 
**Icon** | **string** | Icon URL for the service | 
**Variables** | Pointer to [**[]BlueprintVariableRequest**](BlueprintVariableRequest.md) | Variable overrides to apply in the preview | [optional] 
**SpecOverrides** | Pointer to **map[string]interface{}** | Partial spec overrides merged on top of the blueprint manifest | [optional] 

## Methods

### NewBlueprintUpdatePreviewRequest

`func NewBlueprintUpdatePreviewRequest(name string, tag string, icon string, ) *BlueprintUpdatePreviewRequest`

NewBlueprintUpdatePreviewRequest instantiates a new BlueprintUpdatePreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdatePreviewRequestWithDefaults

`func NewBlueprintUpdatePreviewRequestWithDefaults() *BlueprintUpdatePreviewRequest`

NewBlueprintUpdatePreviewRequestWithDefaults instantiates a new BlueprintUpdatePreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintUpdatePreviewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintUpdatePreviewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintUpdatePreviewRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *BlueprintUpdatePreviewRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BlueprintUpdatePreviewRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BlueprintUpdatePreviewRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetIcon

`func (o *BlueprintUpdatePreviewRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BlueprintUpdatePreviewRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BlueprintUpdatePreviewRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetVariables

`func (o *BlueprintUpdatePreviewRequest) GetVariables() []BlueprintVariableRequest`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *BlueprintUpdatePreviewRequest) GetVariablesOk() (*[]BlueprintVariableRequest, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *BlueprintUpdatePreviewRequest) SetVariables(v []BlueprintVariableRequest)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *BlueprintUpdatePreviewRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *BlueprintUpdatePreviewRequest) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *BlueprintUpdatePreviewRequest) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetSpecOverrides

`func (o *BlueprintUpdatePreviewRequest) GetSpecOverrides() map[string]interface{}`

GetSpecOverrides returns the SpecOverrides field if non-nil, zero value otherwise.

### GetSpecOverridesOk

`func (o *BlueprintUpdatePreviewRequest) GetSpecOverridesOk() (*map[string]interface{}, bool)`

GetSpecOverridesOk returns a tuple with the SpecOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecOverrides

`func (o *BlueprintUpdatePreviewRequest) SetSpecOverrides(v map[string]interface{})`

SetSpecOverrides sets SpecOverrides field to given value.

### HasSpecOverrides

`func (o *BlueprintUpdatePreviewRequest) HasSpecOverrides() bool`

HasSpecOverrides returns a boolean if a field has been set.

### SetSpecOverridesNil

`func (o *BlueprintUpdatePreviewRequest) SetSpecOverridesNil(b bool)`

 SetSpecOverridesNil sets the value for SpecOverrides to be an explicit nil

### UnsetSpecOverrides
`func (o *BlueprintUpdatePreviewRequest) UnsetSpecOverrides()`

UnsetSpecOverrides ensures that no value is present for SpecOverrides, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


