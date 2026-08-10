# McpServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Url** | **string** | HTTPS URL of the remote MCP server | 
**HeaderNames** | **[]string** | Names of the configured HTTP headers. Header values are never returned. | 

## Methods

### NewMcpServerResponse

`func NewMcpServerResponse(id string, createdAt time.Time, updatedAt time.Time, name string, description string, url string, headerNames []string, ) *McpServerResponse`

NewMcpServerResponse instantiates a new McpServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerResponseWithDefaults

`func NewMcpServerResponseWithDefaults() *McpServerResponse`

NewMcpServerResponseWithDefaults instantiates a new McpServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McpServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *McpServerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *McpServerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *McpServerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *McpServerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *McpServerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *McpServerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *McpServerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *McpServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *McpServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *McpServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *McpServerResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaderNames

`func (o *McpServerResponse) GetHeaderNames() []string`

GetHeaderNames returns the HeaderNames field if non-nil, zero value otherwise.

### GetHeaderNamesOk

`func (o *McpServerResponse) GetHeaderNamesOk() (*[]string, bool)`

GetHeaderNamesOk returns a tuple with the HeaderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderNames

`func (o *McpServerResponse) SetHeaderNames(v []string)`

SetHeaderNames sets HeaderNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


