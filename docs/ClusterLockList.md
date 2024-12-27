# ClusterLockList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ClusterLock**](ClusterLock.md) |  | [optional] 

## Methods

### NewClusterLockList

`func NewClusterLockList() *ClusterLockList`

NewClusterLockList instantiates a new ClusterLockList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLockListWithDefaults

`func NewClusterLockListWithDefaults() *ClusterLockList`

NewClusterLockListWithDefaults instantiates a new ClusterLockList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ClusterLockList) GetResults() []ClusterLock`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ClusterLockList) GetResultsOk() (*[]ClusterLock, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ClusterLockList) SetResults(v []ClusterLock)`

SetResults sets Results field to given value.

### HasResults

`func (o *ClusterLockList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


