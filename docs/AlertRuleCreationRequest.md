# AlertRuleCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** | Organization identifier | 
**ClusterId** | **string** |  Cluster identifier where the rule will be deployed | 
**Name** | **string** | Name of the alert rule | 
**Description** | **string** | Description of what the alert monitors  | 
**PromqlExpr** | **string** | PromQL expression to evaluate | 
**ForDuration** | **string** | Duration the condition must be true before firing (ISO-8601 duration format) | 
**Severity** | [**AlertSeverity**](AlertSeverity.md) |  | 
**Enabled** | **bool** | Whether the alert rule is enabled | 
**AlertReceiverIds** | **[]string** | List of alert receiver IDs to send notifications to | 
**Presentation** | [**AlertPresentation**](AlertPresentation.md) |  | 

## Methods

### NewAlertRuleCreationRequest

`func NewAlertRuleCreationRequest(organizationId string, clusterId string, name string, description string, promqlExpr string, forDuration string, severity AlertSeverity, enabled bool, alertReceiverIds []string, presentation AlertPresentation, ) *AlertRuleCreationRequest`

NewAlertRuleCreationRequest instantiates a new AlertRuleCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleCreationRequestWithDefaults

`func NewAlertRuleCreationRequestWithDefaults() *AlertRuleCreationRequest`

NewAlertRuleCreationRequestWithDefaults instantiates a new AlertRuleCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *AlertRuleCreationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertRuleCreationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertRuleCreationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetClusterId

`func (o *AlertRuleCreationRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AlertRuleCreationRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AlertRuleCreationRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *AlertRuleCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRuleCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRuleCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertRuleCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRuleCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRuleCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPromqlExpr

`func (o *AlertRuleCreationRequest) GetPromqlExpr() string`

GetPromqlExpr returns the PromqlExpr field if non-nil, zero value otherwise.

### GetPromqlExprOk

`func (o *AlertRuleCreationRequest) GetPromqlExprOk() (*string, bool)`

GetPromqlExprOk returns a tuple with the PromqlExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromqlExpr

`func (o *AlertRuleCreationRequest) SetPromqlExpr(v string)`

SetPromqlExpr sets PromqlExpr field to given value.


### GetForDuration

`func (o *AlertRuleCreationRequest) GetForDuration() string`

GetForDuration returns the ForDuration field if non-nil, zero value otherwise.

### GetForDurationOk

`func (o *AlertRuleCreationRequest) GetForDurationOk() (*string, bool)`

GetForDurationOk returns a tuple with the ForDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDuration

`func (o *AlertRuleCreationRequest) SetForDuration(v string)`

SetForDuration sets ForDuration field to given value.


### GetSeverity

`func (o *AlertRuleCreationRequest) GetSeverity() AlertSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleCreationRequest) GetSeverityOk() (*AlertSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleCreationRequest) SetSeverity(v AlertSeverity)`

SetSeverity sets Severity field to given value.


### GetEnabled

`func (o *AlertRuleCreationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRuleCreationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRuleCreationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAlertReceiverIds

`func (o *AlertRuleCreationRequest) GetAlertReceiverIds() []string`

GetAlertReceiverIds returns the AlertReceiverIds field if non-nil, zero value otherwise.

### GetAlertReceiverIdsOk

`func (o *AlertRuleCreationRequest) GetAlertReceiverIdsOk() (*[]string, bool)`

GetAlertReceiverIdsOk returns a tuple with the AlertReceiverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertReceiverIds

`func (o *AlertRuleCreationRequest) SetAlertReceiverIds(v []string)`

SetAlertReceiverIds sets AlertReceiverIds field to given value.


### GetPresentation

`func (o *AlertRuleCreationRequest) GetPresentation() AlertPresentation`

GetPresentation returns the Presentation field if non-nil, zero value otherwise.

### GetPresentationOk

`func (o *AlertRuleCreationRequest) GetPresentationOk() (*AlertPresentation, bool)`

GetPresentationOk returns a tuple with the Presentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentation

`func (o *AlertRuleCreationRequest) SetPresentation(v AlertPresentation)`

SetPresentation sets Presentation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


