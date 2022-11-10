# ListEnvironmentLogs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Details** | [**ListEnvironmentLogs200ResponseInnerDetails**](ListEnvironmentLogs200ResponseInnerDetails.md) |  | 
**Error** | Pointer to [**NullableListEnvironmentLogs200ResponseInnerError**](ListEnvironmentLogs200ResponseInnerError.md) |  | [optional] 
**Message** | Pointer to [**NullableListEnvironmentLogs200ResponseInnerMessage**](ListEnvironmentLogs200ResponseInnerMessage.md) |  | [optional] 

## Methods

### NewListEnvironmentLogs200ResponseInner

`func NewListEnvironmentLogs200ResponseInner(id string, timestamp time.Time, details ListEnvironmentLogs200ResponseInnerDetails, ) *ListEnvironmentLogs200ResponseInner`

NewListEnvironmentLogs200ResponseInner instantiates a new ListEnvironmentLogs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironmentLogs200ResponseInnerWithDefaults

`func NewListEnvironmentLogs200ResponseInnerWithDefaults() *ListEnvironmentLogs200ResponseInner`

NewListEnvironmentLogs200ResponseInnerWithDefaults instantiates a new ListEnvironmentLogs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListEnvironmentLogs200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListEnvironmentLogs200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListEnvironmentLogs200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *ListEnvironmentLogs200ResponseInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListEnvironmentLogs200ResponseInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListEnvironmentLogs200ResponseInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetDetails

`func (o *ListEnvironmentLogs200ResponseInner) GetDetails() ListEnvironmentLogs200ResponseInnerDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListEnvironmentLogs200ResponseInner) GetDetailsOk() (*ListEnvironmentLogs200ResponseInnerDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListEnvironmentLogs200ResponseInner) SetDetails(v ListEnvironmentLogs200ResponseInnerDetails)`

SetDetails sets Details field to given value.


### GetError

`func (o *ListEnvironmentLogs200ResponseInner) GetError() ListEnvironmentLogs200ResponseInnerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListEnvironmentLogs200ResponseInner) GetErrorOk() (*ListEnvironmentLogs200ResponseInnerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListEnvironmentLogs200ResponseInner) SetError(v ListEnvironmentLogs200ResponseInnerError)`

SetError sets Error field to given value.

### HasError

`func (o *ListEnvironmentLogs200ResponseInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ListEnvironmentLogs200ResponseInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListEnvironmentLogs200ResponseInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetMessage

`func (o *ListEnvironmentLogs200ResponseInner) GetMessage() ListEnvironmentLogs200ResponseInnerMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListEnvironmentLogs200ResponseInner) GetMessageOk() (*ListEnvironmentLogs200ResponseInnerMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListEnvironmentLogs200ResponseInner) SetMessage(v ListEnvironmentLogs200ResponseInnerMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListEnvironmentLogs200ResponseInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ListEnvironmentLogs200ResponseInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ListEnvironmentLogs200ResponseInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


