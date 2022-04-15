# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CloudProvider** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | 
**Region** | **string** |  | 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**InstanceType** | Pointer to **string** | the instance type to be used for this cluster. The list of values can be retrieved via the endpoint /{CloudProvider}/instanceType | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**CostPerMonthInCents** | Pointer to **NullableInt32** |  | [optional] 
**CostPerMonth** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsValueUpdatable** | Pointer to **bool** |  | [optional] [default to false]
**AcceptedValues** | Pointer to [**[]interface{}**](interface{}.md) |  | [optional] 
**EstimatedCloudProviderCost** | Pointer to **int32** | This is an estimation of the cost this cluster will represent on your cloud proider bill, based on your current configuration | [optional] 
**Status** | Pointer to [**ClusterStatusEnum**](ClusterStatusEnum.md) |  | [optional] 
**HasAccess** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, createdAt time.Time, name string, cloudProvider CloudProviderEnum, region string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Cluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Cluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCloudProvider

`func (o *Cluster) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *Cluster) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *Cluster) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *Cluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Cluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Cluster) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAutoUpdate

`func (o *Cluster) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *Cluster) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *Cluster) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *Cluster) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetCpu

`func (o *Cluster) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Cluster) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Cluster) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Cluster) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *Cluster) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Cluster) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Cluster) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Cluster) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningNodes

`func (o *Cluster) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *Cluster) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *Cluster) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *Cluster) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *Cluster) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *Cluster) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *Cluster) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *Cluster) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.

### GetInstanceType

`func (o *Cluster) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Cluster) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Cluster) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Cluster) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetTitle

`func (o *Cluster) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Cluster) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Cluster) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Cluster) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCostPerMonthInCents

`func (o *Cluster) GetCostPerMonthInCents() int32`

GetCostPerMonthInCents returns the CostPerMonthInCents field if non-nil, zero value otherwise.

### GetCostPerMonthInCentsOk

`func (o *Cluster) GetCostPerMonthInCentsOk() (*int32, bool)`

GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonthInCents

`func (o *Cluster) SetCostPerMonthInCents(v int32)`

SetCostPerMonthInCents sets CostPerMonthInCents field to given value.

### HasCostPerMonthInCents

`func (o *Cluster) HasCostPerMonthInCents() bool`

HasCostPerMonthInCents returns a boolean if a field has been set.

### SetCostPerMonthInCentsNil

`func (o *Cluster) SetCostPerMonthInCentsNil(b bool)`

 SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil

### UnsetCostPerMonthInCents
`func (o *Cluster) UnsetCostPerMonthInCents()`

UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
### GetCostPerMonth

`func (o *Cluster) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *Cluster) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *Cluster) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *Cluster) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### SetCostPerMonthNil

`func (o *Cluster) SetCostPerMonthNil(b bool)`

 SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil

### UnsetCostPerMonth
`func (o *Cluster) UnsetCostPerMonth()`

UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
### GetCurrencyCode

`func (o *Cluster) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Cluster) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Cluster) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Cluster) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *Cluster) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *Cluster) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetValueType

`func (o *Cluster) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Cluster) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Cluster) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Cluster) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *Cluster) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Cluster) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Cluster) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Cluster) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Cluster) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Cluster) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsValueUpdatable

`func (o *Cluster) GetIsValueUpdatable() bool`

GetIsValueUpdatable returns the IsValueUpdatable field if non-nil, zero value otherwise.

### GetIsValueUpdatableOk

`func (o *Cluster) GetIsValueUpdatableOk() (*bool, bool)`

GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueUpdatable

`func (o *Cluster) SetIsValueUpdatable(v bool)`

SetIsValueUpdatable sets IsValueUpdatable field to given value.

### HasIsValueUpdatable

`func (o *Cluster) HasIsValueUpdatable() bool`

HasIsValueUpdatable returns a boolean if a field has been set.

### GetAcceptedValues

`func (o *Cluster) GetAcceptedValues() []interface{}`

GetAcceptedValues returns the AcceptedValues field if non-nil, zero value otherwise.

### GetAcceptedValuesOk

`func (o *Cluster) GetAcceptedValuesOk() (*[]interface{}, bool)`

GetAcceptedValuesOk returns a tuple with the AcceptedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedValues

`func (o *Cluster) SetAcceptedValues(v []interface{})`

SetAcceptedValues sets AcceptedValues field to given value.

### HasAcceptedValues

`func (o *Cluster) HasAcceptedValues() bool`

HasAcceptedValues returns a boolean if a field has been set.

### GetEstimatedCloudProviderCost

`func (o *Cluster) GetEstimatedCloudProviderCost() int32`

GetEstimatedCloudProviderCost returns the EstimatedCloudProviderCost field if non-nil, zero value otherwise.

### GetEstimatedCloudProviderCostOk

`func (o *Cluster) GetEstimatedCloudProviderCostOk() (*int32, bool)`

GetEstimatedCloudProviderCostOk returns a tuple with the EstimatedCloudProviderCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCloudProviderCost

`func (o *Cluster) SetEstimatedCloudProviderCost(v int32)`

SetEstimatedCloudProviderCost sets EstimatedCloudProviderCost field to given value.

### HasEstimatedCloudProviderCost

`func (o *Cluster) HasEstimatedCloudProviderCost() bool`

HasEstimatedCloudProviderCost returns a boolean if a field has been set.

### GetStatus

`func (o *Cluster) GetStatus() ClusterStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*ClusterStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v ClusterStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHasAccess

`func (o *Cluster) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *Cluster) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *Cluster) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *Cluster) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsDefault

`func (o *Cluster) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Cluster) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Cluster) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Cluster) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


