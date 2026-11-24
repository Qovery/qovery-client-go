# AgentActionsByType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The audit event type, as in &#x60;OrganizationEventType&#x60;. | 
**TargetType** | **string** | The audit target type, as in &#x60;OrganizationEventTargetType&#x60;. | 
**Count** | **int64** |  | 
**Denied** | **int64** | How many of them were refused. | 

## Methods

### NewAgentActionsByType

`func NewAgentActionsByType(event string, targetType string, count int64, denied int64, ) *AgentActionsByType`

NewAgentActionsByType instantiates a new AgentActionsByType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActionsByTypeWithDefaults

`func NewAgentActionsByTypeWithDefaults() *AgentActionsByType`

NewAgentActionsByTypeWithDefaults instantiates a new AgentActionsByType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AgentActionsByType) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AgentActionsByType) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AgentActionsByType) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTargetType

`func (o *AgentActionsByType) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AgentActionsByType) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AgentActionsByType) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetCount

`func (o *AgentActionsByType) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AgentActionsByType) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AgentActionsByType) SetCount(v int64)`

SetCount sets Count field to given value.


### GetDenied

`func (o *AgentActionsByType) GetDenied() int64`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *AgentActionsByType) GetDeniedOk() (*int64, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *AgentActionsByType) SetDenied(v int64)`

SetDenied sets Denied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


