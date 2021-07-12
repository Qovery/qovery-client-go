# ProjectCurrentCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**[]GenericObjectCurrentCostResponse**](GenericObjectCurrentCostResponse.md) |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**BillingStartedOn** | **time.Time** |  | 
**BillingEndedOn** | **time.Time** |  | 
**ConsumedTimeInSeconds** | **int32** |  | 
**Cost** | [**CostResponse**](CostResponse.md) |  | 

## Methods

### NewProjectCurrentCostResponse

`func NewProjectCurrentCostResponse(id string, name string, billingStartedOn time.Time, billingEndedOn time.Time, consumedTimeInSeconds int32, cost CostResponse, ) *ProjectCurrentCostResponse`

NewProjectCurrentCostResponse instantiates a new ProjectCurrentCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCurrentCostResponseWithDefaults

`func NewProjectCurrentCostResponseWithDefaults() *ProjectCurrentCostResponse`

NewProjectCurrentCostResponseWithDefaults instantiates a new ProjectCurrentCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *ProjectCurrentCostResponse) GetEnvironments() []GenericObjectCurrentCostResponse`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectCurrentCostResponse) GetEnvironmentsOk() (*[]GenericObjectCurrentCostResponse, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectCurrentCostResponse) SetEnvironments(v []GenericObjectCurrentCostResponse)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectCurrentCostResponse) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *ProjectCurrentCostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCurrentCostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCurrentCostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectCurrentCostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCurrentCostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCurrentCostResponse) SetName(v string)`

SetName sets Name field to given value.


### GetBillingStartedOn

`func (o *ProjectCurrentCostResponse) GetBillingStartedOn() time.Time`

GetBillingStartedOn returns the BillingStartedOn field if non-nil, zero value otherwise.

### GetBillingStartedOnOk

`func (o *ProjectCurrentCostResponse) GetBillingStartedOnOk() (*time.Time, bool)`

GetBillingStartedOnOk returns a tuple with the BillingStartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartedOn

`func (o *ProjectCurrentCostResponse) SetBillingStartedOn(v time.Time)`

SetBillingStartedOn sets BillingStartedOn field to given value.


### GetBillingEndedOn

`func (o *ProjectCurrentCostResponse) GetBillingEndedOn() time.Time`

GetBillingEndedOn returns the BillingEndedOn field if non-nil, zero value otherwise.

### GetBillingEndedOnOk

`func (o *ProjectCurrentCostResponse) GetBillingEndedOnOk() (*time.Time, bool)`

GetBillingEndedOnOk returns a tuple with the BillingEndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndedOn

`func (o *ProjectCurrentCostResponse) SetBillingEndedOn(v time.Time)`

SetBillingEndedOn sets BillingEndedOn field to given value.


### GetConsumedTimeInSeconds

`func (o *ProjectCurrentCostResponse) GetConsumedTimeInSeconds() int32`

GetConsumedTimeInSeconds returns the ConsumedTimeInSeconds field if non-nil, zero value otherwise.

### GetConsumedTimeInSecondsOk

`func (o *ProjectCurrentCostResponse) GetConsumedTimeInSecondsOk() (*int32, bool)`

GetConsumedTimeInSecondsOk returns a tuple with the ConsumedTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedTimeInSeconds

`func (o *ProjectCurrentCostResponse) SetConsumedTimeInSeconds(v int32)`

SetConsumedTimeInSeconds sets ConsumedTimeInSeconds field to given value.


### GetCost

`func (o *ProjectCurrentCostResponse) GetCost() CostResponse`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProjectCurrentCostResponse) GetCostOk() (*CostResponse, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProjectCurrentCostResponse) SetCost(v CostResponse)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


