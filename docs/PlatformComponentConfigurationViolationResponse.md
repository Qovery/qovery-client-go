# PlatformComponentConfigurationViolationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Stable machine-readable violation code. | 
**FieldPath** | **string** | Path of the offending field. Convention: profile configuration fields use the bare field key (e.g. &#x60;storage&#x60;); cluster-provided inputs are prefixed with &#x60;clusterInputs.&#x60; (e.g. &#x60;clusterInputs.s3BucketName&#x60;). Consumers match violations to form fields using these exact paths. | 
**Message** | **string** | Customer-facing message describing how to fix the violation. | 

## Methods

### NewPlatformComponentConfigurationViolationResponse

`func NewPlatformComponentConfigurationViolationResponse(code string, fieldPath string, message string, ) *PlatformComponentConfigurationViolationResponse`

NewPlatformComponentConfigurationViolationResponse instantiates a new PlatformComponentConfigurationViolationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentConfigurationViolationResponseWithDefaults

`func NewPlatformComponentConfigurationViolationResponseWithDefaults() *PlatformComponentConfigurationViolationResponse`

NewPlatformComponentConfigurationViolationResponseWithDefaults instantiates a new PlatformComponentConfigurationViolationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PlatformComponentConfigurationViolationResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlatformComponentConfigurationViolationResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlatformComponentConfigurationViolationResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetFieldPath

`func (o *PlatformComponentConfigurationViolationResponse) GetFieldPath() string`

GetFieldPath returns the FieldPath field if non-nil, zero value otherwise.

### GetFieldPathOk

`func (o *PlatformComponentConfigurationViolationResponse) GetFieldPathOk() (*string, bool)`

GetFieldPathOk returns a tuple with the FieldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPath

`func (o *PlatformComponentConfigurationViolationResponse) SetFieldPath(v string)`

SetFieldPath sets FieldPath field to given value.


### GetMessage

`func (o *PlatformComponentConfigurationViolationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlatformComponentConfigurationViolationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlatformComponentConfigurationViolationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


