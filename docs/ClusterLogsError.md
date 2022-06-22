# ClusterLogsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** | log error tag | [optional] 
**UserLogMessage** | Pointer to **string** | log details about the error | [optional] 
**Link** | Pointer to **string** | link to our documentation | [optional] 
**HintMessage** | Pointer to **string** | hint the user can follow | [optional] 
**EventDetails** | Pointer to [**ClusterLogsErrorEventDetails**](ClusterLogsErrorEventDetails.md) |  | [optional] 

## Methods

### NewClusterLogsError

`func NewClusterLogsError() *ClusterLogsError`

NewClusterLogsError instantiates a new ClusterLogsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsErrorWithDefaults

`func NewClusterLogsErrorWithDefaults() *ClusterLogsError`

NewClusterLogsErrorWithDefaults instantiates a new ClusterLogsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ClusterLogsError) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ClusterLogsError) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ClusterLogsError) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ClusterLogsError) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUserLogMessage

`func (o *ClusterLogsError) GetUserLogMessage() string`

GetUserLogMessage returns the UserLogMessage field if non-nil, zero value otherwise.

### GetUserLogMessageOk

`func (o *ClusterLogsError) GetUserLogMessageOk() (*string, bool)`

GetUserLogMessageOk returns a tuple with the UserLogMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLogMessage

`func (o *ClusterLogsError) SetUserLogMessage(v string)`

SetUserLogMessage sets UserLogMessage field to given value.

### HasUserLogMessage

`func (o *ClusterLogsError) HasUserLogMessage() bool`

HasUserLogMessage returns a boolean if a field has been set.

### GetLink

`func (o *ClusterLogsError) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ClusterLogsError) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ClusterLogsError) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ClusterLogsError) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetHintMessage

`func (o *ClusterLogsError) GetHintMessage() string`

GetHintMessage returns the HintMessage field if non-nil, zero value otherwise.

### GetHintMessageOk

`func (o *ClusterLogsError) GetHintMessageOk() (*string, bool)`

GetHintMessageOk returns a tuple with the HintMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHintMessage

`func (o *ClusterLogsError) SetHintMessage(v string)`

SetHintMessage sets HintMessage field to given value.

### HasHintMessage

`func (o *ClusterLogsError) HasHintMessage() bool`

HasHintMessage returns a boolean if a field has been set.

### GetEventDetails

`func (o *ClusterLogsError) GetEventDetails() ClusterLogsErrorEventDetails`

GetEventDetails returns the EventDetails field if non-nil, zero value otherwise.

### GetEventDetailsOk

`func (o *ClusterLogsError) GetEventDetailsOk() (*ClusterLogsErrorEventDetails, bool)`

GetEventDetailsOk returns a tuple with the EventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetails

`func (o *ClusterLogsError) SetEventDetails(v ClusterLogsErrorEventDetails)`

SetEventDetails sets EventDetails field to given value.

### HasEventDetails

`func (o *ClusterLogsError) HasEventDetails() bool`

HasEventDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


