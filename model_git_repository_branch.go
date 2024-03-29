/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// checks if the GitRepositoryBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitRepositoryBranch{}

// GitRepositoryBranch struct for GitRepositoryBranch
type GitRepositoryBranch struct {
	Name string `json:"name"`
}

// NewGitRepositoryBranch instantiates a new GitRepositoryBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepositoryBranch(name string) *GitRepositoryBranch {
	this := GitRepositoryBranch{}
	this.Name = name
	return &this
}

// NewGitRepositoryBranchWithDefaults instantiates a new GitRepositoryBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepositoryBranchWithDefaults() *GitRepositoryBranch {
	this := GitRepositoryBranch{}
	return &this
}

// GetName returns the Name field value
func (o *GitRepositoryBranch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitRepositoryBranch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitRepositoryBranch) SetName(v string) {
	o.Name = v
}

func (o GitRepositoryBranch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitRepositoryBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGitRepositoryBranch struct {
	value *GitRepositoryBranch
	isSet bool
}

func (v NullableGitRepositoryBranch) Get() *GitRepositoryBranch {
	return v.value
}

func (v *NullableGitRepositoryBranch) Set(val *GitRepositoryBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepositoryBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepositoryBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepositoryBranch(val *GitRepositoryBranch) *NullableGitRepositoryBranch {
	return &NullableGitRepositoryBranch{value: val, isSet: true}
}

func (v NullableGitRepositoryBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepositoryBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
