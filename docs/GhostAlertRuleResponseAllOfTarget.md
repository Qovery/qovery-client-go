# GhostAlertRuleResponseAllOfTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | [**AlertTargetType**](AlertTargetType.md) |  | 
**TargetId** | **string** |  | 
**Service** | Pointer to [**ServiceLightResponse**](ServiceLightResponse.md) |  | [optional] 

## Methods

### NewGhostAlertRuleResponseAllOfTarget

`func NewGhostAlertRuleResponseAllOfTarget(targetType AlertTargetType, targetId string, ) *GhostAlertRuleResponseAllOfTarget`

NewGhostAlertRuleResponseAllOfTarget instantiates a new GhostAlertRuleResponseAllOfTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGhostAlertRuleResponseAllOfTargetWithDefaults

`func NewGhostAlertRuleResponseAllOfTargetWithDefaults() *GhostAlertRuleResponseAllOfTarget`

NewGhostAlertRuleResponseAllOfTargetWithDefaults instantiates a new GhostAlertRuleResponseAllOfTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *GhostAlertRuleResponseAllOfTarget) GetTargetType() AlertTargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *GhostAlertRuleResponseAllOfTarget) GetTargetTypeOk() (*AlertTargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *GhostAlertRuleResponseAllOfTarget) SetTargetType(v AlertTargetType)`

SetTargetType sets TargetType field to given value.


### GetTargetId

`func (o *GhostAlertRuleResponseAllOfTarget) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *GhostAlertRuleResponseAllOfTarget) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *GhostAlertRuleResponseAllOfTarget) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetService

`func (o *GhostAlertRuleResponseAllOfTarget) GetService() ServiceLightResponse`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GhostAlertRuleResponseAllOfTarget) GetServiceOk() (*ServiceLightResponse, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GhostAlertRuleResponseAllOfTarget) SetService(v ServiceLightResponse)`

SetService sets Service field to given value.

### HasService

`func (o *GhostAlertRuleResponseAllOfTarget) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


