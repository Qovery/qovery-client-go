# DatabaseEndpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Database hostname, from the DB_ADDRESS terraform output | 
**Port** | Pointer to **NullableInt32** | Database port, from the DB_PORT terraform output; null when none is reported | [optional] 

## Methods

### NewDatabaseEndpointsResponse

`func NewDatabaseEndpointsResponse(host string, ) *DatabaseEndpointsResponse`

NewDatabaseEndpointsResponse instantiates a new DatabaseEndpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseEndpointsResponseWithDefaults

`func NewDatabaseEndpointsResponseWithDefaults() *DatabaseEndpointsResponse`

NewDatabaseEndpointsResponseWithDefaults instantiates a new DatabaseEndpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DatabaseEndpointsResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseEndpointsResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseEndpointsResponse) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DatabaseEndpointsResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseEndpointsResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseEndpointsResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DatabaseEndpointsResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DatabaseEndpointsResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DatabaseEndpointsResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


