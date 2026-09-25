# SelfManagedClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Production** | Pointer to **bool** |  | [optional] [default to false]
**Provider** | **string** |  | 
**Region** | **string** |  | 
**Credentials** | [**SelfManagedClusterCredentials**](SelfManagedClusterCredentials.md) |  | 
**Platform** | [**PlatformSelection**](PlatformSelection.md) |  | 
**ClusterInputs** | Pointer to **map[string]map[string]string** | String values keyed first by component key and then by input key | [optional] 

## Methods

### NewSelfManagedClusterRequest

`func NewSelfManagedClusterRequest(name string, provider string, region string, credentials SelfManagedClusterCredentials, platform PlatformSelection, ) *SelfManagedClusterRequest`

NewSelfManagedClusterRequest instantiates a new SelfManagedClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfManagedClusterRequestWithDefaults

`func NewSelfManagedClusterRequestWithDefaults() *SelfManagedClusterRequest`

NewSelfManagedClusterRequestWithDefaults instantiates a new SelfManagedClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SelfManagedClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelfManagedClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelfManagedClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProduction

`func (o *SelfManagedClusterRequest) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *SelfManagedClusterRequest) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *SelfManagedClusterRequest) SetProduction(v bool)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *SelfManagedClusterRequest) HasProduction() bool`

HasProduction returns a boolean if a field has been set.

### GetProvider

`func (o *SelfManagedClusterRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SelfManagedClusterRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SelfManagedClusterRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *SelfManagedClusterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SelfManagedClusterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SelfManagedClusterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCredentials

`func (o *SelfManagedClusterRequest) GetCredentials() SelfManagedClusterCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SelfManagedClusterRequest) GetCredentialsOk() (*SelfManagedClusterCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SelfManagedClusterRequest) SetCredentials(v SelfManagedClusterCredentials)`

SetCredentials sets Credentials field to given value.


### GetPlatform

`func (o *SelfManagedClusterRequest) GetPlatform() PlatformSelection`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SelfManagedClusterRequest) GetPlatformOk() (*PlatformSelection, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SelfManagedClusterRequest) SetPlatform(v PlatformSelection)`

SetPlatform sets Platform field to given value.


### GetClusterInputs

`func (o *SelfManagedClusterRequest) GetClusterInputs() map[string]map[string]string`

GetClusterInputs returns the ClusterInputs field if non-nil, zero value otherwise.

### GetClusterInputsOk

`func (o *SelfManagedClusterRequest) GetClusterInputsOk() (*map[string]map[string]string, bool)`

GetClusterInputsOk returns a tuple with the ClusterInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInputs

`func (o *SelfManagedClusterRequest) SetClusterInputs(v map[string]map[string]string)`

SetClusterInputs sets ClusterInputs field to given value.

### HasClusterInputs

`func (o *SelfManagedClusterRequest) HasClusterInputs() bool`

HasClusterInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


