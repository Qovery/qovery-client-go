# ReferenceObjectStatus

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
**DeploymentRequestsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewReferenceObjectStatus

`func NewReferenceObjectStatus(id string, state StateEnum, serviceDeploymentStatus ServiceDeploymentStatusEnum, statusDetails StatusDetails, ) *ReferenceObjectStatus`

NewReferenceObjectStatus instantiates a new ReferenceObjectStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceObjectStatusWithDefaults

`func NewReferenceObjectStatusWithDefaults() *ReferenceObjectStatus`

NewReferenceObjectStatusWithDefaults instantiates a new ReferenceObjectStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferenceObjectStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferenceObjectStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferenceObjectStatus) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *ReferenceObjectStatus) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReferenceObjectStatus) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReferenceObjectStatus) SetState(v StateEnum)`

SetState sets State field to given value.


### GetServiceDeploymentStatus

`func (o *ReferenceObjectStatus) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum`

GetServiceDeploymentStatus returns the ServiceDeploymentStatus field if non-nil, zero value otherwise.

### GetServiceDeploymentStatusOk

`func (o *ReferenceObjectStatus) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool)`

GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeploymentStatus

`func (o *ReferenceObjectStatus) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum)`

SetServiceDeploymentStatus sets ServiceDeploymentStatus field to given value.


### GetLastDeploymentDate

`func (o *ReferenceObjectStatus) GetLastDeploymentDate() time.Time`

GetLastDeploymentDate returns the LastDeploymentDate field if non-nil, zero value otherwise.

### GetLastDeploymentDateOk

`func (o *ReferenceObjectStatus) GetLastDeploymentDateOk() (*time.Time, bool)`

GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentDate

`func (o *ReferenceObjectStatus) SetLastDeploymentDate(v time.Time)`

SetLastDeploymentDate sets LastDeploymentDate field to given value.

### HasLastDeploymentDate

`func (o *ReferenceObjectStatus) HasLastDeploymentDate() bool`

HasLastDeploymentDate returns a boolean if a field has been set.

### GetIsPartLastDeployment

`func (o *ReferenceObjectStatus) GetIsPartLastDeployment() bool`

GetIsPartLastDeployment returns the IsPartLastDeployment field if non-nil, zero value otherwise.

### GetIsPartLastDeploymentOk

`func (o *ReferenceObjectStatus) GetIsPartLastDeploymentOk() (*bool, bool)`

GetIsPartLastDeploymentOk returns a tuple with the IsPartLastDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartLastDeployment

`func (o *ReferenceObjectStatus) SetIsPartLastDeployment(v bool)`

SetIsPartLastDeployment sets IsPartLastDeployment field to given value.

### HasIsPartLastDeployment

`func (o *ReferenceObjectStatus) HasIsPartLastDeployment() bool`

HasIsPartLastDeployment returns a boolean if a field has been set.

### GetSteps

`func (o *ReferenceObjectStatus) GetSteps() ServiceStepMetrics`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReferenceObjectStatus) GetStepsOk() (*ServiceStepMetrics, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReferenceObjectStatus) SetSteps(v ServiceStepMetrics)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReferenceObjectStatus) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetExecutionId

`func (o *ReferenceObjectStatus) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ReferenceObjectStatus) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ReferenceObjectStatus) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ReferenceObjectStatus) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetStatusDetails

`func (o *ReferenceObjectStatus) GetStatusDetails() StatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *ReferenceObjectStatus) GetStatusDetailsOk() (*StatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *ReferenceObjectStatus) SetStatusDetails(v StatusDetails)`

SetStatusDetails sets StatusDetails field to given value.


### GetDeploymentRequestId

`func (o *ReferenceObjectStatus) GetDeploymentRequestId() string`

GetDeploymentRequestId returns the DeploymentRequestId field if non-nil, zero value otherwise.

### GetDeploymentRequestIdOk

`func (o *ReferenceObjectStatus) GetDeploymentRequestIdOk() (*string, bool)`

GetDeploymentRequestIdOk returns a tuple with the DeploymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRequestId

`func (o *ReferenceObjectStatus) SetDeploymentRequestId(v string)`

SetDeploymentRequestId sets DeploymentRequestId field to given value.

### HasDeploymentRequestId

`func (o *ReferenceObjectStatus) HasDeploymentRequestId() bool`

HasDeploymentRequestId returns a boolean if a field has been set.

### SetDeploymentRequestIdNil

`func (o *ReferenceObjectStatus) SetDeploymentRequestIdNil(b bool)`

 SetDeploymentRequestIdNil sets the value for DeploymentRequestId to be an explicit nil

### UnsetDeploymentRequestId
`func (o *ReferenceObjectStatus) UnsetDeploymentRequestId()`

UnsetDeploymentRequestId ensures that no value is present for DeploymentRequestId, not even an explicit nil
### GetDeploymentRequestsCount

`func (o *ReferenceObjectStatus) GetDeploymentRequestsCount() int32`

GetDeploymentRequestsCount returns the DeploymentRequestsCount field if non-nil, zero value otherwise.

### GetDeploymentRequestsCountOk

`func (o *ReferenceObjectStatus) GetDeploymentRequestsCountOk() (*int32, bool)`

GetDeploymentRequestsCountOk returns a tuple with the DeploymentRequestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRequestsCount

`func (o *ReferenceObjectStatus) SetDeploymentRequestsCount(v int32)`

SetDeploymentRequestsCount sets DeploymentRequestsCount field to given value.

### HasDeploymentRequestsCount

`func (o *ReferenceObjectStatus) HasDeploymentRequestsCount() bool`

HasDeploymentRequestsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


