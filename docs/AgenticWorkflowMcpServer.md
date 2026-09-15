# AgenticWorkflowMcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Required** | **bool** | Whether the initial configuration or template requires this MCP server. | [default to false]

## Methods

### NewAgenticWorkflowMcpServer

`func NewAgenticWorkflowMcpServer(id string, required bool, ) *AgenticWorkflowMcpServer`

NewAgenticWorkflowMcpServer instantiates a new AgenticWorkflowMcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowMcpServerWithDefaults

`func NewAgenticWorkflowMcpServerWithDefaults() *AgenticWorkflowMcpServer`

NewAgenticWorkflowMcpServerWithDefaults instantiates a new AgenticWorkflowMcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgenticWorkflowMcpServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgenticWorkflowMcpServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgenticWorkflowMcpServer) SetId(v string)`

SetId sets Id field to given value.


### GetRequired

`func (o *AgenticWorkflowMcpServer) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AgenticWorkflowMcpServer) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AgenticWorkflowMcpServer) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


