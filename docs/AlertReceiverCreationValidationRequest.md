# AlertReceiverCreationValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertReceiver** | [**AlertReceiverCreationRequest**](AlertReceiverCreationRequest.md) |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertReceiverCreationValidationRequest

`func NewAlertReceiverCreationValidationRequest(alertReceiver AlertReceiverCreationRequest, ) *AlertReceiverCreationValidationRequest`

NewAlertReceiverCreationValidationRequest instantiates a new AlertReceiverCreationValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReceiverCreationValidationRequestWithDefaults

`func NewAlertReceiverCreationValidationRequestWithDefaults() *AlertReceiverCreationValidationRequest`

NewAlertReceiverCreationValidationRequestWithDefaults instantiates a new AlertReceiverCreationValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertReceiver

`func (o *AlertReceiverCreationValidationRequest) GetAlertReceiver() AlertReceiverCreationRequest`

GetAlertReceiver returns the AlertReceiver field if non-nil, zero value otherwise.

### GetAlertReceiverOk

`func (o *AlertReceiverCreationValidationRequest) GetAlertReceiverOk() (*AlertReceiverCreationRequest, bool)`

GetAlertReceiverOk returns a tuple with the AlertReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertReceiver

`func (o *AlertReceiverCreationValidationRequest) SetAlertReceiver(v AlertReceiverCreationRequest)`

SetAlertReceiver sets AlertReceiver field to given value.


### GetMessage

`func (o *AlertReceiverCreationValidationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertReceiverCreationValidationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertReceiverCreationValidationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertReceiverCreationValidationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


