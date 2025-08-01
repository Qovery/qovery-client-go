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
)

// checks if the OrganizationChangePlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationChangePlanRequest{}

// OrganizationChangePlanRequest struct for OrganizationChangePlanRequest
type OrganizationChangePlanRequest struct {
	Plan                 *string `json:"plan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationChangePlanRequest OrganizationChangePlanRequest

// NewOrganizationChangePlanRequest instantiates a new OrganizationChangePlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationChangePlanRequest() *OrganizationChangePlanRequest {
	this := OrganizationChangePlanRequest{}
	return &this
}

// NewOrganizationChangePlanRequestWithDefaults instantiates a new OrganizationChangePlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationChangePlanRequestWithDefaults() *OrganizationChangePlanRequest {
	this := OrganizationChangePlanRequest{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *OrganizationChangePlanRequest) GetPlan() string {
	if o == nil || IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationChangePlanRequest) GetPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *OrganizationChangePlanRequest) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *OrganizationChangePlanRequest) SetPlan(v string) {
	o.Plan = &v
}

func (o OrganizationChangePlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationChangePlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationChangePlanRequest) UnmarshalJSON(data []byte) (err error) {
	varOrganizationChangePlanRequest := _OrganizationChangePlanRequest{}

	err = json.Unmarshal(data, &varOrganizationChangePlanRequest)

	if err != nil {
		return err
	}

	*o = OrganizationChangePlanRequest(varOrganizationChangePlanRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "plan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationChangePlanRequest struct {
	value *OrganizationChangePlanRequest
	isSet bool
}

func (v NullableOrganizationChangePlanRequest) Get() *OrganizationChangePlanRequest {
	return v.value
}

func (v *NullableOrganizationChangePlanRequest) Set(val *OrganizationChangePlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationChangePlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationChangePlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationChangePlanRequest(val *OrganizationChangePlanRequest) *NullableOrganizationChangePlanRequest {
	return &NullableOrganizationChangePlanRequest{value: val, isSet: true}
}

func (v NullableOrganizationChangePlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationChangePlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
