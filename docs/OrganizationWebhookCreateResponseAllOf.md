# OrganizationWebhookCreateResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**Kind**](Kind.md) |  | [optional] 
**TargetUrl** | Pointer to **string** | Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with &#x60;http://&#x60; or &#x60;https://&#x60;  | [optional] 
**TargetSecretSet** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Turn on or off your endpoint. | [optional] 
**Events** | Pointer to [**[]Items**](Items.md) |  | [optional] 
**ProjectIdFilter** | Pointer to **[]string** |  | [optional] 
**EnvironmentTypesFilter** | Pointer to [**[]EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 

## Methods

### NewOrganizationWebhookCreateResponseAllOf

`func NewOrganizationWebhookCreateResponseAllOf() *OrganizationWebhookCreateResponseAllOf`

NewOrganizationWebhookCreateResponseAllOf instantiates a new OrganizationWebhookCreateResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWebhookCreateResponseAllOfWithDefaults

`func NewOrganizationWebhookCreateResponseAllOfWithDefaults() *OrganizationWebhookCreateResponseAllOf`

NewOrganizationWebhookCreateResponseAllOfWithDefaults instantiates a new OrganizationWebhookCreateResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *OrganizationWebhookCreateResponseAllOf) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrganizationWebhookCreateResponseAllOf) SetKind(v Kind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrganizationWebhookCreateResponseAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTargetUrl

`func (o *OrganizationWebhookCreateResponseAllOf) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *OrganizationWebhookCreateResponseAllOf) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *OrganizationWebhookCreateResponseAllOf) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTargetSecretSet

`func (o *OrganizationWebhookCreateResponseAllOf) GetTargetSecretSet() bool`

GetTargetSecretSet returns the TargetSecretSet field if non-nil, zero value otherwise.

### GetTargetSecretSetOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetTargetSecretSetOk() (*bool, bool)`

GetTargetSecretSetOk returns a tuple with the TargetSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecretSet

`func (o *OrganizationWebhookCreateResponseAllOf) SetTargetSecretSet(v bool)`

SetTargetSecretSet sets TargetSecretSet field to given value.

### HasTargetSecretSet

`func (o *OrganizationWebhookCreateResponseAllOf) HasTargetSecretSet() bool`

HasTargetSecretSet returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationWebhookCreateResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationWebhookCreateResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationWebhookCreateResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationWebhookCreateResponseAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationWebhookCreateResponseAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationWebhookCreateResponseAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvents

`func (o *OrganizationWebhookCreateResponseAllOf) GetEvents() []Items`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetEventsOk() (*[]Items, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrganizationWebhookCreateResponseAllOf) SetEvents(v []Items)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrganizationWebhookCreateResponseAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetProjectIdFilter

`func (o *OrganizationWebhookCreateResponseAllOf) GetProjectIdFilter() []string`

GetProjectIdFilter returns the ProjectIdFilter field if non-nil, zero value otherwise.

### GetProjectIdFilterOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetProjectIdFilterOk() (*[]string, bool)`

GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdFilter

`func (o *OrganizationWebhookCreateResponseAllOf) SetProjectIdFilter(v []string)`

SetProjectIdFilter sets ProjectIdFilter field to given value.

### HasProjectIdFilter

`func (o *OrganizationWebhookCreateResponseAllOf) HasProjectIdFilter() bool`

HasProjectIdFilter returns a boolean if a field has been set.

### GetEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateResponseAllOf) GetEnvironmentTypesFilter() []EnvironmentModeEnum`

GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field if non-nil, zero value otherwise.

### GetEnvironmentTypesFilterOk

`func (o *OrganizationWebhookCreateResponseAllOf) GetEnvironmentTypesFilterOk() (*[]EnvironmentModeEnum, bool)`

GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateResponseAllOf) SetEnvironmentTypesFilter(v []EnvironmentModeEnum)`

SetEnvironmentTypesFilter sets EnvironmentTypesFilter field to given value.

### HasEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateResponseAllOf) HasEnvironmentTypesFilter() bool`

HasEnvironmentTypesFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


