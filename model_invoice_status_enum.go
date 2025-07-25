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
	"fmt"
)

// InvoiceStatusEnum the model 'InvoiceStatusEnum'
type InvoiceStatusEnum string

// List of InvoiceStatusEnum
const (
	INVOICESTATUSENUM_NOT_PAID    InvoiceStatusEnum = "NOT_PAID"
	INVOICESTATUSENUM_PAID        InvoiceStatusEnum = "PAID"
	INVOICESTATUSENUM_PAYMENT_DUE InvoiceStatusEnum = "PAYMENT_DUE"
	INVOICESTATUSENUM_PENDING     InvoiceStatusEnum = "PENDING"
	INVOICESTATUSENUM_POSTED      InvoiceStatusEnum = "POSTED"
	INVOICESTATUSENUM_UNKNOWN     InvoiceStatusEnum = "UNKNOWN"
	INVOICESTATUSENUM_VOIDED      InvoiceStatusEnum = "VOIDED"
)

// All allowed values of InvoiceStatusEnum enum
var AllowedInvoiceStatusEnumEnumValues = []InvoiceStatusEnum{
	"NOT_PAID",
	"PAID",
	"PAYMENT_DUE",
	"PENDING",
	"POSTED",
	"UNKNOWN",
	"VOIDED",
}

func (v *InvoiceStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvoiceStatusEnum(value)
	for _, existing := range AllowedInvoiceStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvoiceStatusEnum", value)
}

// NewInvoiceStatusEnumFromValue returns a pointer to a valid InvoiceStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvoiceStatusEnumFromValue(v string) (*InvoiceStatusEnum, error) {
	ev := InvoiceStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvoiceStatusEnum: valid values are %v", v, AllowedInvoiceStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvoiceStatusEnum) IsValid() bool {
	for _, existing := range AllowedInvoiceStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvoiceStatusEnum value
func (v InvoiceStatusEnum) Ptr() *InvoiceStatusEnum {
	return &v
}

type NullableInvoiceStatusEnum struct {
	value *InvoiceStatusEnum
	isSet bool
}

func (v NullableInvoiceStatusEnum) Get() *InvoiceStatusEnum {
	return v.value
}

func (v *NullableInvoiceStatusEnum) Set(val *InvoiceStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceStatusEnum(val *InvoiceStatusEnum) *NullableInvoiceStatusEnum {
	return &NullableInvoiceStatusEnum{value: val, isSet: true}
}

func (v NullableInvoiceStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
