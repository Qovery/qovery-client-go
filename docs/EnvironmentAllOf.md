# EnvironmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Project** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**LastUpdatedBy** | Pointer to **string** | uuid of the user that made the last update | [optional] 
**CloudProvider** | [**EnvironmentAllOfCloudProvider**](EnvironmentAllOfCloudProvider.md) |  | 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 
**ClusterId** | **string** |  | 
**ClusterName** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentAllOf

`func NewEnvironmentAllOf(name string, cloudProvider EnvironmentAllOfCloudProvider, mode EnvironmentModeEnum, clusterId string, ) *EnvironmentAllOf`

NewEnvironmentAllOf instantiates a new EnvironmentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAllOfWithDefaults

`func NewEnvironmentAllOfWithDefaults() *EnvironmentAllOf`

NewEnvironmentAllOfWithDefaults instantiates a new EnvironmentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *EnvironmentAllOf) GetProject() ReferenceObject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentAllOf) GetProjectOk() (*ReferenceObject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentAllOf) SetProject(v ReferenceObject)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentAllOf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EnvironmentAllOf) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EnvironmentAllOf) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EnvironmentAllOf) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EnvironmentAllOf) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetCloudProvider

`func (o *EnvironmentAllOf) GetCloudProvider() EnvironmentAllOfCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EnvironmentAllOf) GetCloudProviderOk() (*EnvironmentAllOfCloudProvider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EnvironmentAllOf) SetCloudProvider(v EnvironmentAllOfCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.


### GetMode

`func (o *EnvironmentAllOf) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EnvironmentAllOf) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EnvironmentAllOf) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.


### GetClusterId

`func (o *EnvironmentAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnvironmentAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnvironmentAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *EnvironmentAllOf) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *EnvironmentAllOf) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *EnvironmentAllOf) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *EnvironmentAllOf) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


