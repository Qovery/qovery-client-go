# SlackAlertReceiverCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**AlertReceiverType**](AlertReceiverType.md) |  | 
**SendResolved** | **bool** |  | 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | **string** |  | 

## Methods

### NewSlackAlertReceiverCreationRequest

`func NewSlackAlertReceiverCreationRequest(organizationId string, name string, description string, type_ AlertReceiverType, sendResolved bool, webhookUrl string, ) *SlackAlertReceiverCreationRequest`

NewSlackAlertReceiverCreationRequest instantiates a new SlackAlertReceiverCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackAlertReceiverCreationRequestWithDefaults

`func NewSlackAlertReceiverCreationRequestWithDefaults() *SlackAlertReceiverCreationRequest`

NewSlackAlertReceiverCreationRequestWithDefaults instantiates a new SlackAlertReceiverCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *SlackAlertReceiverCreationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SlackAlertReceiverCreationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SlackAlertReceiverCreationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *SlackAlertReceiverCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlackAlertReceiverCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlackAlertReceiverCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SlackAlertReceiverCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlackAlertReceiverCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlackAlertReceiverCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *SlackAlertReceiverCreationRequest) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlackAlertReceiverCreationRequest) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlackAlertReceiverCreationRequest) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *SlackAlertReceiverCreationRequest) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *SlackAlertReceiverCreationRequest) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *SlackAlertReceiverCreationRequest) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetOwner

`func (o *SlackAlertReceiverCreationRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SlackAlertReceiverCreationRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SlackAlertReceiverCreationRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SlackAlertReceiverCreationRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *SlackAlertReceiverCreationRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *SlackAlertReceiverCreationRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *SlackAlertReceiverCreationRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SlackAlertReceiverCreationRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SlackAlertReceiverCreationRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SlackAlertReceiverCreationRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *SlackAlertReceiverCreationRequest) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *SlackAlertReceiverCreationRequest) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetWebhookUrl

`func (o *SlackAlertReceiverCreationRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *SlackAlertReceiverCreationRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *SlackAlertReceiverCreationRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


