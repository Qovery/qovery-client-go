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

// checks if the ContainerImageCheckRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerImageCheckRequest{}

// ContainerImageCheckRequest struct for ContainerImageCheckRequest
type ContainerImageCheckRequest struct {
	RegistryId           *string `json:"registry_id,omitempty"`
	ImageName            string  `json:"image_name"`
	Tag                  string  `json:"tag"`
	AdditionalProperties map[string]interface{}
}

type _ContainerImageCheckRequest ContainerImageCheckRequest

// NewContainerImageCheckRequest instantiates a new ContainerImageCheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerImageCheckRequest(imageName string, tag string) *ContainerImageCheckRequest {
	this := ContainerImageCheckRequest{}
	this.ImageName = imageName
	this.Tag = tag
	return &this
}

// NewContainerImageCheckRequestWithDefaults instantiates a new ContainerImageCheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerImageCheckRequestWithDefaults() *ContainerImageCheckRequest {
	this := ContainerImageCheckRequest{}
	return &this
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *ContainerImageCheckRequest) GetRegistryId() string {
	if o == nil || IsNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageCheckRequest) GetRegistryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryId) {
		return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *ContainerImageCheckRequest) HasRegistryId() bool {
	if o != nil && !IsNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *ContainerImageCheckRequest) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetImageName returns the ImageName field value
func (o *ContainerImageCheckRequest) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *ContainerImageCheckRequest) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *ContainerImageCheckRequest) SetImageName(v string) {
	o.ImageName = v
}

// GetTag returns the Tag field value
func (o *ContainerImageCheckRequest) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ContainerImageCheckRequest) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ContainerImageCheckRequest) SetTag(v string) {
	o.Tag = v
}

func (o ContainerImageCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerImageCheckRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegistryId) {
		toSerialize["registry_id"] = o.RegistryId
	}
	toSerialize["image_name"] = o.ImageName
	toSerialize["tag"] = o.Tag

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContainerImageCheckRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image_name",
		"tag",
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

	varContainerImageCheckRequest := _ContainerImageCheckRequest{}

	err = json.Unmarshal(data, &varContainerImageCheckRequest)

	if err != nil {
		return err
	}

	*o = ContainerImageCheckRequest(varContainerImageCheckRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "registry_id")
		delete(additionalProperties, "image_name")
		delete(additionalProperties, "tag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerImageCheckRequest struct {
	value *ContainerImageCheckRequest
	isSet bool
}

func (v NullableContainerImageCheckRequest) Get() *ContainerImageCheckRequest {
	return v.value
}

func (v *NullableContainerImageCheckRequest) Set(val *ContainerImageCheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerImageCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerImageCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerImageCheckRequest(val *ContainerImageCheckRequest) *NullableContainerImageCheckRequest {
	return &NullableContainerImageCheckRequest{value: val, isSet: true}
}

func (v NullableContainerImageCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerImageCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
