# EmailAlertReceiverEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**AlertReceiverType**](AlertReceiverType.md) |  | 
**SendResolved** | **bool** |  | 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**To** | **string** |  | 
**From** | **string** |  | 
**Smarthost** | **string** | SMTP server in format &#39;host:port&#39; | 
**AuthUsername** | Pointer to **NullableString** |  | [optional] 
**AuthPassword** | Pointer to **NullableString** | SMTP password. If null, keeps existing password. | [optional] 
**RequireTls** | **bool** |  | 

## Methods

### NewEmailAlertReceiverEditRequest

`func NewEmailAlertReceiverEditRequest(name string, description string, type_ AlertReceiverType, sendResolved bool, to string, from string, smarthost string, requireTls bool, ) *EmailAlertReceiverEditRequest`

NewEmailAlertReceiverEditRequest instantiates a new EmailAlertReceiverEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAlertReceiverEditRequestWithDefaults

`func NewEmailAlertReceiverEditRequestWithDefaults() *EmailAlertReceiverEditRequest`

NewEmailAlertReceiverEditRequestWithDefaults instantiates a new EmailAlertReceiverEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailAlertReceiverEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAlertReceiverEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAlertReceiverEditRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EmailAlertReceiverEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailAlertReceiverEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailAlertReceiverEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *EmailAlertReceiverEditRequest) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailAlertReceiverEditRequest) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailAlertReceiverEditRequest) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *EmailAlertReceiverEditRequest) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *EmailAlertReceiverEditRequest) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *EmailAlertReceiverEditRequest) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetOwner

`func (o *EmailAlertReceiverEditRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EmailAlertReceiverEditRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EmailAlertReceiverEditRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EmailAlertReceiverEditRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *EmailAlertReceiverEditRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *EmailAlertReceiverEditRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *EmailAlertReceiverEditRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EmailAlertReceiverEditRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EmailAlertReceiverEditRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EmailAlertReceiverEditRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *EmailAlertReceiverEditRequest) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *EmailAlertReceiverEditRequest) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetTo

`func (o *EmailAlertReceiverEditRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailAlertReceiverEditRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailAlertReceiverEditRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *EmailAlertReceiverEditRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailAlertReceiverEditRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailAlertReceiverEditRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSmarthost

`func (o *EmailAlertReceiverEditRequest) GetSmarthost() string`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *EmailAlertReceiverEditRequest) GetSmarthostOk() (*string, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *EmailAlertReceiverEditRequest) SetSmarthost(v string)`

SetSmarthost sets Smarthost field to given value.


### GetAuthUsername

`func (o *EmailAlertReceiverEditRequest) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *EmailAlertReceiverEditRequest) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *EmailAlertReceiverEditRequest) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *EmailAlertReceiverEditRequest) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *EmailAlertReceiverEditRequest) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *EmailAlertReceiverEditRequest) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAuthPassword

`func (o *EmailAlertReceiverEditRequest) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *EmailAlertReceiverEditRequest) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *EmailAlertReceiverEditRequest) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *EmailAlertReceiverEditRequest) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### SetAuthPasswordNil

`func (o *EmailAlertReceiverEditRequest) SetAuthPasswordNil(b bool)`

 SetAuthPasswordNil sets the value for AuthPassword to be an explicit nil

### UnsetAuthPassword
`func (o *EmailAlertReceiverEditRequest) UnsetAuthPassword()`

UnsetAuthPassword ensures that no value is present for AuthPassword, not even an explicit nil
### GetRequireTls

`func (o *EmailAlertReceiverEditRequest) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *EmailAlertReceiverEditRequest) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *EmailAlertReceiverEditRequest) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


