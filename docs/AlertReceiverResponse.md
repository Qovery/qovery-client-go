# AlertReceiverResponse

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

## Methods

### NewAlertReceiverResponse

`func NewAlertReceiverResponse(id string, createdAt time.Time, name string, description string, type_ AlertReceiverType, sendResolved bool, ) *AlertReceiverResponse`

NewAlertReceiverResponse instantiates a new AlertReceiverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReceiverResponseWithDefaults

`func NewAlertReceiverResponseWithDefaults() *AlertReceiverResponse`

NewAlertReceiverResponseWithDefaults instantiates a new AlertReceiverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertReceiverResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertReceiverResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertReceiverResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AlertReceiverResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertReceiverResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertReceiverResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AlertReceiverResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AlertReceiverResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AlertReceiverResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AlertReceiverResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *AlertReceiverResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertReceiverResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertReceiverResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertReceiverResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertReceiverResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertReceiverResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *AlertReceiverResponse) GetType() AlertReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertReceiverResponse) GetTypeOk() (*AlertReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertReceiverResponse) SetType(v AlertReceiverType)`

SetType sets Type field to given value.


### GetSendResolved

`func (o *AlertReceiverResponse) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *AlertReceiverResponse) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *AlertReceiverResponse) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


