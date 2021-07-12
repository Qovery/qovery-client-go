# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Commit** | Pointer to [**CommitResponse**](CommitResponse.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Type** | Pointer to **string** | DRAFT - we have to specify here all the possible events | [optional] 
**Log** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEventResponse

`func NewEventResponse(id string, createdAt time.Time, ) *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *EventResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *EventResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCommit

`func (o *EventResponse) GetCommit() CommitResponse`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *EventResponse) GetCommitOk() (*CommitResponse, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *EventResponse) SetCommit(v CommitResponse)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *EventResponse) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetStatus

`func (o *EventResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EventResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLog

`func (o *EventResponse) GetLog() ReferenceObject`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *EventResponse) GetLogOk() (*ReferenceObject, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *EventResponse) SetLog(v ReferenceObject)`

SetLog sets Log field to given value.

### HasLog

`func (o *EventResponse) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetId

`func (o *EventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EventResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EventResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


