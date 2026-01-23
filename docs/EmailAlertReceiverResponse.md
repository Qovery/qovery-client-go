# EmailAlertReceiverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**AlertReceiverType**](AlertReceiverType.md) |  | 
**SendResolved** | **bool** |  | 
**To** | **string** |  | 
**From** | **string** |  | 
**Smarthost** | **string** |  | 
**AuthUsername** | **NullableString** |  | 
**RequireTls** | **bool** |  | 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailAlertReceiverResponse

`func NewEmailAlertReceiverResponse(id string, createdAt time.Time, name string, description string, type_ AlertReceiverType, sendResolved bool, to string, from string, smarthost string, authUsername NullableString, requireTls bool, ) *EmailAlertReceiverResponse`

NewEmailAlertReceiverResponse instantiates a new EmailAlertReceiverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAlertReceiverResponseWithDefaults

`func NewEmailAlertReceiverResponseWithDefaults() *EmailAlertReceiverResponse`

NewEmailAlertReceiverResponseWithDefaults instantiates a new EmailAlertReceiverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailAlertReceiverResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailAlertReceiverResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailAlertReceiverResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EmailAlertReceiverResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailAlertReceiverResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailAlertReceiverResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EmailAlertReceiverResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmailAlertReceiverResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmailAlertReceiverResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EmailAlertReceiverResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *EmailAlertReceiverResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAlertReceiverResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAlertReceiverResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EmailAlertReceiverResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailAlertReceiverResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailAlertReceiverResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *EmailAlertReceiverResponse) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailAlertReceiverResponse) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailAlertReceiverResponse) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *EmailAlertReceiverResponse) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *EmailAlertReceiverResponse) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *EmailAlertReceiverResponse) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetTo

`func (o *EmailAlertReceiverResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailAlertReceiverResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailAlertReceiverResponse) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *EmailAlertReceiverResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailAlertReceiverResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailAlertReceiverResponse) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSmarthost

`func (o *EmailAlertReceiverResponse) GetSmarthost() string`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *EmailAlertReceiverResponse) GetSmarthostOk() (*string, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *EmailAlertReceiverResponse) SetSmarthost(v string)`

SetSmarthost sets Smarthost field to given value.


### GetAuthUsername

`func (o *EmailAlertReceiverResponse) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *EmailAlertReceiverResponse) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *EmailAlertReceiverResponse) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.


### SetAuthUsernameNil

`func (o *EmailAlertReceiverResponse) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *EmailAlertReceiverResponse) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetRequireTls

`func (o *EmailAlertReceiverResponse) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *EmailAlertReceiverResponse) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *EmailAlertReceiverResponse) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.


### GetOwner

`func (o *EmailAlertReceiverResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EmailAlertReceiverResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EmailAlertReceiverResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EmailAlertReceiverResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *EmailAlertReceiverResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *EmailAlertReceiverResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *EmailAlertReceiverResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EmailAlertReceiverResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EmailAlertReceiverResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EmailAlertReceiverResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *EmailAlertReceiverResponse) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *EmailAlertReceiverResponse) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


