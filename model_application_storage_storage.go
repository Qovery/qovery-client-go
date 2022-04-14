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

// ApplicationStorageStorage struct for ApplicationStorageStorage
type ApplicationStorageStorage struct {
	Id   *string         `json:"id,omitempty"`
	Type StorageTypeEnum `json:"type"`
	// unit is GB
	Size       int32  `json:"size"`
	MountPoint string `json:"mount_point"`
}

// NewApplicationStorageStorage instantiates a new ApplicationStorageStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationStorageStorage(type_ StorageTypeEnum, size int32, mountPoint string) *ApplicationStorageStorage {
	this := ApplicationStorageStorage{}
	this.Type = type_
	this.Size = size
	this.MountPoint = mountPoint
	return &this
}

// NewApplicationStorageStorageWithDefaults instantiates a new ApplicationStorageStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationStorageStorageWithDefaults() *ApplicationStorageStorage {
	this := ApplicationStorageStorage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationStorageStorage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationStorageStorage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationStorageStorage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationStorageStorage) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *ApplicationStorageStorage) GetType() StorageTypeEnum {
	if o == nil {
		var ret StorageTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationStorageStorage) GetTypeOk() (*StorageTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationStorageStorage) SetType(v StorageTypeEnum) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *ApplicationStorageStorage) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ApplicationStorageStorage) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ApplicationStorageStorage) SetSize(v int32) {
	o.Size = v
}

// GetMountPoint returns the MountPoint field value
func (o *ApplicationStorageStorage) GetMountPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value
// and a boolean to check if the value has been set.
func (o *ApplicationStorageStorage) GetMountPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPoint, true
}

// SetMountPoint sets field value
func (o *ApplicationStorageStorage) SetMountPoint(v string) {
	o.MountPoint = v
}

func (o ApplicationStorageStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["mount_point"] = o.MountPoint
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationStorageStorage struct {
	value *ApplicationStorageStorage
	isSet bool
}

func (v NullableApplicationStorageStorage) Get() *ApplicationStorageStorage {
	return v.value
}

func (v *NullableApplicationStorageStorage) Set(val *ApplicationStorageStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationStorageStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationStorageStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationStorageStorage(val *ApplicationStorageStorage) *NullableApplicationStorageStorage {
	return &NullableApplicationStorageStorage{value: val, isSet: true}
}

func (v NullableApplicationStorageStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationStorageStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}