# ClusterFailureContextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ClusterId** | **string** |  | 
**GatewayStatus** | [**GatewayStatusResponse**](GatewayStatusResponse.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewClusterFailureContextResponse

`func NewClusterFailureContextResponse(id string, clusterId string, gatewayStatus GatewayStatusResponse, createdAt time.Time, ) *ClusterFailureContextResponse`

NewClusterFailureContextResponse instantiates a new ClusterFailureContextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFailureContextResponseWithDefaults

`func NewClusterFailureContextResponseWithDefaults() *ClusterFailureContextResponse`

NewClusterFailureContextResponseWithDefaults instantiates a new ClusterFailureContextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterFailureContextResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterFailureContextResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterFailureContextResponse) SetId(v string)`

SetId sets Id field to given value.


### GetClusterId

`func (o *ClusterFailureContextResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterFailureContextResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterFailureContextResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGatewayStatus

`func (o *ClusterFailureContextResponse) GetGatewayStatus() GatewayStatusResponse`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *ClusterFailureContextResponse) GetGatewayStatusOk() (*GatewayStatusResponse, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *ClusterFailureContextResponse) SetGatewayStatus(v GatewayStatusResponse)`

SetGatewayStatus sets GatewayStatus field to given value.


### GetCreatedAt

`func (o *ClusterFailureContextResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterFailureContextResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterFailureContextResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


