# \ReferralRewardsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountReferral**](ReferralRewardsApi.md#GetAccountReferral) | **Get** /account/referral | Get your referral information
[**PostAccountRewardClaim**](ReferralRewardsApi.md#PostAccountRewardClaim) | **Post** /account/rewardClaim | Claim a reward



## GetAccountReferral

> ReferralResponse GetAccountReferral(ctx).Execute()

Get your referral information

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReferralRewardsApi.GetAccountReferral(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferralRewardsApi.GetAccountReferral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountReferral`: ReferralResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferralRewardsApi.GetAccountReferral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountReferralRequest struct via the builder pattern


### Return type

[**ReferralResponse**](ReferralResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountRewardClaim

> PostAccountRewardClaim(ctx).RewardClaimResponse(rewardClaimResponse).Execute()

Claim a reward



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
    rewardClaimResponse := *openapiclient.NewRewardClaimResponse() // RewardClaimResponse |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReferralRewardsApi.PostAccountRewardClaim(context.Background()).RewardClaimResponse(rewardClaimResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferralRewardsApi.PostAccountRewardClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountRewardClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewardClaimResponse** | [**RewardClaimResponse**](RewardClaimResponse.md) |  | 

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

