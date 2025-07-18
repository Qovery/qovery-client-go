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

// checks if the OrganizationEventTargetResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationEventTargetResponseList{}

// OrganizationEventTargetResponseList struct for OrganizationEventTargetResponseList
type OrganizationEventTargetResponseList struct {
	Targets              []ClusterCloudProviderInfoCredentials `json:"targets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationEventTargetResponseList OrganizationEventTargetResponseList

// NewOrganizationEventTargetResponseList instantiates a new OrganizationEventTargetResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEventTargetResponseList() *OrganizationEventTargetResponseList {
	this := OrganizationEventTargetResponseList{}
	return &this
}

// NewOrganizationEventTargetResponseListWithDefaults instantiates a new OrganizationEventTargetResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEventTargetResponseListWithDefaults() *OrganizationEventTargetResponseList {
	this := OrganizationEventTargetResponseList{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *OrganizationEventTargetResponseList) GetTargets() []ClusterCloudProviderInfoCredentials {
	if o == nil || IsNil(o.Targets) {
		var ret []ClusterCloudProviderInfoCredentials
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventTargetResponseList) GetTargetsOk() ([]ClusterCloudProviderInfoCredentials, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *OrganizationEventTargetResponseList) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []ClusterCloudProviderInfoCredentials and assigns it to the Targets field.
func (o *OrganizationEventTargetResponseList) SetTargets(v []ClusterCloudProviderInfoCredentials) {
	o.Targets = v
}

func (o OrganizationEventTargetResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationEventTargetResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationEventTargetResponseList) UnmarshalJSON(data []byte) (err error) {
	varOrganizationEventTargetResponseList := _OrganizationEventTargetResponseList{}

	err = json.Unmarshal(data, &varOrganizationEventTargetResponseList)

	if err != nil {
		return err
	}

	*o = OrganizationEventTargetResponseList(varOrganizationEventTargetResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "targets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationEventTargetResponseList struct {
	value *OrganizationEventTargetResponseList
	isSet bool
}

func (v NullableOrganizationEventTargetResponseList) Get() *OrganizationEventTargetResponseList {
	return v.value
}

func (v *NullableOrganizationEventTargetResponseList) Set(val *OrganizationEventTargetResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventTargetResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventTargetResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventTargetResponseList(val *OrganizationEventTargetResponseList) *NullableOrganizationEventTargetResponseList {
	return &NullableOrganizationEventTargetResponseList{value: val, isSet: true}
}

func (v NullableOrganizationEventTargetResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventTargetResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
