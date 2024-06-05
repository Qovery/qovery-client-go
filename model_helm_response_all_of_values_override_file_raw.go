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
)

// checks if the HelmResponseAllOfValuesOverrideFileRaw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmResponseAllOfValuesOverrideFileRaw{}

// HelmResponseAllOfValuesOverrideFileRaw struct for HelmResponseAllOfValuesOverrideFileRaw
type HelmResponseAllOfValuesOverrideFileRaw struct {
	Values []HelmResponseAllOfValuesOverrideFileRawValues `json:"values"`
}

type _HelmResponseAllOfValuesOverrideFileRaw HelmResponseAllOfValuesOverrideFileRaw

// NewHelmResponseAllOfValuesOverrideFileRaw instantiates a new HelmResponseAllOfValuesOverrideFileRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmResponseAllOfValuesOverrideFileRaw(values []HelmResponseAllOfValuesOverrideFileRawValues) *HelmResponseAllOfValuesOverrideFileRaw {
	this := HelmResponseAllOfValuesOverrideFileRaw{}
	this.Values = values
	return &this
}

// NewHelmResponseAllOfValuesOverrideFileRawWithDefaults instantiates a new HelmResponseAllOfValuesOverrideFileRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmResponseAllOfValuesOverrideFileRawWithDefaults() *HelmResponseAllOfValuesOverrideFileRaw {
	this := HelmResponseAllOfValuesOverrideFileRaw{}
	return &this
}

// GetValues returns the Values field value
func (o *HelmResponseAllOfValuesOverrideFileRaw) GetValues() []HelmResponseAllOfValuesOverrideFileRawValues {
	if o == nil {
		var ret []HelmResponseAllOfValuesOverrideFileRawValues
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfValuesOverrideFileRaw) GetValuesOk() ([]HelmResponseAllOfValuesOverrideFileRawValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *HelmResponseAllOfValuesOverrideFileRaw) SetValues(v []HelmResponseAllOfValuesOverrideFileRawValues) {
	o.Values = v
}

func (o HelmResponseAllOfValuesOverrideFileRaw) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmResponseAllOfValuesOverrideFileRaw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *HelmResponseAllOfValuesOverrideFileRaw) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
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

	varHelmResponseAllOfValuesOverrideFileRaw := _HelmResponseAllOfValuesOverrideFileRaw{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHelmResponseAllOfValuesOverrideFileRaw)

	if err != nil {
		return err
	}

	*o = HelmResponseAllOfValuesOverrideFileRaw(varHelmResponseAllOfValuesOverrideFileRaw)

	return err
}

type NullableHelmResponseAllOfValuesOverrideFileRaw struct {
	value *HelmResponseAllOfValuesOverrideFileRaw
	isSet bool
}

func (v NullableHelmResponseAllOfValuesOverrideFileRaw) Get() *HelmResponseAllOfValuesOverrideFileRaw {
	return v.value
}

func (v *NullableHelmResponseAllOfValuesOverrideFileRaw) Set(val *HelmResponseAllOfValuesOverrideFileRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmResponseAllOfValuesOverrideFileRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmResponseAllOfValuesOverrideFileRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmResponseAllOfValuesOverrideFileRaw(val *HelmResponseAllOfValuesOverrideFileRaw) *NullableHelmResponseAllOfValuesOverrideFileRaw {
	return &NullableHelmResponseAllOfValuesOverrideFileRaw{value: val, isSet: true}
}

func (v NullableHelmResponseAllOfValuesOverrideFileRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmResponseAllOfValuesOverrideFileRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
