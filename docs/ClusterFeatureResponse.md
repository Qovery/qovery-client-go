# ClusterFeatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CostPerMonthInCents** | Pointer to **int32** |  | [optional] 
**CostPerMonth** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**IsCloudProviderPayingFeature** | Pointer to **bool** |  | [optional] 
**CloudProviderFeatureDocumentation** | Pointer to **string** |  | [optional] 
**IsQoveryPayingFeature** | Pointer to **bool** |  | [optional] 
**QoveryFeatureDocumentation** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**ValueObject** | Pointer to [**ClusterFeatureResponseValueObject**](ClusterFeatureResponseValueObject.md) |  | [optional] 
**IsValueUpdatable** | Pointer to **bool** |  | [optional] [default to false]
**AcceptedValues** | Pointer to [**[]ClusterFeatureResponseAcceptedValuesInner**](ClusterFeatureResponseAcceptedValuesInner.md) |  | [optional] 

## Methods

### NewClusterFeatureResponse

`func NewClusterFeatureResponse() *ClusterFeatureResponse`

NewClusterFeatureResponse instantiates a new ClusterFeatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureResponseWithDefaults

`func NewClusterFeatureResponseWithDefaults() *ClusterFeatureResponse`

NewClusterFeatureResponseWithDefaults instantiates a new ClusterFeatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterFeatureResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterFeatureResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterFeatureResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterFeatureResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ClusterFeatureResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterFeatureResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterFeatureResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClusterFeatureResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterFeatureResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterFeatureResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterFeatureResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterFeatureResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCostPerMonthInCents

`func (o *ClusterFeatureResponse) GetCostPerMonthInCents() int32`

GetCostPerMonthInCents returns the CostPerMonthInCents field if non-nil, zero value otherwise.

### GetCostPerMonthInCentsOk

`func (o *ClusterFeatureResponse) GetCostPerMonthInCentsOk() (*int32, bool)`

GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonthInCents

`func (o *ClusterFeatureResponse) SetCostPerMonthInCents(v int32)`

SetCostPerMonthInCents sets CostPerMonthInCents field to given value.

### HasCostPerMonthInCents

`func (o *ClusterFeatureResponse) HasCostPerMonthInCents() bool`

HasCostPerMonthInCents returns a boolean if a field has been set.

### GetCostPerMonth

