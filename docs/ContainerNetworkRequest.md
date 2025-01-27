# ContainerNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StickySession** | Pointer to **bool** | Specify if the sticky session option (also called persistant session) is activated or not for this container. If activated, user will be redirected by the load balancer to the same instance each time he access to the container.  | [optional] [default to false]

## Methods

### NewContainerNetworkRequest

`func NewContainerNetworkRequest() *ContainerNetworkRequest`

NewContainerNetworkRequest instantiates a new ContainerNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerNetworkRequestWithDefaults

`func NewContainerNetworkRequestWithDefaults() *ContainerNetworkRequest`

NewContainerNetworkRequestWithDefaults instantiates a new ContainerNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStickySession

`func (o *ContainerNetworkRequest) GetStickySession() bool`

GetStickySession returns the StickySession field if non-nil, zero value otherwise.

### GetStickySessionOk

`func (o *ContainerNetworkRequest) GetStickySessionOk() (*bool, bool)`

GetStickySessionOk returns a tuple with the StickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySession

`func (o *ContainerNetworkRequest) SetStickySession(v bool)`

SetStickySession sets StickySession field to given value.

### HasStickySession

`func (o *ContainerNetworkRequest) HasStickySession() bool`

HasStickySession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


