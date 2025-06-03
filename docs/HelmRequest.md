# HelmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]HelmPortRequestPortsInner**](HelmPortRequestPortsInner.md) |  | [optional] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**TimeoutSec** | Pointer to **int32** | Maximum number of seconds allowed for helm to run before killing it and mark it as failed  | [optional] [default to 600]
**AutoPreview** | Pointer to **NullableBool** | Indicates if the &#39;environment preview option&#39; is enabled.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called or when a new commit is updated. If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | [optional] 
**AutoDeploy** | **bool** | Specify if the helm will be automatically updated after receiving a new image tag or a new commit according to the source type.  | 
**Source** | [**HelmRequestAllOfSource**](HelmRequestAllOfSource.md) |  | 
**Arguments** | **[]string** | The extra arguments to pass to helm | 
**AllowClusterWideResources** | Pointer to **bool** | If we should allow the chart to deploy object outside his specified namespace. Setting this flag to true, requires special rights  | [optional] [default to false]
**ValuesOverride** | [**HelmRequestAllOfValuesOverride**](HelmRequestAllOfValuesOverride.md) |  | 
**IconUri** | Pointer to **string** | Icon URI representing the helm service. | [optional] 

## Methods

### NewHelmRequest

`func NewHelmRequest(name string, autoDeploy bool, source HelmRequestAllOfSource, arguments []string, valuesOverride HelmRequestAllOfValuesOverride, ) *HelmRequest`

NewHelmRequest instantiates a new HelmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRequestWithDefaults

`func NewHelmRequestWithDefaults() *HelmRequest`

NewHelmRequestWithDefaults instantiates a new HelmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *HelmRequest) GetPorts() []HelmPortRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HelmRequest) GetPortsOk() (*[]HelmPortRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HelmRequest) SetPorts(v []HelmPortRequestPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HelmRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetName

`func (o *HelmRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HelmRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelmRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelmRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelmRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *HelmRequest) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *HelmRequest) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *HelmRequest) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *HelmRequest) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetAutoPreview

`func (o *HelmRequest) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *HelmRequest) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *HelmRequest) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *HelmRequest) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### SetAutoPreviewNil

`func (o *HelmRequest) SetAutoPreviewNil(b bool)`

 SetAutoPreviewNil sets the value for AutoPreview to be an explicit nil

### UnsetAutoPreview
`func (o *HelmRequest) UnsetAutoPreview()`

UnsetAutoPreview ensures that no value is present for AutoPreview, not even an explicit nil
### GetAutoDeploy

`func (o *HelmRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *HelmRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *HelmRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.


### GetSource

`func (o *HelmRequest) GetSource() HelmRequestAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HelmRequest) GetSourceOk() (*HelmRequestAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HelmRequest) SetSource(v HelmRequestAllOfSource)`

SetSource sets Source field to given value.


### GetArguments

`func (o *HelmRequest) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *HelmRequest) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *HelmRequest) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetAllowClusterWideResources

`func (o *HelmRequest) GetAllowClusterWideResources() bool`

GetAllowClusterWideResources returns the AllowClusterWideResources field if non-nil, zero value otherwise.

### GetAllowClusterWideResourcesOk

`func (o *HelmRequest) GetAllowClusterWideResourcesOk() (*bool, bool)`

GetAllowClusterWideResourcesOk returns a tuple with the AllowClusterWideResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClusterWideResources

`func (o *HelmRequest) SetAllowClusterWideResources(v bool)`

SetAllowClusterWideResources sets AllowClusterWideResources field to given value.

### HasAllowClusterWideResources

`func (o *HelmRequest) HasAllowClusterWideResources() bool`

HasAllowClusterWideResources returns a boolean if a field has been set.

### GetValuesOverride

`func (o *HelmRequest) GetValuesOverride() HelmRequestAllOfValuesOverride`

GetValuesOverride returns the ValuesOverride field if non-nil, zero value otherwise.

### GetValuesOverrideOk

`func (o *HelmRequest) GetValuesOverrideOk() (*HelmRequestAllOfValuesOverride, bool)`

GetValuesOverrideOk returns a tuple with the ValuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesOverride

`func (o *HelmRequest) SetValuesOverride(v HelmRequestAllOfValuesOverride)`

SetValuesOverride sets ValuesOverride field to given value.


### GetIconUri

`func (o *HelmRequest) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *HelmRequest) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *HelmRequest) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *HelmRequest) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


