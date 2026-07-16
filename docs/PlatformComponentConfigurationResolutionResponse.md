# PlatformComponentConfigurationResolutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentKey** | **string** |  | 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) |  | 
**Requirements** | [**[]PlatformComponentInputRequirementResponse**](PlatformComponentInputRequirementResponse.md) |  | 
**ComponentBindings** | [**[]PlatformComponentOutputBindingResponse**](PlatformComponentOutputBindingResponse.md) |  | 
**Violations** | [**[]PlatformComponentConfigurationViolationResponse**](PlatformComponentConfigurationViolationResponse.md) |  | 

## Methods

### NewPlatformComponentConfigurationResolutionResponse

`func NewPlatformComponentConfigurationResolutionResponse(componentKey string, fields []FieldSchemaResponse, requirements []PlatformComponentInputRequirementResponse, componentBindings []PlatformComponentOutputBindingResponse, violations []PlatformComponentConfigurationViolationResponse, ) *PlatformComponentConfigurationResolutionResponse`

NewPlatformComponentConfigurationResolutionResponse instantiates a new PlatformComponentConfigurationResolutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentConfigurationResolutionResponseWithDefaults

`func NewPlatformComponentConfigurationResolutionResponseWithDefaults() *PlatformComponentConfigurationResolutionResponse`

NewPlatformComponentConfigurationResolutionResponseWithDefaults instantiates a new PlatformComponentConfigurationResolutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentKey

`func (o *PlatformComponentConfigurationResolutionResponse) GetComponentKey() string`

GetComponentKey returns the ComponentKey field if non-nil, zero value otherwise.

### GetComponentKeyOk

`func (o *PlatformComponentConfigurationResolutionResponse) GetComponentKeyOk() (*string, bool)`

GetComponentKeyOk returns a tuple with the ComponentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentKey

`func (o *PlatformComponentConfigurationResolutionResponse) SetComponentKey(v string)`

SetComponentKey sets ComponentKey field to given value.


### GetFields

`func (o *PlatformComponentConfigurationResolutionResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PlatformComponentConfigurationResolutionResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PlatformComponentConfigurationResolutionResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.


### GetRequirements

`func (o *PlatformComponentConfigurationResolutionResponse) GetRequirements() []PlatformComponentInputRequirementResponse`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *PlatformComponentConfigurationResolutionResponse) GetRequirementsOk() (*[]PlatformComponentInputRequirementResponse, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *PlatformComponentConfigurationResolutionResponse) SetRequirements(v []PlatformComponentInputRequirementResponse)`

SetRequirements sets Requirements field to given value.


### GetComponentBindings

`func (o *PlatformComponentConfigurationResolutionResponse) GetComponentBindings() []PlatformComponentOutputBindingResponse`

GetComponentBindings returns the ComponentBindings field if non-nil, zero value otherwise.

### GetComponentBindingsOk

`func (o *PlatformComponentConfigurationResolutionResponse) GetComponentBindingsOk() (*[]PlatformComponentOutputBindingResponse, bool)`

GetComponentBindingsOk returns a tuple with the ComponentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentBindings

`func (o *PlatformComponentConfigurationResolutionResponse) SetComponentBindings(v []PlatformComponentOutputBindingResponse)`

SetComponentBindings sets ComponentBindings field to given value.


### GetViolations

`func (o *PlatformComponentConfigurationResolutionResponse) GetViolations() []PlatformComponentConfigurationViolationResponse`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *PlatformComponentConfigurationResolutionResponse) GetViolationsOk() (*[]PlatformComponentConfigurationViolationResponse, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *PlatformComponentConfigurationResolutionResponse) SetViolations(v []PlatformComponentConfigurationViolationResponse)`

SetViolations sets Violations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


