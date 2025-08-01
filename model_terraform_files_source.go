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

// checks if the TerraformFilesSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerraformFilesSource{}

// TerraformFilesSource struct for TerraformFilesSource
type TerraformFilesSource struct {
	Git                  *TerraformFilesSourceGit `json:"git,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TerraformFilesSource TerraformFilesSource

// NewTerraformFilesSource instantiates a new TerraformFilesSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformFilesSource() *TerraformFilesSource {
	this := TerraformFilesSource{}
	return &this
}

// NewTerraformFilesSourceWithDefaults instantiates a new TerraformFilesSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformFilesSourceWithDefaults() *TerraformFilesSource {
	this := TerraformFilesSource{}
	return &this
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *TerraformFilesSource) GetGit() TerraformFilesSourceGit {
	if o == nil || IsNil(o.Git) {
		var ret TerraformFilesSourceGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformFilesSource) GetGitOk() (*TerraformFilesSourceGit, bool) {
	if o == nil || IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *TerraformFilesSource) HasGit() bool {
	if o != nil && !IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given TerraformFilesSourceGit and assigns it to the Git field.
func (o *TerraformFilesSource) SetGit(v TerraformFilesSourceGit) {
	o.Git = &v
}

func (o TerraformFilesSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerraformFilesSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TerraformFilesSource) UnmarshalJSON(data []byte) (err error) {
	varTerraformFilesSource := _TerraformFilesSource{}

	err = json.Unmarshal(data, &varTerraformFilesSource)

	if err != nil {
		return err
	}

	*o = TerraformFilesSource(varTerraformFilesSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "git")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerraformFilesSource struct {
	value *TerraformFilesSource
	isSet bool
}

func (v NullableTerraformFilesSource) Get() *TerraformFilesSource {
	return v.value
}

func (v *NullableTerraformFilesSource) Set(val *TerraformFilesSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformFilesSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformFilesSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformFilesSource(val *TerraformFilesSource) *NullableTerraformFilesSource {
	return &NullableTerraformFilesSource{value: val, isSet: true}
}

func (v NullableTerraformFilesSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformFilesSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
