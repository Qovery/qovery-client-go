# OrganizationPolicyApiTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**OpaPolicy** | **string** | Open Policy Agent (rego) rule definitions, without a &#x60;package&#x60; declaration: Qovery prepends a per-token package so that one token&#39;s rules cannot authorize another&#39;s. The policy must define an &#x60;allow&#x60; rule, and the request is denied unless it evaluates to true. | 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationPolicyApiTokenCreateRequest

`func NewOrganizationPolicyApiTokenCreateRequest(name string, opaPolicy string, ) *OrganizationPolicyApiTokenCreateRequest`

NewOrganizationPolicyApiTokenCreateRequest instantiates a new OrganizationPolicyApiTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPolicyApiTokenCreateRequestWithDefaults

`func NewOrganizationPolicyApiTokenCreateRequestWithDefaults() *OrganizationPolicyApiTokenCreateRequest`

NewOrganizationPolicyApiTokenCreateRequestWithDefaults instantiates a new OrganizationPolicyApiTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationPolicyApiTokenCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationPolicyApiTokenCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationPolicyApiTokenCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationPolicyApiTokenCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationPolicyApiTokenCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationPolicyApiTokenCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationPolicyApiTokenCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOpaPolicy

`func (o *OrganizationPolicyApiTokenCreateRequest) GetOpaPolicy() string`

GetOpaPolicy returns the OpaPolicy field if non-nil, zero value otherwise.

### GetOpaPolicyOk

`func (o *OrganizationPolicyApiTokenCreateRequest) GetOpaPolicyOk() (*string, bool)`

GetOpaPolicyOk returns a tuple with the OpaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaPolicy

`func (o *OrganizationPolicyApiTokenCreateRequest) SetOpaPolicy(v string)`

SetOpaPolicy sets OpaPolicy field to given value.


### GetExpiresAt

`func (o *OrganizationPolicyApiTokenCreateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationPolicyApiTokenCreateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationPolicyApiTokenCreateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationPolicyApiTokenCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OrganizationPolicyApiTokenCreateRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OrganizationPolicyApiTokenCreateRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


