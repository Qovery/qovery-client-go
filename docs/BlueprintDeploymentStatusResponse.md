# BlueprintDeploymentStatusResponse

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

### NewBlueprintDeploymentStatusResponse

`func NewBlueprintDeploymentStatusResponse(id string, executionId string, status string, startedAt time.Time, terminatedAt NullableTime, errorMessage NullableString, ) *BlueprintDeploymentStatusResponse`

NewBlueprintDeploymentStatusResponse instantiates a new BlueprintDeploymentStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDeploymentStatusResponseWithDefaults

`func NewBlueprintDeploymentStatusResponseWithDefaults() *BlueprintDeploymentStatusResponse`

NewBlueprintDeploymentStatusResponseWithDefaults instantiates a new BlueprintDeploymentStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlueprintDeploymentStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlueprintDeploymentStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlueprintDeploymentStatusResponse) SetId(v string)`

SetId sets Id field to given value.


### GetExecutionId

`func (o *BlueprintDeploymentStatusResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BlueprintDeploymentStatusResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BlueprintDeploymentStatusResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetStatus

`func (o *BlueprintDeploymentStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintDeploymentStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintDeploymentStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *BlueprintDeploymentStatusResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BlueprintDeploymentStatusResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BlueprintDeploymentStatusResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTerminatedAt

`func (o *BlueprintDeploymentStatusResponse) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *BlueprintDeploymentStatusResponse) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *BlueprintDeploymentStatusResponse) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.


### SetTerminatedAtNil

`func (o *BlueprintDeploymentStatusResponse) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *BlueprintDeploymentStatusResponse) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetErrorMessage

`func (o *BlueprintDeploymentStatusResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BlueprintDeploymentStatusResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BlueprintDeploymentStatusResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *BlueprintDeploymentStatusResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BlueprintDeploymentStatusResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


