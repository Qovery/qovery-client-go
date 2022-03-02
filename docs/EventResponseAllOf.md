# EventResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Commit** | Pointer to [**CommitResponse**](CommitResponse.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Type** | Pointer to **string** | DRAFT - we have to specify here all the possible events | [optional] 
**Log** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 

## Methods

### NewEventResponseAllOf

`func NewEventResponseAllOf() *EventResponseAllOf`

NewEventResponseAllOf instantiates a new EventResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseAllOfWithDefaults

`func NewEventResponseAllOfWithDefaults() *EventResponseAllOf`

NewEventResponseAllOfWithDefaults instantiates a new EventResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *EventResponseAllOf) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventResponseAllOf) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventResponseAllOf) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *EventResponseAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCommit

`func (o *EventResponseAllOf) GetCommit() CommitResponse`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *EventResponseAllOf) GetCommitOk() (*CommitResponse, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *EventResponseAllOf) SetCommit(v CommitResponse)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *EventResponseAllOf) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetStatus

`func (o *EventResponseAllOf) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventResponseAllOf) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventResponseAllOf) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EventResponseAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponseAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponseAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventResponseAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLog

`func (o *EventResponseAllOf) GetLog() ReferenceObject`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *EventResponseAllOf) GetLogOk() (*ReferenceObject, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *EventResponseAllOf) SetLog(v ReferenceObject)`

SetLog sets Log field to given value.

### HasLog

`func (o *EventResponseAllOf) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


