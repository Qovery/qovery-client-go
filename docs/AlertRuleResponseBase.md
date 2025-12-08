# AlertRuleResponseBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**AlertRuleSource**](AlertRuleSource.md) |  | 
**Name** | **string** | Name of the alert rule | 
**State** | [**AlertRuleState**](AlertRuleState.md) |  | 

## Methods

### NewAlertRuleResponseBase

`func NewAlertRuleResponseBase(source AlertRuleSource, name string, state AlertRuleState, ) *AlertRuleResponseBase`

NewAlertRuleResponseBase instantiates a new AlertRuleResponseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleResponseBaseWithDefaults

`func NewAlertRuleResponseBaseWithDefaults() *AlertRuleResponseBase`

NewAlertRuleResponseBaseWithDefaults instantiates a new AlertRuleResponseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AlertRuleResponseBase) GetSource() AlertRuleSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AlertRuleResponseBase) GetSourceOk() (*AlertRuleSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AlertRuleResponseBase) SetSource(v AlertRuleSource)`

SetSource sets Source field to given value.


### GetName

`func (o *AlertRuleResponseBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRuleResponseBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRuleResponseBase) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AlertRuleResponseBase) GetState() AlertRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AlertRuleResponseBase) GetStateOk() (*AlertRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AlertRuleResponseBase) SetState(v AlertRuleState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


