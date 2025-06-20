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

// checks if the OrganizationAvailableRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAvailableRole{}

// OrganizationAvailableRole struct for OrganizationAvailableRole
type OrganizationAvailableRole struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationAvailableRole OrganizationAvailableRole

// NewOrganizationAvailableRole instantiates a new OrganizationAvailableRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAvailableRole(id string, name string) *OrganizationAvailableRole {
	this := OrganizationAvailableRole{}
	this.Id = id
	this.Name = name
	return &this
}

// NewOrganizationAvailableRoleWithDefaults instantiates a new OrganizationAvailableRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAvailableRoleWithDefaults() *OrganizationAvailableRole {
	this := OrganizationAvailableRole{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationAvailableRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationAvailableRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationAvailableRole) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OrganizationAvailableRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationAvailableRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationAvailableRole) SetName(v string) {
	o.Name = v
}

func (o OrganizationAvailableRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAvailableRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationAvailableRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varOrganizationAvailableRole := _OrganizationAvailableRole{}

	err = json.Unmarshal(data, &varOrganizationAvailableRole)

	if err != nil {
		return err
	}

	*o = OrganizationAvailableRole(varOrganizationAvailableRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationAvailableRole struct {
	value *OrganizationAvailableRole
	isSet bool
}

func (v NullableOrganizationAvailableRole) Get() *OrganizationAvailableRole {
	return v.value
}

func (v *NullableOrganizationAvailableRole) Set(val *OrganizationAvailableRole) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAvailableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAvailableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAvailableRole(val *OrganizationAvailableRole) *NullableOrganizationAvailableRole {
	return &NullableOrganizationAvailableRole{value: val, isSet: true}
}

func (v NullableOrganizationAvailableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAvailableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
