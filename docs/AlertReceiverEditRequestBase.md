# AlertReceiverEditRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**AlertReceiverType**](AlertReceiverType.md) |  | 
**SendResolved** | **bool** |  | 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlertReceiverEditRequestBase

`func NewAlertReceiverEditRequestBase(name string, description string, type_ AlertReceiverType, sendResolved bool, ) *AlertReceiverEditRequestBase`

NewAlertReceiverEditRequestBase instantiates a new AlertReceiverEditRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReceiverEditRequestBaseWithDefaults

`func NewAlertReceiverEditRequestBaseWithDefaults() *AlertReceiverEditRequestBase`

NewAlertReceiverEditRequestBaseWithDefaults instantiates a new AlertReceiverEditRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertReceiverEditRequestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertReceiverEditRequestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertReceiverEditRequestBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertReceiverEditRequestBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertReceiverEditRequestBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertReceiverEditRequestBase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *AlertReceiverEditRequestBase) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertReceiverEditRequestBase) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertReceiverEditRequestBase) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *AlertReceiverEditRequestBase) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *AlertReceiverEditRequestBase) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *AlertReceiverEditRequestBase) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.


### GetOwner

`func (o *AlertReceiverEditRequestBase) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AlertReceiverEditRequestBase) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AlertReceiverEditRequestBase) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AlertReceiverEditRequestBase) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *AlertReceiverEditRequestBase) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *AlertReceiverEditRequestBase) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSeverity

`func (o *AlertReceiverEditRequestBase) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertReceiverEditRequestBase) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertReceiverEditRequestBase) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertReceiverEditRequestBase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *AlertReceiverEditRequestBase) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *AlertReceiverEditRequestBase) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


