# ClusterRoutingTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | [**[]ClusterRoutingTableResultsInner**](ClusterRoutingTableResultsInner.md) |  | 

## Methods

### NewClusterRoutingTableRequest

`func NewClusterRoutingTableRequest(routes []ClusterRoutingTableResultsInner, ) *ClusterRoutingTableRequest`

NewClusterRoutingTableRequest instantiates a new ClusterRoutingTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoutingTableRequestWithDefaults

`func NewClusterRoutingTableRequestWithDefaults() *ClusterRoutingTableRequest`

NewClusterRoutingTableRequestWithDefaults instantiates a new ClusterRoutingTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *ClusterRoutingTableRequest) GetRoutes() []ClusterRoutingTableResultsInner`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ClusterRoutingTableRequest) GetRoutesOk() (*[]ClusterRoutingTableResultsInner, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ClusterRoutingTableRequest) SetRoutes(v []ClusterRoutingTableResultsInner)`

SetRoutes sets Routes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


