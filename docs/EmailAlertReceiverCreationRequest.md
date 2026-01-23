# EmailAlertReceiverCreationRequest

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
**To** | **string** | Recipient email address | 
**From** | **string** | Sender email address | 
**Smarthost** | **string** | SMTP server in format &#39;host:port&#39; | 
**AuthUsername** | Pointer to **NullableString** | SMTP authentication username. Defaults to &#39;from&#39; if not provided. | [optional] 
**AuthPassword** | **string** | SMTP authentication password | 
**RequireTls** | **bool** | Whether to require TLS for SMTP connection | 

## Methods

### NewEmailAlertReceiverCreationRequest

`func NewEmailAlertReceiverCreationRequest(organizationId string, name string, description string, type_ AlertReceiverType, sendResolved bool, to string, from string, smarthost string, authPassword string, requireTls bool, ) *EmailAlertReceiverCreationRequest`

NewEmailAlertReceiverCreationRequest instantiates a new EmailAlertReceiverCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAlertReceiverCreationRequestWithDefaults

`func NewEmailAlertReceiverCreationRequestWithDefaults() *EmailAlertReceiverCreationRequest`

NewEmailAlertReceiverCreationRequestWithDefaults instantiates a new EmailAlertReceiverCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *EmailAlertReceiverCreationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *EmailAlertReceiverCreationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *EmailAlertReceiverCreationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *EmailAlertReceiverCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAlertReceiverCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAlertReceiverCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EmailAlertReceiverCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailAlertReceiverCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailAlertReceiverCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *EmailAlertReceiverCreationRequest) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailAlertReceiverCreationRequest) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailAlertReceiverCreationRequest) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *EmailAlertReceiverCreationRequest) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *EmailAlertReceiverCreationRequest) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *EmailAlertReceiverCreationRequest) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetOwner

`func (o *EmailAlertReceiverCreationRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EmailAlertReceiverCreationRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EmailAlertReceiverCreationRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EmailAlertReceiverCreationRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *EmailAlertReceiverCreationRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *EmailAlertReceiverCreationRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *EmailAlertReceiverCreationRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EmailAlertReceiverCreationRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EmailAlertReceiverCreationRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EmailAlertReceiverCreationRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *EmailAlertReceiverCreationRequest) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *EmailAlertReceiverCreationRequest) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetTo

`func (o *EmailAlertReceiverCreationRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailAlertReceiverCreationRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailAlertReceiverCreationRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *EmailAlertReceiverCreationRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailAlertReceiverCreationRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailAlertReceiverCreationRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSmarthost

`func (o *EmailAlertReceiverCreationRequest) GetSmarthost() string`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *EmailAlertReceiverCreationRequest) GetSmarthostOk() (*string, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *EmailAlertReceiverCreationRequest) SetSmarthost(v string)`

SetSmarthost sets Smarthost field to given value.


### GetAuthUsername

`func (o *EmailAlertReceiverCreationRequest) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *EmailAlertReceiverCreationRequest) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *EmailAlertReceiverCreationRequest) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *EmailAlertReceiverCreationRequest) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *EmailAlertReceiverCreationRequest) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *EmailAlertReceiverCreationRequest) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAuthPassword

`func (o *EmailAlertReceiverCreationRequest) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *EmailAlertReceiverCreationRequest) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *EmailAlertReceiverCreationRequest) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.


### GetRequireTls

`func (o *EmailAlertReceiverCreationRequest) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *EmailAlertReceiverCreationRequest) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *EmailAlertReceiverCreationRequest) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


