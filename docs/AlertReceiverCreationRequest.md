# AlertReceiverCreationRequest

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
**To** | **string** | Recipient email address | 
**From** | **string** | Sender email address | 
**Smarthost** | **string** | SMTP server in format &#39;host:port&#39; | 
**AuthUsername** | Pointer to **NullableString** | SMTP authentication username. Defaults to &#39;from&#39; if not provided. | [optional] 
**AuthPassword** | **string** | SMTP authentication password | 
**RequireTls** | **bool** | Whether to require TLS for SMTP connection | 

## Methods

### NewAlertReceiverCreationRequest

`func NewAlertReceiverCreationRequest(organizationId string, name string, description string, type_ AlertReceiverType, sendResolved bool, webhookUrl string, to string, from string, smarthost string, authPassword string, requireTls bool, ) *AlertReceiverCreationRequest`

NewAlertReceiverCreationRequest instantiates a new AlertReceiverCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReceiverCreationRequestWithDefaults

`func NewAlertReceiverCreationRequestWithDefaults() *AlertReceiverCreationRequest`

NewAlertReceiverCreationRequestWithDefaults instantiates a new AlertReceiverCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *AlertReceiverCreationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertReceiverCreationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertReceiverCreationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *AlertReceiverCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertReceiverCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertReceiverCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertReceiverCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertReceiverCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertReceiverCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *AlertReceiverCreationRequest) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertReceiverCreationRequest) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertReceiverCreationRequest) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *AlertReceiverCreationRequest) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *AlertReceiverCreationRequest) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *AlertReceiverCreationRequest) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetOwner

`func (o *AlertReceiverCreationRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AlertReceiverCreationRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AlertReceiverCreationRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AlertReceiverCreationRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *AlertReceiverCreationRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *AlertReceiverCreationRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *AlertReceiverCreationRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertReceiverCreationRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertReceiverCreationRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertReceiverCreationRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *AlertReceiverCreationRequest) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *AlertReceiverCreationRequest) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetWebhookUrl

`func (o *AlertReceiverCreationRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AlertReceiverCreationRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AlertReceiverCreationRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetTo

`func (o *AlertReceiverCreationRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AlertReceiverCreationRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AlertReceiverCreationRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *AlertReceiverCreationRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AlertReceiverCreationRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AlertReceiverCreationRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSmarthost

`func (o *AlertReceiverCreationRequest) GetSmarthost() string`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *AlertReceiverCreationRequest) GetSmarthostOk() (*string, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *AlertReceiverCreationRequest) SetSmarthost(v string)`

SetSmarthost sets Smarthost field to given value.


### GetAuthUsername

`func (o *AlertReceiverCreationRequest) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *AlertReceiverCreationRequest) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *AlertReceiverCreationRequest) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *AlertReceiverCreationRequest) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *AlertReceiverCreationRequest) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *AlertReceiverCreationRequest) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAuthPassword

`func (o *AlertReceiverCreationRequest) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *AlertReceiverCreationRequest) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *AlertReceiverCreationRequest) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.


### GetRequireTls

`func (o *AlertReceiverCreationRequest) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *AlertReceiverCreationRequest) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *AlertReceiverCreationRequest) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


