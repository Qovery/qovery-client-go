# AgentActivityTokenPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** | Start of the hour or day. Buckets are cut in UTC; empty buckets are omitted. | 
**Actions** | **int64** |  | 
**Denials** | **int64** |  | 

## Methods

### NewAgentActivityTokenPoint

`func NewAgentActivityTokenPoint(start time.Time, actions int64, denials int64, ) *AgentActivityTokenPoint`

NewAgentActivityTokenPoint instantiates a new AgentActivityTokenPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityTokenPointWithDefaults

`func NewAgentActivityTokenPointWithDefaults() *AgentActivityTokenPoint`

NewAgentActivityTokenPointWithDefaults instantiates a new AgentActivityTokenPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *AgentActivityTokenPoint) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AgentActivityTokenPoint) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AgentActivityTokenPoint) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetActions

`func (o *AgentActivityTokenPoint) GetActions() int64`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AgentActivityTokenPoint) GetActionsOk() (*int64, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AgentActivityTokenPoint) SetActions(v int64)`

SetActions sets Actions field to given value.


### GetDenials

`func (o *AgentActivityTokenPoint) GetDenials() int64`

GetDenials returns the Denials field if non-nil, zero value otherwise.

### GetDenialsOk

`func (o *AgentActivityTokenPoint) GetDenialsOk() (*int64, bool)`

GetDenialsOk returns a tuple with the Denials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenials

`func (o *AgentActivityTokenPoint) SetDenials(v int64)`

SetDenials sets Denials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


