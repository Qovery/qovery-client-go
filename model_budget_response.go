/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// BudgetResponse struct for BudgetResponse
type BudgetResponse struct {
	TotalInCents *int32   `json:"total_in_cents,omitempty"`
	Total        *float32 `json:"total,omitempty"`
	CurrencyCode *string  `json:"currency_code,omitempty"`
}

// NewBudgetResponse instantiates a new BudgetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetResponse() *BudgetResponse {
	this := BudgetResponse{}
	return &this
}

// NewBudgetResponseWithDefaults instantiates a new BudgetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetResponseWithDefaults() *BudgetResponse {
	this := BudgetResponse{}
	return &this
}

// GetTotalInCents returns the TotalInCents field value if set, zero value otherwise.
func (o *BudgetResponse) GetTotalInCents() int32 {
	if o == nil || o.TotalInCents == nil {
		var ret int32
		return ret
	}
	return *o.TotalInCents
}

// GetTotalInCentsOk returns a tuple with the TotalInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetResponse) GetTotalInCentsOk() (*int32, bool) {
	if o == nil || o.TotalInCents == nil {
		return nil, false
	}
	return o.TotalInCents, true
}

// HasTotalInCents returns a boolean if a field has been set.
func (o *BudgetResponse) HasTotalInCents() bool {
	if o != nil && o.TotalInCents != nil {
		return true
	}

	return false
}

// SetTotalInCents gets a reference to the given int32 and assigns it to the TotalInCents field.
func (o *BudgetResponse) SetTotalInCents(v int32) {
	o.TotalInCents = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BudgetResponse) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetResponse) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BudgetResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *BudgetResponse) SetTotal(v float32) {
	o.Total = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BudgetResponse) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BudgetResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BudgetResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o BudgetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalInCents != nil {
		toSerialize["total_in_cents"] = o.TotalInCents
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableBudgetResponse struct {
	value *BudgetResponse
	isSet bool
}

func (v NullableBudgetResponse) Get() *BudgetResponse {
	return v.value
}

func (v *NullableBudgetResponse) Set(val *BudgetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetResponse(val *BudgetResponse) *NullableBudgetResponse {
	return &NullableBudgetResponse{value: val, isSet: true}
}

func (v NullableBudgetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
