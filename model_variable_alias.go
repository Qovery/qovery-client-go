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

// VariableAlias struct for VariableAlias
type VariableAlias struct {
	Id           string               `json:"id"`
	Key          string               `json:"key"`
	Value        NullableString       `json:"value,omitempty"`
	MountPath    string               `json:"mount_path"`
	Scope        APIVariableScopeEnum `json:"scope"`
	VariableType APIVariableTypeEnum  `json:"variable_type"`
}

// NewVariableAlias instantiates a new VariableAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableAlias(id string, key string, mountPath string, scope APIVariableScopeEnum, variableType APIVariableTypeEnum) *VariableAlias {
	this := VariableAlias{}
	this.Id = id
	this.Key = key
	this.MountPath = mountPath
	this.Scope = scope
	this.VariableType = variableType
	return &this
}

// NewVariableAliasWithDefaults instantiates a new VariableAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableAliasWithDefaults() *VariableAlias {
	this := VariableAlias{}
	return &this
}

// GetId returns the Id field value
func (o *VariableAlias) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VariableAlias) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *VariableAlias) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *VariableAlias) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableAlias) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableAlias) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *VariableAlias) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *VariableAlias) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *VariableAlias) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *VariableAlias) UnsetValue() {
	o.Value.Unset()
}

// GetMountPath returns the MountPath field value
func (o *VariableAlias) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPath, true
}

// SetMountPath sets field value
func (o *VariableAlias) SetMountPath(v string) {
	o.MountPath = v
}

// GetScope returns the Scope field value
func (o *VariableAlias) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VariableAlias) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetVariableType returns the VariableType field value
func (o *VariableAlias) GetVariableType() APIVariableTypeEnum {
	if o == nil {
		var ret APIVariableTypeEnum
		return ret
	}

	return o.VariableType
}

// GetVariableTypeOk returns a tuple with the VariableType field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetVariableTypeOk() (*APIVariableTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableType, true
}

// SetVariableType sets field value
func (o *VariableAlias) SetVariableType(v APIVariableTypeEnum) {
	o.VariableType = v
}

func (o VariableAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if true {
		toSerialize["mount_path"] = o.MountPath
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["variable_type"] = o.VariableType
	}
	return json.Marshal(toSerialize)
}

type NullableVariableAlias struct {
	value *VariableAlias
	isSet bool
}

func (v NullableVariableAlias) Get() *VariableAlias {
	return v.value
}

func (v *NullableVariableAlias) Set(val *VariableAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableAlias(val *VariableAlias) *NullableVariableAlias {
	return &NullableVariableAlias{value: val, isSet: true}
}

func (v NullableVariableAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}