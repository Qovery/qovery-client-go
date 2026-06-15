# BlueprintUpdatePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewId** | **string** | Identifier of the preview execution | 
**ServiceType** | **string** | Type of the underlying service backing this blueprint | 

## Methods

### NewBlueprintUpdatePreviewResponse

`func NewBlueprintUpdatePreviewResponse(previewId string, serviceType string, ) *BlueprintUpdatePreviewResponse`

NewBlueprintUpdatePreviewResponse instantiates a new BlueprintUpdatePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdatePreviewResponseWithDefaults

`func NewBlueprintUpdatePreviewResponseWithDefaults() *BlueprintUpdatePreviewResponse`

NewBlueprintUpdatePreviewResponseWithDefaults instantiates a new BlueprintUpdatePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviewId

`func (o *BlueprintUpdatePreviewResponse) GetPreviewId() string`

GetPreviewId returns the PreviewId field if non-nil, zero value otherwise.

### GetPreviewIdOk

`func (o *BlueprintUpdatePreviewResponse) GetPreviewIdOk() (*string, bool)`

GetPreviewIdOk returns a tuple with the PreviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewId

`func (o *BlueprintUpdatePreviewResponse) SetPreviewId(v string)`

SetPreviewId sets PreviewId field to given value.


### GetServiceType

`func (o *BlueprintUpdatePreviewResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *BlueprintUpdatePreviewResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *BlueprintUpdatePreviewResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


