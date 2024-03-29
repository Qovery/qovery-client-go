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

// checks if the ContainerDeployRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerDeployRequest{}

// ContainerDeployRequest struct for ContainerDeployRequest
type ContainerDeployRequest struct {
	// Image tag to deploy
	ImageTag string `json:"image_tag"`
}

// NewContainerDeployRequest instantiates a new ContainerDeployRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerDeployRequest(imageTag string) *ContainerDeployRequest {
	this := ContainerDeployRequest{}
	this.ImageTag = imageTag
	return &this
}

// NewContainerDeployRequestWithDefaults instantiates a new ContainerDeployRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerDeployRequestWithDefaults() *ContainerDeployRequest {
	this := ContainerDeployRequest{}
	return &this
}

// GetImageTag returns the ImageTag field value
func (o *ContainerDeployRequest) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *ContainerDeployRequest) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *ContainerDeployRequest) SetImageTag(v string) {
	o.ImageTag = v
}

func (o ContainerDeployRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerDeployRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image_tag"] = o.ImageTag
	return toSerialize, nil
}

type NullableContainerDeployRequest struct {
	value *ContainerDeployRequest
	isSet bool
}

func (v NullableContainerDeployRequest) Get() *ContainerDeployRequest {
	return v.value
}

func (v *NullableContainerDeployRequest) Set(val *ContainerDeployRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerDeployRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerDeployRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerDeployRequest(val *ContainerDeployRequest) *NullableContainerDeployRequest {
	return &NullableContainerDeployRequest{value: val, isSet: true}
}

func (v NullableContainerDeployRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerDeployRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
