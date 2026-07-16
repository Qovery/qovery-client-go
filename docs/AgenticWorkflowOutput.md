# AgenticWorkflowOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to [**[]AgenticWorkflowHeader**](AgenticWorkflowHeader.md) |  | [optional] [default to []]
**Instructions** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewAgenticWorkflowOutput

`func NewAgenticWorkflowOutput(name string, ) *AgenticWorkflowOutput`

NewAgenticWorkflowOutput instantiates a new AgenticWorkflowOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowOutputWithDefaults

`func NewAgenticWorkflowOutputWithDefaults() *AgenticWorkflowOutput`

NewAgenticWorkflowOutputWithDefaults instantiates a new AgenticWorkflowOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgenticWorkflowOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgenticWorkflowOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgenticWorkflowOutput) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *AgenticWorkflowOutput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgenticWorkflowOutput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgenticWorkflowOutput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AgenticWorkflowOutput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AgenticWorkflowOutput) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AgenticWorkflowOutput) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetHeaders

`func (o *AgenticWorkflowOutput) GetHeaders() []AgenticWorkflowHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AgenticWorkflowOutput) GetHeadersOk() (*[]AgenticWorkflowHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AgenticWorkflowOutput) SetHeaders(v []AgenticWorkflowHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AgenticWorkflowOutput) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetInstructions

`func (o *AgenticWorkflowOutput) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgenticWorkflowOutput) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgenticWorkflowOutput) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgenticWorkflowOutput) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


