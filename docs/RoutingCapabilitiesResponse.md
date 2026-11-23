# RoutingCapabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanShare** | **bool** | Whether several services can route the same custom domain. &#x60;false&#x60; on a &#x60;PER_ROUTER_LISTENER_SET&#x60; environment, where a domain belongs to one service. | 

## Methods

### NewRoutingCapabilitiesResponse

`func NewRoutingCapabilitiesResponse(canShare bool, ) *RoutingCapabilitiesResponse`

NewRoutingCapabilitiesResponse instantiates a new RoutingCapabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingCapabilitiesResponseWithDefaults

`func NewRoutingCapabilitiesResponseWithDefaults() *RoutingCapabilitiesResponse`

NewRoutingCapabilitiesResponseWithDefaults instantiates a new RoutingCapabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanShare

`func (o *RoutingCapabilitiesResponse) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *RoutingCapabilitiesResponse) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *RoutingCapabilitiesResponse) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


