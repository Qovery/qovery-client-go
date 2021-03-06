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

// BillingInfo struct for BillingInfo
type BillingInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// email used for billing, and to receive all invoices by email
	Email   string `json:"email"`
	Address string `json:"address"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	// only for US
	State *string `json:"state,omitempty"`
	// ISO code of the country
	CountryCode string `json:"country_code"`
	// name of the company to bill
	Company   *string `json:"company,omitempty"`
	VatNumber *string `json:"vat_number,omitempty"`
}

// NewBillingInfo instantiates a new BillingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfo(firstName string, lastName string, email string, address string, city string, zip string, countryCode string) *BillingInfo {
	this := BillingInfo{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Address = address
	this.City = city
	this.Zip = zip
	this.CountryCode = countryCode
	return &this
}

// NewBillingInfoWithDefaults instantiates a new BillingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoWithDefaults() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *BillingInfo) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *BillingInfo) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *BillingInfo) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *BillingInfo) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *BillingInfo) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BillingInfo) SetEmail(v string) {
	o.Email = v
}

// GetAddress returns the Address field value
func (o *BillingInfo) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BillingInfo) SetAddress(v string) {
	o.Address = v
}

// GetCity returns the City field value
func (o *BillingInfo) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *BillingInfo) SetCity(v string) {
	o.City = v
}

// GetZip returns the Zip field value
func (o *BillingInfo) GetZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zip
}

// GetZipOk returns a tuple with the Zip field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zip, true
}

// SetZip sets field value
func (o *BillingInfo) SetZip(v string) {
	o.Zip = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BillingInfo) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BillingInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BillingInfo) SetState(v string) {
	o.State = &v
}

// GetCountryCode returns the CountryCode field value
func (o *BillingInfo) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *BillingInfo) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BillingInfo) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BillingInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *BillingInfo) SetCompany(v string) {
	o.Company = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *BillingInfo) GetVatNumber() string {
	if o == nil || o.VatNumber == nil {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetVatNumberOk() (*string, bool) {
	if o == nil || o.VatNumber == nil {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *BillingInfo) HasVatNumber() bool {
	if o != nil && o.VatNumber != nil {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *BillingInfo) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o BillingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["zip"] = o.Zip
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.VatNumber != nil {
		toSerialize["vat_number"] = o.VatNumber
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfo struct {
	value *BillingInfo
	isSet bool
}

func (v NullableBillingInfo) Get() *BillingInfo {
	return v.value
}

func (v *NullableBillingInfo) Set(val *BillingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfo(val *BillingInfo) *NullableBillingInfo {
	return &NullableBillingInfo{value: val, isSet: true}
}

func (v NullableBillingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
