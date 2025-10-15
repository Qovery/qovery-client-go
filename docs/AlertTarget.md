# AlertTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | [**AlertTargetType**](AlertTargetType.md) |  | 
**TargetId** | **string** |  | 

## Methods

### NewAlertTarget

`func NewAlertTarget(targetType AlertTargetType, targetId string, ) *AlertTarget`

NewAlertTarget instantiates a new AlertTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertTargetWithDefaults

`func NewAlertTargetWithDefaults() *AlertTarget`

NewAlertTargetWithDefaults instantiates a new AlertTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *AlertTarget) GetTargetType() AlertTargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AlertTarget) GetTargetTypeOk() (*AlertTargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AlertTarget) SetTargetType(v AlertTargetType)`

SetTargetType sets TargetType field to given value.


### GetTargetId

`func (o *AlertTarget) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AlertTarget) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AlertTarget) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


