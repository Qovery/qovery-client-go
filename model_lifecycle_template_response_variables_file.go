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
	"fmt"
)

// checks if the LifecycleTemplateResponseVariablesFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleTemplateResponseVariablesFile{}

// LifecycleTemplateResponseVariablesFile If present, the variable should be a file instead of a raw value
type LifecycleTemplateResponseVariablesFile struct {
	Path                 string `json:"path"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleTemplateResponseVariablesFile LifecycleTemplateResponseVariablesFile

// NewLifecycleTemplateResponseVariablesFile instantiates a new LifecycleTemplateResponseVariablesFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleTemplateResponseVariablesFile(path string) *LifecycleTemplateResponseVariablesFile {
	this := LifecycleTemplateResponseVariablesFile{}
	this.Path = path
	return &this
}

// NewLifecycleTemplateResponseVariablesFileWithDefaults instantiates a new LifecycleTemplateResponseVariablesFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleTemplateResponseVariablesFileWithDefaults() *LifecycleTemplateResponseVariablesFile {
	this := LifecycleTemplateResponseVariablesFile{}
	return &this
}

// GetPath returns the Path field value
func (o *LifecycleTemplateResponseVariablesFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponseVariablesFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *LifecycleTemplateResponseVariablesFile) SetPath(v string) {
	o.Path = v
}

func (o LifecycleTemplateResponseVariablesFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleTemplateResponseVariablesFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleTemplateResponseVariablesFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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

	varLifecycleTemplateResponseVariablesFile := _LifecycleTemplateResponseVariablesFile{}

	err = json.Unmarshal(data, &varLifecycleTemplateResponseVariablesFile)

	if err != nil {
		return err
	}

	*o = LifecycleTemplateResponseVariablesFile(varLifecycleTemplateResponseVariablesFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleTemplateResponseVariablesFile struct {
	value *LifecycleTemplateResponseVariablesFile
	isSet bool
}

func (v NullableLifecycleTemplateResponseVariablesFile) Get() *LifecycleTemplateResponseVariablesFile {
	return v.value
}

func (v *NullableLifecycleTemplateResponseVariablesFile) Set(val *LifecycleTemplateResponseVariablesFile) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleTemplateResponseVariablesFile) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleTemplateResponseVariablesFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleTemplateResponseVariablesFile(val *LifecycleTemplateResponseVariablesFile) *NullableLifecycleTemplateResponseVariablesFile {
	return &NullableLifecycleTemplateResponseVariablesFile{value: val, isSet: true}
}

func (v NullableLifecycleTemplateResponseVariablesFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleTemplateResponseVariablesFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
