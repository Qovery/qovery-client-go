# CustomDomainErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**CustomDomainErrorCodeEnum**](CustomDomainErrorCodeEnum.md) |  | 
**Field** | **string** | The request field to blame, named as on the wire, such as &#x60;domain&#x60;, &#x60;path&#x60;, &#x60;port_id&#x60;, &#x60;generate_certificate&#x60; or &#x60;use_cdn&#x60;. | 
**Path** | **NullableString** | The path in conflict. Null when the error is not about a path. | 
**ServiceId** | **string** | Id of the service of the request that cannot hold the domain. | 
**ConflictingService** | [**NullableCustomDomainErrorResponseConflictingService**](CustomDomainErrorResponseConflictingService.md) |  | 
**Message** | **string** | A message to show to the user. | 

## Methods

### NewCustomDomainErrorResponse

`func NewCustomDomainErrorResponse(code CustomDomainErrorCodeEnum, field string, path NullableString, serviceId string, conflictingService NullableCustomDomainErrorResponseConflictingService, message string, ) *CustomDomainErrorResponse`

NewCustomDomainErrorResponse instantiates a new CustomDomainErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainErrorResponseWithDefaults

`func NewCustomDomainErrorResponseWithDefaults() *CustomDomainErrorResponse`

NewCustomDomainErrorResponseWithDefaults instantiates a new CustomDomainErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CustomDomainErrorResponse) GetCode() CustomDomainErrorCodeEnum`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomDomainErrorResponse) GetCodeOk() (*CustomDomainErrorCodeEnum, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomDomainErrorResponse) SetCode(v CustomDomainErrorCodeEnum)`

SetCode sets Code field to given value.


### GetField

`func (o *CustomDomainErrorResponse) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CustomDomainErrorResponse) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CustomDomainErrorResponse) SetField(v string)`

SetField sets Field field to given value.


### GetPath

`func (o *CustomDomainErrorResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CustomDomainErrorResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CustomDomainErrorResponse) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *CustomDomainErrorResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CustomDomainErrorResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetServiceId

`func (o *CustomDomainErrorResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CustomDomainErrorResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CustomDomainErrorResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetConflictingService

`func (o *CustomDomainErrorResponse) GetConflictingService() CustomDomainErrorResponseConflictingService`

GetConflictingService returns the ConflictingService field if non-nil, zero value otherwise.

### GetConflictingServiceOk

`func (o *CustomDomainErrorResponse) GetConflictingServiceOk() (*CustomDomainErrorResponseConflictingService, bool)`

GetConflictingServiceOk returns a tuple with the ConflictingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingService

`func (o *CustomDomainErrorResponse) SetConflictingService(v CustomDomainErrorResponseConflictingService)`

SetConflictingService sets ConflictingService field to given value.


### SetConflictingServiceNil

`func (o *CustomDomainErrorResponse) SetConflictingServiceNil(b bool)`

 SetConflictingServiceNil sets the value for ConflictingService to be an explicit nil

### UnsetConflictingService
`func (o *CustomDomainErrorResponse) UnsetConflictingService()`

UnsetConflictingService ensures that no value is present for ConflictingService, not even an explicit nil
### GetMessage

`func (o *CustomDomainErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CustomDomainErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CustomDomainErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


