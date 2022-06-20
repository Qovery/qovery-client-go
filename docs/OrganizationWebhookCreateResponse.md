# OrganizationWebhookCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Kind** | Pointer to [**Kind**](Kind.md) |  | [optional] 
**TargetUrl** | Pointer to **string** | Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with &#x60;http://&#x60; or &#x60;https://&#x60;  | [optional] 
**TargetSecretSet** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Turn on or off your endpoint. | [optional] 
**Events** | Pointer to [**[]Items**](Items.md) |  | [optional] 
**ProjectIdFilter** | Pointer to **[]string** |  | [optional] 
**EnvironmentTypesFilter** | Pointer to [**[]EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 

## Methods

### NewOrganizationWebhookCreateResponse

`func NewOrganizationWebhookCreateResponse(id string, createdAt time.Time, ) *OrganizationWebhookCreateResponse`

NewOrganizationWebhookCreateResponse instantiates a new OrganizationWebhookCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWebhookCreateResponseWithDefaults

`func NewOrganizationWebhookCreateResponseWithDefaults() *OrganizationWebhookCreateResponse`

NewOrganizationWebhookCreateResponseWithDefaults instantiates a new OrganizationWebhookCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationWebhookCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationWebhookCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationWebhookCreateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationWebhookCreateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationWebhookCreateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationWebhookCreateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationWebhookCreateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationWebhookCreateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationWebhookCreateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationWebhookCreateResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetKind

`func (o *OrganizationWebhookCreateResponse) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrganizationWebhookCreateResponse) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrganizationWebhookCreateResponse) SetKind(v Kind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrganizationWebhookCreateResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTargetUrl

`func (o *OrganizationWebhookCreateResponse) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *OrganizationWebhookCreateResponse) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *OrganizationWebhookCreateResponse) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *OrganizationWebhookCreateResponse) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTargetSecretSet

`func (o *OrganizationWebhookCreateResponse) GetTargetSecretSet() bool`

GetTargetSecretSet returns the TargetSecretSet field if non-nil, zero value otherwise.

### GetTargetSecretSetOk

`func (o *OrganizationWebhookCreateResponse) GetTargetSecretSetOk() (*bool, bool)`

GetTargetSecretSetOk returns a tuple with the TargetSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecretSet

`func (o *OrganizationWebhookCreateResponse) SetTargetSecretSet(v bool)`

SetTargetSecretSet sets TargetSecretSet field to given value.

### HasTargetSecretSet

`func (o *OrganizationWebhookCreateResponse) HasTargetSecretSet() bool`

HasTargetSecretSet returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationWebhookCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationWebhookCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationWebhookCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationWebhookCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationWebhookCreateResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationWebhookCreateResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationWebhookCreateResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationWebhookCreateResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvents

`func (o *OrganizationWebhookCreateResponse) GetEvents() []Items`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrganizationWebhookCreateResponse) GetEventsOk() (*[]Items, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrganizationWebhookCreateResponse) SetEvents(v []Items)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrganizationWebhookCreateResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetProjectIdFilter

`func (o *OrganizationWebhookCreateResponse) GetProjectIdFilter() []string`

GetProjectIdFilter returns the ProjectIdFilter field if non-nil, zero value otherwise.

### GetProjectIdFilterOk

`func (o *OrganizationWebhookCreateResponse) GetProjectIdFilterOk() (*[]string, bool)`

GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdFilter

`func (o *OrganizationWebhookCreateResponse) SetProjectIdFilter(v []string)`

SetProjectIdFilter sets ProjectIdFilter field to given value.

### HasProjectIdFilter

`func (o *OrganizationWebhookCreateResponse) HasProjectIdFilter() bool`

HasProjectIdFilter returns a boolean if a field has been set.

### GetEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateResponse) GetEnvironmentTypesFilter() []EnvironmentModeEnum`

GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field if non-nil, zero value otherwise.

### GetEnvironmentTypesFilterOk

`func (o *OrganizationWebhookCreateResponse) GetEnvironmentTypesFilterOk() (*[]EnvironmentModeEnum, bool)`

GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateResponse) SetEnvironmentTypesFilter(v []EnvironmentModeEnum)`

SetEnvironmentTypesFilter sets EnvironmentTypesFilter field to given value.

### HasEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateResponse) HasEnvironmentTypesFilter() bool`

HasEnvironmentTypesFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


