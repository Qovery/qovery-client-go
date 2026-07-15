# PlatformComponentConfigurationPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ComponentKey** | **string** |  | 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) |  | 
**Requirements** | [**[]PlatformComponentInputRequirementResponse**](PlatformComponentInputRequirementResponse.md) |  | 
**ComponentBindings** | [**[]PlatformComponentOutputBindingResponse**](PlatformComponentOutputBindingResponse.md) |  | 
**Violations** | [**[]PlatformComponentConfigurationViolationResponse**](PlatformComponentConfigurationViolationResponse.md) |  | 

## Methods

### NewPlatformComponentConfigurationPreviewResponse

`func NewPlatformComponentConfigurationPreviewResponse(clusterId string, componentKey string, fields []FieldSchemaResponse, requirements []PlatformComponentInputRequirementResponse, componentBindings []PlatformComponentOutputBindingResponse, violations []PlatformComponentConfigurationViolationResponse, ) *PlatformComponentConfigurationPreviewResponse`

NewPlatformComponentConfigurationPreviewResponse instantiates a new PlatformComponentConfigurationPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentConfigurationPreviewResponseWithDefaults

`func NewPlatformComponentConfigurationPreviewResponseWithDefaults() *PlatformComponentConfigurationPreviewResponse`

NewPlatformComponentConfigurationPreviewResponseWithDefaults instantiates a new PlatformComponentConfigurationPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *PlatformComponentConfigurationPreviewResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PlatformComponentConfigurationPreviewResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PlatformComponentConfigurationPreviewResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetComponentKey

`func (o *PlatformComponentConfigurationPreviewResponse) GetComponentKey() string`

GetComponentKey returns the ComponentKey field if non-nil, zero value otherwise.

### GetComponentKeyOk

`func (o *PlatformComponentConfigurationPreviewResponse) GetComponentKeyOk() (*string, bool)`

GetComponentKeyOk returns a tuple with the ComponentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentKey

`func (o *PlatformComponentConfigurationPreviewResponse) SetComponentKey(v string)`

SetComponentKey sets ComponentKey field to given value.


### GetFields

`func (o *PlatformComponentConfigurationPreviewResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PlatformComponentConfigurationPreviewResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PlatformComponentConfigurationPreviewResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.


### GetRequirements

`func (o *PlatformComponentConfigurationPreviewResponse) GetRequirements() []PlatformComponentInputRequirementResponse`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *PlatformComponentConfigurationPreviewResponse) GetRequirementsOk() (*[]PlatformComponentInputRequirementResponse, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *PlatformComponentConfigurationPreviewResponse) SetRequirements(v []PlatformComponentInputRequirementResponse)`

SetRequirements sets Requirements field to given value.


### GetComponentBindings

`func (o *PlatformComponentConfigurationPreviewResponse) GetComponentBindings() []PlatformComponentOutputBindingResponse`

GetComponentBindings returns the ComponentBindings field if non-nil, zero value otherwise.

### GetComponentBindingsOk

`func (o *PlatformComponentConfigurationPreviewResponse) GetComponentBindingsOk() (*[]PlatformComponentOutputBindingResponse, bool)`

GetComponentBindingsOk returns a tuple with the ComponentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentBindings

`func (o *PlatformComponentConfigurationPreviewResponse) SetComponentBindings(v []PlatformComponentOutputBindingResponse)`

SetComponentBindings sets ComponentBindings field to given value.


### GetViolations

`func (o *PlatformComponentConfigurationPreviewResponse) GetViolations() []PlatformComponentConfigurationViolationResponse`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *PlatformComponentConfigurationPreviewResponse) GetViolationsOk() (*[]PlatformComponentConfigurationViolationResponse, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *PlatformComponentConfigurationPreviewResponse) SetViolations(v []PlatformComponentConfigurationViolationResponse)`

SetViolations sets Violations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


