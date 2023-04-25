# OrganizationEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**EventType** | Pointer to [**OrganizationEventType**](OrganizationEventType.md) |  | [optional] 
**TargetId** | Pointer to **NullableString** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to [**OrganizationEventTargetType**](OrganizationEventTargetType.md) |  | [optional] 
**SubTargetType** | Pointer to [**OrganizationEventSubTargetType**](OrganizationEventSubTargetType.md) |  | [optional] 
**Change** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | [optional] 
**TriggeredBy** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrganizationEventResponse

`func NewOrganizationEventResponse() *OrganizationEventResponse`

NewOrganizationEventResponse instantiates a new OrganizationEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationEventResponseWithDefaults

`func NewOrganizationEventResponseWithDefaults() *OrganizationEventResponse`

NewOrganizationEventResponseWithDefaults instantiates a new OrganizationEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationEventResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationEventResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *OrganizationEventResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrganizationEventResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrganizationEventResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OrganizationEventResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEventType

`func (o *OrganizationEventResponse) GetEventType() OrganizationEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OrganizationEventResponse) GetEventTypeOk() (*OrganizationEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OrganizationEventResponse) SetEventType(v OrganizationEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OrganizationEventResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTargetId

`func (o *OrganizationEventResponse) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *OrganizationEventResponse) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *OrganizationEventResponse) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *OrganizationEventResponse) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *OrganizationEventResponse) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *OrganizationEventResponse) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *OrganizationEventResponse) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *OrganizationEventResponse) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *OrganizationEventResponse) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *OrganizationEventResponse) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *OrganizationEventResponse) GetTargetType() OrganizationEventTargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *OrganizationEventResponse) GetTargetTypeOk() (*OrganizationEventTargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *OrganizationEventResponse) SetTargetType(v OrganizationEventTargetType)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *OrganizationEventResponse) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetSubTargetType

`func (o *OrganizationEventResponse) GetSubTargetType() OrganizationEventSubTargetType`

GetSubTargetType returns the SubTargetType field if non-nil, zero value otherwise.

### GetSubTargetTypeOk

`func (o *OrganizationEventResponse) GetSubTargetTypeOk() (*OrganizationEventSubTargetType, bool)`

GetSubTargetTypeOk returns a tuple with the SubTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTargetType

`func (o *OrganizationEventResponse) SetSubTargetType(v OrganizationEventSubTargetType)`

SetSubTargetType sets SubTargetType field to given value.

### HasSubTargetType

`func (o *OrganizationEventResponse) HasSubTargetType() bool`

HasSubTargetType returns a boolean if a field has been set.

### GetChange

`func (o *OrganizationEventResponse) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *OrganizationEventResponse) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *OrganizationEventResponse) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *OrganizationEventResponse) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetOrigin

`func (o *OrganizationEventResponse) GetOrigin() OrganizationEventOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *OrganizationEventResponse) GetOriginOk() (*OrganizationEventOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *OrganizationEventResponse) SetOrigin(v OrganizationEventOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *OrganizationEventResponse) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *OrganizationEventResponse) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *OrganizationEventResponse) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *OrganizationEventResponse) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *OrganizationEventResponse) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetProjectId

`func (o *OrganizationEventResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OrganizationEventResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OrganizationEventResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OrganizationEventResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *OrganizationEventResponse) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *OrganizationEventResponse) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetEnvironmentId

`func (o *OrganizationEventResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *OrganizationEventResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *OrganizationEventResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *OrganizationEventResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *OrganizationEventResponse) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *OrganizationEventResponse) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


