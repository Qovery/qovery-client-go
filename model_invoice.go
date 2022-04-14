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
	"time"
)

// Invoice struct for Invoice
type Invoice struct {
	TotalInCents int32             `json:"total_in_cents"`
	Total        float32           `json:"total"`
	CurrencyCode string            `json:"currency_code"`
	Id           string            `json:"id"`
	CreatedAt    time.Time         `json:"created_at"`
	Status       InvoiceStatusEnum `json:"status"`
}

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice(totalInCents int32, total float32, currencyCode string, id string, createdAt time.Time, status InvoiceStatusEnum) *Invoice {
	this := Invoice{}
	this.TotalInCents = totalInCents
	this.Total = total
	this.CurrencyCode = currencyCode
	this.Id = id
	this.CreatedAt = createdAt
	this.Status = status
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetTotalInCents returns the TotalInCents field value
func (o *Invoice) GetTotalInCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalInCents
}

// GetTotalInCentsOk returns a tuple with the TotalInCents field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetTotalInCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalInCents, true
}

// SetTotalInCents sets field value
func (o *Invoice) SetTotalInCents(v int32) {
	o.TotalInCents = v
}

// GetTotal returns the Total field value
func (o *Invoice) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Invoice) SetTotal(v float32) {
	o.Total = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *Invoice) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *Invoice) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetId returns the Id field value
func (o *Invoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Invoice) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Invoice) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Invoice) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStatus returns the Status field value
func (o *Invoice) GetStatus() InvoiceStatusEnum {
	if o == nil {
		var ret InvoiceStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusOk() (*InvoiceStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Invoice) SetStatus(v InvoiceStatusEnum) {
	o.Status = v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_in_cents"] = o.TotalInCents
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}