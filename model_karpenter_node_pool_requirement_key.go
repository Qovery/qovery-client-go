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

// KarpenterNodePoolRequirementKey the model 'KarpenterNodePoolRequirementKey'
type KarpenterNodePoolRequirementKey string

// List of KarpenterNodePoolRequirementKey
const (
	KARPENTERNODEPOOLREQUIREMENTKEY_INSTANCE_FAMILY KarpenterNodePoolRequirementKey = "InstanceFamily"
	KARPENTERNODEPOOLREQUIREMENTKEY_INSTANCE_SIZE   KarpenterNodePoolRequirementKey = "InstanceSize"
	KARPENTERNODEPOOLREQUIREMENTKEY_ARCH            KarpenterNodePoolRequirementKey = "Arch"
)

// All allowed values of KarpenterNodePoolRequirementKey enum
var AllowedKarpenterNodePoolRequirementKeyEnumValues = []KarpenterNodePoolRequirementKey{
	"InstanceFamily",
	"InstanceSize",
	"Arch",
}

func (v *KarpenterNodePoolRequirementKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KarpenterNodePoolRequirementKey(value)
	for _, existing := range AllowedKarpenterNodePoolRequirementKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KarpenterNodePoolRequirementKey", value)
}

// NewKarpenterNodePoolRequirementKeyFromValue returns a pointer to a valid KarpenterNodePoolRequirementKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKarpenterNodePoolRequirementKeyFromValue(v string) (*KarpenterNodePoolRequirementKey, error) {
	ev := KarpenterNodePoolRequirementKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KarpenterNodePoolRequirementKey: valid values are %v", v, AllowedKarpenterNodePoolRequirementKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KarpenterNodePoolRequirementKey) IsValid() bool {
	for _, existing := range AllowedKarpenterNodePoolRequirementKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KarpenterNodePoolRequirementKey value
func (v KarpenterNodePoolRequirementKey) Ptr() *KarpenterNodePoolRequirementKey {
	return &v
}

type NullableKarpenterNodePoolRequirementKey struct {
	value *KarpenterNodePoolRequirementKey
	isSet bool
}

func (v NullableKarpenterNodePoolRequirementKey) Get() *KarpenterNodePoolRequirementKey {
	return v.value
}

func (v *NullableKarpenterNodePoolRequirementKey) Set(val *KarpenterNodePoolRequirementKey) {
	v.value = val
	v.isSet = true
}

func (v NullableKarpenterNodePoolRequirementKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKarpenterNodePoolRequirementKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKarpenterNodePoolRequirementKey(val *KarpenterNodePoolRequirementKey) *NullableKarpenterNodePoolRequirementKey {
	return &NullableKarpenterNodePoolRequirementKey{value: val, isSet: true}
}

func (v NullableKarpenterNodePoolRequirementKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKarpenterNodePoolRequirementKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}