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

// checks if the KarpenterDefaultNodePoolOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KarpenterDefaultNodePoolOverride{}

// KarpenterDefaultNodePoolOverride struct for KarpenterDefaultNodePoolOverride
type KarpenterDefaultNodePoolOverride struct {
	Limits               *KarpenterNodePoolLimits `json:"limits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KarpenterDefaultNodePoolOverride KarpenterDefaultNodePoolOverride

// NewKarpenterDefaultNodePoolOverride instantiates a new KarpenterDefaultNodePoolOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKarpenterDefaultNodePoolOverride() *KarpenterDefaultNodePoolOverride {
	this := KarpenterDefaultNodePoolOverride{}
	return &this
}

// NewKarpenterDefaultNodePoolOverrideWithDefaults instantiates a new KarpenterDefaultNodePoolOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKarpenterDefaultNodePoolOverrideWithDefaults() *KarpenterDefaultNodePoolOverride {
	this := KarpenterDefaultNodePoolOverride{}
	return &this
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *KarpenterDefaultNodePoolOverride) GetLimits() KarpenterNodePoolLimits {
	if o == nil || IsNil(o.Limits) {
		var ret KarpenterNodePoolLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KarpenterDefaultNodePoolOverride) GetLimitsOk() (*KarpenterNodePoolLimits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *KarpenterDefaultNodePoolOverride) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given KarpenterNodePoolLimits and assigns it to the Limits field.
func (o *KarpenterDefaultNodePoolOverride) SetLimits(v KarpenterNodePoolLimits) {
	o.Limits = &v
}

func (o KarpenterDefaultNodePoolOverride) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KarpenterDefaultNodePoolOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KarpenterDefaultNodePoolOverride) UnmarshalJSON(data []byte) (err error) {
	varKarpenterDefaultNodePoolOverride := _KarpenterDefaultNodePoolOverride{}

	err = json.Unmarshal(data, &varKarpenterDefaultNodePoolOverride)

	if err != nil {
		return err
	}

	*o = KarpenterDefaultNodePoolOverride(varKarpenterDefaultNodePoolOverride)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "limits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKarpenterDefaultNodePoolOverride struct {
	value *KarpenterDefaultNodePoolOverride
	isSet bool
}

func (v NullableKarpenterDefaultNodePoolOverride) Get() *KarpenterDefaultNodePoolOverride {
	return v.value
}

func (v *NullableKarpenterDefaultNodePoolOverride) Set(val *KarpenterDefaultNodePoolOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableKarpenterDefaultNodePoolOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableKarpenterDefaultNodePoolOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKarpenterDefaultNodePoolOverride(val *KarpenterDefaultNodePoolOverride) *NullableKarpenterDefaultNodePoolOverride {
	return &NullableKarpenterDefaultNodePoolOverride{value: val, isSet: true}
}

func (v NullableKarpenterDefaultNodePoolOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKarpenterDefaultNodePoolOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
