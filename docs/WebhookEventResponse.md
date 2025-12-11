# WebhookEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier | 
**CreatedAt** | **time.Time** | Timestamp when the webhook event was created | 
**Kind** | [**OrganizationWebhookKindEnum**](OrganizationWebhookKindEnum.md) |  | 
**MatchedEvent** | [**OrganizationWebhookEventEnum**](OrganizationWebhookEventEnum.md) |  | 
**TargetUrlUsed** | **string** | The webhook target URL that was invoked | 
**Request** | **map[string]interface{}** | The request payload sent to the webhook | 
**TargetResponseStatusCode** | **int32** | HTTP status code returned by the webhook target | 
**TargetResponseBody** | Pointer to **NullableString** | Response body from the webhook target | [optional] 

## Methods

### NewWebhookEventResponse

`func NewWebhookEventResponse(id string, createdAt time.Time, kind OrganizationWebhookKindEnum, matchedEvent OrganizationWebhookEventEnum, targetUrlUsed string, request map[string]interface{}, targetResponseStatusCode int32, ) *WebhookEventResponse`

NewWebhookEventResponse instantiates a new WebhookEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventResponseWithDefaults

`func NewWebhookEventResponseWithDefaults() *WebhookEventResponse`

NewWebhookEventResponseWithDefaults instantiates a new WebhookEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *WebhookEventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetKind

`func (o *WebhookEventResponse) GetKind() OrganizationWebhookKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebhookEventResponse) GetKindOk() (*OrganizationWebhookKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebhookEventResponse) SetKind(v OrganizationWebhookKindEnum)`

SetKind sets Kind field to given value.


### GetMatchedEvent

`func (o *WebhookEventResponse) GetMatchedEvent() OrganizationWebhookEventEnum`

GetMatchedEvent returns the MatchedEvent field if non-nil, zero value otherwise.

### GetMatchedEventOk

`func (o *WebhookEventResponse) GetMatchedEventOk() (*OrganizationWebhookEventEnum, bool)`

GetMatchedEventOk returns a tuple with the MatchedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedEvent

`func (o *WebhookEventResponse) SetMatchedEvent(v OrganizationWebhookEventEnum)`

SetMatchedEvent sets MatchedEvent field to given value.


### GetTargetUrlUsed

`func (o *WebhookEventResponse) GetTargetUrlUsed() string`

GetTargetUrlUsed returns the TargetUrlUsed field if non-nil, zero value otherwise.

### GetTargetUrlUsedOk

`func (o *WebhookEventResponse) GetTargetUrlUsedOk() (*string, bool)`

GetTargetUrlUsedOk returns a tuple with the TargetUrlUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrlUsed

`func (o *WebhookEventResponse) SetTargetUrlUsed(v string)`

SetTargetUrlUsed sets TargetUrlUsed field to given value.


### GetRequest

`func (o *WebhookEventResponse) GetRequest() map[string]interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WebhookEventResponse) GetRequestOk() (*map[string]interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WebhookEventResponse) SetRequest(v map[string]interface{})`

SetRequest sets Request field to given value.


### GetTargetResponseStatusCode

`func (o *WebhookEventResponse) GetTargetResponseStatusCode() int32`

GetTargetResponseStatusCode returns the TargetResponseStatusCode field if non-nil, zero value otherwise.

### GetTargetResponseStatusCodeOk

`func (o *WebhookEventResponse) GetTargetResponseStatusCodeOk() (*int32, bool)`

GetTargetResponseStatusCodeOk returns a tuple with the TargetResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResponseStatusCode

`func (o *WebhookEventResponse) SetTargetResponseStatusCode(v int32)`

SetTargetResponseStatusCode sets TargetResponseStatusCode field to given value.


### GetTargetResponseBody

`func (o *WebhookEventResponse) GetTargetResponseBody() string`

GetTargetResponseBody returns the TargetResponseBody field if non-nil, zero value otherwise.

### GetTargetResponseBodyOk

`func (o *WebhookEventResponse) GetTargetResponseBodyOk() (*string, bool)`

GetTargetResponseBodyOk returns a tuple with the TargetResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResponseBody

`func (o *WebhookEventResponse) SetTargetResponseBody(v string)`

SetTargetResponseBody sets TargetResponseBody field to given value.

### HasTargetResponseBody

`func (o *WebhookEventResponse) HasTargetResponseBody() bool`

HasTargetResponseBody returns a boolean if a field has been set.

### SetTargetResponseBodyNil

`func (o *WebhookEventResponse) SetTargetResponseBodyNil(b bool)`

 SetTargetResponseBodyNil sets the value for TargetResponseBody to be an explicit nil

### UnsetTargetResponseBody
`func (o *WebhookEventResponse) UnsetTargetResponseBody()`

UnsetTargetResponseBody ensures that no value is present for TargetResponseBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


