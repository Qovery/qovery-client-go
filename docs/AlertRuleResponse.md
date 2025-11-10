# AlertRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrganizationId** | **string** | Organization identifier | 
**ClusterId** | **string** |  Cluster identifier | 
**Name** | **string** | Name of the alert rule  | 
**Description** | **string** | Description of what the alert monitors | 
**PromqlExpr** | **string** | PromQL expression to evaluate | 
**ForDuration** | **string** | Duration the condition must be true before firing (ISO-8601 duration format) | 
**Severity** | [**AlertSeverity**](AlertSeverity.md) |  | 
**Enabled** | **bool** | Whether the alert rule is enabled | 
**AlertReceiverIds** | **[]string** | List of alert receiver IDs to send notifications to | 
**Presentation** | [**AlertPresentationResponse**](AlertPresentationResponse.md) |  | 
**Target** | [**AlertTarget**](AlertTarget.md) |  | 
**State** | [**AlertRuleState**](AlertRuleState.md) |  | 
**IsUpToDate** | **bool** | Indicates whether the current version of the alert has been synced with the alerting system. If false, an outdated version is currently deployed. | 

## Methods

### NewAlertRuleResponse

`func NewAlertRuleResponse(id string, createdAt time.Time, organizationId string, clusterId string, name string, description string, promqlExpr string, forDuration string, severity AlertSeverity, enabled bool, alertReceiverIds []string, presentation AlertPresentationResponse, target AlertTarget, state AlertRuleState, isUpToDate bool, ) *AlertRuleResponse`

NewAlertRuleResponse instantiates a new AlertRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleResponseWithDefaults

`func NewAlertRuleResponseWithDefaults() *AlertRuleResponse`

NewAlertRuleResponseWithDefaults instantiates a new AlertRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AlertRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AlertRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AlertRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AlertRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AlertRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AlertRuleResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertRuleResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertRuleResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetClusterId

`func (o *AlertRuleResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AlertRuleResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AlertRuleResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *AlertRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPromqlExpr

`func (o *AlertRuleResponse) GetPromqlExpr() string`

GetPromqlExpr returns the PromqlExpr field if non-nil, zero value otherwise.

### GetPromqlExprOk

`func (o *AlertRuleResponse) GetPromqlExprOk() (*string, bool)`

GetPromqlExprOk returns a tuple with the PromqlExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromqlExpr

`func (o *AlertRuleResponse) SetPromqlExpr(v string)`

SetPromqlExpr sets PromqlExpr field to given value.


### GetForDuration

`func (o *AlertRuleResponse) GetForDuration() string`

GetForDuration returns the ForDuration field if non-nil, zero value otherwise.

### GetForDurationOk

`func (o *AlertRuleResponse) GetForDurationOk() (*string, bool)`

GetForDurationOk returns a tuple with the ForDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDuration

`func (o *AlertRuleResponse) SetForDuration(v string)`

SetForDuration sets ForDuration field to given value.


### GetSeverity

`func (o *AlertRuleResponse) GetSeverity() AlertSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleResponse) GetSeverityOk() (*AlertSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleResponse) SetSeverity(v AlertSeverity)`

SetSeverity sets Severity field to given value.


### GetEnabled

`func (o *AlertRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAlertReceiverIds

`func (o *AlertRuleResponse) GetAlertReceiverIds() []string`

GetAlertReceiverIds returns the AlertReceiverIds field if non-nil, zero value otherwise.

### GetAlertReceiverIdsOk

`func (o *AlertRuleResponse) GetAlertReceiverIdsOk() (*[]string, bool)`

GetAlertReceiverIdsOk returns a tuple with the AlertReceiverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertReceiverIds

`func (o *AlertRuleResponse) SetAlertReceiverIds(v []string)`

SetAlertReceiverIds sets AlertReceiverIds field to given value.


### GetPresentation

`func (o *AlertRuleResponse) GetPresentation() AlertPresentationResponse`

GetPresentation returns the Presentation field if non-nil, zero value otherwise.

### GetPresentationOk

`func (o *AlertRuleResponse) GetPresentationOk() (*AlertPresentationResponse, bool)`

GetPresentationOk returns a tuple with the Presentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentation

`func (o *AlertRuleResponse) SetPresentation(v AlertPresentationResponse)`

SetPresentation sets Presentation field to given value.


### GetTarget

`func (o *AlertRuleResponse) GetTarget() AlertTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AlertRuleResponse) GetTargetOk() (*AlertTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AlertRuleResponse) SetTarget(v AlertTarget)`

SetTarget sets Target field to given value.


### GetState

`func (o *AlertRuleResponse) GetState() AlertRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AlertRuleResponse) GetStateOk() (*AlertRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AlertRuleResponse) SetState(v AlertRuleState)`

SetState sets State field to given value.


### GetIsUpToDate

`func (o *AlertRuleResponse) GetIsUpToDate() bool`

GetIsUpToDate returns the IsUpToDate field if non-nil, zero value otherwise.

### GetIsUpToDateOk

`func (o *AlertRuleResponse) GetIsUpToDateOk() (*bool, bool)`

GetIsUpToDateOk returns a tuple with the IsUpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpToDate

`func (o *AlertRuleResponse) SetIsUpToDate(v bool)`

SetIsUpToDate sets IsUpToDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


