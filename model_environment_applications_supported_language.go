/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// EnvironmentApplicationsSupportedLanguage struct for EnvironmentApplicationsSupportedLanguage
type EnvironmentApplicationsSupportedLanguage struct {
	Name string `json:"name"`
}

// NewEnvironmentApplicationsSupportedLanguage instantiates a new EnvironmentApplicationsSupportedLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsSupportedLanguage(name string) *EnvironmentApplicationsSupportedLanguage {
	this := EnvironmentApplicationsSupportedLanguage{}
	this.Name = name
	return &this
}

// NewEnvironmentApplicationsSupportedLanguageWithDefaults instantiates a new EnvironmentApplicationsSupportedLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsSupportedLanguageWithDefaults() *EnvironmentApplicationsSupportedLanguage {
	this := EnvironmentApplicationsSupportedLanguage{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentApplicationsSupportedLanguage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsSupportedLanguage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentApplicationsSupportedLanguage) SetName(v string) {
	o.Name = v
}

func (o EnvironmentApplicationsSupportedLanguage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsSupportedLanguage struct {
	value *EnvironmentApplicationsSupportedLanguage
	isSet bool
}

func (v NullableEnvironmentApplicationsSupportedLanguage) Get() *EnvironmentApplicationsSupportedLanguage {
	return v.value
}

func (v *NullableEnvironmentApplicationsSupportedLanguage) Set(val *EnvironmentApplicationsSupportedLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsSupportedLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsSupportedLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsSupportedLanguage(val *EnvironmentApplicationsSupportedLanguage) *NullableEnvironmentApplicationsSupportedLanguage {
	return &NullableEnvironmentApplicationsSupportedLanguage{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsSupportedLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsSupportedLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
