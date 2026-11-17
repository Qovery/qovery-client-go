# CustomDomainEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the environment. | 
**Name** | **string** | Name of the environment. | 
**RoutingMode** | [**EnvironmentRoutingModeEnum**](EnvironmentRoutingModeEnum.md) |  | 

## Methods

### NewCustomDomainEnvironmentResponse

`func NewCustomDomainEnvironmentResponse(id string, name string, routingMode EnvironmentRoutingModeEnum, ) *CustomDomainEnvironmentResponse`

NewCustomDomainEnvironmentResponse instantiates a new CustomDomainEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainEnvironmentResponseWithDefaults

`func NewCustomDomainEnvironmentResponseWithDefaults() *CustomDomainEnvironmentResponse`

NewCustomDomainEnvironmentResponseWithDefaults instantiates a new CustomDomainEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomainEnvironmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainEnvironmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainEnvironmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomDomainEnvironmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDomainEnvironmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDomainEnvironmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRoutingMode

`func (o *CustomDomainEnvironmentResponse) GetRoutingMode() EnvironmentRoutingModeEnum`

GetRoutingMode returns the RoutingMode field if non-nil, zero value otherwise.

### GetRoutingModeOk

`func (o *CustomDomainEnvironmentResponse) GetRoutingModeOk() (*EnvironmentRoutingModeEnum, bool)`

GetRoutingModeOk returns a tuple with the RoutingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMode

`func (o *CustomDomainEnvironmentResponse) SetRoutingMode(v EnvironmentRoutingModeEnum)`

SetRoutingMode sets RoutingMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


