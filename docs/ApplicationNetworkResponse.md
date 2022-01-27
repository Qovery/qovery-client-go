# ApplicationNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StickySession** | Pointer to **bool** | Specify if the sticky session option (also called persistant session) is activated or not for this application. If activated, user will be redirected by the load balancer to the same instance each time he access to the application.   | [optional] [default to false]

## Methods

### NewApplicationNetworkResponse

`func NewApplicationNetworkResponse() *ApplicationNetworkResponse`

NewApplicationNetworkResponse instantiates a new ApplicationNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationNetworkResponseWithDefaults

`func NewApplicationNetworkResponseWithDefaults() *ApplicationNetworkResponse`

NewApplicationNetworkResponseWithDefaults instantiates a new ApplicationNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStickySession

`func (o *ApplicationNetworkResponse) GetStickySession() bool`

GetStickySession returns the StickySession field if non-nil, zero value otherwise.

### GetStickySessionOk

`func (o *ApplicationNetworkResponse) GetStickySessionOk() (*bool, bool)`

GetStickySessionOk returns a tuple with the StickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySession

`func (o *ApplicationNetworkResponse) SetStickySession(v bool)`

SetStickySession sets StickySession field to given value.

### HasStickySession

`func (o *ApplicationNetworkResponse) HasStickySession() bool`

HasStickySession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


