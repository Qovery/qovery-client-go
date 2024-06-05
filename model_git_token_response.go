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

// checks if the GitTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitTokenResponse{}

// GitTokenResponse struct for GitTokenResponse
type GitTokenResponse struct {
	Id          string          `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Type        GitProviderEnum `json:"type"`
	ExpiredAt   *string         `json:"expired_at,omitempty"`
	// Mandatory only for BITBUCKET git provider
	Workspace *string `json:"workspace,omitempty"`
	// The number of services using this git token
	AssociatedServicesCount float32 `json:"associated_services_count"`
}

type _GitTokenResponse GitTokenResponse

// NewGitTokenResponse instantiates a new GitTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTokenResponse(id string, createdAt time.Time, name string, type_ GitProviderEnum, associatedServicesCount float32) *GitTokenResponse {
	this := GitTokenResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Type = type_
	this.AssociatedServicesCount = associatedServicesCount
	return &this
}

// NewGitTokenResponseWithDefaults instantiates a new GitTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTokenResponseWithDefaults() *GitTokenResponse {
	this := GitTokenResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GitTokenResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GitTokenResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GitTokenResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GitTokenResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GitTokenResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GitTokenResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GitTokenResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *GitTokenResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitTokenResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GitTokenResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GitTokenResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GitTokenResponse) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *GitTokenResponse) GetType() GitProviderEnum {
	if o == nil {
		var ret GitProviderEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetTypeOk() (*GitProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GitTokenResponse) SetType(v GitProviderEnum) {
	o.Type = v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *GitTokenResponse) GetExpiredAt() string {
	if o == nil || IsNil(o.ExpiredAt) {
		var ret string
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetExpiredAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiredAt) {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *GitTokenResponse) HasExpiredAt() bool {
	if o != nil && !IsNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given string and assigns it to the ExpiredAt field.
func (o *GitTokenResponse) SetExpiredAt(v string) {
	o.ExpiredAt = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *GitTokenResponse) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *GitTokenResponse) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *GitTokenResponse) SetWorkspace(v string) {
	o.Workspace = &v
}

// GetAssociatedServicesCount returns the AssociatedServicesCount field value
func (o *GitTokenResponse) GetAssociatedServicesCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AssociatedServicesCount
}

// GetAssociatedServicesCountOk returns a tuple with the AssociatedServicesCount field value
// and a boolean to check if the value has been set.
func (o *GitTokenResponse) GetAssociatedServicesCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociatedServicesCount, true
}

// SetAssociatedServicesCount sets field value
func (o *GitTokenResponse) SetAssociatedServicesCount(v float32) {
	o.AssociatedServicesCount = v
}

func (o GitTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.ExpiredAt) {
		toSerialize["expired_at"] = o.ExpiredAt
	}
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	toSerialize["associated_services_count"] = o.AssociatedServicesCount
	return toSerialize, nil
}

func (o *GitTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"name",
		"type",
		"associated_services_count",
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

	varGitTokenResponse := _GitTokenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGitTokenResponse)

	if err != nil {
		return err
	}

	*o = GitTokenResponse(varGitTokenResponse)

	return err
}

type NullableGitTokenResponse struct {
	value *GitTokenResponse
	isSet bool
}

func (v NullableGitTokenResponse) Get() *GitTokenResponse {
	return v.value
}

func (v *NullableGitTokenResponse) Set(val *GitTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTokenResponse(val *GitTokenResponse) *NullableGitTokenResponse {
	return &NullableGitTokenResponse{value: val, isSet: true}
}

func (v NullableGitTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
