# AgenticWorkflowConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | **string** |  | 
**Headers** | Pointer to [**[]AgenticWorkflowHeader**](AgenticWorkflowHeader.md) |  | [optional] [default to []]
**Instructions** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewAgenticWorkflowConnector

`func NewAgenticWorkflowConnector(name string, url string, ) *AgenticWorkflowConnector`

NewAgenticWorkflowConnector instantiates a new AgenticWorkflowConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowConnectorWithDefaults

`func NewAgenticWorkflowConnectorWithDefaults() *AgenticWorkflowConnector`

NewAgenticWorkflowConnectorWithDefaults instantiates a new AgenticWorkflowConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgenticWorkflowConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgenticWorkflowConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgenticWorkflowConnector) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *AgenticWorkflowConnector) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgenticWorkflowConnector) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgenticWorkflowConnector) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *AgenticWorkflowConnector) GetHeaders() []AgenticWorkflowHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AgenticWorkflowConnector) GetHeadersOk() (*[]AgenticWorkflowHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AgenticWorkflowConnector) SetHeaders(v []AgenticWorkflowHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AgenticWorkflowConnector) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetInstructions

`func (o *AgenticWorkflowConnector) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgenticWorkflowConnector) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgenticWorkflowConnector) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgenticWorkflowConnector) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