`func (o *ClusterFeatureResponse) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *ClusterFeatureResponse) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *ClusterFeatureResponse) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *ClusterFeatureResponse) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ClusterFeatureResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ClusterFeatureResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ClusterFeatureResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ClusterFeatureResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetIsCloudProviderPayingFeature

`func (o *ClusterFeatureResponse) GetIsCloudProviderPayingFeature() bool`

GetIsCloudProviderPayingFeature returns the IsCloudProviderPayingFeature field if non-nil, zero value otherwise.

### GetIsCloudProviderPayingFeatureOk

`func (o *ClusterFeatureResponse) GetIsCloudProviderPayingFeatureOk() (*bool, bool)`

GetIsCloudProviderPayingFeatureOk returns a tuple with the IsCloudProviderPayingFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProviderPayingFeature

`func (o *ClusterFeatureResponse) SetIsCloudProviderPayingFeature(v bool)`

SetIsCloudProviderPayingFeature sets IsCloudProviderPayingFeature field to given value.

### HasIsCloudProviderPayingFeature

`func (o *ClusterFeatureResponse) HasIsCloudProviderPayingFeature() bool`

HasIsCloudProviderPayingFeature returns a boolean if a field has been set.

### GetCloudProviderFeatureDocumentation

`func (o *ClusterFeatureResponse) GetCloudProviderFeatureDocumentation() string`

GetCloudProviderFeatureDocumentation returns the CloudProviderFeatureDocumentation field if non-nil, zero value otherwise.

### GetCloudProviderFeatureDocumentationOk

`func (o *ClusterFeatureResponse) GetCloudProviderFeatureDocumentationOk() (*string, bool)`

GetCloudProviderFeatureDocumentationOk returns a tuple with the CloudProviderFeatureDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderFeatureDocumentation

`func (o *ClusterFeatureResponse) SetCloudProviderFeatureDocumentation(v string)`

SetCloudProviderFeatureDocumentation sets CloudProviderFeatureDocumentation field to given value.

### HasCloudProviderFeatureDocumentation

`func (o *ClusterFeatureResponse) HasCloudProviderFeatureDocumentation() bool`

HasCloudProviderFeatureDocumentation returns a boolean if a field has been set.

### GetIsQoveryPayingFeature

`func (o *ClusterFeatureResponse) GetIsQoveryPayingFeature() bool`

GetIsQoveryPayingFeature returns the IsQoveryPayingFeature field if non-nil, zero value otherwise.

### GetIsQoveryPayingFeatureOk

`func (o *ClusterFeatureResponse) GetIsQoveryPayingFeatureOk() (*bool, bool)`

GetIsQoveryPayingFeatureOk returns a tuple with the IsQoveryPayingFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQoveryPayingFeature

`func (o *ClusterFeatureResponse) SetIsQoveryPayingFeature(v bool)`

SetIsQoveryPayingFeature sets IsQoveryPayingFeature field to given value.

### HasIsQoveryPayingFeature

`func (o *ClusterFeatureResponse) HasIsQoveryPayingFeature() bool`

HasIsQoveryPayingFeature returns a boolean if a field has been set.

### GetQoveryFeatureDocumentation

`func (o *ClusterFeatureResponse) GetQoveryFeatureDocumentation() string`

GetQoveryFeatureDocumentation returns the QoveryFeatureDocumentation field if non-nil, zero value otherwise.

### GetQoveryFeatureDocumentationOk

`func (o *ClusterFeatureResponse) GetQoveryFeatureDocumentationOk() (*string, bool)`

GetQoveryFeatureDocumentationOk returns a tuple with the QoveryFeatureDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryFeatureDocumentation

`func (o *ClusterFeatureResponse) SetQoveryFeatureDocumentation(v string)`

SetQoveryFeatureDocumentation sets QoveryFeatureDocumentation field to given value.

### HasQoveryFeatureDocumentation

`func (o *ClusterFeatureResponse) HasQoveryFeatureDocumentation() bool`

HasQoveryFeatureDocumentation returns a boolean if a field has been set.

### GetValueType

`func (o *ClusterFeatureResponse) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ClusterFeatureResponse) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ClusterFeatureResponse) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ClusterFeatureResponse) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValueObject

`func (o *ClusterFeatureResponse) GetValueObject() ClusterFeatureResponseValueObject`

GetValueObject returns the ValueObject field if non-nil, zero value otherwise.

### GetValueObjectOk

`func (o *ClusterFeatureResponse) GetValueObjectOk() (*ClusterFeatureResponseValueObject, bool)`

GetValueObjectOk returns a tuple with the ValueObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueObject

`func (o *ClusterFeatureResponse) SetValueObject(v ClusterFeatureResponseValueObject)`

SetValueObject sets ValueObject field to given value.

### HasValueObject

`func (o *ClusterFeatureResponse) HasValueObject() bool`

HasValueObject returns a boolean if a field has been set.

### GetIsValueUpdatable

`func (o *ClusterFeatureResponse) GetIsValueUpdatable() bool`

GetIsValueUpdatable returns the IsValueUpdatable field if non-nil, zero value otherwise.

### GetIsValueUpdatableOk

`func (o *ClusterFeatureResponse) GetIsValueUpdatableOk() (*bool, bool)`

GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueUpdatable

`func (o *ClusterFeatureResponse) SetIsValueUpdatable(v bool)`

SetIsValueUpdatable sets IsValueUpdatable field to given value.

### HasIsValueUpdatable

`func (o *ClusterFeatureResponse) HasIsValueUpdatable() bool`

HasIsValueUpdatable returns a boolean if a field has been set.

### GetAcceptedValues

`func (o *ClusterFeatureResponse) GetAcceptedValues() []ClusterFeatureResponseAcceptedValuesInner`

GetAcceptedValues returns the AcceptedValues field if non-nil, zero value otherwise.

### GetAcceptedValuesOk

`func (o *ClusterFeatureResponse) GetAcceptedValuesOk() (*[]ClusterFeatureResponseAcceptedValuesInner, bool)`

GetAcceptedValuesOk returns a tuple with the AcceptedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedValues

`func (o *ClusterFeatureResponse) SetAcceptedValues(v []ClusterFeatureResponseAcceptedValuesInner)`

SetAcceptedValues sets AcceptedValues field to given value.

### HasAcceptedValues

`func (o *ClusterFeatureResponse) HasAcceptedValues() bool`

HasAcceptedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


