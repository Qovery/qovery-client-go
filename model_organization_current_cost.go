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
	"time"
)

// checks if the OrganizationCurrentCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCurrentCost{}

// OrganizationCurrentCost struct for OrganizationCurrentCost
type OrganizationCurrentCost struct {
	Plan *PlanEnum `json:"plan,omitempty"`
	// number of days remaining before the end of the trial period
	RemainingTrialDay *int32 `json:"remaining_trial_day,omitempty"`
	// date when the current plan will be renewed
	RenewalAt            NullableTime `json:"renewal_at,omitempty"`
	Cost                 *Cost        `json:"cost,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationCurrentCost OrganizationCurrentCost

// NewOrganizationCurrentCost instantiates a new OrganizationCurrentCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCurrentCost() *OrganizationCurrentCost {
	this := OrganizationCurrentCost{}
	return &this
}

// NewOrganizationCurrentCostWithDefaults instantiates a new OrganizationCurrentCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCurrentCostWithDefaults() *OrganizationCurrentCost {
	this := OrganizationCurrentCost{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetPlan() PlanEnum {
	if o == nil || IsNil(o.Plan) {
		var ret PlanEnum
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetPlanOk() (*PlanEnum, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given PlanEnum and assigns it to the Plan field.
func (o *OrganizationCurrentCost) SetPlan(v PlanEnum) {
	o.Plan = &v
}

// GetRemainingTrialDay returns the RemainingTrialDay field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetRemainingTrialDay() int32 {
	if o == nil || IsNil(o.RemainingTrialDay) {
		var ret int32
		return ret
	}
	return *o.RemainingTrialDay
}

// GetRemainingTrialDayOk returns a tuple with the RemainingTrialDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetRemainingTrialDayOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainingTrialDay) {
		return nil, false
	}
	return o.RemainingTrialDay, true
}

// HasRemainingTrialDay returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasRemainingTrialDay() bool {
	if o != nil && !IsNil(o.RemainingTrialDay) {
		return true
	}

	return false
}

// SetRemainingTrialDay gets a reference to the given int32 and assigns it to the RemainingTrialDay field.
func (o *OrganizationCurrentCost) SetRemainingTrialDay(v int32) {
	o.RemainingTrialDay = &v
}

// GetRenewalAt returns the RenewalAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationCurrentCost) GetRenewalAt() time.Time {
	if o == nil || IsNil(o.RenewalAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RenewalAt.Get()
}

// GetRenewalAtOk returns a tuple with the RenewalAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationCurrentCost) GetRenewalAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalAt.Get(), o.RenewalAt.IsSet()
}

// HasRenewalAt returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasRenewalAt() bool {
	if o != nil && o.RenewalAt.IsSet() {
		return true
	}

	return false
}

// SetRenewalAt gets a reference to the given NullableTime and assigns it to the RenewalAt field.
func (o *OrganizationCurrentCost) SetRenewalAt(v time.Time) {
	o.RenewalAt.Set(&v)
}

// SetRenewalAtNil sets the value for RenewalAt to be an explicit nil
func (o *OrganizationCurrentCost) SetRenewalAtNil() {
	o.RenewalAt.Set(nil)
}

// UnsetRenewalAt ensures that no value is present for RenewalAt, not even an explicit nil
func (o *OrganizationCurrentCost) UnsetRenewalAt() {
	o.RenewalAt.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *OrganizationCurrentCost) GetCost() Cost {
	if o == nil || IsNil(o.Cost) {
		var ret Cost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCurrentCost) GetCostOk() (*Cost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *OrganizationCurrentCost) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given Cost and assigns it to the Cost field.
func (o *OrganizationCurrentCost) SetCost(v Cost) {
	o.Cost = &v
}

func (o OrganizationCurrentCost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCurrentCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.RemainingTrialDay) {
		toSerialize["remaining_trial_day"] = o.RemainingTrialDay
	}
	if o.RenewalAt.IsSet() {
		toSerialize["renewal_at"] = o.RenewalAt.Get()
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationCurrentCost) UnmarshalJSON(data []byte) (err error) {
	varOrganizationCurrentCost := _OrganizationCurrentCost{}

	err = json.Unmarshal(data, &varOrganizationCurrentCost)

	if err != nil {
		return err
	}

	*o = OrganizationCurrentCost(varOrganizationCurrentCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "plan")
		delete(additionalProperties, "remaining_trial_day")
		delete(additionalProperties, "renewal_at")
		delete(additionalProperties, "cost")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationCurrentCost struct {
	value *OrganizationCurrentCost
	isSet bool
}

func (v NullableOrganizationCurrentCost) Get() *OrganizationCurrentCost {
	return v.value
}

func (v *NullableOrganizationCurrentCost) Set(val *OrganizationCurrentCost) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCurrentCost) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCurrentCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCurrentCost(val *OrganizationCurrentCost) *NullableOrganizationCurrentCost {
	return &NullableOrganizationCurrentCost{value: val, isSet: true}
}

func (v NullableOrganizationCurrentCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCurrentCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
