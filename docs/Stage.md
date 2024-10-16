# Stage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | stage name | 
**Steps** | Pointer to [**StageStepMetrics**](StageStepMetrics.md) |  | [optional] 
**Description** | **string** |  | 
**Status** | Pointer to [**StageStatusEnum**](StageStatusEnum.md) |  | [optional] 

## Methods

### NewStage

`func NewStage(id string, name string, description string, ) *Stage`

NewStage instantiates a new Stage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageWithDefaults

`func NewStageWithDefaults() *Stage`

NewStageWithDefaults instantiates a new Stage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stage) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Stage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stage) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *Stage) GetSteps() StageStepMetrics`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Stage) GetStepsOk() (*StageStepMetrics, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Stage) SetSteps(v StageStepMetrics)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Stage) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetDescription

`func (o *Stage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Stage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Stage) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *Stage) GetStatus() StageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Stage) GetStatusOk() (*StageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Stage) SetStatus(v StageStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Stage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


