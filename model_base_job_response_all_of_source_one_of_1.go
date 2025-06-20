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

// checks if the BaseJobResponseAllOfSourceOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseJobResponseAllOfSourceOneOf1{}

// BaseJobResponseAllOfSourceOneOf1 struct for BaseJobResponseAllOfSourceOneOf1
type BaseJobResponseAllOfSourceOneOf1 struct {
	Docker               JobSourceDockerResponse `json:"docker"`
	AdditionalProperties map[string]interface{}
}

type _BaseJobResponseAllOfSourceOneOf1 BaseJobResponseAllOfSourceOneOf1

// NewBaseJobResponseAllOfSourceOneOf1 instantiates a new BaseJobResponseAllOfSourceOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseJobResponseAllOfSourceOneOf1(docker JobSourceDockerResponse) *BaseJobResponseAllOfSourceOneOf1 {
	this := BaseJobResponseAllOfSourceOneOf1{}
	this.Docker = docker
	return &this
}

// NewBaseJobResponseAllOfSourceOneOf1WithDefaults instantiates a new BaseJobResponseAllOfSourceOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseJobResponseAllOfSourceOneOf1WithDefaults() *BaseJobResponseAllOfSourceOneOf1 {
	this := BaseJobResponseAllOfSourceOneOf1{}
	return &this
}

// GetDocker returns the Docker field value
func (o *BaseJobResponseAllOfSourceOneOf1) GetDocker() JobSourceDockerResponse {
	if o == nil {
		var ret JobSourceDockerResponse
		return ret
	}

	return o.Docker
}

// GetDockerOk returns a tuple with the Docker field value
// and a boolean to check if the value has been set.
func (o *BaseJobResponseAllOfSourceOneOf1) GetDockerOk() (*JobSourceDockerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Docker, true
}

// SetDocker sets field value
func (o *BaseJobResponseAllOfSourceOneOf1) SetDocker(v JobSourceDockerResponse) {
	o.Docker = v
}

func (o BaseJobResponseAllOfSourceOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseJobResponseAllOfSourceOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["docker"] = o.Docker

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseJobResponseAllOfSourceOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"docker",
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

	varBaseJobResponseAllOfSourceOneOf1 := _BaseJobResponseAllOfSourceOneOf1{}

	err = json.Unmarshal(data, &varBaseJobResponseAllOfSourceOneOf1)

	if err != nil {
		return err
	}

	*o = BaseJobResponseAllOfSourceOneOf1(varBaseJobResponseAllOfSourceOneOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "docker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseJobResponseAllOfSourceOneOf1 struct {
	value *BaseJobResponseAllOfSourceOneOf1
	isSet bool
}

func (v NullableBaseJobResponseAllOfSourceOneOf1) Get() *BaseJobResponseAllOfSourceOneOf1 {
	return v.value
}

func (v *NullableBaseJobResponseAllOfSourceOneOf1) Set(val *BaseJobResponseAllOfSourceOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseJobResponseAllOfSourceOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseJobResponseAllOfSourceOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseJobResponseAllOfSourceOneOf1(val *BaseJobResponseAllOfSourceOneOf1) *NullableBaseJobResponseAllOfSourceOneOf1 {
	return &NullableBaseJobResponseAllOfSourceOneOf1{value: val, isSet: true}
}

func (v NullableBaseJobResponseAllOfSourceOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseJobResponseAllOfSourceOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
