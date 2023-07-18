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

// VariableResponseAllOf struct for VariableResponseAllOf
type VariableResponseAllOf struct {
	OverriddenVariable *EnvironmentVariableOverride `json:"overridden_variable,omitempty"`
	AliasedVariable    *EnvironmentVariableAlias    `json:"aliased_variable,omitempty"`
	OverriddenSecret   *SecretOverride              `json:"overridden_secret,omitempty"`
	AliasedSecret      *SecretAlias                 `json:"aliased_secret,omitempty"`
	Scope              APIVariableScopeEnum         `json:"scope"`
	VariableType       *APIVariableTypeEnum         `json:"variable_type,omitempty"`
	// present only for `BUILT_IN` variable
	ServiceId *string `json:"service_id,omitempty"`
	// present only for `BUILT_IN` variable
	ServiceName *string                `json:"service_name,omitempty"`
	ServiceType *LinkedServiceTypeEnum `json:"service_type,omitempty"`
	// Entity that created/own the variable (i.e: Qovery, Doppler)
	OwnedBy *string `json:"owned_by,omitempty"`
}

// NewVariableResponseAllOf instantiates a new VariableResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableResponseAllOf(scope APIVariableScopeEnum) *VariableResponseAllOf {
	this := VariableResponseAllOf{}
	this.Scope = scope
	return &this
}

// NewVariableResponseAllOfWithDefaults instantiates a new VariableResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableResponseAllOfWithDefaults() *VariableResponseAllOf {
	this := VariableResponseAllOf{}
	return &this
}

// GetOverriddenVariable returns the OverriddenVariable field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetOverriddenVariable() EnvironmentVariableOverride {
	if o == nil || o.OverriddenVariable == nil {
		var ret EnvironmentVariableOverride
		return ret
	}
	return *o.OverriddenVariable
}

// GetOverriddenVariableOk returns a tuple with the OverriddenVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetOverriddenVariableOk() (*EnvironmentVariableOverride, bool) {
	if o == nil || o.OverriddenVariable == nil {
		return nil, false
	}
	return o.OverriddenVariable, true
}

// HasOverriddenVariable returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasOverriddenVariable() bool {
	if o != nil && o.OverriddenVariable != nil {
		return true
	}

	return false
}

// SetOverriddenVariable gets a reference to the given EnvironmentVariableOverride and assigns it to the OverriddenVariable field.
func (o *VariableResponseAllOf) SetOverriddenVariable(v EnvironmentVariableOverride) {
	o.OverriddenVariable = &v
}

// GetAliasedVariable returns the AliasedVariable field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetAliasedVariable() EnvironmentVariableAlias {
	if o == nil || o.AliasedVariable == nil {
		var ret EnvironmentVariableAlias
		return ret
	}
	return *o.AliasedVariable
}

// GetAliasedVariableOk returns a tuple with the AliasedVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetAliasedVariableOk() (*EnvironmentVariableAlias, bool) {
	if o == nil || o.AliasedVariable == nil {
		return nil, false
	}
	return o.AliasedVariable, true
}

// HasAliasedVariable returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasAliasedVariable() bool {
	if o != nil && o.AliasedVariable != nil {
		return true
	}

	return false
}

// SetAliasedVariable gets a reference to the given EnvironmentVariableAlias and assigns it to the AliasedVariable field.
func (o *VariableResponseAllOf) SetAliasedVariable(v EnvironmentVariableAlias) {
	o.AliasedVariable = &v
}

// GetOverriddenSecret returns the OverriddenSecret field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetOverriddenSecret() SecretOverride {
	if o == nil || o.OverriddenSecret == nil {
		var ret SecretOverride
		return ret
	}
	return *o.OverriddenSecret
}

// GetOverriddenSecretOk returns a tuple with the OverriddenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetOverriddenSecretOk() (*SecretOverride, bool) {
	if o == nil || o.OverriddenSecret == nil {
		return nil, false
	}
	return o.OverriddenSecret, true
}

