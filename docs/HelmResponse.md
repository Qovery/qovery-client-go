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
**TimeoutSec** | Pointer to **int32** | Maximum number of seconds allowed for helm to run before killing it and mark it as failed  | [optional] [default to 600]
**AutoPreview** | **bool** | Indicates if the &#39;environment preview option&#39; is enabled.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | 
**AutoDeploy** | **bool** | Specify if the service will be automatically updated after receiving a new image tag or a new commit according to the source type.  | 
**Ports** | Pointer to [**[]HelmResponseAllOfPorts**](HelmResponseAllOfPorts.md) |  | [optional] 
**Source** | [**HelmResponseAllOfSource**](HelmResponseAllOfSource.md) |  | 
**Arguments** | **[]string** | The extra arguments to pass to helm | 
**AllowClusterWideResources** | **bool** | If we should allow the chart to deploy object outside his specified namespace. Setting this flag to true, requires special rights  | [default to false]
**ValuesOverride** | [**HelmResponseAllOfValuesOverride**](HelmResponseAllOfValuesOverride.md) |  | 
**IconUri** | **string** | Icon URI representing the helm service. | 
**ServiceType** | Pointer to [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | [optional] 

## Methods

### NewHelmResponse

`func NewHelmResponse(id string, createdAt time.Time, environment ReferenceObject, name string, autoPreview bool, autoDeploy bool, source HelmResponseAllOfSource, arguments []string, allowClusterWideResources bool, valuesOverride HelmResponseAllOfValuesOverride, iconUri string, ) *HelmResponse`

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

### GetTimeoutSec

`func (o *HelmResponse) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *HelmResponse) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *HelmResponse) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *HelmResponse) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

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


### GetPorts

`func (o *HelmResponse) GetPorts() []HelmResponseAllOfPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HelmResponse) GetPortsOk() (*[]HelmResponseAllOfPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HelmResponse) SetPorts(v []HelmResponseAllOfPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HelmResponse) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSource

`func (o *HelmResponse) GetSource() HelmResponseAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HelmResponse) GetSourceOk() (*HelmResponseAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HelmResponse) SetSource(v HelmResponseAllOfSource)`

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


### GetValuesOverride

`func (o *HelmResponse) GetValuesOverride() HelmResponseAllOfValuesOverride`

GetValuesOverride returns the ValuesOverride field if non-nil, zero value otherwise.

### GetValuesOverrideOk

`func (o *HelmResponse) GetValuesOverrideOk() (*HelmResponseAllOfValuesOverride, bool)`

GetValuesOverrideOk returns a tuple with the ValuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesOverride

`func (o *HelmResponse) SetValuesOverride(v HelmResponseAllOfValuesOverride)`

SetValuesOverride sets ValuesOverride field to given value.


### GetIconUri

`func (o *HelmResponse) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *HelmResponse) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *HelmResponse) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.


### GetServiceType

`func (o *HelmResponse) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *HelmResponse) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *HelmResponse) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *HelmResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


