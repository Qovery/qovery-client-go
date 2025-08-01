/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.4
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// checks if the OrganizationEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationEventResponse{}

// OrganizationEventResponse
type OrganizationEventResponse struct {
	Id                   *string                                `json:"id,omitempty"`
	Timestamp            *time.Time                             `json:"timestamp,omitempty"`
	EventType            *OrganizationEventType                 `json:"event_type,omitempty"`
	TargetId             NullableString                         `json:"target_id,omitempty"`
	TargetName           *string                                `json:"target_name,omitempty"`
	TargetType           *OrganizationEventTargetType           `json:"target_type,omitempty"`
	SubTargetType        NullableOrganizationEventSubTargetType `json:"sub_target_type,omitempty"`
	Change               *string                                `json:"change,omitempty"`
	Origin               *OrganizationEventOrigin               `json:"origin,omitempty"`
	TriggeredBy          *string                                `json:"triggered_by,omitempty"`
	ProjectId            NullableString                         `json:"project_id,omitempty"`
	ProjectName          *string                                `json:"project_name,omitempty"`
	EnvironmentId        NullableString                         `json:"environment_id,omitempty"`
	EnvironmentName      *string                                `json:"environment_name,omitempty"`
	UserAgent            NullableString                         `json:"user_agent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationEventResponse OrganizationEventResponse

// NewOrganizationEventResponse instantiates a new OrganizationEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEventResponse() *OrganizationEventResponse {
	this := OrganizationEventResponse{}
	return &this
}

// NewOrganizationEventResponseWithDefaults instantiates a new OrganizationEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEventResponseWithDefaults() *OrganizationEventResponse {
	this := OrganizationEventResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationEventResponse) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *OrganizationEventResponse) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetEventType() OrganizationEventType {
	if o == nil || IsNil(o.EventType) {
		var ret OrganizationEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetEventTypeOk() (*OrganizationEventType, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given OrganizationEventType and assigns it to the EventType field.
func (o *OrganizationEventResponse) SetEventType(v OrganizationEventType) {
	o.EventType = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEventResponse) GetTargetId() string {
	if o == nil || IsNil(o.TargetId.Get()) {
		var ret string
		return ret
	}
	return *o.TargetId.Get()
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEventResponse) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetId.Get(), o.TargetId.IsSet()
}

// HasTargetId returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasTargetId() bool {
	if o != nil && o.TargetId.IsSet() {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given NullableString and assigns it to the TargetId field.
func (o *OrganizationEventResponse) SetTargetId(v string) {
	o.TargetId.Set(&v)
}

// SetTargetIdNil sets the value for TargetId to be an explicit nil
func (o *OrganizationEventResponse) SetTargetIdNil() {
	o.TargetId.Set(nil)
}

// UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
func (o *OrganizationEventResponse) UnsetTargetId() {
	o.TargetId.Unset()
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *OrganizationEventResponse) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetTargetType() OrganizationEventTargetType {
	if o == nil || IsNil(o.TargetType) {
		var ret OrganizationEventTargetType
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetTargetTypeOk() (*OrganizationEventTargetType, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given OrganizationEventTargetType and assigns it to the TargetType field.
func (o *OrganizationEventResponse) SetTargetType(v OrganizationEventTargetType) {
	o.TargetType = &v
}

// GetSubTargetType returns the SubTargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEventResponse) GetSubTargetType() OrganizationEventSubTargetType {
	if o == nil || IsNil(o.SubTargetType.Get()) {
		var ret OrganizationEventSubTargetType
		return ret
	}
	return *o.SubTargetType.Get()
}

// GetSubTargetTypeOk returns a tuple with the SubTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEventResponse) GetSubTargetTypeOk() (*OrganizationEventSubTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubTargetType.Get(), o.SubTargetType.IsSet()
}

// HasSubTargetType returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasSubTargetType() bool {
	if o != nil && o.SubTargetType.IsSet() {
		return true
	}

	return false
}

// SetSubTargetType gets a reference to the given NullableOrganizationEventSubTargetType and assigns it to the SubTargetType field.
func (o *OrganizationEventResponse) SetSubTargetType(v OrganizationEventSubTargetType) {
	o.SubTargetType.Set(&v)
}

// SetSubTargetTypeNil sets the value for SubTargetType to be an explicit nil
func (o *OrganizationEventResponse) SetSubTargetTypeNil() {
	o.SubTargetType.Set(nil)
}

// UnsetSubTargetType ensures that no value is present for SubTargetType, not even an explicit nil
func (o *OrganizationEventResponse) UnsetSubTargetType() {
	o.SubTargetType.Unset()
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetChange() string {
	if o == nil || IsNil(o.Change) {
		var ret string
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given string and assigns it to the Change field.
func (o *OrganizationEventResponse) SetChange(v string) {
	o.Change = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetOrigin() OrganizationEventOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret OrganizationEventOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetOriginOk() (*OrganizationEventOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given OrganizationEventOrigin and assigns it to the Origin field.
func (o *OrganizationEventResponse) SetOrigin(v OrganizationEventOrigin) {
	o.Origin = &v
}

// GetTriggeredBy returns the TriggeredBy field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetTriggeredBy() string {
	if o == nil || IsNil(o.TriggeredBy) {
		var ret string
		return ret
	}
	return *o.TriggeredBy
}

// GetTriggeredByOk returns a tuple with the TriggeredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetTriggeredByOk() (*string, bool) {
	if o == nil || IsNil(o.TriggeredBy) {
		return nil, false
	}
	return o.TriggeredBy, true
}

// HasTriggeredBy returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasTriggeredBy() bool {
	if o != nil && !IsNil(o.TriggeredBy) {
		return true
	}

	return false
}

// SetTriggeredBy gets a reference to the given string and assigns it to the TriggeredBy field.
func (o *OrganizationEventResponse) SetTriggeredBy(v string) {
	o.TriggeredBy = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEventResponse) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEventResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *OrganizationEventResponse) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}

// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *OrganizationEventResponse) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *OrganizationEventResponse) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *OrganizationEventResponse) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEventResponse) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEventResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *OrganizationEventResponse) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *OrganizationEventResponse) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *OrganizationEventResponse) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *OrganizationEventResponse) GetEnvironmentName() string {
	if o == nil || IsNil(o.EnvironmentName) {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponse) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentName) {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasEnvironmentName() bool {
	if o != nil && !IsNil(o.EnvironmentName) {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *OrganizationEventResponse) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEventResponse) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEventResponse) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *OrganizationEventResponse) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableString and assigns it to the UserAgent field.
func (o *OrganizationEventResponse) SetUserAgent(v string) {
	o.UserAgent.Set(&v)
}

// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *OrganizationEventResponse) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *OrganizationEventResponse) UnsetUserAgent() {
	o.UserAgent.Unset()
}

func (o OrganizationEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if o.TargetId.IsSet() {
		toSerialize["target_id"] = o.TargetId.Get()
	}
	if !IsNil(o.TargetName) {
		toSerialize["target_name"] = o.TargetName
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	if o.SubTargetType.IsSet() {
		toSerialize["sub_target_type"] = o.SubTargetType.Get()
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.TriggeredBy) {
		toSerialize["triggered_by"] = o.TriggeredBy
	}
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if !IsNil(o.ProjectName) {
		toSerialize["project_name"] = o.ProjectName
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if !IsNil(o.EnvironmentName) {
		toSerialize["environment_name"] = o.EnvironmentName
	}
	if o.UserAgent.IsSet() {
		toSerialize["user_agent"] = o.UserAgent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationEventResponse) UnmarshalJSON(data []byte) (err error) {
	varOrganizationEventResponse := _OrganizationEventResponse{}

	err = json.Unmarshal(data, &varOrganizationEventResponse)

	if err != nil {
		return err
	}

	*o = OrganizationEventResponse(varOrganizationEventResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "event_type")
		delete(additionalProperties, "target_id")
		delete(additionalProperties, "target_name")
		delete(additionalProperties, "target_type")
		delete(additionalProperties, "sub_target_type")
		delete(additionalProperties, "change")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "triggered_by")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "project_name")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "environment_name")
		delete(additionalProperties, "user_agent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationEventResponse struct {
	value *OrganizationEventResponse
	isSet bool
}

func (v NullableOrganizationEventResponse) Get() *OrganizationEventResponse {
	return v.value
}

func (v *NullableOrganizationEventResponse) Set(val *OrganizationEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventResponse(val *OrganizationEventResponse) *NullableOrganizationEventResponse {
	return &NullableOrganizationEventResponse{value: val, isSet: true}
}

func (v NullableOrganizationEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
