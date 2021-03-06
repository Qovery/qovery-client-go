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

// DeploymentHistoryAllOf struct for DeploymentHistoryAllOf
type DeploymentHistoryAllOf struct {
	// name of the service
	Name   *string                      `json:"name,omitempty"`
	Commit *Commit                      `json:"commit,omitempty"`
	Status *DeploymentHistoryStatusEnum `json:"status,omitempty"`
}

// NewDeploymentHistoryAllOf instantiates a new DeploymentHistoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryAllOf() *DeploymentHistoryAllOf {
	this := DeploymentHistoryAllOf{}
	return &this
}

// NewDeploymentHistoryAllOfWithDefaults instantiates a new DeploymentHistoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryAllOfWithDefaults() *DeploymentHistoryAllOf {
	this := DeploymentHistoryAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryAllOf) SetName(v string) {
	o.Name = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistoryAllOf) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAllOf) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryAllOf) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *DeploymentHistoryAllOf) SetCommit(v Commit) {
	o.Commit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryAllOf) GetStatus() DeploymentHistoryStatusEnum {
	if o == nil || o.Status == nil {
		var ret DeploymentHistoryStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAllOf) GetStatusOk() (*DeploymentHistoryStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentHistoryStatusEnum and assigns it to the Status field.
func (o *DeploymentHistoryAllOf) SetStatus(v DeploymentHistoryStatusEnum) {
	o.Status = &v
}

func (o DeploymentHistoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryAllOf struct {
	value *DeploymentHistoryAllOf
	isSet bool
}

func (v NullableDeploymentHistoryAllOf) Get() *DeploymentHistoryAllOf {
	return v.value
}

func (v *NullableDeploymentHistoryAllOf) Set(val *DeploymentHistoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryAllOf(val *DeploymentHistoryAllOf) *NullableDeploymentHistoryAllOf {
	return &NullableDeploymentHistoryAllOf{value: val, isSet: true}
}

func (v NullableDeploymentHistoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
