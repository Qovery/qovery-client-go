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

// StorageTypeEnum the model 'StorageTypeEnum'
type StorageTypeEnum string

// List of StorageTypeEnum
const (
	AWS           StorageTypeEnum = "AWS"
	DIGITAL_OCEAN StorageTypeEnum = "DIGITAL_OCEAN"
	SCALEWAY      StorageTypeEnum = "SCALEWAY"
)

// All allowed values of StorageTypeEnum enum
var AllowedStorageTypeEnumEnumValues = []StorageTypeEnum{
	"AWS",
	"DIGITAL_OCEAN",
	"SCALEWAY",
}

func (v *StorageTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageTypeEnum(value)
	for _, existing := range AllowedStorageTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageTypeEnum", value)
}

// NewStorageTypeEnumFromValue returns a pointer to a valid StorageTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageTypeEnumFromValue(v string) (*StorageTypeEnum, error) {
	ev := StorageTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageTypeEnum: valid values are %v", v, AllowedStorageTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageTypeEnum) IsValid() bool {
	for _, existing := range AllowedStorageTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageTypeEnum value
func (v StorageTypeEnum) Ptr() *StorageTypeEnum {
	return &v
}

type NullableStorageTypeEnum struct {
	value *StorageTypeEnum
	isSet bool
}

func (v NullableStorageTypeEnum) Get() *StorageTypeEnum {
	return v.value
}

func (v *NullableStorageTypeEnum) Set(val *StorageTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageTypeEnum(val *StorageTypeEnum) *NullableStorageTypeEnum {
	return &NullableStorageTypeEnum{value: val, isSet: true}
}

func (v NullableStorageTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}