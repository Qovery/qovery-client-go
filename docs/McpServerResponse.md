# McpServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Url** | **string** | HTTPS URL of the remote MCP server | 
**HeaderNames** | **[]string** | Names of the configured HTTP headers. Header values are never returned. | 
**Scope** | [**McpServerScope**](McpServerScope.md) |  | 
**OwnerUserSub** | Pointer to **NullableString** | Identity of the owning member. Null for an ORGANIZATION connector. | [optional] 
**OwnerName** | Pointer to **NullableString** | Display name of the owning member. Null for an ORGANIZATION connector. | [optional] 
**Attachable** | **bool** | Whether the member making this request may attach the connector to an agentic workflow. Computed per caller: an organization admin sees every USER connector but can attach none of them, so a picker must use this rather than scope alone. | 

## Methods

### NewMcpServerResponse

`func NewMcpServerResponse(id string, createdAt time.Time, updatedAt time.Time, name string, description string, url string, headerNames []string, scope McpServerScope, attachable bool, ) *McpServerResponse`

NewMcpServerResponse instantiates a new McpServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerResponseWithDefaults

`func NewMcpServerResponseWithDefaults() *McpServerResponse`

NewMcpServerResponseWithDefaults instantiates a new McpServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McpServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *McpServerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *McpServerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *McpServerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *McpServerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *McpServerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *McpServerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *McpServerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *McpServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *McpServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *McpServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *McpServerResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaderNames

`func (o *McpServerResponse) GetHeaderNames() []string`

GetHeaderNames returns the HeaderNames field if non-nil, zero value otherwise.

### GetHeaderNamesOk

`func (o *McpServerResponse) GetHeaderNamesOk() (*[]string, bool)`

GetHeaderNamesOk returns a tuple with the HeaderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderNames

`func (o *McpServerResponse) SetHeaderNames(v []string)`

SetHeaderNames sets HeaderNames field to given value.


### GetScope

`func (o *McpServerResponse) GetScope() McpServerScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *McpServerResponse) GetScopeOk() (*McpServerScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *McpServerResponse) SetScope(v McpServerScope)`

SetScope sets Scope field to given value.


### GetOwnerUserSub

`func (o *McpServerResponse) GetOwnerUserSub() string`

GetOwnerUserSub returns the OwnerUserSub field if non-nil, zero value otherwise.

### GetOwnerUserSubOk

`func (o *McpServerResponse) GetOwnerUserSubOk() (*string, bool)`

GetOwnerUserSubOk returns a tuple with the OwnerUserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserSub

`func (o *McpServerResponse) SetOwnerUserSub(v string)`

SetOwnerUserSub sets OwnerUserSub field to given value.

### HasOwnerUserSub

`func (o *McpServerResponse) HasOwnerUserSub() bool`

HasOwnerUserSub returns a boolean if a field has been set.

### SetOwnerUserSubNil

`func (o *McpServerResponse) SetOwnerUserSubNil(b bool)`

 SetOwnerUserSubNil sets the value for OwnerUserSub to be an explicit nil

### UnsetOwnerUserSub
`func (o *McpServerResponse) UnsetOwnerUserSub()`

UnsetOwnerUserSub ensures that no value is present for OwnerUserSub, not even an explicit nil
### GetOwnerName

`func (o *McpServerResponse) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *McpServerResponse) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *McpServerResponse) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *McpServerResponse) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *McpServerResponse) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *McpServerResponse) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetAttachable

`func (o *McpServerResponse) GetAttachable() bool`

GetAttachable returns the Attachable field if non-nil, zero value otherwise.

### GetAttachableOk

`func (o *McpServerResponse) GetAttachableOk() (*bool, bool)`

GetAttachableOk returns a tuple with the Attachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachable

`func (o *McpServerResponse) SetAttachable(v bool)`

SetAttachable sets Attachable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


