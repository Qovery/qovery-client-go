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
)

// checks if the DeploymentHistoryServiceDetailsOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryServiceDetailsOneOf3{}

// DeploymentHistoryServiceDetailsOneOf3 HelmDeploymentHistoryDetails
type DeploymentHistoryServiceDetailsOneOf3 struct {
	Commit               NullableCommit                                   `json:"commit,omitempty"`
	Repository           *DeploymentHistoryServiceDetailsOneOf3Repository `json:"repository,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentHistoryServiceDetailsOneOf3 DeploymentHistoryServiceDetailsOneOf3

// NewDeploymentHistoryServiceDetailsOneOf3 instantiates a new DeploymentHistoryServiceDetailsOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryServiceDetailsOneOf3() *DeploymentHistoryServiceDetailsOneOf3 {
	this := DeploymentHistoryServiceDetailsOneOf3{}
	return &this
}

// NewDeploymentHistoryServiceDetailsOneOf3WithDefaults instantiates a new DeploymentHistoryServiceDetailsOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryServiceDetailsOneOf3WithDefaults() *DeploymentHistoryServiceDetailsOneOf3 {
	this := DeploymentHistoryServiceDetailsOneOf3{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentHistoryServiceDetailsOneOf3) GetCommit() Commit {
	if o == nil || IsNil(o.Commit.Get()) {
		var ret Commit
		return ret
	}
	return *o.Commit.Get()
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentHistoryServiceDetailsOneOf3) GetCommitOk() (*Commit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commit.Get(), o.Commit.IsSet()
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3) HasCommit() bool {
	if o != nil && o.Commit.IsSet() {
		return true
	}

	return false
}

// SetCommit gets a reference to the given NullableCommit and assigns it to the Commit field.
func (o *DeploymentHistoryServiceDetailsOneOf3) SetCommit(v Commit) {
	o.Commit.Set(&v)
}

// SetCommitNil sets the value for Commit to be an explicit nil
func (o *DeploymentHistoryServiceDetailsOneOf3) SetCommitNil() {
	o.Commit.Set(nil)
}

// UnsetCommit ensures that no value is present for Commit, not even an explicit nil
func (o *DeploymentHistoryServiceDetailsOneOf3) UnsetCommit() {
	o.Commit.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DeploymentHistoryServiceDetailsOneOf3) GetRepository() DeploymentHistoryServiceDetailsOneOf3Repository {
	if o == nil || IsNil(o.Repository) {
		var ret DeploymentHistoryServiceDetailsOneOf3Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3) GetRepositoryOk() (*DeploymentHistoryServiceDetailsOneOf3Repository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given DeploymentHistoryServiceDetailsOneOf3Repository and assigns it to the Repository field.
func (o *DeploymentHistoryServiceDetailsOneOf3) SetRepository(v DeploymentHistoryServiceDetailsOneOf3Repository) {
	o.Repository = &v
}

func (o DeploymentHistoryServiceDetailsOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryServiceDetailsOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Commit.IsSet() {
		toSerialize["commit"] = o.Commit.Get()
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentHistoryServiceDetailsOneOf3) UnmarshalJSON(data []byte) (err error) {
	varDeploymentHistoryServiceDetailsOneOf3 := _DeploymentHistoryServiceDetailsOneOf3{}

	err = json.Unmarshal(data, &varDeploymentHistoryServiceDetailsOneOf3)

	if err != nil {
		return err
	}

	*o = DeploymentHistoryServiceDetailsOneOf3(varDeploymentHistoryServiceDetailsOneOf3)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commit")
		delete(additionalProperties, "repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentHistoryServiceDetailsOneOf3 struct {
	value *DeploymentHistoryServiceDetailsOneOf3
	isSet bool
}

func (v NullableDeploymentHistoryServiceDetailsOneOf3) Get() *DeploymentHistoryServiceDetailsOneOf3 {
	return v.value
}

func (v *NullableDeploymentHistoryServiceDetailsOneOf3) Set(val *DeploymentHistoryServiceDetailsOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryServiceDetailsOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryServiceDetailsOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryServiceDetailsOneOf3(val *DeploymentHistoryServiceDetailsOneOf3) *NullableDeploymentHistoryServiceDetailsOneOf3 {
	return &NullableDeploymentHistoryServiceDetailsOneOf3{value: val, isSet: true}
}

func (v NullableDeploymentHistoryServiceDetailsOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryServiceDetailsOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
