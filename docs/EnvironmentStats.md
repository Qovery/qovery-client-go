# EnvironmentStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ServiceTotalNumber** | Pointer to **float32** |  | [optional] 

## Methods

### NewEnvironmentStats

`func NewEnvironmentStats(id string, ) *EnvironmentStats`

NewEnvironmentStats instantiates a new EnvironmentStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatsWithDefaults

`func NewEnvironmentStatsWithDefaults() *EnvironmentStats`

NewEnvironmentStatsWithDefaults instantiates a new EnvironmentStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentStats) SetId(v string)`

SetId sets Id field to given value.


### GetServiceTotalNumber

`func (o *EnvironmentStats) GetServiceTotalNumber() float32`

GetServiceTotalNumber returns the ServiceTotalNumber field if non-nil, zero value otherwise.

### GetServiceTotalNumberOk

`func (o *EnvironmentStats) GetServiceTotalNumberOk() (*float32, bool)`

GetServiceTotalNumberOk returns a tuple with the ServiceTotalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTotalNumber

`func (o *EnvironmentStats) SetServiceTotalNumber(v float32)`

SetServiceTotalNumber sets ServiceTotalNumber field to given value.

### HasServiceTotalNumber

`func (o *EnvironmentStats) HasServiceTotalNumber() bool`

HasServiceTotalNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


