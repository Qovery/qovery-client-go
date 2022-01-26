# GenericObjectCurrentCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ConsumedTimeInSeconds** | **int32** |  | 
**Cost** | [**CostResponse**](CostResponse.md) |  | 

## Methods

### NewGenericObjectCurrentCostResponse

`func NewGenericObjectCurrentCostResponse(id string, name string, consumedTimeInSeconds int32, cost CostResponse, ) *GenericObjectCurrentCostResponse`

NewGenericObjectCurrentCostResponse instantiates a new GenericObjectCurrentCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericObjectCurrentCostResponseWithDefaults

`func NewGenericObjectCurrentCostResponseWithDefaults() *GenericObjectCurrentCostResponse`

NewGenericObjectCurrentCostResponseWithDefaults instantiates a new GenericObjectCurrentCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenericObjectCurrentCostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericObjectCurrentCostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericObjectCurrentCostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GenericObjectCurrentCostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericObjectCurrentCostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericObjectCurrentCostResponse) SetName(v string)`

SetName sets Name field to given value.


### GetConsumedTimeInSeconds

`func (o *GenericObjectCurrentCostResponse) GetConsumedTimeInSeconds() int32`

GetConsumedTimeInSeconds returns the ConsumedTimeInSeconds field if non-nil, zero value otherwise.

### GetConsumedTimeInSecondsOk

`func (o *GenericObjectCurrentCostResponse) GetConsumedTimeInSecondsOk() (*int32, bool)`

GetConsumedTimeInSecondsOk returns a tuple with the ConsumedTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedTimeInSeconds

`func (o *GenericObjectCurrentCostResponse) SetConsumedTimeInSeconds(v int32)`

SetConsumedTimeInSeconds sets ConsumedTimeInSeconds field to given value.


### GetCost

`func (o *GenericObjectCurrentCostResponse) GetCost() CostResponse`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GenericObjectCurrentCostResponse) GetCostOk() (*CostResponse, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GenericObjectCurrentCostResponse) SetCost(v CostResponse)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


