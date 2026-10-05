# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Message** | **string** |  | 
**PodName** | Pointer to **string** |  | [optional] 
**ApplicationCommitId** | Pointer to **string** |  | [optional] 

## Methods

### NewLog

`func NewLog(id string, createdAt time.Time, message string, ) *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Log) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Log) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Log) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Log) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Log) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Log) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMessage

`func (o *Log) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Log) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Log) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPodName

`func (o *Log) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *Log) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *Log) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *Log) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetApplicationCommitId

`func (o *Log) GetApplicationCommitId() string`

GetApplicationCommitId returns the ApplicationCommitId field if non-nil, zero value otherwise.

### GetApplicationCommitIdOk

`func (o *Log) GetApplicationCommitIdOk() (*string, bool)`

GetApplicationCommitIdOk returns a tuple with the ApplicationCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCommitId

`func (o *Log) SetApplicationCommitId(v string)`

SetApplicationCommitId sets ApplicationCommitId field to given value.

### HasApplicationCommitId

`func (o *Log) HasApplicationCommitId() bool`

HasApplicationCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


