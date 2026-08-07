# OrganizationAgentApiTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**OpaPolicy** | **string** | Open Policy Agent (rego) rule definitions, without a &#x60;package&#x60; declaration: Qovery prepends a per-token package so that one token&#39;s rules cannot authorize another&#39;s. The policy must define an &#x60;allow&#x60; rule, and the request is denied unless it evaluates to true. | 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationAgentApiTokenCreateRequest

`func NewOrganizationAgentApiTokenCreateRequest(name string, opaPolicy string, ) *OrganizationAgentApiTokenCreateRequest`

NewOrganizationAgentApiTokenCreateRequest instantiates a new OrganizationAgentApiTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAgentApiTokenCreateRequestWithDefaults

`func NewOrganizationAgentApiTokenCreateRequestWithDefaults() *OrganizationAgentApiTokenCreateRequest`

NewOrganizationAgentApiTokenCreateRequestWithDefaults instantiates a new OrganizationAgentApiTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationAgentApiTokenCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAgentApiTokenCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAgentApiTokenCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationAgentApiTokenCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationAgentApiTokenCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationAgentApiTokenCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationAgentApiTokenCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOpaPolicy

`func (o *OrganizationAgentApiTokenCreateRequest) GetOpaPolicy() string`

GetOpaPolicy returns the OpaPolicy field if non-nil, zero value otherwise.

### GetOpaPolicyOk

`func (o *OrganizationAgentApiTokenCreateRequest) GetOpaPolicyOk() (*string, bool)`

GetOpaPolicyOk returns a tuple with the OpaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaPolicy

`func (o *OrganizationAgentApiTokenCreateRequest) SetOpaPolicy(v string)`

SetOpaPolicy sets OpaPolicy field to given value.


### GetExpiresAt

`func (o *OrganizationAgentApiTokenCreateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationAgentApiTokenCreateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationAgentApiTokenCreateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationAgentApiTokenCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OrganizationAgentApiTokenCreateRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OrganizationAgentApiTokenCreateRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


