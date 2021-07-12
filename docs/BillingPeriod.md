# BillingPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingStartedOn** | Pointer to **time.Time** |  | [optional] 
**BillingEndedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBillingPeriod

`func NewBillingPeriod() *BillingPeriod`

NewBillingPeriod instantiates a new BillingPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPeriodWithDefaults

`func NewBillingPeriodWithDefaults() *BillingPeriod`

NewBillingPeriodWithDefaults instantiates a new BillingPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingStartedOn

`func (o *BillingPeriod) GetBillingStartedOn() time.Time`

GetBillingStartedOn returns the BillingStartedOn field if non-nil, zero value otherwise.

### GetBillingStartedOnOk

`func (o *BillingPeriod) GetBillingStartedOnOk() (*time.Time, bool)`

GetBillingStartedOnOk returns a tuple with the BillingStartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartedOn

`func (o *BillingPeriod) SetBillingStartedOn(v time.Time)`

SetBillingStartedOn sets BillingStartedOn field to given value.

### HasBillingStartedOn

`func (o *BillingPeriod) HasBillingStartedOn() bool`

HasBillingStartedOn returns a boolean if a field has been set.

### GetBillingEndedOn

`func (o *BillingPeriod) GetBillingEndedOn() time.Time`

GetBillingEndedOn returns the BillingEndedOn field if non-nil, zero value otherwise.

### GetBillingEndedOnOk

`func (o *BillingPeriod) GetBillingEndedOnOk() (*time.Time, bool)`

GetBillingEndedOnOk returns a tuple with the BillingEndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndedOn

`func (o *BillingPeriod) SetBillingEndedOn(v time.Time)`

SetBillingEndedOn sets BillingEndedOn field to given value.

### HasBillingEndedOn

`func (o *BillingPeriod) HasBillingEndedOn() bool`

HasBillingEndedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


