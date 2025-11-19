# AlertRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AlertRuleConditionKind**](AlertRuleConditionKind.md) |  | 
**Operator** | Pointer to [**AlertRuleConditionOperator**](AlertRuleConditionOperator.md) |  | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 
**Function** | Pointer to [**AlertRuleConditionFunction**](AlertRuleConditionFunction.md) |  | [optional] 
**Promql** | **string** |  | 

## Methods

### NewAlertRuleCondition

`func NewAlertRuleCondition(kind AlertRuleConditionKind, promql string, ) *AlertRuleCondition`

NewAlertRuleCondition instantiates a new AlertRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleConditionWithDefaults

`func NewAlertRuleConditionWithDefaults() *AlertRuleCondition`

NewAlertRuleConditionWithDefaults instantiates a new AlertRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AlertRuleCondition) GetKind() AlertRuleConditionKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlertRuleCondition) GetKindOk() (*AlertRuleConditionKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlertRuleCondition) SetKind(v AlertRuleConditionKind)`

SetKind sets Kind field to given value.


### GetOperator

`func (o *AlertRuleCondition) GetOperator() AlertRuleConditionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AlertRuleCondition) GetOperatorOk() (*AlertRuleConditionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AlertRuleCondition) SetOperator(v AlertRuleConditionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AlertRuleCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *AlertRuleCondition) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AlertRuleCondition) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AlertRuleCondition) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *AlertRuleCondition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetFunction

`func (o *AlertRuleCondition) GetFunction() AlertRuleConditionFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *AlertRuleCondition) GetFunctionOk() (*AlertRuleConditionFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *AlertRuleCondition) SetFunction(v AlertRuleConditionFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *AlertRuleCondition) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetPromql

`func (o *AlertRuleCondition) GetPromql() string`

GetPromql returns the Promql field if non-nil, zero value otherwise.

### GetPromqlOk

`func (o *AlertRuleCondition) GetPromqlOk() (*string, bool)`

GetPromqlOk returns a tuple with the Promql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromql

`func (o *AlertRuleCondition) SetPromql(v string)`

SetPromql sets Promql field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


