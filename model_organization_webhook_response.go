/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the OrganizationWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationWebhookResponse{}

// OrganizationWebhookResponse struct for OrganizationWebhookResponse
type OrganizationWebhookResponse struct {
	Id        string                       `json:"id"`
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt *time.Time                   `json:"updated_at,omitempty"`
	Kind      *OrganizationWebhookKindEnum `json:"kind,omitempty"`
	// Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with `http://` or `https://`
	TargetUrl       *string `json:"target_url,omitempty"`
	TargetSecretSet *bool   `json:"target_secret_set,omitempty"`
	Description     *string `json:"description,omitempty"`
	// Turn on or off your endpoint.
	Enabled *bool                          `json:"enabled,omitempty"`
	Events  []OrganizationWebhookEventEnum `json:"events,omitempty"`
	// Specify the project names you want to filter to.  This webhook will be triggered only if the event is coming from the specified Project IDs. Notes: 1. Wildcard is accepted E.g. `product*`. 2. Name is case insensitive.
	ProjectNamesFilter []string `json:"project_names_filter,omitempty"`
	// Specify the environment modes you want to filter to. This webhook will be triggered only if the event is coming from an environment with the specified mode.
	EnvironmentTypesFilter []EnvironmentModeEnum `json:"environment_types_filter,omitempty"`
}

type _OrganizationWebhookResponse OrganizationWebhookResponse

// NewOrganizationWebhookResponse instantiates a new OrganizationWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationWebhookResponse(id string, createdAt time.Time) *OrganizationWebhookResponse {
	this := OrganizationWebhookResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationWebhookResponseWithDefaults instantiates a new OrganizationWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWebhookResponseWithDefaults() *OrganizationWebhookResponse {
	this := OrganizationWebhookResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationWebhookResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationWebhookResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationWebhookResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationWebhookResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationWebhookResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetKind() OrganizationWebhookKindEnum {
	if o == nil || IsNil(o.Kind) {
		var ret OrganizationWebhookKindEnum
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetKindOk() (*OrganizationWebhookKindEnum, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given OrganizationWebhookKindEnum and assigns it to the Kind field.
func (o *OrganizationWebhookResponse) SetKind(v OrganizationWebhookKindEnum) {
	o.Kind = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *OrganizationWebhookResponse) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTargetSecretSet returns the TargetSecretSet field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetTargetSecretSet() bool {
	if o == nil || IsNil(o.TargetSecretSet) {
		var ret bool
		return ret
	}
	return *o.TargetSecretSet
}

// GetTargetSecretSetOk returns a tuple with the TargetSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetTargetSecretSetOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetSecretSet) {
		return nil, false
	}
	return o.TargetSecretSet, true
}

// HasTargetSecretSet returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasTargetSecretSet() bool {
	if o != nil && !IsNil(o.TargetSecretSet) {
		return true
	}

	return false
}

// SetTargetSecretSet gets a reference to the given bool and assigns it to the TargetSecretSet field.
func (o *OrganizationWebhookResponse) SetTargetSecretSet(v bool) {
	o.TargetSecretSet = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationWebhookResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationWebhookResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetEvents() []OrganizationWebhookEventEnum {
	if o == nil || IsNil(o.Events) {
		var ret []OrganizationWebhookEventEnum
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetEventsOk() ([]OrganizationWebhookEventEnum, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []OrganizationWebhookEventEnum and assigns it to the Events field.
func (o *OrganizationWebhookResponse) SetEvents(v []OrganizationWebhookEventEnum) {
	o.Events = v
}

// GetProjectNamesFilter returns the ProjectNamesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetProjectNamesFilter() []string {
	if o == nil || IsNil(o.ProjectNamesFilter) {
		var ret []string
		return ret
	}
	return o.ProjectNamesFilter
}

// GetProjectNamesFilterOk returns a tuple with the ProjectNamesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetProjectNamesFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectNamesFilter) {
		return nil, false
	}
	return o.ProjectNamesFilter, true
}

// HasProjectNamesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasProjectNamesFilter() bool {
	if o != nil && !IsNil(o.ProjectNamesFilter) {
		return true
	}

	return false
}

// SetProjectNamesFilter gets a reference to the given []string and assigns it to the ProjectNamesFilter field.
func (o *OrganizationWebhookResponse) SetProjectNamesFilter(v []string) {
	o.ProjectNamesFilter = v
}

// GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookResponse) GetEnvironmentTypesFilter() []EnvironmentModeEnum {
	if o == nil || IsNil(o.EnvironmentTypesFilter) {
		var ret []EnvironmentModeEnum
		return ret
	}
	return o.EnvironmentTypesFilter
}

// GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookResponse) GetEnvironmentTypesFilterOk() ([]EnvironmentModeEnum, bool) {
	if o == nil || IsNil(o.EnvironmentTypesFilter) {
		return nil, false
	}
	return o.EnvironmentTypesFilter, true
}

// HasEnvironmentTypesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookResponse) HasEnvironmentTypesFilter() bool {
	if o != nil && !IsNil(o.EnvironmentTypesFilter) {
		return true
	}

	return false
}

// SetEnvironmentTypesFilter gets a reference to the given []EnvironmentModeEnum and assigns it to the EnvironmentTypesFilter field.
func (o *OrganizationWebhookResponse) SetEnvironmentTypesFilter(v []EnvironmentModeEnum) {
	o.EnvironmentTypesFilter = v
}

func (o OrganizationWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["target_url"] = o.TargetUrl
	}
	if !IsNil(o.TargetSecretSet) {
		toSerialize["target_secret_set"] = o.TargetSecretSet
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.ProjectNamesFilter) {
		toSerialize["project_names_filter"] = o.ProjectNamesFilter
	}
	if !IsNil(o.EnvironmentTypesFilter) {
		toSerialize["environment_types_filter"] = o.EnvironmentTypesFilter
	}
	return toSerialize, nil
}

func (o *OrganizationWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrganizationWebhookResponse := _OrganizationWebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationWebhookResponse)

	if err != nil {
		return err
	}

	*o = OrganizationWebhookResponse(varOrganizationWebhookResponse)

	return err
}

type NullableOrganizationWebhookResponse struct {
	value *OrganizationWebhookResponse
	isSet bool
}

func (v NullableOrganizationWebhookResponse) Get() *OrganizationWebhookResponse {
	return v.value
}

func (v *NullableOrganizationWebhookResponse) Set(val *OrganizationWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWebhookResponse(val *OrganizationWebhookResponse) *NullableOrganizationWebhookResponse {
	return &NullableOrganizationWebhookResponse{value: val, isSet: true}
}

func (v NullableOrganizationWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
