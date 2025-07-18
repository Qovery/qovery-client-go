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

// checks if the ContainerVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerVersionResponse{}

// ContainerVersionResponse struct for ContainerVersionResponse
type ContainerVersionResponse struct {
	ImageName            *string  `json:"image_name,omitempty"`
	Versions             []string `json:"versions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContainerVersionResponse ContainerVersionResponse

// NewContainerVersionResponse instantiates a new ContainerVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerVersionResponse() *ContainerVersionResponse {
	this := ContainerVersionResponse{}
	return &this
}

// NewContainerVersionResponseWithDefaults instantiates a new ContainerVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerVersionResponseWithDefaults() *ContainerVersionResponse {
	this := ContainerVersionResponse{}
	return &this
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ContainerVersionResponse) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerVersionResponse) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ContainerVersionResponse) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ContainerVersionResponse) SetImageName(v string) {
	o.ImageName = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ContainerVersionResponse) GetVersions() []string {
	if o == nil || IsNil(o.Versions) {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerVersionResponse) GetVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ContainerVersionResponse) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *ContainerVersionResponse) SetVersions(v []string) {
	o.Versions = v
}

func (o ContainerVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContainerVersionResponse) UnmarshalJSON(data []byte) (err error) {
	varContainerVersionResponse := _ContainerVersionResponse{}

	err = json.Unmarshal(data, &varContainerVersionResponse)

	if err != nil {
		return err
	}

	*o = ContainerVersionResponse(varContainerVersionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "image_name")
		delete(additionalProperties, "versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerVersionResponse struct {
	value *ContainerVersionResponse
	isSet bool
}

func (v NullableContainerVersionResponse) Get() *ContainerVersionResponse {
	return v.value
}

func (v *NullableContainerVersionResponse) Set(val *ContainerVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerVersionResponse(val *ContainerVersionResponse) *NullableContainerVersionResponse {
	return &NullableContainerVersionResponse{value: val, isSet: true}
}

func (v NullableContainerVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
