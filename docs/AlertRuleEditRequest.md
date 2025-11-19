# AlertRuleEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the alert rule | 
**Description** | **string** | Description of what the alert monitors | 
**Tag** | **string** |  | 
**Condition** | [**AlertRuleCondition**](AlertRuleCondition.md) |  | 
**ForDuration** | **string** | Duration the condition must be true before firing (ISO-8601 duration format) | 
**Severity** | [**AlertSeverity**](AlertSeverity.md) |  | 
**Enabled** | **bool** | Whether the alert rule is enabled | 
**AlertReceiverIds** | **[]string** | List of alert receiver IDs to send notifications to | 
**Presentation** | [**AlertPresentation**](AlertPresentation.md) |  | 

## Methods

### NewAlertRuleEditRequest

`func NewAlertRuleEditRequest(name string, description string, tag string, condition AlertRuleCondition, forDuration string, severity AlertSeverity, enabled bool, alertReceiverIds []string, presentation AlertPresentation, ) *AlertRuleEditRequest`

NewAlertRuleEditRequest instantiates a new AlertRuleEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleEditRequestWithDefaults

`func NewAlertRuleEditRequestWithDefaults() *AlertRuleEditRequest`

NewAlertRuleEditRequestWithDefaults instantiates a new AlertRuleEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertRuleEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRuleEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRuleEditRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertRuleEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRuleEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRuleEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *AlertRuleEditRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AlertRuleEditRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AlertRuleEditRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetCondition

`func (o *AlertRuleEditRequest) GetCondition() AlertRuleCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AlertRuleEditRequest) GetConditionOk() (*AlertRuleCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AlertRuleEditRequest) SetCondition(v AlertRuleCondition)`

SetCondition sets Condition field to given value.


### GetForDuration

`func (o *AlertRuleEditRequest) GetForDuration() string`

GetForDuration returns the ForDuration field if non-nil, zero value otherwise.

### GetForDurationOk

`func (o *AlertRuleEditRequest) GetForDurationOk() (*string, bool)`

GetForDurationOk returns a tuple with the ForDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDuration

`func (o *AlertRuleEditRequest) SetForDuration(v string)`

SetForDuration sets ForDuration field to given value.


### GetSeverity

`func (o *AlertRuleEditRequest) GetSeverity() AlertSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleEditRequest) GetSeverityOk() (*AlertSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleEditRequest) SetSeverity(v AlertSeverity)`

SetSeverity sets Severity field to given value.


### GetEnabled

`func (o *AlertRuleEditRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRuleEditRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRuleEditRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAlertReceiverIds

`func (o *AlertRuleEditRequest) GetAlertReceiverIds() []string`

GetAlertReceiverIds returns the AlertReceiverIds field if non-nil, zero value otherwise.

### GetAlertReceiverIdsOk

`func (o *AlertRuleEditRequest) GetAlertReceiverIdsOk() (*[]string, bool)`

GetAlertReceiverIdsOk returns a tuple with the AlertReceiverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertReceiverIds

`func (o *AlertRuleEditRequest) SetAlertReceiverIds(v []string)`

SetAlertReceiverIds sets AlertReceiverIds field to given value.


### GetPresentation

`func (o *AlertRuleEditRequest) GetPresentation() AlertPresentation`

GetPresentation returns the Presentation field if non-nil, zero value otherwise.

### GetPresentationOk

`func (o *AlertRuleEditRequest) GetPresentationOk() (*AlertPresentation, bool)`

GetPresentationOk returns a tuple with the Presentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentation

`func (o *AlertRuleEditRequest) SetPresentation(v AlertPresentation)`

SetPresentation sets Presentation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


