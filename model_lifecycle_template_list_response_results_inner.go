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

// checks if the LifecycleTemplateListResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleTemplateListResponseResultsInner{}

// LifecycleTemplateListResponseResultsInner struct for LifecycleTemplateListResponseResultsInner
type LifecycleTemplateListResponseResultsInner struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleTemplateListResponseResultsInner LifecycleTemplateListResponseResultsInner

// NewLifecycleTemplateListResponseResultsInner instantiates a new LifecycleTemplateListResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleTemplateListResponseResultsInner(id string, name string, description string) *LifecycleTemplateListResponseResultsInner {
	this := LifecycleTemplateListResponseResultsInner{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewLifecycleTemplateListResponseResultsInnerWithDefaults instantiates a new LifecycleTemplateListResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleTemplateListResponseResultsInnerWithDefaults() *LifecycleTemplateListResponseResultsInner {
	this := LifecycleTemplateListResponseResultsInner{}
	return &this
}

// GetId returns the Id field value
func (o *LifecycleTemplateListResponseResultsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateListResponseResultsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LifecycleTemplateListResponseResultsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LifecycleTemplateListResponseResultsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateListResponseResultsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LifecycleTemplateListResponseResultsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *LifecycleTemplateListResponseResultsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateListResponseResultsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LifecycleTemplateListResponseResultsInner) SetDescription(v string) {
	o.Description = v
}

func (o LifecycleTemplateListResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleTemplateListResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleTemplateListResponseResultsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
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

	varLifecycleTemplateListResponseResultsInner := _LifecycleTemplateListResponseResultsInner{}

	err = json.Unmarshal(data, &varLifecycleTemplateListResponseResultsInner)

	if err != nil {
		return err
	}

	*o = LifecycleTemplateListResponseResultsInner(varLifecycleTemplateListResponseResultsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleTemplateListResponseResultsInner struct {
	value *LifecycleTemplateListResponseResultsInner
	isSet bool
}

func (v NullableLifecycleTemplateListResponseResultsInner) Get() *LifecycleTemplateListResponseResultsInner {
	return v.value
}

func (v *NullableLifecycleTemplateListResponseResultsInner) Set(val *LifecycleTemplateListResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleTemplateListResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleTemplateListResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleTemplateListResponseResultsInner(val *LifecycleTemplateListResponseResultsInner) *NullableLifecycleTemplateListResponseResultsInner {
	return &NullableLifecycleTemplateListResponseResultsInner{value: val, isSet: true}
}

func (v NullableLifecycleTemplateListResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleTemplateListResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
