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

// checks if the JobResponseAllOfSourceOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobResponseAllOfSourceOneOf1{}

// JobResponseAllOfSourceOneOf1 struct for JobResponseAllOfSourceOneOf1
type JobResponseAllOfSourceOneOf1 struct {
	Docker *JobResponseAllOfSourceOneOf1Docker `json:"docker,omitempty"`
}

// NewJobResponseAllOfSourceOneOf1 instantiates a new JobResponseAllOfSourceOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResponseAllOfSourceOneOf1() *JobResponseAllOfSourceOneOf1 {
	this := JobResponseAllOfSourceOneOf1{}
	return &this
}

// NewJobResponseAllOfSourceOneOf1WithDefaults instantiates a new JobResponseAllOfSourceOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResponseAllOfSourceOneOf1WithDefaults() *JobResponseAllOfSourceOneOf1 {
	this := JobResponseAllOfSourceOneOf1{}
	return &this
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *JobResponseAllOfSourceOneOf1) GetDocker() JobResponseAllOfSourceOneOf1Docker {
	if o == nil || IsNil(o.Docker) {
		var ret JobResponseAllOfSourceOneOf1Docker
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponseAllOfSourceOneOf1) GetDockerOk() (*JobResponseAllOfSourceOneOf1Docker, bool) {
	if o == nil || IsNil(o.Docker) {
		return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *JobResponseAllOfSourceOneOf1) HasDocker() bool {
	if o != nil && !IsNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given JobResponseAllOfSourceOneOf1Docker and assigns it to the Docker field.
func (o *JobResponseAllOfSourceOneOf1) SetDocker(v JobResponseAllOfSourceOneOf1Docker) {
	o.Docker = &v
}

func (o JobResponseAllOfSourceOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobResponseAllOfSourceOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Docker) {
		toSerialize["docker"] = o.Docker
	}
	return toSerialize, nil
}

type NullableJobResponseAllOfSourceOneOf1 struct {
	value *JobResponseAllOfSourceOneOf1
	isSet bool
}

func (v NullableJobResponseAllOfSourceOneOf1) Get() *JobResponseAllOfSourceOneOf1 {
	return v.value
}

func (v *NullableJobResponseAllOfSourceOneOf1) Set(val *JobResponseAllOfSourceOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResponseAllOfSourceOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResponseAllOfSourceOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResponseAllOfSourceOneOf1(val *JobResponseAllOfSourceOneOf1) *NullableJobResponseAllOfSourceOneOf1 {
	return &NullableJobResponseAllOfSourceOneOf1{value: val, isSet: true}
}

func (v NullableJobResponseAllOfSourceOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResponseAllOfSourceOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}