# \MembersAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInviteMember**](MembersAPI.md#DeleteInviteMember) | **Delete** /organization/{organizationId}/inviteMember/{inviteId} | Remove an invited member
[**DeleteMember**](MembersAPI.md#DeleteMember) | **Delete** /organization/{organizationId}/member | Remove a member
[**EditOrganizationMemberRole**](MembersAPI.md#EditOrganizationMemberRole) | **Put** /organization/{organizationId}/member | Edit an organization member role
[**GetMemberInvitation**](MembersAPI.md#GetMemberInvitation) | **Get** /organization/{organizationId}/inviteMember/{inviteId} | Get member invitation
[**GetOrganizationInvitedMembers**](MembersAPI.md#GetOrganizationInvitedMembers) | **Get** /organization/{organizationId}/inviteMember | Get invited members
[**GetOrganizationMembers**](MembersAPI.md#GetOrganizationMembers) | **Get** /organization/{organizationId}/member | Get organization members
[**PostAcceptInviteMember**](MembersAPI.md#PostAcceptInviteMember) | **Post** /organization/{organizationId}/inviteMember/{inviteId} | Accept Invite in the organization
[**PostInviteMember**](MembersAPI.md#PostInviteMember) | **Post** /organization/{organizationId}/inviteMember | Invite someone in the organization
[**PostOrganizationTransferOwnership**](MembersAPI.md#PostOrganizationTransferOwnership) | **Post** /organization/{organizationId}/transferOwnership | Transfer organization ownership to another user



## DeleteInviteMember

> DeleteInviteMember(ctx, organizationId, inviteId).Execute()

Remove an invited member

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
    inviteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invite ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersAPI.DeleteInviteMember(context.Background(), organizationId, inviteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteInviteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**inviteId** | **string** | Invite ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInviteMemberRequest struct via the builder pattern


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


## DeleteMember

> DeleteMember(ctx, organizationId).DeleteMemberRequest(deleteMemberRequest).Execute()

Remove a member

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
    deleteMemberRequest := *openapiclient.NewDeleteMemberRequest("UserId_example") // DeleteMemberRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersAPI.DeleteMember(context.Background(), organizationId).DeleteMemberRequest(deleteMemberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteMemberRequest** | [**DeleteMemberRequest**](DeleteMemberRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationMemberRole

> EditOrganizationMemberRole(ctx, organizationId).MemberRoleUpdateRequest(memberRoleUpdateRequest).Execute()

Edit an organization member role



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
    memberRoleUpdateRequest := *openapiclient.NewMemberRoleUpdateRequest("UserId_example", "RoleId_example") // MemberRoleUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersAPI.EditOrganizationMemberRole(context.Background(), organizationId).MemberRoleUpdateRequest(memberRoleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.EditOrganizationMemberRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberRoleUpdateRequest** | [**MemberRoleUpdateRequest**](MemberRoleUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemberInvitation

> InviteMember GetMemberInvitation(ctx, organizationId, inviteId).Execute()

Get member invitation

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
    inviteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invite ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.GetMemberInvitation(context.Background(), organizationId, inviteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetMemberInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemberInvitation`: InviteMember
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetMemberInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**inviteId** | **string** | Invite ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InviteMember**](InviteMember.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInvitedMembers

> InviteMemberResponseList GetOrganizationInvitedMembers(ctx, organizationId).Execute()

Get invited members

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
    resp, r, err := apiClient.MembersAPI.GetOrganizationInvitedMembers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetOrganizationInvitedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInvitedMembers`: InviteMemberResponseList
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetOrganizationInvitedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInvitedMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InviteMemberResponseList**](InviteMemberResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationMembers

> MemberResponseList GetOrganizationMembers(ctx, organizationId).Execute()

Get organization members

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
    resp, r, err := apiClient.MembersAPI.GetOrganizationMembers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetOrganizationMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationMembers`: MemberResponseList
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MemberResponseList**](MemberResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAcceptInviteMember

> InviteMember PostAcceptInviteMember(ctx, organizationId, inviteId).Execute()

Accept Invite in the organization

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
    inviteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invite ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.PostAcceptInviteMember(context.Background(), organizationId, inviteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.PostAcceptInviteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAcceptInviteMember`: InviteMember
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.PostAcceptInviteMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**inviteId** | **string** | Invite ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAcceptInviteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InviteMember**](InviteMember.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInviteMember

> InviteMember PostInviteMember(ctx, organizationId).InviteMemberRequest(inviteMemberRequest).Execute()

Invite someone in the organization

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
    inviteMemberRequest := *openapiclient.NewInviteMemberRequest("Email_example") // InviteMemberRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.PostInviteMember(context.Background(), organizationId).InviteMemberRequest(inviteMemberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.PostInviteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostInviteMember`: InviteMember
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.PostInviteMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInviteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteMemberRequest** | [**InviteMemberRequest**](InviteMemberRequest.md) |  | 

### Return type

[**InviteMember**](InviteMember.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationTransferOwnership

> PostOrganizationTransferOwnership(ctx, organizationId).TransferOwnershipRequest(transferOwnershipRequest).Execute()

Transfer organization ownership to another user

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
    transferOwnershipRequest := *openapiclient.NewTransferOwnershipRequest("UserId_example") // TransferOwnershipRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersAPI.PostOrganizationTransferOwnership(context.Background(), organizationId).TransferOwnershipRequest(transferOwnershipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.PostOrganizationTransferOwnership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationTransferOwnershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferOwnershipRequest** | [**TransferOwnershipRequest**](TransferOwnershipRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

