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

// checks if the GitTokenAssociatedServicesResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitTokenAssociatedServicesResponseList{}

// GitTokenAssociatedServicesResponseList struct for GitTokenAssociatedServicesResponseList
type GitTokenAssociatedServicesResponseList struct {
	Results              []GitTokenAssociatedServiceResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GitTokenAssociatedServicesResponseList GitTokenAssociatedServicesResponseList

// NewGitTokenAssociatedServicesResponseList instantiates a new GitTokenAssociatedServicesResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTokenAssociatedServicesResponseList() *GitTokenAssociatedServicesResponseList {
	this := GitTokenAssociatedServicesResponseList{}
	return &this
}

// NewGitTokenAssociatedServicesResponseListWithDefaults instantiates a new GitTokenAssociatedServicesResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTokenAssociatedServicesResponseListWithDefaults() *GitTokenAssociatedServicesResponseList {
	this := GitTokenAssociatedServicesResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GitTokenAssociatedServicesResponseList) GetResults() []GitTokenAssociatedServiceResponse {
	if o == nil || IsNil(o.Results) {
		var ret []GitTokenAssociatedServiceResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenAssociatedServicesResponseList) GetResultsOk() ([]GitTokenAssociatedServiceResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GitTokenAssociatedServicesResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GitTokenAssociatedServiceResponse and assigns it to the Results field.
func (o *GitTokenAssociatedServicesResponseList) SetResults(v []GitTokenAssociatedServiceResponse) {
	o.Results = v
}

func (o GitTokenAssociatedServicesResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitTokenAssociatedServicesResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GitTokenAssociatedServicesResponseList) UnmarshalJSON(data []byte) (err error) {
	varGitTokenAssociatedServicesResponseList := _GitTokenAssociatedServicesResponseList{}

	err = json.Unmarshal(data, &varGitTokenAssociatedServicesResponseList)

	if err != nil {
		return err
	}

	*o = GitTokenAssociatedServicesResponseList(varGitTokenAssociatedServicesResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGitTokenAssociatedServicesResponseList struct {
	value *GitTokenAssociatedServicesResponseList
	isSet bool
}

func (v NullableGitTokenAssociatedServicesResponseList) Get() *GitTokenAssociatedServicesResponseList {
	return v.value
}

func (v *NullableGitTokenAssociatedServicesResponseList) Set(val *GitTokenAssociatedServicesResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTokenAssociatedServicesResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTokenAssociatedServicesResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTokenAssociatedServicesResponseList(val *GitTokenAssociatedServicesResponseList) *NullableGitTokenAssociatedServicesResponseList {
	return &NullableGitTokenAssociatedServicesResponseList{value: val, isSet: true}
}

func (v NullableGitTokenAssociatedServicesResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTokenAssociatedServicesResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
