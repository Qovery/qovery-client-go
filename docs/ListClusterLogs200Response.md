# ListClusterLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ClusterLogs**](ClusterLogs.md) |  | [optional] 

## Methods

### NewListClusterLogs200Response

`func NewListClusterLogs200Response() *ListClusterLogs200Response`

NewListClusterLogs200Response instantiates a new ListClusterLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLogs200ResponseWithDefaults

`func NewListClusterLogs200ResponseWithDefaults() *ListClusterLogs200Response`

NewListClusterLogs200ResponseWithDefaults instantiates a new ListClusterLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListClusterLogs200Response) GetResults() []ClusterLogs`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListClusterLogs200Response) GetResultsOk() (*[]ClusterLogs, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListClusterLogs200Response) SetResults(v []ClusterLogs)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListClusterLogs200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


