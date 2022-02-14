# \BackupsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBackupDatabase**](BackupsApi.md#AddBackupDatabase) | **Post** /database/{databaseId}/backup | Add a backup to the Database 
[**ListDatabaseBackup**](BackupsApi.md#ListDatabaseBackup) | **Get** /database/{databaseId}/backup | List database  backups
[**RemoveDatabaseBackup**](BackupsApi.md#RemoveDatabaseBackup) | **Delete** /database/{databaseId}/backup/{backupId} | Remove database  backup



## AddBackupDatabase

> BackupResponse AddBackupDatabase(ctx, databaseId).BackupRequest(backupRequest).Execute()

Add a backup to the Database 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    backupRequest := *openapiclient.NewBackupRequest("Name_example", "Message_example") // BackupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.AddBackupDatabase(context.Background(), databaseId).BackupRequest(backupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.AddBackupDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBackupDatabase`: BackupResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.AddBackupDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBackupDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupRequest** | [**BackupRequest**](BackupRequest.md) |  | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseBackup

> BackupPaginatedResponseList ListDatabaseBackup(ctx, databaseId).StartId(startId).Execute()

List database  backups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    startId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Starting point after which to return results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.ListDatabaseBackup(context.Background(), databaseId).StartId(startId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ListDatabaseBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseBackup`: BackupPaginatedResponseList
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ListDatabaseBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | **string** | Starting point after which to return results | 

### Return type

[**BackupPaginatedResponseList**](BackupPaginatedResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDatabaseBackup

> RemoveDatabaseBackup(ctx, databaseId, backupId).Execute()

Remove database  backup

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    backupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database Backup ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupsApi.RemoveDatabaseBackup(context.Background(), databaseId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.RemoveDatabaseBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 
**backupId** | **string** | Database Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDatabaseBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

