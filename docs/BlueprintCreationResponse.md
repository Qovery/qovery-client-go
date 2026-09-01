# BlueprintCreationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CatalogUrl** | **string** | URL to the blueprint catalog entry | 
**Tag** | **string** |  | 
**EnvironmentId** | **string** |  | 
**DeploymentId** | **string** | Identifier of the dispatch started by this creation. Resolve its status with GET /blueprint/{blueprintId}, and match it against &#x60;latest_deployment.id&#x60; there to confirm that dispatch is this one rather than a later re-dispatch. | 
**ExecutionId** | **string** | Engine execution identifier for this dispatch | 

## Methods

### NewBlueprintCreationResponse

`func NewBlueprintCreationResponse(id string, catalogUrl string, tag string, environmentId string, deploymentId string, executionId string, ) *BlueprintCreationResponse`

NewBlueprintCreationResponse instantiates a new BlueprintCreationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCreationResponseWithDefaults

`func NewBlueprintCreationResponseWithDefaults() *BlueprintCreationResponse`

NewBlueprintCreationResponseWithDefaults instantiates a new BlueprintCreationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlueprintCreationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlueprintCreationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlueprintCreationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCatalogUrl

`func (o *BlueprintCreationResponse) GetCatalogUrl() string`

GetCatalogUrl returns the CatalogUrl field if non-nil, zero value otherwise.

### GetCatalogUrlOk

`func (o *BlueprintCreationResponse) GetCatalogUrlOk() (*string, bool)`

GetCatalogUrlOk returns a tuple with the CatalogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUrl

`func (o *BlueprintCreationResponse) SetCatalogUrl(v string)`

SetCatalogUrl sets CatalogUrl field to given value.


### GetTag

`func (o *BlueprintCreationResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BlueprintCreationResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BlueprintCreationResponse) SetTag(v string)`

SetTag sets Tag field to given value.


### GetEnvironmentId

`func (o *BlueprintCreationResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BlueprintCreationResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BlueprintCreationResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDeploymentId

`func (o *BlueprintCreationResponse) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *BlueprintCreationResponse) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *BlueprintCreationResponse) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetExecutionId

`func (o *BlueprintCreationResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BlueprintCreationResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BlueprintCreationResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


