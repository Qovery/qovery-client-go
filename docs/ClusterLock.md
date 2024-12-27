# ClusterLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**TtlInDays** | Pointer to **float32** |  | [optional] 
**ClusterId** | **string** |  | 
**LockedAt** | **float32** |  | 
**OwnerName** | **string** |  | 

## Methods

### NewClusterLock

`func NewClusterLock(reason string, clusterId string, lockedAt float32, ownerName string, ) *ClusterLock`

NewClusterLock instantiates a new ClusterLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLockWithDefaults

`func NewClusterLockWithDefaults() *ClusterLock`

NewClusterLockWithDefaults instantiates a new ClusterLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ClusterLock) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterLock) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterLock) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTtlInDays

`func (o *ClusterLock) GetTtlInDays() float32`

GetTtlInDays returns the TtlInDays field if non-nil, zero value otherwise.

### GetTtlInDaysOk

`func (o *ClusterLock) GetTtlInDaysOk() (*float32, bool)`

GetTtlInDaysOk returns a tuple with the TtlInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlInDays

`func (o *ClusterLock) SetTtlInDays(v float32)`

SetTtlInDays sets TtlInDays field to given value.

### HasTtlInDays

`func (o *ClusterLock) HasTtlInDays() bool`

HasTtlInDays returns a boolean if a field has been set.

### GetClusterId

`func (o *ClusterLock) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterLock) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterLock) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetLockedAt

`func (o *ClusterLock) GetLockedAt() float32`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *ClusterLock) GetLockedAtOk() (*float32, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *ClusterLock) SetLockedAt(v float32)`

SetLockedAt sets LockedAt field to given value.


### GetOwnerName

`func (o *ClusterLock) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ClusterLock) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ClusterLock) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


