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
	"fmt"
)

// checks if the GitRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitRepository{}

// GitRepository struct for GitRepository
type GitRepository struct {
	Id                   string  `json:"id"`
	Name                 string  `json:"name"`
	Url                  string  `json:"url"`
	DefaultBranch        *string `json:"default_branch,omitempty"`
	IsPrivate            *bool   `json:"is_private,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GitRepository GitRepository

// NewGitRepository instantiates a new GitRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepository(id string, name string, url string) *GitRepository {
	this := GitRepository{}
	this.Id = id
	this.Name = name
	this.Url = url
	return &this
}

// NewGitRepositoryWithDefaults instantiates a new GitRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepositoryWithDefaults() *GitRepository {
	this := GitRepository{}
	return &this
}

// GetId returns the Id field value
func (o *GitRepository) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GitRepository) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GitRepository) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GitRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitRepository) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *GitRepository) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GitRepository) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GitRepository) SetUrl(v string) {
	o.Url = v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *GitRepository) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetDefaultBranchOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *GitRepository) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *GitRepository) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *GitRepository) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *GitRepository) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *GitRepository) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

func (o GitRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.DefaultBranch) {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	if !IsNil(o.IsPrivate) {
		toSerialize["is_private"] = o.IsPrivate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GitRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"url",
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

	varGitRepository := _GitRepository{}

	err = json.Unmarshal(data, &varGitRepository)

	if err != nil {
		return err
	}

	*o = GitRepository(varGitRepository)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "default_branch")
		delete(additionalProperties, "is_private")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGitRepository struct {
	value *GitRepository
	isSet bool
}

func (v NullableGitRepository) Get() *GitRepository {
	return v.value
}

func (v *NullableGitRepository) Set(val *GitRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepository(val *GitRepository) *NullableGitRepository {
	return &NullableGitRepository{value: val, isSet: true}
}

func (v NullableGitRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
