# AlertRuleListResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**AlertRuleSource**](AlertRuleSource.md) |  | 
**Name** | **string** | Name of the alert rule | 
**State** | [**AlertRuleState**](AlertRuleState.md) |  | 
**Target** | Pointer to [**GhostAlertRuleResponseAllOfTarget**](GhostAlertRuleResponseAllOfTarget.md) |  | [optional] 
**StartsAt** | Pointer to **NullableTime** | When the ghost alert started firing | [optional] 

## Methods

### NewAlertRuleListResultsInner

`func NewAlertRuleListResultsInner(source AlertRuleSource, name string, state AlertRuleState, ) *AlertRuleListResultsInner`

NewAlertRuleListResultsInner instantiates a new AlertRuleListResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleListResultsInnerWithDefaults

`func NewAlertRuleListResultsInnerWithDefaults() *AlertRuleListResultsInner`

NewAlertRuleListResultsInnerWithDefaults instantiates a new AlertRuleListResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AlertRuleListResultsInner) GetSource() AlertRuleSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AlertRuleListResultsInner) GetSourceOk() (*AlertRuleSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AlertRuleListResultsInner) SetSource(v AlertRuleSource)`

SetSource sets Source field to given value.


### GetName

`func (o *AlertRuleListResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRuleListResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRuleListResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AlertRuleListResultsInner) GetState() AlertRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AlertRuleListResultsInner) GetStateOk() (*AlertRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AlertRuleListResultsInner) SetState(v AlertRuleState)`

SetState sets State field to given value.


### GetTarget

`func (o *AlertRuleListResultsInner) GetTarget() GhostAlertRuleResponseAllOfTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AlertRuleListResultsInner) GetTargetOk() (*GhostAlertRuleResponseAllOfTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AlertRuleListResultsInner) SetTarget(v GhostAlertRuleResponseAllOfTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AlertRuleListResultsInner) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStartsAt

`func (o *AlertRuleListResultsInner) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *AlertRuleListResultsInner) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *AlertRuleListResultsInner) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *AlertRuleListResultsInner) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *AlertRuleListResultsInner) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *AlertRuleListResultsInner) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


