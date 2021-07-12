# ProjectStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ServiceTotalNumber** | Pointer to **float32** |  | [optional] 
**EnvironmentTotalNumber** | Pointer to **float32** |  | [optional] 

## Methods

### NewProjectStatsResponse

`func NewProjectStatsResponse(id string, ) *ProjectStatsResponse`

NewProjectStatsResponse instantiates a new ProjectStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectStatsResponseWithDefaults

`func NewProjectStatsResponseWithDefaults() *ProjectStatsResponse`

NewProjectStatsResponseWithDefaults instantiates a new ProjectStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectStatsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectStatsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectStatsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetServiceTotalNumber

`func (o *ProjectStatsResponse) GetServiceTotalNumber() float32`

GetServiceTotalNumber returns the ServiceTotalNumber field if non-nil, zero value otherwise.

### GetServiceTotalNumberOk

`func (o *ProjectStatsResponse) GetServiceTotalNumberOk() (*float32, bool)`

GetServiceTotalNumberOk returns a tuple with the ServiceTotalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTotalNumber

`func (o *ProjectStatsResponse) SetServiceTotalNumber(v float32)`

SetServiceTotalNumber sets ServiceTotalNumber field to given value.

### HasServiceTotalNumber

`func (o *ProjectStatsResponse) HasServiceTotalNumber() bool`

HasServiceTotalNumber returns a boolean if a field has been set.

### GetEnvironmentTotalNumber

`func (o *ProjectStatsResponse) GetEnvironmentTotalNumber() float32`

GetEnvironmentTotalNumber returns the EnvironmentTotalNumber field if non-nil, zero value otherwise.

### GetEnvironmentTotalNumberOk

`func (o *ProjectStatsResponse) GetEnvironmentTotalNumberOk() (*float32, bool)`

GetEnvironmentTotalNumberOk returns a tuple with the EnvironmentTotalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTotalNumber

`func (o *ProjectStatsResponse) SetEnvironmentTotalNumber(v float32)`

SetEnvironmentTotalNumber sets EnvironmentTotalNumber field to given value.

### HasEnvironmentTotalNumber

`func (o *ProjectStatsResponse) HasEnvironmentTotalNumber() bool`

HasEnvironmentTotalNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


