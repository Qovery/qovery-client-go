# LifecycleTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**SourceUrl** | **string** | location of the template | 
**CloudProvider** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | 
**Events** | [**[]LifecycleTemplateResponseEventsInner**](LifecycleTemplateResponseEventsInner.md) | lis of pre-defined command for each event | 
**MaxDurationInSec** | **int32** | Job max allowed duration in seconds. After this allowed time, the job is going to be killed. | 
**Resources** | [**LifecycleTemplateResponseResources**](LifecycleTemplateResponseResources.md) |  | 
**Variables** | [**[]LifecycleTemplateResponseVariablesInner**](LifecycleTemplateResponseVariablesInner.md) | Variables to inject at the creation of this lifecycle job | 
**Dockerfile** | **string** | Dockerfile of the template | 

## Methods

### NewLifecycleTemplateResponse

`func NewLifecycleTemplateResponse(id string, name string, description string, sourceUrl string, cloudProvider CloudProviderEnum, events []LifecycleTemplateResponseEventsInner, maxDurationInSec int32, resources LifecycleTemplateResponseResources, variables []LifecycleTemplateResponseVariablesInner, dockerfile string, ) *LifecycleTemplateResponse`

NewLifecycleTemplateResponse instantiates a new LifecycleTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleTemplateResponseWithDefaults

`func NewLifecycleTemplateResponseWithDefaults() *LifecycleTemplateResponse`

NewLifecycleTemplateResponseWithDefaults instantiates a new LifecycleTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LifecycleTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LifecycleTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LifecycleTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LifecycleTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LifecycleTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LifecycleTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LifecycleTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LifecycleTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LifecycleTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSourceUrl

`func (o *LifecycleTemplateResponse) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *LifecycleTemplateResponse) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *LifecycleTemplateResponse) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetCloudProvider

`func (o *LifecycleTemplateResponse) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *LifecycleTemplateResponse) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *LifecycleTemplateResponse) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.


### GetEvents

`func (o *LifecycleTemplateResponse) GetEvents() []LifecycleTemplateResponseEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LifecycleTemplateResponse) GetEventsOk() (*[]LifecycleTemplateResponseEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LifecycleTemplateResponse) SetEvents(v []LifecycleTemplateResponseEventsInner)`

SetEvents sets Events field to given value.


### GetMaxDurationInSec

`func (o *LifecycleTemplateResponse) GetMaxDurationInSec() int32`

GetMaxDurationInSec returns the MaxDurationInSec field if non-nil, zero value otherwise.

### GetMaxDurationInSecOk

`func (o *LifecycleTemplateResponse) GetMaxDurationInSecOk() (*int32, bool)`

GetMaxDurationInSecOk returns a tuple with the MaxDurationInSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInSec

`func (o *LifecycleTemplateResponse) SetMaxDurationInSec(v int32)`

SetMaxDurationInSec sets MaxDurationInSec field to given value.


### GetResources

`func (o *LifecycleTemplateResponse) GetResources() LifecycleTemplateResponseResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LifecycleTemplateResponse) GetResourcesOk() (*LifecycleTemplateResponseResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LifecycleTemplateResponse) SetResources(v LifecycleTemplateResponseResources)`

SetResources sets Resources field to given value.


### GetVariables

`func (o *LifecycleTemplateResponse) GetVariables() []LifecycleTemplateResponseVariablesInner`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *LifecycleTemplateResponse) GetVariablesOk() (*[]LifecycleTemplateResponseVariablesInner, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *LifecycleTemplateResponse) SetVariables(v []LifecycleTemplateResponseVariablesInner)`

SetVariables sets Variables field to given value.


### GetDockerfile

`func (o *LifecycleTemplateResponse) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *LifecycleTemplateResponse) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *LifecycleTemplateResponse) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


