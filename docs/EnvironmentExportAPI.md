# \EnvironmentExportAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportEnvironmentConfigurationIntoTerraform**](EnvironmentExportAPI.md#ExportEnvironmentConfigurationIntoTerraform) | **Get** /environment/{environmentId}/terraformExport | Export full environment and its resources into Terraform manifests



## ExportEnvironmentConfigurationIntoTerraform

> *os.File ExportEnvironmentConfigurationIntoTerraform(ctx, environmentId).ExportSecrets(exportSecrets).Execute()

Export full environment and its resources into Terraform manifests

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    exportSecrets := true // bool | export Secrets from configuration and include them into Terraform export (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentExportAPI.ExportEnvironmentConfigurationIntoTerraform(context.Background(), environmentId).ExportSecrets(exportSecrets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentExportAPI.ExportEnvironmentConfigurationIntoTerraform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportEnvironmentConfigurationIntoTerraform`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentExportAPI.ExportEnvironmentConfigurationIntoTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportEnvironmentConfigurationIntoTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportSecrets** | **bool** | export Secrets from configuration and include them into Terraform export | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

