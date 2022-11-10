# EnvironmentLogsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Transmitter** | Pointer to [**EnvironmentLogsDetailsTransmitter**](EnvironmentLogsDetailsTransmitter.md) |  | [optional] 
**Stage** | Pointer to [**EnvironmentLogsDetailsStage**](EnvironmentLogsDetailsStage.md) |  | [optional] 

## Methods

### NewEnvironmentLogsDetails

`func NewEnvironmentLogsDetails() *EnvironmentLogsDetails`

NewEnvironmentLogsDetails instantiates a new EnvironmentLogsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLogsDetailsWithDefaults

`func NewEnvironmentLogsDetailsWithDefaults() *EnvironmentLogsDetails`

NewEnvironmentLogsDetailsWithDefaults instantiates a new EnvironmentLogsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *EnvironmentLogsDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *EnvironmentLogsDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *EnvironmentLogsDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *EnvironmentLogsDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetClusterId

`func (o *EnvironmentLogsDetails) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnvironmentLogsDetails) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnvironmentLogsDetails) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnvironmentLogsDetails) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetExecutionId

`func (o *EnvironmentLogsDetails) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EnvironmentLogsDetails) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EnvironmentLogsDetails) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EnvironmentLogsDetails) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetTransmitter

`func (o *EnvironmentLogsDetails) GetTransmitter() EnvironmentLogsDetailsTransmitter`

GetTransmitter returns the Transmitter field if non-nil, zero value otherwise.

### GetTransmitterOk

`func (o *EnvironmentLogsDetails) GetTransmitterOk() (*EnvironmentLogsDetailsTransmitter, bool)`

GetTransmitterOk returns a tuple with the Transmitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitter

`func (o *EnvironmentLogsDetails) SetTransmitter(v EnvironmentLogsDetailsTransmitter)`

SetTransmitter sets Transmitter field to given value.

### HasTransmitter

`func (o *EnvironmentLogsDetails) HasTransmitter() bool`

HasTransmitter returns a boolean if a field has been set.

### GetStage

`func (o *EnvironmentLogsDetails) GetStage() EnvironmentLogsDetailsStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *EnvironmentLogsDetails) GetStageOk() (*EnvironmentLogsDetailsStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *EnvironmentLogsDetails) SetStage(v EnvironmentLogsDetailsStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *EnvironmentLogsDetails) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


