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

// checks if the Annotation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Annotation{}

// Annotation struct for Annotation
type Annotation struct {
	Key                  string `json:"key"`
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _Annotation Annotation

// NewAnnotation instantiates a new Annotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotation(key string, value string) *Annotation {
	this := Annotation{}
	this.Key = key
	this.Value = value
	return &this
}

// NewAnnotationWithDefaults instantiates a new Annotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationWithDefaults() *Annotation {
	this := Annotation{}
	return &this
}

// GetKey returns the Key field value
func (o *Annotation) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Annotation) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *Annotation) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Annotation) SetValue(v string) {
	o.Value = v
}

func (o Annotation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Annotation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Annotation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varAnnotation := _Annotation{}

	err = json.Unmarshal(data, &varAnnotation)

	if err != nil {
		return err
	}

	*o = Annotation(varAnnotation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnnotation struct {
	value *Annotation
	isSet bool
}

func (v NullableAnnotation) Get() *Annotation {
	return v.value
}

func (v *NullableAnnotation) Set(val *Annotation) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotation(val *Annotation) *NullableAnnotation {
	return &NullableAnnotation{value: val, isSet: true}
}

func (v NullableAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
