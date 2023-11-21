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

// checks if the BaseJobResponseAllOfSourceOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseJobResponseAllOfSourceOneOf{}

// BaseJobResponseAllOfSourceOneOf struct for BaseJobResponseAllOfSourceOneOf
type BaseJobResponseAllOfSourceOneOf struct {
	Image *ContainerSource `json:"image,omitempty"`
}

// NewBaseJobResponseAllOfSourceOneOf instantiates a new BaseJobResponseAllOfSourceOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseJobResponseAllOfSourceOneOf() *BaseJobResponseAllOfSourceOneOf {
	this := BaseJobResponseAllOfSourceOneOf{}
	return &this
}

// NewBaseJobResponseAllOfSourceOneOfWithDefaults instantiates a new BaseJobResponseAllOfSourceOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseJobResponseAllOfSourceOneOfWithDefaults() *BaseJobResponseAllOfSourceOneOf {
	this := BaseJobResponseAllOfSourceOneOf{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BaseJobResponseAllOfSourceOneOf) GetImage() ContainerSource {
	if o == nil || IsNil(o.Image) {
		var ret ContainerSource
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseJobResponseAllOfSourceOneOf) GetImageOk() (*ContainerSource, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BaseJobResponseAllOfSourceOneOf) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ContainerSource and assigns it to the Image field.
func (o *BaseJobResponseAllOfSourceOneOf) SetImage(v ContainerSource) {
	o.Image = &v
}

func (o BaseJobResponseAllOfSourceOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseJobResponseAllOfSourceOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableBaseJobResponseAllOfSourceOneOf struct {
	value *BaseJobResponseAllOfSourceOneOf
	isSet bool
}

func (v NullableBaseJobResponseAllOfSourceOneOf) Get() *BaseJobResponseAllOfSourceOneOf {
	return v.value
}

func (v *NullableBaseJobResponseAllOfSourceOneOf) Set(val *BaseJobResponseAllOfSourceOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseJobResponseAllOfSourceOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseJobResponseAllOfSourceOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseJobResponseAllOfSourceOneOf(val *BaseJobResponseAllOfSourceOneOf) *NullableBaseJobResponseAllOfSourceOneOf {
	return &NullableBaseJobResponseAllOfSourceOneOf{value: val, isSet: true}
}

func (v NullableBaseJobResponseAllOfSourceOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseJobResponseAllOfSourceOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}