# AlertReceiverCreationRequestBase

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

## Methods

### NewAlertReceiverCreationRequestBase

`func NewAlertReceiverCreationRequestBase(organizationId string, name string, description string, type_ AlertReceiverType, sendResolved bool, ) *AlertReceiverCreationRequestBase`

NewAlertReceiverCreationRequestBase instantiates a new AlertReceiverCreationRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReceiverCreationRequestBaseWithDefaults

`func NewAlertReceiverCreationRequestBaseWithDefaults() *AlertReceiverCreationRequestBase`

NewAlertReceiverCreationRequestBaseWithDefaults instantiates a new AlertReceiverCreationRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *AlertReceiverCreationRequestBase) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertReceiverCreationRequestBase) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertReceiverCreationRequestBase) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *AlertReceiverCreationRequestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertReceiverCreationRequestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertReceiverCreationRequestBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertReceiverCreationRequestBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertReceiverCreationRequestBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertReceiverCreationRequestBase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *AlertReceiverCreationRequestBase) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertReceiverCreationRequestBase) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertReceiverCreationRequestBase) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *AlertReceiverCreationRequestBase) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *AlertReceiverCreationRequestBase) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *AlertReceiverCreationRequestBase) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetOwner

`func (o *AlertReceiverCreationRequestBase) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AlertReceiverCreationRequestBase) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AlertReceiverCreationRequestBase) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AlertReceiverCreationRequestBase) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *AlertReceiverCreationRequestBase) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *AlertReceiverCreationRequestBase) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *AlertReceiverCreationRequestBase) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertReceiverCreationRequestBase) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertReceiverCreationRequestBase) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertReceiverCreationRequestBase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *AlertReceiverCreationRequestBase) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *AlertReceiverCreationRequestBase) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


