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

// checks if the ContainerRegistryAssociatedServicesResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRegistryAssociatedServicesResponseList{}

// ContainerRegistryAssociatedServicesResponseList struct for ContainerRegistryAssociatedServicesResponseList
type ContainerRegistryAssociatedServicesResponseList struct {
	Results              *ContainerRegistryAssociatedServicesResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContainerRegistryAssociatedServicesResponseList ContainerRegistryAssociatedServicesResponseList

// NewContainerRegistryAssociatedServicesResponseList instantiates a new ContainerRegistryAssociatedServicesResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistryAssociatedServicesResponseList() *ContainerRegistryAssociatedServicesResponseList {
	this := ContainerRegistryAssociatedServicesResponseList{}
	return &this
}

// NewContainerRegistryAssociatedServicesResponseListWithDefaults instantiates a new ContainerRegistryAssociatedServicesResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryAssociatedServicesResponseListWithDefaults() *ContainerRegistryAssociatedServicesResponseList {
	this := ContainerRegistryAssociatedServicesResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ContainerRegistryAssociatedServicesResponseList) GetResults() ContainerRegistryAssociatedServicesResponse {
	if o == nil || IsNil(o.Results) {
		var ret ContainerRegistryAssociatedServicesResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryAssociatedServicesResponseList) GetResultsOk() (*ContainerRegistryAssociatedServicesResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ContainerRegistryAssociatedServicesResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given ContainerRegistryAssociatedServicesResponse and assigns it to the Results field.
func (o *ContainerRegistryAssociatedServicesResponseList) SetResults(v ContainerRegistryAssociatedServicesResponse) {
	o.Results = &v
}

func (o ContainerRegistryAssociatedServicesResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRegistryAssociatedServicesResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContainerRegistryAssociatedServicesResponseList) UnmarshalJSON(data []byte) (err error) {
	varContainerRegistryAssociatedServicesResponseList := _ContainerRegistryAssociatedServicesResponseList{}

	err = json.Unmarshal(data, &varContainerRegistryAssociatedServicesResponseList)

	if err != nil {
		return err
	}

	*o = ContainerRegistryAssociatedServicesResponseList(varContainerRegistryAssociatedServicesResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerRegistryAssociatedServicesResponseList struct {
	value *ContainerRegistryAssociatedServicesResponseList
	isSet bool
}

func (v NullableContainerRegistryAssociatedServicesResponseList) Get() *ContainerRegistryAssociatedServicesResponseList {
	return v.value
}

func (v *NullableContainerRegistryAssociatedServicesResponseList) Set(val *ContainerRegistryAssociatedServicesResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistryAssociatedServicesResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistryAssociatedServicesResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistryAssociatedServicesResponseList(val *ContainerRegistryAssociatedServicesResponseList) *NullableContainerRegistryAssociatedServicesResponseList {
	return &NullableContainerRegistryAssociatedServicesResponseList{value: val, isSet: true}
}

func (v NullableContainerRegistryAssociatedServicesResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistryAssociatedServicesResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
