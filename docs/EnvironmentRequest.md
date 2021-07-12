# EnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Cluster** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentRequest

`func NewEnvironmentRequest(name string, ) *EnvironmentRequest`

NewEnvironmentRequest instantiates a new EnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentRequestWithDefaults

`func NewEnvironmentRequestWithDefaults() *EnvironmentRequest`

NewEnvironmentRequestWithDefaults instantiates a new EnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCluster

`func (o *EnvironmentRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *EnvironmentRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *EnvironmentRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *EnvironmentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMode

`func (o *EnvironmentRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EnvironmentRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EnvironmentRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EnvironmentRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


