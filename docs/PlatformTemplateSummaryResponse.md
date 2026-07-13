# PlatformTemplateSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Version** | **string** |  | 
**Status** | [**PlatformTemplateReleaseStatus**](PlatformTemplateReleaseStatus.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Layers** | [**[]PlatformTemplateLayerResponse**](PlatformTemplateLayerResponse.md) |  | 

## Methods

### NewPlatformTemplateSummaryResponse

`func NewPlatformTemplateSummaryResponse(key string, version string, status PlatformTemplateReleaseStatus, layers []PlatformTemplateLayerResponse, ) *PlatformTemplateSummaryResponse`

NewPlatformTemplateSummaryResponse instantiates a new PlatformTemplateSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTemplateSummaryResponseWithDefaults

`func NewPlatformTemplateSummaryResponseWithDefaults() *PlatformTemplateSummaryResponse`

NewPlatformTemplateSummaryResponseWithDefaults instantiates a new PlatformTemplateSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformTemplateSummaryResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformTemplateSummaryResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformTemplateSummaryResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *PlatformTemplateSummaryResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlatformTemplateSummaryResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlatformTemplateSummaryResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *PlatformTemplateSummaryResponse) GetStatus() PlatformTemplateReleaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformTemplateSummaryResponse) GetStatusOk() (*PlatformTemplateReleaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformTemplateSummaryResponse) SetStatus(v PlatformTemplateReleaseStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *PlatformTemplateSummaryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformTemplateSummaryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformTemplateSummaryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformTemplateSummaryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlatformTemplateSummaryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlatformTemplateSummaryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLayers

`func (o *PlatformTemplateSummaryResponse) GetLayers() []PlatformTemplateLayerResponse`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *PlatformTemplateSummaryResponse) GetLayersOk() (*[]PlatformTemplateLayerResponse, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *PlatformTemplateSummaryResponse) SetLayers(v []PlatformTemplateLayerResponse)`

SetLayers sets Layers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


