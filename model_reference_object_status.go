/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ReferenceObjectStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceObjectStatus{}

// ReferenceObjectStatus struct for ReferenceObjectStatus
type ReferenceObjectStatus struct {
	Id                      string                      `json:"id"`
	State                   StateEnum                   `json:"state"`
	ServiceDeploymentStatus ServiceDeploymentStatusEnum `json:"service_deployment_status"`
	LastDeploymentDate      *time.Time                  `json:"last_deployment_date,omitempty"`
	IsPartLastDeployment    *bool                       `json:"is_part_last_deployment,omitempty"`
	Steps                   *ServiceStepMetrics         `json:"steps,omitempty"`
}

type _ReferenceObjectStatus ReferenceObjectStatus

// NewReferenceObjectStatus instantiates a new ReferenceObjectStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceObjectStatus(id string, state StateEnum, serviceDeploymentStatus ServiceDeploymentStatusEnum) *ReferenceObjectStatus {
	this := ReferenceObjectStatus{}
	this.Id = id
	this.State = state
	this.ServiceDeploymentStatus = serviceDeploymentStatus
	return &this
}

// NewReferenceObjectStatusWithDefaults instantiates a new ReferenceObjectStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceObjectStatusWithDefaults() *ReferenceObjectStatus {
	this := ReferenceObjectStatus{}
	return &this
}

// GetId returns the Id field value
func (o *ReferenceObjectStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReferenceObjectStatus) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *ReferenceObjectStatus) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReferenceObjectStatus) SetState(v StateEnum) {
	o.State = v
}

// GetServiceDeploymentStatus returns the ServiceDeploymentStatus field value
func (o *ReferenceObjectStatus) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum {
	if o == nil {
		var ret ServiceDeploymentStatusEnum
		return ret
	}

	return o.ServiceDeploymentStatus
}

// GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceDeploymentStatus, true
}

// SetServiceDeploymentStatus sets field value
func (o *ReferenceObjectStatus) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum) {
	o.ServiceDeploymentStatus = v
}

// GetLastDeploymentDate returns the LastDeploymentDate field value if set, zero value otherwise.
func (o *ReferenceObjectStatus) GetLastDeploymentDate() time.Time {
	if o == nil || IsNil(o.LastDeploymentDate) {
		var ret time.Time
		return ret
	}
	return *o.LastDeploymentDate
}

// GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetLastDeploymentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastDeploymentDate) {
		return nil, false
	}
	return o.LastDeploymentDate, true
}

// HasLastDeploymentDate returns a boolean if a field has been set.
func (o *ReferenceObjectStatus) HasLastDeploymentDate() bool {
	if o != nil && !IsNil(o.LastDeploymentDate) {
		return true
	}

	return false
}

// SetLastDeploymentDate gets a reference to the given time.Time and assigns it to the LastDeploymentDate field.
func (o *ReferenceObjectStatus) SetLastDeploymentDate(v time.Time) {
	o.LastDeploymentDate = &v
}

// GetIsPartLastDeployment returns the IsPartLastDeployment field value if set, zero value otherwise.
func (o *ReferenceObjectStatus) GetIsPartLastDeployment() bool {
	if o == nil || IsNil(o.IsPartLastDeployment) {
		var ret bool
		return ret
	}
	return *o.IsPartLastDeployment
}

// GetIsPartLastDeploymentOk returns a tuple with the IsPartLastDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetIsPartLastDeploymentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartLastDeployment) {
		return nil, false
	}
	return o.IsPartLastDeployment, true
}

// HasIsPartLastDeployment returns a boolean if a field has been set.
func (o *ReferenceObjectStatus) HasIsPartLastDeployment() bool {
	if o != nil && !IsNil(o.IsPartLastDeployment) {
		return true
	}

	return false
}

// SetIsPartLastDeployment gets a reference to the given bool and assigns it to the IsPartLastDeployment field.
func (o *ReferenceObjectStatus) SetIsPartLastDeployment(v bool) {
	o.IsPartLastDeployment = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ReferenceObjectStatus) GetSteps() ServiceStepMetrics {
	if o == nil || IsNil(o.Steps) {
		var ret ServiceStepMetrics
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetStepsOk() (*ServiceStepMetrics, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ReferenceObjectStatus) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given ServiceStepMetrics and assigns it to the Steps field.
func (o *ReferenceObjectStatus) SetSteps(v ServiceStepMetrics) {
	o.Steps = &v
}

func (o ReferenceObjectStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceObjectStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["service_deployment_status"] = o.ServiceDeploymentStatus
	if !IsNil(o.LastDeploymentDate) {
		toSerialize["last_deployment_date"] = o.LastDeploymentDate
	}
	if !IsNil(o.IsPartLastDeployment) {
		toSerialize["is_part_last_deployment"] = o.IsPartLastDeployment
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

func (o *ReferenceObjectStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"service_deployment_status",
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

	varReferenceObjectStatus := _ReferenceObjectStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReferenceObjectStatus)

	if err != nil {
		return err
	}

	*o = ReferenceObjectStatus(varReferenceObjectStatus)

	return err
}

type NullableReferenceObjectStatus struct {
	value *ReferenceObjectStatus
	isSet bool
}

func (v NullableReferenceObjectStatus) Get() *ReferenceObjectStatus {
	return v.value
}

func (v *NullableReferenceObjectStatus) Set(val *ReferenceObjectStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceObjectStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceObjectStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceObjectStatus(val *ReferenceObjectStatus) *NullableReferenceObjectStatus {
	return &NullableReferenceObjectStatus{value: val, isSet: true}
}

func (v NullableReferenceObjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceObjectStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
