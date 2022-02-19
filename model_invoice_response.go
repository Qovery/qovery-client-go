/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.2
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// InvoiceResponse struct for InvoiceResponse
type InvoiceResponse struct {
	Id           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Status       string    `json:"status"`
	TotalInCents int32     `json:"total_in_cents"`
	Total        float32   `json:"total"`
	CurrencyCode string    `json:"currency_code"`
}

// NewInvoiceResponse instantiates a new InvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceResponse(id string, createdAt time.Time, status string, totalInCents int32, total float32, currencyCode string) *InvoiceResponse {
	this := InvoiceResponse{}
	this.TotalInCents = totalInCents
	this.Total = total
	this.CurrencyCode = currencyCode
	return &this
}

// NewInvoiceResponseWithDefaults instantiates a new InvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceResponseWithDefaults() *InvoiceResponse {
	this := InvoiceResponse{}
	return &this
}

// GetId returns the Id field value
func (o *InvoiceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InvoiceResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InvoiceResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStatus returns the Status field value
func (o *InvoiceResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InvoiceResponse) SetStatus(v string) {
	o.Status = v
}

// GetTotalInCents returns the TotalInCents field value
func (o *InvoiceResponse) GetTotalInCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalInCents
}

// GetTotalInCentsOk returns a tuple with the TotalInCents field value
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetTotalInCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalInCents, true
}

// SetTotalInCents sets field value
func (o *InvoiceResponse) SetTotalInCents(v int32) {
	o.TotalInCents = v
}

// GetTotal returns the Total field value
func (o *InvoiceResponse) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InvoiceResponse) SetTotal(v float32) {
	o.Total = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *InvoiceResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *InvoiceResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *InvoiceResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o InvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["total_in_cents"] = o.TotalInCents
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceResponse struct {
	value *InvoiceResponse
	isSet bool
}

func (v NullableInvoiceResponse) Get() *InvoiceResponse {
	return v.value
}

func (v *NullableInvoiceResponse) Set(val *InvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceResponse(val *InvoiceResponse) *NullableInvoiceResponse {
	return &NullableInvoiceResponse{value: val, isSet: true}
}

func (v NullableInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
