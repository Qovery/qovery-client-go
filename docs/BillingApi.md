# \BillingApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCreditCard**](BillingApi.md#AddCreditCard) | **Post** /organization/{organizationId}/creditCard | Add credit card
[**AddCreditCode**](BillingApi.md#AddCreditCode) | **Post** /organization/{organizationId}/creditCode | Add credit code
[**EditOrganizationBillingInfo**](BillingApi.md#EditOrganizationBillingInfo) | **Put** /organization/{organizationId}/billingInfo | Edit Organization Billing Info
[**EditOrganizationBudget**](BillingApi.md#EditOrganizationBudget) | **Put** /organization/{organizationId}/costBudget | Edit Organization Budget
[**GetClusterCurrentCost**](BillingApi.md#GetClusterCurrentCost) | **Get** /organization/{organizationId}/cluster/{clusterId}/currentCost | Get cluster current cost
[**GetOrganizationBillingInfo**](BillingApi.md#GetOrganizationBillingInfo) | **Get** /organization/{organizationId}/billingInfo | Get organization billing info
[**GetOrganizationBillingStatus**](BillingApi.md#GetOrganizationBillingStatus) | **Get** /organization/{organizationId}/billingStatus | Get organization billing status
[**GetOrganizationCostBudget**](BillingApi.md#GetOrganizationCostBudget) | **Get** /organization/{organizationId}/costBudget | Get organization cost Budget
[**GetOrganizationCurrentCost**](BillingApi.md#GetOrganizationCurrentCost) | **Get** /organization/{organizationId}/currentCost | Get organization current cost
[**GetOrganizationInvoice**](BillingApi.md#GetOrganizationInvoice) | **Get** /organization/{organizationId}/invoice/{invoiceId} | Get organization invoice
[**GetOrganizationInvoicePDF**](BillingApi.md#GetOrganizationInvoicePDF) | **Get** /organization/{organizationId}/invoice/{invoiceId}/download | Get invoice link
[**ListOrganizationCreditCards**](BillingApi.md#ListOrganizationCreditCards) | **Get** /organization/{organizationId}/creditCard | List organization credit cards
[**ListOrganizationInvoice**](BillingApi.md#ListOrganizationInvoice) | **Get** /organization/{organizationId}/invoice | List organization invoices
[**OrganizationDownloadAllInvoices**](BillingApi.md#OrganizationDownloadAllInvoices) | **Post** /organization/{organizationId}/downloadInvoices | Download all invoices
[**OrganizationOrganizationIdCreditCardCreditCardIdDelete**](BillingApi.md#OrganizationOrganizationIdCreditCardCreditCardIdDelete) | **Delete** /organization/{organizationId}/creditCard/{creditCardId} | Delete credit card



## AddCreditCard

> CreditCardResponse AddCreditCard(ctx, organizationId).CreditCardRequest(creditCardRequest).Execute()

Add credit card

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
    organizationId := TODO // string | Organization ID
    creditCardRequest := *openapiclient.NewCreditCardRequest("Number_example", "Cvv_example", int32(6), int32(2025)) // CreditCardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.AddCreditCard(context.Background(), organizationId).CreditCardRequest(creditCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.AddCreditCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCreditCard`: CreditCardResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.AddCreditCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCreditCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditCardRequest** | [**CreditCardRequest**](CreditCardRequest.md) |  | 

### Return type

[**CreditCardResponse**](CreditCardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCreditCode

> AddCreditCode(ctx, organizationId).OrganizationCreditCodeRequest(organizationCreditCodeRequest).Execute()

Add credit code

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
    organizationId := TODO // string | Organization ID
    organizationCreditCodeRequest := *openapiclient.NewOrganizationCreditCodeRequest() // OrganizationCreditCodeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.AddCreditCode(context.Background(), organizationId).OrganizationCreditCodeRequest(organizationCreditCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.AddCreditCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCreditCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationCreditCodeRequest** | [**OrganizationCreditCodeRequest**](OrganizationCreditCodeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationBillingInfo

> BillingInfoResponse EditOrganizationBillingInfo(ctx, organizationId).BillingInfoRequest(billingInfoRequest).Execute()

Edit Organization Billing Info

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
    organizationId := TODO // string | Organization ID
    billingInfoRequest := *openapiclient.NewBillingInfoRequest("Forrest", "Gump", "forrest@gump.com", "21 Jenny Street", "Greenbow", "36744", "US") // BillingInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.EditOrganizationBillingInfo(context.Background(), organizationId).BillingInfoRequest(billingInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.EditOrganizationBillingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationBillingInfo`: BillingInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.EditOrganizationBillingInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationBillingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingInfoRequest** | [**BillingInfoRequest**](BillingInfoRequest.md) |  | 

### Return type

[**BillingInfoResponse**](BillingInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationBudget

> CostResponse EditOrganizationBudget(ctx, organizationId).OrganizationBudgetEditRequest(organizationBudgetEditRequest).Execute()

Edit Organization Budget

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
    organizationId := TODO // string | Organization ID
    organizationBudgetEditRequest := *openapiclient.NewOrganizationBudgetEditRequest() // OrganizationBudgetEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.EditOrganizationBudget(context.Background(), organizationId).OrganizationBudgetEditRequest(organizationBudgetEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.EditOrganizationBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationBudget`: CostResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.EditOrganizationBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationBudgetEditRequest** | [**OrganizationBudgetEditRequest**](OrganizationBudgetEditRequest.md) |  | 

### Return type

[**CostResponse**](CostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterCurrentCost

> CostResponse GetClusterCurrentCost(ctx, organizationId, clusterId).Execute()

Get cluster current cost

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetClusterCurrentCost(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetClusterCurrentCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterCurrentCost`: CostResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetClusterCurrentCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterCurrentCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CostResponse**](CostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBillingInfo

> BillingInfoResponse GetOrganizationBillingInfo(ctx, organizationId).Execute()

Get organization billing info

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetOrganizationBillingInfo(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetOrganizationBillingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBillingInfo`: BillingInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetOrganizationBillingInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBillingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingInfoResponse**](BillingInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBillingStatus

> BillingStatus GetOrganizationBillingStatus(ctx, organizationId).Execute()

Get organization billing status



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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetOrganizationBillingStatus(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetOrganizationBillingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBillingStatus`: BillingStatus
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetOrganizationBillingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBillingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingStatus**](BillingStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCostBudget

> CostResponse GetOrganizationCostBudget(ctx, organizationId).Execute()

Get organization cost Budget

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetOrganizationCostBudget(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetOrganizationCostBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCostBudget`: CostResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetOrganizationCostBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCostBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CostResponse**](CostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCurrentCost

> OrganizationCurrentCostResponse GetOrganizationCurrentCost(ctx, organizationId).Execute()

Get organization current cost

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetOrganizationCurrentCost(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetOrganizationCurrentCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCurrentCost`: OrganizationCurrentCostResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetOrganizationCurrentCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCurrentCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationCurrentCostResponse**](OrganizationCurrentCostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInvoice

> InvoiceResponse GetOrganizationInvoice(ctx, organizationId, invoiceId).Execute()

Get organization invoice

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
    organizationId := TODO // string | Organization ID
    invoiceId := TODO // string | Invoice ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetOrganizationInvoice(context.Background(), organizationId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetOrganizationInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInvoice`: InvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetOrganizationInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**invoiceId** | [**string**](.md) | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInvoicePDF

> LinkResponse GetOrganizationInvoicePDF(ctx, organizationId, invoiceId).Execute()

Get invoice link



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
    organizationId := TODO // string | Organization ID
    invoiceId := TODO // string | Invoice ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetOrganizationInvoicePDF(context.Background(), organizationId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetOrganizationInvoicePDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInvoicePDF`: LinkResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetOrganizationInvoicePDF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**invoiceId** | [**string**](.md) | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInvoicePDFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LinkResponse**](LinkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCreditCards

> CreditCardResponseList ListOrganizationCreditCards(ctx, organizationId).Execute()

List organization credit cards

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.ListOrganizationCreditCards(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListOrganizationCreditCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationCreditCards`: CreditCardResponseList
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListOrganizationCreditCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationCreditCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditCardResponseList**](CreditCardResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationInvoice

> InvoiceResponseList ListOrganizationInvoice(ctx, organizationId).Execute()

List organization invoices

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.ListOrganizationInvoice(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListOrganizationInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationInvoice`: InvoiceResponseList
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListOrganizationInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceResponseList**](InvoiceResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationDownloadAllInvoices

> OrganizationDownloadAllInvoices(ctx, organizationId).Execute()

Download all invoices

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.OrganizationDownloadAllInvoices(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.OrganizationDownloadAllInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationDownloadAllInvoicesRequest struct via the builder pattern


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


## OrganizationOrganizationIdCreditCardCreditCardIdDelete

> OrganizationOrganizationIdCreditCardCreditCardIdDelete(ctx, organizationId, creditCardId).Execute()

Delete credit card

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
    organizationId := TODO // string | Organization ID
    creditCardId := TODO // string | Credit Card ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.OrganizationOrganizationIdCreditCardCreditCardIdDelete(context.Background(), organizationId, creditCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.OrganizationOrganizationIdCreditCardCreditCardIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**creditCardId** | [**string**](.md) | Credit Card ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationIdCreditCardCreditCardIdDeleteRequest struct via the builder pattern


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

