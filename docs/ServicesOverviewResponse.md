# ServicesOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCount** | **int32** |  | 
**ManagedBy** | **string** |  | 

## Methods

### NewServicesOverviewResponse

`func NewServicesOverviewResponse(serviceCount int32, managedBy string, ) *ServicesOverviewResponse`

NewServicesOverviewResponse instantiates a new ServicesOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesOverviewResponseWithDefaults

`func NewServicesOverviewResponseWithDefaults() *ServicesOverviewResponse`

NewServicesOverviewResponseWithDefaults instantiates a new ServicesOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceCount

`func (o *ServicesOverviewResponse) GetServiceCount() int32`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *ServicesOverviewResponse) GetServiceCountOk() (*int32, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCount

`func (o *ServicesOverviewResponse) SetServiceCount(v int32)`

SetServiceCount sets ServiceCount field to given value.


### GetManagedBy

`func (o *ServicesOverviewResponse) GetManagedBy() string`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *ServicesOverviewResponse) GetManagedByOk() (*string, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *ServicesOverviewResponse) SetManagedBy(v string)`

SetManagedBy sets ManagedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


