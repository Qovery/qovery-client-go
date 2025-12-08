# GhostAlertRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to [**GhostAlertRuleResponseAllOfTarget**](GhostAlertRuleResponseAllOfTarget.md) |  | [optional] 
**StartsAt** | Pointer to **NullableTime** | When the ghost alert started firing | [optional] 

## Methods

### NewGhostAlertRuleResponse

`func NewGhostAlertRuleResponse() *GhostAlertRuleResponse`

NewGhostAlertRuleResponse instantiates a new GhostAlertRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGhostAlertRuleResponseWithDefaults

`func NewGhostAlertRuleResponseWithDefaults() *GhostAlertRuleResponse`

NewGhostAlertRuleResponseWithDefaults instantiates a new GhostAlertRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *GhostAlertRuleResponse) GetTarget() GhostAlertRuleResponseAllOfTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GhostAlertRuleResponse) GetTargetOk() (*GhostAlertRuleResponseAllOfTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GhostAlertRuleResponse) SetTarget(v GhostAlertRuleResponseAllOfTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GhostAlertRuleResponse) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStartsAt

`func (o *GhostAlertRuleResponse) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GhostAlertRuleResponse) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GhostAlertRuleResponse) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GhostAlertRuleResponse) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GhostAlertRuleResponse) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GhostAlertRuleResponse) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


