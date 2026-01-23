# GitWebhookStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The webhook configuration status: - ACTIVE: Webhook is properly configured with all required events - NOT_CONFIGURED: No Qovery webhook found on the git rep - MISCONFIGURED: Webhook exists but is missing required events - UNABLE_TO_VERIFY: Could not check webhook status (auth error, rate limit, etc.)  | 
**Provider** | **string** | The git provider where the webhook is configured | 
**MissingEvents** | Pointer to **[]string** | List of required events that are missing from the webhook configuration (only present when status is MISCONFIGURED) | [optional] 

## Methods

### NewGitWebhookStatusResponse

`func NewGitWebhookStatusResponse(status string, provider string, ) *GitWebhookStatusResponse`

NewGitWebhookStatusResponse instantiates a new GitWebhookStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitWebhookStatusResponseWithDefaults

`func NewGitWebhookStatusResponseWithDefaults() *GitWebhookStatusResponse`

NewGitWebhookStatusResponseWithDefaults instantiates a new GitWebhookStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GitWebhookStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GitWebhookStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GitWebhookStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *GitWebhookStatusResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitWebhookStatusResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitWebhookStatusResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMissingEvents

`func (o *GitWebhookStatusResponse) GetMissingEvents() []string`

GetMissingEvents returns the MissingEvents field if non-nil, zero value otherwise.

### GetMissingEventsOk

`func (o *GitWebhookStatusResponse) GetMissingEventsOk() (*[]string, bool)`

GetMissingEventsOk returns a tuple with the MissingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingEvents

`func (o *GitWebhookStatusResponse) SetMissingEvents(v []string)`

SetMissingEvents sets MissingEvents field to given value.

### HasMissingEvents

`func (o *GitWebhookStatusResponse) HasMissingEvents() bool`

HasMissingEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


