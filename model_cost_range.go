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

// CostRange struct for CostRange
type CostRange struct {
	MinCostInCents *int32   `json:"min_cost_in_cents,omitempty"`
	MinCost        *float32 `json:"min_cost,omitempty"`
	MaxCostInCents *int32   `json:"max_cost_in_cents,omitempty"`
	MaxCost        *float32 `json:"max_cost,omitempty"`
	CurrencyCode   string   `json:"currency_code"`
}

// NewCostRange instantiates a new CostRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostRange(currencyCode string) *CostRange {
	this := CostRange{}
	this.CurrencyCode = currencyCode
	return &this
}

// NewCostRangeWithDefaults instantiates a new CostRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostRangeWithDefaults() *CostRange {
	this := CostRange{}
	return &this
}

// GetMinCostInCents returns the MinCostInCents field value if set, zero value otherwise.
func (o *CostRange) GetMinCostInCents() int32 {
	if o == nil || o.MinCostInCents == nil {
		var ret int32
		return ret
	}
	return *o.MinCostInCents
}

// GetMinCostInCentsOk returns a tuple with the MinCostInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRange) GetMinCostInCentsOk() (*int32, bool) {
	if o == nil || o.MinCostInCents == nil {
		return nil, false
	}
	return o.MinCostInCents, true
}

// HasMinCostInCents returns a boolean if a field has been set.
func (o *CostRange) HasMinCostInCents() bool {
	if o != nil && o.MinCostInCents != nil {
		return true
	}

	return false
}

// SetMinCostInCents gets a reference to the given int32 and assigns it to the MinCostInCents field.
func (o *CostRange) SetMinCostInCents(v int32) {
	o.MinCostInCents = &v
}

// GetMinCost returns the MinCost field value if set, zero value otherwise.
func (o *CostRange) GetMinCost() float32 {
	if o == nil || o.MinCost == nil {
		var ret float32
		return ret
	}
	return *o.MinCost
}

// GetMinCostOk returns a tuple with the MinCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRange) GetMinCostOk() (*float32, bool) {
	if o == nil || o.MinCost == nil {
		return nil, false
	}
	return o.MinCost, true
}

// HasMinCost returns a boolean if a field has been set.
func (o *CostRange) HasMinCost() bool {
	if o != nil && o.MinCost != nil {
		return true
	}

	return false
}

// SetMinCost gets a reference to the given float32 and assigns it to the MinCost field.
func (o *CostRange) SetMinCost(v float32) {
	o.MinCost = &v
}

// GetMaxCostInCents returns the MaxCostInCents field value if set, zero value otherwise.
func (o *CostRange) GetMaxCostInCents() int32 {
	if o == nil || o.MaxCostInCents == nil {
		var ret int32
		return ret
	}
	return *o.MaxCostInCents
}

// GetMaxCostInCentsOk returns a tuple with the MaxCostInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRange) GetMaxCostInCentsOk() (*int32, bool) {
	if o == nil || o.MaxCostInCents == nil {
		return nil, false
	}
	return o.MaxCostInCents, true
}

// HasMaxCostInCents returns a boolean if a field has been set.
func (o *CostRange) HasMaxCostInCents() bool {
	if o != nil && o.MaxCostInCents != nil {
		return true
	}

	return false
}

// SetMaxCostInCents gets a reference to the given int32 and assigns it to the MaxCostInCents field.
func (o *CostRange) SetMaxCostInCents(v int32) {
	o.MaxCostInCents = &v
}

// GetMaxCost returns the MaxCost field value if set, zero value otherwise.
func (o *CostRange) GetMaxCost() float32 {
	if o == nil || o.MaxCost == nil {
		var ret float32
		return ret
	}
	return *o.MaxCost
}

// GetMaxCostOk returns a tuple with the MaxCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRange) GetMaxCostOk() (*float32, bool) {
	if o == nil || o.MaxCost == nil {
		return nil, false
	}
	return o.MaxCost, true
}

// HasMaxCost returns a boolean if a field has been set.
func (o *CostRange) HasMaxCost() bool {
	if o != nil && o.MaxCost != nil {
		return true
	}

	return false
}

// SetMaxCost gets a reference to the given float32 and assigns it to the MaxCost field.
func (o *CostRange) SetMaxCost(v float32) {
	o.MaxCost = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CostRange) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CostRange) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CostRange) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o CostRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinCostInCents != nil {
		toSerialize["min_cost_in_cents"] = o.MinCostInCents
	}
	if o.MinCost != nil {
		toSerialize["min_cost"] = o.MinCost
	}
	if o.MaxCostInCents != nil {
		toSerialize["max_cost_in_cents"] = o.MaxCostInCents
	}
	if o.MaxCost != nil {
		toSerialize["max_cost"] = o.MaxCost
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableCostRange struct {
	value *CostRange
	isSet bool
}

func (v NullableCostRange) Get() *CostRange {
	return v.value
}

func (v *NullableCostRange) Set(val *CostRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCostRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCostRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostRange(val *CostRange) *NullableCostRange {
	return &NullableCostRange{value: val, isSet: true}
}

func (v NullableCostRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