// HasOverriddenSecret returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasOverriddenSecret() bool {
	if o != nil && o.OverriddenSecret != nil {
		return true
	}

	return false
}

// SetOverriddenSecret gets a reference to the given SecretOverride and assigns it to the OverriddenSecret field.
func (o *VariableResponseAllOf) SetOverriddenSecret(v SecretOverride) {
	o.OverriddenSecret = &v
}

// GetAliasedSecret returns the AliasedSecret field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetAliasedSecret() SecretAlias {
	if o == nil || o.AliasedSecret == nil {
		var ret SecretAlias
		return ret
	}
	return *o.AliasedSecret
}

// GetAliasedSecretOk returns a tuple with the AliasedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetAliasedSecretOk() (*SecretAlias, bool) {
	if o == nil || o.AliasedSecret == nil {
		return nil, false
	}
	return o.AliasedSecret, true
}

// HasAliasedSecret returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasAliasedSecret() bool {
	if o != nil && o.AliasedSecret != nil {
		return true
	}

	return false
}

// SetAliasedSecret gets a reference to the given SecretAlias and assigns it to the AliasedSecret field.
func (o *VariableResponseAllOf) SetAliasedSecret(v SecretAlias) {
	o.AliasedSecret = &v
}

// GetScope returns the Scope field value
func (o *VariableResponseAllOf) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VariableResponseAllOf) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetVariableType returns the VariableType field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetVariableType() APIVariableTypeEnum {
	if o == nil || o.VariableType == nil {
		var ret APIVariableTypeEnum
		return ret
	}
	return *o.VariableType
}

// GetVariableTypeOk returns a tuple with the VariableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetVariableTypeOk() (*APIVariableTypeEnum, bool) {
	if o == nil || o.VariableType == nil {
		return nil, false
	}
	return o.VariableType, true
}

// HasVariableType returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasVariableType() bool {
	if o != nil && o.VariableType != nil {
		return true
	}

	return false
}

// SetVariableType gets a reference to the given APIVariableTypeEnum and assigns it to the VariableType field.
func (o *VariableResponseAllOf) SetVariableType(v APIVariableTypeEnum) {
	o.VariableType = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *VariableResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *VariableResponseAllOf) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetServiceType() LinkedServiceTypeEnum {
	if o == nil || o.ServiceType == nil {
		var ret LinkedServiceTypeEnum
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given LinkedServiceTypeEnum and assigns it to the ServiceType field.
func (o *VariableResponseAllOf) SetServiceType(v LinkedServiceTypeEnum) {
	o.ServiceType = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *VariableResponseAllOf) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableResponseAllOf) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *VariableResponseAllOf) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *VariableResponseAllOf) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

func (o VariableResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OverriddenVariable != nil {
		toSerialize["overridden_variable"] = o.OverriddenVariable
	}
	if o.AliasedVariable != nil {
		toSerialize["aliased_variable"] = o.AliasedVariable
	}
	if o.OverriddenSecret != nil {
		toSerialize["overridden_secret"] = o.OverriddenSecret
	}
	if o.AliasedSecret != nil {
		toSerialize["aliased_secret"] = o.AliasedSecret
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if o.VariableType != nil {
		toSerialize["variable_type"] = o.VariableType
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	if o.OwnedBy != nil {
		toSerialize["owned_by"] = o.OwnedBy
	}
	return json.Marshal(toSerialize)
}

type NullableVariableResponseAllOf struct {
	value *VariableResponseAllOf
	isSet bool
}

func (v NullableVariableResponseAllOf) Get() *VariableResponseAllOf {
	return v.value
}

func (v *NullableVariableResponseAllOf) Set(val *VariableResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableResponseAllOf(val *VariableResponseAllOf) *NullableVariableResponseAllOf {
	return &NullableVariableResponseAllOf{value: val, isSet: true}
}

func (v NullableVariableResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
