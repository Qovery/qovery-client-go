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

// ClusterRoutingTableResultsInner struct for ClusterRoutingTableResultsInner
type ClusterRoutingTableResultsInner struct {
	Destination *string `json:"destination,omitempty"`
	Target      *string `json:"target,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewClusterRoutingTableResultsInner instantiates a new ClusterRoutingTableResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRoutingTableResultsInner() *ClusterRoutingTableResultsInner {
	this := ClusterRoutingTableResultsInner{}
	return &this
}

// NewClusterRoutingTableResultsInnerWithDefaults instantiates a new ClusterRoutingTableResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRoutingTableResultsInnerWithDefaults() *ClusterRoutingTableResultsInner {
	this := ClusterRoutingTableResultsInner{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ClusterRoutingTableResultsInner) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableResultsInner) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ClusterRoutingTableResultsInner) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *ClusterRoutingTableResultsInner) SetDestination(v string) {
	o.Destination = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ClusterRoutingTableResultsInner) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableResultsInner) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ClusterRoutingTableResultsInner) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ClusterRoutingTableResultsInner) SetTarget(v string) {
	o.Target = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterRoutingTableResultsInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRoutingTableResultsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterRoutingTableResultsInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterRoutingTableResultsInner) SetDescription(v string) {
	o.Description = &v
}

func (o ClusterRoutingTableResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableClusterRoutingTableResultsInner struct {
	value *ClusterRoutingTableResultsInner
	isSet bool
}

func (v NullableClusterRoutingTableResultsInner) Get() *ClusterRoutingTableResultsInner {
	return v.value
}

func (v *NullableClusterRoutingTableResultsInner) Set(val *ClusterRoutingTableResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRoutingTableResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRoutingTableResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRoutingTableResultsInner(val *ClusterRoutingTableResultsInner) *NullableClusterRoutingTableResultsInner {
	return &NullableClusterRoutingTableResultsInner{value: val, isSet: true}
}

func (v NullableClusterRoutingTableResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRoutingTableResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
