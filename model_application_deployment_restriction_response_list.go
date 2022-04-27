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

// ApplicationDeploymentRestrictionResponseList struct for ApplicationDeploymentRestrictionResponseList
type ApplicationDeploymentRestrictionResponseList struct {
	DeploymentRestrictions []ApplicationDeploymentRestriction `json:"deployment_restrictions,omitempty"`
}

// NewApplicationDeploymentRestrictionResponseList instantiates a new ApplicationDeploymentRestrictionResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationDeploymentRestrictionResponseList() *ApplicationDeploymentRestrictionResponseList {
	this := ApplicationDeploymentRestrictionResponseList{}
	return &this
}

// NewApplicationDeploymentRestrictionResponseListWithDefaults instantiates a new ApplicationDeploymentRestrictionResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDeploymentRestrictionResponseListWithDefaults() *ApplicationDeploymentRestrictionResponseList {
	this := ApplicationDeploymentRestrictionResponseList{}
	return &this
}

// GetDeploymentRestrictions returns the DeploymentRestrictions field value if set, zero value otherwise.
func (o *ApplicationDeploymentRestrictionResponseList) GetDeploymentRestrictions() []ApplicationDeploymentRestriction {
	if o == nil || o.DeploymentRestrictions == nil {
		var ret []ApplicationDeploymentRestriction
		return ret
	}
	return o.DeploymentRestrictions
}

// GetDeploymentRestrictionsOk returns a tuple with the DeploymentRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDeploymentRestrictionResponseList) GetDeploymentRestrictionsOk() ([]ApplicationDeploymentRestriction, bool) {
	if o == nil || o.DeploymentRestrictions == nil {
		return nil, false
	}
	return o.DeploymentRestrictions, true
}

// HasDeploymentRestrictions returns a boolean if a field has been set.
func (o *ApplicationDeploymentRestrictionResponseList) HasDeploymentRestrictions() bool {
	if o != nil && o.DeploymentRestrictions != nil {
		return true
	}

	return false
}

// SetDeploymentRestrictions gets a reference to the given []ApplicationDeploymentRestriction and assigns it to the DeploymentRestrictions field.
func (o *ApplicationDeploymentRestrictionResponseList) SetDeploymentRestrictions(v []ApplicationDeploymentRestriction) {
	o.DeploymentRestrictions = v
}

func (o ApplicationDeploymentRestrictionResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentRestrictions != nil {
		toSerialize["deployment_restrictions"] = o.DeploymentRestrictions
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationDeploymentRestrictionResponseList struct {
	value *ApplicationDeploymentRestrictionResponseList
	isSet bool
}

func (v NullableApplicationDeploymentRestrictionResponseList) Get() *ApplicationDeploymentRestrictionResponseList {
	return v.value
}

func (v *NullableApplicationDeploymentRestrictionResponseList) Set(val *ApplicationDeploymentRestrictionResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDeploymentRestrictionResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDeploymentRestrictionResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDeploymentRestrictionResponseList(val *ApplicationDeploymentRestrictionResponseList) *NullableApplicationDeploymentRestrictionResponseList {
	return &NullableApplicationDeploymentRestrictionResponseList{value: val, isSet: true}
}

func (v NullableApplicationDeploymentRestrictionResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDeploymentRestrictionResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
