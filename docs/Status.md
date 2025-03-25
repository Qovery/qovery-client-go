# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**ServiceDeploymentStatus** | [**ServiceDeploymentStatusEnum**](ServiceDeploymentStatusEnum.md) |  | 
**LastDeploymentDate** | Pointer to **time.Time** |  | [optional] 
**IsPartLastDeployment** | Pointer to **bool** |  | [optional] 
**Steps** | Pointer to [**ServiceStepMetrics**](ServiceStepMetrics.md) |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**StatusDetails** | [**StatusDetails**](StatusDetails.md) |  | 
**DeploymentRequestId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStatus

`func NewStatus(id string, state StateEnum, serviceDeploymentStatus ServiceDeploymentStatusEnum, statusDetails StatusDetails, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Status) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Status) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Status) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *Status) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Status) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Status) SetState(v StateEnum)`

SetState sets State field to given value.


### GetServiceDeploymentStatus

`func (o *Status) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum`

GetServiceDeploymentStatus returns the ServiceDeploymentStatus field if non-nil, zero value otherwise.

### GetServiceDeploymentStatusOk

`func (o *Status) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool)`

GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeploymentStatus

`func (o *Status) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum)`

SetServiceDeploymentStatus sets ServiceDeploymentStatus field to given value.


### GetLastDeploymentDate

`func (o *Status) GetLastDeploymentDate() time.Time`

GetLastDeploymentDate returns the LastDeploymentDate field if non-nil, zero value otherwise.

### GetLastDeploymentDateOk

`func (o *Status) GetLastDeploymentDateOk() (*time.Time, bool)`

GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentDate

`func (o *Status) SetLastDeploymentDate(v time.Time)`

SetLastDeploymentDate sets LastDeploymentDate field to given value.

### HasLastDeploymentDate

`func (o *Status) HasLastDeploymentDate() bool`

HasLastDeploymentDate returns a boolean if a field has been set.

### GetIsPartLastDeployment

`func (o *Status) GetIsPartLastDeployment() bool`

GetIsPartLastDeployment returns the IsPartLastDeployment field if non-nil, zero value otherwise.

### GetIsPartLastDeploymentOk

`func (o *Status) GetIsPartLastDeploymentOk() (*bool, bool)`

GetIsPartLastDeploymentOk returns a tuple with the IsPartLastDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartLastDeployment

`func (o *Status) SetIsPartLastDeployment(v bool)`

SetIsPartLastDeployment sets IsPartLastDeployment field to given value.

### HasIsPartLastDeployment

`func (o *Status) HasIsPartLastDeployment() bool`

HasIsPartLastDeployment returns a boolean if a field has been set.

### GetSteps

`func (o *Status) GetSteps() ServiceStepMetrics`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Status) GetStepsOk() (*ServiceStepMetrics, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Status) SetSteps(v ServiceStepMetrics)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Status) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetExecutionId

`func (o *Status) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Status) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Status) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Status) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetStatusDetails

`func (o *Status) GetStatusDetails() StatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Status) GetStatusDetailsOk() (*StatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Status) SetStatusDetails(v StatusDetails)`

SetStatusDetails sets StatusDetails field to given value.


### GetDeploymentRequestId

`func (o *Status) GetDeploymentRequestId() string`

GetDeploymentRequestId returns the DeploymentRequestId field if non-nil, zero value otherwise.

### GetDeploymentRequestIdOk

`func (o *Status) GetDeploymentRequestIdOk() (*string, bool)`

GetDeploymentRequestIdOk returns a tuple with the DeploymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRequestId

`func (o *Status) SetDeploymentRequestId(v string)`

SetDeploymentRequestId sets DeploymentRequestId field to given value.

### HasDeploymentRequestId

`func (o *Status) HasDeploymentRequestId() bool`

HasDeploymentRequestId returns a boolean if a field has been set.

### SetDeploymentRequestIdNil

`func (o *Status) SetDeploymentRequestIdNil(b bool)`

 SetDeploymentRequestIdNil sets the value for DeploymentRequestId to be an explicit nil

### UnsetDeploymentRequestId
`func (o *Status) UnsetDeploymentRequestId()`

UnsetDeploymentRequestId ensures that no value is present for DeploymentRequestId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


