# \MCPServersAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMcpServer**](MCPServersAPI.md#CreateMcpServer) | **Post** /organization/{organizationId}/mcpServer | Create an MCP server
[**DeleteMcpServer**](MCPServersAPI.md#DeleteMcpServer) | **Delete** /mcpServer/{mcpServerId} | Delete an MCP server
[**EditMcpServer**](MCPServersAPI.md#EditMcpServer) | **Put** /mcpServer/{mcpServerId} | Edit an MCP server
[**GetMcpServer**](MCPServersAPI.md#GetMcpServer) | **Get** /mcpServer/{mcpServerId} | Get an MCP server
[**ListMcpServers**](MCPServersAPI.md#ListMcpServers) | **Get** /organization/{organizationId}/mcpServer | List organization MCP servers



## CreateMcpServer

> McpServerResponse CreateMcpServer(ctx, organizationId).McpServerRequest(mcpServerRequest).Execute()

Create an MCP server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	mcpServerRequest := *openapiclient.NewMcpServerRequest("Name_example", "Url_example") // McpServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPServersAPI.CreateMcpServer(context.Background(), organizationId).McpServerRequest(mcpServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPServersAPI.CreateMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMcpServer`: McpServerResponse
	fmt.Fprintf(os.Stdout, "Response from `MCPServersAPI.CreateMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mcpServerRequest** | [**McpServerRequest**](McpServerRequest.md) |  | 

### Return type

[**McpServerResponse**](McpServerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMcpServer

> DeleteMcpServer(ctx, mcpServerId).Execute()

Delete an MCP server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	mcpServerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | MCP Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MCPServersAPI.DeleteMcpServer(context.Background(), mcpServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPServersAPI.DeleteMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mcpServerId** | **string** | MCP Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMcpServer

> McpServerResponse EditMcpServer(ctx, mcpServerId).McpServerRequest(mcpServerRequest).Execute()

Edit an MCP server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	mcpServerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | MCP Server ID
	mcpServerRequest := *openapiclient.NewMcpServerRequest("Name_example", "Url_example") // McpServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPServersAPI.EditMcpServer(context.Background(), mcpServerId).McpServerRequest(mcpServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPServersAPI.EditMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditMcpServer`: McpServerResponse
	fmt.Fprintf(os.Stdout, "Response from `MCPServersAPI.EditMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mcpServerId** | **string** | MCP Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mcpServerRequest** | [**McpServerRequest**](McpServerRequest.md) |  | 

### Return type

[**McpServerResponse**](McpServerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMcpServer

> McpServerResponse GetMcpServer(ctx, mcpServerId).Execute()

Get an MCP server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	mcpServerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | MCP Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPServersAPI.GetMcpServer(context.Background(), mcpServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPServersAPI.GetMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMcpServer`: McpServerResponse
	fmt.Fprintf(os.Stdout, "Response from `MCPServersAPI.GetMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mcpServerId** | **string** | MCP Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**McpServerResponse**](McpServerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMcpServers

> McpServerResponseList ListMcpServers(ctx, organizationId).Execute()

List organization MCP servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPServersAPI.ListMcpServers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPServersAPI.ListMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMcpServers`: McpServerResponseList
	fmt.Fprintf(os.Stdout, "Response from `MCPServersAPI.ListMcpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**McpServerResponseList**](McpServerResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

