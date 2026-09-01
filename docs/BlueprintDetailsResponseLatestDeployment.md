# BlueprintDetailsResponseLatestDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExecutionId** | **string** | Engine execution identifier for this dispatch | 
**Status** | **string** | Status of the dispatch. Only the states a blueprint dispatch can reach, which is a subset of the deployment statuses used elsewhere in the API. | 
**StartedAt** | **time.Time** |  | 
**TerminatedAt** | **NullableTime** | When the dispatch reached a terminal state. Null while it is still running. | 
**ErrorMessage** | **NullableString** | The engine&#39;s failure message. Null unless the dispatch failed. | 

## Methods

### NewBlueprintDetailsResponseLatestDeployment

`func NewBlueprintDetailsResponseLatestDeployment(id string, executionId string, status string, startedAt time.Time, terminatedAt NullableTime, errorMessage NullableString, ) *BlueprintDetailsResponseLatestDeployment`

NewBlueprintDetailsResponseLatestDeployment instantiates a new BlueprintDetailsResponseLatestDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDetailsResponseLatestDeploymentWithDefaults

`func NewBlueprintDetailsResponseLatestDeploymentWithDefaults() *BlueprintDetailsResponseLatestDeployment`

NewBlueprintDetailsResponseLatestDeploymentWithDefaults instantiates a new BlueprintDetailsResponseLatestDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlueprintDetailsResponseLatestDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlueprintDetailsResponseLatestDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlueprintDetailsResponseLatestDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetExecutionId

`func (o *BlueprintDetailsResponseLatestDeployment) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BlueprintDetailsResponseLatestDeployment) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BlueprintDetailsResponseLatestDeployment) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetStatus

`func (o *BlueprintDetailsResponseLatestDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintDetailsResponseLatestDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintDetailsResponseLatestDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *BlueprintDetailsResponseLatestDeployment) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BlueprintDetailsResponseLatestDeployment) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BlueprintDetailsResponseLatestDeployment) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTerminatedAt

`func (o *BlueprintDetailsResponseLatestDeployment) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *BlueprintDetailsResponseLatestDeployment) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *BlueprintDetailsResponseLatestDeployment) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.


### SetTerminatedAtNil

`func (o *BlueprintDetailsResponseLatestDeployment) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *BlueprintDetailsResponseLatestDeployment) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetErrorMessage

`func (o *BlueprintDetailsResponseLatestDeployment) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BlueprintDetailsResponseLatestDeployment) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BlueprintDetailsResponseLatestDeployment) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *BlueprintDetailsResponseLatestDeployment) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BlueprintDetailsResponseLatestDeployment) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


