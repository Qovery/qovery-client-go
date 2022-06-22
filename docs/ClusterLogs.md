# ClusterLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | log level | [optional] 
**Timestamp** | Pointer to **time.Time** | log date creation | [optional] 
**Step** | Pointer to **string** | log step | [optional] 
**Message** | Pointer to [**ClusterLogsMessage**](ClusterLogsMessage.md) |  | [optional] 
**Error** | Pointer to [**ClusterLogsError**](ClusterLogsError.md) |  | [optional] 
**Details** | Pointer to [**ClusterLogsDetails**](ClusterLogsDetails.md) |  | [optional] 

## Methods

### NewClusterLogs

`func NewClusterLogs() *ClusterLogs`

NewClusterLogs instantiates a new ClusterLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsWithDefaults

`func NewClusterLogsWithDefaults() *ClusterLogs`

NewClusterLogsWithDefaults instantiates a new ClusterLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterLogs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterLogs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterLogs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterLogs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimestamp

`func (o *ClusterLogs) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClusterLogs) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClusterLogs) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ClusterLogs) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStep

`func (o *ClusterLogs) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *ClusterLogs) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *ClusterLogs) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *ClusterLogs) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterLogs) GetMessage() ClusterLogsMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterLogs) GetMessageOk() (*ClusterLogsMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterLogs) SetMessage(v ClusterLogsMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterLogs) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *ClusterLogs) GetError() ClusterLogsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClusterLogs) GetErrorOk() (*ClusterLogsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClusterLogs) SetError(v ClusterLogsError)`

SetError sets Error field to given value.

### HasError

`func (o *ClusterLogs) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetails

`func (o *ClusterLogs) GetDetails() ClusterLogsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ClusterLogs) GetDetailsOk() (*ClusterLogsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ClusterLogs) SetDetails(v ClusterLogsDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ClusterLogs) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


