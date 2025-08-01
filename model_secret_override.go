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

// checks if the SecretOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretOverride{}

// SecretOverride struct for SecretOverride
type SecretOverride struct {
	Id           string               `json:"id"`
	Key          string               `json:"key"`
	MountPath    string               `json:"mount_path"`
	Scope        APIVariableScopeEnum `json:"scope"`
	VariableType APIVariableTypeEnum  `json:"variable_type"`
	// optional variable description (255 characters maximum)
	Description               NullableString `json:"description,omitempty"`
	EnableInterpolationInFile NullableBool   `json:"enable_interpolation_in_file,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _SecretOverride SecretOverride

// NewSecretOverride instantiates a new SecretOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretOverride(id string, key string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum) *SecretOverride {
	this := SecretOverride{}
	this.Id = id
	this.Key = key
	this.MountPath = mountPath
	this.Scope = scope
	this.VariableType = variableType
	return &this
}

// NewSecretOverrideWithDefaults instantiates a new SecretOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretOverrideWithDefaults() *SecretOverride {
	this := SecretOverride{}
	return &this
}

// GetId returns the Id field value
func (o *SecretOverride) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretOverride) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretOverride) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *SecretOverride) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecretOverride) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SecretOverride) SetKey(v string) {
	o.Key = v
}

// GetMountPath returns the MountPath field value
func (o *SecretOverride) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *SecretOverride) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPath, true
}

// SetMountPath sets field value
func (o *SecretOverride) SetMountPath(v string) {
	o.MountPath = v
}

// GetScope returns the Scope field value
func (o *SecretOverride) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SecretOverride) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SecretOverride) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetVariableType returns the VariableType field value
func (o *SecretOverride) GetVariableType() APIVariableTypeEnum {
	if o == nil {
		var ret APIVariableTypeEnum
		return ret
	}

	return o.VariableType
}

// GetVariableTypeOk returns a tuple with the VariableType field value
// and a boolean to check if the value has been set.
func (o *SecretOverride) GetVariableTypeOk() (*APIVariableTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableType, true
}

// SetVariableType sets field value
func (o *SecretOverride) SetVariableType(v APIVariableTypeEnum) {
	o.VariableType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecretOverride) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecretOverride) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SecretOverride) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SecretOverride) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SecretOverride) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SecretOverride) UnsetDescription() {
	o.Description.Unset()
}

// GetEnableInterpolationInFile returns the EnableInterpolationInFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecretOverride) GetEnableInterpolationInFile() bool {
	if o == nil || IsNil(o.EnableInterpolationInFile.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableInterpolationInFile.Get()
}

// GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecretOverride) GetEnableInterpolationInFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableInterpolationInFile.Get(), o.EnableInterpolationInFile.IsSet()
}

// HasEnableInterpolationInFile returns a boolean if a field has been set.
func (o *SecretOverride) HasEnableInterpolationInFile() bool {
	if o != nil && o.EnableInterpolationInFile.IsSet() {
		return true
	}

	return false
}

// SetEnableInterpolationInFile gets a reference to the given NullableBool and assigns it to the EnableInterpolationInFile field.
func (o *SecretOverride) SetEnableInterpolationInFile(v bool) {
	o.EnableInterpolationInFile.Set(&v)
}

// SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil
func (o *SecretOverride) SetEnableInterpolationInFileNil() {
	o.EnableInterpolationInFile.Set(nil)
}

// UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil
func (o *SecretOverride) UnsetEnableInterpolationInFile() {
	o.EnableInterpolationInFile.Unset()
}

func (o SecretOverride) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["mount_path"] = o.MountPath
	toSerialize["scope"] = o.Scope
	toSerialize["variable_type"] = o.VariableType
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EnableInterpolationInFile.IsSet() {
		toSerialize["enable_interpolation_in_file"] = o.EnableInterpolationInFile.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecretOverride) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"mount_path",
		"scope",
		"variable_type",
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

	varSecretOverride := _SecretOverride{}

	err = json.Unmarshal(data, &varSecretOverride)

	if err != nil {
		return err
	}

	*o = SecretOverride(varSecretOverride)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "mount_path")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "variable_type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enable_interpolation_in_file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecretOverride struct {
	value *SecretOverride
	isSet bool
}

func (v NullableSecretOverride) Get() *SecretOverride {
	return v.value
}

func (v *NullableSecretOverride) Set(val *SecretOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretOverride(val *SecretOverride) *NullableSecretOverride {
	return &NullableSecretOverride{value: val, isSet: true}
}

func (v NullableSecretOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
