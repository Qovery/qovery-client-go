# EventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Type** | Pointer to **string** | DRAFT - we have to specify here all the possible events | [optional] 
**Log** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 

## Methods

### NewEventAllOf

`func NewEventAllOf() *EventAllOf`

NewEventAllOf instantiates a new EventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAllOfWithDefaults

`func NewEventAllOfWithDefaults() *EventAllOf`

NewEventAllOfWithDefaults instantiates a new EventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *EventAllOf) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventAllOf) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventAllOf) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *EventAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCommit

`func (o *EventAllOf) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *EventAllOf) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *EventAllOf) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *EventAllOf) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetStatus

`func (o *EventAllOf) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventAllOf) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventAllOf) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLog

`func (o *EventAllOf) GetLog() ReferenceObject`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *EventAllOf) GetLogOk() (*ReferenceObject, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *EventAllOf) SetLog(v ReferenceObject)`

SetLog sets Log field to given value.

### HasLog

`func (o *EventAllOf) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


