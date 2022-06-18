# CreateOrganizationWebhook201ResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**Object**](Object.md) |  | [optional] 
**TargetUrl** | Pointer to **string** | Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with &#x60;http://&#x60; or &#x60;https://&#x60;  | [optional] 
**TargetSecretSet** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Turn on or off your endpoint. | [optional] 
**Events** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**ProjectIdFilter** | Pointer to **[]string** |  | [optional] 
**EnvironmentTypesFilter** | Pointer to [**[]EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 

## Methods

### NewCreateOrganizationWebhook201ResponseAllOf

`func NewCreateOrganizationWebhook201ResponseAllOf() *CreateOrganizationWebhook201ResponseAllOf`

NewCreateOrganizationWebhook201ResponseAllOf instantiates a new CreateOrganizationWebhook201ResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationWebhook201ResponseAllOfWithDefaults

`func NewCreateOrganizationWebhook201ResponseAllOfWithDefaults() *CreateOrganizationWebhook201ResponseAllOf`

NewCreateOrganizationWebhook201ResponseAllOfWithDefaults instantiates a new CreateOrganizationWebhook201ResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetKind() Object`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetKindOk() (*Object, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetKind(v Object)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTargetUrl

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTargetSecretSet

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetSecretSet() bool`

GetTargetSecretSet returns the TargetSecretSet field if non-nil, zero value otherwise.

### GetTargetSecretSetOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetSecretSetOk() (*bool, bool)`

GetTargetSecretSetOk returns a tuple with the TargetSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecretSet

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetTargetSecretSet(v bool)`

SetTargetSecretSet sets TargetSecretSet field to given value.

### HasTargetSecretSet

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasTargetSecretSet() bool`

HasTargetSecretSet returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvents

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetEvents() []Object`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetEventsOk() (*[]Object, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetEvents(v []Object)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetProjectIdFilter

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetProjectIdFilter() []string`

GetProjectIdFilter returns the ProjectIdFilter field if non-nil, zero value otherwise.

### GetProjectIdFilterOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetProjectIdFilterOk() (*[]string, bool)`

GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdFilter

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetProjectIdFilter(v []string)`

SetProjectIdFilter sets ProjectIdFilter field to given value.

### HasProjectIdFilter

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasProjectIdFilter() bool`

HasProjectIdFilter returns a boolean if a field has been set.

### GetEnvironmentTypesFilter

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnvironmentTypesFilter() []EnvironmentModeEnum`

GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field if non-nil, zero value otherwise.

### GetEnvironmentTypesFilterOk

`func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnvironmentTypesFilterOk() (*[]EnvironmentModeEnum, bool)`

GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTypesFilter

`func (o *CreateOrganizationWebhook201ResponseAllOf) SetEnvironmentTypesFilter(v []EnvironmentModeEnum)`

SetEnvironmentTypesFilter sets EnvironmentTypesFilter field to given value.

### HasEnvironmentTypesFilter

`func (o *CreateOrganizationWebhook201ResponseAllOf) HasEnvironmentTypesFilter() bool`

HasEnvironmentTypesFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


