# BlueprintDatabaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**DatabaseTypeEnum**](DatabaseTypeEnum.md) |  | 
**ServiceId** | **NullableString** | Terraform service the blueprint created; null before its first dispatch | 
**Endpoint** | [**NullableBlueprintDatabaseResponseEndpoint**](BlueprintDatabaseResponseEndpoint.md) |  | 

## Methods

### NewBlueprintDatabaseResponse

`func NewBlueprintDatabaseResponse(kind DatabaseTypeEnum, serviceId NullableString, endpoint NullableBlueprintDatabaseResponseEndpoint, ) *BlueprintDatabaseResponse`

NewBlueprintDatabaseResponse instantiates a new BlueprintDatabaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDatabaseResponseWithDefaults

`func NewBlueprintDatabaseResponseWithDefaults() *BlueprintDatabaseResponse`

NewBlueprintDatabaseResponseWithDefaults instantiates a new BlueprintDatabaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BlueprintDatabaseResponse) GetKind() DatabaseTypeEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlueprintDatabaseResponse) GetKindOk() (*DatabaseTypeEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlueprintDatabaseResponse) SetKind(v DatabaseTypeEnum)`

SetKind sets Kind field to given value.


### GetServiceId

`func (o *BlueprintDatabaseResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *BlueprintDatabaseResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *BlueprintDatabaseResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### SetServiceIdNil

`func (o *BlueprintDatabaseResponse) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *BlueprintDatabaseResponse) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetEndpoint

`func (o *BlueprintDatabaseResponse) GetEndpoint() BlueprintDatabaseResponseEndpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BlueprintDatabaseResponse) GetEndpointOk() (*BlueprintDatabaseResponseEndpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BlueprintDatabaseResponse) SetEndpoint(v BlueprintDatabaseResponseEndpoint)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *BlueprintDatabaseResponse) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *BlueprintDatabaseResponse) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


