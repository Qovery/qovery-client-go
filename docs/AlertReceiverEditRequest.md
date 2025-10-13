# AlertReceiverEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | [**AlertReceiverType**](AlertReceiverType.md) |  | 
**SendResolved** | **bool** |  | 

## Methods

### NewAlertReceiverEditRequest

`func NewAlertReceiverEditRequest(name string, description string, type_ AlertReceiverType, sendResolved bool, ) *AlertReceiverEditRequest`

NewAlertReceiverEditRequest instantiates a new AlertReceiverEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReceiverEditRequestWithDefaults

`func NewAlertReceiverEditRequestWithDefaults() *AlertReceiverEditRequest`

NewAlertReceiverEditRequestWithDefaults instantiates a new AlertReceiverEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertReceiverEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertReceiverEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertReceiverEditRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertReceiverEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertReceiverEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertReceiverEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *AlertReceiverEditRequest) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertReceiverEditRequest) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertReceiverEditRequest) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *AlertReceiverEditRequest) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *AlertReceiverEditRequest) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *AlertReceiverEditRequest) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


