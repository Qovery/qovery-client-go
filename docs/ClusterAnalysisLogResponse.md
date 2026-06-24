# ClusterAnalysisLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Level** | **string** |  | 
**Message** | **string** |  | 
**LineOrder** | **int32** |  | 

## Methods

### NewClusterAnalysisLogResponse

`func NewClusterAnalysisLogResponse(timestamp time.Time, level string, message string, lineOrder int32, ) *ClusterAnalysisLogResponse`

NewClusterAnalysisLogResponse instantiates a new ClusterAnalysisLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAnalysisLogResponseWithDefaults

`func NewClusterAnalysisLogResponseWithDefaults() *ClusterAnalysisLogResponse`

NewClusterAnalysisLogResponseWithDefaults instantiates a new ClusterAnalysisLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ClusterAnalysisLogResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClusterAnalysisLogResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClusterAnalysisLogResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetLevel

`func (o *ClusterAnalysisLogResponse) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ClusterAnalysisLogResponse) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ClusterAnalysisLogResponse) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *ClusterAnalysisLogResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterAnalysisLogResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterAnalysisLogResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLineOrder

`func (o *ClusterAnalysisLogResponse) GetLineOrder() int32`

GetLineOrder returns the LineOrder field if non-nil, zero value otherwise.

### GetLineOrderOk

`func (o *ClusterAnalysisLogResponse) GetLineOrderOk() (*int32, bool)`

GetLineOrderOk returns a tuple with the LineOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOrder

`func (o *ClusterAnalysisLogResponse) SetLineOrder(v int32)`

SetLineOrder sets LineOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


