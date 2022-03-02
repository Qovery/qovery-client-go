# EnvironmentResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Project** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**LastUpdatedBy** | Pointer to **string** | uuid of the user that made the last update | [optional] 
**CloudProvider** | [**EnvironmentResponseAllOfCloudProvider**](EnvironmentResponseAllOfCloudProvider.md) |  | 
**Mode** | **string** |  | 
**ClusterId** | **string** |  | 

## Methods

### NewEnvironmentResponseAllOf

`func NewEnvironmentResponseAllOf(name string, cloudProvider EnvironmentResponseAllOfCloudProvider, mode string, clusterId string, ) *EnvironmentResponseAllOf`

NewEnvironmentResponseAllOf instantiates a new EnvironmentResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResponseAllOfWithDefaults

`func NewEnvironmentResponseAllOfWithDefaults() *EnvironmentResponseAllOf`

NewEnvironmentResponseAllOfWithDefaults instantiates a new EnvironmentResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentResponseAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *EnvironmentResponseAllOf) GetProject() ReferenceObject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentResponseAllOf) GetProjectOk() (*ReferenceObject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentResponseAllOf) SetProject(v ReferenceObject)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentResponseAllOf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EnvironmentResponseAllOf) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EnvironmentResponseAllOf) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EnvironmentResponseAllOf) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EnvironmentResponseAllOf) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetCloudProvider

`func (o *EnvironmentResponseAllOf) GetCloudProvider() EnvironmentResponseAllOfCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EnvironmentResponseAllOf) GetCloudProviderOk() (*EnvironmentResponseAllOfCloudProvider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EnvironmentResponseAllOf) SetCloudProvider(v EnvironmentResponseAllOfCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.


### GetMode

`func (o *EnvironmentResponseAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EnvironmentResponseAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EnvironmentResponseAllOf) SetMode(v string)`

SetMode sets Mode field to given value.


### GetClusterId

`func (o *EnvironmentResponseAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnvironmentResponseAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnvironmentResponseAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


