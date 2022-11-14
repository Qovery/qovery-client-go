# EnvironmentLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Details** | [**EnvironmentLogsDetails**](EnvironmentLogsDetails.md) |  | 
**Error** | Pointer to [**NullableEnvironmentLogsError**](EnvironmentLogsError.md) |  | [optional] 
**Message** | Pointer to [**NullableEnvironmentLogsMessage**](EnvironmentLogsMessage.md) |  | [optional] 

## Methods

### NewEnvironmentLogs

`func NewEnvironmentLogs(type_ string, timestamp time.Time, details EnvironmentLogsDetails, ) *EnvironmentLogs`

NewEnvironmentLogs instantiates a new EnvironmentLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLogsWithDefaults

`func NewEnvironmentLogsWithDefaults() *EnvironmentLogs`

NewEnvironmentLogsWithDefaults instantiates a new EnvironmentLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnvironmentLogs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentLogs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentLogs) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *EnvironmentLogs) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnvironmentLogs) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnvironmentLogs) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetDetails

`func (o *EnvironmentLogs) GetDetails() EnvironmentLogsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EnvironmentLogs) GetDetailsOk() (*EnvironmentLogsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EnvironmentLogs) SetDetails(v EnvironmentLogsDetails)`

SetDetails sets Details field to given value.


### GetError

`func (o *EnvironmentLogs) GetError() EnvironmentLogsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EnvironmentLogs) GetErrorOk() (*EnvironmentLogsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EnvironmentLogs) SetError(v EnvironmentLogsError)`

SetError sets Error field to given value.

### HasError

`func (o *EnvironmentLogs) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EnvironmentLogs) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EnvironmentLogs) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetMessage

`func (o *EnvironmentLogs) GetMessage() EnvironmentLogsMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EnvironmentLogs) GetMessageOk() (*EnvironmentLogsMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EnvironmentLogs) SetMessage(v EnvironmentLogsMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EnvironmentLogs) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EnvironmentLogs) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EnvironmentLogs) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


