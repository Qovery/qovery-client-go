# ClusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedCloudProviderCost** | Pointer to **float32** | This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**HasAccess** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CloudProvider** | **string** |  | 
**Region** | **string** |  | 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **float32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **float32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**Title** | Pointer to **string** |  | [optional] 
**CostPerMonthInCents** | Pointer to **NullableInt32** |  | [optional] 
**CostPerMonth** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsValueUpdatable** | Pointer to **bool** |  | [optional] [default to false]
**AcceptedValues** | Pointer to [**[]interface{}**](interface{}.md) |  | [optional] 

## Methods

### NewClusterResponse

`func NewClusterResponse(id string, createdAt time.Time, name string, cloudProvider string, region string, ) *ClusterResponse`

NewClusterResponse instantiates a new ClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResponseWithDefaults

`func NewClusterResponseWithDefaults() *ClusterResponse`

NewClusterResponseWithDefaults instantiates a new ClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedCloudProviderCost

`func (o *ClusterResponse) GetEstimatedCloudProviderCost() float32`

GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field if non-nil, zero value otherwise.

### GetEstimatedCloudProviderCostOk

`func (o *ClusterResponse) GetEstimatedCloudProviderCostOk() (*float32, bool)`

GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCloudProviderCost

`func (o *ClusterResponse) SetEstimatedCloudProviderCost(v float32)`

SetEstimatedCloudProviderCost sets EstimatedCloudProviderCost field to given value.

### HasEstimatedCloudProviderCost

`func (o *ClusterResponse) HasEstimatedCloudProviderCost() bool`

HasEstimatedCloudProviderCost returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHasAccess

`func (o *ClusterResponse) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *ClusterResponse) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *ClusterResponse) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *ClusterResponse) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *ClusterResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ClusterResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ClusterResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClusterResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ClusterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCloudProvider

`func (o *ClusterResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *ClusterResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAutoUpdate

`func (o *ClusterResponse) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *ClusterResponse) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *ClusterResponse) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *ClusterResponse) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetCpu

`func (o *ClusterResponse) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterResponse) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterResponse) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ClusterResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ClusterResponse) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterResponse) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterResponse) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ClusterResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningNodes

`func (o *ClusterResponse) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *ClusterResponse) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *ClusterResponse) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *ClusterResponse) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *ClusterResponse) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *ClusterResponse) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *ClusterResponse) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *ClusterResponse) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.

### GetTitle

`func (o *ClusterResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClusterResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCostPerMonthInCents

`func (o *ClusterResponse) GetCostPerMonthInCents() int32`

GetCostPerMonthInCents returns the CostPerMonthInCents field if non-nil, zero value otherwise.

### GetCostPerMonthInCentsOk

`func (o *ClusterResponse) GetCostPerMonthInCentsOk() (*int32, bool)`

GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonthInCents

`func (o *ClusterResponse) SetCostPerMonthInCents(v int32)`

SetCostPerMonthInCents sets CostPerMonthInCents field to given value.

### HasCostPerMonthInCents

`func (o *ClusterResponse) HasCostPerMonthInCents() bool`

HasCostPerMonthInCents returns a boolean if a field has been set.

### SetCostPerMonthInCentsNil

`func (o *ClusterResponse) SetCostPerMonthInCentsNil(b bool)`

 SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil

### UnsetCostPerMonthInCents
`func (o *ClusterResponse) UnsetCostPerMonthInCents()`

UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
### GetCostPerMonth

`func (o *ClusterResponse) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *ClusterResponse) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *ClusterResponse) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *ClusterResponse) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### SetCostPerMonthNil

`func (o *ClusterResponse) SetCostPerMonthNil(b bool)`

 SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil

### UnsetCostPerMonth
`func (o *ClusterResponse) UnsetCostPerMonth()`

UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
### GetCurrencyCode

`func (o *ClusterResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ClusterResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ClusterResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ClusterResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *ClusterResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *ClusterResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetValueType

`func (o *ClusterResponse) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ClusterResponse) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ClusterResponse) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ClusterResponse) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *ClusterResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClusterResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClusterResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClusterResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsValueUpdatable

`func (o *ClusterResponse) GetIsValueUpdatable() bool`

GetIsValueUpdatable returns the IsValueUpdatable field if non-nil, zero value otherwise.

### GetIsValueUpdatableOk

`func (o *ClusterResponse) GetIsValueUpdatableOk() (*bool, bool)`

GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueUpdatable

`func (o *ClusterResponse) SetIsValueUpdatable(v bool)`

SetIsValueUpdatable sets IsValueUpdatable field to given value.

### HasIsValueUpdatable

`func (o *ClusterResponse) HasIsValueUpdatable() bool`

HasIsValueUpdatable returns a boolean if a field has been set.

### GetAcceptedValues

`func (o *ClusterResponse) GetAcceptedValues() []interface{}`

GetAcceptedValues returns the AcceptedValues field if non-nil, zero value otherwise.

### GetAcceptedValuesOk

`func (o *ClusterResponse) GetAcceptedValuesOk() (*[]interface{}, bool)`

GetAcceptedValuesOk returns a tuple with the AcceptedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedValues

`func (o *ClusterResponse) SetAcceptedValues(v []interface{})`

SetAcceptedValues sets AcceptedValues field to given value.

### HasAcceptedValues

`func (o *ClusterResponse) HasAcceptedValues() bool`

HasAcceptedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


