# BlueprintManifestEngineHelm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Chart** | [**BlueprintManifestChartConfig**](BlueprintManifestChartConfig.md) |  | 
**Timeout** | Pointer to **NullableInt64** |  | [optional] 
**Resources** | Pointer to [**BlueprintManifestResourcesConfig**](BlueprintManifestResourcesConfig.md) |  | [optional] 

## Methods

### NewBlueprintManifestEngineHelm

`func NewBlueprintManifestEngineHelm(type_ string, chart BlueprintManifestChartConfig, ) *BlueprintManifestEngineHelm`

NewBlueprintManifestEngineHelm instantiates a new BlueprintManifestEngineHelm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestEngineHelmWithDefaults

`func NewBlueprintManifestEngineHelmWithDefaults() *BlueprintManifestEngineHelm`

NewBlueprintManifestEngineHelmWithDefaults instantiates a new BlueprintManifestEngineHelm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintManifestEngineHelm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintManifestEngineHelm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintManifestEngineHelm) SetType(v string)`

SetType sets Type field to given value.


### GetChart

`func (o *BlueprintManifestEngineHelm) GetChart() BlueprintManifestChartConfig`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *BlueprintManifestEngineHelm) GetChartOk() (*BlueprintManifestChartConfig, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *BlueprintManifestEngineHelm) SetChart(v BlueprintManifestChartConfig)`

SetChart sets Chart field to given value.


### GetTimeout

`func (o *BlueprintManifestEngineHelm) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BlueprintManifestEngineHelm) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BlueprintManifestEngineHelm) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BlueprintManifestEngineHelm) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BlueprintManifestEngineHelm) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BlueprintManifestEngineHelm) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetResources

`func (o *BlueprintManifestEngineHelm) GetResources() BlueprintManifestResourcesConfig`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintManifestEngineHelm) GetResourcesOk() (*BlueprintManifestResourcesConfig, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintManifestEngineHelm) SetResources(v BlueprintManifestResourcesConfig)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintManifestEngineHelm) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


