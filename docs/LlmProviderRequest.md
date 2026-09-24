# LlmProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | LLM provider name, unique per scope owner within the organization | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Type** | [**LlmProviderType**](LlmProviderType.md) |  | 
**Credential** | Pointer to **string** | The provider credential. Encrypted at rest and never returned by the API. Blank means unchanged: on create no credential is stored, on edit the stored one is kept. Sending a nonblank value rotates it. | [optional] [default to ""]
**Scope** | Pointer to [**LlmProviderScope**](LlmProviderScope.md) | Cannot be changed after creation. On create, omitting it means ORGANIZATION, which requires the MANAGE_INFRASTRUCTURE permission; creating a USER provider requires CREATE_PROJECT. On edit, omitting it leaves the provider&#39;s scope unchanged, and stating a scope that differs from the provider&#39;s is refused with 400. | [optional] 
**Region** | Pointer to **NullableString** | On edit, omitting it or sending null clears the stored region. This differs from credential, where a blank value keeps the stored one. AWS region the Bedrock client calls, for example us-east-1 or eu-west-1; model availability differs by region. Only allowed for a BEDROCK provider. Any sent string, blank included, must match the pattern. Null keeps the engine default region. | [optional] 

## Methods

### NewLlmProviderRequest

`func NewLlmProviderRequest(name string, type_ LlmProviderType, ) *LlmProviderRequest`

NewLlmProviderRequest instantiates a new LlmProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmProviderRequestWithDefaults

`func NewLlmProviderRequestWithDefaults() *LlmProviderRequest`

NewLlmProviderRequestWithDefaults instantiates a new LlmProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LlmProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LlmProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LlmProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LlmProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LlmProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LlmProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LlmProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *LlmProviderRequest) GetType() LlmProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LlmProviderRequest) GetTypeOk() (*LlmProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LlmProviderRequest) SetType(v LlmProviderType)`

SetType sets Type field to given value.


### GetCredential

`func (o *LlmProviderRequest) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *LlmProviderRequest) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *LlmProviderRequest) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *LlmProviderRequest) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetScope

`func (o *LlmProviderRequest) GetScope() LlmProviderScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LlmProviderRequest) GetScopeOk() (*LlmProviderScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LlmProviderRequest) SetScope(v LlmProviderScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *LlmProviderRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRegion

`func (o *LlmProviderRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LlmProviderRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LlmProviderRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LlmProviderRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *LlmProviderRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *LlmProviderRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


