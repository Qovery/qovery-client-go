# HelmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**AutoPreview** | **bool** | Indicates if the &#39;environment preview option&#39; is enabled.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | 
**AutoDeploy** | **bool** | Specify if the service will be automatically updated after receiving a new image tag or a new commit according to the source type.   | 
**Source** | [**HelmRequestAllOfSource**](HelmRequestAllOfSource.md) |  | 
**Arguments** | Pointer to **[]string** | The extra arguments to pass to helm | [optional] 
**AllowClusterWideResources** | Pointer to **bool** | If we should allow the chart to deploy object outside his specified namespace. Setting this flag to true, requires special rights  | [optional] [default to false]
**ValuesOverride** | Pointer to [**HelmRequestAllOfValuesOverride**](HelmRequestAllOfValuesOverride.md) |  | [optional] 

## Methods

### NewHelmResponse

`func NewHelmResponse(id string, createdAt time.Time, environment ReferenceObject, name string, autoPreview bool, autoDeploy bool, source HelmRequestAllOfSource, ) *HelmResponse`

NewHelmResponse instantiates a new HelmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmResponseWithDefaults

`func NewHelmResponseWithDefaults() *HelmResponse`

NewHelmResponseWithDefaults instantiates a new HelmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HelmResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelmResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelmResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HelmResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HelmResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HelmResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HelmResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HelmResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HelmResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HelmResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *HelmResponse) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *HelmResponse) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *HelmResponse) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *HelmResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HelmResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelmResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelmResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelmResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoPreview

`func (o *HelmResponse) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *HelmResponse) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *HelmResponse) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.


### GetAutoDeploy

`func (o *HelmResponse) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *HelmResponse) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *HelmResponse) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.


### GetSource

`func (o *HelmResponse) GetSource() HelmRequestAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HelmResponse) GetSourceOk() (*HelmRequestAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HelmResponse) SetSource(v HelmRequestAllOfSource)`

SetSource sets Source field to given value.


### GetArguments

`func (o *HelmResponse) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *HelmResponse) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *HelmResponse) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *HelmResponse) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetAllowClusterWideResources

`func (o *HelmResponse) GetAllowClusterWideResources() bool`

GetAllowClusterWideResources returns the AllowClusterWideResources field if non-nil, zero value otherwise.

### GetAllowClusterWideResourcesOk

`func (o *HelmResponse) GetAllowClusterWideResourcesOk() (*bool, bool)`

GetAllowClusterWideResourcesOk returns a tuple with the AllowClusterWideResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClusterWideResources

`func (o *HelmResponse) SetAllowClusterWideResources(v bool)`

SetAllowClusterWideResources sets AllowClusterWideResources field to given value.

### HasAllowClusterWideResources

`func (o *HelmResponse) HasAllowClusterWideResources() bool`

HasAllowClusterWideResources returns a boolean if a field has been set.

### GetValuesOverride

`func (o *HelmResponse) GetValuesOverride() HelmRequestAllOfValuesOverride`

GetValuesOverride returns the ValuesOverride field if non-nil, zero value otherwise.

### GetValuesOverrideOk

`func (o *HelmResponse) GetValuesOverrideOk() (*HelmRequestAllOfValuesOverride, bool)`

GetValuesOverrideOk returns a tuple with the ValuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesOverride

`func (o *HelmResponse) SetValuesOverride(v HelmRequestAllOfValuesOverride)`

SetValuesOverride sets ValuesOverride field to given value.

### HasValuesOverride

`func (o *HelmResponse) HasValuesOverride() bool`

HasValuesOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

