# ClusterLockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**TtlInDays** | Pointer to **float32** |  | [optional] 

## Methods

### NewClusterLockRequest

`func NewClusterLockRequest(reason string, ) *ClusterLockRequest`

NewClusterLockRequest instantiates a new ClusterLockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLockRequestWithDefaults

`func NewClusterLockRequestWithDefaults() *ClusterLockRequest`

NewClusterLockRequestWithDefaults instantiates a new ClusterLockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ClusterLockRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterLockRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterLockRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTtlInDays

`func (o *ClusterLockRequest) GetTtlInDays() float32`

GetTtlInDays returns the TtlInDays field if non-nil, zero value otherwise.

### GetTtlInDaysOk

`func (o *ClusterLockRequest) GetTtlInDaysOk() (*float32, bool)`

GetTtlInDaysOk returns a tuple with the TtlInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlInDays

`func (o *ClusterLockRequest) SetTtlInDays(v float32)`

SetTtlInDays sets TtlInDays field to given value.

### HasTtlInDays

`func (o *ClusterLockRequest) HasTtlInDays() bool`

HasTtlInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


