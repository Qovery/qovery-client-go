# ProjectCurrentCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ConsumedTimeInSeconds** | **int32** |  | 
**Cost** | [**Cost**](Cost.md) |  | 
**Environments** | Pointer to [**[]GenericObjectCurrentCost**](GenericObjectCurrentCost.md) |  | [optional] 

## Methods

### NewProjectCurrentCost

`func NewProjectCurrentCost(id string, name string, consumedTimeInSeconds int32, cost Cost, ) *ProjectCurrentCost`

NewProjectCurrentCost instantiates a new ProjectCurrentCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCurrentCostWithDefaults

`func NewProjectCurrentCostWithDefaults() *ProjectCurrentCost`

NewProjectCurrentCostWithDefaults instantiates a new ProjectCurrentCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCurrentCost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCurrentCost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCurrentCost) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectCurrentCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCurrentCost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCurrentCost) SetName(v string)`

SetName sets Name field to given value.


### GetConsumedTimeInSeconds

`func (o *ProjectCurrentCost) GetConsumedTimeInSeconds() int32`

GetConsumedTimeInSeconds returns the ConsumedTimeInSeconds field if non-nil, zero value otherwise.

### GetConsumedTimeInSecondsOk

`func (o *ProjectCurrentCost) GetConsumedTimeInSecondsOk() (*int32, bool)`

GetConsumedTimeInSecondsOk returns a tuple with the ConsumedTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedTimeInSeconds

`func (o *ProjectCurrentCost) SetConsumedTimeInSeconds(v int32)`

SetConsumedTimeInSeconds sets ConsumedTimeInSeconds field to given value.


### GetCost

`func (o *ProjectCurrentCost) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProjectCurrentCost) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProjectCurrentCost) SetCost(v Cost)`

SetCost sets Cost field to given value.


### GetEnvironments

`func (o *ProjectCurrentCost) GetEnvironments() []GenericObjectCurrentCost`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectCurrentCost) GetEnvironmentsOk() (*[]GenericObjectCurrentCost, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectCurrentCost) SetEnvironments(v []GenericObjectCurrentCost)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectCurrentCost) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


