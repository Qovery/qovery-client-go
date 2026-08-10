# McpServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique MCP server name within the organization | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Url** | **string** | HTTPS URL of the remote MCP server | 
**Headers** | Pointer to **map[string]string** | HTTP headers sent to the MCP server. Header values are encrypted and never returned by the API. | [optional] [default to {}]

## Methods

### NewMcpServerRequest

`func NewMcpServerRequest(name string, url string, ) *McpServerRequest`

NewMcpServerRequest instantiates a new McpServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestWithDefaults

`func NewMcpServerRequestWithDefaults() *McpServerRequest`

NewMcpServerRequestWithDefaults instantiates a new McpServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *McpServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *McpServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *McpServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *McpServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *McpServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *McpServerRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *McpServerRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *McpServerRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *McpServerRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *McpServerRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


